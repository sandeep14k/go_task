package main

import (
	"CF_PROJECT/store"
	"CF_PROJECT/web"
	"CF_PROJECT/worker"
	"log"
	"sync"
)

func main() {
	mongoStore := new(store.MongoStore)
	mongoStore.OpenConnectionWithMongoDB()
	var wg sync.WaitGroup

	wg.Add(2)

	go worker.PerformWork(mongoStore, &wg)
	go server1(&wg)
	wg.Wait()
}
func server1(wg *sync.WaitGroup) {
	mongoStore := new(store.MongoStore)
	mongoStore.OpenConnectionWithMongoDB()
	defer wg.Done()
	srv := web.CreateWebServer(mongoStore)
	port := ":8082"
	if err := srv.StartListeningRequests(port); err != nil {
		log.Printf("Error occurred while starting the server: %v", err)
	}
}
