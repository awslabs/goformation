package cloudformation

// AWSSecretsManagerRotationSchedule_RotationRules AWS CloudFormation Resource (AWS::SecretsManager::RotationSchedule.RotationRules)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-rotationrules.html
type AWSSecretsManagerRotationSchedule_RotationRules struct {

	// AutomaticallyAfterDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-rotationrules.html#cfn-secretsmanager-rotationschedule-rotationrules-automaticallyafterdays
	AutomaticallyAfterDays int `json:"AutomaticallyAfterDays,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSecretsManagerRotationSchedule_RotationRules) AWSCloudFormationType() string {
	return "AWS::SecretsManager::RotationSchedule.RotationRules"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSSecretsManagerRotationSchedule_RotationRules) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
