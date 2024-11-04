package flags

import "flag"

const (
	CONFIGS_FLAG                  = "configs"
	CERTS_FLAG                    = "certs"
	LOGGER_OUTPUT_PATH_FLAG       = "logger-output-path"
	LOGGER_ERROR_OUTPUT_PATH_FLAG = "logger-error-output-path"

	CONFIGS_PATH_DEFAULT             = "configs"
	CERTS_PATH_DEFAULT               = "certs"
	LOGGER_OUTPUT_PATH_DEFAULT       = "logs/output.log"
	LOGGER_ERROR_OUTPUT_PATH_DEFAULT = "logs/error.log"
)

type ParsedFlags struct {
	AppMode               *string
	ConfigsPath           *string
	CertsPath             *string
	LoggerOutputPath      *string
	LoggerErrorOutputPath *string
}

type FlagsLoggerConfig struct {
	OutputPath      string
	ErrorOutputPath string
}

type FlagsConfig struct {
	Mode        AppMode
	ConfigsPath string
	CertsPath   string
	Logger      FlagsLoggerConfig
}

func ParseFlags() *ParsedFlags {
	appModeFlag := flag.String(APP_MODE_FLAG, string(AppModeDev), "application mode")
	configsPathFlag := flag.String(CONFIGS_FLAG, CONFIGS_PATH_DEFAULT, "configs path")
	certsPathFlag := flag.String(CERTS_FLAG, CERTS_PATH_DEFAULT, "pool api certificates path")
	loggerOutputPath := flag.String(LOGGER_OUTPUT_PATH_FLAG, LOGGER_OUTPUT_PATH_DEFAULT, "logger output logs file path")
	loggerErrorOutputPath := flag.String(LOGGER_ERROR_OUTPUT_PATH_FLAG, LOGGER_ERROR_OUTPUT_PATH_DEFAULT, "logger output error logs file path")
	parsedFlags := &ParsedFlags{
		AppMode:               appModeFlag,
		ConfigsPath:           configsPathFlag,
		CertsPath:             certsPathFlag,
		LoggerOutputPath:      loggerOutputPath,
		LoggerErrorOutputPath: loggerErrorOutputPath,
	}

	flag.Parse()

	return parsedFlags
}

func SetupFlags(parsedFlags *ParsedFlags) *FlagsConfig {
	appMode := AppModeDev
	configsPath := CONFIGS_PATH_DEFAULT
	certsPath := CERTS_PATH_DEFAULT
	loggerConfig := FlagsLoggerConfig{
		OutputPath:      LOGGER_OUTPUT_PATH_DEFAULT,
		ErrorOutputPath: LOGGER_ERROR_OUTPUT_PATH_DEFAULT,
	}

	if parsedFlags.AppMode != nil {
		appMode = checkAppMode(*parsedFlags.AppMode)
	}

	if parsedFlags.ConfigsPath != nil {
		configsPath = *parsedFlags.ConfigsPath
	}

	if parsedFlags.CertsPath != nil {
		certsPath = *parsedFlags.CertsPath
	}

	if parsedFlags.LoggerOutputPath != nil {
		loggerConfig.OutputPath = *parsedFlags.LoggerOutputPath
	}

	if parsedFlags.LoggerErrorOutputPath != nil {
		loggerConfig.ErrorOutputPath = *parsedFlags.LoggerErrorOutputPath
	}

	return &FlagsConfig{
		Mode:        appMode,
		ConfigsPath: configsPath,
		CertsPath:   certsPath,
		Logger:      loggerConfig,
	}
}
