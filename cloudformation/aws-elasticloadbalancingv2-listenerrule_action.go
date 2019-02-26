package cloudformation

// AWSElasticLoadBalancingV2ListenerRule_Action AWS CloudFormation Resource (AWS::ElasticLoadBalancingV2::ListenerRule.Action)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html
type AWSElasticLoadBalancingV2ListenerRule_Action struct {

	// AuthenticateCognitoConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listenerrule-action-authenticatecognitoconfig
	AuthenticateCognitoConfig *AWSElasticLoadBalancingV2ListenerRule_AuthenticateCognitoConfig `json:"AuthenticateCognitoConfig,omitempty"`

	// AuthenticateOidcConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listenerrule-action-authenticateoidcconfig
	AuthenticateOidcConfig *AWSElasticLoadBalancingV2ListenerRule_AuthenticateOidcConfig `json:"AuthenticateOidcConfig,omitempty"`

	// FixedResponseConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listenerrule-action-fixedresponseconfig
	FixedResponseConfig *AWSElasticLoadBalancingV2ListenerRule_FixedResponseConfig `json:"FixedResponseConfig,omitempty"`

	// Order AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listenerrule-action-order
	Order int `json:"Order,omitempty"`

	// RedirectConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listenerrule-action-redirectconfig
	RedirectConfig *AWSElasticLoadBalancingV2ListenerRule_RedirectConfig `json:"RedirectConfig,omitempty"`

	// TargetGroupArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-targetgrouparn
	TargetGroupArn string `json:"TargetGroupArn,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-type
	Type string `json:"Type,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2ListenerRule_Action) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::ListenerRule.Action"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSElasticLoadBalancingV2ListenerRule_Action) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
