package main

import (
	"encoding/json"
	"fmt"

	catalyst "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"
)

// client is Catalyst Center API client
var client *catalyst.Client

func responseInterfaceToString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}
	return fmt.Sprint(string(b))
}

func main() {
	var err error
	// projectID := ""
	// templateId := ""
	fmt.Println("Authenticating...")
	client, err = catalyst.NewClientWithOptions("https://100.119.103.190",
		"cloverhound_user", "LABchsys!23$",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("Printing GetsAListOfProjects")
	// getListOfProjectsQueryParams := &catalyst.GetsAListOfProjectsQueryParams{}
	// respListOfProjects, _, err := client.ConfigurationTemplates.GetsAListOfProjects(getListOfProjectsQueryParams)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if respListOfProjects != nil {
	// 	fmt.Println(respListOfProjects)
	// 	projectID = (*respListOfProjects)[0].ID
	// } else {
	// 	fmt.Println("There's no data on GetsAListOfProjects response")
	// 	return
	// }

	// fmt.Println("Printing DetailsOfAGivenProject")
	// respDetailsOfAGivenProject, _, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(projectID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if respDetailsOfAGivenProject != nil {
	// 	fmt.Println(respDetailsOfAGivenProject)
	// } else {
	// 	fmt.Println("There's no data on DetailsOfAGivenProject response")
	// 	return
	// }

	// fmt.Println("Printing GetsTheTemplatesAvailable")
	// getTemplatesAvailableQueryParams := &catalyst.GetsTheTemplatesAvailableQueryParams{}
	// respTemplatesAvailable, _, err := client.ConfigurationTemplates.GetsTheTemplatesAvailable(getTemplatesAvailableQueryParams)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if respTemplatesAvailable != nil {
	// 	fmt.Println(respTemplatesAvailable)
	// 	templateId = (*respTemplatesAvailable)[0].TemplateID
	// } else {
	// 	fmt.Println("There's no data on GetsTheTemplatesAvailable response")
	// 	return
	// }

	// //Print StatusOfTemplateDeployment can't test

	// fmt.Println("Printng GetsAllTheVersionsOfAGivenTemplate")
	// respAllTheVersionsOfAGivenTemplate, _, err := client.ConfigurationTemplates.GetsAllTheVersionsOfAGivenTemplate(templateId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if respAllTheVersionsOfAGivenTemplate != nil {
	// 	fmt.Println(respAllTheVersionsOfAGivenTemplate)
	// } else {
	// 	fmt.Println("There's no data on GetsAllTheVersionsOfAGivenTemplate response")
	// 	return
	// }

	resp, _, err := client.ConfigurationTemplates.GetProjectsDetails(&catalyst.GetProjectsDetailsQueryParams{
		Name: "Onboarding Configuration",
	})

	println(responseInterfaceToString(resp))

	// /v2 endpoints can't test

	/*fmt.Println("Creates a CLONE")
	respCreatesAClone, _, err := client.ConfigurationTemplates.CreatesACloneOfTheGivenTemplate("DMVPN for Cloud Router - System Default", templateId, projectID)

	if err != nil {
		fmt.Println(err)
		return
	}

	if respCreatesAClone != nil {
		fmt.Println(respCreatesAClone)
	} else {
		fmt.Println("There's no data on Creates a CLONE response")
		return
	}

	fmt.Println("Creates a Project")

	createTime := 1
	lastUpdateTime := 1

	//createProject:=&catalyst.RequestConfigurationTemplatesCreateProjectTemplates{}
	reqBodyProject := &catalyst.RequestConfigurationTemplatesCreateProject{
		CreateTime:     &createTime,
		Description:    "DESCRIPTION",
		ID:             "1",
		LastUpdateTime: &lastUpdateTime,
		Name:           "Name",
		//Templates:      &createProject,
	}

	respCreatesAProject, _, err := client.ConfigurationTemplates.CreateProject(reqBodyProject)

	if err != nil {
		fmt.Println(err)
		return
	}

	if respCreatesAClone != nil {
		fmt.Println(respCreatesAProject)
	} else {
		fmt.Println("There's no data on Creates a CLONE response")
		return
	}
	*/
}
