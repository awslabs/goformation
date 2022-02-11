package backup

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Framework_ControlInputParameter AWS CloudFormation Resource (AWS::Backup::Framework.ControlInputParameter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-framework-controlinputparameter.html
type Framework_ControlInputParameter struct {

	// ParameterName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-framework-controlinputparameter.html#cfn-backup-framework-controlinputparameter-parametername
	ParameterName string `json:"ParameterName"`

	// ParameterValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-framework-controlinputparameter.html#cfn-backup-framework-controlinputparameter-parametervalue
	ParameterValue string `json:"ParameterValue"`

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
func (r *Framework_ControlInputParameter) AWSCloudFormationType() string {
	return "AWS::Backup::Framework.ControlInputParameter"
}
