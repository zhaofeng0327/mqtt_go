package main

import (
	"../db_mysql"
	"../protobuf"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
	//"strings"
	"time"
    "strconv"
)

var db *db_mysql.MysqlDB
var c mqtt.Client

const MQTT_SERVER_ADDR = "tcp://120.79.223.61:1883"
const MQTT_USER_NAME = "admin"
const MQTT_PASSWORD = "public"
const MQTT_ALIVE_PERIOD = 60
const MQTT_PING_TIMEOUT = 2


const TOPIC_CTS = "CTS/#"
const TOPIC_STC = "STC/#"
const TOPIC_CTS_TYPE = "CTS"
const TOPIC_STC_CHARGE_SETTING = "STC/CHARGE"
const CLIENT_ID_CGI = "client_id_cgi"

var mqttPublishHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	topic := msg.Topic()
	//i := strings.Index(topic, "/")
	//topic1 := topic[:i]
	//topic2 := topic[i+1:]
	fmt.Println("recv topic : ", topic)

}

var mqttConnectionLostHanler = func(client mqtt.Client, err error) {
	fmt.Println("mqtt connection lost!!!")
}

func connMqttServer() int32{

	for i := 0; i < 3; i++ {

		if token := c.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		} else {
			return 0
		}

		time.Sleep(2 * time.Second)
	}

	return -1
}

func main() {

	config := db_mysql.MySQLConfig{
		db_mysql.DATABASE_USER,
		db_mysql.DATABASE_PASSWORD,
		db_mysql.DATABASE_ADDR,
		db_mysql.DATABASE_PORT,
		"",
	}

	dbt, err := db_mysql.NewMySQLDB(config)
	if err != nil {
		fmt.Println("mysql database init error")
		os.Exit(1)
	} else {
		db = dbt
	}

	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker(MQTT_SERVER_ADDR).SetClientID(CLIENT_ID_CGI)
	opts.SetAutoReconnect(true)
	opts.SetKeepAlive(MQTT_ALIVE_PERIOD * time.Second)
	opts.SetDefaultPublishHandler(mqttPublishHandler)
	opts.SetPingTimeout(MQTT_PING_TIMEOUT * time.Second)
	opts.SetUsername(MQTT_USER_NAME)
	opts.SetPassword(MQTT_PASSWORD)
	opts.SetConnectionLostHandler(mqttConnectionLostHanler)

	c = mqtt.NewClient(opts)
	if 0 > connMqttServer() {
		fmt.Println("fail to connect to mqtt server");
		return
	}

    dev_sn := os.Args[1]

    i64, err := strconv.ParseInt(os.Args[2], 10, 32)
    if err != nil {
        fmt.Printf("strconv to int err", err);
        return
    }
    slot_num := int32(i64)

    i64, err = strconv.ParseInt(os.Args[3], 10, 32)
    if err != nil {
        fmt.Printf("strconv to int err", err);
        return
    }
    var cmd battery_ageing.DISCHARGE_CMD
    //i32 := int32(i64)
    cmd = battery_ageing.DISCHARGE_CMD(i64)

    i64, err = strconv.ParseInt(os.Args[4], 10, 32)
    if err != nil {
        fmt.Printf("strconv2 to int err", err);
        return
    }
    level := int32(i64)

	var q struct{};
	var b []byte;

	t := time.Now().Local()
    timestamp := t.Format("20060102120102")
	user := "zf";
	
	setting := &battery_ageing.DischargeSetting {
		&timestamp,
		&user,
		proto.Int32(slot_num),
		&cmd,
		proto.Int32(level),
		q,
		b,
		0,
	}

	msg_type := battery_ageing.MSG_TYPE_DISCHARGE_SETTING
	msg_body := &battery_ageing.MSG_BODY {
		Type: &msg_type,
        DeviceSn : &dev_sn,
		DisSetting : setting,
	}

	data, err := proto.Marshal(msg_body)
		if err != nil {
			fmt.Println("marshaling error: ", err)
	}

	token := c.Publish(TOPIC_STC_CHARGE_SETTING, 0, false, data)
			token.Wait()


	c.Disconnect(250)
}
