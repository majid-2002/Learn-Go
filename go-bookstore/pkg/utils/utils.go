package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &x)
}
