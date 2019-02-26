package cloudformation

// AWSDLMLifecyclePolicy_RetainRule AWS CloudFormation Resource (AWS::DLM::LifecyclePolicy.RetainRule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retainrule.html
type AWSDLMLifecyclePolicy_RetainRule struct {

	// Count AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retainrule.html#cfn-dlm-lifecyclepolicy-retainrule-count
	Count int `json:"Count,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDLMLifecyclePolicy_RetainRule) AWSCloudFormationType() string {
	return "AWS::DLM::LifecyclePolicy.RetainRule"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSDLMLifecyclePolicy_RetainRule) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
