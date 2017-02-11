package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type statusCode struct {
	Code int
	Description string
}

func main() {
	var data []statusCode
	rcvd := `[{"Code":200,"Description":"StatusOK"},{"Code":301,"Description":"StatusMovedPermanently"},{"Code":302,"Description":"StatusFound"},{"Code":303,"Description":"StatusSeeOther"},{"Code":307,"Description":"StatusTemporaryRedirect"},{"Code":400,"Description":"StatusBadRequest"},{"Code":401,"Description":"StatusUnauthorized"},{"Code":402,"Description":"StatusPaymentRequired"},{"Code":403,"Description":"StatusForbidden"},{"Code":404,"Description":"StatusNotFound"},{"Code":405,"Description":"StatusMethodNotAllowed"},{"Code":418,"Description":"StatusTeapot"},{"Code":500,"Description":"StatusInternalServerError"}]`

	if err := json.Unmarshal([]byte(rcvd), &data); err != nil {
		log.Fatalln(err)
	}

	for _, sc := range data {
		fmt.Printf("Code: %d, Description: %s\n", sc.Code, sc.Description)
	}
}
