package util

import (
	"errors"
	"os/exec"
	"runtime"

	"hrubos.dev/collectorsden/internal/logger"
)

// OpenPath opens a file or folder in the default program/explorer
func OpenPath(path string) error {
    var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// "start" is a shell builtin, so we need cmd.exe
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", path)

	case "darwin":
		// macOS
		cmd = exec.Command("open", path)

	case "linux":
		// Most Linux distros support xdg-open
		cmd = exec.Command("xdg-open", path)

	default:
		errMsg := "Could not open file in default program (" + runtime.GOOS + ")"
		err := errors.New(errMsg)
		logger.Fatal(errMsg, err, logger.CatOther)
		return err
	}

    return cmd.Start()
}
