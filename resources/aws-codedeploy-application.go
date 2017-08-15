package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::Application AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-application.html
type AWSCodeDeployApplication struct {

	// ApplicationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-application.html#cfn-codedeploy-application-applicationname
	ApplicationName string `json:"ApplicationName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployApplication) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::Application"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployApplication) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployApplicationResources retrieves all AWSCodeDeployApplication items from a CloudFormation template
func GetAllAWSCodeDeployApplicationResources(template *Template) map[string]*AWSCodeDeployApplication {

	results := map[string]*AWSCodeDeployApplication{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployApplication{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployApplicationWithName retrieves all AWSCodeDeployApplication items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCodeDeployApplicationWithName(name string, template *Template) (*AWSCodeDeployApplication, error) {

	result := &AWSCodeDeployApplication{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployApplication{}, errors.New("resource not found")

}
