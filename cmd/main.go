package main

import (
	"log"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory_builder"
)

func main() {
	serviceName := "user-service"
	err := directorybuilder.CreateSkeletonDirectories(serviceName)
	if err != nil {
		log.Fatalln(err)
	}

}
