package resources

// AWS::OpsWorks::Stack.ChefConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html
type AWSOpsWorksStackChefConfiguration struct {

	// BerkshelfVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html#cfn-opsworks-chefconfiguration-berkshelfversion
	BerkshelfVersion string `json:"BerkshelfVersion"`

	// ManageBerkshelf AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html#cfn-opsworks-chefconfiguration-berkshelfversion
	ManageBerkshelf bool `json:"ManageBerkshelf"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksStackChefConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Stack.ChefConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksStackChefConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
