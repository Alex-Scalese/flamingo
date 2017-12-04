package locale

import (
	"go.aoe.com/flamingo/core/locale/interfaces/templatefunctions"
	"go.aoe.com/flamingo/framework/config"
	"go.aoe.com/flamingo/framework/dingo"
	"go.aoe.com/flamingo/framework/template"
)

type (
	// Module registers our profiler
	Module struct {
	}
)

// Configure the product URL
func (m *Module) Configure(injector *dingo.Injector) {

	injector.BindMulti((*template.Function)(nil)).To(templatefunctions.Label{})
	injector.BindMulti((*template.Function)(nil)).To(templatefunctions.PriceFormatFunc{})
	injector.BindMulti((*template.Function)(nil)).To(templatefunctions.DateTime{})
}

// DefaultConfig for this module
func (m *Module) DefaultConfig() config.Map {
	return config.Map{
		"locale.translationFile": "translations/en-US.all.json",
		"locale.locale":          "en-US",
		"locale.accounting": config.Map{
			"decimal":    ",",
			"thousand":   ".",
			"formatZero": "%s -,-",
			"format":     "%s %v",
		},
		"locale.date": config.Map{
			"dateFormat":     "02 Jan 2006",
			"timeFormat":     "15:04:05",
			"dateTimeFormat": "02 Jan 2006 15:04:05",
			"location":       "Europe/London",
		},
	}
}
