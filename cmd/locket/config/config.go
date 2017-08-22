package config

import (
	"encoding/json"
	"os"

	"code.cloudfoundry.org/debugserver"
	loggingclient "code.cloudfoundry.org/diego-logging-client"
	"code.cloudfoundry.org/lager/lagerflags"
)

type LocketConfig struct {
	CaFile                     string               `json:"ca_file"`
	CertFile                   string               `json:"cert_file"`
	ConsulCluster              string               `json:"consul_cluster,omitempty"`
	DatabaseConnectionString   string               `json:"database_connection_string"`
	MaxOpenDatabaseConnections int                  `json:"max_open_database_connections,omitempty"`
	DatabaseDriver             string               `json:"database_driver,omitempty"`
	DropsondePort              int                  `json:"dropsonde_port,omitempty"`
	KeyFile                    string               `json:"key_file"`
	ListenAddress              string               `json:"listen_address"`
	SQLCACertFile              string               `json:"sql_ca_cert_file,omitempty"`
	LoggregatorConfig          loggingclient.Config `json:"loggregator"`
	debugserver.DebugServerConfig
	lagerflags.LagerConfig
}

func DefaultLocketConfig() LocketConfig {
	return LocketConfig{
		LagerConfig:    lagerflags.DefaultLagerConfig(),
		DatabaseDriver: "mysql",
	}
}

func NewLocketConfig(configPath string) (LocketConfig, error) {
	locketConfig := DefaultLocketConfig()
	configFile, err := os.Open(configPath)
	if err != nil {
		return LocketConfig{}, err
	}

	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	err = decoder.Decode(&locketConfig)
	if err != nil {
		return LocketConfig{}, err
	}

	return locketConfig, nil
}
