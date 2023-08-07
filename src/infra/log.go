package infra

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"
)

func InitLogger() (*logrus.Logger, error) {
	cfg := LoadEnvVars()
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
