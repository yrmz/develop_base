package controllers

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Result string `json:"result"`
}

func Index(w http.ResponseWriter, r *http.Request) ([]byte, int) {

	var result = Result{
		Result: "ok",
	}

	output, err := json.Marshal(result)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return output, http.StatusOK
}
