#include <stdio.h>
#include <string.h>
#include "battery_ageing.pb-c.h"
#include "protobuf-c.h"

int main(int argc, char **argv)
{

    AMessage msg;
    char buf[1024] = { 0 };
    amessage__init(&msg);
    msg.a = 1;
    msg.has_b = 1;
    msg.b = 2;
    msg.n_d = 2;
    int data[2] = {100, 200};
    msg.d = data;


    BMessage bmsg[2];
    bmessage__init(&bmsg[0]);
    bmessage__init(&bmsg[1]);

    bmsg[0].m = "ma";
    bmsg[0].n = "na";
    bmsg[1].m = "mb";
    bmsg[1].n = "nb";

    msg.n_msg = 2;

    BMessage *bpmsg[2];
    msg.msg = bpmsg;
    msg.msg[0] = &bmsg[0];
    msg.msg[1] = &bmsg[1];

    amessage__pack(&msg, buf);
    int len = amessage__get_packed_size(&msg);


    AMessage *pmsg = amessage__unpack(NULL, len, buf);
    printf("a = %d b = %d\n \
            n_d = %zu d[0] = %d d[1] = %d\n \
            n_msg=%zu \n \
            msg[0].m = %s msg[0].n = %s\n \
            msg[1].m = %s msg[1].n = %s\n",
            pmsg->a, pmsg->b,
            pmsg->n_d, pmsg->d[0], pmsg->d[1],
            pmsg->n_msg,
            pmsg->msg[0]->m, pmsg->msg[0]->n,
            pmsg->msg[1]->m, pmsg->msg[1]->n );

	return 0;
}

