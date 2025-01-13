package backend

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello Dayo")
}

func Run(addr string)  {
	http.HandleFunc("/", helloworld)
	fmt.Println(" server started and listening on port ", addr)
	log.Fatal(http.ListenAndServe(addr,nil))
}