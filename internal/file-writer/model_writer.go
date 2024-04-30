package filewriter

import (
	"fmt"
	"os"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory-builder"
	projectinfocollector "github.com/Alptahta/go-microservice-builder/internal/project-info-collector"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FillModelFile(projectInfo projectinfocollector.ProjectInformation) error {
	modelPath := createModelFilePath(projectInfo.RepositoryName, projectInfo.DomainName)

	f, err := os.OpenFile(modelPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	s := CreateModelCodes(projectInfo.DomainName)

	for _, v := range s {
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

func createModelFilePath(serviceName, domainName string) string {
	modelPath := fmt.Sprintf("%s/%s/%s/%s.go", serviceName, directorybuilder.INTERNAL, directorybuilder.MODEL, domainName)
	return modelPath
}

func CreateModelCodes(domainName string) [][]byte {
	var s [][]byte
	s = append(s,
		[]byte("package model\n\n"),
		[]byte(CreateModelStructString(domainName)),
	)
	return s
}

func CreateModelStructString(domainName string) string {
	return fmt.Sprintf("type %s struct {}", cases.Title(language.English, cases.NoLower).String(domainName))
}
