package main

import (
	"log"
	"time"

	"github.com/awslabs/goformation"
	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/intrinsics"
)

func main() {

	// Open a template from file (can be JSON or YAML)
	template, err := goformation.OpenWithOptions("/Users/zhouy4l/go/src/stash.abc-dev.net.au/ter/serverless-event-management/deployments/aws/ingestor-sam.yaml", &intrinsics.ProcessorOptions{
		ParameterOverrides: map[string]interface{}{
			"HostEnv":        "staging",
			"DeployTime":     time.Now().String(),
			"QueueStackName": "QueueStack",
		},
	}, map[string]func() cloudformation.Resource{
		"Custom::PublishLambdaVersion": func() cloudformation.Resource {
			return &cloudformation.BasicCustomResource{}
		},
		"Custom::SoftStackRef": func() cloudformation.Resource {
			return &cloudformation.BasicCustomResource{}
		},
	},
	)
	if err != nil {
		log.Fatalf("There was an error processing the template: %s", err)
	}

	// You can extract all resources of a certain type
	// Each AWS CloudFormation resource is a strongly typed struct
	functions := template.GetAllAWSServerlessFunctionResources()
	for name, function := range functions {

		// E.g. Found a AWS::Serverless::Function named GetHelloWorld (runtime: nodejs6.10)
		log.Printf("Found a %s named %s (runtime: %s)\n", function.AWSCloudFormationType(), name, function.Runtime)
		log.Printf("\tEnvironment: %v\n", function.Environment.Variables)

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
