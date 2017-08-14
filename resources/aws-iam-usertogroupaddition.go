package resources

// AWS::IAM::UserToGroupAddition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
type AWSIAMUserToGroupAddition struct {

	// GroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-groupname

	GroupName string `json:"GroupName"`

	// Users AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-users

	Users []string `json:"Users"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMUserToGroupAddition) AWSCloudFormationType() string {
	return "AWS::IAM::UserToGroupAddition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMUserToGroupAddition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
