package config

import "time"

type Config struct {
	Database      DatabaseConfig `yaml:"database"`
	Cache         CacheConfig    `yaml:"cache"`
	Server        ServerConfig   `yaml:"server"`
	NewRelic      NewRelicConfig `yaml:"newrelic"`
	Logger        LoggerConfig   `yaml:"logger"`
	SwaggerConfig SwaggerConfig  `yaml:"swagger"`
}

type DatabaseConfig struct {
	Driver          string `yaml:"driver"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	Name            string `yaml:"name"`
	Schema          string `yaml:"schema"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	SSLMode         string `yaml:"sslmode"`
	SSLRootCert     string `yaml:"sslrootcert"`
	MaxOpenConns    int    `yaml:"maxopenconns"`
	MaxIdleConns    int    `yaml:"maxidleconns"`
	MaxConnLifetime int    `yaml:"maxconnlifetime"`
}

type CacheConfig struct {
	TLSEnabled   bool          `yaml:"tlsenabled"`
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	Password     string        `yaml:"password"`
	DB           string        `yaml:"db"`
	PoolSize     int           `yaml:"pool_size"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
	PoolTimeout  time.Duration `yaml:"pool_timeout"`
	DialTimeout  time.Duration `yaml:"dial_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	Expiration   time.Duration `yaml:"expiration"`
	MaxRetries   int           `yaml:"max_retries"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type NewRelicConfig struct {
	LicenseKey string `yaml:"license_key"`
	AppName    string `yaml:"app_name"`
}

type LoggerConfig struct {
	Level     string `yaml:"level"`
	Formatter struct {
		Type             string `yaml:"type"`
		DisableTimestamp bool   `yaml:"disable_timestamp"`
		FullTimestamp    bool   `yaml:"full_timestamp"`
		TimestampFormat  string `yaml:"timestamp_format"`
	} `yaml:"formatter"`
	Output struct {
		Type string `yaml:"type"`
		Path string `yaml:"path"`
	} `yaml:"output"`
}

type SwaggerConfig struct {
	SwaggerUiPath string `yaml:"swagger_ui_path"`
	JsonPath      string `yaml:"json_path"`
	DocsPath      string `yaml:"docs_path"`
	StaticDir     string `yaml:"static_dir"`
	BasePath      string `yaml:"base_path"`
}
