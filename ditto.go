/*
 * DITTO-TRX - A maltego transform server for IDN homograph attacks
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
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dreadl0ck/cryptoutils"
	"github.com/dreadl0ck/maltego"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const storage = "/tmp"

var ditto = func(status string, hasIP bool, args ...string) http.HandlerFunc {

	return maltego.MakeHandler(func(w http.ResponseWriter, r *http.Request, t *maltego.Transform) {

		// get host that was queried
		host := t.RequestMessage.Entities.Items[0].Value

		fmt.Println("got request from", r.RemoteAddr, "to lookup:", host)

		if !govalidator.IsDNSName(host) {
			t.AddUIMessage("invalid domain: "+host, maltego.UIMessageFatal)
			return
		}

		id, err := cryptoutils.RandomString(10)
		if err != nil {
			fmt.Println("failed to generate id:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		file := filepath.Join(storage, id+"-"+host)

		finalArgs := []string{"-domain", host, "-no-progress-bar", "-csv", file, "-throttle=1000", "-workers=4", "-whois"}
		for _, a := range args {
			finalArgs = append(finalArgs, a)
		}

		start := time.Now()

		// run ditto
		// we are running inside a docker container, the ditto binary has been copied into it at build time.
		// TODO: drop privileges
		out, err := exec.Command("/root/ditto",  finalArgs...).CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println("failed to run ditto:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("failed to read file:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rd := csv.NewReader(bytes.NewReader(data))
		records, err := rd.ReadAll()
		if err != nil {
			fmt.Println("failed to read CSV records:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("results for", host, "=", len(records), "in", time.Since(start))

		// process results
		for i, rec := range records {
			//fmt.Println(rec)
			if i == 0 {
				// skip CSV header
				continue
			}

			// handle status if provided
			if status != "" {
				if rec[2] != status {
					continue
				}
			}

			if hasIP {
				if rec[3] == "" {
					continue
				}
			}

			addEntity(t, rec)
		}

		_ = os.Remove(file)
	})
}

func addEntity(t *maltego.Transform, rec []string) {
	e := t.AddEntity("dittotrx.IDNDomain", rec[0])
	e.AddProp("unicode", rec[0])
	e.AddProp("ascii", rec[1])
	e.AddProp("status", rec[2])
	e.AddProp("ips", rec[3])
	e.AddProp("names", rec[4])

	e.AddProp("registrar", rec[5])
	e.AddProp("created_at", rec[6])
	e.AddProp("updated_at", rec[7])
	e.AddProp("expires_at", rec[8])
	e.AddProp("nameservers", rec[9])
}
