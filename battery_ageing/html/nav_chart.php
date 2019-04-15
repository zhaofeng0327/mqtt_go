<!DOCTYPE html>
<html lang="en">
<head>
  <title>Bootstrap Example</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <link rel="stylesheet" href="bootstrap/4.0.0/css/bootstrap.min.css">
  <script src="ajax/libs/jquery/3.3.1/jquery.min.js"></script>
  <script src="bootstrap/4.0.0/js/bootstrap.min.js"></script>

  <style>
    /* Remove the navbar's default margin-bottom and rounded borders */ 
    .navbar {
      margin-bottom: 0;
      border-radius: 0;
    }
    
    /* Set height of the grid so .sidenav can be 100% (adjust as needed) */
    .row.content {height: 100%}
    
    /* Set gray background color and 100% height */
    .sidenav {
      padding-top: 20px;
      background-color: #f1f1f1;
      height: 100%;
    }
    
    /* Set black background color, white text and some padding */
    footer {
      position:absolute;
      bottom:0;
      background-color: #555;
      color: white;
      padding: 15px;
    }
    
    /* On small screens, set height to 'auto' for sidenav and grid */
    @media screen and (max-width: 767px) {
      .sidenav {
        height: auto;
        padding: 15px;
      }
      .row.content {height:auto;} 
    }
  </style>
</head>
<body>

<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
  <a class="navbar-brand" href="#">
    <img src="bootstrap/4.0.0/brand/jiedian.png" width="30" height="30" class="d-inline-block align-top" alt="">
    街电
  </a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>

  <div class="collapse navbar-collapse" id="navbarSupportedContent">
    <ul class="navbar-nav mr-auto">
      <li class="nav-item active">
        <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">Link</a>
      </li>
      <li class="nav-item dropdown">
        <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
          Dropdown
        </a>
        <div class="dropdown-menu" aria-labelledby="navbarDropdown">
          <a class="dropdown-item" href="#">Action</a>
          <a class="dropdown-item" href="#">Another action</a>
          <div class="dropdown-divider"></div>
          <a class="dropdown-item" href="#">Something else here</a>
        </div>
      </li>
      <li class="nav-item">
        <a class="nav-link disabled" href="#">Disabled</a>
      </li>
    </ul>
    <form class="form-inline my-2 my-lg-0">
      <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search">
      <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
    </form>
  </div>
</nav>  
  
<div class="container-fluid text-center">    
  <div class="row content">
    <div class="col-sm-2 sidenav">
      <p><a href="#">Link</a></p>
      <p><a href="#">Link</a></p>
      <p><a href="#">Link</a></p>
    </div>
    <div class="col-sm-8 text-left"> 
    <script>
      function drawVoltage() {

	var start_time = document.getElementById("start_time").value; 
	var end_time = document.getElementById("end_time").value; 
	var slot_num = document.getElementById("slot_num").value; 
	alert(start_time);

	var jsonData = $.ajax({
          url: "getVoltage.php",
          dataType: "json",
          async: false,
          type: "post",
          data: {"slot_num" : slot_num, "start_time" : start_time, "end_time" : end_time}
          }).responseText;
	var data = new google.visualization.DataTable(jsonData);

        var options = {
          title: '电压',
          hAxis: {title: '时间',  titleTextStyle: {color: '#333'}},
          vAxis: {minValue: 0}
        };

        var chart = new google.visualization.AreaChart(document.getElementById('chart_voltage'));
        chart.draw(data, options);

	//chart.draw(data, {width: 400, height: 240});

      }

      function drawCurrent() {

	var start_time = document.getElementById("start_time").value; 
	var end_time = document.getElementById("end_time").value; 
	var slot_num = document.getElementById("slot_num").value; 

	var jsonData = $.ajax({
          url: "getCurrent.php",
          dataType: "json",
          async: false,
          type: "post",
          data: {"slot_num" : slot_num, "start_time" : start_time, "end_time" : end_time}
          }).responseText;
	var data = new google.visualization.DataTable(jsonData);

        var options = {
          title: '电流',
          hAxis: {title: '时间',  titleTextStyle: {color: '#333'}},
          vAxis: {minValue: 0}
        };

        var chart = new google.visualization.AreaChart(document.getElementById('chart_current'));
        chart.draw(data, options);

	//chart.draw(data, {width: 400, height: 240});

      }

      function drawCharts() {
	drawVoltage();
	drawCurrent();
      }

    </script>


    <label>开始时间</label>
<div class="container">
    <div class="row">
        <div class='col-sm-6'>
            <div class="form-group">
                <div class='input-group date' id='datetimepicker1'>
                    <input type='text' class="form-control" />
                    <span class="input-group-addon">
                        <span class="glyphicon glyphicon-calendar"></span>
                    </span>
                </div>
            </div>
        </div>
        <script type="text/javascript">
            $(function () {
                $('#datetimepicker1').datetimepicker({
                   //inline: true,
                   //sideBySide: true
                });
            });
        </script>
    </div>
</div>





    <input type="datetime-local" id="start_time">

    <label>结束时间</label>
    <input type="datetime-local" id="end_time">


    <label>卡槽</label>
    <select id="slot_num">
    <option value="1">1</option>
    <option value="2">2</option>
    <option value="3">3</option>
    <option value="4">4</option>
    <option value="5">5</option>
    <option value="6">6</option>
    </select>

    <hr>
    <input type="button" class="btn" name="drawchart" value="提交" onclick="drawCharts()" />

    <div id="chart_voltage" style="width: 100%; height: 500px;"></div>
    <div id="chart_current" style="width: 100%; height: 500px;"></div>

    </div>

    <div class="col-sm-2 sidenav">
      <div class="well">
        <p>ADS</p>
      </div>
      <div class="well">
        <p>ADS</p>
      </div>
    </div>
  </div>
</div>

<footer class="container-fluid text-center">
  <p>© 2019 Copyright: www.jiediankeji.com </p>
</footer>

</body>
</html>
