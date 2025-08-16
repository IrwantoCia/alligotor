package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/IrwantoCia/alligotor/cmd/gotor/libs"
)

//go:embed all:templates/*
var templatesFS embed.FS

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Help:")
		fmt.Println("	Usage: gotor init <project_name>")
		fmt.Println("	Usage: gotor context <context_name>")
		os.Exit(1)
	}

	command := os.Args[1]
	contextName := os.Args[2]

	switch command {
	case "init":
		libs.InitProject(contextName, templatesFS)
	case "context":
		libs.CreateContext(contextName, templatesFS)
	}
	os.Exit(0)
}
