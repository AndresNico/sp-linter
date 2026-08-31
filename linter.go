package linter

import (
	"github.com/AndresNico/sp-linter/analyzers/annotation"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("serviceprovider_linter", New)
}

type Settings struct {
	// Disable lists analyzer names to skip
	Disable []string `json:"disable"`
}

type ServiceProviderLinter struct {
	settings Settings
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}
	return &ServiceProviderLinter{settings: s}, nil
}

func (f *ServiceProviderLinter) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	all := []*analysis.Analyzer{
		annotation.Analyzer,
	}

	disabled := make(map[string]bool, len(f.settings.Disable))
	for _, name := range f.settings.Disable {
		disabled[name] = true
	}

	enabled := make([]*analysis.Analyzer, 0, len(all))
	for _, a := range all {
		if !disabled[a.Name] {
			enabled = append(enabled, a)
		}
	}
	return enabled, nil
}

func (f *ServiceProviderLinter) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
