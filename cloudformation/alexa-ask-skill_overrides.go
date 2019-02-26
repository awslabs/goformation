package cloudformation

// AlexaASKSkill_Overrides AWS CloudFormation Resource (Alexa::ASK::Skill.Overrides)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-overrides.html
type AlexaASKSkill_Overrides struct {

	// Manifest AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-overrides.html#cfn-ask-skill-overrides-manifest
	Manifest interface{} `json:"Manifest,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AlexaASKSkill_Overrides) AWSCloudFormationType() string {
	return "Alexa::ASK::Skill.Overrides"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AlexaASKSkill_Overrides) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
