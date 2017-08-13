package resources

// AWS::ElasticBeanstalk::ApplicationVersion AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html
type AWSElasticBeanstalkApplicationVersion struct {
    
    // ApplicationName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-applicationname
    ApplicationName string `json:"ApplicationName"`
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-description
    Description string `json:"Description"`
    
    // SourceBundle AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
    SourceBundle AWSElasticBeanstalkApplicationVersionSourceBundle `json:"SourceBundle"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplicationVersion) AWSCloudFormationType() string {
    return "AWS::ElasticBeanstalk::ApplicationVersion"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkApplicationVersion) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}