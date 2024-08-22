package main

import (
	"github.com/kiriminaja/kaj-log-activities-sdk/cmd/app"
	kaj_log_activities_sdk "github.com/kiriminaja/kaj-log-activities-sdk/sdk"
)

func main() {
	client := kaj_log_activities_sdk.NewClient(kaj_log_activities_sdk.Config{
		URL: "http://localhost:8081",
	})

	app.Start(client)
}
