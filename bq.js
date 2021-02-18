const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

var WebSocket = require('ws');

ws = new WebSocket("ws://127.0.0.1:8080");

ws.onopen = function() {}
ws.onmessage = function(m) {
    console.log(m)
}

delay(2000000);
