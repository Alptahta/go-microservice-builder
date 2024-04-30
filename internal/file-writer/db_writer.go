package filewriter

import (
	"fmt"
	"os"

	directorybuilder "github.com/Alptahta/go-microservice-builder/internal/directory-builder"
	projectinfocollector "github.com/Alptahta/go-microservice-builder/internal/project-info-collector"
)

func FillDBFile(projectInfo projectinfocollector.ProjectInformation) error {
	dbFilePath := createDBFilePath(projectInfo.RepositoryName, projectInfo.DatabaseName)

	f, err := os.OpenFile(dbFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	s := CreateDBCodes(projectInfo)

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

func createDBFilePath(serviceName, dbName string) string {
	dbFilePath := fmt.Sprintf("%s/%s/%s/%s.go", serviceName, directorybuilder.INTERNAL, directorybuilder.REPOSITORY, dbName)
	return dbFilePath
}

func CreateDBCodes(projectInfo projectinfocollector.ProjectInformation) [][]byte {
	var s [][]byte
	s = append(s,
		[]byte("package repository\n\n"),
		[]byte("import \"context\"\n\n"),
		[]byte("type Repository struct {\n\tctx context.Context\n}\n\n"),
		[]byte("func Create(ctx context.Context) *Repository {\n\treturn &Repository{ctx: ctx}\n}\n\n"),
	)
	return s
}
