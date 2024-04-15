package main

import (
	"errors"
	"log"
	"strconv"

	"golang.org/x/net/websocket"
)

type SubscribeMessage struct {
	Id     int      `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type Order struct {
	Price    float64
	Quantity float64
}

type OrderBook struct {
	Symbol string
	Bids   []Order
	Asks   []Order
}

type SearchedOrder struct {
	Symbol       string
	BuyPrice     float64
	BuyQuantity  float64
	SellPrice    float64
	SellQuantity float64
}

// search price until volume is greater than 100
// if it's looped until the end and not found, return error
func (o *OrderBook) search() (SearchedOrder, error) {
	var searchedOrder SearchedOrder
	searchedOrder.Symbol = o.Symbol
	for i := 0; i < len(o.Bids); i++ {
		if o.Bids[i].Quantity > 100 {
			searchedOrder.BuyPrice = o.Bids[i].Price
			searchedOrder.BuyQuantity = o.Bids[i].Quantity
			break
		}

		if i == len(o.Bids)-1 {
			return searchedOrder, errors.New("Not found")
		}
	}

	for i := 0; i < len(o.Asks); i++ {
		if o.Asks[i].Quantity > 100 {
			searchedOrder.SellPrice = o.Asks[i].Price
			searchedOrder.SellQuantity = o.Asks[i].Quantity
			break
		}

		if i == len(o.Asks)-1 {
			return searchedOrder, errors.New("Not found")
		}
	}

	if searchedOrder.BuyPrice == 0 || searchedOrder.SellPrice == 0 {
		return searchedOrder, errors.New("Not found")
	}

	searchedOrder.Symbol = o.Symbol
	return searchedOrder, nil
}

type OrderBookDepthMessage struct {
	EventType     string     `json:"e"`
	EventTime     int        `json:"E"`
	Symbol        string     `json:"s"`
	FirstUpdateId int        `json:"U"`
	FinalUpdateId int        `json:"u"`
	Bids          [][]string `json:"b"`
	Asks          [][]string `json:"a"`
}

func (o *OrderBookDepthMessage) Parse() OrderBook {
	var orderBook OrderBook
	orderBook.Symbol = o.Symbol
	for _, bid := range o.Bids {
		price, _ := strconv.ParseFloat(bid[0], 64)
		quantity, _ := strconv.ParseFloat(bid[1], 64)
		orderBook.Bids = append(orderBook.Bids, Order{price, quantity})
	}
	for _, ask := range o.Asks {
		price, _ := strconv.ParseFloat(ask[0], 64)
		quantity, _ := strconv.ParseFloat(ask[1], 64)
		orderBook.Asks = append(orderBook.Asks, Order{price, quantity})
	}

	return orderBook
}

type SubscribeResponse struct {
	Result string `json:"result"`
	Id     int    `json:"id"`
}

func main() {
	ws, err := websocket.Dial("wss://stream.binance.com:9443/ws/stream", "", "https://stream.binance.com")
	if err != nil {
		log.Fatal(err)
	}
	// loop forever
	subscribeMessage := SubscribeMessage{1, "SUBSCRIBE", []string{"btcusdt@depth"}}
	err = websocket.JSON.Send(ws, subscribeMessage)
	if err != nil {
		log.Fatal(err)
	}

    var message map[string]interface{}
	for {
        websocket.JSON.Receive(ws, &message)

        if message["e"] == "depthUpdate" {
            var orderbook OrderBookDepthMessage = OrderBookDepthMessage{
                EventType:     message["e"].(string),
                EventTime:     int(message["E"].(float64)),
                Symbol:        message["s"].(string),
                FirstUpdateId: int(message["U"].(float64)),
                FinalUpdateId: int(message["u"].(float64)),
                Bids:          message["b"].([][]string),
                Asks:          message["a"].([][]string),
            }
            log.Println(orderbook)
        }

        if message["id"] == 1 {
            subscribeResponse := SubscribeResponse{
                Result: message["result"].(string),
                Id:     int(message["id"].(float64)),
            }
            log.Println(subscribeResponse)
        }

        log.Println(message)
	}
}
