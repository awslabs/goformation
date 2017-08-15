package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SSM::Association AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html
type AWSSSMAssociation struct {

	// DocumentVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
	DocumentVersion string `json:"DocumentVersion"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
	InstanceId string `json:"InstanceId"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
	Name string `json:"Name"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
	Parameters map[string]AWSSSMAssociation_ParameterValues `json:"Parameters"`

	// ScheduleExpression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
	ScheduleExpression string `json:"ScheduleExpression"`

	// Targets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
	Targets []AWSSSMAssociation_Target `json:"Targets"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociation) AWSCloudFormationType() string {
	return "AWS::SSM::Association"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSSMAssociationResources retrieves all AWSSSMAssociation items from a CloudFormation template
func GetAllAWSSSMAssociation(template *Template) map[string]*AWSSSMAssociation {

	results := map[string]*AWSSSMAssociation{}
	for name, resource := range template.Resources {
		result := &AWSSSMAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSSMAssociationWithName retrieves all AWSSSMAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSSSMAssociation(name string, template *Template) (*AWSSSMAssociation, error) {

	result := &AWSSSMAssociation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSSMAssociation{}, errors.New("resource not found")

}
