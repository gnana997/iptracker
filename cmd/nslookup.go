package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nsCmd)
}

var nsCmd = &cobra.Command{
  Use:   "nslookup",
  Short: "lookup the addr recieved",
  Long:  `Will see the trace of the nameservers`,
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
		for _, addr := range args {
			res := getNameServers(addr)
			for i, ns := range res {
				fmt.Printf("%d. %s \n", i, ns.Host)
			}
		}
	} else {
		fmt.Println("please provie IPs to trace")
	}
  },
}

func getNameServers(addr string) []*net.NS {
	nsr, err := net.LookupNS(addr)

	if err != nil {
		fmt.Println("Unable to lookup the addr")
	}

	return nsr
}