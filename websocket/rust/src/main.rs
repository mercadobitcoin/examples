use url::Url;
use tungstenite::{connect, Message};


fn main() {

    let (mut socket, response) = connect(
        Url::parse("wss://ws.mercadobitcoin.net/ws").unwrap()
    ).expect("Can't connect");

    socket.write_message(Message::Text(r#"{"type": "subscribe","subscription": {"name": "orderbook","id": "BRLBTC", "limit": 10}}"#.into()));

    loop {
        let msg = socket.read_message().expect("Error reading message");
        println!("Received: {}", msg);
    }
}
