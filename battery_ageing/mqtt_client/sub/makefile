all:
	#gcc sample.c -o client -g -I /home/zf/work/paho.mqtt.c/src -L /home/zf/work/paho.mqtt.c/build/output  -lpaho-mqtt3as
	gcc sample.c -o client -g -I /home/zf/work/paho.mqtt.c/src -L /home/zf/work/paho.mqtt.c/build/output -lpaho-mqtt3cs
	export LD_LIBRARY_PATH=/home/zf/work/paho.mqtt.c/build/output
clean:
	rm client
