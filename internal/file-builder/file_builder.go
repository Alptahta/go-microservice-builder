package filebuilder

import (
	"fmt"
	"os"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory-builder"
	projectinfocollector "github.com/Alptahta/go-microservice-builder/internal/project-info-collector"
)

func CreateGoFiles(projectInformation projectinfocollector.ProjectInformation) error {
	err := createMainFile(projectInformation.RepositoryName)
	if err != nil {
		return err
	}

	err = createModelFile(projectInformation.RepositoryName, projectInformation.DomainName)
	if err != nil {
		return err
	}

	err = createRepositoryFile(projectInformation.RepositoryName)
	if err != nil {
		return err
	}

	err = createDBFile(projectInformation.RepositoryName, projectInformation.DatabaseName)
	if err != nil {
		return err
	}

	err = createServiceFile(projectInformation.RepositoryName)
	if err != nil {
		return err
	}

	err = createHandlerFile(projectInformation.RepositoryName)
	if err != nil {
		return err
	}

	return nil
}

func createMainFile(serviceName string) error {
	path := fmt.Sprintf("%s/%s/%s/%s", serviceName, directorybuilder.CMD, directorybuilder.API_SERVER, MAIN)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func createModelFile(serviceName, domainName string) error {
	path := fmt.Sprintf("%s/%s/%s/%s.go", serviceName, directorybuilder.INTERNAL, directorybuilder.MODEL, domainName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func createRepositoryFile(serviceName string) error {
	path := fmt.Sprintf("%s/%s/%s/%s", serviceName, directorybuilder.INTERNAL, directorybuilder.REPOSITORY, REPOSITORY)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func createDBFile(serviceName, dbName string) error {
	path := fmt.Sprintf("%s/%s/%s/%s.go", serviceName, directorybuilder.INTERNAL, directorybuilder.REPOSITORY, dbName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func createServiceFile(serviceName string) error {
	path := fmt.Sprintf("%s/%s/%s/%s", serviceName, directorybuilder.INTERNAL, directorybuilder.SERVICE, SERVICE)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func createHandlerFile(serviceName string) error {
	path := fmt.Sprintf("%s/%s/%s/%s", serviceName, directorybuilder.INTERNAL, directorybuilder.HANDLER, HANDLER)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
