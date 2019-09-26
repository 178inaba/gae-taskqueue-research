package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
)

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := appengine.NewContext(r)

	t := taskqueue.NewPOSTTask("/task", nil)
	t.Header.Set("Host", "taskqueue-research-task")
	log.Infof(ctx, "Header: %v.", t.Header)
	if _, err := taskqueue.Add(ctx, t, "taskqueue-research-queue"); err != nil {
		msg := fmt.Sprintf("Add task: %v.", err)
		log.Errorf(ctx, msg)
		fmt.Fprint(w, msg)
		return
	}

	log.Infof(ctx, "post task!!")
	fmt.Fprint(w, "Hello, World!")
}
