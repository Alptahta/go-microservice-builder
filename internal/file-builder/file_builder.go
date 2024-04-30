package filebuilder

import (
	"fmt"
	"os"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory-builder"
)

func CreateGoFiles(serviceName string) error {
	err := createMainFile(serviceName)
	if err != nil {
		return err
	}

	err = createRepositoryFile(serviceName)
	if err != nil {
		return err
	}

	err = createServiceFile(serviceName)
	if err != nil {
		return err
	}

	err = createHandlerFile(serviceName)
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

func createRepositoryFile(serviceName string) error {
	path := fmt.Sprintf("%s/%s/%s/%s", serviceName, directorybuilder.INTERNAL, directorybuilder.REPOSITORY, REPOSITORY)
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
