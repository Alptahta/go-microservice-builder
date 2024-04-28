package filewriter

import (
	"fmt"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory-builder"
	filebuilder "github.com/Alptahta/go-microservice-builder/internal/file-builder"
	projectinfocollector "github.com/Alptahta/go-microservice-builder/internal/project-info-collector"
)

func FillRepositoryFile(projectInfo projectinfocollector.ProjectInformation) error {
	repositoryPath := createRepositoryFilePath(projectInfo.RepositoryName)

	f, err := os.OpenFile(repositoryPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	s := CreateRepositoryCodes(projectInfo)

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

func createRepositoryFilePath(serviceName string) string {
	repositoryPath := fmt.Sprintf("%s/%s/%s/%s", serviceName, directorybuilder.INTERNAL, directorybuilder.REPOSITORY, filebuilder.REPOSITORY)
	return repositoryPath
}

func CreateRepositoryInterfaceString(domainName string) string {
	return fmt.Sprintf("type RepositoryInterface interface {\n\tCreate%s() error\n}", cases.Title(language.English, cases.NoLower).String(domainName))
}

func CreateRepositoryCodes(projectInfo projectinfocollector.ProjectInformation) [][]byte {
	var s [][]byte
	s = append(s,
		[]byte("package repository\n\n"),
		[]byte("import \"context\"\n\n"),
		[]byte("type Repository struct {\n\tctx context.Context\n}\n\n"),
		[]byte("func CreateRepository(ctx context.Context) *Repository {\n\treturn &Repository{ctx: ctx}\n}\n\n"),
		[]byte(CreateRepositoryInterfaceString(projectInfo.DomainName)),
	)
	return s
}
