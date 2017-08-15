package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::TargetGroup.TargetDescription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html
type AWSElasticLoadBalancingV2TargetGroup_TargetDescription struct {

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-id
	Id string `json:"Id"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-port
	Port int64 `json:"Port"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2TargetGroup_TargetDescription) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::TargetGroup.TargetDescription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2TargetGroup_TargetDescription) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2TargetGroup_TargetDescriptionResources retrieves all AWSElasticLoadBalancingV2TargetGroup_TargetDescription items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2TargetGroup_TargetDescription(template *Template) map[string]*AWSElasticLoadBalancingV2TargetGroup_TargetDescription {

	results := map[string]*AWSElasticLoadBalancingV2TargetGroup_TargetDescription{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2TargetGroup_TargetDescription{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2TargetGroup_TargetDescriptionWithName retrieves all AWSElasticLoadBalancingV2TargetGroup_TargetDescription items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2TargetGroup_TargetDescription(name string, template *Template) (*AWSElasticLoadBalancingV2TargetGroup_TargetDescription, error) {

	result := &AWSElasticLoadBalancingV2TargetGroup_TargetDescription{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2TargetGroup_TargetDescription{}, errors.New("resource not found")

}
