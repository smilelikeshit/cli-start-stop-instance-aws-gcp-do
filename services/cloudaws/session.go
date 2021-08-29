package cloudaws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Instance struct {
	InstanceId []string
	Sess       *session.Session
	Svc        *ec2.EC2
	Region     string
}

func NewSession(region string) *Instance {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	svc := ec2.New(sess)

	return &Instance{
		Sess:   sess,
		Svc:    svc,
		Region: region,
	}

}
