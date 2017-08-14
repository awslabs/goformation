package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Stack.ElasticIp AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-elasticip.html
type AWSOpsWorksStack_ElasticIp struct {

	// Ip AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-elasticip.html#cfn-opsworks-stack-elasticip-ip

	Ip string `json:"Ip"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-elasticip.html#cfn-opsworks-stack-elasticip-name

	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksStack_ElasticIp) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Stack.ElasticIp"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksStack_ElasticIp) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksStack_ElasticIpResources retrieves all AWSOpsWorksStack_ElasticIp items from a CloudFormation template
func GetAllAWSOpsWorksStack_ElasticIp(template *Template) map[string]*AWSOpsWorksStack_ElasticIp {

	results := map[string]*AWSOpsWorksStack_ElasticIp{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksStack_ElasticIp{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksStack_ElasticIpWithName retrieves all AWSOpsWorksStack_ElasticIp items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksStack_ElasticIp(name string, template *Template) (*AWSOpsWorksStack_ElasticIp, error) {

	result := &AWSOpsWorksStack_ElasticIp{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksStack_ElasticIp{}, errors.New("resource not found")

}
