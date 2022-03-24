const WebSocket = require('ws');
const client = new WebSocket('wss://ws.mercadobitcoin.net/ws');

client.on('open', function () {
    client.send(JSON.stringify({"type": "subscribe","subscription": {"name": "orderbook","id": "BRLBTC", "limit": 10}}));
})

client.on('message', function (msg) {
    console.log("%s", msg);
})
