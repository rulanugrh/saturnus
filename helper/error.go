package helper

import "log"

func PrintError(err error) error {
	if err != nil {
		log.Printf("something error %v", err)
	}

	return nil
}