package resources

// AWS::DynamoDB::Table.LocalSecondaryIndex AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html
type AWSDynamoDBTableLocalSecondaryIndex struct {

	// IndexName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-indexname
	IndexName string `json:"IndexName"`

	// KeySchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-keyschema
	KeySchema []AWSDynamoDBTableLocalSecondaryIndexKeySchema `json:"KeySchema"`

	// Projection AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-projection
	Projection AWSDynamoDBTableLocalSecondaryIndexProjection `json:"Projection"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTableLocalSecondaryIndex) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.LocalSecondaryIndex"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTableLocalSecondaryIndex) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
