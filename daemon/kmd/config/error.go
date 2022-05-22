package config

import "fmt"

// ErrSQLiteWalletNotAbsolute is returned when the passed sqlite wallet directory is relative
var ErrSQLiteWalletNotAbsolute = fmt.Errorf("sqlite wallets path must be absolute path")
