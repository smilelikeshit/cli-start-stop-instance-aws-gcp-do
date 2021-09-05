package digitalocean

import (
	"encoding/json"
	"fmt"
)

// input channel from struct
func InputChan(droplet []int) <-chan int {

	ch := make(chan int)

	go func() {
		for c := range droplet {
			ch <- droplet[c]
		}

		close(ch)
	}()
	return ch
}

// stop
func (d *Digitalocean) Stop() {

	for i := 0; i < d.Consumer; i++ {
		d.Wg.Add(1)
		go func() {

			for val := range d.DropletID {

				action, response, err := d.Client.DropletActions.PowerOff(d.Ctx, d.DropletID[val])
				if err != nil {
					fmt.Printf("Something bad happened: %s\n\n", err)
				}

				resp := &DigitalOceanResponse{action.ID, action.Status, response.Rate.Remaining}
				js, err := json.Marshal(&resp)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(js))
			}

			d.Wg.Done()
		}()
	}
	d.Wg.Wait()
}

// start
func (d *Digitalocean) Start() {

	for i := 0; i < d.Consumer; i++ {
		d.Wg.Add(1)

		go func() {

			for val := range d.DropletID {

				action, response, err := d.Client.DropletActions.PowerOn(d.Ctx, d.DropletID[val])
				if err != nil {
					fmt.Printf("Something bad happened: %s\n\n", err)
				}

				resp := &DigitalOceanResponse{action.ID, action.Status, response.Rate.Remaining}
				js, err := json.Marshal(&resp)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(string(js))

			}

			d.Wg.Done()
		}()
	}

	d.Wg.Wait()

}
