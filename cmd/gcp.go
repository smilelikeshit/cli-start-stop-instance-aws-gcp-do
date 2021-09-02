package cmd

import (
	"context"
	"fmt"
	"learn1/services/cloudgcp"
	"log"
	"strings"

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
			Project:  "<PLEASE_PUT_PROJECT_ID_HERE>",
			Ctx:      context.Background(),
		}

		if strings.ToLower(action) == "start" {

			err := gc.Start()
			if err != nil {
				log.Fatalf(err.Error())
			}

		} else if strings.ToLower(action) == "stop" {

			err := gc.Stop()
			if err != nil {
				log.Fatalf(err.Error())
			}

		} else if strings.ToLower(action) == "list" {
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
