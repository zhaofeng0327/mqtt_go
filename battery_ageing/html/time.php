<?php
$t = time();
print ($t . "\xA");
$t1 = strtotime("20181214182000\xA");
print ($t1. "\xA");
$s="2018-12-13T01:23";
list($year, $mon, $day, $hour, $min) = sscanf($s, "%d-%d-%dT%d:%d");
$ns=sprintf("%04d%02d%02d%02d%02d00", $year, $mon, $day, $hour, $min);
print $ns. "\xA";

$s1="2018-12-13T02:45";
list($year, $mon, $day, $hour, $min) = sscanf($s1, "%d-%d-%dT%d:%d");
$ns=sprintf("%04d%02d%02d%02d%02d00", $year, $mon, $day, $hour, $min);
print $ns. "\xA"
?>
