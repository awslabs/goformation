package resources

// AWS::OpsWorks::Stack AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html
type AWSOpsWorksStack struct {
    
    // AgentVersion AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-agentversion
    AgentVersion string `json:"AgentVersion"`
    
    // Attributes AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-attributes
    Attributes map[string]string `json:"Attributes"`
    
    // ChefConfiguration AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-chefconfiguration
    ChefConfiguration AWSOpsWorksStackChefConfiguration `json:"ChefConfiguration"`
    
    // CloneAppIds AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-cloneappids
    CloneAppIds []string `json:"CloneAppIds"`
    
    // ClonePermissions AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-clonepermissions
    ClonePermissions bool `json:"ClonePermissions"`
    
    // ConfigurationManager AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-configmanager
    ConfigurationManager AWSOpsWorksStackStackConfigurationManager `json:"ConfigurationManager"`
    
    // CustomCookbooksSource AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-custcookbooksource
    CustomCookbooksSource AWSOpsWorksStackSource `json:"CustomCookbooksSource"`
    
    // CustomJson AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-custjson
    CustomJson string `json:"CustomJson"`
    
    // DefaultAvailabilityZone AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultaz
    DefaultAvailabilityZone string `json:"DefaultAvailabilityZone"`
    
    // DefaultInstanceProfileArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultinstanceprof
    DefaultInstanceProfileArn string `json:"DefaultInstanceProfileArn"`
    
    // DefaultOs AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultos
    DefaultOs string `json:"DefaultOs"`
    
    // DefaultRootDeviceType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultrootdevicetype
    DefaultRootDeviceType string `json:"DefaultRootDeviceType"`
    
    // DefaultSshKeyName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultsshkeyname
    DefaultSshKeyName string `json:"DefaultSshKeyName"`
    
    // DefaultSubnetId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#defaultsubnet
    DefaultSubnetId string `json:"DefaultSubnetId"`
    
    // EcsClusterArn AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-ecsclusterarn
    EcsClusterArn string `json:"EcsClusterArn"`
    
    // ElasticIps AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-elasticips
    ElasticIps []AWSOpsWorksStackElasticIp `json:"ElasticIps"`
    
    // HostnameTheme AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-hostnametheme
    HostnameTheme string `json:"HostnameTheme"`
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-name
    Name string `json:"Name"`
    
    // RdsDbInstances AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-rdsdbinstances
    RdsDbInstances []AWSOpsWorksStackRdsDbInstance `json:"RdsDbInstances"`
    
    // ServiceRoleArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-servicerolearn
    ServiceRoleArn string `json:"ServiceRoleArn"`
    
    // SourceStackId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-sourcestackid
    SourceStackId string `json:"SourceStackId"`
    
    // UseCustomCookbooks AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#usecustcookbooks
    UseCustomCookbooks bool `json:"UseCustomCookbooks"`
    
    // UseOpsworksSecurityGroups AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-useopsworkssecuritygroups
    UseOpsworksSecurityGroups bool `json:"UseOpsworksSecurityGroups"`
    
    // VpcId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-vpcid
    VpcId string `json:"VpcId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksStack) AWSCloudFormationType() string {
    return "AWS::OpsWorks::Stack"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksStack) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}