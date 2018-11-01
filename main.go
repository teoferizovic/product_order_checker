package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"product_order_checker/model"
	"github.com/BurntSushi/toml"
)

var conf model.Config

func task() {
	fmt.Println("I am runnning task.")
}

func init(){
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", conf)
}

func main() {
	s := gocron.NewScheduler()
	s.Every(2).Seconds().Do(task)
	<- s.Start()
}
