package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::AutoScaling::AutoScalingGroup.MetricsCollection AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
type AWSAutoScalingAutoScalingGroup_MetricsCollection struct {

	// Granularity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-granularity

	Granularity string `json:"Granularity"`

	// Metrics AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-metrics

	Metrics []string `json:"Metrics"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup_MetricsCollection) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.MetricsCollection"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingAutoScalingGroup_MetricsCollection) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSAutoScalingAutoScalingGroup_MetricsCollectionResources retrieves all AWSAutoScalingAutoScalingGroup_MetricsCollection items from a CloudFormation template
func GetAllAWSAutoScalingAutoScalingGroup_MetricsCollection(template *Template) map[string]*AWSAutoScalingAutoScalingGroup_MetricsCollection {

	results := map[string]*AWSAutoScalingAutoScalingGroup_MetricsCollection{}
	for name, resource := range template.Resources {
		result := &AWSAutoScalingAutoScalingGroup_MetricsCollection{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSAutoScalingAutoScalingGroup_MetricsCollectionWithName retrieves all AWSAutoScalingAutoScalingGroup_MetricsCollection items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSAutoScalingAutoScalingGroup_MetricsCollection(name string, template *Template) (*AWSAutoScalingAutoScalingGroup_MetricsCollection, error) {

	result := &AWSAutoScalingAutoScalingGroup_MetricsCollection{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSAutoScalingAutoScalingGroup_MetricsCollection{}, errors.New("resource not found")

}
