package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer.AccessLoggingPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html
type AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy struct {

	// EmitInterval AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html#cfn-elb-accessloggingpolicy-emitinterval
	EmitInterval int64 `json:"EmitInterval"`

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html#cfn-elb-accessloggingpolicy-enabled
	Enabled bool `json:"Enabled"`

	// S3BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html#cfn-elb-accessloggingpolicy-s3bucketname
	S3BucketName string `json:"S3BucketName"`

	// S3BucketPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html#cfn-elb-accessloggingpolicy-s3bucketprefix
	S3BucketPrefix string `json:"S3BucketPrefix"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.AccessLoggingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicyResources retrieves all AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicyWithName retrieves all AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy, error) {

	result := &AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy{}, errors.New("resource not found")

}
