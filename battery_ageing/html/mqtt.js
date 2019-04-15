var mqtt;
var reconnectTimeout = 3000;
/*
var host="iot.eclipse.org";
var port=80;
*/
var host="120.79.223.61";
var port=8083;

function subsuccess() {
	console.log("subscribe success");
}

function subfail() {
	console.log("subscribe fail");
}

function onConnect() {
	console.log("connected");
	var ops = {
		onSuccess:subsuccess,
		onFailure:subfail,
	};
	mqtt.subscribe("CTS/#", ops);
}

function onFailure(message) {
	console.log("connection failed");
	setTimeout(MQTTconnect, reconnectTimeout);
}

function onConnectionLost(responseObject) {
  if (responseObject.errorCode !== 0) {
    console.log("onConnectionLost:"+responseObject.errorMessage);
  }
}

function onMessageArrived(msg) {
	console.log("message arrived");
}

function MQTTconnect() {
	console.log("connecting to " + host + " " + port);
	var dt = new Date();
	mqtt =  new Paho.MQTT.Client(host, port, "clientjsp" + dt.getTime());
	var options = {
		userName:"admin",
		password:"public",
		timeout:5,
		keepAliveInterval:100,
		useSSL:false,
		cleanSession:true,
		onSuccess:onConnect,
		onFailure:onFailure,
	};
	mqtt.onMessageArrived=onMessageArrived
	mqtt.onConnectionlost=onConnectionLost
	mqtt.connect(options);
}


