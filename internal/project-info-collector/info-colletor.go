package projectinfocollector

import "fmt"

func GetProjectInfo() ProjectInformation {
	projectInformation := ProjectInformation{}

	fmt.Println("Welcome to Microservice creator with Golang")
	fmt.Println("Before creating your project, Tell me how do you want to keep it")
	fmt.Println("For Remote Hosts: 1")
	fmt.Println("For Local: 2")

	fmt.Scanln(&projectInformation.HostType)

	var projectHostNameNumber int

	if projectInformation.HostType == 1 {
		fmt.Println("Which remote host do you want to use.(We currently only have Github)")
		fmt.Println("For Github: 1")
		fmt.Scanln(&projectHostNameNumber)
		projectInformation.ProjectHostName = RemoteHostNames[projectHostNameNumber]
		if projectInformation.ProjectHostName == RemoteHostNames[1] {
			fmt.Println("Tell us your Github username")
			fmt.Scanln(&projectInformation.RemoteHostUsername)
		}
	}
	fmt.Println("What should we name your repository")
	fmt.Println("Example: user-service")
	fmt.Scanln(&projectInformation.RepositoryName)

	fmt.Println("What should we name your domain")
	fmt.Println("Example: user")
	fmt.Scanln(&projectInformation.DomainName)

	fmt.Printf("The Project will be initialize with go mod init %s.com/%s/%s", projectInformation.ProjectHostName, projectInformation.RemoteHostUsername, projectInformation.RepositoryName)
	return projectInformation
}

type ProjectInformation struct {
	HostType           int
	ProjectHostName    string
	RemoteHostUsername string
	RepositoryName     string
	DomainName         string
}

var RemoteHostNames = map[int]string{
	1: "github",
	2: "gitlab",
}
