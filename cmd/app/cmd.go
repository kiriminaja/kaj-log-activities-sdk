package app

import (
	"log"

	"github.com/kiriminaja/kaj-log-activities-sdk/cmd/event"
	"github.com/kiriminaja/kaj-log-activities-sdk/cmd/packages"
	kaj_log_activities_sdk "github.com/kiriminaja/kaj-log-activities-sdk/sdk"
	"github.com/spf13/cobra"
)

func Start(client *kaj_log_activities_sdk.Client) {
	rootCmd := &cobra.Command{}
	// ctx, cancel := context.WithCancel(context.Background())
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	// go func() {
	// 	<-quit
	// 	cancel()
	// }()

	cmd := []*cobra.Command{
		event.MainCommand(),
		packages.MainCommand(client),
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
