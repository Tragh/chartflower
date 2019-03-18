package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	csvFiles := chooseCsvFiles()
	database := getDatabase()
	table := getTable(database, csvFiles)
	chooseColumns(table)
}

func chooseColumns(table *table) {
	println()
	for i, columnName := range table.csvColumnNames {
		fmt.Printf("%v. %v\n", i, columnName)
	}
	print("\nChoose columns:")
	input := getConsoleText()
	choices := strings.Fields(input)
	fmt.Println(choices)
}

func getTable(database *sql.DB, filenames []string) *table {
	var table *table
	if len(filenames) == 1 {
		tableName := strings.TrimSuffix(filenames[0], ".csv")
		table = createTable(tableName, database, filenames[0])
	} else if len(filenames) == 3 {
		fmt.Println("No support for 3+ tables yet")
		os.Exit(0)
	} else {
		table = createJoinedTable(database, filenames)
	}
	return table
}

func createJoinedTable(database *sql.DB, filenames []string) *table {
	var tables []*table
	for _, filename := range filenames {
		tableName := strings.TrimSuffix(filename, ".csv")
		table := createTable(tableName, database, filename)
		tables = append(tables, table)
	}
	joinedTable := new(table)
	joinedTable.database = database
	joinedTable.sqlTableName = "Joined"
	for _, table := range tables {
		joinedTable.csvColumnNames = append(joinedTable.csvColumnNames, table.csvColumnNames...)
		joinedTable.sqlColumnNames = append(joinedTable.sqlColumnNames, table.sqlColumnNames...)
	}
	joinedTable.numberOfColumns = len(joinedTable.sqlColumnNames)

	statementString := "CREATE TABLE IF NOT EXISTS " + joinedTable.sqlTableName + " AS SELECT * FROM " + tables[0].sqlTableName + " JOIN " + tables[1].sqlTableName + " ON " + tables[0].sqlTableName + "." + tables[0].sqlColumnNames[0] + " = " + tables[1].sqlTableName + "." + tables[1].sqlColumnNames[0]
	statement, error := joinedTable.database.Prepare(statementString)
	if error != nil {
		fmt.Println(error)
	}
	statement.Exec()
	statement.Close()
	return joinedTable
}

func createTable(tableName string, database *sql.DB, filename string) *table {
	table := new(table)
	table.database = database
	table.sqlTableName = tableName
	csvFilename := "./csv/" + filename
	file, _ := os.Open(csvFilename)
	defer file.Close()
	reader := csv.NewReader(file)
	table.csvColumnNames, _ = reader.Read()
	table.sqlColumnNames = []string{}
	for i := range table.csvColumnNames {
		table.sqlColumnNames = append(table.sqlColumnNames, table.sqlTableName+"_"+strconv.Itoa(i))
	}
	table.numberOfColumns = len(table.sqlColumnNames)

	tableLayout := ""
	for i := 0; i < table.numberOfColumns-1; i++ {
		tableLayout += table.sqlColumnNames[i] + " TEXT, "
	}
	tableLayout += table.sqlColumnNames[table.numberOfColumns-1] + " TEXT"
	statement, _ := table.database.Prepare("CREATE TABLE IF NOT EXISTS " + table.sqlTableName + " (" + tableLayout + ")")
	statement.Exec()
	statement.Close()
	statement, _ = table.database.Prepare("INSERT INTO " + table.sqlTableName + " VALUES (" + getCommaSeparatedString("?", table.numberOfColumns) + ")")
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		statement.Exec(convertStringsToInterfaces(row)...)
	}
	statement.Close()

	return table
}

func convertStringsToInterfaces(strings []string) []interface{} {
	interfaces := make([]interface{}, len(strings))
	for i, value := range strings {
		interfaces[i] = value
	}
	return interfaces
}

func getCommaSeparatedString(value string, numberOfColumns int) string {
	commaSeparatedString := ""
	if numberOfColumns <= 0 {
		return commaSeparatedString
	}
	for i := 0; i < numberOfColumns-1; i++ {
		commaSeparatedString += value + ","
	}
	commaSeparatedString += value
	return commaSeparatedString
}

type table struct {
	database        *sql.DB
	sqlTableName    string
	csvColumnNames  []string
	sqlColumnNames  []string
	numberOfColumns int
}

func getDatabase() *sql.DB {
	database, _ := sql.Open("sqlite3", ":memory:")
	return database
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func convertCSVToArray(filename string) [][]string {
	filename = "./csv/" + filename
	file, _ := os.Open(filename)
	defer file.Close()
	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()
	return rows
}

func chooseCsvFiles() []string {
	csvFiles := getCsvFilesinFolder()
	println()
	for index, csvFiles := range csvFiles {
		csvTrimmed := strings.TrimSuffix(csvFiles, ".csv")
		fmt.Printf("%v. %v\n", index, csvTrimmed)
	}
	print("\nChoose CSVs:")
	input := getConsoleText()
	choices := strings.Fields(input)
	var chosenCsvFiles []string
	for _, choice := range choices {
		choiceIndex, _ := strconv.Atoi(choice)
		chosenCsvFiles = append(chosenCsvFiles, csvFiles[choiceIndex])
	}
	return chosenCsvFiles
}

func getCsvFilesinFolder() []string {
	var csvFiles []string
	filesInFolder, _ := ioutil.ReadDir("./csv/")
	for _, fileInFolder := range filesInFolder {
		if strings.HasSuffix(fileInFolder.Name(), "csv") {
			csvFiles = append(csvFiles, fileInFolder.Name())
		}
	}
	return csvFiles
}

func getConsoleText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = removeNewLine(text)
	typeQToQuit(text)
	return text
}

func removeNewLine(input string) string {
	switch os := runtime.GOOS; os {
	case "windows":
		newLineRemoved := strings.TrimSuffix(input, "\r\n")
		return newLineRemoved
	default:
		newLineRemoved := strings.TrimSuffix(input, "\n")
		return newLineRemoved
	}
}

func typeQToQuit(input string) {
	if input == "q" {
		os.Exit(0)
	}
}

func (table table) printTable(columns ...int) {
	isColumnPrinted := make(map[int]bool)
	if columns != nil {
		for _, c := range columns {
			isColumnPrinted[c] = true
		}
	} else {
		for i := range table.csvColumnNames {
			isColumnPrinted[i] = true
		}
	}
	fmt.Println("## Column Names ##")

	for i, columnName := range table.csvColumnNames {
		if isColumnPrinted[i] {
			fmt.Print(columnName, ", ")
		}
	}
	fmt.Println()
	fmt.Println("## Rows ##")
	rows, _ := table.database.Query("SELECT * FROM " + table.sqlTableName)
	rowData := make([]interface{}, table.numberOfColumns)
	rowDataPointers := make([]interface{}, table.numberOfColumns)

	for i := range rowData {
		rowDataPointers[i] = &rowData[i]
	}

	for rows.Next() {
		rows.Scan(rowDataPointers...)
		for i, data := range rowData {
			if isColumnPrinted[i] {
				//fmt.Print(fmt.Sprint(data)+", ") //This also works
				if dataAsString, ok := data.(string); ok {
					fmt.Print(dataAsString + ", ")
				} else {
					fmt.Print(" , ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
