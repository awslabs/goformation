package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::RDS::DBSecurityGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html
type AWSRDSDBSecurityGroup struct {

	// DBSecurityGroupIngress AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html#cfn-rds-dbsecuritygroup-dbsecuritygroupingress

	DBSecurityGroupIngress []AWSRDSDBSecurityGroup_Ingress `json:"DBSecurityGroupIngress"`

	// EC2VpcId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html#cfn-rds-dbsecuritygroup-ec2vpcid

	EC2VpcId string `json:"EC2VpcId"`

	// GroupDescription AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html#cfn-rds-dbsecuritygroup-groupdescription

	GroupDescription string `json:"GroupDescription"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html#cfn-rds-dbsecuritygroup-tags

	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSDBSecurityGroup) AWSCloudFormationType() string {
	return "AWS::RDS::DBSecurityGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSDBSecurityGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSDBSecurityGroupResources retrieves all AWSRDSDBSecurityGroup items from a CloudFormation template
func GetAllAWSRDSDBSecurityGroup(template *Template) map[string]*AWSRDSDBSecurityGroup {

	results := map[string]*AWSRDSDBSecurityGroup{}
	for name, resource := range template.Resources {
		result := &AWSRDSDBSecurityGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSDBSecurityGroupWithName retrieves all AWSRDSDBSecurityGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRDSDBSecurityGroup(name string, template *Template) (*AWSRDSDBSecurityGroup, error) {

	result := &AWSRDSDBSecurityGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSDBSecurityGroup{}, errors.New("resource not found")

}
