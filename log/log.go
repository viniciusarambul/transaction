package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"
	"github.com/viniciusarambul/transaction/config"
)

func InitLogger() (*logrus.Logger, error) {
	cfg := config.LoadEnvVars()
	log := logrus.New()

	log.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	log.SetLevel(level)

	// Instrument logrus.
	log.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	)))

	return log, nil
}
