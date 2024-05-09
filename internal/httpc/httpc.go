package httpc

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/adzsx/dirsgover/internal/utils"
)

var (
	client http.Client
)

func Status(host string) int {
	resp, err := client.Get(host)

	if err != nil {
		return 404
	}

	return resp.StatusCode
}

func Up(host string) bool {
	client.Timeout = time.Second * 3
	utils.Verbose(1, "Checking status of host...")
	_, err := client.Get(host)

	return err != nil
}

func Robots(host string) {
	bleepbloop := host + "/robots.txt"
	status := Status(bleepbloop)
	if status == 200 {
		resp, _ := http.Get(bleepbloop)

		resBody, err := io.ReadAll(resp.Body)
		body := string(resBody)

		var ent string

		for _, entry := range strings.Split(body, "\n") {
			ent = utils.FilterChar(entry, ":", false)
			if len(ent) > 1 {
				if string(ent[1]) == "/" {

					bleep = append(bleep, ent[2:])
					wg.Add(1)
				}
			}

		}
		utils.Err(err)
	} else if status == 404 {
		log.Fatalln("robots.txt does not exist on this host")
	}
}
