package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Body) // output : &{[] {0xc00018e000} <nil> <nil>}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("body :", body) // output bertipe slice of byte

	if err != nil {
		log.Fatalln(err)
	}
	// untuk mencegah kebocoran memori atau memory leak.
	defer res.Body.Close()

	stringBody := string(body)
	fmt.Println("response body :", stringBody)
	/* Contoh output :
	   response body : {
	     "userId": 1,
	     "id": 1,
	     "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
	     "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
	   }
	*/
	log.Println(stringBody)
}
