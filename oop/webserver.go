package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type egps float32

func (e egps) String() string {
	return fmt.Sprintf("%g", e)
}

type databse map[string]egps

///// Solution 1 : context by function injector /////
// func main() {
// 	db := databse{"shoes": 300, "socks": 25}
// 	http.HandleFunc("/list", DbProvider(db, ListPrices))
// 	http.HandleFunc("/add", DbProvider(db, AddPrices))
// 	http.ListenAndServe("localhost:8000", nil)

// }

// func DbProvider(db databse, handler HandlerWithDb) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		handler(w, req, db)
// 	}
// }

// type HandlerWithDb func(w http.ResponseWriter, req *http.Request, db databse)

// func ListPrices(w http.ResponseWriter, req *http.Request, db databse) {
// 	encoder := json.NewEncoder(w)
// 	err := encoder.Encode(db)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func AddPrices(w http.ResponseWriter, req *http.Request, db databse) {
// 	decoder := json.NewDecoder(req.Body)
// 	var newPrices map[string]egps
// 	err := decoder.Decode(&newPrices)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for k, v := range newPrices {
// 		db[k] = v
// 	}
// }

// /// Solution 2: context by methods on object /////
func main() {
	db := databse{"shoes": 300, "socks": 25}
	http.HandleFunc("/list", db.ListPrices)
	http.HandleFunc("/patch", db.PatchPrices)
	http.ListenAndServe("localhost:8000", nil)
}

func (db databse) ListPrices(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(db)
	if err != nil {
		panic(err)
	}
}

func (db databse) PatchPrices(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var newPrices map[string]egps
	err := decoder.Decode(&newPrices)
	if err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal") {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		panic(err)
	}
	for k, v := range newPrices {
		db[k] = v
	}
}

// /// Solution 3: built-in context on request object /////
