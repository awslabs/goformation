package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::Cluster AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html
type AWSECSCluster struct {

	// ClusterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-clustername
	ClusterName string `json:"ClusterName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSCluster) AWSCloudFormationType() string {
	return "AWS::ECS::Cluster"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSCluster) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSClusterResources retrieves all AWSECSCluster items from a CloudFormation template
func GetAllAWSECSCluster(template *Template) map[string]*AWSECSCluster {

	results := map[string]*AWSECSCluster{}
	for name, resource := range template.Resources {
		result := &AWSECSCluster{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSClusterWithName retrieves all AWSECSCluster items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSCluster(name string, template *Template) (*AWSECSCluster, error) {

	result := &AWSECSCluster{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSCluster{}, errors.New("resource not found")

}
