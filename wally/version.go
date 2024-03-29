package wally

import "runtime"

// GetAppVersion returns the current version number.
func GetAppVersion() string {
	switch runtime.GOOS {
	case "darwin":
		return "2.1.0"
	case "linux":
		return "2.1.3"
	case "windows":
		return "2.1.0"
	default:
		return "2.1.0"
	}
}
