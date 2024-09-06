package banking

import (
	"encoding/json"
	"net/http"

	"golang.org/x/exp/rand"
)

const creditScoreMin = 500
const creditScoreMax = 900

type credit_rating struct {
	CreditRating int `json:"credit_rating"`
}

func GetCreditScore(w http.ResponseWriter, r *http.Request) {
	var creditRating = credit_rating{
		CreditRating: (rand.Intn(creditScoreMax-creditScoreMin) + creditScoreMin),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(creditRating)
}
