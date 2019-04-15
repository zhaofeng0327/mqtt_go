<html>

<head>
<title>(Type a title for your page here)</title>

<script type="text/javascript"> 
function reload()
{

var val1=document.form1.fname.value ;
var val2=document.form1.mname.value ;
var val3=document.form1.lname.value ;
//// For radio button value to collect ///
for(var i=0; i < document.form1.type.length; i++){
if(document.form1.type[i].checked)
var val4=document.form1.type[i].value 
}

self.location=window.location.pathname + '?fname=' + val1 + '&mname=' + val2 + '&lname=' + val3 + '&type=' + val4;

}

</script>
</head>

<body >

<?php

$fname=@$_GET['fname'];
$mname=@$_GET['mname'];
$lname=@$_GET['lname'];
$type=@$_GET['type'];
if($type=="male"){$maleck="checked";
$femaleck="";}
else{$femaleck="checked";
$maleck="";}

echo "<table border='0' width='50%' cellspacing='0' cellpadding='0' ><form name=form1 method=post action=index.php><input type=hidden name=todo value=post>


<tr ><td align=center ><font face='verdana' size='2'><br>
First Name <input type=text name=fname value='$fname'><br>
Middle Name <input type=text name=mname value='$mname'><br>
Last Name <input type=text name=lname value='$lname'><br>

<b>Type</b><input type=radio name=type value='male' $maleck>Male </font><input type=radio name=type value='female' $femaleck>Female</td></tr>

<tr bgcolor='#ffffff'><td align=center ><input type=button onclick='reload()'; value=Submit> <input type=reset value=Reset></td></tr>
</table></form>
";
?>

</body>

</html>
