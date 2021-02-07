package main

import (
	"fmt"
	"github.com/dreadl0ck/maltego"
	"log"
	"os"
	"time"
)

// This is an example for a local transformation that does a reverse name lookup for a given address.
// It will take an IP address and return the hostnames associated with it as maltego entities.
func main() {

	log.Println(os.Args[1:])

	// parse arguments
	lt := maltego.ParseLocalArguments(os.Args[1:])

	log.Println(lt.Values)
	date := lt.Values["created_at"]

	ti, err := time.Parse(time.RFC3339, date)
	if err != nil {
		maltego.Die(err.Error(), "invalid date")
	}

	// create new transform
	t := maltego.Transform{}

	y, m, d := ti.Date()
	t.AddEntity(maltego.DateTime, fmt.Sprintf("%d %s %d", d, m, y))

	t.AddUIMessage("complete", maltego.UIMessageInform)

	// return output to stdout and exit cleanly (exit code 0)
	fmt.Println(t.ReturnOutput())
}

