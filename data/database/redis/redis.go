package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	// Connect to Redis server:	
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkErr(err)
	//conn.Do("AUTH", "password")
	defer conn.Close()
	
	// Use SET command to put a single value:
	_, err = conn.Do("SET", "email", "i@RincLiu.com")
	checkErr(err)
	// Get bool value:
	exists, err := redis.Bool(conn.Do("EXISTS", "email"))
	checkErr(err)
	fmt.Printf("Is email exists: %#v\n", exists)
	if exists {
		// Use GET command to get a single value:
		email, err := redis.String(conn.Do("GET", "email"))
		checkErr(err)
		fmt.Printf("email: %s\n", email)
	}
	
	// Use RPUSH command to add items in list:
	_, err = conn.Do("RPUSH", "platformList", "Android")
	checkErr(err)
	_, err = conn.Do("RPUSH", "platformList", "iOS")
	checkErr(err)
	_, err = conn.Do("RPUSH", "platformList", "WindowsPhone")
	checkErr(err)
	// Use LRANGE command to get list items:
	listItems, err := redis.Values(conn.Do("LRANGE", "platformList", 0, 2))
	checkErr(err)
	fmt.Println("Items in List:")
	for _, itemV := range listItems {
		fmt.Printf("%s\n", itemV)
	}

	// Use SADD command to add items in set:
	_, err = conn.Do("SADD", "platformSet", "Android")
	checkErr(err)
	_, err = conn.Do("SADD", "platformSet", "iOS")
	checkErr(err)
	_, err = conn.Do("SADD", "platformSet", "WindowsPhone")
	checkErr(err)
	// Use SMENMBERS command to get set items:
	setMembers, err := redis.Values(conn.Do("SMEMBERS", "platformSet"))
	checkErr(err)
	fmt.Println("Items in Set:")
	for _, itemV := range setMembers {
		fmt.Printf("%s\n", itemV)
	}

	// Use HSET command to add k-v to HashMap:
	_, err = conn.Do("HSET", "userHashMap", "name", "Rinc")
	checkErr(err)
	_, err = conn.Do("HSET", "userHashMap", "email", "i@RincLiu.com")
	checkErr(err)
	// Use HKEYS command to get HashMap keys:
	mapKeys, err := redis.Values(conn.Do("HKEYS", "userHashMap"))
	checkErr(err)
	fmt.Println("Items in HashMap:")
	for _, itemKey := range mapKeys {
		// Use HGET command to get item value in HashMap by key:
		itemValue, err := conn.Do("HGET", "userHashMap", itemKey)
		checkErr(err)
		fmt.Printf("%s: %s\n", itemKey, itemValue)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
