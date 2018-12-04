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
	"strings"
	"time"
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
const CLIENT_ID_SERVER = "client_id_server"

var mqttPublishHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	topic := msg.Topic()
	i := strings.Index(topic, "/")
	topic1 := topic[:i]
	//topic2 := topic[i+1:]
	fmt.Println("recv topic : ", topic)

	if (topic1 != TOPIC_CTS_TYPE) {
		fmt.Println("topic not wanted")
		return
	}

	data := msg.Payload()
	cmsg := &battery_ageing.MSG_BODY{}

	err := proto.Unmarshal(data, cmsg)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	} else {

		if (battery_ageing.MSG_TYPE_UPLOAD_INFO != cmsg.GetType()) {
			fmt.Println("msg type not wanted")
			return
		}

        dev_sn := cmsg.GetDeviceSn();
        if (10 != len(dev_sn)) {
            fmt.Println("device sn len err \n", dev_sn)
            return
        }

        up := cmsg.GetUpInfo();
        if (up == nil) {
            fmt.Println("get upload info err\n")
            return
        }

		a := up.GetBatteryAgeingInfo();
        if (a == nil) {
            fmt.Println("get battery ageing info err\n")
            return
        }

        l := len(a);
        fmt.Println("cnt ", l);

        for i:=0; i < l; i++ {

            fmt.Println(
				"Timestamp"           , *a[i].Timestamp          ,
				"BatterySn"           , *a[i].BatterySn          ,
				"SlotNum"             , *a[i].SlotNum            ,
				"Voltage"             , *a[i].Voltage            ,
				"Current"             , *a[i].Current            ,
				"Temprature"          , *a[i].Temprature         ,
				"Elapsed"             , *a[i].Elapsed            ,
				"Discharging"         , *a[i].Discharging        ,
				"XTemprature"         , *a[i].XTemprature        ,
				"XVoltage"            , *a[i].XVoltage           ,
				"XFullChargecapacity" , *a[i].XFullChargecapacity,
				"XRemainingcapacity"  , *a[i].XRemainingcapacity ,
				"XAveragecurrent"     , *a[i].XAveragecurrent    ,
				"XCyclecount"         , *a[i].XCyclecount        ,
				"XBmssafetyStatus"    , *a[i].XBmssafetyStatus   ,
				"XBmsflags"           , *a[i].XBmsflags          ,
				"XBatterystatus"      , *a[i].XBatterystatus     ,
				"XChargestatus"       , *a[i].XChargestatus      ,
				"XEnablestatus"       , *a[i].XEnablestatus      ,
				"XSlotstatus"         , *a[i].XSlotstatus        ,
				"XDestroyed"          , *a[i].XDestroyed         ,
				"XHasbms"             , *a[i].XHasbms            ,
				"XRadio"              , *a[i].XRadio             )

            timestamp           := *a[i].Timestamp
            batterySn           := *a[i].BatterySn
            slotNum             := *a[i].SlotNum
            voltage             := *a[i].Voltage
            current             := *a[i].Current
			temprature          := *a[i].Temprature
            elapsed             := *a[i].Elapsed
            discharging         := *a[i].Discharging
            xtemprature         := *a[i].XTemprature
            xvoltage            := *a[i].XVoltage
            xfullchargecapacity := *a[i].XFullChargecapacity
            xremainingcapacity  := *a[i].XRemainingcapacity
            xaveragecurrent     := *a[i].XAveragecurrent
            xcyclecount         := *a[i].XCyclecount
            xbmssafetystatus    := *a[i].XBmssafetyStatus
            xbmsflags           := *a[i].XBmsflags
            xbatterystatus      := int32(*a[i].XBatterystatus)
            xchargestatus       := int32(*a[i].XChargestatus)
            xenablestatus       := int32(*a[i].XEnablestatus)
            xslotstatus         := *a[i].XSlotstatus
            xdestroyed          := *a[i].XDestroyed
            xhasbms             := *a[i].XHasbms
            xradio              := *a[i].XRadio

			msg := db_mysql.MqttMsg{
				0                   ,
				timestamp           ,
				batterySn           ,
				slotNum             ,
				voltage             ,
				current             ,
				temprature          ,
				elapsed             ,
				discharging         ,
				xtemprature         ,
				xvoltage            ,
				xfullchargecapacity ,
				xremainingcapacity  ,
				xaveragecurrent     ,
				xcyclecount         ,
				xbmssafetystatus    ,
				xbmsflags           ,
				xbatterystatus      ,
				xchargestatus       ,
				xenablestatus       ,
				xslotstatus         ,
				xdestroyed          ,
				xhasbms             ,
				xradio              ,
            };

			//save to mysql
            config := db_mysql.MySQLConfig{
                db_mysql.DATABASE_USER,
                db_mysql.DATABASE_PASSWORD,
                db_mysql.DATABASE_ADDR,
                db_mysql.DATABASE_PORT,
                "",
            }
			t := time.Now().Local()
			table_name := dev_sn + "_" + t.Format("20060102")
			id, err := db.InsertMqttMsg(config, table_name, &msg)
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

	if token := c.Subscribe(TOPIC_CTS, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
/*
	if token := c.Subscribe(TOPIC_STC, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
*/
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
		fmt.Println("mysql database init error :", err)
		os.Exit(1)
	} else {
		db = dbt
	}

	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker(MQTT_SERVER_ADDR).SetClientID(CLIENT_ID_SERVER)
	opts.SetAutoReconnect(true)
	opts.SetKeepAlive(MQTT_ALIVE_PERIOD * time.Second)
	opts.SetDefaultPublishHandler(mqttPublishHandler)
	opts.SetPingTimeout(MQTT_PING_TIMEOUT * time.Second)
	opts.SetUsername(MQTT_USER_NAME)
	opts.SetPassword(MQTT_PASSWORD)
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
