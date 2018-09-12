package cloudformation

import (
	"encoding/json"
)

// AWSElasticLoadBalancingV2TargetGroup_TargetDescription AWS CloudFormation Resource (AWS::ElasticLoadBalancingV2::TargetGroup.TargetDescription)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html
type AWSElasticLoadBalancingV2TargetGroup_TargetDescription struct {

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-availabilityzone
	AvailabilityZone *Value `json:"AvailabilityZone,omitempty"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-id
	Id *Value `json:"Id,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-port
	Port *Value `json:"Port,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2TargetGroup_TargetDescription) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::TargetGroup.TargetDescription"
}

func (r *AWSElasticLoadBalancingV2TargetGroup_TargetDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
