// Package testgrp is for learning.
package testgrp

import (
	"context"
	"encoding/json"
	"net/http"
)

func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return json.NewEncoder(w).Encode(status)
}
