
function getDevicesn(url_string) {
	var url_string = "http://www.example.com/t.html?a=1&b=3&c=m2-m3-m4-m5"; //window.location.href
	var url = new URL(url_string);
	var c = url.searchParams.get("device_sn");
	console.log(c);
}

