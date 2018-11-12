package db_mysql

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"
)

var db *MysqlDB

func TestInit(t *testing.T) {

	config := MySQLConfig{
		"root",
		"password",
		"127.0.0.1",
		3306,
		"",
	}

	db, err := NewMySQLDB(config)
	if err != nil {
		fmt.Println("mysql database init error")
		os.Exit(1)
	}

	msg := MqttMsg{
		0,
		"17:23:13.123",
		"b8dfd052-99e8-4a67-ab68-2adcf81977a9",
		"SERTODEV",
		"OPENDEV",
		base64.StdEncoding.EncodeToString([]byte{1, 2, 3, 4}),
	}

	table_name := "mqtt" + time.Now().Local().Format("20060102")
	id, err := db.InsertMqttMsg(table_name, &msg)
	if err != nil {
		fmt.Println("insert err :", err)
	} else {
		fmt.Printf("insert id %d\n", id)
	}

	msgs, err := db.ListMqttMessages(table_name)
	if err != nil {
		fmt.Println("list err :", err)
	} else {
		for id, m := range msgs {
			fmt.Printf("list id %d msg %v\n", id, m)
		}
	}
	msgs, err = db.ListMqttmsgsOfCmd(table_name, "OPENDEV")
	if err != nil {
		fmt.Println("list err :", err)
	} else {
		for id, m := range msgs {
			fmt.Printf("list by cmd id %d msg %v\n", id, m)
		}
	}

	msg1, err := db.GetMqttMsgById(table_name, 1)
	if err != nil {
		fmt.Println("list err :", err)
	} else {
		fmt.Printf("list by id %d msg : %v\n", 1, msg1)
		str, err := base64.StdEncoding.DecodeString(msg1.Msgdata)
		if err != nil {
			fmt.Println("decode err ", err)
		} else {
			fmt.Println("decode data ", str)
		}
	}

	err = db.DeleteMqttMsg(table_name, 1)
	if err != nil {
		fmt.Println("delete id err :", err)
	} else {
		fmt.Println("delete msg ok")
	}

	db.Close()
}
