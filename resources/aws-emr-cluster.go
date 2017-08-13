package resources

// AWS::EMR::Cluster AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html
type AWSEMRCluster struct {
    
    // AdditionalInfo AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-additionalinfo
    AdditionalInfo string `json:"AdditionalInfo"`
    
    // Applications AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-applications
    Applications []AWSEMRClusterApplication `json:"Applications"`
    
    // AutoScalingRole AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-elasticmapreduce-cluster-autoscalingrole
    AutoScalingRole string `json:"AutoScalingRole"`
    
    // BootstrapActions AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-bootstrapactions
    BootstrapActions []AWSEMRClusterBootstrapActionConfig `json:"BootstrapActions"`
    
    // Configurations AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-configurations
    Configurations []AWSEMRClusterConfiguration `json:"Configurations"`
    
    // Instances AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-instances
    Instances AWSEMRClusterJobFlowInstancesConfig `json:"Instances"`
    
    // JobFlowRole AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-jobflowrole
    JobFlowRole string `json:"JobFlowRole"`
    
    // LogUri AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-loguri
    LogUri string `json:"LogUri"`
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-name
    Name string `json:"Name"`
    
    // ReleaseLabel AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-releaselabel
    ReleaseLabel string `json:"ReleaseLabel"`
    
    // ScaleDownBehavior AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-elasticmapreduce-cluster-scaledownbehavior
    ScaleDownBehavior string `json:"ScaleDownBehavior"`
    
    // SecurityConfiguration AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-securityconfiguration
    SecurityConfiguration string `json:"SecurityConfiguration"`
    
    // ServiceRole AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-servicerole
    ServiceRole string `json:"ServiceRole"`
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-elasticmapreduce-cluster-tags
    Tags []AWSEMRClusterTag `json:"Tags"`
    
    // VisibleToAllUsers AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-visibletoallusers
    VisibleToAllUsers bool `json:"VisibleToAllUsers"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}