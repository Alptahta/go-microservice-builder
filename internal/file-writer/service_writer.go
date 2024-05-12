package filewriter

import (
	"fmt"
	"os"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory-builder"
	filebuilder "github.com/Alptahta/go-microservice-builder/internal/file-builder"
	projectinfocollector "github.com/Alptahta/go-microservice-builder/internal/project-info-collector"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FillServiceFile(projectInfo projectinfocollector.ProjectInformation) error {
	serviceFilePath := createServiceFilePath(projectInfo.RepositoryName)

	f, err := os.OpenFile(serviceFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	s := CreateServiceCodes(projectInfo)

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

func createServiceFilePath(serviceName string) string {
	serviceFilePath := fmt.Sprintf("%s/%s/%s/%s", serviceName, directorybuilder.INTERNAL, directorybuilder.SERVICE, filebuilder.SERVICE)
	return serviceFilePath
}

func CreateServiceInterfaceString(domainName string) string {
	return fmt.Sprintf("type ServiceInterface interface {\n\tCreate%s(ctx context.Context) error\n}", cases.Title(language.English, cases.NoLower).String(domainName))
}

func CreateServiceCodes(projectInfo projectinfocollector.ProjectInformation) [][]byte {
	var s [][]byte
	s = append(s,
		[]byte("package service\n\n"),
		[]byte("import \"context\"\n\n"),
		[]byte(CreateServiceInterfaceString(projectInfo.DomainName)),
	)
	return s
}
