#!/usr/bin/perl

local ($buffer, @pairs, $pair, $name, $value, %FORM);
# Read in text
$ENV{'REQUEST_METHOD'} =~ tr/a-z/A-Z/;

if ($ENV{'REQUEST_METHOD'} eq "POST") {
   read(STDIN, $buffer, $ENV{'CONTENT_LENGTH'});
} else {
   $buffer = $ENV{'QUERY_STRING'};
}

# Split information into name/value pairs
@pairs = split(/&/, $buffer);

foreach $pair (@pairs) {
   ($name, $value) = split(/=/, $pair);
   $value =~ tr/+/ /;
   $value =~ s/%(..)/pack("C", hex($1))/eg;
   $FORM{$name} = $value;
}

$device_sn = $FORM{device_sn};
$slot_num = $FORM{slot_num};
$option = $FORM{option};
$current_level = $FORM{current_level};
use Time::gmtime;
$gm = gmtime();
$timestamp = sprintf("%04d%02d%02d%02d%02d%02d",
	$gm->year + 1900,
	$gm->mon,
	$gm->mday,
	$gm->hour,
	$gm->min,
	$gm->sec);
$user="zf";

print "Content-type:text/html\r\n\r\n";
print "<html>";
print "<head>";
print "<title>Radio - Fourth CGI Program</title>";
print "</head>";
print "<body>";
#print "<h2> Selected level is $current_level</h2>";
#print "<h2> Selected slot is $slot_num</h2>";
#print "button $option\n";
print "柜机 $device_sn 卡槽 $slot_num 操作 $option 电流 $current_level";
print "</body>";
print "</html>";



system("/usr/lib/cgi-bin/bat_pub.cgi", $timestamp, $user, $device_sn, $slot_num, $option, $current_level);

1;
