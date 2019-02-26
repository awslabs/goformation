package cloudformation

// AWSApiGatewayV2RouteResponse_ParameterConstraints AWS CloudFormation Resource (AWS::ApiGatewayV2::RouteResponse.ParameterConstraints)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routeresponse-parameterconstraints.html
type AWSApiGatewayV2RouteResponse_ParameterConstraints struct {

	// Required AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routeresponse-parameterconstraints.html#cfn-apigatewayv2-routeresponse-parameterconstraints-required
	Required bool `json:"Required,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayV2RouteResponse_ParameterConstraints) AWSCloudFormationType() string {
	return "AWS::ApiGatewayV2::RouteResponse.ParameterConstraints"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSApiGatewayV2RouteResponse_ParameterConstraints) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
