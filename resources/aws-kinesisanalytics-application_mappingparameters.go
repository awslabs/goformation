package resources

// AWS::KinesisAnalytics::Application.MappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html
type AWSKinesisAnalyticsApplication_MappingParameters struct {

	// CSVMappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html#cfn-kinesisanalytics-application-mappingparameters-csvmappingparameters
	CSVMappingParameters AWSKinesisAnalyticsApplication_CSVMappingParameters `json:"CSVMappingParameters"`

	// JSONMappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html#cfn-kinesisanalytics-application-mappingparameters-jsonmappingparameters
	JSONMappingParameters AWSKinesisAnalyticsApplication_JSONMappingParameters `json:"JSONMappingParameters"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_MappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.MappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_MappingParameters) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
