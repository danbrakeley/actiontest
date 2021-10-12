package main

import (
	"fmt"

	"github.com/danbrakeley/actiontest/internal/buildvar"
)

func main() {
	fmt.Printf("NEW ACTION TEST PRIME MARK 2 _FINAL\n")
	fmt.Printf("Version: %s\n", buildvar.Version)
	fmt.Printf("Build Time: %s\n", buildvar.BuildTime)
	fmt.Printf("Release Name: %s\n", buildvar.RelaseName)
	fmt.Printf("Release Description: %s\n", buildvar.ReleaseDescription)
	fmt.Printf("Release URL: %s\n", buildvar.ReleaseURL)
}
