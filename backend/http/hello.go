package http

import (
	"fmt"
	"net/http"
)

func HelloGo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO")
}
