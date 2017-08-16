package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCodeCommitRepository AWS CloudFormation Resource (AWS::CodeCommit::Repository)
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
func (t *Template) GetAllAWSCodeCommitRepositoryResources() map[string]*AWSCodeCommitRepository {

	results := map[string]*AWSCodeCommitRepository{}
	for name, resource := range t.Resources {
		result := &AWSCodeCommitRepository{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeCommitRepositoryWithName retrieves all AWSCodeCommitRepository items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeCommitRepositoryWithName(name string) (*AWSCodeCommitRepository, error) {

	result := &AWSCodeCommitRepository{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeCommitRepository{}, errors.New("resource not found")

}
