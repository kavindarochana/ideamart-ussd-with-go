package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kavindarochana/ussdapp/utils"
)

func sender(out Out) {
	url := "http://localhost:7000/ussd/send/"

	message := map[string]interface{}{
		"applicationId":      out.ApplicationId,
		"password":           out.Password,
		"version":            out.Version,
		"message":            out.Message,
		"sessionId":          out.SessionId,
		"ussdOperation":      out.UssdOperation,
		"destinationAddress": out.DestinationAddress,
		"encoding":           out.Encoding,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	utils.Debug("api res - ", result)
}
