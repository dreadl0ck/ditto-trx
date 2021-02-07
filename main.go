/*
 * MALTEGO - Go package that provides datastructures for interacting with the Maltego graphical link analysis tool.
 * Copyright (c) 2021 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"flag"
	"fmt"
	"github.com/dreadl0ck/maltego"
	"log"
	"net/http"
)

var (
	flagAddr   = flag.String("addr", ":8081", "server listen address")
)

func main() {

	flag.Parse()

	const (
		// ditto status names
		registered = "registered"
		available = "available"
	)

	// register transforms to http.DefaultServeMux

	// all similar domains
	maltego.RegisterTransform(ditto("", false), "similarDomains")

	// only registered domains
	maltego.RegisterTransform(ditto(registered, false), "registeredDomains")

	// only registered domains that resolve to an IP
	maltego.RegisterTransform(ditto(registered, true), "liveDomains")

	// only show domains that are available and not registered
	maltego.RegisterTransform(ditto(available, false), "availableDomains")

	// only live domains that resolve for all TLDs
	maltego.RegisterTransform(ditto(registered, true, "-tld"), "liveDomainsTLD")

	// register catch all handler to serve home page
	http.HandleFunc("/", maltego.Home)

	fmt.Println("serving at", *flagAddr)

	s := &http.Server{
		Addr:              *flagAddr,
		Handler:           http.DefaultServeMux,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
	}

	// start server
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("failed to serve HTTP: ", err)
	}
}
