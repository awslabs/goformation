package cloudformation

// AWSApiGatewayDeployment_CanarySetting AWS CloudFormation Resource (AWS::ApiGateway::Deployment.CanarySetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html
type AWSApiGatewayDeployment_CanarySetting struct {

	// PercentTraffic AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-percenttraffic
	PercentTraffic float64 `json:"PercentTraffic,omitempty"`

	// StageVariableOverrides AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-stagevariableoverrides
	StageVariableOverrides map[string]string `json:"StageVariableOverrides,omitempty"`

	// UseStageCache AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-usestagecache
	UseStageCache bool `json:"UseStageCache,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDeployment_CanarySetting) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Deployment.CanarySetting"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSApiGatewayDeployment_CanarySetting) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
