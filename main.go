package main

import (
	"os"

	dacranepdk "github.com/SIOS-Technology-Inc/dacrane-pdk"
)

func main() {
	dacranepdk.ExecPluginJob(dacranepdk.Plugin{
		Config: dacranepdk.NewDefaultPluginConfig(),
		Resources: dacranepdk.MapToFunc(map[string]dacranepdk.Resource{
			"file": FileResource,
		}),
		Data: dacranepdk.MapToFunc(map[string]dacranepdk.Data{}),
	})
}

var FileResource = dacranepdk.Resource{
	Create: func(parameter any, _ dacranepdk.PluginMeta) (any, error) {
		params := parameter.(map[string]any)
		contents := params["contents"].(string)
		filename := params["filename"].(string)

		e := os.WriteFile(filename, []byte(contents), 0644)
		if e != nil {
			return nil, e
		}

		return parameter, nil
	},
	Delete: func(parameter any, _ dacranepdk.PluginMeta) error {
		params := parameter.(map[string]any)
		filename := params["filename"].(string)
		err := os.Remove(filename)
		if err != nil {
			return err
		}
		return nil
	},
}
