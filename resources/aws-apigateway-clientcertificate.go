package resources

// AWS::ApiGateway::ClientCertificate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html
type AWSApiGatewayClientCertificate struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html#cfn-apigateway-clientcertificate-description

	Description string `json:"Description"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayClientCertificate) AWSCloudFormationType() string {
	return "AWS::ApiGateway::ClientCertificate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayClientCertificate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
