package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const (
	kmdConfigFilename = "kmd_config.json"
	defaultSessionLifetimeSecs = 60
	defaultScryptN = 65536
	defaultScryptR = 1
	defaultScryptP = 32
)

// KMDConfig contains global configuration information for kmd
type KMDConfig struct {
	DataDir string `json:"-"`
	DriverConfig DriverConfig `json:"drivers"`
	SessionLifetimeSecs uint64 `json:"session_lifetime_secs"`
	Address string `json:"address"`
	AllowedOrigins []string `json:"allowed_origins"`
}

// DriverConfig contains config info specific to each wallet driver
type DriverConfig struct {
	SQLiteWalletDriverConfig SQLiteWalletDriverConfig `json:"sqlite"`
}

// SQLiteWalletDriverConfig is configuration specific to the SQLWalletDriver
type SQLiteWalletDriverConfig struct {
	WalletsDir string `json:"wallets_dir"`
	UnsafeScrypt bool `json:"allow_unsafe_scrypt"`
	ScryptParams ScryptParams `json:"scrypt"`
}

// ScryptParams stores the parameters used for key derivation. This allows
// upgrading security parameters over time
type ScryptParams struct {
	ScryptN int `json:"scrypt_n"`
	ScryptR int `json:"scrypt_r"`
	ScryptP int `json:"scrypt_p"`
}

// defaultConfig returns the default KMDConfig
func defaultConfig(dataDir string) KMDConfig {
	return KMDConfig{
		DataDir:             dataDir,
		SessionLifetimeSecs: defaultSessionLifetimeSecs,
		DriverConfig: DriverConfig{SQLiteWalletDriverConfig: SQLiteWalletDriverConfig{
			ScryptParams: ScryptParams{
				ScryptN: defaultScryptN,
				ScryptR: defaultScryptR,
				ScryptP: defaultScryptP,
			},
		}},
	}
}

// Validate ensures that the current configuration is valid, returning an error
// if it's not
func (k KMDConfig) Validate() error {
	sqlWalletDir := k.DriverConfig.SQLiteWalletDriverConfig.WalletsDir
	if sqlWalletDir != "" {
		if !filepath.IsAbs(sqlWalletDir) {
			return ErrSQLiteWalletNotAbsolute
		}
	}
	return nil
}

// LoadKMDConfig tries to read the kmd configuration from disk, merging the
// default kmd configuration with what it finds
func LoadKMDConfig(dataDir string) (cfg KMDConfig, err error) {
	cfg = defaultConfig(dataDir)
	configFileName := filepath.Join(dataDir, kmdConfigFilename)
	dat, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return cfg, nil
	}

	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return
	}
	err = cfg.Validate()
	return
}
