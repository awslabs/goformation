package cloudformation

// AWSApiGatewayStage_AccessLogSetting AWS CloudFormation Resource (AWS::ApiGateway::Stage.AccessLogSetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-accesslogsetting.html
type AWSApiGatewayStage_AccessLogSetting struct {

	// DestinationArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-accesslogsetting.html#cfn-apigateway-stage-accesslogsetting-destinationarn
	DestinationArn string `json:"DestinationArn,omitempty"`

	// Format AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-accesslogsetting.html#cfn-apigateway-stage-accesslogsetting-format
	Format string `json:"Format,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayStage_AccessLogSetting) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Stage.AccessLogSetting"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSApiGatewayStage_AccessLogSetting) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
