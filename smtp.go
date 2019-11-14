package main

import (
	"fmt"
	"io/ioutil"
	"github.com/alash3al/go-smtpsrv"
	"time"
	"strings"
	"os"
)

func main() {
	handler := func(req *smtpsrv.Request) error {

		headerStrings := []string{""}
		header := req.Message.Header
		for k := range header {
		    headerStrings = append(headerStrings, fmt.Sprintf("%s: %s\n", k, header[k]) )
		}

		outputString := strings.Join(headerStrings,"")

               fn := "emails/" + time.Now().Format(time.RFC3339Nano) + ".eml"
               ioutil.WriteFile(fn, []byte(outputString), os.ModePerm)
               fmt.Println("Wrote to " + fn)
               return nil
	}

	srv := &smtpsrv.Server{
		Name: "mail.my.server",
		Addr:        ":25025",
		MaxBodySize:  100 * 1024,
		Handler:     handler,
	}
	fmt.Println(srv.ListenAndServe())
}
