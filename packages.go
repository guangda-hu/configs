package packages

import (
	// Used by Sublime plugin https://github.com/liuhewei/gotools-sublime
	_ "github.com/mdempsky/gocode"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/cmd/gorename"
	_ "golang.org/x/tools/cmd/guru"

	// OAuth2l
	_ "github.com/google/oauth2l"
)
