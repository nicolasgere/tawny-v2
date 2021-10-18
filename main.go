package main

import (
	"github.com/hashicorp/consul/api"
	"io/ioutil"
	"net/http"
)
import "fmt"

func main() {

	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	c, m, e := client.Catalog().Service("front", "", &api.QueryOptions{
		AllowStale: true,
	})
	fmt.Println(c, m, e)

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	http.HandleFunc("/watch", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
	})
	http.ListenAndServe(":8000", nil)
}
