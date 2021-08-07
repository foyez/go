package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// https://swapi.dev/api/people
/*
{
	"results": [
		{
			"name": "Luke Skywalker",
			"homeworld": "https://swapi.dev/api/planets/1/",
		}
	]
}
*/

// https://swapi.dev/api/planets/1/
/*
{
	"name": "Tatooine",
	"terrain": "desert",
	"population": "200000",
}
*/

const BaseURL = "https://swapi.dev/api/"

type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

type AllPeople struct {
	People []Person `json:"results"`
}

func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Failed to parse response body")
	}

	if err := json.Unmarshal(bytes, &p.Homeworld); err != nil {
		log.Print("Error parsing json")
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "getting people")
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	// fmt.Println(res)

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse response body")
	}

	// fmt.Println(string(bytes))

	var people AllPeople

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	// fmt.Println(people)

	for _, person := range people.People {
		person.getHomeworld()
		fmt.Println(person)
	}
}

func main() {
	http.HandleFunc("/people", getPeople)

	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
