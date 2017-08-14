package resources

// AWS::IAM::InstanceProfile AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
type AWSIAMInstanceProfile struct {

	// InstanceProfileName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-instanceprofilename

	InstanceProfileName string `json:"InstanceProfileName"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-path

	Path string `json:"Path"`

	// Roles AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-roles

	Roles []string `json:"Roles"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMInstanceProfile) AWSCloudFormationType() string {
	return "AWS::IAM::InstanceProfile"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMInstanceProfile) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
