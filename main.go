package main

import (
	"fmt"
	"os"

	"github.com/0xJacky/Nginx-UI/cmd"
)

// Version information, injected at build time via ldflags
var (
	Version   = "dev"
	BuildTime = "unknown"
	Commit    = "unknown"
)

func main() {
	// Inject version info into the cmd package
	cmd.Version = Version
	cmd.BuildTime = BuildTime
	cmd.Commit = Commit

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
