package resources

// AWS::ECR::Repository AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html
type AWSECRRepository struct {

	// RepositoryName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositoryname

	RepositoryName string `json:"RepositoryName"`

	// RepositoryPolicyText AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositorypolicytext

	RepositoryPolicyText interface{} `json:"RepositoryPolicyText"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECRRepository) AWSCloudFormationType() string {
	return "AWS::ECR::Repository"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECRRepository) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
