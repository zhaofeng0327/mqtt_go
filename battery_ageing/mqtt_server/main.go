package main

import (
	"../db_mysql"
	"../protobuf"
	//"encoding/base64"
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

var mqttPublishHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
/*
	topic := msg.Topic()
	i := strings.Index(topic, "/")
	topic1 := topic[:i]
	deviceid := topic[i+1:]
*/
	data := msg.Payload()
	cmsg := &battery_ageing.UploadInfo{}

	err := proto.Unmarshal(data, cmsg)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	} else {
        a := cmsg.GetBatteryAgeingInfo();
        l := len(a);
        fmt.Println("cnt ", l);

        for i:=0; i < l; i++ {

            fmt.Println("timestamp : ", *a[i].Timestamp, "BatterySn :", *a[i].BatterySn, "SlotNum", *a[i].SlotNum, "Temprature", *a[i].Temprature,
            "Voltage", *a[i].Voltage, "Current", *a[i].Current, "Elapsed", *a[i].Elapsed);

            timestamp := *a[i].Timestamp;
            batterySn := *a[i].BatterySn;
            slotNum := *a[i].SlotNum;
            temprature := *a[i].Temprature;
            voltage := *a[i].Voltage;
            current :=  *a[i].Current;
            elapsed := *a[i].Elapsed;

			msg := db_mysql.MqttMsg{
				0,
                timestamp,
                batterySn,
                slotNum,
                temprature,
                voltage,
                current,
                elapsed,
            };

			//save to mysql
			t := time.Now().Local()
			table_name := "battery_ageing" + t.Format("20060102")
			id, err := db.InsertMqttMsg(table_name, &msg)
			if err != nil {
				fmt.Printf("insert id %d err\n",id)
			} else {
				fmt.Printf("insert id %d\n",id)
			}
        }
        /*
		head := cmsg.GetMsgHead()
		if head == nil {
			fmt.Println("get msg head error")
		} else {
			cmd := head.GetCmd()

			//save to mysql
			t := time.Now().Local()
			table_name := "mqtt" + t.Format("20060102")
			timestamp := t.Format("15:04:05.000 2006-01-02")
			msg := db_mysql.MqttMsg{
				0,
				timestamp,
				deviceid,
				topic1,
				msginfo.CMD_name[int32(cmd)],
				base64.StdEncoding.EncodeToString(data),
			}
			id, err := db.InsertMqttMsg(table_name, &msg)
			if err != nil {
				fmt.Printf("%-64s%-32s%-8d err\n", topic, cmd, id)
			} else {
				fmt.Printf("%-64s%-32s%-8d\n", topic, cmd, id)
			}
		}
    */
	}

}

var mqttConnectionLostHanler = func(client mqtt.Client, err error) {
	fmt.Println("mqtt connection lost!!!")
}

func connMqttServer() {

	for {

		if token := c.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		} else {
			break
		}

		time.Sleep(5 * time.Second)
	}

	if token := c.Subscribe("REQ/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	if token := c.Subscribe("RES/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
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
	connMqttServer()

	for {
		time.Sleep(60 * time.Second)
	}

	/*
		if token := c.Unsubscribe("DEVTOSER/#"); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

		if token := c.Unsubscribe("SERTOSER/#"); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

		c.Disconnect(250)

		time.Sleep(1 * time.Second)

		db.Close()
	*/

}
