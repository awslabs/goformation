package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Layer.Recipes AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html
type AWSOpsWorksLayer_Recipes struct {

	// Configure AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-configure

	Configure []string `json:"Configure"`

	// Deploy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-deploy

	Deploy []string `json:"Deploy"`

	// Setup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-setup

	Setup []string `json:"Setup"`

	// Shutdown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-shutdown

	Shutdown []string `json:"Shutdown"`

	// Undeploy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-undeploy

	Undeploy []string `json:"Undeploy"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_Recipes) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.Recipes"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayer_Recipes) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksLayer_RecipesResources retrieves all AWSOpsWorksLayer_Recipes items from a CloudFormation template
func GetAllAWSOpsWorksLayer_Recipes(template *Template) map[string]*AWSOpsWorksLayer_Recipes {

	results := map[string]*AWSOpsWorksLayer_Recipes{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksLayer_Recipes{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksLayer_RecipesWithName retrieves all AWSOpsWorksLayer_Recipes items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksLayer_Recipes(name string, template *Template) (*AWSOpsWorksLayer_Recipes, error) {

	result := &AWSOpsWorksLayer_Recipes{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksLayer_Recipes{}, errors.New("resource not found")

}
