package main

import (
	"fmt"
	"net/http"
)

type MessageHandler struct {
	Message string
}

func main() {
	messageHandler := &MessageHandler{
		Message: "hello from web server",
	}

	http.HandleFunc("/", messageHandler.helloWorldHandler)
	http.HandleFunc("/books", messageHandler.bookHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}

func (m *MessageHandler) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.Message)
}

func (m *MessageHandler)bookHandler(w http.ResponseWriter,r *http.Request){

	switch r.Method{
	case "GET":
		genre:=r.URL.Query().Get("genre")
		if genre!=""{
			fmt.Fprintf(w,"Here is the list of %s genre books",genre)
		}else{
			fmt.Fprintf(w,"Here is the list of all books")
		}
		
	case "POST":
		fmt.Fprintf(w,"Adding new book")
		
	default:
		http.Error(w,"method not supported",http.StatusMethodNotAllowed)
	}
	
}
