package resources

// AWS::DynamoDB::Table.AttributeDefinition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html
type AWSDynamoDBTable_AttributeDefinition struct {

	// AttributeName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename

	AttributeName string `json:"AttributeName"`

	// AttributeType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename-attributetype

	AttributeType string `json:"AttributeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_AttributeDefinition) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.AttributeDefinition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_AttributeDefinition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
