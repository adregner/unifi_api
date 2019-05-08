package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/bah2830/unifi_api/pkg/client/operations"

	"github.com/bah2830/unifi_api"
)

var (
	host = flag.String("host", "192.168.1.10:8443", "host and port of the unifi controller")
	user = flag.String("username", "admin", "username for the unifi controller")
	pass = flag.String("password", "password", "password for the unifi controller")
)

func main() {
	flag.Parse()

	c, err := unifi_api.NewClient(*host, *user, *pass)
	if err != nil {
		panic(err)
	}

	r, err := c.Operations.SitesList(&operations.SitesListParams{
		Context: context.Background(),
	})
	if err != nil {
		panic(err)
	}

	for _, site := range r.Payload.Data {
		fmt.Printf("%s(%s)\n", site.Name, site.Desc)
	}
}
