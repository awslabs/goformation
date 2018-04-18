package main

import (
	"log"

	"github.com/awslabs/goformation"
)

func main() {

	// Open a template from file (can be JSON or YAML)
	template, err := goformation.Open("template.yaml")
	if err != nil {
		log.Fatalf("There was an error processing the template: %s", err)
	}

	// You can extract all resources of a certain type
	// Each AWS CloudFormation resource is a strongly typed struct
	functions := template.GetAllAWSServerlessFunctionResources()
	for name, function := range functions {

		// E.g. Found a AWS::Serverless::Function named GetHelloWorld (runtime: nodejs6.10)
		log.Printf("Found a %s named %s (runtime: %s)\n", function.AWSCloudFormationType(), name, function.Runtime)

	}

	// You can also search for specific resources by their logicalId
	search := "GetHelloWorld"
	function, err := template.GetAWSServerlessFunctionWithName(search)
	if err != nil {
		log.Fatalf("Function not found")
	}

	// E.g. Found a AWS::Serverless::Function named GetHelloWorld (runtime: nodejs6.10)
	log.Printf("Found a %s named %s (runtime: %s)\n", function.AWSCloudFormationType(), search, function.Runtime)

}
