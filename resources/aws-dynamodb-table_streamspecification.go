package resources

// AWS::DynamoDB::Table.StreamSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html
type AWSDynamoDBTable_StreamSpecification struct {

	// StreamViewType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html#cfn-dynamodb-streamspecification-streamviewtype

	StreamViewType string `json:"StreamViewType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_StreamSpecification) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.StreamSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_StreamSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
