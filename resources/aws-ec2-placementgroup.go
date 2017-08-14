package resources

// AWS::EC2::PlacementGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html
type AWSEC2PlacementGroup struct {

	// Strategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-strategy

	Strategy string `json:"Strategy"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2PlacementGroup) AWSCloudFormationType() string {
	return "AWS::EC2::PlacementGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2PlacementGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
