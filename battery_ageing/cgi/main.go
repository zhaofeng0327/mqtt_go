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
)

var db *db_mysql.MysqlDB
var c mqtt.Client

const TOPIC_CTS = "CTS/#"
const TOPIC_STC = "STC/#"
const TOPIC_CTS_TYPE = "CTS"
const TOPIC_STC_CHARGE_SETTING = "STC/CHARGE"

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
		"root",
		"password",
		"127.0.0.1",
		3306,
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

	opts := mqtt.NewClientOptions().AddBroker("tcp://120.79.223.61:1883").SetClientID("123")
	opts.SetAutoReconnect(true)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(mqttPublishHandler)
	opts.SetPingTimeout(2 * time.Second)
	opts.SetUsername("admin")
	opts.SetPassword("public")
	opts.SetConnectionLostHandler(mqttConnectionLostHanler)

	c = mqtt.NewClient(opts)
	if 0 > connMqttServer() {
		fmt.Println("fail to connect to mqtt server");
		return
	}

	cmd := battery_ageing.DISCHARGE_CMD_START_DISCHARGE
	var q struct{};
	var b []byte;
	
	setting := &battery_ageing.DischargeSetting {
		&cmd,
		proto.Int32(1),
		q,
		b,
		0,
	}
	
	msg_type := battery_ageing.MSG_TYPE_DISCHARGE_SETTING
	msg_body := &battery_ageing.MSG_BODY {
		Type: &msg_type,
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
