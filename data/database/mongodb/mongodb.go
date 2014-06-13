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

	// Delete values:
	// DELETE FROM users WHERE age >= 0;
	change, err := collection.RemoveAll(bson.M{"age": bson.M{"$gte": 0}})
	checkErr(err)
	fmt.Printf("updated %d and removed %d after executing RemoveAll().\n", change.Updated, change.Removed)
	
	// Insert values:
	err = collection.Insert(&User{"Adam", 14}, &User{"Bob", 19}, &User{"Chris", 10}, &User{"Dennis", 35}, &User{"Emma", 26}, &User{"Frank", 41})
	checkErr(err)

	// Update values:
	// UPDATE users SET age = 15 WHERE name = 'Adam';
	err = collection.Update(bson.M{"name": "Adam"}, bson.M{"$set": bson.M{"age": 15}})
	checkErr(err)
	
	// Query and sort:
	// SELECT * FROM users WHERE age >= 10 && age <= 50 ORDER BY name ASC, age DESC;
	query := collection.Find(bson.M{"age": bson.M{"$gte": 10, "$lte": 50}}).Sort("name", "-age")
	// Travel result set:
	iter := query.Iter()
	u := User{}
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
