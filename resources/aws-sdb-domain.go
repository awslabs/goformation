package resources

// AWS::SDB::Domain AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html
type AWSSDBDomain struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html#cfn-sdb-domain-description

	Description string `json:"Description"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSDBDomain) AWSCloudFormationType() string {
	return "AWS::SDB::Domain"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSDBDomain) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
