package resources

// AWS::EC2::InternetGateway AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html
type AWSEC2InternetGateway struct {
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html#cfn-ec2-internetgateway-tags
    Tags []AWSEC2InternetGatewayTag `json:"Tags"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2InternetGateway) AWSCloudFormationType() string {
    return "AWS::EC2::InternetGateway"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2InternetGateway) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}