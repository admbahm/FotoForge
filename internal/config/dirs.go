package config

import "os"

// UserConfigDir returns the platform-specific user configuration directory.
var UserConfigDir = os.UserConfigDir
