package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.Application AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html
type AWSEMRCluster_Application struct {

	// AdditionalInfo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-additionalinfo
	AdditionalInfo map[string]string `json:"AdditionalInfo"`

	// Args AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-args
	Args []string `json:"Args"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-name
	Name string `json:"Name"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-version
	Version string `json:"Version"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_Application) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.Application"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_Application) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_ApplicationResources retrieves all AWSEMRCluster_Application items from a CloudFormation template
func GetAllAWSEMRCluster_Application(template *Template) map[string]*AWSEMRCluster_Application {

	results := map[string]*AWSEMRCluster_Application{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_Application{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_ApplicationWithName retrieves all AWSEMRCluster_Application items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_Application(name string, template *Template) (*AWSEMRCluster_Application, error) {

	result := &AWSEMRCluster_Application{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_Application{}, errors.New("resource not found")

}
