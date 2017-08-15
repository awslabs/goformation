package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::RDS::DBSecurityGroup.Ingress AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html
type AWSRDSDBSecurityGroup_Ingress struct {

	// CIDRIP AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-cidrip
	CIDRIP string `json:"CIDRIP"`

	// EC2SecurityGroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupid
	EC2SecurityGroupId string `json:"EC2SecurityGroupId"`

	// EC2SecurityGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupname
	EC2SecurityGroupName string `json:"EC2SecurityGroupName"`

	// EC2SecurityGroupOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupownerid
	EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSDBSecurityGroup_Ingress) AWSCloudFormationType() string {
	return "AWS::RDS::DBSecurityGroup.Ingress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSDBSecurityGroup_Ingress) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSDBSecurityGroup_IngressResources retrieves all AWSRDSDBSecurityGroup_Ingress items from a CloudFormation template
func GetAllAWSRDSDBSecurityGroup_Ingress(template *Template) map[string]*AWSRDSDBSecurityGroup_Ingress {

	results := map[string]*AWSRDSDBSecurityGroup_Ingress{}
	for name, resource := range template.Resources {
		result := &AWSRDSDBSecurityGroup_Ingress{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSDBSecurityGroup_IngressWithName retrieves all AWSRDSDBSecurityGroup_Ingress items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRDSDBSecurityGroup_Ingress(name string, template *Template) (*AWSRDSDBSecurityGroup_Ingress, error) {

	result := &AWSRDSDBSecurityGroup_Ingress{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSDBSecurityGroup_Ingress{}, errors.New("resource not found")

}
