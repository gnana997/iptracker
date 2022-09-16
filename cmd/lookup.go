package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ipCmd)
}

var ipCmd = &cobra.Command{
  Use:   "iplookup",
  Short: "return the ip of the addr recieved",
  Long:  `Will see the ip of the given host`,
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
		for _, addr := range args {
			res := getIp(addr)
			for i, ip := range res {
				fmt.Printf("%d. %s \n", i, ip)
			}
		}
	} else {
		fmt.Println("please provie IPs to trace")
	}
  },
}

func getIp(addr string) []net.IP {
	ip, err := net.LookupIP(addr)

	if err != nil {
		fmt.Println("Unable to lookup the addr")
	}

	return ip
}