package polishserver

import (
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

var durationServerWaitUntilDisconnect = 3 * time.Second

func disconnectAfterSomeTime(w http.ResponseWriter, r *http.Request) {
	time.Sleep(durationServerWaitUntilDisconnect)
}

func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/poland", disconnectAfterSomeTime)
	http.ListenAndServe(":81", mux)
}
