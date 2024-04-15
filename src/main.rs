use std::u64;

use serde::{Deserialize, Serialize};
use serde_with::serde_as;
use tungstenite::Message;

#[derive(Debug, Serialize, Deserialize)]
struct SearchResult {
    symbol: String,
    buy_price: f64,
    buy_quantity: f64,
    sell_price: f64,
    sell_quantity: f64,
    last_update_id: i64,
}

impl SearchResult {
    fn new(depth: DepthUpdate) -> Self {
        let symbol = depth.get_symbol().to_string();
        let (buy_price, buy_quantity) = depth.data.bids.iter().next().unwrap();
        let (sell_price, sell_quantity) = depth.data.asks.iter().next().unwrap();

        return SearchResult {
            symbol,
            buy_price: *buy_price,
            buy_quantity: *buy_quantity,
            sell_price: *sell_price,
            sell_quantity: *sell_quantity,
            last_update_id: depth.data.last_update_id,
        };
    }
}

#[derive(Debug, Serialize, Deserialize)]
pub struct SubscribeMessage {
    method: String,
    params: Vec<String>,
    id: i32,
}

// Deserialize bids and asks as vectors of f64
#[derive(Debug, Serialize, Deserialize)]
pub struct SubscribtionSuccess {
    result: Option<String>,
    id: i32,
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(untagged)]
pub enum BinanceMessage {
    SubscribtionSuccess(SubscribtionSuccess),
    DepthUpdate(DepthUpdate),
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let subscribe_message = SubscribeMessage {
        method: "SUBSCRIBE".to_string(),
        params: vec!["btcusdt@depth20@100ms".to_string()],
        id: 1,
    };

    let url = url::Url::parse("wss://stream.binance.com:9443/stream").unwrap();
    println!("Connecting to {}", url);

    let (mut socket, response) = tungstenite::connect(url).expect("asdf");

    println!("Connected to {}", response.status());
    for (ref header, ref value) in response.headers() {
        println!("{}: {}", header, value.to_str().unwrap());
    }

    socket
        .send(Message::Text(
            serde_json::to_string(&subscribe_message).unwrap(),
        ))
        .unwrap();

    loop {
        let events = socket.read().unwrap();
        let message = match events {
            Message::Ping(_) => {
                socket.send(Message::Pong(vec![])).unwrap();
                continue;
            }
            Message::Close(_) => {
                eprintln!("Connection closed");
                return Ok(());
            }
            Message::Text(message) => message,
            _ => {
                eprintln!("Unexpected message: {:?}", events);
                continue;
            }
        };
        let message: BinanceMessage = match serde_json::from_str(&message) {
            Ok(message) => message,
            Err(e) => {
                eprintln!("Error: {}", e);
                continue;
            }
        };
        match message {
            BinanceMessage::DepthUpdate(message) => {
                let search_result = SearchResult::new(message);
                println!("{:#?}", search_result);
            }
            BinanceMessage::SubscribtionSuccess(message) => {
                println!("{:#?}", message);
            }
        };
    }
}

#[derive(Debug, Serialize, Deserialize)]
pub struct DepthUpdate {
    stream: String,
    data: DepthUpdateData,
}

impl DepthUpdate {
    pub fn get_symbol(&self) -> &str {
        return self.stream.split("usdt").next().unwrap();
    }
}

/*
* Response from Binance
{
    lastUpdateId: 160,
    bids: [
        [ "0.0024", "10.0000" ],
        ...
    ],
    asks: [
        [ "0.0025", "10.0000" ],
        ...
    ]
}
*/
#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
#[serde_as]
pub struct DepthUpdateData {
    last_update_id: i64,
    #[serde_as(as = "Vec<(DisplayFromStr, DisplayFromStr)>")]
    bids: Vec<(f64, f64)>,
    asks: Vec<(f64, f64)>,
}

impl DepthUpdateData {}
