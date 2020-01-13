package server

import (
	"Blog/cmd/server/hook"
	"net/http"

	"github.com/bilibili/twirp"
)

var hooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewLog(),
)

var loginHooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewCheckLogin(),
	hook.NewLog(),
)

func initMux(mux *http.ServeMux) {
}

func initInternalMux(mux *http.ServeMux) {
}
