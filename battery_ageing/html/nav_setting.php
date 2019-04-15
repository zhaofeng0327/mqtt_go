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

    /* table */
    table {
        border-collapse: collapse;
        width: 100%;
    }
    
    th, td {
        text-align: left;
        padding: 8px;
    }
    
    tr:nth-child(even){background-color: #f2f2f2}
    
    th {
        background-color: #4CAF50;
        color: white;
    }

    input:focus {
        outline:none;
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
      <p><a href=<?php print "nav_chart.php?device_sn=" . $_GET["device_sn"] ?>>历史数据</a></p>
      
    </div>

    <div class="col-sm-8 text-left"> 

    <?php
	$servername = "localhost";
	$username = "phpmyadmin";
	$password = "PHP@password123";
	$dbname = "battery_ageing";

	$conn = mysqli_connect($servername, $username, $password, $dbname);

	if (!$conn) {
	    die("Connection failed: " . mysqli_connect_error());
	}

	$sql = "SHOW TABLES";
	$result = mysqli_query($conn, $sql);

	$tables = array();
	if (mysqli_num_rows($result) > 0) {
		while ($row = mysqli_fetch_row($result)) {
			if (strpos($row[0], $_GET["device_sn"]) === 0) {
				$tables[] = $row[0];
			}
		}
	}

	if (count($tables) === 0) {
		mysqli_close($conn);
		die("no tables found");
	}

	rsort($tables);

	echo "<table>";
	echo "<tr>";
	echo "<th>时间</th>";
	echo "<th>卡槽ID</th>";
	echo "<th>电池SN</th>";
	echo "<th>外部电压</th>";
	echo "<th>放电电流</th>";
	echo "<th>温度</th>";
	echo "<th>放电累计</th>";
	echo "<th>剩余电量</th>";
	echo "</tr>";

	for ($x = 0; $x < 6; $x++) {
		$sql = "SELECT * FROM $tables[0] where slotnum=$x order by timestamp DESC limit 1";
		$result = mysqli_query($conn, $sql);

		if (mysqli_num_rows($result) > 0) {
			// output data of each row
			while($row = mysqli_fetch_assoc($result)) {
				$t = strtotime($row["timestamp"]) + 8*3600;
				if ($t + 120 >= time()) {
					$d = gmdate("Y-m-d H:i:s", $t);
					echo "<tr>";
					echo "<td>" . $d . "</td>";
					echo "<td>" . $row["slotnum"] . "</td>";
					echo "<td>" . $row["batterysn"] . "</td>";
					echo "<td>" . $row["voltage"] . "</td>";
					echo "<td>" . $row["current"] . "</td>";
					echo "<td>" . $row["temprature"] . "</td>";
					echo "<td>" . $row["elapsed"] . "</td>";
					echo "<td>" . $row["xradio"] . "</td>";
					echo "</tr>";
				}
			}
		} else {
			//	echo "0 results";
		}
	}
	echo "</table><br>";


	mysqli_close($conn);

	?>

	<form name = form1 action = "/cgi-bin/battery_ageing.cgi" method = "POST" target="formDestination">

	柜机SN <input type="text" id="sn" name="device_sn" readonly value="<?php echo htmlentities($_GET["device_sn"]); ?>"/>

	<script>
	function reload() {
		var v_sn = document.form1.sn.value;
		var v_slot_num = "1";
		var v_current_level = "0";
		var v_option = "1";

		for(var i = 0; i < document.form1.slot_num.length; i++) {
			if(document.form1.slot_num[i].checked) {
				v_slot_num = document.form1.slot_num[i].value; 
				break;
			}
		}

		for(var i = 0; i < document.form1.current_level.length; i++) {
			if(document.form1.current_level[i].checked) {
				v_current_level = document.form1.current_level[i].value; 
				break;
			}
		}

		for(var i = 0; i < document.form1.option.length; i++) {
			if(document.form1.option[i].checked) {
				v_option = document.form1.option[i].value; 
				break;
			}
		}

		self.location=window.location.pathname + '?device_sn=' + v_sn + '&slot_num=' + v_slot_num + '&current_level=' + v_current_level + '&option=' + v_option;

		return false;
	}
	</script>


	<hr>
	<h4>卡槽号</h4>

	<div class="btn-group btn-group-toggle" data-toggle="buttons">

	<label class="btn btn-secondary <?php echo $_GET['slot_num'] == "1" ? "active" : "";?>">
	<input type="radio" name="slot_num" value = "1" autocomplete=off onclick="reload()" /> 1
	</label>

	<label class="btn btn-secondary <?php echo $_GET['slot_num'] == "2" ? "active" : "";?>">
	<input type="radio" name="slot_num" value = "2" autocomplete=off onclick="reload()" /> 2
	</label>

	<label class="btn btn-secondary <?php echo $_GET['slot_num'] == "3" ? "active" : "";?>">
	<input type="radio" name="slot_num" value = "3" autocomplete=off onclick="reload()" /> 3
	</label>

	<label class="btn btn-secondary <?php echo $_GET['slot_num'] == "4" ? "active" : "";?>">
	<input type="radio" name="slot_num" value = "4" autocomplete=off onclick="reload()" /> 4
	</label>

	<label class="btn btn-secondary <?php echo $_GET['slot_num'] == "5" ? "active" : "";?>">
	<input type="radio" name="slot_num" value = "5" autocomplete=off onclick="reload()" /> 5
	</label>

	<label class="btn btn-secondary <?php echo $_GET['slot_num'] == "6" ? "active" : "";?>">
	<input type="radio" name="slot_num" value = "6" autocomplete=off onclick="reload()" /> 6
	</label>

	</div>

	<br><br>

	<!-- input type="radio" class="option-input radio" name="slot_num" value = "0"/ 全部-->

	<h4 >放电电流</h4>

	<div class="btn-group btn-group-toggle" data-toggle="buttons">

	<label class="btn btn-secondary <?php echo $_GET['current_level'] == "0" ? "active" : "";?>">
	<input type="radio" name="current_level" value = "0" autocomplete=off onclick="reload()" /> 低
	</label>

	<label class="btn btn-secondary <?php echo $_GET['current_level'] == "1" ? "active" : "";?>">
	<input type="radio" name="current_level" value = "1" autocomplete=off onclick="reload()" /> 中
	</label>

	<label class="btn btn-secondary <?php echo $_GET['current_level'] == "2" ? "active" : "";?>">
	<input type="radio" name="current_level" value = "2" autocomplete=off onclick="reload()" /> 高
	</label>
	</div>

	<br><br>

	<h4 >操作</h4>

	<div class="btn-group btn-group-toggle" data-toggle="buttons">

	<label class="btn btn-secondary <?php echo $_GET['option'] == "1" ? "active" : "";?>">
	<input type="radio" name="option" value = "1" autocomplete=off onclick="reload()" /> 开始放电
	</label>

	<label class="btn btn-secondary" <?php echo $_GET['option'] == "1" ? "active" : "";?>>
	<input type="radio" name="option" value = "2" autocomplete=off onclick="reload()"/> 停止放电
	</label>

	</div>

	<hr>
	<input type = "submit" style = "margin:20px;" value = "提交" class = "btn">
	</form>

	<iframe name="formDestination" width="1000" height="40" align = center frameborder= no></iframe>

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
