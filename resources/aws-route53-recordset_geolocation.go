package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::RecordSet.GeoLocation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
type AWSRoute53RecordSet_GeoLocation struct {

	// ContinentCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-continentcode

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
func (r *AWSRoute53RecordSet_GeoLocation) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSet.GeoLocation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53RecordSet_GeoLocation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53RecordSet_GeoLocationResources retrieves all AWSRoute53RecordSet_GeoLocation items from a CloudFormation template
func GetAllAWSRoute53RecordSet_GeoLocation(template *Template) map[string]*AWSRoute53RecordSet_GeoLocation {

	results := map[string]*AWSRoute53RecordSet_GeoLocation{}
	for name, resource := range template.Resources {
		result := &AWSRoute53RecordSet_GeoLocation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53RecordSet_GeoLocationWithName retrieves all AWSRoute53RecordSet_GeoLocation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53RecordSet_GeoLocation(name string, template *Template) (*AWSRoute53RecordSet_GeoLocation, error) {

	result := &AWSRoute53RecordSet_GeoLocation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53RecordSet_GeoLocation{}, errors.New("resource not found")

}
