package refactorspaces

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Service_UrlEndpointInput AWS CloudFormation Resource (AWS::RefactorSpaces::Service.UrlEndpointInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-service-urlendpointinput.html
type Service_UrlEndpointInput struct {

	// HealthUrl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-service-urlendpointinput.html#cfn-refactorspaces-service-urlendpointinput-healthurl
	HealthUrl *string `json:"HealthUrl,omitempty"`

	// Url AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-service-urlendpointinput.html#cfn-refactorspaces-service-urlendpointinput-url
	Url string `json:"Url"`

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
func (r *Service_UrlEndpointInput) AWSCloudFormationType() string {
	return "AWS::RefactorSpaces::Service.UrlEndpointInput"
}
