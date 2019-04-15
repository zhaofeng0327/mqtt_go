<?php
$servername = "localhost";
$username = "phpmyadmin";
$password = "PHP@password123";
$dbname = "battery_ageing";

$conn = mysqli_connect($servername, $username, $password, $dbname);

if (!$conn) {
    die("Connection failed: " . mysqli_connect_error());
}

//echo '<table cellpadding="1" cellspacing="20" class="db-table">';
echo '<table>';
echo '<tr> <th>timestamp</th> <th>slotnum</th> <th>batterysn</th> <th>voltage</th> </tr>';
//$sql = "SELECT * FROM A484MB1358_20181208 where timestamp=(select MAX(timestamp) from A484MB1358_20181208)";
for ($x = 0; $x < 6; $x++) {
	$sql = "SELECT * FROM A484MB1358_20181208 where slotnum=$x order by timestamp DESC limit 1";
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
		//echo "0 results";
	}
}
echo '</table>';
mysqli_close($conn);
?>
