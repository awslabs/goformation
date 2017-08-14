package resources

// AWS::EC2::SpotFleet.SpotFleetMonitoring AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html
type AWSEC2SpotFleet_SpotFleetMonitoring struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html#cfn-ec2-spotfleet-spotfleetmonitoring-enabled

	Enabled bool `json:"Enabled"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_SpotFleetMonitoring) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.SpotFleetMonitoring"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet_SpotFleetMonitoring) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
