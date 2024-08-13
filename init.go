package main

import (
	"github.com/Delta456/box-cli-maker/v2"
	"urlshorten.kwikbill.in/config"
	"urlshorten.kwikbill.in/constants"
	"urlshorten.kwikbill.in/databases"
)

func init() {
	config.AppConfig = config.GetConfig()

	boxConfig := box.Config{
		Px:           10,
		Py:           1,
		Type:         "Round",
		Color:        "Yellow",
		TitleColor:   "Cyan",
		ContentAlign: "Center",
	}

	Box := box.New(boxConfig)

	if _, err := databases.InitDB(); err != nil {
		Box.Print("KwikBill v2 - Url Shorten Server", "Database Connection Failed!")
		panic(err)
	} else {
		Box.Print("KwikBill v2 - Url Shorten Server Started!", "Database Connected! \nServer Mode: "+constants.SERVER_MODE)
	}
}
