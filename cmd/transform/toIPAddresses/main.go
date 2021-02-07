package main

import (
	"fmt"
	"github.com/dreadl0ck/maltego"
	"log"
	"os"
	"strings"
)

// This is an example for a local transformation that does a reverse name lookup for a given address.
// It will take an IP address and return the hostnames associated with it as maltego entities.
func main() {

	log.Println(os.Args[1:])

	// parse arguments
	lt := maltego.ParseLocalArguments(os.Args[1:])

	log.Println(lt.Values)
	ips := lt.Values["ips"]

	// create new transform
	t := maltego.Transform{}

	if ips != "" {
		for _, ip := range strings.Split(ips, ",") {
			t.AddEntity(maltego.IPv4Address, ip)
		}
	}

	t.AddUIMessage("complete", maltego.UIMessageInform)

	// return output to stdout and exit cleanly (exit code 0)
	fmt.Println(t.ReturnOutput())
}

