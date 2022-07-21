package kafka

import (
	"strconv"

	"github.com/awesome-sphere/as-booking-consumer/kafka/interfaces"
	"github.com/awesome-sphere/as-booking-consumer/utils"

	"github.com/segmentio/kafka-go"
)

var BOOKING_TOPIC string
var CANCELING_TOPIC string
var BOOKING_PARTITION int
var CANCELING_PARTITION int
var KAFKA_ADDR string

func ListTopic(connector *kafka.Conn) map[string]*interfaces.TopicInterface {
	partitions, err := connector.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}
	m := make(map[string]*interfaces.TopicInterface)
	for _, p := range partitions {
		if _, ok := m[p.Topic]; ok {
			m[p.Topic].Partition += 1
		} else {
			m[p.Topic] = &interfaces.TopicInterface{Partition: 1}
		}
	}
	return m

}

func doesTopicExist(connector *kafka.Conn, topic_name string) bool {
	topic_map := ListTopic(connector)
	_, ok := topic_map[topic_name]
	return ok

}

func ConnectKafka() *kafka.Conn {
	conn, err := kafka.Dial("tcp", KAFKA_ADDR)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

func InitKafkaTopic() {
	BOOKING_TOPIC = utils.GetenvOr("KAFKA_BOOKING_TOPIC_NAME", "booking")
	CANCELING_TOPIC = utils.GetenvOr("KAFKA_CANCELING_TOPIC_NAME", "canceling")
	KAFKA_ADDR = utils.GetenvOr("KAFKA_ADDR", "localhost:9092")
	partition_num, err := strconv.Atoi(utils.GetenvOr("KAFKA_BOOKING_TOPIC_PARTITION", "5"))
	if err != nil {
		panic(err.Error())
	}
	BOOKING_PARTITION = partition_num
	CANCELING_partition_num, err := strconv.Atoi(utils.GetenvOr("KAFKA_CANCELING_TOPIC_PARTITION", "5"))
	if err != nil {
		panic(err.Error())
	}
	CANCELING_PARTITION = CANCELING_partition_num

	conn := ConnectKafka()
	defer conn.Close()

	if !doesTopicExist(conn, BOOKING_TOPIC) {
		updateBookingtopicConfigs := []kafka.TopicConfig{
			{
				Topic:             BOOKING_TOPIC,
				NumPartitions:     BOOKING_PARTITION,
				ReplicationFactor: 1,
			},
		}
		err := conn.CreateTopics(updateBookingtopicConfigs...)
		if err != nil {
			panic("Booking Topic: " + err.Error())
		}
	}

	if !doesTopicExist(conn, CANCELING_TOPIC) {
		cancelTopicConfigs := []kafka.TopicConfig{
			{
				Topic:             CANCELING_TOPIC,
				NumPartitions:     CANCELING_PARTITION,
				ReplicationFactor: 1,
			},
		}
		err := conn.CreateTopics(cancelTopicConfigs...)
		if err != nil {
			panic("Cancel Topic: " + err.Error())
		}
	}

	Consume(BOOKING_TOPIC, "booking-consumer", BOOKING_PARTITION)
	Consume(CANCELING_TOPIC, "canceling-consumer", CANCELING_PARTITION)
}
