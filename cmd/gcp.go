package cmd

import (
	"context"
	"fmt"
	"learn1/services/cloudgcp"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var gcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Command to user provider GCP",
	Long:  `Provider Google Cloud Platform`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var instanceGCPCmd = &cobra.Command{
	Use:   "instance",
	Short: "for instance using GCP",
	Long:  `instances for GCP `,
	Run: func(cmd *cobra.Command, args []string) {

		action, _ := cmd.Flags().GetString("action")
		zone, _ := cmd.Flags().GetString("region")

		gc := &cloudgcp.Gcs{
			Instance: args,
			Zone:     zone,
			Project:  os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
			Ctx:      context.Background(),
		}

		if action == "start" || action == "START" {

			err := gc.Start()
			if err != nil {
				log.Fatalf(err.Error())
			}

		} else if action == "stop" || action == "STOP" {

			err := gc.Stop()
			if err != nil {
				log.Fatalf(err.Error())
			}

		} else if action == "list" || action == "LIST" {
			err := gc.List()
			if err != nil {
				log.Fatalf(err.Error())
			}

		} else {
			fmt.Println("failed command please read docs")

		}

	},
}

func init() {
	rootCmd.AddCommand(gcpCmd)
	gcpCmd.AddCommand(instanceGCPCmd)
	gcpCmd.PersistentFlags().StringP("action", "a", "", "Possible values: start, stop, list")
	gcpCmd.PersistentFlags().StringP("region", "r", "", "Possible values: asia-southeast1-a")
	gcpCmd.MarkPersistentFlagRequired("action")
	gcpCmd.MarkPersistentFlagRequired("region")
}
