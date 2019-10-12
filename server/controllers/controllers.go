package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Controllers : Struct offering controller methods and containing related
//resources, such as websocket write channel
type Controllers struct {
	SocketChannel chan []byte
}

func (c Controllers) WriteToChannel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Fatal("API supports only POST, was:", r.Method)
		http.Error(w, "Wrong method", http.StatusBadRequest)
	}

	bytes, err := ioutil.ReadAll(r.Body)
	fmt.Printf("BOOO! %v", r)
	if err != nil {
		log.Fatal("Error when calling writeToChannel", err)
		http.Error(w, "No body in call.", http.StatusBadRequest)
	}
	defer r.Body.Close()
	//This pushes incoming bytes to our websocket for easy testing
	c.SocketChannel <- bytes
	resp := "OK"
	w.Write([]byte(resp))

}
