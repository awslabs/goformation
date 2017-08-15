package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::Service.LoadBalancer AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html
type AWSECSService_LoadBalancer struct {

	// ContainerName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html#cfn-ecs-service-loadbalancers-containername
	ContainerName string `json:"ContainerName"`

	// ContainerPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html#cfn-ecs-service-loadbalancers-containerport
	ContainerPort int64 `json:"ContainerPort"`

	// LoadBalancerName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html#cfn-ecs-service-loadbalancers-loadbalancername
	LoadBalancerName string `json:"LoadBalancerName"`

	// TargetGroupArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html#cfn-ecs-service-loadbalancers-targetgrouparn
	TargetGroupArn string `json:"TargetGroupArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSService_LoadBalancer) AWSCloudFormationType() string {
	return "AWS::ECS::Service.LoadBalancer"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSService_LoadBalancer) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSService_LoadBalancerResources retrieves all AWSECSService_LoadBalancer items from a CloudFormation template
func GetAllAWSECSService_LoadBalancer(template *Template) map[string]*AWSECSService_LoadBalancer {

	results := map[string]*AWSECSService_LoadBalancer{}
	for name, resource := range template.Resources {
		result := &AWSECSService_LoadBalancer{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSService_LoadBalancerWithName retrieves all AWSECSService_LoadBalancer items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSService_LoadBalancer(name string, template *Template) (*AWSECSService_LoadBalancer, error) {

	result := &AWSECSService_LoadBalancer{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSService_LoadBalancer{}, errors.New("resource not found")

}
