let WebSocketAsPromised = require('websocket-as-promised');
const WebSocket = require('ws');


const wsUrl = 'ws://127.0.0.1:8080/ws';


const wsp = new WebSocketAsPromised(wsUrl, {
    createWebSocket: url => new WebSocket(url),
    extractMessageData: event => event, // <- this is important
    packMessage: data => (new Uint8Array(data)).buffer,
    unpackMessage: data => new Uint8Array(data),
    //binaryType: "arraybuffer",
});


wsp.onUnpackedMessage.addListener(message => {
    if (message.data instanceof ArrayBuffer) {
        // string received instead of a buffer
        console.log("---------------");
    }
    console.log("msg", message);
});

wsp.open()
    // .then ( () => {
    //     wsp.onMessage.addListener(message => console.log(message));
    // })
    .then(() => {

        wsp.sendPacked([1, 2, 3]);


    })
    .then(() => wsp.close())
    .catch(e => console.error("error", e));


