package handler

import (
	"fmt"
	"net/http"
)

func ValidateQuery(w http.ResponseWriter, r *http.Request, validateQuery ...string) bool {
	if r.Method != validateQuery[0] {
		jsonResponse(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("bad %v method", r.Method))
		return false
	}
	for _, v := range validateQuery[1:] {
		if !r.URL.Query().Has(v) {
			jsonResponse(true, w, http.StatusBadRequest, "not parameters "+v)
			return false
		}
	}
	return true
}
