package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type person struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
	Phone string `bson:"phone"`
}

func main() {
	fmt.Println("Connecting to mongodb...")

	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	peopleCollection := session.DB("dino").C("people")

	people := []interface{}{person{
		Name:  "Michael",
		Age:   35,
		Email: "michael_voo@gmail.com",
		Phone: "1234",
	}, person{
		Name:  "Bianca",
		Age:   19,
		Email: "bianca2008@gmail.com",
		Phone: "367192",
	}}

	err = peopleCollection.Insert(people...)

	if err != nil {
		log.Fatal(err)
	}

	_, err = peopleCollection.UpdateAll(bson.M{"name": "Peter"}, bson.M{"$set": bson.M{"age": 10}})
	if err != nil {
		log.Fatal(err)
	}

	_, err = peopleCollection.RemoveAll(bson.M{"name": "Luke"})
	if err != nil {
		log.Fatal(err)
	}

	query := bson.M{
		"age": bson.M{
			"$lt": 40,
		},
		//"age": bson.M{
		//	"$in": []int{19, 28},
		//},
	}

	results := []person{}
	peopleCollection.Find(query).All(&results)

	for _, v := range results {
		fmt.Println(v.Name)
	}
	fmt.Println(results)
}
