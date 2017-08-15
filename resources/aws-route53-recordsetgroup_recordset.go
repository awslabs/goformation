package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::RecordSetGroup.RecordSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html
type AWSRoute53RecordSetGroup_RecordSet struct {

	// AliasTarget AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-aliastarget
	AliasTarget AWSRoute53RecordSetGroup_AliasTarget `json:"AliasTarget"`

	// Comment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-comment
	Comment string `json:"Comment"`

	// Failover AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-failover
	Failover string `json:"Failover"`

	// GeoLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-geolocation
	GeoLocation AWSRoute53RecordSetGroup_GeoLocation `json:"GeoLocation"`

	// HealthCheckId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-healthcheckid
	HealthCheckId string `json:"HealthCheckId"`

	// HostedZoneId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzoneid
	HostedZoneId string `json:"HostedZoneId"`

	// HostedZoneName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzonename
	HostedZoneName string `json:"HostedZoneName"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-name
	Name string `json:"Name"`

	// Region AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-region
	Region string `json:"Region"`

	// ResourceRecords AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-resourcerecords
	ResourceRecords []string `json:"ResourceRecords"`

	// SetIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-setidentifier
	SetIdentifier string `json:"SetIdentifier"`

	// TTL AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-ttl
	TTL string `json:"TTL"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-type
	Type string `json:"Type"`

	// Weight AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-weight
	Weight int64 `json:"Weight"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53RecordSetGroup_RecordSet) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSetGroup.RecordSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53RecordSetGroup_RecordSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53RecordSetGroup_RecordSetResources retrieves all AWSRoute53RecordSetGroup_RecordSet items from a CloudFormation template
func GetAllAWSRoute53RecordSetGroup_RecordSet(template *Template) map[string]*AWSRoute53RecordSetGroup_RecordSet {

	results := map[string]*AWSRoute53RecordSetGroup_RecordSet{}
	for name, resource := range template.Resources {
		result := &AWSRoute53RecordSetGroup_RecordSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53RecordSetGroup_RecordSetWithName retrieves all AWSRoute53RecordSetGroup_RecordSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53RecordSetGroup_RecordSet(name string, template *Template) (*AWSRoute53RecordSetGroup_RecordSet, error) {

	result := &AWSRoute53RecordSetGroup_RecordSet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53RecordSetGroup_RecordSet{}, errors.New("resource not found")

}
