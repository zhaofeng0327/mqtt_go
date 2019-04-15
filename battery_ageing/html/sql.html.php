<!DOCTYPE html>
<meta charset="UTF-8">
<html>
<head>
<style>
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

<table>

<tr>
<th>timestamp</th>
<th>slotnum</th>
<th>batterysn</th>
<th>voltage</th>
</tr>

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
		$tables[] = $row[0];
	}
}

if (count($tables) === 0) {
	die("no tables found");
}

rsort($tables);

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
			echo "<td>" . $row["xvoltage"] . "</td>";
			echo "</tr>";
		}
	} else {
		echo "0 results";
	}
}
mysqli_close($conn);
?>

<form action="sql.php" method="post" target="formDestination">
  <input type="submit" value="click on me!">
</form>


</body>
</html>
