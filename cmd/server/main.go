package main

import (
	"fmt"
	"os"
	"path/filepath"

	"echo-starter/internal"
	startup "echo-starter/internal/startup"

	runtime "github.com/fluffy-bunny/fluffycore/echo/runtime"
	log "github.com/rs/zerolog/log"
)

var version = "Development"

func main() {
	processDirectory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	internal.RootFolder = processDirectory
	fmt.Println("Version:" + version)
	DumpPath("./")
	r := runtime.New(startup.NewStartup())
	err = r.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run the application")
	}
}
