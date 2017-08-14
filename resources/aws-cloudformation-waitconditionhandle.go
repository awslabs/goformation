package resources

// AWS::CloudFormation::WaitConditionHandle AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html
type AWSCloudFormationWaitConditionHandle struct {
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFormationWaitConditionHandle) AWSCloudFormationType() string {
	return "AWS::CloudFormation::WaitConditionHandle"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFormationWaitConditionHandle) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
