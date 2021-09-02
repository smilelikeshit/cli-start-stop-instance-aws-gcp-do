package cmd

import (
	"fmt"
	"learn1/services/cloudaws"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Command to use provider AWS",
	Long:  `Provider Amazon Web Services`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var instanceAWSCmd = &cobra.Command{
	Use:   "instance",
	Short: "for instance using AWS",
	Long:  `instances for AWS `,
	Run: func(cmd *cobra.Command, args []string) {

		action, _ := cmd.Flags().GetString("action")
		region, _ := cmd.Flags().GetString("region")

		s := &cloudaws.Instance{
			Region:     region,
			InstanceId: args,
		}

		if strings.ToLower(action) == "start" {

			err := s.Start(false)
			if err != nil {
				log.Fatalf(err.Error())
			}
			s.Start(true)

		} else if strings.ToLower(action) == "stop" {

			err := s.Stop(false)
			if err != nil {
				log.Fatalf(err.Error())
			}
			s.Stop(true)

		} else {
			fmt.Println("failed command please read docs")

		}

	},
}

func init() {

	rootCmd.AddCommand(awsCmd)
	awsCmd.AddCommand(instanceAWSCmd)
	awsCmd.PersistentFlags().StringP("action", "a", "", "Possible values: start, stop, dry-run")
	awsCmd.PersistentFlags().StringP("region", "r", "", "Possible values: us-west-1, ap-southeast-1")
	awsCmd.MarkPersistentFlagRequired("action")
	awsCmd.MarkPersistentFlagRequired("region")

}
