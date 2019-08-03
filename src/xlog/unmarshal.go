package xlog

import (
	"config"

	"github.com/go-kit/kit/log"
	"go.uber.org/fx"
)

// Unmarshal loads an Options from a Viper instance and produces a go-kit Logger
func Unmarshal(key string, u config.KeyUnmarshaller) (log.Logger, error) {
	var o Options
	if err := u.UnmarshalKey(key, &o); err != nil {
		return nil, err
	}

	return New(o)
}

// Unmarshaller produces an optioner strategy that loads the logger from configuration and
// emits it as an uber/fx component.  If supplied, the pf closure is used to construct
// an fx.Printer which may use the created logger.  If pf is nil, no fx.Printer is configured.
func Unmarshaller(key string, pf func(log.Logger, config.Environment) fx.Printer) config.Optioner {
	return func(e config.Environment) fx.Option {
		logger, err := Unmarshal(key, e.KeyUnmarshaller)
		if logger == nil {
			logger = Default()
		}

		options := []fx.Option{
			fx.Provide(
				func() (log.Logger, error) { return logger, err },
			),
		}

		if pf != nil {
			options = append(options,
				fx.Logger(pf(logger, e)),
			)
		}

		return fx.Options(options...)
	}
}
