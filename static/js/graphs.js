//setup
const data = {
  labels: ['January', 'February', 'March', 'April', 'May', 'June'],
  datasets: [{
      label: 'Half Year Earnings in millions',
      data: [12, 19, 3, 5, 2, 3],
      backgroundColor: [
          'rgba(255, 99, 132, 0.2)',
          'rgba(54, 162, 235, 0.2)',
          'rgba(255, 206, 86, 0.2)',
          'rgba(75, 192, 192, 0.2)',
          'rgba(153, 102, 255, 0.2)',
          'rgba(255, 159, 64, 0.2)'
      ],
      borderColor: [
          'rgba(255, 99, 132, 1)',
          'rgba(54, 162, 235, 1)',
          'rgba(255, 206, 86, 1)',
          'rgba(75, 192, 192, 1)',
          'rgba(153, 102, 255, 1)',
          'rgba(255, 159, 64, 1)'
      ],
      borderWidth: 1
  }]
}
// config
const config = {
  type: 'bar',
  data,
  options: {
      scales: {
          y: {
              beginAtZero: true
          }
      }
  }
}
//render block
const myChart = new Chart(document.getElementById('myAreaChart').getContext('2d'),config);
const dataPie = {
  labels: ['Walk in Clients', 'Referal Client ', ' Online Booking '],
  datasets: [{
      label: 'Total Earnngs of each vehicle',
      data: [12, 19, 10],
      backgroundColor: [
          'rgba(255, 99, 132, 0.2)',
          'rgba(54, 162, 235, 0.2)',
          'rgba(255, 206, 86, 0.2)'
      ],
      borderColor: [
          'rgba(255, 99, 132, 1)',
          'rgba(54, 162, 235, 1)',
          'rgba(255, 206, 86, 1)'
      ],
      borderWidth: 1
  }]
};
const configPie = {
  type: 'pie',
  data:dataPie,
  options: {
    responsive: true,
    maintainAspectRatio : false
  }
}
const myPieChart = new Chart(document.getElementById('myPieChart'),configPie);
