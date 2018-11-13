package api

import (
	"encoding/json"
	"net/http"
)

func GetCandidates(w http.ResponseWriter, r *http.Request) {
	cState, err := GetStateForRequest(r)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(Response{
			Code: 404,
			Log:  "State for given height not found",
		})
		return
	}

	candidates := cState.GetStateCandidates().GetData()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var result []Candidate

	for _, candidate := range candidates {
		result = append(result, makeResponseCandidate(candidate, false))
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(Response{
		Code:   0,
		Result: result,
	})
}
