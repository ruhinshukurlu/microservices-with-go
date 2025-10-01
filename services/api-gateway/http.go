package main

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "userid is required", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: "ok"}

	// TODO: Call Trip service
	writeJSON(w, http.StatusCreated, response)
}
