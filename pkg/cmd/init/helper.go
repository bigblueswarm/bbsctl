package init

import (
	"os"
)

func fileAlreadyExists(path string) bool {
	info, _ := os.Stat(path)
	return info != nil
}
