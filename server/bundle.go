package server

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
)

func Bundle(key string) error {

	result := api.Build(api.BuildOptions{
		NodePaths:         []string{"node_modules"},
		EntryPoints:       []string{fmt.Sprintf("code/%s/index.js", key)},
		Bundle:            true,
		Format:            api.FormatIIFE,
		Outfile:           fmt.Sprintf("js/%s.js", key),
		Write:             true,
		GlobalName:        "global",
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		//External:          []string{"node_modules"},
	})

	if len(result.Errors) != 0 {
		return fmt.Errorf(result.Errors[0].Text)
	}

	return nil
}
