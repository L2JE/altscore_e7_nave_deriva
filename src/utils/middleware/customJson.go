package middleware

import (
	"encoding/json"
	"math"
	"net/http"
	"reflect"
)

func EncodeRoundingFloats(w http.ResponseWriter, target any, precission int) {
	reflectedTarget := reflect.ValueOf(target).Elem()

	for i := 0; i < reflectedTarget.NumField(); i++ {
		field := reflectedTarget.Field(i)

		if field.IsValid() && field.CanSet() && (field.Kind() == reflect.Float64 || field.Kind() == reflect.Float32) {
			field.SetFloat(toFixed(field.Float(), precission))
		}
	}

	json.NewEncoder(w).Encode(target)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
