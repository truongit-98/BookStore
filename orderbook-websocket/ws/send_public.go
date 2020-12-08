package ws

import (
	"BookStore/models"
	"encoding/json"
	"fmt"
	"log"
)
const WS_EVENT_PUBLIC_ORDER_BOOK = "WS_EVENT_PUBLIC_ORDER_BOOK"

func SendOrder2Ws(orderBook *models.Order) {
	Hub.Broadcast <- MakeChannelMessageWithWrapper(WrapChannel{
		Channel: fmt.Sprintf("%s", WS_EVENT_PUBLIC_ORDER_BOOK),
		Data:    orderBook,
	})
}


func MakeChannelMessageWithWrapper(wrapperChannel WrapChannel) Message {
	bData, err := json.Marshal(wrapperChannel)
	if err != nil {
		log.Println(err.Error(), "-- err.Error() orderbook-websocket/kafka/consumer/consumer.go:232")
		return Message{
			Channel: wrapperChannel.Channel,
			Data:    []byte{},
		}
	}

	return Message{
		Channel: wrapperChannel.Channel,
		Data:    bData,
	}
}
