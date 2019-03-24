package main

import "fmt"

func barChartTemplate(labels string, values string, valuesLabel string) string {
	html := fmt.Sprintf(`
<html>
<head>
</head>
<body>
	<canvas id="myChart"></canvas>

	<script src="./js/Chart.bundle.js"></script>
	<script src="./js/chartjs-plugin-colorschemes.js"></script>

	<script>
		var ctx = document.getElementById('myChart');
		var myChart = new Chart(ctx, {
			type: 'bar',
			data: {
				labels: [%v],
				datasets: [{
					label: %v,
					data: [%v],
				}]
			},
			options: {
				scales: {
					yAxes: [{
						ticks: {
							beginAtZero: false
						}
					}]
				}
			}
		});
</script>
</body>
</html>
`, labels, valuesLabel, values)
	return html
}

func pieChartTemplate(labels string, values string, valuesLabel string) string {
	html := fmt.Sprintf(`
<html>
<head>
</head>
<body>
	<canvas id="myChart"></canvas>

	<script src="./js/Chart.bundle.js"></script>
	<script src="./js/chartjs-plugin-colorschemes.js"></script>

	<script>
		var ctx = document.getElementById('myChart');
		var myChart = new Chart(ctx, {
			type: 'pie',
			data: {
				labels: [%v],
				datasets: [{
					label: %v,
					data: [%v],
				}]
			},
			options: {
				plugins: {
					colorschemes: {
					  scheme: 'brewer.Paired12'
					}
				  }
			}
		});
</script>
</body>
</html>
`, labels, valuesLabel, values)
	return html
}
