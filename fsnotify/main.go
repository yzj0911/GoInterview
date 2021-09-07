package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watch.Close()
	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			select {
			case event, ok := <-watch.Events:
				if !ok {
					return
				}
				log.Printf("%s %s \n", event.Name, event.Op)
			case err, ok := <-watch.Errors:
				if !ok {
					return
				}
				log.Printf("Error: %s", err)
			}
		}
	}()
	err = watch.Add("./")
	if err != nil {
		log.Fatalf("Add Faild:%s", err)
	}
	<-done
}
