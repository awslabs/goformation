package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSRedshiftClusterSecurityGroupIngress AWS CloudFormation Resource (AWS::Redshift::ClusterSecurityGroupIngress)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html
type AWSRedshiftClusterSecurityGroupIngress struct {

	// CIDRIP AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-cidrip
	CIDRIP string `json:"CIDRIP"`

	// ClusterSecurityGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-clustersecuritygroupname
	ClusterSecurityGroupName string `json:"ClusterSecurityGroupName"`

	// EC2SecurityGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-ec2securitygroupname
	EC2SecurityGroupName string `json:"EC2SecurityGroupName"`

	// EC2SecurityGroupOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-ec2securitygroupownerid
	EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterSecurityGroupIngress) AWSCloudFormationType() string {
	return "AWS::Redshift::ClusterSecurityGroupIngress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterSecurityGroupIngress) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRedshiftClusterSecurityGroupIngressResources retrieves all AWSRedshiftClusterSecurityGroupIngress items from a CloudFormation template
func GetAllAWSRedshiftClusterSecurityGroupIngressResources(template *Template) map[string]*AWSRedshiftClusterSecurityGroupIngress {

	results := map[string]*AWSRedshiftClusterSecurityGroupIngress{}
	for name, resource := range template.Resources {
		result := &AWSRedshiftClusterSecurityGroupIngress{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRedshiftClusterSecurityGroupIngressWithName retrieves all AWSRedshiftClusterSecurityGroupIngress items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSRedshiftClusterSecurityGroupIngressWithName(name string, template *Template) (*AWSRedshiftClusterSecurityGroupIngress, error) {

	result := &AWSRedshiftClusterSecurityGroupIngress{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRedshiftClusterSecurityGroupIngress{}, errors.New("resource not found")

}
