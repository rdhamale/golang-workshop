package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type info struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type data struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

var wg sync.WaitGroup

type handler struct {
	id   string
	name string
	data string
}

func (handle *handler) serviceInfo() http.HandlerFunc {

	// handler.
	return func(w http.ResponseWriter, r *http.Request) {
		info := info{
			Id:   handle.id,
			Name: handle.name,
		}
		responseJson, err := json.MarshalIndent(info, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		w.Write(responseJson)
	}
}

func (handle *handler) serviceData() http.HandlerFunc {

	// handler.
	return func(w http.ResponseWriter, r *http.Request) {
		data := data{
			Id:   handle.id,
			Data: handle.data,
		}
		responseJson, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		w.Write(responseJson)
	}
}

func createHttpServer(ctx context.Context, ip string, port string) {
	mux := mux.NewRouter()
	handle := &handler{
		id:   port,
		data: "This is a dummy server-" + port,
		name: "server-" + port,
	}
	mux.HandleFunc("/info", handle.serviceInfo())
	mux.HandleFunc("/data", handle.serviceData())
	serverAddr := ip + ":" + port
	server := &http.Server{Addr: serverAddr, Handler: mux}
	go func() {
		wg.Add(1)
		log.Fatal(server.ListenAndServeTLS("server.crt", "server.key"))
	}()
	for {
		select {
		case <-ctx.Done():
			server.Shutdown(ctx)
			wg.Done()
			return
		}
	}
}

func main() {
	numberOfServers := flag.Int("servers", 10, "number of servers")
	flag.Parse()
	port := 2000
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	go func(ctx context.Context) {
		for iter := 0; iter < *numberOfServers; iter = iter + 1 {
			go createHttpServer(ctx, "127.0.0.1", strconv.Itoa(port))
			port = port + 1
		}
	}(ctx)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	for {
		select {
		case <-signalCh:
			fmt.Println("cancellation received!")
			cancelCtx()
			wg.Wait()
			return
		}
	}
}
