package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance.SsmAssociation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html
type AWSEC2Instance_SsmAssociation struct {

	// AssociationParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html#cfn-ec2-instance-ssmassociations-associationparameters

	AssociationParameters []AWSEC2Instance_AssociationParameter `json:"AssociationParameters"`

	// DocumentName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html#cfn-ec2-instance-ssmassociations-documentname

	DocumentName string `json:"DocumentName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_SsmAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.SsmAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_SsmAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2Instance_SsmAssociationResources retrieves all AWSEC2Instance_SsmAssociation items from a CloudFormation template
func GetAllAWSEC2Instance_SsmAssociation(template *Template) map[string]*AWSEC2Instance_SsmAssociation {

	results := map[string]*AWSEC2Instance_SsmAssociation{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance_SsmAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2Instance_SsmAssociationWithName retrieves all AWSEC2Instance_SsmAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Instance_SsmAssociation(name string, template *Template) (*AWSEC2Instance_SsmAssociation, error) {

	result := &AWSEC2Instance_SsmAssociation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance_SsmAssociation{}, errors.New("resource not found")

}
