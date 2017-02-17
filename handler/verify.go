package handler

import (
	log "github.com/cihub/seelog"
	"github.com/HailoOSS/platform/errors"
	"github.com/HailoOSS/platform/server"
	verify "github.com/HailoOSS/oauth-client-service/proto/verify"
	provider "github.com/HailoOSS/oauth-client-service/provider"
	"github.com/HailoOSS/protobuf/proto"
)

// Verify authenticate a Oauth token with the provider to check its validity
func Verify(req *server.Request) (proto.Message, errors.Error) {
	log.Infof("verifying token %+v", req)

	request := &verify.Request{}
	if err := req.Unmarshal(request); err != nil {
		return nil, errors.BadRequest(server.Name+".verify", err.Error())
	}

	log.Debugf("Received token=%v", request.GetToken())
	info, err := provider.CheckProviderToken(request.GetProvider(), request.GetToken())
	if err != nil {
		return nil, errors.BadRequest(server.Name+".verify", err.Error())
	}

	rsp := &verify.Response{
		Valid: proto.Bool(info.IsValid()),
	}

	return rsp, nil
}
