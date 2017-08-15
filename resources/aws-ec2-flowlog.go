package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2FlowLog AWS CloudFormation Resource (AWS::EC2::FlowLog)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html
type AWSEC2FlowLog struct {

	// DeliverLogsPermissionArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-deliverlogspermissionarn
	DeliverLogsPermissionArn string `json:"DeliverLogsPermissionArn"`

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-loggroupname
	LogGroupName string `json:"LogGroupName"`

	// ResourceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourceid
	ResourceId string `json:"ResourceId"`

	// ResourceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourcetype
	ResourceType string `json:"ResourceType"`

	// TrafficType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-traffictype
	TrafficType string `json:"TrafficType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2FlowLog) AWSCloudFormationType() string {
	return "AWS::EC2::FlowLog"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2FlowLog) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2FlowLogResources retrieves all AWSEC2FlowLog items from a CloudFormation template
func GetAllAWSEC2FlowLogResources(template *Template) map[string]*AWSEC2FlowLog {

	results := map[string]*AWSEC2FlowLog{}
	for name, resource := range template.Resources {
		result := &AWSEC2FlowLog{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2FlowLogWithName retrieves all AWSEC2FlowLog items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2FlowLogWithName(name string, template *Template) (*AWSEC2FlowLog, error) {

	result := &AWSEC2FlowLog{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2FlowLog{}, errors.New("resource not found")

}
