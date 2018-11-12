package main

import (
	"../devproto"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
	"time"
)

var c, c1 mqtt.Client

var mqttPublishHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	topic := msg.Topic()

	data := msg.Payload()
	cmsg := &msginfo.CMsg{}

	err := proto.Unmarshal(data, cmsg)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	} else {
		head := cmsg.GetMsgHead()
		if head == nil {
			fmt.Println("get msg head error")
		} else {
			cmd := head.GetCmd()
			fmt.Printf("%-64s%-32s\n", topic, cmd)

			//transfer to local mqtt server
			token := c1.Publish(topic, 0, false, data)
			token.Wait()
		}
	}

}

var mqttConnectionLostHanler = func(client mqtt.Client, err error) {
	fmt.Println("anker mqtt connection lost!!!")
}

var mqttConnectionLostHanler1 = func(client mqtt.Client, err error) {
	fmt.Println("local mqtt connection lost!!!")
}

func connAnkerMqttServer() {

	for {

		if token := c.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		} else {
			break
		}

		time.Sleep(5 * time.Second)
	}

	if token := c.Subscribe("DEVTOSER/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	if token := c.Subscribe("SERTODEV/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func connLocalMqttServer() {

	for {

		if token := c1.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		} else {
			break
		}

		time.Sleep(5 * time.Second)
	}
}

func main() {

	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	//S323PA4121
	opts1 := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetClientID("d93f3165-d239-4583-9d48-ccc62b8b3d00")
	opts1.SetAutoReconnect(true)
	opts1.SetKeepAlive(60 * time.Second)
	opts1.SetPingTimeout(2 * time.Second)
	opts1.SetUsername("admin")
	opts1.SetPassword("public")
	opts1.SetConnectionLostHandler(mqttConnectionLostHanler1)

	c1 = mqtt.NewClient(opts1)
	connLocalMqttServer()

	opts := mqtt.NewClientOptions().AddBroker("tcp://58.87.68.11:1883").SetClientID("d93f3165-d239-4583-9d48-ccc62b8b3d00")
	opts1.SetAutoReconnect(true)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(mqttPublishHandler)
	opts.SetPingTimeout(2 * time.Second)
	opts.SetUsername("Y1533_2.7.40")
	opts.SetPassword("khKE@24!oW4@cFS1WyjZ")
	opts.SetConnectionLostHandler(mqttConnectionLostHanler)

	c = mqtt.NewClient(opts)
	connAnkerMqttServer()

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
	*/
}
