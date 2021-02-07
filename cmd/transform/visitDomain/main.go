package main

import (
	"fmt"
	"github.com/dreadl0ck/maltego"
	"log"
	"os"
	"os/exec"
	"github.com/asaskevich/govalidator"
)

func main() {
	lt := maltego.ParseLocalArguments(os.Args)

	log.Println(lt.Values, os.Args)
	domain := lt.Values["ascii"]

	if !govalidator.IsDNSName(domain) {
		maltego.Die(domain, "is not a valid domain")
	}

	log.Println("open", domain)

	err := exec.Command("open", "http://" + domain).Run()
	if err != nil {
		maltego.Die(err.Error(), "failed to open domain")
	}

	tr := maltego.Transform{}
	tr.AddUIMessage("done", maltego.UIMessageInform)
	fmt.Println(tr.ReturnOutput())
}
