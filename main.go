package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func main() {

	brokers := strings.Split(os.Getenv("KAFKA_BROKER_ADDRS"), ",")

	id := 0

	for {
		topic := fmt.Sprintf("mytopic-%d", id)

		fmt.Printf("Subscribing to topic: %s\n", topic)

		config := sarama.NewConfig()
		config.Version = sarama.V3_1_0_0

		saramaConsumer, err := sarama.NewConsumer(brokers, config)

		partitions, err := saramaConsumer.Partitions(topic)

		if err != nil {
			log.Fatalf("Error retrieving partitions for topic %s: %s\n", topic, err)
		}

		_, err = saramaConsumer.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)

		if err != nil {
			topics, err2 := saramaConsumer.Topics()

			if err2 != nil {
				log.Fatalf("Error listing topics: %s", err)
			}

			log.Printf("Could not consume partition %d in topic %s: %s\n"+
				"\tTopics in broker: %v\n"+
				"\tPartitions for topic: %v\n", partitions[0], topic, err, topics, partitions)
		}

		saramaConsumer.Close()

		id++

		time.Sleep(4 * time.Second)
	}
}
