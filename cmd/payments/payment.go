package payments

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
	kaj_log_activities_sdk "github.com/kiriminaja/kaj-log-activities-sdk/sdk"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func MainCommand(client *kaj_log_activities_sdk.Client) *cobra.Command {
	event := &cobra.Command{}
	event.Use = "payment"
	event.Short = "payment [command]"
	event.AddCommand(search(client))
	return event
}

func search(client *kaj_log_activities_sdk.Client) *cobra.Command {
	// client.Packages.Search()
	search := &cobra.Command{}
	search.Flags().StringP("search", "s", "", "Payment ID | TrxID")
	search.Flags().StringP("detail", "d", "", "Show Detail")
	search.Use = "search"
	search.Short = "search"
	search.Run = func(cmd *cobra.Command, args []string) {
		search, err := cmd.Flags().GetString("search")
		if err != nil {
			log.Fatalf("Argument search: %v", err)
		}
		data, err := client.Payments.Search(search)
		if err != nil {
			log.Fatalf("Fecth data error: %v", err)
		}

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl := table.New("Source", "PaymentID", "TrxID")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		tbl.AddRow(data.Source, data.PaymentID, data.TrxID)
		tbl.Print()
		tblHistory := table.New("Event Tracking ID", "Event Name")
		tblHistory.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for _, i := range data.History {
			tblHistory.AddRow(i.EventTrackingID, i.EventName)
		}
		tblHistory.Print()
		detail, _ := cmd.Flags().GetString("detail")
		if detail != "" {
			for _, i := range data.History {
				if i.EventTrackingID == detail {
					tblData := table.New("Event Tracking ID", "Event Name")
					tblData.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
					tblData.AddRow(i.EventTrackingID, i.EventName)
					payload, _ := json.MarshalIndent(i.Data, "", "  ")
					tblData.Print()
					fmt.Println(string(payload))
					break
				}

			}
		}
	}
	return search
}
