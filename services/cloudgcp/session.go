package cloudgcp

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	google "golang.org/x/oauth2/google"
)

type Gcs struct {
	Instance []string
	Gclient  *http.Client
	Project  string
	Zone     string
	Ctx      context.Context
}

func NewGCS(ctx context.Context) *Gcs {

	data, err := ioutil.ReadFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		log.Fatal(err)
	}

	// https://cloud.google.com/compute/docs/api/how-tos/authorization
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/compute")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx)

	return &Gcs{
		Gclient: client,
		Ctx:     ctx,
	}

}
