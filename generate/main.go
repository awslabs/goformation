package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("GoFormation Resource Generator\n")

	// cloudformationSpec := "https://d1uauaxba7bl26.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json"
	cloudformationSpec := "file://generate/cfn-2019-10-26.json"

	// // TODO(pmaddox): commented out until I work out what SAMPT means in the spec
	// otherSpecs := map[string]string{
	// 	"sam": "file://generate/sam-2016-10-31.json",
	// }
	otherSpecs := map[string]string{}

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
