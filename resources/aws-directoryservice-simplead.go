package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSDirectoryServiceSimpleAD AWS CloudFormation Resource (AWS::DirectoryService::SimpleAD)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html
type AWSDirectoryServiceSimpleAD struct {

	// CreateAlias AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-createalias
	CreateAlias bool `json:"CreateAlias"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-description
	Description string `json:"Description"`

	// EnableSso AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-enablesso
	EnableSso bool `json:"EnableSso"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-name
	Name string `json:"Name"`

	// Password AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-password
	Password string `json:"Password"`

	// ShortName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-shortname
	ShortName string `json:"ShortName"`

	// Size AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-size
	Size string `json:"Size"`

	// VpcSettings AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-vpcsettings
	VpcSettings AWSDirectoryServiceSimpleAD_VpcSettings `json:"VpcSettings"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceSimpleAD) AWSCloudFormationType() string {
	return "AWS::DirectoryService::SimpleAD"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDirectoryServiceSimpleAD) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDirectoryServiceSimpleADResources retrieves all AWSDirectoryServiceSimpleAD items from a CloudFormation template
func (t *Template) GetAllAWSDirectoryServiceSimpleADResources() map[string]*AWSDirectoryServiceSimpleAD {

	results := map[string]*AWSDirectoryServiceSimpleAD{}
	for name, resource := range t.Resources {
		result := &AWSDirectoryServiceSimpleAD{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDirectoryServiceSimpleADWithName retrieves all AWSDirectoryServiceSimpleAD items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDirectoryServiceSimpleADWithName(name string) (*AWSDirectoryServiceSimpleAD, error) {

	result := &AWSDirectoryServiceSimpleAD{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDirectoryServiceSimpleAD{}, errors.New("resource not found")

}
