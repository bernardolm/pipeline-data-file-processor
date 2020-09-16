module github.com/bernardolm/pipelinedatafilprocessor

go 1.13

replace github.com/bernardolm/pipelinedatafilprocessor => ../pipeline-data-file-processor

require (
	github.com/google/martian v2.1.0+incompatible
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/text v0.3.3
	google.golang.org/appengine v1.6.6
	gopkg.in/yaml.v2 v2.3.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
