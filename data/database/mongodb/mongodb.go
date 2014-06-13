package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Name string
	Age int
}

func main() {
	// Connect to MongoDB server to get a session:
	session, err := mgo.Dial("localhost:27017")
	checkErr(err)
	defer session.Close()
	//session.SetMode(mgo.Monotonic, true)
	
	// Select DB and collection:
	collection := session.DB("test").C("users")
	
	// Insert values:
	err = collection.Insert(&User{"Adam", 14}, &User{"Bob", 19}, &User{"Chris", 10}, &User{"Dennis", 35}, &User{"Emma", 26}, &User{"Frank", 41})
	checkErr(err)

	u := User{}// Temp var to store query result
	
	// Query and sort:
	// SELECT * FROM users WHERE age >= 10 && age <= 50 ORDER BY name ASC, age DESC;
	query := collection.Find(bson.M{"age": bson.M{"$gte": 10, "$lte": 50}}).Sort("name", "-age")
	// Travel result set:
	iter := query.Iter()
	for iter.Next(&u) {
    		fmt.Printf("%s: %d\n", u.Name, u.Age)
	}
	// Close iter:
	err = iter.Close()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
