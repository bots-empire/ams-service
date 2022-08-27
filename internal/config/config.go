package config

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServicePort   string          `yaml:"service_port"`
	MetricsPort   string          `yaml:"metrics_port"`
	RepositoryCfg *pgxpool.Config `yaml:"repository_cfg"`

	DBConnConf dbConnConfig `yaml:"db_conn_conf"`
}

type dbConnConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	DBName       string `yaml:"db_name"`
	PoolMaxConns string `yaml:"pool_max_conns"`
}

func InitConfig() (*Config, error) {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		return nil, fmt.Errorf("config path is empty (help: set CONFIG_PATH=<path>)")
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, errors.Wrap(err, "didn't find base config")
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?pool_max_conns=%s",
		cfg.DBConnConf.User,
		cfg.DBConnConf.Password,
		cfg.DBConnConf.Host,
		cfg.DBConnConf.DBName,
		cfg.DBConnConf.PoolMaxConns)

	repCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "parse data base config")
	}

	cfg.RepositoryCfg = repCfg

	return &cfg, nil
}
