package helper

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
}
