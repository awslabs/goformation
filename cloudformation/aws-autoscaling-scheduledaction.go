package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSAutoScalingScheduledAction AWS CloudFormation Resource (AWS::AutoScaling::ScheduledAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html
type AWSAutoScalingScheduledAction struct {

	// AutoScalingGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-asgname
	AutoScalingGroupName *Value `json:"AutoScalingGroupName,omitempty"`

	// DesiredCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-desiredcapacity
	DesiredCapacity *Value `json:"DesiredCapacity,omitempty"`

	// EndTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-endtime
	EndTime *Value `json:"EndTime,omitempty"`

	// MaxSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-maxsize
	MaxSize *Value `json:"MaxSize,omitempty"`

	// MinSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-minsize
	MinSize *Value `json:"MinSize,omitempty"`

	// Recurrence AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-recurrence
	Recurrence *Value `json:"Recurrence,omitempty"`

	// StartTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-starttime
	StartTime *Value `json:"StartTime,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingScheduledAction) AWSCloudFormationType() string {
	return "AWS::AutoScaling::ScheduledAction"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSAutoScalingScheduledAction) MarshalJSON() ([]byte, error) {
	type Properties AWSAutoScalingScheduledAction
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSAutoScalingScheduledAction) UnmarshalJSON(b []byte) error {
	type Properties AWSAutoScalingScheduledAction
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSAutoScalingScheduledAction(*res.Properties)
	}

	return nil
}

// GetAllAWSAutoScalingScheduledActionResources retrieves all AWSAutoScalingScheduledAction items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingScheduledActionResources() map[string]AWSAutoScalingScheduledAction {
	results := map[string]AWSAutoScalingScheduledAction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSAutoScalingScheduledAction:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AutoScaling::ScheduledAction" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSAutoScalingScheduledAction{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSAutoScalingScheduledActionWithName retrieves all AWSAutoScalingScheduledAction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingScheduledActionWithName(name string) (AWSAutoScalingScheduledAction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSAutoScalingScheduledAction:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AutoScaling::ScheduledAction" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSAutoScalingScheduledAction{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSAutoScalingScheduledAction{}, errors.New("resource not found")
}
