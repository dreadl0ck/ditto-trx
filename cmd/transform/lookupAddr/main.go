package main

import (
	"fmt"
	"github.com/dreadl0ck/maltego"
	"log"
	"net"
	"os"
)

// This is an example for a local transformation that does a reverse name lookup for a given address.
// It will take an IP address and return the hostnames associated with it as maltego entities.
func main() {

	log.Println(os.Args[1:])

	// parse arguments
	lt := maltego.ParseLocalArguments(os.Args[1:])

	// ensure the provided address is valid
	ip := net.ParseIP(lt.Value)
	if ip == nil {
		maltego.Die("invalid ip", lt.Value+" is not a valid IP address")
	}

	// lookup provided ip address
	names, err := net.LookupAddr(lt.Value)
	if err != nil {
		maltego.Die(err.Error(), "failed to lookup address")
	}

	// create new transform
	t := maltego.Transform{}

	// iterate over lookup results
	for _, host := range names {
		e := t.AddEntity(maltego.DNSName, host)
		e.AddProperty("hostname", "Hostname", maltego.Strict, host)
	}

	t.AddUIMessage("complete", maltego.UIMessageInform)

	// return output to stdout and exit cleanly (exit code 0)
	fmt.Println(t.ReturnOutput())
}

