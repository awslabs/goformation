package frauddetector

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
	"github.com/awslabs/goformation/v5/cloudformation/tags"
)

// EventType_EntityType AWS CloudFormation Resource (AWS::FraudDetector::EventType.EntityType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-eventtype-entitytype.html
type EventType_EntityType struct {

	// Arn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-eventtype-entitytype.html#cfn-frauddetector-eventtype-entitytype-arn
	Arn string `json:"Arn,omitempty"`

	// CreatedTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-eventtype-entitytype.html#cfn-frauddetector-eventtype-entitytype-createdtime
	CreatedTime string `json:"CreatedTime,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-eventtype-entitytype.html#cfn-frauddetector-eventtype-entitytype-description
	Description string `json:"Description,omitempty"`

	// Inline AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-eventtype-entitytype.html#cfn-frauddetector-eventtype-entitytype-inline
	Inline bool `json:"Inline,omitempty"`

	// LastUpdatedTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-eventtype-entitytype.html#cfn-frauddetector-eventtype-entitytype-lastupdatedtime
	LastUpdatedTime string `json:"LastUpdatedTime,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-eventtype-entitytype.html#cfn-frauddetector-eventtype-entitytype-name
	Name string `json:"Name,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-eventtype-entitytype.html#cfn-frauddetector-eventtype-entitytype-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

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
func (r *EventType_EntityType) AWSCloudFormationType() string {
	return "AWS::FraudDetector::EventType.EntityType"
}
