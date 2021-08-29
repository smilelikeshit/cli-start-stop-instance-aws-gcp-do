package cloudaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (i *Instance) Start(dryrun bool) (err error) {

	for c := 0; c < len(i.InstanceId); c++ {
		input := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String(i.InstanceId[c]),
			},
			DryRun: aws.Bool(dryrun),
		}

		result, err := NewSession(i.Region).Svc.StartInstancesWithContext(context.Background(), input)

		if err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Printf("%s", result.StartingInstances)
		}
	}
	return
}

func (i *Instance) Stop(dryrun bool) (err error) {
	for c := 0; c < len(i.InstanceId); c++ {
		input := &ec2.StopInstancesInput{
			InstanceIds: []*string{
				aws.String(i.InstanceId[c]),
			},
			DryRun: aws.Bool(dryrun),
		}

		result, err := NewSession(i.Region).Svc.StopInstancesWithContext(context.Background(), input)

		if err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Printf("%s", result.StoppingInstances)
		}
	}
	return
}
