package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSAppSyncDataSource AWS CloudFormation Resource (AWS::AppSync::DataSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html
type AWSAppSyncDataSource struct {

	// ApiId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-apiid
	ApiId *Value `json:"ApiId,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-description
	Description *Value `json:"Description,omitempty"`

	// DynamoDBConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-dynamodbconfig
	DynamoDBConfig *AWSAppSyncDataSource_DynamoDBConfig `json:"DynamoDBConfig,omitempty"`

	// ElasticsearchConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-elasticsearchconfig
	ElasticsearchConfig *AWSAppSyncDataSource_ElasticsearchConfig `json:"ElasticsearchConfig,omitempty"`

	// HttpConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-httpconfig
	HttpConfig *AWSAppSyncDataSource_HttpConfig `json:"HttpConfig,omitempty"`

	// LambdaConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-lambdaconfig
	LambdaConfig *AWSAppSyncDataSource_LambdaConfig `json:"LambdaConfig,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-name
	Name *Value `json:"Name,omitempty"`

	// ServiceRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-servicerolearn
	ServiceRoleArn *Value `json:"ServiceRoleArn,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-type
	Type *Value `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncDataSource) AWSCloudFormationType() string {
	return "AWS::AppSync::DataSource"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSAppSyncDataSource) MarshalJSON() ([]byte, error) {
	type Properties AWSAppSyncDataSource
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
func (r *AWSAppSyncDataSource) UnmarshalJSON(b []byte) error {
	type Properties AWSAppSyncDataSource
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
		*r = AWSAppSyncDataSource(*res.Properties)
	}

	return nil
}

// GetAllAWSAppSyncDataSourceResources retrieves all AWSAppSyncDataSource items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncDataSourceResources() map[string]AWSAppSyncDataSource {
	results := map[string]AWSAppSyncDataSource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSAppSyncDataSource:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AppSync::DataSource" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSAppSyncDataSource{}
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

// GetAWSAppSyncDataSourceWithName retrieves all AWSAppSyncDataSource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncDataSourceWithName(name string) (AWSAppSyncDataSource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSAppSyncDataSource:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AppSync::DataSource" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSAppSyncDataSource{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSAppSyncDataSource{}, errors.New("resource not found")
}
