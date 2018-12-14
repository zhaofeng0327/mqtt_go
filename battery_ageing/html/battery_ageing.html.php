<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="refresh" content="5">
  <link rel="stylesheet" href="css/style.css">

<style>

.item1 { grid-area: header; }
.item2 { grid-area: menu; }
.item3 { grid-area: main; }
.item4 { grid-area: right; }
.item5 { grid-area: footer; }

.grid-container {
display: grid;
	 grid-template:
		 'header header header header header header'
		 'main main main main main main'
		 'footer footer footer footer footer footer';
	 grid-gap: 10px;
	 background-color: #2196F3;
padding: 10px;
}



.grid-container>div {
	text-color: #000000;
	background-color: rgba(255, 255, 255, 0.8);
	text-align: center;
        padding:20px 0;
	font-size: 30px;
}

.button {
    background-color: #4CAF50;
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    margin: 4px 2px;
    cursor: pointer;
}

input[type=text] {
  width: 25%;
  text-align: center;
  padding: 12px 20px;
  margin: 8px 0;
  box-sizing: border-box;
  border: none;
  background-color: #ADFF2F;
  color: white;
  font-size:24px;
}

input[type="radio"] {
  margin-left: 50px;
}

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
</style>
</head>


<body>

<div class="grid-container">
<div class="item3">  

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
				echo "<tr>";
				echo "<td>" . $row["timestamp"] . "</td>";
				echo "<td>" . $row["slotnum"] . "</td>";
				echo "<td>" . $row["batterysn"] . "</td>";
				echo "<td>" . $row["voltage"] . "</td>";
				echo "<td>" . $row["current"] . "</td>";
				echo "<td>" . $row["temprature"] . "</td>";
				echo "<td>" . $row["elapsed"] . "</td>";
				echo "<td>" . $row["xradio"] . "</td>";
				echo "</tr>";
			}
		} else {
			//	echo "0 results";
		}
	}
	echo "</table><br>";


	mysqli_close($conn);

?>


<form name = form1 action = "/cgi-bin/battery_ageing.cgi" method = "POST" target="formDestination">

    <label for="fname">柜机SN</label> 
    <input type="text" id="sn" name="device_sn" readonly value="<?php echo htmlentities($_GET["device_sn"]); ?>"/> 

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
	}
	</script>


    <h3 >卡槽号</h3>

    <input type="radio" class="option-input radio" name="slot_num" value = "1" onclick="reload()" <?php echo $_GET['slot_num'] == "1" ? "checked" : "";?>/> 1
    <input type="radio" class="option-input radio" name="slot_num" value = "2" onclick="reload()" <?php echo $_GET['slot_num'] == "2" ? "checked" : "";?>/> 2
    <input type="radio" class="option-input radio" name="slot_num" value = "3" onclick="reload()" <?php echo $_GET['slot_num'] == "3" ? "checked" : "";?>/> 3
    <br> <br>

    <input type="radio" class="option-input radio" name="slot_num" value = "4" onclick="reload()" <?php echo $_GET['slot_num'] == "4" ? "checked" : "";?>/> 4
    <input type="radio" class="option-input radio" name="slot_num" value = "5" onclick="reload()" <?php echo $_GET['slot_num'] == "5" ? "checked" : "";?>/> 5
    <input type="radio" class="option-input radio" name="slot_num" value = "6" onclick="reload()" <?php echo $_GET['slot_num'] == "6" ? "checked" : "";?>/> 6
    <br><br>

    <!-- input type="radio" class="option-input radio" name="slot_num" value = "0"/ 全部-->

    <h3 >放电电流</h3>

    <input type="radio" class="option-input radio" name="current_level" value = "0" onclick="reload()" <?php echo $_GET['current_level'] == "0" ? "checked" : "";?>/> 低
    <input type="radio" class="option-input radio" name="current_level" value = "1" onclick="reload()" <?php echo $_GET['current_level'] == "1" ? "checked" : "";?>/> 中
    <input type="radio" class="option-input radio" name="current_level" value = "2" onclick="reload()" <?php echo $_GET['current_level'] == "2" ? "checked" : "";?>/> 高

    <br><br>

    <h3 >操作</h3>

    <label>
    <input type="radio" class="option-input radio" name="option" value = "1" onclick="reload()" <?php echo $_GET['option'] == "1" ? "checked" : "";?>/> 开始放电
    </label>

    <label>
    <input type="radio" class="option-input radio" name="option" value = "2" onclick="reload()" <?php echo $_GET['option'] == "2" ? "checked" : "";?>/> 停止放电
    </label>

<br><br>
<input type = "submit" style = "margin:20px;" value = "提交" class = "button">
</form>
</div>

<div class="item5">
<iframe name="formDestination" width="1000" height="40" align = center frameborder= no></iframe>
</div>


</div>

</body>
</html>
