package ec2

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// LaunchTemplate_PrivateDnsNameOptions AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.PrivateDnsNameOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions.html
type LaunchTemplate_PrivateDnsNameOptions struct {

	// EnableResourceNameDnsAAAARecord AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions-enableresourcenamednsaaaarecord
	EnableResourceNameDnsAAAARecord *bool `json:"EnableResourceNameDnsAAAARecord,omitempty"`

	// EnableResourceNameDnsARecord AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions-enableresourcenamednsarecord
	EnableResourceNameDnsARecord *bool `json:"EnableResourceNameDnsARecord,omitempty"`

	// HostnameType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions-hostnametype
	HostnameType *string `json:"HostnameType,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *LaunchTemplate_PrivateDnsNameOptions) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.PrivateDnsNameOptions"
}
