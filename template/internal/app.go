package internal

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

type (
	flags struct {
		configPath *string
	}
)

// runHandler define any custom logic in run handler
var runHandler = func(fv flags, logger *zerolog.Logger) {}

var metricsAddr = ":8081"

func Run() {
	var fv = initFlags()
	var logger = initLogger()
	go initMetrics(&logger)

	runHandler(fv, &logger)
}

func initFlags() (fv flags) {
	fv.configPath = flag.String("config", "config.yaml", "path to config.yml")

	flag.Parse()
	return fv
}

func initLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func initMetrics(logger *zerolog.Logger) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(metricsAddr, nil); err != nil {
		logger.Error().Msgf("metrics not served")
	}
}
