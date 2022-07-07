package mpopenvidu

import (
	"flag"
	"fmt"
	"os"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

type OpenViduPlugin struct {
	prefix      string
	endpointUrl string
	secret      string
}

func Do() {
	var (
		optPrefix      = flag.String("metric-key-prefix", "", "Metric key prefix")
		optEndpointUrl = flag.String("endpoint-url", "", "API Endpoint URL")
		optSecret      = flag.String("secret", "", "OpenVidu Secret")
	)
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage: %s", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	mp.NewMackerelPlugin(&OpenViduPlugin{
		prefix:      *optPrefix,
		endpointUrl: *optEndpointUrl,
		secret:      *optSecret,
	}).Run()
}
