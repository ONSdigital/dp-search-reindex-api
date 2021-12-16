package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-search-reindex-api
type Config struct {
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	MaxReindexJobRuntime       time.Duration `envconfig:"MAX_REINDEX_JOB_RUNTIME"`
	MongoConfig                MongoConfig
	DefaultMaxLimit            int    `envconfig:"DEFAULT_MAXIMUM_LIMIT"`
	DefaultLimit               int    `envconfig:"DEFAULT_LIMIT"`
	DefaultOffset              int    `envconfig:"DEFAULT_OFFSET"`
	ZebedeeURL                 string `envconfig:"ZEBEDEE_URL"`
	TaskNameValues             string `envconfig:"TASK_NAME_VALUES"`
	SearchApiURL               string `envconfig:"SEARCH_API_URL"`
	ServiceAuthToken           string `envconfig:"SERVICE_AUTH_TOKEN"   json:"-"`
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	BindAddr        string `envconfig:"MONGODB_BIND_ADDR"   json:"-"`
	JobsCollection  string `envconfig:"MONGODB_JOBS_COLLECTION"`
	LocksCollection string `envconfig:"MONGODB_LOCKS_COLLECTION"`
	TasksCollection string `envconfig:"MONGODB_TASKS_COLLECTION"`
	Database        string `envconfig:"MONGODB_DATABASE"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		BindAddr:                   "localhost:25700",
		GracefulShutdownTimeout:    20 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		MaxReindexJobRuntime:       3600 * time.Second,
		MongoConfig: MongoConfig{
			BindAddr:        "localhost:27017",
			JobsCollection:  "jobs",
			LocksCollection: "jobs_locks",
			TasksCollection: "tasks",
			Database:        "search",
		},
		DefaultMaxLimit:  1000,
		DefaultLimit:     20,
		DefaultOffset:    0,
		ZebedeeURL:       "http://localhost:8082",
		TaskNameValues:   "dataset-api,zebedee",
		SearchApiURL:     "http://localhost:23900",
		ServiceAuthToken: "d940e905cbbc752007b9d1b15df4c6926db2eb8b61aa204b4b208182dab28bdc",
	}

	return cfg, envconfig.Process("", cfg)
}
