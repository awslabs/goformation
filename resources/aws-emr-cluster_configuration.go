package resources

// AWS::EMR::Cluster.Configuration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html
type AWSEMRClusterConfiguration struct {
    
    // Classification AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-classification
    Classification string `json:"Classification"`
    
    // ConfigurationProperties AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurationproperties
    ConfigurationProperties map[string]string `json:"ConfigurationProperties"`
    
    // Configurations AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurations
    Configurations []AWSEMRClusterConfigurationConfiguration `json:"Configurations"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterConfiguration) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.Configuration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}