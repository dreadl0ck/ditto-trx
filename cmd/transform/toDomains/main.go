package main

import (
	"fmt"
	"github.com/dreadl0ck/maltego"
	"log"
	"os"
)

// This is an example for a local transformation that does a reverse name lookup for a given address.
// It will take an IP address and return the hostnames associated with it as maltego entities.
func main() {

	log.Println(os.Args[1:])

	// parse arguments
	lt := maltego.ParseLocalArguments(os.Args[1:])
	names := lt.Values["names"]

	// create new transform
	t := maltego.Transform{}

	if names != "" {
		t.AddEntity(maltego.Company, names)
	}

	t.AddUIMessage("complete", maltego.UIMessageInform)

	// return output to stdout and exit cleanly (exit code 0)
	fmt.Println(t.ReturnOutput())
}

