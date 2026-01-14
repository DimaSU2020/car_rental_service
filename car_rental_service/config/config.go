package config

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Duration time.Duration

type Config struct {
	DBPath       string   `json:"-"`
	Port         string   `json:"port"`
	ReadTimeout  Duration `json:"read_timeout"`
	WriteTimeout Duration `json:"write_timeout"`
}

var (
	ErrDBPathNotSet        = errors.New("DATABASE_PATH is not set")
	ErrConfigNotFound      = errors.New("config file not found")
	ErrInvalidReadTimeout  = errors.New("read_timeout must be > 0")
	ErrInvalidWriteTimeout = errors.New("write_timeout must be > 0")
	ErrInvalidDuration     = errors.New("invalid duration")
)

func New(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, errors.Join(ErrConfigNotFound, err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		return nil, ErrDBPathNotSet
	}
	cfg.DBPath = dbPath

	if cfg.ReadTimeout <= 0 {
		return nil, ErrInvalidReadTimeout
	}

	if cfg.WriteTimeout <= 0 {
		return nil, ErrInvalidWriteTimeout
	}

	return &cfg, nil
}

func MustNewCfg(configPath string) *Config {
	cfg, err := New(configPath)
	if err != nil {
		panic(err)
	}
	return cfg
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return errors.Join(ErrInvalidDuration, err)
	}

	parsed, err := time.ParseDuration(s)
	if err != nil {
		return errors.Join(ErrInvalidDuration, err)
	}

	if parsed <= 0 {
		return ErrInvalidDuration
	}

	*d = Duration(parsed)

	return nil
}
