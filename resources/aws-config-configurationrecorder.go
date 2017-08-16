package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSConfigConfigurationRecorder AWS CloudFormation Resource (AWS::Config::ConfigurationRecorder)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html
type AWSConfigConfigurationRecorder struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-name
	Name string `json:"Name"`

	// RecordingGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-recordinggroup
	RecordingGroup AWSConfigConfigurationRecorder_RecordingGroup `json:"RecordingGroup"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-rolearn
	RoleARN string `json:"RoleARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigurationRecorder) AWSCloudFormationType() string {
	return "AWS::Config::ConfigurationRecorder"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigurationRecorder) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSConfigConfigurationRecorderResources retrieves all AWSConfigConfigurationRecorder items from a CloudFormation template
func (t *Template) GetAllAWSConfigConfigurationRecorderResources() map[string]*AWSConfigConfigurationRecorder {

	results := map[string]*AWSConfigConfigurationRecorder{}
	for name, resource := range t.Resources {
		result := &AWSConfigConfigurationRecorder{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSConfigConfigurationRecorderWithName retrieves all AWSConfigConfigurationRecorder items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigurationRecorderWithName(name string) (*AWSConfigConfigurationRecorder, error) {

	result := &AWSConfigConfigurationRecorder{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSConfigConfigurationRecorder{}, errors.New("resource not found")

}
