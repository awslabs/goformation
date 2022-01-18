package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("GoFormation Resource Generator\n")

	// Fetch and process the AWS published CloudFormation Resource Specification
	cloudformationSpec := "https://d1uauaxba7bl26.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json"

	otherSpecs := map[string]string{
		// We have a manually generated SAM specification in this repo too
		// which needs to be manually updated when the SAM spec changes
		"sam": "file://generate/sam-2016-10-31.json",
		// We have a manually generated CDK specification too, which currently
		// only includes the AWS::CDK::Metadata resource type, however it could
		// include others in the future if more resources are published under the
		// AWS::CDK:: namespace.
		"cdk": "file://generate/cdk.json",
	}

	rg, err := NewResourceGenerator(cloudformationSpec, otherSpecs)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	if err := rg.Generate(); err != nil {
		fmt.Printf("ERROR: Failed to generate resources: %s\n", err)
		os.Exit(1)
	}

	if len(rg.Results.UpdatedResources) > 0 {
		fmt.Printf("\nUpdated the following AWS CloudFormation resources:\n\n")
		for _, updated := range rg.Results.UpdatedResources {
			fmt.Printf(" - %s\n", updated.Name)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Processed %d resources\n", rg.Results.ProcessedCount)

}
