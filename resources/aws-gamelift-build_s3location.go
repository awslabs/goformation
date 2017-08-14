package resources

// AWS::GameLift::Build.S3Location AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html
type AWSGameLiftBuild_S3Location struct {

	// Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-bucket

	Bucket string `json:"Bucket"`

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-key

	Key string `json:"Key"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-rolearn

	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftBuild_S3Location) AWSCloudFormationType() string {
	return "AWS::GameLift::Build.S3Location"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSGameLiftBuild_S3Location) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
