package digitalocean

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/digitalocean/godo"
)

type Digitalocean struct {
	DropletID []int
	Client    *godo.Client
	Ctx       context.Context
	Consumer  int
	Wg        *sync.WaitGroup
}

type DigitalOceanResponse struct {
	ID     int
	Status string
	Rate   int
}

func NewDigitalOcean() *Digitalocean {

	cl := godo.NewFromToken(os.Getenv("DO_TOKEN"))
	ctx := context.Background()

	fmt.Println("Has Starting connect to digitalocean API")

	return &Digitalocean{
		Client: cl,
		Ctx:    ctx,
	}
}
