package appflow

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// ConnectorProfile_RedshiftConnectorProfileProperties AWS CloudFormation Resource (AWS::AppFlow::ConnectorProfile.RedshiftConnectorProfileProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-redshiftconnectorprofileproperties.html
type ConnectorProfile_RedshiftConnectorProfileProperties struct {

	// BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-redshiftconnectorprofileproperties.html#cfn-appflow-connectorprofile-redshiftconnectorprofileproperties-bucketname
	BucketName string `json:"BucketName"`

	// BucketPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-redshiftconnectorprofileproperties.html#cfn-appflow-connectorprofile-redshiftconnectorprofileproperties-bucketprefix
	BucketPrefix *string `json:"BucketPrefix,omitempty"`

	// DatabaseUrl AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-redshiftconnectorprofileproperties.html#cfn-appflow-connectorprofile-redshiftconnectorprofileproperties-databaseurl
	DatabaseUrl string `json:"DatabaseUrl"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-redshiftconnectorprofileproperties.html#cfn-appflow-connectorprofile-redshiftconnectorprofileproperties-rolearn
	RoleArn string `json:"RoleArn"`

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
func (r *ConnectorProfile_RedshiftConnectorProfileProperties) AWSCloudFormationType() string {
	return "AWS::AppFlow::ConnectorProfile.RedshiftConnectorProfileProperties"
}
