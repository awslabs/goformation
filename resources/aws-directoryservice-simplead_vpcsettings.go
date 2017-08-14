package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DirectoryService::SimpleAD.VpcSettings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html
type AWSDirectoryServiceSimpleAD_VpcSettings struct {

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-subnetids

	SubnetIds []string `json:"SubnetIds"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-vpcid

	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceSimpleAD_VpcSettings) AWSCloudFormationType() string {
	return "AWS::DirectoryService::SimpleAD.VpcSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDirectoryServiceSimpleAD_VpcSettings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDirectoryServiceSimpleAD_VpcSettingsResources retrieves all AWSDirectoryServiceSimpleAD_VpcSettings items from a CloudFormation template
func GetAllAWSDirectoryServiceSimpleAD_VpcSettings(template *Template) map[string]*AWSDirectoryServiceSimpleAD_VpcSettings {

	results := map[string]*AWSDirectoryServiceSimpleAD_VpcSettings{}
	for name, resource := range template.Resources {
		result := &AWSDirectoryServiceSimpleAD_VpcSettings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDirectoryServiceSimpleAD_VpcSettingsWithName retrieves all AWSDirectoryServiceSimpleAD_VpcSettings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDirectoryServiceSimpleAD_VpcSettings(name string, template *Template) (*AWSDirectoryServiceSimpleAD_VpcSettings, error) {

	result := &AWSDirectoryServiceSimpleAD_VpcSettings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDirectoryServiceSimpleAD_VpcSettings{}, errors.New("resource not found")

}
