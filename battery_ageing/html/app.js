$(function() {
    var source = new EventSource('/events');
    source.onopen = function (event) {
        console.log("eventsource connection open");
    };
    source.onerror = function() {
        if (event.target.readyState === 0) {
            console.log("reconnecting to eventsource");
        } else {
            console.log("eventsource error");
        }
    };
    source.onmessage = function(event) {
        $('<div>', {
            text: event.data,
            css: {
                display: "none"
            }
        })
            .prependTo("#messages")
            .show("fast")
        ;
    };
});
