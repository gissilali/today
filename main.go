package main

import (
	"github.com/gissilali/today/cmd"
	"github.com/gissilali/today/repositories"
)

func main() {
	repositories.InitDB()
	cmd.Execute()
}
