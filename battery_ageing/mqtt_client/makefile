all:
	#gcc battery_ageing.pb-c.c battery_ageing.pb-c.h client.c -o client -g -I /home/zf/work/paho.mqtt.c/src -L /home/zf/work/paho.mqtt.c/build/output  -lpaho-mqtt3as
	gcc battery_ageing.pb-c.c battery_ageing.pb-c.h client.c -o client -g -I /usr/local/include -L /usr/local/lib -lprotobuf-c -I /home/zf/work/paho.mqtt.c/src -L /home/zf/work/paho.mqtt.c/build/output -lpaho-mqtt3c
clean:
	rm client

 #export LD_LIBRARY_PATH=/home/zf/work/paho.mqtt.c/build/output
