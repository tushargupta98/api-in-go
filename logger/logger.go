package logger

import (
	"os"
	"sync"

	"github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"

	"github.com/tushargupta98/api-in-go/config"
)

var (
	once   sync.Once
	Logger *log.Logger
)

func InitNewRelic(newRelicConfig config.NewRelicConfig) *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(newRelicConfig.AppName),
		newrelic.ConfigLicense(newRelicConfig.LicenseKey),
	)
	if err != nil {
		log.Warn("Unable to create New Relic application, logging to file instead: ", err)
		return nil
	}
	return app
}

func InitLogger(config config.Config) *log.Logger {
	level, err := log.ParseLevel(config.Logger.Level)
	if err != nil {
		log.Fatal("Error parsing log level: ", err)
	}
	logger := log.New()
	logger.SetLevel(level)

	var output *os.File
	if config.Logger.Output.Type == "file" {
		var err error
		output, err = os.OpenFile(config.Logger.Output.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("Error opening log file: ", err)
		}
	} else {
		output = os.Stdout
	}
	logger.SetOutput(output)

	var formatter log.Formatter
	if config.Logger.Formatter.Type == "text" {
		formatter = &log.TextFormatter{
			DisableTimestamp: config.Logger.Formatter.DisableTimestamp,
			FullTimestamp:    config.Logger.Formatter.FullTimestamp,
			TimestampFormat:  config.Logger.Formatter.TimestampFormat,
		}
	} else if config.Logger.Formatter.Type == "json" {
		formatter = &log.JSONFormatter{}
	} else {
		log.Fatal("Unknown formatter type: ", config.Logger.Formatter.Type)
	}
	logger.SetFormatter(formatter)

	// Hooking up logger with New Relic
	/*	app := InitNewRelic(config.NewRelic)
		if app != nil {
			logger.AddHook(&NewRelicHook{app})
		}
	*/

	return logger
}

type NewRelicHook struct {
	app *newrelic.Application
}

func (h *NewRelicHook) Levels() []log.Level {
	return []log.Level{log.PanicLevel, log.FatalLevel, log.ErrorLevel, log.WarnLevel}
}

func (h *NewRelicHook) Fire(entry *log.Entry) error {
	txn := h.app.StartTransaction(entry.Message)
	defer txn.End()
	txn.AddAttribute("level", entry.Level.String())
	txn.AddAttribute("fields", entry.Data)
	return nil
}

func init() {
	once.Do(func() {
		Logger = InitLogger(*config.GetConfig())
	})
}

func WithFields(fields log.Fields) *log.Entry {
	return Logger.WithFields(fields)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}
