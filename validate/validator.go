package validate

import (
	"encoding/json"
	"io"
	"net/http"
)

type ValidateCardRequestBody struct {
    CardNumber int64 `json:"cardNumber"`
}

type ValidateCardResponseBody struct {
    Result bool `json:"isValid"`
}

func validateRequest(request *http.Request) bool {
    return request.Method == "GET"
}

func validateCardNumber(number int64) bool {
    return true // TODO implement Luhn's Algorithm
}

func HandleValidateRequest (writer http.ResponseWriter, request *http.Request) {
    if !validateRequest(request) {
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

    body, readErr := io.ReadAll(request.Body)

    if readErr != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        return
    }

    defer request.Body.Close()
    

    var requestBody ValidateCardRequestBody
    err := json.Unmarshal(body, &requestBody)

    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        return
    }

    result := validateCardNumber(requestBody.CardNumber)
    writer.WriteHeader(http.StatusOK)
    writer.Header().Set("Content-Type", "application/json")

    response := ValidateCardResponseBody { Result: result }

    json.NewEncoder(writer).Encode(response)
}
