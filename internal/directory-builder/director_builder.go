package directorybuilder

import (
	"fmt"
	"os"
)

func CreateSkeletonDirectories(serviceName string) error {
	err := createServiceName(serviceName)
	if err != nil {
		return err
	}

	err = createCMDFolder(serviceName)
	if err != nil {
		return err
	}

	err = createInternalFolder(serviceName)
	if err != nil {
		return err
	}

	err = createRepositoryFolder(serviceName)
	if err != nil {
		return err
	}

	err = createServiceFolder(serviceName)
	if err != nil {
		return err
	}

	err = createHandlerFolder(serviceName)
	if err != nil {
		return err
	}

	return nil
}

func createServiceName(serviceName string) error {
	err := os.Mkdir(serviceName, os.ModePerm)
	return err
}

func createCMDFolder(serviceName string) error {
	path := fmt.Sprintf("%s/%s", serviceName, CMD)
	err := os.Mkdir(path, os.ModePerm)
	return err
}

func createInternalFolder(serviceName string) error {
	path := fmt.Sprintf("%s/%s", serviceName, INTERNAL)
	err := os.Mkdir(path, os.ModePerm)
	return err
}

func createRepositoryFolder(serviceName string) error {
	path := fmt.Sprintf("%s/%s/%s", serviceName, INTERNAL, REPOSITORY)
	err := os.Mkdir(path, os.ModePerm)
	return err
}

func createServiceFolder(serviceName string) error {
	path := fmt.Sprintf("%s/%s/%s", serviceName, INTERNAL, SERVICE)
	err := os.Mkdir(path, os.ModePerm)
	return err
}

func createHandlerFolder(serviceName string) error {
	path := fmt.Sprintf("%s/%s/%s", serviceName, INTERNAL, HANDLER)
	err := os.Mkdir(path, os.ModePerm)
	return err
}
