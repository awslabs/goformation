package cloudformation

import (
	"encoding/json"
)

// AWSRoute53RecordSetGroup_GeoLocation AWS CloudFormation Resource (AWS::Route53::RecordSetGroup.GeoLocation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
type AWSRoute53RecordSetGroup_GeoLocation struct {

	// ContinentCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordsetgroup-geolocation-continentcode
	ContinentCode *Value `json:"ContinentCode,omitempty"`

	// CountryCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-countrycode
	CountryCode *Value `json:"CountryCode,omitempty"`

	// SubdivisionCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-subdivisioncode
	SubdivisionCode *Value `json:"SubdivisionCode,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53RecordSetGroup_GeoLocation) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSetGroup.GeoLocation"
}

func (r *AWSRoute53RecordSetGroup_GeoLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
