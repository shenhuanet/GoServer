package main

import (
	"../bean"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	resp := new(bean.Response)
	resp.Code = 0
	resp.Message = "success"
	user := new(bean.User)
	user.Name = "Go"
	user.Age = 12
	resp.Data = user
	result, _ := json.Marshal(resp)
	_, _ = io.WriteString(w, string(result))
}
