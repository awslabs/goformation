package resources

// AWS::KinesisAnalytics::Application.MappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html
type AWSKinesisAnalyticsApplicationMappingParameters struct {
    
    // CSVMappingParameters AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html#cfn-kinesisanalytics-application-mappingparameters-csvmappingparameters
    CSVMappingParameters AWSKinesisAnalyticsApplicationMappingParametersCSVMappingParameters `json:"CSVMappingParameters"`
    
    // JSONMappingParameters AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html#cfn-kinesisanalytics-application-mappingparameters-jsonmappingparameters
    JSONMappingParameters AWSKinesisAnalyticsApplicationMappingParametersJSONMappingParameters `json:"JSONMappingParameters"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationMappingParameters) AWSCloudFormationType() string {
    return "AWS::KinesisAnalytics::Application.MappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationMappingParameters) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}