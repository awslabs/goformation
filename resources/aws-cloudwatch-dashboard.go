package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCloudWatchDashboard AWS CloudFormation Resource (AWS::CloudWatch::Dashboard)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html
type AWSCloudWatchDashboard struct {

	// DashboardBody AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html#cfn-cloudwatch-dashboard-dashboardbody
	DashboardBody string `json:"DashboardBody"`

	// DashboardName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html#cfn-cloudwatch-dashboard-dashboardname
	DashboardName string `json:"DashboardName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudWatchDashboard) AWSCloudFormationType() string {
	return "AWS::CloudWatch::Dashboard"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudWatchDashboard) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudWatchDashboardResources retrieves all AWSCloudWatchDashboard items from a CloudFormation template
func (t *Template) GetAllAWSCloudWatchDashboardResources() map[string]*AWSCloudWatchDashboard {

	results := map[string]*AWSCloudWatchDashboard{}
	for name, resource := range t.Resources {
		result := &AWSCloudWatchDashboard{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudWatchDashboardWithName retrieves all AWSCloudWatchDashboard items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudWatchDashboardWithName(name string) (*AWSCloudWatchDashboard, error) {

	result := &AWSCloudWatchDashboard{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudWatchDashboard{}, errors.New("resource not found")

}
