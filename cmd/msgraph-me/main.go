package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yaegashi/msgraph.go/auth"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

const (
	defaultTenantID = "common"
	defaultClientID = "45c7f99c-0a94-42ff-a6d8-a8d657229e8c"
	defaultScope    = "openid profile user.readwrite files.readwrite.all"
)

func dump(o interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

func main() {
	var tenantID, clientID string
	flag.StringVar(&tenantID, "tenant_id", defaultTenantID, "Tenant ID")
	flag.StringVar(&clientID, "client_id", defaultClientID, "Client ID")
	flag.Parse()

	da, err := auth.NewDeviceAuth(tenantID, clientID, defaultScope)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(da.Message())

	dt, err := da.Poll()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	cli := dt.Client(ctx)
	serv := msgraph.NewClient(cli)

	{
		r := serv.Me().Request()
		log.Printf("GET %s", r.URL())
		x, err := r.Get()
		if err == nil {
			dump(x)
		} else {
			log.Println(err)
		}
	}

	{
		r := serv.Me().Drive().Root().Children().Request()
		r.Filter("file ne null")
		r.Select("name,file,size,webUrl")
		log.Printf("GET %s", r.URL())
		x, err := r.Get()
		if err == nil {
			dump(x)
		} else {
			log.Println(err)
		}
	}
}
