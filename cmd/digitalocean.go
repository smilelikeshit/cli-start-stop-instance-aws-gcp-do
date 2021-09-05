package cmd

import (
	"context"
	"fmt"
	"learn1/services/digitalocean"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var digitaloceanCmd = &cobra.Command{
	Use:   "digitalocean",
	Short: "Command to user provider GCP",
	Long:  `Provider Google Cloud Platform`,
	Run: func(cmd *cobra.Command, args []string) {
		// none
	},
}

var instanceDigitalOceanCmd = &cobra.Command{
	Use:   "instance",
	Short: "for instance using Digitalocean",
	Long:  `instances for Digitalocean `,
	Run: func(cmd *cobra.Command, args []string) {

		action, _ := cmd.Flags().GetString("action")

		var newargs = []int{}
		for _, i := range args {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			newargs = append(newargs, j)
		}

		data := &digitalocean.Digitalocean{
			DropletID: newargs,
			Client:    digitalocean.NewDigitalOcean().Client,
			Ctx:       context.Background(),
			Consumer:  1,
			Wg:        &sync.WaitGroup{},
		}

		if strings.ToLower(action) == "start" {
			data.Start()

		} else if strings.ToLower(action) == "stop" {
			data.Stop()

		} else if strings.ToLower(action) == "status" {
			fmt.Println("Need Contribute :p")

		} else {
			fmt.Println("failed command please read docs")

		}

	},
}

func init() {
	rootCmd.AddCommand(digitaloceanCmd)
	digitaloceanCmd.AddCommand(instanceDigitalOceanCmd)
	digitaloceanCmd.PersistentFlags().StringP("action", "a", "", "Possible values: start, stop, status")
	digitaloceanCmd.MarkPersistentFlagRequired("action")
}
