package resources

// AWS::OpsWorks::Stack.RdsDbInstance AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html
type AWSOpsWorksStackRdsDbInstance struct {
    
    // DbPassword AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbpassword
    DbPassword string `json:"DbPassword"`
    
    // DbUser AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbuser
    DbUser string `json:"DbUser"`
    
    // RdsDbInstanceArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-rdsdbinstancearn
    RdsDbInstanceArn string `json:"RdsDbInstanceArn"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksStackRdsDbInstance) AWSCloudFormationType() string {
    return "AWS::OpsWorks::Stack.RdsDbInstance"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksStackRdsDbInstance) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}