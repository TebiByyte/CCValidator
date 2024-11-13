package validate

import (
	"io"
	"net/http"
)

type ValidateCardRequestBody struct {
    CardNumber int64 `json:"cardNumber"`
}

func validateRequest(request *http.Request) bool {
    return request.Method == "GET"
}

func validateCardNumber(int64 number) bool {
    return true // TODO implement Luhn's Algorithm
}

func HandleValidateRequest (writer http.ResponseWriter, request *http.Request) {
    if !validateRequest(request) {
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

    var requestBody ValidateCardRequestBody


    
    if !validateCardNumber

    writer.WriteHeader(http.StatusOK)
}
