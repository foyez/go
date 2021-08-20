package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}

	http.NotFound(w, r)
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello myroute!") // send data to client side
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method) // get request method

	if r.Method == "GET" {
		t, _ := template.ParseFiles("register.gohtml")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		// parse to int
		age, _ := strconv.Atoi(r.Form.Get("age"))
		// match, err := regexp.MatchString("^[0-9]+$", r.Form.Get("age"))

		if age >= 18 {
			fmt.Println("Adult")
		}

		cities := []string{"dhaka", "cumilla", "feni"}
		validCity := func(city string) bool {
			for _, v := range cities {
				if v == r.Form.Get("city") {
					return true
				}
			}
			return false
		}
		if correctCity := validCity(r.Form.Get("city")); correctCity {
			fmt.Println("city is correct")
		}

		genderMap := map[string]string{
			"male":   "Male",
			"female": "Female",
		}
		if gender, ok := genderMap[r.Form.Get("gender")]; ok {
			fmt.Println("Gender:", gender)
		}

		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("Go launched at %s\n", t.Local())

		fmt.Println("username:", r.Form["username"])
		fmt.Println("age:", r.Form["age"])
		fmt.Println("interest:", r.Form["interest"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	// http.HandleFunc("/", sayHelloName) // set router
	// mux := &MyMux{}
	http.HandleFunc("/register", register)

	fmt.Println("Serving on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil)) // set listen port
	// log.Fatal(http.ListenAndServe(":9090", mux))
}
