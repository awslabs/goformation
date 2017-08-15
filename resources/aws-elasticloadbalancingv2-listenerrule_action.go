package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::ListenerRule.Action AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html
type AWSElasticLoadBalancingV2ListenerRule_Action struct {

	// TargetGroupArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-targetgrouparn
	TargetGroupArn string `json:"TargetGroupArn"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2ListenerRule_Action) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::ListenerRule.Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2ListenerRule_Action) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
