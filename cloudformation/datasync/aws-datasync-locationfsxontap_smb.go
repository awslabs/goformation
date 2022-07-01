// Code generated by "go generate". Please don't change this file directly.

package datasync

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// LocationFSxONTAP_SMB AWS CloudFormation Resource (AWS::DataSync::LocationFSxONTAP.SMB)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html
type LocationFSxONTAP_SMB struct {

	// Domain AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-domain
	Domain *string `json:"Domain,omitempty"`

	// MountOptions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-mountoptions
	MountOptions *LocationFSxONTAP_SmbMountOptions `json:"MountOptions"`

	// Password AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-password
	Password string `json:"Password"`

	// User AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-user
	User string `json:"User"`

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
func (r *LocationFSxONTAP_SMB) AWSCloudFormationType() string {
	return "AWS::DataSync::LocationFSxONTAP.SMB"
}
