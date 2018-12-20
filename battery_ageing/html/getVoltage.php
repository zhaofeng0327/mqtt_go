<?php
/*
            cols: [{id: 'a', label: 'NEW A', type: 'string'},
                   {id: 'b', label: 'B-label', type: 'number'},
                   {id: 'c', label: 'C-label', type: 'number'}
            ],
            rows: [{c:[{v: 'a'},
                       {v: 1.0, f: 'One'},
                       {v: 100}
                  ]},
                   {c:[{v: 'b'},
                       {v: 2.0, f: 'Two'},
                       {v: 200}
                  ]},
                   {c:[{v: 'c'},
                       {v: 3.0, f: 'Three'},
                       {v: 300}
                  ]}
            ]
          }
*/

function getVoltageData($slot_num, $start_time, $end_time)
{
	list($year, $mon, $day, $hour, $min) = sscanf($start_time, "%d-%d-%dT%d:%d");
	$start_t=sprintf("%04d%02d%02d%02d%02d00", $year, $mon, $day, $hour, $min);

	list($year, $mon, $day, $hour, $min) = sscanf($end_time, "%d-%d-%dT%d:%d");
	$end_t=sprintf("%04d%02d%02d%02d%02d00", $year, $mon, $day, $hour, $min);
	
	$r = array();

	$servername = "localhost";
	$username = "phpmyadmin";
	$password = "PHP@password123";
	$dbname = "battery_ageing";

	$conn = mysqli_connect($servername, $username, $password, $dbname);

	if (!$conn) {
	    die("Connection failed: " . mysqli_connect_error());
	}
/*
	$sql = "SHOW TABLES";
	$result = mysqli_query($conn, $sql);

	$tables = array();

	$sql = "SELECT * FROM A484MB1358_20181220 where slotnum=$slot_num order by timestamp DESC limit 100";
*/

	$sql = "SELECT * FROM A484MB1358_20181220 where slotnum=$slot_num and timestamp between $start_t and $end_t order by timestamp";

	$result = mysqli_query($conn, $sql);

	if (mysqli_num_rows($result) > 0) {
		// output data of each row
		while($row = mysqli_fetch_assoc($result)) {
			$t = $row["timestamp"];
			$v = $row["voltage"];
			$xv = $row["xvoltage"];
			$rs = array(array("v"=>$t), array("v"=>$v), array("v"=>$xv));
			$rc = array("c"=>$rs);
			array_push($r, $rc);
		}
	}

	mysqli_close($conn);

	$h1=array("id"=>"a", "label"=>"time", "type"=>"string");
	$h2=array("id"=>"b", "label"=>"外部测量", "type"=>"number");
	$h3=array("id"=>"c", "label"=>"内部测量", "type"=>"number");
	//$h=array("cols"=>array($h1, $h2, $h3));

	//$c1=array(array("v"=>"a"), array("v"=>1000), array("v"=>400));
	//$c2=array(array("v"=>"a"), array("v"=>1170), array("v"=>460));
	//$c3=array(array("v"=>"a"), array("v"=>660), array("v"=>1120));
	//$c4=array(array("v"=>"a"), array("v"=>1030), array("v"=>540));
	//$c=array("rows"=>array(array("c"=>$c1), array("c"=>$c2), array("c"=>$c3), array("c"=>$c4)));

	//$data=array("cols"=>array($h1, $h2, $h3), "rows"=>array(array("c"=>$c1), array("c"=>$c2), array("c"=>$c3), array("c"=>$c4)));

	$data=array("cols"=>array($h1, $h2, $h3), "rows"=>$r);
	$json_data=json_encode($data, JSON_PRETTY_PRINT);

	return $json_data;
}

if (isset($_POST['slot_num']) && isset($_POST['start_time']) && isset($_POST['end_time'])) {
	echo getVoltageData($_POST['slot_num'], $_POST['start_time'], $_POST['end_time']);
}

?>
