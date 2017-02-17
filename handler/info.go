package handler

import (
	log "github.com/cihub/seelog"
	"github.com/HailoOSS/platform/errors"
	"github.com/HailoOSS/platform/server"
	info "github.com/HailoOSS/oauth-client-service/proto/info"
	provider "github.com/HailoOSS/oauth-client-service/provider"
	"github.com/HailoOSS/protobuf/proto"
)

// Info gets information about a Oauth token user
func Info(req *server.Request) (proto.Message, errors.Error) {
	log.Infof("user info %+v", req)

	request := &info.Request{}
	if err := req.Unmarshal(request); err != nil {
		return nil, errors.BadRequest(server.Name+".info", err.Error())
	}

	log.Debugf("Received token=%v", request.GetToken())
	user, err := provider.TokenUserInfo(request.GetProvider(), request.GetToken())
	if err != nil {
		return nil, errors.BadRequest(server.Name+".info", err.Error())
	}

	rsp := &info.Response{
		Email:      proto.String(user.Email()),
		GivenName:  proto.String(user.GivenName()),
		FamilyName: proto.String(user.FamilyName()),
	}

	return rsp, nil
}
