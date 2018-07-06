package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/podnov/range-value-broker/pkg"
)

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Lookup("stderrthreshold").Value.Set("INFO")

	server, err := pkg.NewServer()
	if err == nil {
		server.Start()
	} else {
		glog.Errorf("Could not init server: %s", err.Error())
	}
}
