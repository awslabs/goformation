package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Step.HadoopJarStepConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html
type AWSEMRStep_HadoopJarStepConfig struct {

	// Args AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-args

	Args []string `json:"Args"`

	// Jar AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-jar

	Jar string `json:"Jar"`

	// MainClass AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-mainclass

	MainClass string `json:"MainClass"`

	// StepProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-stepproperties

	StepProperties []AWSEMRStep_KeyValue `json:"StepProperties"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRStep_HadoopJarStepConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Step.HadoopJarStepConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRStep_HadoopJarStepConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRStep_HadoopJarStepConfigResources retrieves all AWSEMRStep_HadoopJarStepConfig items from a CloudFormation template
func GetAllAWSEMRStep_HadoopJarStepConfig(template *Template) map[string]*AWSEMRStep_HadoopJarStepConfig {

	results := map[string]*AWSEMRStep_HadoopJarStepConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRStep_HadoopJarStepConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRStep_HadoopJarStepConfigWithName retrieves all AWSEMRStep_HadoopJarStepConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRStep_HadoopJarStepConfig(name string, template *Template) (*AWSEMRStep_HadoopJarStepConfig, error) {

	result := &AWSEMRStep_HadoopJarStepConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRStep_HadoopJarStepConfig{}, errors.New("resource not found")

}
