package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSRDSDBSecurityGroupIngress AWS CloudFormation Resource (AWS::RDS::DBSecurityGroupIngress)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html
type AWSRDSDBSecurityGroupIngress struct {

	// CIDRIP AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-cidrip
	CIDRIP string `json:"CIDRIP"`

	// DBSecurityGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-dbsecuritygroupname
	DBSecurityGroupName string `json:"DBSecurityGroupName"`

	// EC2SecurityGroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-ec2securitygroupid
	EC2SecurityGroupId string `json:"EC2SecurityGroupId"`

	// EC2SecurityGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-ec2securitygroupname
	EC2SecurityGroupName string `json:"EC2SecurityGroupName"`

	// EC2SecurityGroupOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-ec2securitygroupownerid
	EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSDBSecurityGroupIngress) AWSCloudFormationType() string {
	return "AWS::RDS::DBSecurityGroupIngress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSDBSecurityGroupIngress) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSDBSecurityGroupIngressResources retrieves all AWSRDSDBSecurityGroupIngress items from a CloudFormation template
func GetAllAWSRDSDBSecurityGroupIngressResources(template *Template) map[string]*AWSRDSDBSecurityGroupIngress {

	results := map[string]*AWSRDSDBSecurityGroupIngress{}
	for name, resource := range template.Resources {
		result := &AWSRDSDBSecurityGroupIngress{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSDBSecurityGroupIngressWithName retrieves all AWSRDSDBSecurityGroupIngress items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSRDSDBSecurityGroupIngressWithName(name string, template *Template) (*AWSRDSDBSecurityGroupIngress, error) {

	result := &AWSRDSDBSecurityGroupIngress{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSDBSecurityGroupIngress{}, errors.New("resource not found")

}
