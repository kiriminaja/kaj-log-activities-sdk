package event

import (
	"os"

	"github.com/spf13/cobra"
)

func MainCommand() *cobra.Command {
	event := &cobra.Command{}
	event.Flags().StringP("port", "p", os.Getenv("SERVER_PORT"), "Add port")
	event.Use = "event"
	event.Short = "Event "
	event.AddCommand(upsert())
	event.AddCommand(list())
	return event
}

func upsert() *cobra.Command {
	upsert := &cobra.Command{}
	upsert.Use = "upsert"
	upsert.Short = "Upsert "
	upsert.Run = func(cmd *cobra.Command, args []string) {

	}
	return upsert
}

func list() *cobra.Command {
	list := &cobra.Command{}
	list.Use = "list"
	list.Short = "List "
	list.Run = func(cmd *cobra.Command, args []string) {

	}
	return list
}
