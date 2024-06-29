// Code generated by "go generate". Please don't change this file directly.

package connect

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// SecurityProfile_Application AWS CloudFormation Resource (AWS::Connect::SecurityProfile.Application)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-application.html
type SecurityProfile_Application struct {

	// ApplicationPermissions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-application.html#cfn-connect-securityprofile-application-applicationpermissions
	ApplicationPermissions []string `json:"ApplicationPermissions"`

	// Namespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-application.html#cfn-connect-securityprofile-application-namespace
	Namespace string `json:"Namespace"`

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
func (r *SecurityProfile_Application) AWSCloudFormationType() string {
	return "AWS::Connect::SecurityProfile.Application"
}