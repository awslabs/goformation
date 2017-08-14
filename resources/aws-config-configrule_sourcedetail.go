package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Config::ConfigRule.SourceDetail AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html
type AWSConfigConfigRule_SourceDetail struct {

	// EventSource AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-source-sourcedetail-eventsource

	EventSource string `json:"EventSource"`

	// MaximumExecutionFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-sourcedetail-maximumexecutionfrequency

	MaximumExecutionFrequency string `json:"MaximumExecutionFrequency"`

	// MessageType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-source-sourcedetail-messagetype

	MessageType string `json:"MessageType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigRule_SourceDetail) AWSCloudFormationType() string {
	return "AWS::Config::ConfigRule.SourceDetail"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigRule_SourceDetail) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSConfigConfigRule_SourceDetailResources retrieves all AWSConfigConfigRule_SourceDetail items from a CloudFormation template
func GetAllAWSConfigConfigRule_SourceDetail(template *Template) map[string]*AWSConfigConfigRule_SourceDetail {

	results := map[string]*AWSConfigConfigRule_SourceDetail{}
	for name, resource := range template.Resources {
		result := &AWSConfigConfigRule_SourceDetail{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSConfigConfigRule_SourceDetailWithName retrieves all AWSConfigConfigRule_SourceDetail items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSConfigConfigRule_SourceDetail(name string, template *Template) (*AWSConfigConfigRule_SourceDetail, error) {

	result := &AWSConfigConfigRule_SourceDetail{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSConfigConfigRule_SourceDetail{}, errors.New("resource not found")

}
