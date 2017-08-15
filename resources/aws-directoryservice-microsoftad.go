package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DirectoryService::MicrosoftAD AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html
type AWSDirectoryServiceMicrosoftAD struct {

	// CreateAlias AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-createalias

	CreateAlias bool `json:"CreateAlias"`

	// EnableSso AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-enablesso

	EnableSso bool `json:"EnableSso"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-name

	Name string `json:"Name"`

	// Password AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-password

	Password string `json:"Password"`

	// ShortName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-shortname

	ShortName string `json:"ShortName"`

	// VpcSettings AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-vpcsettings

	VpcSettings AWSDirectoryServiceMicrosoftAD_VpcSettings `json:"VpcSettings"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceMicrosoftAD) AWSCloudFormationType() string {
	return "AWS::DirectoryService::MicrosoftAD"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDirectoryServiceMicrosoftAD) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDirectoryServiceMicrosoftADResources retrieves all AWSDirectoryServiceMicrosoftAD items from a CloudFormation template
func GetAllAWSDirectoryServiceMicrosoftAD(template *Template) map[string]*AWSDirectoryServiceMicrosoftAD {

	results := map[string]*AWSDirectoryServiceMicrosoftAD{}
	for name, resource := range template.Resources {
		result := &AWSDirectoryServiceMicrosoftAD{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDirectoryServiceMicrosoftADWithName retrieves all AWSDirectoryServiceMicrosoftAD items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDirectoryServiceMicrosoftAD(name string, template *Template) (*AWSDirectoryServiceMicrosoftAD, error) {

	result := &AWSDirectoryServiceMicrosoftAD{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDirectoryServiceMicrosoftAD{}, errors.New("resource not found")

}
