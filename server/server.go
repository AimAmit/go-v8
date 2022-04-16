package server

import (
	"encoding/json"
	"go.kuoruan.net/v8go-polyfills/console"
	"log"
	"net/http"
	v8 "rogchap.com/v8go"
)

var iso = v8.NewIsolate()
var ctx = v8.NewContext(iso)

func serverRequest(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")

	if key == "" {
		w.Write([]byte("key missing"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodPost:
		err := HandlePost(r, key)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		return

	case http.MethodGet:
		res, err := HandleGet(key)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{"data": res})
	}

}

func Server() {

	if err := console.InjectTo(ctx); err != nil {
		panic(err)
	}

	InitializeRedisClient()
	http.HandleFunc("/", serverRequest)

	log.Println("server starting on 4000")
	http.ListenAndServe(":4000", nil)
}
