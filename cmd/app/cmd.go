package app

import (
	"log"

	"github.com/kiriminaja/kaj-log-activities-sdk/cmd/event"
	"github.com/kiriminaja/kaj-log-activities-sdk/cmd/packages"
	"github.com/kiriminaja/kaj-log-activities-sdk/cmd/payments"
	"github.com/kiriminaja/kaj-log-activities-sdk/cmd/withdrawals"
	kaj_log_activities_sdk "github.com/kiriminaja/kaj-log-activities-sdk/sdk"
	"github.com/spf13/cobra"
)

func Start(client *kaj_log_activities_sdk.Client) {
	rootCmd := &cobra.Command{}
	cmd := []*cobra.Command{
		event.MainCommand(),
		packages.MainCommand(client),
		payments.MainCommand(client),
		withdrawals.MainCommand(client),
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
