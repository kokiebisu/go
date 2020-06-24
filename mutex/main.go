package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Due to the fact that the web server is asynchronous,
// we’ll have to guard our counter using a mutex in
// order to prevent us from being hit with race-condition bugs.
var counter int

// instantiated an object (the address so that you don't have to add the & for every mutex afterwards)
var mutex = &sync.Mutex{}

func incrementCounter(rw http.ResponseWriter, r *http.Request) {
	// you don't want the copy to be modifies because that loses
	// the whole point

	// you want to modify the global variable's state
	mutex.Lock()
	counter++
	fmt.Fprint(rw, counter)
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/increment", incrementCounter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
