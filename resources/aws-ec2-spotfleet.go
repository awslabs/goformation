package resources

// AWS::EC2::SpotFleet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html
type AWSEC2SpotFleet struct {
    
    // SpotFleetRequestConfigData AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata
    SpotFleetRequestConfigData AWSEC2SpotFleetSpotFleetRequestConfigData `json:"SpotFleetRequestConfigData"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet) AWSCloudFormationType() string {
    return "AWS::EC2::SpotFleet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}