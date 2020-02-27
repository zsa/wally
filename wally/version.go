package wally

import "runtime"

// GetAppVersion returns the current version number.
func GetAppVersion() string {
	switch runtime.GOOS {
	case "darwin":
		return "1.1.2"
	case "linux":
		return "1.1.1"
	case "windows":
		return "1.1.4"
	default:
		return "1.1.0"
	}
}
