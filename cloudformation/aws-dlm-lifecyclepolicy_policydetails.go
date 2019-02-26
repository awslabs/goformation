package cloudformation

// AWSDLMLifecyclePolicy_PolicyDetails AWS CloudFormation Resource (AWS::DLM::LifecyclePolicy.PolicyDetails)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html
type AWSDLMLifecyclePolicy_PolicyDetails struct {

	// ResourceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-resourcetypes
	ResourceTypes []string `json:"ResourceTypes,omitempty"`

	// Schedules AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-schedules
	Schedules []AWSDLMLifecyclePolicy_Schedule `json:"Schedules,omitempty"`

	// TargetTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-targettags
	TargetTags []Tag `json:"TargetTags,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDLMLifecyclePolicy_PolicyDetails) AWSCloudFormationType() string {
	return "AWS::DLM::LifecyclePolicy.PolicyDetails"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSDLMLifecyclePolicy_PolicyDetails) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
