package https

import (
	"log"
	"net/http"
)

func SmpServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("this is an exp server.\n"))
}

func Handler() {
	http.HandleFunc("/hello", SmpServer)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("Ls", err)
	}
}
