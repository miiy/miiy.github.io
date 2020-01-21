package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func main()  {
	newClient()
	ping()
	string()
	list()
	hash()
	del()
	defer client.Close()
}

func newClient()  {
	client = redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "",
		DB:                 0,
	})

}

func ping()  {
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func string()  {
	err := client.Set("key", "value",0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func list()  {
	val, err := client.LPush("myList", "one", "two", "three").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	val2, err := client.LRange("myList", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val2)
}

func hash() {
	val, err := client.HSet("myHset", "name", "jack").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	val2, err := client.HGet("myHset", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val2)
}

func del()  {
	val, err := client.Del("key", "myList", "myHset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

/*
Redis client for Golang: https://github.com/go-redis/redis

output:

PONG <nil>
key value
key2 does not exist
3
[three two one]
true
jack
3

 */