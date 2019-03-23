package main

import "fmt"

func barChartTemplate(labels string, values string, valuesLabel string) string {
	html := fmt.Sprintf(`
<html>
<head>
</head>
<body>
	<canvas id="myChart"></canvas>

	<script src="https://cdn.jsdelivr.net/npm/chart.js@2.8.0"></script>

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
							beginAtZero: true
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
