package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::ElasticLoadBalancerAttachment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html
type AWSOpsWorksElasticLoadBalancerAttachment struct {

	// ElasticLoadBalancerName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html#cfn-opsworks-elbattachment-elbname
	ElasticLoadBalancerName string `json:"ElasticLoadBalancerName"`

	// LayerId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html#cfn-opsworks-elbattachment-layerid
	LayerId string `json:"LayerId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksElasticLoadBalancerAttachment) AWSCloudFormationType() string {
	return "AWS::OpsWorks::ElasticLoadBalancerAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksElasticLoadBalancerAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksElasticLoadBalancerAttachmentResources retrieves all AWSOpsWorksElasticLoadBalancerAttachment items from a CloudFormation template
func GetAllAWSOpsWorksElasticLoadBalancerAttachmentResources(template *Template) map[string]*AWSOpsWorksElasticLoadBalancerAttachment {

	results := map[string]*AWSOpsWorksElasticLoadBalancerAttachment{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksElasticLoadBalancerAttachment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksElasticLoadBalancerAttachmentWithName retrieves all AWSOpsWorksElasticLoadBalancerAttachment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSOpsWorksElasticLoadBalancerAttachmentWithName(name string, template *Template) (*AWSOpsWorksElasticLoadBalancerAttachment, error) {

	result := &AWSOpsWorksElasticLoadBalancerAttachment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksElasticLoadBalancerAttachment{}, errors.New("resource not found")

}
