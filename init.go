package main

import "urlshorten.kwikbill.in/database"

func init() {
	err := database.InitMongo()

	if err != nil {
		panic(err)
	}
}
