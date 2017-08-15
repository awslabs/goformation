package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeCommit::Repository AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html
type AWSCodeCommitRepository struct {

	// RepositoryDescription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-repositorydescription
	RepositoryDescription string `json:"RepositoryDescription"`

	// RepositoryName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-repositoryname
	RepositoryName string `json:"RepositoryName"`

	// Triggers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-triggers
	Triggers []AWSCodeCommitRepository_RepositoryTrigger `json:"Triggers"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeCommitRepository) AWSCloudFormationType() string {
	return "AWS::CodeCommit::Repository"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeCommitRepository) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeCommitRepositoryResources retrieves all AWSCodeCommitRepository items from a CloudFormation template
func GetAllAWSCodeCommitRepositoryResources(template *Template) map[string]*AWSCodeCommitRepository {

	results := map[string]*AWSCodeCommitRepository{}
	for name, resource := range template.Resources {
		result := &AWSCodeCommitRepository{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeCommitRepositoryWithName retrieves all AWSCodeCommitRepository items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCodeCommitRepositoryWithName(name string, template *Template) (*AWSCodeCommitRepository, error) {

	result := &AWSCodeCommitRepository{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeCommitRepository{}, errors.New("resource not found")

}
