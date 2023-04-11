package consumer

import (
    "fmt"
    "github.com/Shopify/sarama"
)

type Consumer struct{
   self sarama.Consumer
   subscribedTopics []string
}

func (c *Consumer) Init(addr []string) (err error) {
    consumer, err := sarama.NewConsumer(addr, nil)
    if err != nil{
        return err
    }

    c.self = consumer
    return nil
}

func (c *Consumer) SubscribeTopic(topic string) (err error){
    topics, err := c.self.Topics()
    if err != nil{
        return err
    }

    if contains(topics, topic){
        c.subscribedTopics = append(c.subscribedTopics, topic)
        return nil
    }

    return fmt.Errorf("There is no such topic in broker: %s", topic)
}

func contains(s []string, str string) bool{
    for _, v := range s{
        if v == str{
            return true
        }
    }
    
    return false
}
