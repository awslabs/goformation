package goformation_test

import (
	"encoding/json"

	"github.com/sanathkr/yaml"

	"github.com/awslabs/goformation"
	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/intrinsics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Goformation", func() {

	Context("with a Serverless function matching 2016-10-31 specification", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-function-2016-10-31.yaml")
		It("should successfully validate the SAM template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		functions := template.GetAllAWSServerlessFunctionResources()

		It("should have exactly one function", func() {
			Expect(functions).To(HaveLen(1))
			Expect(functions).To(HaveKey("Function20161031"))
		})

		f := functions["Function20161031"]

		It("should correctly parse all of the function properties", func() {

			Expect(f.Handler).To(Equal("file.method"))
			Expect(f.Runtime).To(Equal("nodejs"))
			Expect(f.FunctionName).To(Equal("functionname"))
			Expect(f.Description).To(Equal("description"))
			Expect(f.MemorySize).To(Equal(128))
			Expect(f.Timeout).To(Equal(30))
			Expect(f.Role).To(Equal("aws::arn::123456789012::some/role"))
			Expect(f.Policies.StringArray).To(PointTo(ContainElement("AmazonDynamoDBFullAccess")))
			Expect(f.Environment).ToNot(BeNil())
			Expect(f.Environment.Variables).To(HaveKeyWithValue("NAME", "VALUE"))

		})

		It("should correctly parse all of the function API event sources/endpoints", func() {

			Expect(f.Events).ToNot(BeNil())
			Expect(f.Events).To(HaveKey("TestApi"))
			Expect(f.Events["TestApi"].Type).To(Equal("Api"))
			Expect(f.Events["TestApi"].Properties.ApiEvent).ToNot(BeNil())

			event := f.Events["TestApi"].Properties.ApiEvent
			Expect(event.Method).To(Equal("any"))
			Expect(event.Path).To(Equal("/testing"))

		})

	})

	Context("with an AWS CloudFormation template that contains multiple resources", func() {

		Context("described as Go structs", func() {

			template := cloudformation.NewTemplate()

			template.Resources["MySNSTopic"] = cloudformation.AWSSNSTopic{
				DisplayName: "test-sns-topic-display-name",
				TopicName:   "test-sns-topic-name",
				Subscription: []cloudformation.AWSSNSTopic_Subscription{
					cloudformation.AWSSNSTopic_Subscription{
						Endpoint: "test-sns-topic-subscription-endpoint",
						Protocol: "test-sns-topic-subscription-protocol",
					},
				},
			}

			template.Resources["MyRoute53HostedZone"] = cloudformation.AWSRoute53HostedZone{
				Name: "example.com",
			}

			topics := template.GetAllAWSSNSTopicResources()
			It("should have one AWS::SNS::Topic resource", func() {
				Expect(topics).To(HaveLen(1))
				Expect(topics).To(HaveKey("MySNSTopic"))
			})

			topic, err := template.GetAWSSNSTopicWithName("MySNSTopic")
			It("should be able to find the AWS::SNS::Topic by name", func() {
				Expect(topic).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have the correct AWS::SNS::Topic values", func() {
				Expect(topic.DisplayName).To(Equal("test-sns-topic-display-name"))
				Expect(topic.TopicName).To(Equal("test-sns-topic-name"))
				Expect(topic.Subscription).To(HaveLen(1))
				Expect(topic.Subscription[0].Endpoint).To(Equal("test-sns-topic-subscription-endpoint"))
				Expect(topic.Subscription[0].Protocol).To(Equal("test-sns-topic-subscription-protocol"))
			})

			zones := template.GetAllAWSRoute53HostedZoneResources()
			It("should have one AWS::Route53::HostedZone resource", func() {
				Expect(zones).To(HaveLen(1))
				Expect(zones).To(HaveKey("MyRoute53HostedZone"))
			})

			zone, err := template.GetAWSRoute53HostedZoneWithName("MyRoute53HostedZone")
			It("should be able to find the AWS::Route53::HostedZone by name", func() {
				Expect(zone).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have the correct AWS::Route53::HostedZone values", func() {
				Expect(zone.Name).To(Equal("example.com"))
			})

		})

		Context("described as JSON", func() {

			template := []byte(`{"AWSTemplateFormatVersion":"2010-09-09","Resources":{"MyRoute53HostedZone":{"Type":"AWS::Route53::HostedZone","Properties":{"Name":"example.com"}},"MySNSTopic":{"Type":"AWS::SNS::Topic","Properties":{"DisplayName":"test-sns-topic-display-name","Subscription":[{"Endpoint":"test-sns-topic-subscription-endpoint","Protocol":"test-sns-topic-subscription-protocol"}],"TopicName":"test-sns-topic-name"}}}}`)

			expected := cloudformation.NewTemplate()

			expected.Resources["MySNSTopic"] = cloudformation.AWSSNSTopic{
				DisplayName: "test-sns-topic-display-name",
				TopicName:   "test-sns-topic-name",
				Subscription: []cloudformation.AWSSNSTopic_Subscription{
					cloudformation.AWSSNSTopic_Subscription{
						Endpoint: "test-sns-topic-subscription-endpoint",
						Protocol: "test-sns-topic-subscription-protocol",
					},
				},
			}

			expected.Resources["MyRoute53HostedZone"] = cloudformation.AWSRoute53HostedZone{
				Name: "example.com",
			}

			result, err := goformation.ParseJSON(template)
			It("should marshal to Go structs successfully", func() {
				Expect(err).To(BeNil())
			})

			topics := result.GetAllAWSSNSTopicResources()
			It("should have one AWS::SNS::Topic resource", func() {
				Expect(topics).To(HaveLen(1))
				Expect(topics).To(HaveKey("MySNSTopic"))
			})

			topic, err := result.GetAWSSNSTopicWithName("MySNSTopic")
			It("should be able to find the AWS::SNS::Topic by name", func() {
				Expect(topic).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have the correct AWS::SNS::Topic values", func() {
				Expect(topic.DisplayName).To(Equal("test-sns-topic-display-name"))
				Expect(topic.TopicName).To(Equal("test-sns-topic-name"))
				Expect(topic.Subscription).To(HaveLen(1))
				Expect(topic.Subscription[0].Endpoint).To(Equal("test-sns-topic-subscription-endpoint"))
				Expect(topic.Subscription[0].Protocol).To(Equal("test-sns-topic-subscription-protocol"))
			})

			zones := result.GetAllAWSRoute53HostedZoneResources()
			It("should have one AWS::Route53::HostedZone resource", func() {
				Expect(zones).To(HaveLen(1))
				Expect(zones).To(HaveKey("MyRoute53HostedZone"))
			})

			zone, err := result.GetAWSRoute53HostedZoneWithName("MyRoute53HostedZone")
			It("should be able to find the AWS::Route53::HostedZone by name", func() {
				Expect(zone).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have the correct AWS::Route53::HostedZone values", func() {
				Expect(zone.Name).To(Equal("example.com"))
			})

		})

	})

	Context("with the official AWS SAM example templates", func() {

		inputs := []string{
			"test/yaml/sam-official-samples/alexa_skill/template.yaml",
			"test/yaml/sam-official-samples/api_backend/template.yaml",
			"test/yaml/sam-official-samples/api_swagger_cors/template.yaml",
			"test/yaml/sam-official-samples/encryption_proxy/template.yaml",
			"test/yaml/sam-official-samples/hello_world/template.yaml",
			"test/yaml/sam-official-samples/inline_swagger/template.yaml",
			"test/yaml/sam-official-samples/iot_backend/template.yaml",
			"test/yaml/sam-official-samples/s3_processor/template.yaml",
			"test/yaml/sam-official-samples/schedule/template.yaml",
			"test/yaml/sam-official-samples/stream_processor/template.yaml",
		}

		for _, filename := range inputs {
			Context("including "+filename, func() {
				template, err := goformation.Open(filename)
				It("should successfully parse the SAM template", func() {
					Expect(err).To(BeNil())
					Expect(template).ShouldNot(BeNil())
				})
			})
		}

	})

	Context("with the default AWS CodeStar templates", func() {

		inputs := []string{
			"test/yaml/codestar/nodejs.yml",
			"test/yaml/codestar/python.yml",
			"test/yaml/codestar/java.yml",
		}

		for _, filename := range inputs {
			Context("including "+filename, func() {
				template, err := goformation.Open(filename)
				It("should successfully validate the SAM template", func() {
					Expect(err).To(BeNil())
					Expect(template).ShouldNot(BeNil())
				})
			})
		}
	})

	// pmaddox@ 2017-08-17:
	// Commented out until we have support for YAML tag intrinsic functions (e.g. !Sub)
	Context("with a YAML template containing intrinsic tags (e.g. !Sub)", func() {

		template, err := goformation.Open("test/yaml/yaml-intrinsic-tags.yaml")
		It("should successfully validate the SAM template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(PointTo(BeNil()))
		})

		function, err := template.GetAWSServerlessFunctionWithName("IntrinsicFunctionTest")
		It("should have a function named 'IntrinsicFunctionTest'", func() {
			Expect(function).To(Not(BeNil()))
			Expect(err).To(BeNil())
		})

		It("it should have the correct values", func() {
			Expect(function.Runtime).To(Equal("4.3"))
			Expect(function.Timeout).To(Equal(10))
		})

	})

	Context("with a Serverless template containing different CodeUri formats", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-function-string-or-s3-location.yaml")
		It("should successfully parse the template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		functions := template.GetAllAWSServerlessFunctionResources()

		It("should have exactly three functions", func() {
			Expect(functions).To(HaveLen(3))
			Expect(functions).To(HaveKey("CodeUriWithS3LocationSpecifiedAsString"))
			Expect(functions).To(HaveKey("CodeUriWithS3LocationSpecifiedAsObject"))
			Expect(functions).To(HaveKey("CodeUriWithString"))
		})

		f1 := functions["CodeUriWithS3LocationSpecifiedAsString"]
		It("should parse a CodeUri property with an S3 location specified as a string", func() {
			Expect(f1.CodeUri.String).To(PointTo(Equal("s3://testbucket/testkey.zip")))
		})

		f2 := functions["CodeUriWithS3LocationSpecifiedAsObject"]
		It("should parse a CodeUri property with an S3 location specified as an object", func() {
			Expect(f2.CodeUri.S3Location.Key).To(Equal("testkey.zip"))
			Expect(f2.CodeUri.S3Location.Version).To(Equal(5))
		})

		f3 := functions["CodeUriWithString"]
		It("should parse a CodeUri property with a string", func() {
			Expect(f3.CodeUri.String).To(PointTo(Equal("./testfolder")))
		})

	})

	Context("with a template defined as Go code", func() {

		template := &cloudformation.Template{
			Resources: map[string]interface{}{
				"MyLambdaFunction": cloudformation.AWSLambdaFunction{
					Handler: "nodejs6.10",
				},
			},
		}

		functions := template.GetAllAWSLambdaFunctionResources()
		It("should be able to retrieve all Lambda functions with GetAllAWSLambdaFunction(template)", func() {
			Expect(functions).To(HaveLen(1))
		})

		function, err := template.GetAWSLambdaFunctionWithName("MyLambdaFunction")
		It("should be able to retrieve a specific Lambda function with GetAWSLambdaFunctionWithName(template, name)", func() {
			Expect(err).To(BeNil())
			Expect(function).To(BeAssignableToTypeOf(cloudformation.AWSLambdaFunction{}))
		})

		It("should have the correct Handler property", func() {
			Expect(function.Handler).To(Equal("nodejs6.10"))
		})

	})

	Context("with a template that defines an AWS::Serverless::Function", func() {

		Context("that has a CodeUri property set as an S3 Location", func() {

			template := &cloudformation.Template{
				Resources: map[string]interface{}{
					"MySAMFunction": cloudformation.AWSServerlessFunction{
						Handler: "nodejs6.10",
						CodeUri: &cloudformation.AWSServerlessFunction_CodeUri{
							S3Location: &cloudformation.AWSServerlessFunction_S3Location{
								Bucket:  "test-bucket",
								Key:     "test-key",
								Version: 100,
							},
						},
					},
				},
			}

			function, err := template.GetAWSServerlessFunctionWithName("MySAMFunction")
			It("should have an AWS::Serverless::Function called MySAMFunction", func() {
				Expect(function).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have the correct S3 bucket/key/version", func() {
				Expect(function.CodeUri.S3Location.Bucket).To(Equal("test-bucket"))
				Expect(function.CodeUri.S3Location.Key).To(Equal("test-key"))
				Expect(function.CodeUri.S3Location.Version).To(Equal(100))
			})

		})

		Context("that has a CodeUri property set as a string", func() {

			codeuri := "./some-folder"
			template := &cloudformation.Template{
				Resources: map[string]interface{}{
					"MySAMFunction": cloudformation.AWSServerlessFunction{
						Handler: "nodejs6.10",
						CodeUri: &cloudformation.AWSServerlessFunction_CodeUri{
							String: &codeuri,
						},
					},
				},
			}

			function, err := template.GetAWSServerlessFunctionWithName("MySAMFunction")
			It("should have an AWS::Serverless::Function called MySAMFunction", func() {
				Expect(function).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have the correct CodeUri", func() {
				Expect(function.CodeUri.String).To(PointTo(Equal("./some-folder")))
			})

		})

	})

	Context("with a YAML template that contains AWS::Serverless::SimpleTable resource(s)", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-simpletable.yaml")

		It("should parse the template successfully", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		table, err := template.GetAWSServerlessSimpleTableWithName("TestSimpleTable")
		It("should have a table named 'TestSimpleTable'", func() {
			Expect(table).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have a primary key set", func() {
			Expect(table.PrimaryKey).ToNot(BeNil())
		})

		It("should have the correct value for the primary key name", func() {
			Expect(table.PrimaryKey.Name).To(Equal("test-primary-key-name"))
		})

		It("should have the correct value for the primary key type", func() {
			Expect(table.PrimaryKey.Type).To(Equal("test-primary-key-type"))
		})

		It("should have provisioned throughput set", func() {
			Expect(table.ProvisionedThroughput).ToNot(BeNil())
		})

		It("should have the correct value for ReadCapacityUnits", func() {
			Expect(table.ProvisionedThroughput.ReadCapacityUnits).To(Equal(100))
		})

		It("should have the correct value for WriteCapacityUnits", func() {
			Expect(table.ProvisionedThroughput.WriteCapacityUnits).To(Equal(200))
		})

		It("should have a table named 'TestSimpleTableNoProperties'", func() {
			nopropertiesTable, err := template.GetAWSServerlessSimpleTableWithName("TestSimpleTableNoProperties")
			Expect(nopropertiesTable).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

	})

	Context("with a YAML template that contains AWS::Serverless::Api resource(s)", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-api.yaml")

		It("should parse the template successfully", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		api1, err := template.GetAWSServerlessApiWithName("ServerlessApiWithDefinitionUriAsString")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionUriAsString'", func() {
			Expect(api1).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct value for Name", func() {
			Expect(api1.Name).To(Equal("test-name"))
		})

		It("should have the correct value for StageName", func() {
			Expect(api1.StageName).To(Equal("test-stage-name"))
		})

		It("should have the correct value for DefinitionUri", func() {
			Expect(api1.DefinitionUri.String).To(PointTo(Equal("test-definition-uri")))
		})

		It("should have the correct value for CacheClusterEnabled", func() {
			Expect(api1.CacheClusterEnabled).To(Equal(true))
		})

		It("should have the correct value for CacheClusterSize", func() {
			Expect(api1.CacheClusterSize).To(Equal("test-cache-cluster-size"))
		})

		It("should have the correct value for Variables", func() {
			Expect(api1.Variables).To(HaveKeyWithValue("NAME", "VALUE"))
		})

		api2, err := template.GetAWSServerlessApiWithName("ServerlessApiWithDefinitionUriAsS3Location")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionUriAsS3Location'", func() {
			Expect(api2).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct value for DefinitionUri", func() {
			Expect(api2.DefinitionUri.S3Location.Bucket).To(Equal("test-bucket"))
			Expect(api2.DefinitionUri.S3Location.Key).To(Equal("test-key"))
			Expect(api2.DefinitionUri.S3Location.Version).To(Equal(1))
		})

		api3, err := template.GetAWSServerlessApiWithName("ServerlessApiWithDefinitionBodyAsJSON")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionBodyAsJSON'", func() {
			Expect(api3).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct value for DefinitionBody", func() {
			Expect(api3.DefinitionBody).To(Equal("{\n  \"DefinitionKey\": \"test-definition-value\"\n}\n"))
		})

		api4, err := template.GetAWSServerlessApiWithName("ServerlessApiWithDefinitionBodyAsYAML")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionBodyAsYAML'", func() {
			Expect(api4).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct value for DefinitionBody", func() {
			var expected map[string]interface{}
			expected = map[string]interface{}{
				"DefinitionKey": "test-definition-value",
			}
			Expect(api4.DefinitionBody).To(Equal(expected))
		})

	})

	Context("with a YAML template with paramter overrides", func() {

		template, err := goformation.OpenWithOptions("test/yaml/aws-serverless-function-env-vars.yaml", &intrinsics.ProcessorOptions{
			ParameterOverrides: map[string]interface{}{"ExampleParameter": "SomeNewValue"},
		})

		It("should successfully validate the SAM template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		function, err := template.GetAWSServerlessFunctionWithName("IntrinsicEnvironmentVariableTestFunction")
		It("should have a function named 'IntrinsicEnvironmentVariableTestFunction'", func() {
			Expect(function).To(Not(BeNil()))
			Expect(err).To(BeNil())
		})

		It("it should have the correct values", func() {
			Expect(function.Environment.Variables).To(HaveKeyWithValue("REF_ENV_VAR", "SomeNewValue"))
		})
	})

	Context("with a YAML template with processing disabled", func() {

		template, err := goformation.OpenWithOptions("test/yaml/aws-serverless-function-env-vars.yaml", &intrinsics.ProcessorOptions{
			NoProcess: true,
		})

		It("should successfully validate the SAM template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})
	})

	Context("with a SNS event source", func() {
		event := cloudformation.AWSServerlessFunction_Properties{
			SNSEvent: &cloudformation.AWSServerlessFunction_SNSEvent{
				Topic: "MyTopic",
			},
		}

		It("should marshal properties correctly", func() {
			bytes, err := event.MarshalJSON()
			Expect(err).To(BeNil())
			Expect(string(bytes)).To(Equal(`{"Topic":"MyTopic"}`))
		})
	})

	Context("with an SNS event source created from JSON", func() {
		eventString := `{"Topic":"MyTopic"}`
		eventJson := []byte(eventString)
		event := cloudformation.AWSServerlessFunction_Properties{}
		event.UnmarshalJSON(eventJson)

		It("should marshal properties correctly", func() {
			bytes, err := event.MarshalJSON()
			Expect(err).To(BeNil())
			Expect(string(bytes)).To(Equal(eventString))
		})
	})

	Context("with a template that contains a reference to another resource within the template", func() {

		template := &cloudformation.Template{
			Resources: map[string]interface{}{
				"TestBucket": cloudformation.AWSS3Bucket{
					BucketName: "test-bucket",
				},
				"TestBucketPolicy": cloudformation.AWSS3BucketPolicy{
					Bucket: cloudformation.Ref("TestBucket"),
				},
			},
		}

		It("should have the correct reference object when converted to JSON", func() {

			data, err := template.JSON()
			Expect(err).To(BeNil())

			var result map[string]interface{}
			if err := json.Unmarshal(data, &result); err != nil {
				Fail(err.Error())
			}

			resources, ok := result["Resources"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			bucket, ok := resources["TestBucketPolicy"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			properties, ok := bucket["Properties"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			reference, ok := properties["Bucket"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			Expect(reference["Ref"]).To(Equal("TestBucket"))

		})

		It("should have the correct reference object when converted to YAML", func() {

			data, err := template.YAML()
			Expect(err).To(BeNil())

			var result map[string]interface{}
			if err := yaml.Unmarshal(data, &result); err != nil {
				Fail(err.Error())
			}

			resources, ok := result["Resources"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			bucket, ok := resources["TestBucketPolicy"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			properties, ok := bucket["Properties"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			reference, ok := properties["Bucket"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			Expect(reference["Ref"]).To(Equal("TestBucket"))

		})

	})

	Context("with a template that is composed with all of the intrinsics", func() {

		tests := []struct {
			Name     string
			Input    string
			Expected map[string]interface{}
		}{
			{
				Name:  "Ref",
				Input: cloudformation.Ref("test-reference"),
				Expected: map[string]interface{}{
					"Ref": "test-reference",
				},
			},
			{
				Name:  "Fn::GetAtt",
				Input: cloudformation.GetAtt("resource", "property"),
				Expected: map[string]interface{}{
					"Fn::GetAtt": []interface{}{"resource", "property"},
				},
			},
			{
				Name:  "Fn::ImportValue",
				Input: cloudformation.ImportValue("test-import"),
				Expected: map[string]interface{}{
					"Fn::ImportValue": "test-import",
				},
			},
			{
				Name:  "Fn::Base64",
				Input: cloudformation.Base64("test-base64"),
				Expected: map[string]interface{}{
					"Fn::Base64": "test-base64",
				},
			},
			{
				Name:  "Fn::Cidr",
				Input: cloudformation.CIDR("test-ip-block", "test-count", "test-cidr-bits"),
				Expected: map[string]interface{}{
					"Fn::Cidr": []interface{}{"test-ip-block", "test-count", "test-cidr-bits"},
				},
			},
			{
				Name:  "Fn::FindInMap",
				Input: cloudformation.FindInMap("test-map", "test-top-level-key", "test-second-level-key"),
				Expected: map[string]interface{}{
					"Fn::FindInMap": []interface{}{"test-map", "test-top-level-key", "test-second-level-key"},
				},
			},
			{
				Name:  "Fn::GetAZs",
				Input: cloudformation.GetAZs("test-region"),
				Expected: map[string]interface{}{
					"Fn::GetAZs": "test-region",
				},
			},
			{
				Name:  "Fn::Join",
				Input: cloudformation.Join("test-delimiter", []string{"test-join-value-1", "test-join-value-2"}),
				Expected: map[string]interface{}{
					"Fn::Join": []interface{}{
						"test-delimiter",
						[]interface{}{
							"test-join-value-1",
							"test-join-value-2",
						},
					},
				},
			},
			{
				Name:  "Fn::Select",
				Input: cloudformation.Select("test-index", []string{"test-select-value-1", "test-select-value-2"}),
				Expected: map[string]interface{}{
					"Fn::Select": []interface{}{
						"test-index",
						[]interface{}{
							"test-select-value-1",
							"test-select-value-2",
						},
					},
				},
			},
			{
				Name:  "Fn::Split",
				Input: cloudformation.Split("test-delimiter", "test-split-source"),
				Expected: map[string]interface{}{
					"Fn::Split": []interface{}{"test-delimiter", "test-split-source"},
				},
			},
			{
				Name:  "Fn::Sub",
				Input: cloudformation.Sub("test-sub"),
				Expected: map[string]interface{}{
					"Fn::Sub": "test-sub",
				},
			},
			{
				Name:  "Fn::And",
				Input: cloudformation.And([]string{"test-and-first", "test-and-second", "test-and-third"}),
				Expected: map[string]interface{}{
					"Fn::And": []interface{}{"test-and-first", "test-and-second", "test-and-third"},
				},
			},
			{
				Name:  "Fn::Equals",
				Input: cloudformation.Equals("test-equals-value1", "test-equals-value2"),
				Expected: map[string]interface{}{
					"Fn::Equals": []interface{}{"test-equals-value1", "test-equals-value2"},
				},
			},
			{
				Name:  "Fn::If",
				Input: cloudformation.If("test-if-value1", "test-iftrue-value", "test-ifnottrue-value"),
				Expected: map[string]interface{}{
					"Fn::If": []interface{}{"test-if-value1", "test-iftrue-value", "test-ifnottrue-value"},
				},
			},
			{
				Name:  "Fn::Not",
				Input: cloudformation.Not([]string{"test-not-value1", "test-not-value2", "test-not-value3"}),
				Expected: map[string]interface{}{
					"Fn::Not": []interface{}{"test-not-value1", "test-not-value2", "test-not-value3"},
				},
			},
			{
				Name:  "Fn::Or",
				Input: cloudformation.Or([]string{"test-or-value1", "test-or-value2", "test-or-value3"}),
				Expected: map[string]interface{}{
					"Fn::Or": []interface{}{"test-or-value1", "test-or-value2", "test-or-value3"},
				},
			},
		}

		for _, test := range tests {
			test := test // https://github.com/onsi/ginkgo/issues/175
			It(test.Name+" should have the correct values", func() {

				template := &cloudformation.Template{
					Resources: map[string]interface{}{
						"Intrinsic_" + test.Name: test.Input,
					},
				}

				data, _ := template.JSON()
				var result map[string]interface{}
				json.Unmarshal(data, &result)

				resources, ok := result["Resources"].(map[string]interface{})
				Expect(ok).To(BeTrue())

				intr, ok := resources["Intrinsic_"+test.Name].(map[string]interface{})
				Expect(ok).To(BeTrue())
				Expect(intr).To(HaveLen(1))
				Expect(intr).To(BeEquivalentTo(test.Expected))

			})
		}

	})

	Context("with a template that contains nested intrinsics", func() {

		template := &cloudformation.Template{
			Resources: map[string]interface{}{
				"TestBucket": cloudformation.AWSS3Bucket{
					BucketName: cloudformation.Join("/", []string{
						cloudformation.Join("-", []string{"test", "bucket"}),
					}),
				},
			},
		}

		It("should have the correct nested intrinsics reference object when converted to JSON", func() {

			data, err := template.JSON()
			Expect(err).To(BeNil())

			var result map[string]interface{}
			if err := json.Unmarshal(data, &result); err != nil {
				Fail(err.Error())
			}

			resources, ok := result["Resources"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			bucket, ok := resources["TestBucket"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			properties, ok := bucket["Properties"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			value, ok := properties["BucketName"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			firstJoin, ok := value["Fn::Join"].([]interface{})
			Expect(ok).To(BeTrue())
			Expect(firstJoin).To(HaveLen(2))
			Expect(firstJoin[0]).To(Equal("/"))

			firstJoinElements, ok := firstJoin[1].([]interface{})
			Expect(ok).To(BeTrue())
			Expect(firstJoinElements).To(HaveLen(1))

			firstJoinFirstElement, ok := firstJoinElements[0].(map[string]interface{})
			Expect(ok).To(BeTrue())
			Expect(firstJoinFirstElement).To(HaveLen(1))

			secondJoin, ok := firstJoinFirstElement["Fn::Join"].([]interface{})
			Expect(ok).To(BeTrue())
			Expect(secondJoin).To(HaveLen(2))
			Expect(secondJoin[0]).To(Equal("-"))

			secondJoinElements, ok := secondJoin[1].([]interface{})
			Expect(secondJoinElements).To(HaveLen(2))
			Expect(secondJoinElements[0]).To(Equal("test"))
			Expect(secondJoinElements[1]).To(Equal("bucket"))

		})

	})

	Context("with a template that contains a Fn::GetAtt reference to another resource within the template", func() {

		template := &cloudformation.Template{
			Resources: map[string]interface{}{
				"TestBucket": cloudformation.AWSS3Bucket{
					BucketName: "test-bucket",
				},
				"TestBucketPolicy": cloudformation.AWSS3BucketPolicy{
					Bucket: cloudformation.GetAtt("TestBucket", "WebsiteURL"),
				},
			},
		}

		It("should have the correct Fn::GetAtt reference object when converted to JSON", func() {

			data, err := template.JSON()
			Expect(err).To(BeNil())

			var result map[string]interface{}
			if err := json.Unmarshal(data, &result); err != nil {
				Fail(err.Error())
			}

			resources, ok := result["Resources"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			bucket, ok := resources["TestBucketPolicy"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			properties, ok := bucket["Properties"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			reference, ok := properties["Bucket"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			Expect(reference["Fn::GetAtt"]).To(ContainElement("TestBucket"))
			Expect(reference["Fn::GetAtt"]).To(ContainElement("WebsiteURL"))

		})

		It("should have the correct Fn::GetAtt reference object when converted to YAML", func() {

			data, err := template.YAML()
			Expect(err).To(BeNil())

			var result map[string]interface{}
			if err := yaml.Unmarshal(data, &result); err != nil {
				Fail(err.Error())
			}

			resources, ok := result["Resources"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			bucket, ok := resources["TestBucketPolicy"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			properties, ok := bucket["Properties"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			reference, ok := properties["Bucket"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			Expect(reference["Fn::GetAtt"]).To(ContainElement("TestBucket"))
			Expect(reference["Fn::GetAtt"]).To(ContainElement("WebsiteURL"))

		})

	})

})
