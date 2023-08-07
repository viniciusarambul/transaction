package infra

import (
	"os"

	"github.com/sirupsen/logrus"
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

	return log, nil
}
