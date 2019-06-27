package main

import (
	"github.com/narrowizard/nirvana-cms/services"

	"github.com/narrowizard/nirvana-cms/api"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/log"
)

func main() {
	var config = nirvana.NewDefaultConfig()
	config.Configure(nirvana.Descriptor(api.User))
	config.Configure(nirvana.Descriptor(api.Menu))
	config.Configure(nirvana.Descriptor(api.Role))
	config.Configure(nirvana.Port(services.ConfigInfo().Port))
	log.Infof("listening on %s:%d", config.IP(), config.Port())
	if err := nirvana.NewServer(config).Serve(); err != nil {
		log.Fatal(err)
	}
}
