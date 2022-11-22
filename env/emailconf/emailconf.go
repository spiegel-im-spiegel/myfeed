package emailconf

import (
	"encoding/json"
	"os"

	"github.com/goark/errs"
)

type Config struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Encrypt  bool   `json:"encrypt"`
	From     string `json:"from"`
	To       string `json:"to"`
}

// ImportFile returns ChemConfig instance from config file.
func ImportFile(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return Import(b)
}

// Import returns ChemConfig instance from config data.
func Import(b []byte) (*Config, error) {
	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, errs.Wrap(err)
	}
	return &cfg, nil
}
