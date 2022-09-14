package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "trace",
  Short: "Trace the Ip",
  Long:  `Will see the trace of the ip`,
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
		for _, ip := range args {
			res := getData(ip)
			fmt.Println(res)
		}
	} else {
		fmt.Println("please provie IPs to trace")
	}
  },
}

type IP struct {
	Ip string `json::"ip"`
	HostName string `json::"hostname"`
	City string `json::"city"`
	Region string `json::"region"`
	Country string `json::"country"`
	GeoCoordinates string `json::"loc"`
	TimeZone string `json::"timezone"`
}

func getData(ip string) IP {
	url := "http://ipinfo.io/" + ip + "/geo"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Unable to call the api")
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Unable to read the api")
	}

	var res IP
	json.Unmarshal(responseBytes, &res)

	return res
}