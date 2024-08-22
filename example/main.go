package main

import (
	"fmt"
	"log"

	kaj_log_activities_sdk "github.com/kiriminaja/kaj-log-activities-sdk/sdk"
)

func main() {
	client := kaj_log_activities_sdk.NewClient(kaj_log_activities_sdk.Config{
		URL: "URL",
	})
	// payload := map[string]interface{}{
	// 	"source":     "kaj-prd-core-srvc",
	// 	"package_id": 11429683,
	// 	"order_id":   "OID-7951969212",
	// 	"history": map[string]interface{}{
	// 		"event_name": "final_state",
	// 		"data": map[string]interface{}{
	// 			"response": map[string]interface{}{
	// 				"id":       11429683,
	// 				"order_id": "OID-7951969212",
	// 				"status":   200,
	// 			},
	// 		},
	// 	},
	// }
	// result, err := client.Events.Upsert("request_pickup", payload)
	// if err != nil {
	// 	log.Fatalf("ERROR: %s", err.Error())
	// }
	// fmt.Println("RESULT:> ", result)

	data, err := client.Events.List()
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	fmt.Println(data["disbursement"])

}
