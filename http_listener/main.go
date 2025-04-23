package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	remote_address := r.RemoteAddr
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Forbidden Method Request %s from %s \n", r.Method, remote_address)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil{
		fmt.Printf("Error reading body from %s \n", remote_address)
		return
	}
	fmt.Printf("%s \n", body)
	fmt.Fprintf(w, "This is your request text:\n%s\n", body)
}

func main(){
	mux := http.NewServeMux();
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr : "0.0.0.0:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil{
		fmt.Printf("FATAL ListenAndServe Error")
	}
	fmt.Println("Server started on Port 8080")
}
