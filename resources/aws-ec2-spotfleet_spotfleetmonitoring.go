package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SpotFleet.SpotFleetMonitoring AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html
type AWSEC2SpotFleet_SpotFleetMonitoring struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html#cfn-ec2-spotfleet-spotfleetmonitoring-enabled
	Enabled bool `json:"Enabled"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_SpotFleetMonitoring) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.SpotFleetMonitoring"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet_SpotFleetMonitoring) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SpotFleet_SpotFleetMonitoringResources retrieves all AWSEC2SpotFleet_SpotFleetMonitoring items from a CloudFormation template
func GetAllAWSEC2SpotFleet_SpotFleetMonitoring(template *Template) map[string]*AWSEC2SpotFleet_SpotFleetMonitoring {

	results := map[string]*AWSEC2SpotFleet_SpotFleetMonitoring{}
	for name, resource := range template.Resources {
		result := &AWSEC2SpotFleet_SpotFleetMonitoring{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SpotFleet_SpotFleetMonitoringWithName retrieves all AWSEC2SpotFleet_SpotFleetMonitoring items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2SpotFleet_SpotFleetMonitoring(name string, template *Template) (*AWSEC2SpotFleet_SpotFleetMonitoring, error) {

	result := &AWSEC2SpotFleet_SpotFleetMonitoring{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SpotFleet_SpotFleetMonitoring{}, errors.New("resource not found")

}
