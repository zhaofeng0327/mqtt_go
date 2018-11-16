#include "MQTTClient.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <unistd.h>
#include "battery_ageing.pb-c.h"

//#define USE_SSL 1

#ifdef USE_SSL
#define ADDRESS "ssl://120.79.223.61:8883"
#else
#define ADDRESS "tcp://120.79.223.61:1883"
#endif 

#define ALI_MQTT_USER_NAME "admin"
#define ALI_MQTT_PASSWORD "public"

#define CLIENTID "0123456789"
#define TOPIC1 "REQ/#"
#define TOPIC2 "RES/#"
#define PAYLOAD "Hello World!"
#define QOS 1
#define TIMEOUT 10000L

static MQTTClient client;
volatile MQTTClient_deliveryToken deliveredtoken;

void delivered(void *context, MQTTClient_deliveryToken dt)
{
	printf("Message with token value %d delivery confirmed\n", dt);
	deliveredtoken = dt;
}

int publish_msg(char *topic, char *msg)
{
	MQTTClient_message pubmsg = MQTTClient_message_initializer;
	MQTTClient_deliveryToken token;

    AMessage Amsg;
    char buf[1024] = { 0 };
    amessage__init(&Amsg);
    Amsg.a = 1;
    Amsg.has_b = 1;
    Amsg.b = 2;
    Amsg.n_d = 2;
    int data[2] = {100, 200};
    Amsg.d = data;
    amessage__pack(&Amsg, buf);
    int len = amessage__get_packed_size(&Amsg);

	pubmsg.payload = buf;
	pubmsg.payloadlen = len;
	pubmsg.qos = QOS;
	pubmsg.retained = 0;
	deliveredtoken = 0;

	MQTTClient_publishMessage(client, topic, &pubmsg, &token);
	printf("Waiting for publication of %s"
	       "on topic %s for client with ClientID: %s\n",
	       PAYLOAD, topic, CLIENTID);
	while (deliveredtoken != token);
}

int msgarrvd(void *context, char *topicName, int topicLen,
	     MQTTClient_message *message)
{
	int i;
	char *payloadptr;
	printf("Message arrived topic: %s\n", topicName);
	payloadptr = message->payload;
#if 0
	for (i = 0; i < message->payloadlen; i++) {
		putchar(*payloadptr++);
	}
	putchar('\n');

	publish_msg(topicName, message->payload);
#endif
	MQTTClient_freeMessage(&message);
	MQTTClient_free(topicName);

	return 1;
}

void connlost(void *context, char *cause)
{
	printf("\nConnection lost\n");
	printf("     cause: %s\n", cause);
}

int main(int argc, char *argv[])
{
	MQTTClient_connectOptions conn_opts = MQTTClient_connectOptions_initializer;
	int rc;
	int ch;
	MQTTClient_create(&client, ADDRESS, CLIENTID,
			  MQTTCLIENT_PERSISTENCE_NONE, NULL);

	conn_opts.keepAliveInterval = 120;
	conn_opts.cleansession = 1;
	conn_opts.connectTimeout = 20;
	conn_opts.reliable = false;
	conn_opts.username = ALI_MQTT_USER_NAME;
	conn_opts.password = ALI_MQTT_PASSWORD;

#ifdef USE_SSL
	MQTTClient_SSLOptions ssl_opts = MQTTClient_SSLOptions_initializer;
	conn_opts.ssl = &ssl_opts;
	conn_opts.ssl->trustStore = "/home/zf/work/program/mqtt/ssl/ca.crt";
	conn_opts.ssl->enabledCipherSuites = "TLSv1.2";
#endif

	MQTTClient_setCallbacks(client, NULL, connlost, msgarrvd, delivered);
	if ((rc = MQTTClient_connect(client, &conn_opts)) !=
	    MQTTCLIENT_SUCCESS) {
		printf("Failed to connect, return code %d\n", rc);
		exit(EXIT_FAILURE);
    } else {
        printf("connet to %s success\n", ADDRESS);
    }

	printf("Subscribing to topic %s\nfor client %s using QoS%d\n",
	       TOPIC1, CLIENTID, QOS);
	MQTTClient_subscribe(client, TOPIC1, QOS);

	printf("Subscribing to topic %s\nfor client %s using QoS%d\n\n",
	       TOPIC2, CLIENTID, QOS);
	MQTTClient_subscribe(client, TOPIC2, QOS);

    for (int i = 0; i < 3; i++) {
        publish_msg("REQ/123", "H");
        sleep(1);
    }

	do {
		ch = getchar();
	} while (ch != 'Q' && ch != 'q');
	MQTTClient_disconnect(client, 10000);
	MQTTClient_destroy(&client);

	return rc;
}
