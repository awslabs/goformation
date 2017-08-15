package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance.AssociationParameter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html
type AWSEC2Instance_AssociationParameter struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html#cfn-ec2-instance-ssmassociations-associationparameters-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html#cfn-ec2-instance-ssmassociations-associationparameters-value
	Value []string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_AssociationParameter) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.AssociationParameter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_AssociationParameter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2Instance_AssociationParameterResources retrieves all AWSEC2Instance_AssociationParameter items from a CloudFormation template
func GetAllAWSEC2Instance_AssociationParameter(template *Template) map[string]*AWSEC2Instance_AssociationParameter {

	results := map[string]*AWSEC2Instance_AssociationParameter{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance_AssociationParameter{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2Instance_AssociationParameterWithName retrieves all AWSEC2Instance_AssociationParameter items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Instance_AssociationParameter(name string, template *Template) (*AWSEC2Instance_AssociationParameter, error) {

	result := &AWSEC2Instance_AssociationParameter{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance_AssociationParameter{}, errors.New("resource not found")

}
