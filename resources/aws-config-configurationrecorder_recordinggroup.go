package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Config::ConfigurationRecorder.RecordingGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html
type AWSConfigConfigurationRecorder_RecordingGroup struct {

	// AllSupported AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-allsupported
	AllSupported bool `json:"AllSupported"`

	// IncludeGlobalResourceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-includeglobalresourcetypes
	IncludeGlobalResourceTypes bool `json:"IncludeGlobalResourceTypes"`

	// ResourceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-resourcetypes
	ResourceTypes []string `json:"ResourceTypes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigurationRecorder_RecordingGroup) AWSCloudFormationType() string {
	return "AWS::Config::ConfigurationRecorder.RecordingGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigurationRecorder_RecordingGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSConfigConfigurationRecorder_RecordingGroupResources retrieves all AWSConfigConfigurationRecorder_RecordingGroup items from a CloudFormation template
func GetAllAWSConfigConfigurationRecorder_RecordingGroup(template *Template) map[string]*AWSConfigConfigurationRecorder_RecordingGroup {

	results := map[string]*AWSConfigConfigurationRecorder_RecordingGroup{}
	for name, resource := range template.Resources {
		result := &AWSConfigConfigurationRecorder_RecordingGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSConfigConfigurationRecorder_RecordingGroupWithName retrieves all AWSConfigConfigurationRecorder_RecordingGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSConfigConfigurationRecorder_RecordingGroup(name string, template *Template) (*AWSConfigConfigurationRecorder_RecordingGroup, error) {

	result := &AWSConfigConfigurationRecorder_RecordingGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSConfigConfigurationRecorder_RecordingGroup{}, errors.New("resource not found")

}
