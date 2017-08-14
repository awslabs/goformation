package resources

// AWS::ApiGateway::ApiKey.StageKey AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html
type AWSApiGatewayApiKeyStageKey struct {

	// RestApiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-restapiid
	RestApiId string `json:"RestApiId"`

	// StageName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-stagename
	StageName string `json:"StageName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayApiKeyStageKey) AWSCloudFormationType() string {
	return "AWS::ApiGateway::ApiKey.StageKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayApiKeyStageKey) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
