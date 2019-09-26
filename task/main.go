package main

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	http.HandleFunc("/task", taskHandler)
	appengine.Main()
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	for i := 1; ; i++ {
		log.Infof(ctx, "ver v1.0.4! loop: %d, time: %v", i, time.Now())
		time.Sleep(10 * time.Second)
	}

	log.Infof(ctx, "End task.")
	fmt.Fprint(w, "Hello, World!")
}
