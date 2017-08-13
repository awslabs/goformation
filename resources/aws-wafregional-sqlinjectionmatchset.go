package resources

// AWS::WAFRegional::SqlInjectionMatchSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html
type AWSWAFRegionalSqlInjectionMatchSet struct {
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html#cfn-wafregional-sqlinjectionmatchset-name
    Name string `json:"Name"`
    
    // SqlInjectionMatchTuples AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuples
    SqlInjectionMatchTuples []AWSWAFRegionalSqlInjectionMatchSetSqlInjectionMatchTuple `json:"SqlInjectionMatchTuples"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSqlInjectionMatchSet) AWSCloudFormationType() string {
    return "AWS::WAFRegional::SqlInjectionMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSqlInjectionMatchSet) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}