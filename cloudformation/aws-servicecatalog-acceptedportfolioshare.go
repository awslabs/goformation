package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSServiceCatalogAcceptedPortfolioShare AWS CloudFormation Resource (AWS::ServiceCatalog::AcceptedPortfolioShare)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html
type AWSServiceCatalogAcceptedPortfolioShare struct {

	// AcceptLanguage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html#cfn-servicecatalog-acceptedportfolioshare-acceptlanguage
	AcceptLanguage *Value `json:"AcceptLanguage,omitempty"`

	// PortfolioId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html#cfn-servicecatalog-acceptedportfolioshare-portfolioid
	PortfolioId *Value `json:"PortfolioId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceCatalogAcceptedPortfolioShare) AWSCloudFormationType() string {
	return "AWS::ServiceCatalog::AcceptedPortfolioShare"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSServiceCatalogAcceptedPortfolioShare) MarshalJSON() ([]byte, error) {
	type Properties AWSServiceCatalogAcceptedPortfolioShare
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
func (r *AWSServiceCatalogAcceptedPortfolioShare) UnmarshalJSON(b []byte) error {
	type Properties AWSServiceCatalogAcceptedPortfolioShare
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
		*r = AWSServiceCatalogAcceptedPortfolioShare(*res.Properties)
	}

	return nil
}

// GetAllAWSServiceCatalogAcceptedPortfolioShareResources retrieves all AWSServiceCatalogAcceptedPortfolioShare items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogAcceptedPortfolioShareResources() map[string]AWSServiceCatalogAcceptedPortfolioShare {
	results := map[string]AWSServiceCatalogAcceptedPortfolioShare{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSServiceCatalogAcceptedPortfolioShare:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceCatalog::AcceptedPortfolioShare" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSServiceCatalogAcceptedPortfolioShare{}
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

// GetAWSServiceCatalogAcceptedPortfolioShareWithName retrieves all AWSServiceCatalogAcceptedPortfolioShare items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogAcceptedPortfolioShareWithName(name string) (AWSServiceCatalogAcceptedPortfolioShare, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSServiceCatalogAcceptedPortfolioShare:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceCatalog::AcceptedPortfolioShare" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSServiceCatalogAcceptedPortfolioShare{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSServiceCatalogAcceptedPortfolioShare{}, errors.New("resource not found")
}
