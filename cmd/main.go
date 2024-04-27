package main

import (
	"log"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory_builder"
	filebuilder "github.com/Alptahta/go-microservice-builder/internal/file_builder"
	filewriter "github.com/Alptahta/go-microservice-builder/internal/file_writer"
)

func main() {
	serviceName := "user-service"

	err := directorybuilder.CreateSkeletonDirectories(serviceName)
	if err != nil {
		log.Fatalln(err)
	}

	err = filebuilder.CreateGoFiles(serviceName)
	if err != nil {
		log.Fatalln(err)
	}

	err = filewriter.FillRepositoryFile(serviceName)
	if err != nil {
		log.Fatalln(err)
	}
}
