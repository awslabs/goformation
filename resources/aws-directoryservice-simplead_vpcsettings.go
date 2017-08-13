package resources

// AWS::DirectoryService::SimpleAD.VpcSettings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html
type AWSDirectoryServiceSimpleADVpcSettings struct {
    
    // SubnetIds AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-subnetids
    SubnetIds []string `json:"SubnetIds"`
    
    // VpcId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-vpcid
    VpcId string `json:"VpcId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceSimpleADVpcSettings) AWSCloudFormationType() string {
    return "AWS::DirectoryService::SimpleAD.VpcSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDirectoryServiceSimpleADVpcSettings) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}