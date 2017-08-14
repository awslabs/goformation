package resources

// AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html
type AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple-fieldtomatch

	FieldToMatch AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch `json:"FieldToMatch"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple-texttransformation

	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
