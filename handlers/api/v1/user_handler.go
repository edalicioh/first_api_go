package v1

import (
	"fmt"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserHandler v1")
}
