package resources

// AWS::WAF::SqlInjectionMatchSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html
type AWSWAFSqlInjectionMatchSet struct {
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-name
    Name string `json:"Name"`
    
    // SqlInjectionMatchTuples AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples
    SqlInjectionMatchTuples []AWSWAFSqlInjectionMatchSetSqlInjectionMatchTuple `json:"SqlInjectionMatchTuples"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSqlInjectionMatchSet) AWSCloudFormationType() string {
    return "AWS::WAF::SqlInjectionMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFSqlInjectionMatchSet) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}