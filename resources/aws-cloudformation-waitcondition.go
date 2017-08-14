package resources

// AWS::CloudFormation::WaitCondition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html
type AWSCloudFormationWaitCondition struct {

	// Count AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-count

	Count int64 `json:"Count"`

	// Handle AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-handle

	Handle string `json:"Handle"`

	// Timeout AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-timeout

	Timeout string `json:"Timeout"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFormationWaitCondition) AWSCloudFormationType() string {
	return "AWS::CloudFormation::WaitCondition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFormationWaitCondition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
