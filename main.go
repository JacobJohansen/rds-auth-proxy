package main

import (
	"github.com/JacobJohansen/rds-auth-proxy/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
