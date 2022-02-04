package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RandomName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ChuckJoke struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}



func requests(url string, reqType string, c chan string){
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	// would like to make this more dry, these two if statments are essentially repeated
	if reqType == "chuck"{
		var chuckJoke ChuckJoke //TODO look into a shorthand for this
		//with JS, the type wouldn't be a problem and these if statements wouldn't need to be so repetative
		err = json.NewDecoder(res.Body).Decode(&chuckJoke)

		if err != nil {
			log.Fatalln(err)
		}

		c <- chuckJoke.Value.Joke
	}

	if reqType == "name" {
		var randomName RandomName

		err = json.NewDecoder(res.Body).Decode(&randomName)

		if err != nil {
			log.Fatalln(err)
		}

		c <- randomName.FirstName + " " + randomName.LastName
	}

}

func main() {

	c := make(chan string, 2)
	go requests("https://names.mcquay.me/api/v0/", "name", c)
	go requests("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy", "chuck", c)

	name, joke := <-c, <-c

	fmt.Println(name, joke)
}
