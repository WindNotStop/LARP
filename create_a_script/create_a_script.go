package main

import (
	"flag"
	"fmt"
	"github.com/WindNotStop/LARP/model"
	"github.com/go-redis/redis/v7"
	"io/ioutil"
	"log"
)
var scriptName = flag.String("sname", "剧本1", "剧本名")

func main(){
	flag.Parse()
	data, err := ioutil.ReadFile("main/script.json")
	if err != nil{
		log.Fatal(data)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err = client.Set(*scriptName, data, 0).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	val, err := client.Get(*scriptName).Bytes()
	if err != nil {
		log.Fatal(err.Error())
	}
	retrieved := &model.Script{}
	retrieved.Decode(val)
	fmt.Println("成功添加一个新剧本:")
	fmt.Println(retrieved.String())
}
