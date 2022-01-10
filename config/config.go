package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// KafkaTLSProtocolFlag informs service to use TLS protocol for kafka
const KafkaTLSProtocolFlag = "TLS"

var cfg *Config

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
	SearchAPIURL               string `envconfig:"SEARCH_API_URL"`
	ServiceAuthToken           string `envconfig:"SERVICE_AUTH_TOKEN"   json:"-"`
	KafkaConfig                KafkaConfig
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	BindAddr        string `envconfig:"MONGODB_BIND_ADDR"   json:"-"`
	JobsCollection  string `envconfig:"MONGODB_JOBS_COLLECTION"`
	LocksCollection string `envconfig:"MONGODB_LOCKS_COLLECTION"`
	TasksCollection string `envconfig:"MONGODB_TASKS_COLLECTION"`
	Database        string `envconfig:"MONGODB_DATABASE"`
}

// KafkaConfig contains the config required to connect to Kafka
type KafkaConfig struct {
	Brokers               []string `envconfig:"KAFKA_ADDR"                            json:"-"`
	Version               string   `envconfig:"KAFKA_VERSION"`
	SecProtocol           string   `envconfig:"KAFKA_SEC_PROTO"`
	SecCACerts            string   `envconfig:"KAFKA_SEC_CA_CERTS"`
	SecClientKey          string   `envconfig:"KAFKA_SEC_CLIENT_KEY"                  json:"-"`
	SecClientCert         string   `envconfig:"KAFKA_SEC_CLIENT_CERT"`
	SecSkipVerify         bool     `envconfig:"KAFKA_SEC_SKIP_VERIFY"`
	ReindexRequestedTopic string   `envconfig:"KAFKA_REINDEX_REQUESTED_TOPIC"`
}

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
		SearchAPIURL:     "http://localhost:23900",
		ServiceAuthToken: "",
		KafkaConfig: KafkaConfig{
			Brokers:               []string{"localhost:9092"},
			Version:               "1.0.2",
			SecProtocol:           "",
			SecCACerts:            "",
			SecClientCert:         "",
			SecClientKey:          "",
			SecSkipVerify:         false,
			ReindexRequestedTopic: "reindex-requested",
		},
	}

	return cfg, envconfig.Process("", cfg)
}
