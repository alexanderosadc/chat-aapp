package main

import (
    "log"
	"github.com/alexanderosadc/chat-app/pkg/producer"
)

func main() {
	p := producer.Producer{}
    if err := p.Init([]string{"localhost:9092"}); err != nil{
        log.Println(err)
    }

    if err := p.SendMessage("Ana", "fa"); err != nil{
        log.Println(err)
    }

}
