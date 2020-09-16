package main

import (
	processor "github.com/bernardolm/pipelinedatafilprocessor"
	"github.com/bernardolm/pipelinedatafilprocessor/config"
	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	configs := config.Load()
	pp.Println(configs)

	for _, v1 := range configs {
		for _, v2 := range v1.Sources {
			if obj, err := processor.LoadSource(v1.PluginPath, v2.Name); err != nil {
				log.Error(err)
				continue
			} else {
				pp.Println(obj)
			}
		}

		for _, v2 := range v1.Outputs {
			if obj, err := processor.LoadOutput(v1.PluginPath, v2.Name); err != nil {
				log.Error(err)
				continue
			} else {
				pp.Println(obj)
			}
		}
	}
}
