package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSServiceCatalogPortfolioProductAssociation AWS CloudFormation Resource (AWS::ServiceCatalog::PortfolioProductAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html
type AWSServiceCatalogPortfolioProductAssociation struct {

	// AcceptLanguage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html#cfn-servicecatalog-portfolioproductassociation-acceptlanguage
	AcceptLanguage *Value `json:"AcceptLanguage,omitempty"`

	// PortfolioId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html#cfn-servicecatalog-portfolioproductassociation-portfolioid
	PortfolioId *Value `json:"PortfolioId,omitempty"`

	// ProductId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html#cfn-servicecatalog-portfolioproductassociation-productid
	ProductId *Value `json:"ProductId,omitempty"`

	// SourcePortfolioId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html#cfn-servicecatalog-portfolioproductassociation-sourceportfolioid
	SourcePortfolioId *Value `json:"SourcePortfolioId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceCatalogPortfolioProductAssociation) AWSCloudFormationType() string {
	return "AWS::ServiceCatalog::PortfolioProductAssociation"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSServiceCatalogPortfolioProductAssociation) MarshalJSON() ([]byte, error) {
	type Properties AWSServiceCatalogPortfolioProductAssociation
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
func (r *AWSServiceCatalogPortfolioProductAssociation) UnmarshalJSON(b []byte) error {
	type Properties AWSServiceCatalogPortfolioProductAssociation
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
		*r = AWSServiceCatalogPortfolioProductAssociation(*res.Properties)
	}

	return nil
}

// GetAllAWSServiceCatalogPortfolioProductAssociationResources retrieves all AWSServiceCatalogPortfolioProductAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioProductAssociationResources() map[string]AWSServiceCatalogPortfolioProductAssociation {
	results := map[string]AWSServiceCatalogPortfolioProductAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSServiceCatalogPortfolioProductAssociation:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceCatalog::PortfolioProductAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSServiceCatalogPortfolioProductAssociation{}
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

// GetAWSServiceCatalogPortfolioProductAssociationWithName retrieves all AWSServiceCatalogPortfolioProductAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioProductAssociationWithName(name string) (AWSServiceCatalogPortfolioProductAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSServiceCatalogPortfolioProductAssociation:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceCatalog::PortfolioProductAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSServiceCatalogPortfolioProductAssociation{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSServiceCatalogPortfolioProductAssociation{}, errors.New("resource not found")
}
