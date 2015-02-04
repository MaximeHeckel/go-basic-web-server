package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

type Person struct {
	Name string
	Type string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello, world</h1>\nI'm running on %s with an %s CPU the name of this program is : %s", runtime.GOOS, runtime.GOARCH, os.Getenv("NAME"))
}

func main() {
	session, err := mgo.Dial(os.Getenv("MONGO"))

	if err == nil {
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("quicksell").C("people")
		err = c.Insert(&Person{"Golang", "Awesome"})
		if err != nil {
			log.Fatal(err)
		}
		result := Person{}

		err = c.Find(bson.M{"name": "Golang"}).One(&result)
		fmt.Println("Golang is: ", result.Type)
	}
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("running")
}
