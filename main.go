package main

import (
	log "github.com/cihub/seelog"

	"github.com/HailoOSS/oauth-client-service/handler"
	service "github.com/HailoOSS/platform/server"
)

func main() {
	defer log.Flush()

	service.Name = "com.HailoOSS.service.oauth-client"
	service.Description = "Check token validtity and getting information about user when using Oauth"
	service.Version = ServiceVersion
	service.Source = "github.com/HailoOSS/oauth-client"

	service.Init()

	service.Register(&service.Endpoint{
		Name:       "verify",
		Mean:       300,
		Upper95:    600,
		Handler:    handler.Verify,
		Authoriser: service.OpenToTheWorldAuthoriser(),
	})

	service.Register(&service.Endpoint{
		Name:       "info",
		Mean:       300,
		Upper95:    600,
		Handler:    handler.Info,
		Authoriser: service.OpenToTheWorldAuthoriser(),
	})

	service.Run()
}
