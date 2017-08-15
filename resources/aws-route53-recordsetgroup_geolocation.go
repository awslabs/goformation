package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::RecordSetGroup.GeoLocation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
type AWSRoute53RecordSetGroup_GeoLocation struct {

	// ContinentCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordsetgroup-geolocation-continentcode

	ContinentCode string `json:"ContinentCode"`

	// CountryCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-countrycode

	CountryCode string `json:"CountryCode"`

	// SubdivisionCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-subdivisioncode

	SubdivisionCode string `json:"SubdivisionCode"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53RecordSetGroup_GeoLocation) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSetGroup.GeoLocation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53RecordSetGroup_GeoLocation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53RecordSetGroup_GeoLocationResources retrieves all AWSRoute53RecordSetGroup_GeoLocation items from a CloudFormation template
func GetAllAWSRoute53RecordSetGroup_GeoLocation(template *Template) map[string]*AWSRoute53RecordSetGroup_GeoLocation {

	results := map[string]*AWSRoute53RecordSetGroup_GeoLocation{}
	for name, resource := range template.Resources {
		result := &AWSRoute53RecordSetGroup_GeoLocation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53RecordSetGroup_GeoLocationWithName retrieves all AWSRoute53RecordSetGroup_GeoLocation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53RecordSetGroup_GeoLocation(name string, template *Template) (*AWSRoute53RecordSetGroup_GeoLocation, error) {

	result := &AWSRoute53RecordSetGroup_GeoLocation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53RecordSetGroup_GeoLocation{}, errors.New("resource not found")

}
