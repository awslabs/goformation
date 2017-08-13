package resources

// AWS::EC2::TrunkInterfaceAssociation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html
type AWSEC2TrunkInterfaceAssociation struct {
    
    // BranchInterfaceId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html#cfn-ec2-trunkinterfaceassociation-branchinterfaceid
    BranchInterfaceId string `json:"BranchInterfaceId"`
    
    // GREKey AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html#cfn-ec2-trunkinterfaceassociation-grekey
    GREKey int64 `json:"GREKey"`
    
    // TrunkInterfaceId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html#cfn-ec2-trunkinterfaceassociation-trunkinterfaceid
    TrunkInterfaceId string `json:"TrunkInterfaceId"`
    
    // VLANId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html#cfn-ec2-trunkinterfaceassociation-vlanid
    VLANId int64 `json:"VLANId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2TrunkInterfaceAssociation) AWSCloudFormationType() string {
    return "AWS::EC2::TrunkInterfaceAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2TrunkInterfaceAssociation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}