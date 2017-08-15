package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::TargetGroup.Matcher AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html
type AWSElasticLoadBalancingV2TargetGroup_Matcher struct {

	// HttpCode AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html#cfn-elasticloadbalancingv2-targetgroup-matcher-httpcode
	HttpCode string `json:"HttpCode"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2TargetGroup_Matcher) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::TargetGroup.Matcher"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2TargetGroup_Matcher) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2TargetGroup_MatcherResources retrieves all AWSElasticLoadBalancingV2TargetGroup_Matcher items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2TargetGroup_Matcher(template *Template) map[string]*AWSElasticLoadBalancingV2TargetGroup_Matcher {

	results := map[string]*AWSElasticLoadBalancingV2TargetGroup_Matcher{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2TargetGroup_Matcher{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2TargetGroup_MatcherWithName retrieves all AWSElasticLoadBalancingV2TargetGroup_Matcher items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2TargetGroup_Matcher(name string, template *Template) (*AWSElasticLoadBalancingV2TargetGroup_Matcher, error) {

	result := &AWSElasticLoadBalancingV2TargetGroup_Matcher{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2TargetGroup_Matcher{}, errors.New("resource not found")

}
