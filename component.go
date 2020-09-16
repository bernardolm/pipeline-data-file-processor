package processor

import (
	"fmt"
	"plugin"

	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"
)

var (
	tryingLoadDebugMessage string = "trying catch %s as %s component from %s\n"
	cantCastErrorMessage   string = "can't cast %s as %s component from %s, this struct probably doesn't implement the correct interface\n"
)

func loadComponent(path, name string) (plugin.Symbol, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	log.Debugf("plugin loaded: %s", pp.Sprint(p))

	s, err := p.Lookup(name)
	if err != nil {
		return nil, err
	}

	log.Debugf("symbol plugin loaded: %s", pp.Sprint(s))

	return s, nil
}

func LoadSource(path, name string) (ISource, error) {
	log.Debugf(tryingLoadDebugMessage, name, "Source", path)

	sym, err := loadComponent(path, name)
	if err != nil {
		return nil, err
	}

	obj, ok := sym.(ISource)
	if !ok {
		return nil, fmt.Errorf(cantCastErrorMessage, name, "Source", path)
	}

	return obj, nil
}

func LoadOutput(path, name string) (IOutput, error) {
	log.Debugf(tryingLoadDebugMessage, name, "Output", path)

	sym, err := loadComponent(path, name)
	if err != nil {
		return nil, err
	}

	obj, ok := sym.(IOutput)
	if !ok {
		return nil, fmt.Errorf(cantCastErrorMessage, name, "Output", path)
	}

	return obj, nil
}
