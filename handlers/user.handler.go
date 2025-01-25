package handlers

import (
	"fmt"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the users handler!")
}
