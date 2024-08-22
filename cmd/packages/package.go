package packages

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	kaj_log_activities_sdk "github.com/kiriminaja/kaj-log-activities-sdk/sdk"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func MainCommand(client *kaj_log_activities_sdk.Client) *cobra.Command {
	event := &cobra.Command{}
	event.Use = "package"
	event.Short = "package [command]"
	event.AddCommand(search(client))
	return event
}

func search(client *kaj_log_activities_sdk.Client) *cobra.Command {
	// client.Packages.Search()
	search := &cobra.Command{}
	search.Flags().StringP("search", "s", "", "Package ID | OrderID")
	search.Use = "search"
	search.Short = "search"
	search.Run = func(cmd *cobra.Command, args []string) {
		search, err := cmd.Flags().GetString("search")
		if err != nil {
			log.Fatalf("Argument search: %v", err)
		}
		data, err := client.Packages.Search(search)
		if err != nil {
			log.Fatalf("Fecth data error: %v", err)
		}

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl := table.New("Source", "PackageID", "OrderID")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		tbl.AddRow(data.Source, data.PackageID, data.OrderID)
		tbl.Print()
		fmt.Println()
		tblHistory := table.New("Event Tracking ID", "Event Name")
		tblHistory.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for _, i := range data.History {
			tblHistory.AddRow(i.EventTrackingID, i.EventName)
		}
		tblHistory.Print()
	}
	return search
}
