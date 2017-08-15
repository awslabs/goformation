package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFRegionalWebACLAssociation AWS CloudFormation Resource (AWS::WAFRegional::WebACLAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html
type AWSWAFRegionalWebACLAssociation struct {

	// ResourceArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-resourcearn
	ResourceArn string `json:"ResourceArn"`

	// WebACLId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-webaclid
	WebACLId string `json:"WebACLId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalWebACLAssociation) AWSCloudFormationType() string {
	return "AWS::WAFRegional::WebACLAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalWebACLAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalWebACLAssociationResources retrieves all AWSWAFRegionalWebACLAssociation items from a CloudFormation template
func GetAllAWSWAFRegionalWebACLAssociationResources(template *Template) map[string]*AWSWAFRegionalWebACLAssociation {

	results := map[string]*AWSWAFRegionalWebACLAssociation{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalWebACLAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalWebACLAssociationWithName retrieves all AWSWAFRegionalWebACLAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSWAFRegionalWebACLAssociationWithName(name string, template *Template) (*AWSWAFRegionalWebACLAssociation, error) {

	result := &AWSWAFRegionalWebACLAssociation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalWebACLAssociation{}, errors.New("resource not found")

}
