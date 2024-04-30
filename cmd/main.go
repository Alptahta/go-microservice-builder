package main

import (
	"log"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory-builder"
	filebuilder "github.com/Alptahta/go-microservice-builder/internal/file-builder"
	filewriter "github.com/Alptahta/go-microservice-builder/internal/file-writer"
	projectinfocollector "github.com/Alptahta/go-microservice-builder/internal/project-info-collector"
)

func main() {
	projectInformation := projectinfocollector.GetProjectInfo()

	err := directorybuilder.CreateSkeletonDirectories(projectInformation.RepositoryName)
	if err != nil {
		log.Fatalln(err)
	}

	err = filebuilder.CreateGoFiles(projectInformation)
	if err != nil {
		log.Fatalln(err)
	}

	err = filewriter.FillRepositoryFile(projectInformation)
	if err != nil {
		log.Fatalln(err)
	}
}
