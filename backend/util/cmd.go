package util

import (
	"fmt"
	"log"
)

func ReciveCmd() string {
	var response string

	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
