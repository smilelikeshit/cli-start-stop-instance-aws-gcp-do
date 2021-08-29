package cloudgcp

import (
	"fmt"
	"log"

	compute "google.golang.org/api/compute/v1"
)

func (g *Gcs) Start() (err error) {

	newservice, err := compute.NewService(g.Ctx)
	if err != nil {
		log.Fatal(err)
	}

	for c := 0; c < len(g.Instance); c++ {
		resp, err := newservice.Instances.Start(g.Project, g.Zone, g.Instance[c]).Context(g.Ctx).Do()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" Instance ID 	: %d\n", resp.TargetId)
		fmt.Printf(" Status    		: %s\n", resp.Status)
		fmt.Printf(" Progress 		: %d\n", resp.Progress)
		fmt.Printf(" Region		: %s\n", resp.Region)
	}

	return
}

func (g *Gcs) Stop() (err error) {

	newservice, err := compute.NewService(g.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	for c := 0; c < len(g.Instance); c++ {
		resp, err := newservice.Instances.Stop(g.Project, g.Zone, g.Instance[c]).Context(g.Ctx).Do()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" Instance ID 	: %d\n", resp.TargetId)
		fmt.Printf(" Status    		: %s\n", resp.Status)
		fmt.Printf(" Progress 		: %d\n", resp.Progress)
	}

	return
}

func (g *Gcs) List() (err error) {

	newservice, err := compute.NewService(g.Ctx)
	if err != nil {
		log.Fatal(err)
	}

	req := newservice.Instances.List(g.Project, g.Zone)
	if err := req.Pages(g.Ctx, func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			// TODO: Change code below to process each `instance` resource:
			fmt.Printf("Name	 	: %s\n", instance.Name)
			fmt.Printf("Instance ID	: %d\n", instance.Id)
			fmt.Printf("Zone		: %s\n", instance.Zone)
			fmt.Printf("Status		: %s\n", instance.Status)
			fmt.Printf("======================================\n")
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return
}
