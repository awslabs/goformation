package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECR::Repository AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html
type AWSECRRepository struct {

	// RepositoryName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositoryname

	RepositoryName string `json:"RepositoryName"`

	// RepositoryPolicyText AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositorypolicytext

	RepositoryPolicyText interface{} `json:"RepositoryPolicyText"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECRRepository) AWSCloudFormationType() string {
	return "AWS::ECR::Repository"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECRRepository) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECRRepositoryResources retrieves all AWSECRRepository items from a CloudFormation template
func GetAllAWSECRRepository(template *Template) map[string]*AWSECRRepository {

	results := map[string]*AWSECRRepository{}
	for name, resource := range template.Resources {
		result := &AWSECRRepository{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECRRepositoryWithName retrieves all AWSECRRepository items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECRRepository(name string, template *Template) (*AWSECRRepository, error) {

	result := &AWSECRRepository{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECRRepository{}, errors.New("resource not found")

}
