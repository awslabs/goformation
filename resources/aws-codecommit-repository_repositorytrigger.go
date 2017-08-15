package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeCommit::Repository.RepositoryTrigger AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html
type AWSCodeCommitRepository_RepositoryTrigger struct {

	// Branches AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-branches

	Branches []string `json:"Branches"`

	// CustomData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-customdata

	CustomData string `json:"CustomData"`

	// DestinationArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-destinationarn

	DestinationArn string `json:"DestinationArn"`

	// Events AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-events

	Events []string `json:"Events"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-name

	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeCommitRepository_RepositoryTrigger) AWSCloudFormationType() string {
	return "AWS::CodeCommit::Repository.RepositoryTrigger"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeCommitRepository_RepositoryTrigger) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeCommitRepository_RepositoryTriggerResources retrieves all AWSCodeCommitRepository_RepositoryTrigger items from a CloudFormation template
func GetAllAWSCodeCommitRepository_RepositoryTrigger(template *Template) map[string]*AWSCodeCommitRepository_RepositoryTrigger {

	results := map[string]*AWSCodeCommitRepository_RepositoryTrigger{}
	for name, resource := range template.Resources {
		result := &AWSCodeCommitRepository_RepositoryTrigger{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeCommitRepository_RepositoryTriggerWithName retrieves all AWSCodeCommitRepository_RepositoryTrigger items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeCommitRepository_RepositoryTrigger(name string, template *Template) (*AWSCodeCommitRepository_RepositoryTrigger, error) {

	result := &AWSCodeCommitRepository_RepositoryTrigger{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeCommitRepository_RepositoryTrigger{}, errors.New("resource not found")

}
