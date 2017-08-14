package resources

// AWS::DynamoDB::Table.ProvisionedThroughput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
type AWSDynamoDBTableProvisionedThroughput struct {

	// ReadCapacityUnits AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-readcapacityunits
	ReadCapacityUnits int64 `json:"ReadCapacityUnits"`

	// WriteCapacityUnits AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-writecapacityunits
	WriteCapacityUnits int64 `json:"WriteCapacityUnits"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTableProvisionedThroughput) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.ProvisionedThroughput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTableProvisionedThroughput) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
