package resources

// AWS::DirectoryService::SimpleAD AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html
type AWSDirectoryServiceSimpleAD struct {
    
    // CreateAlias AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-createalias
    CreateAlias bool `json:"CreateAlias"`
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-description
    Description string `json:"Description"`
    
    // EnableSso AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-enablesso
    EnableSso bool `json:"EnableSso"`
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-name
    Name string `json:"Name"`
    
    // Password AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-password
    Password string `json:"Password"`
    
    // ShortName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-shortname
    ShortName string `json:"ShortName"`
    
    // Size AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-size
    Size string `json:"Size"`
    
    // VpcSettings AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-vpcsettings
    VpcSettings AWSDirectoryServiceSimpleADVpcSettings `json:"VpcSettings"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceSimpleAD) AWSCloudFormationType() string {
    return "AWS::DirectoryService::SimpleAD"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDirectoryServiceSimpleAD) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}