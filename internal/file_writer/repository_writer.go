package filewriter

import (
	"fmt"
	"os"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory_builder"
	filebuilder "github.com/Alptahta/go-microservice-builder/internal/file_builder"
)

func FillRepositoryFile(serviceName string) error {
	repositoryPath := createRepositoryFilePath(serviceName)

	f, err := os.OpenFile(repositoryPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	for _, v := range RepositoryCodes {
		_, err = f.Write(v)
		if err != nil {
			return err
		}
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func createRepositoryFilePath(serviceName string) string {
	repositoryPath := fmt.Sprintf("%s/%s/%s/%s", serviceName, directorybuilder.INTERNAL, directorybuilder.REPOSITORY, filebuilder.REPOSITORY)
	return repositoryPath
}
