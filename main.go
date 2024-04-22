package main

import (
	"github.com/go-oauth2/oauth2/manage"
	"github.com/go-oauth2/oauth2/store"
)

func main() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())
}
