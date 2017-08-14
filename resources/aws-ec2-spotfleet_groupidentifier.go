package resources

// AWS::EC2::SpotFleet.GroupIdentifier AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html
type AWSEC2SpotFleetGroupIdentifier struct {

	// GroupId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html#cfn-ec2-spotfleet-groupidentifier-groupid
	GroupId string `json:"GroupId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleetGroupIdentifier) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.GroupIdentifier"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleetGroupIdentifier) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
