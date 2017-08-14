package resources

// AWS::CodeDeploy::Application AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-application.html
type AWSCodeDeployApplication struct {

	// ApplicationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-application.html#cfn-codedeploy-application-applicationname

	ApplicationName string `json:"ApplicationName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployApplication) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::Application"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployApplication) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
