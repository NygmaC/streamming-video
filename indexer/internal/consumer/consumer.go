package consumer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/leocrispindev/streaming-video/indexer/internal/indexer"
	"github.com/leocrispindev/streaming-video/indexer/internal/model"
	"github.com/leocrispindev/streaming-video/stream-go-commons/pkg/broker/consumer"
)

var proccessConsumer consumer.Consumer

func Init() {

	proccessConsumer = consumer.CreateConsumer("", os.Getenv("KAFKA_INDEXER_TOPIC"))
	proccessConsumer.ReadMessage(handleMessage)

}

func handleMessage(msgs <-chan *sarama.ConsumerMessage) {
	fmt.Println("Consumer OK")

	for msg := range msgs {
		println("Message")
		var video = model.Document{}

		err := parse(msg.Value, &video)

		if err != nil {
			fmt.Println(err)
			continue

		}

		if video.Action == 1 {
			indexer.Index(video)

		} else if video.Action == 2 {
			indexer.Delete(video)

		} else {
			fmt.Println("Action not mapped")
		}

	}
}

func parse(value []byte, p *model.Document) error {
	fmt.Println(string(value))
	return json.Unmarshal(value, p)
}
