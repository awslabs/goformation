package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DirectoryService::MicrosoftAD.VpcSettings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html
type AWSDirectoryServiceMicrosoftAD_VpcSettings struct {

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html#cfn-directoryservice-microsoftad-vpcsettings-subnetids
	SubnetIds []string `json:"SubnetIds"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html#cfn-directoryservice-microsoftad-vpcsettings-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceMicrosoftAD_VpcSettings) AWSCloudFormationType() string {
	return "AWS::DirectoryService::MicrosoftAD.VpcSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDirectoryServiceMicrosoftAD_VpcSettings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDirectoryServiceMicrosoftAD_VpcSettingsResources retrieves all AWSDirectoryServiceMicrosoftAD_VpcSettings items from a CloudFormation template
func GetAllAWSDirectoryServiceMicrosoftAD_VpcSettings(template *Template) map[string]*AWSDirectoryServiceMicrosoftAD_VpcSettings {

	results := map[string]*AWSDirectoryServiceMicrosoftAD_VpcSettings{}
	for name, resource := range template.Resources {
		result := &AWSDirectoryServiceMicrosoftAD_VpcSettings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDirectoryServiceMicrosoftAD_VpcSettingsWithName retrieves all AWSDirectoryServiceMicrosoftAD_VpcSettings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDirectoryServiceMicrosoftAD_VpcSettings(name string, template *Template) (*AWSDirectoryServiceMicrosoftAD_VpcSettings, error) {

	result := &AWSDirectoryServiceMicrosoftAD_VpcSettings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDirectoryServiceMicrosoftAD_VpcSettings{}, errors.New("resource not found")

}
