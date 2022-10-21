package state

import "runtime"

func GetAppVersion() string {
	switch runtime.GOOS {
	case "darwin":
		return "3.0.0β"
	case "linux":
		return "3.0.0β"
	case "windows":
		return "3.1.0β"
	default:
		return "n/a"
	}
}
