package goformation_test

import (
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
			Expect(err).ToNot(HaveOccurred())
			Expect(template).ShouldNot(BeNil())
		})

		It("should have exactly one function as resource", func() {
			Expect(template.Resources).To(HaveLen(1))
			Expect(template.Resources).To(HaveKey("Function20161031"))
		})

		functions := template.GetAllAWSServerlessFunctionResources()

		It("should have exactly one function", func() {
			Expect(functions).To(HaveLen(1))
			Expect(functions).To(HaveKey("Function20161031"))
		})

		f := functions["Function20161031"]

		It("should correctly parse all of the function properties", func() {
			Expect(f.Handler).ToNot(BeNil())
			Expect(f.Handler.String()).To(Equal("file.method"))
			Expect(f.Runtime.String()).To(Equal("nodejs"))
			Expect(f.FunctionName.String()).To(Equal("functionname"))
			Expect(f.Description.String()).To(Equal("description"))
			Expect(f.MemorySize.String()).To(Equal("128"))
			Expect(f.Timeout.String()).To(Equal("30"))
			Expect(f.Role).To(Equal(cloudformation.NewString("aws::arn::123456789012::some/role")))
			Expect(f.Policies).ToNot(BeNil())
			Expect(f.Policies.Raw()).To(ContainElement("AmazonDynamoDBFullAccess"))
			Expect(f.Environment).ToNot(BeNil())
			Expect(f.Environment.Variables).To(HaveKeyWithValue("NAME", cloudformation.NewString("VALUE")))
		})

		It("should correctly parse all of the function API event sources/endpoints", func() {
			Expect(f.Events).ToNot(BeNil())
			Expect(f.Events).To(HaveKey("TestApi"))
			Expect(f.Events["TestApi"].Type).To(Equal(cloudformation.NewString(("Api"))))
			Expect(f.Events["TestApi"].Properties.ApiEvent).ToNot(BeNil())

			event := f.Events["TestApi"].Properties.ApiEvent
			Expect(event.Method).To(Equal(cloudformation.NewString(("any"))))
			Expect(event.Path).To(Equal(cloudformation.NewString(("/testing"))))
		})

	})

	Context("with an AWS CloudFormation template that contains multiple resources", func() {

		Context("described as Go structs", func() {

			template := cloudformation.NewTemplate()

			template.Resources["MySNSTopic"] = cloudformation.AWSSNSTopic{
				DisplayName: cloudformation.NewString("test-sns-topic-display-name"),
				TopicName:   cloudformation.NewString("test-sns-topic-name"),
				Subscription: []cloudformation.AWSSNSTopic_Subscription{
					cloudformation.AWSSNSTopic_Subscription{
						Endpoint: cloudformation.NewString("test-sns-topic-subscription-endpoint"),
						Protocol: cloudformation.NewString("test-sns-topic-subscription-protocol"),
					},
				},
			}

			template.Resources["MyRoute53HostedZone"] = cloudformation.AWSRoute53HostedZone{
				Name: cloudformation.NewString("example.com"),
			}

			topics := template.GetAllAWSSNSTopicResources()
			It("should have one AWS::SNS::Topic resource", func() {
				Expect(topics).To(HaveLen(1))
				Expect(topics).To(HaveKey("MySNSTopic"))
			})

			topic, err := template.GetAWSSNSTopicWithName("MySNSTopic")
			It("should be able to find the AWS::SNS::Topic by name", func() {
				Expect(topic).ToNot(BeNil())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have the correct AWS::SNS::Topic values", func() {
				Expect(topic.DisplayName).To(Equal(cloudformation.NewString("test-sns-topic-display-name")))
				Expect(topic.TopicName).To(Equal(cloudformation.NewString("test-sns-topic-name")))
				Expect(topic.Subscription).To(HaveLen(1))
				Expect(topic.Subscription[0].Endpoint).To(Equal(cloudformation.NewString("test-sns-topic-subscription-endpoint")))
				Expect(topic.Subscription[0].Protocol).To(Equal(cloudformation.NewString("test-sns-topic-subscription-protocol")))
			})

			zones := template.GetAllAWSRoute53HostedZoneResources()
			It("should have one AWS::Route53::HostedZone resource", func() {
				Expect(zones).To(HaveLen(1))
				Expect(zones).To(HaveKey("MyRoute53HostedZone"))
			})

			zone, err := template.GetAWSRoute53HostedZoneWithName("MyRoute53HostedZone")
			It("should be able to find the AWS::Route53::HostedZone by name", func() {
				Expect(zone).ToNot(BeNil())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have the correct AWS::Route53::HostedZone values", func() {
				Expect(zone.Name).To(Equal(cloudformation.NewString("example.com")))
			})

		})

		Context("described as JSON", func() {

			template := []byte(`{"AWSTemplateFormatVersion":"2010-09-09","Resources":{"MyRoute53HostedZone":{"Type":"AWS::Route53::HostedZone","Properties":{"Name":"example.com"}},"MySNSTopic":{"Type":"AWS::SNS::Topic","Properties":{"DisplayName":"test-sns-topic-display-name","Subscription":[{"Endpoint":"test-sns-topic-subscription-endpoint","Protocol":"test-sns-topic-subscription-protocol"}],"TopicName":"test-sns-topic-name"}}}}`)

			expected := cloudformation.NewTemplate()

			expected.Resources["MySNSTopic"] = cloudformation.AWSSNSTopic{
				DisplayName: cloudformation.NewString("test-sns-topic-display-name"),
				TopicName:   cloudformation.NewString("test-sns-topic-name"),
				Subscription: []cloudformation.AWSSNSTopic_Subscription{
					cloudformation.AWSSNSTopic_Subscription{
						Endpoint: cloudformation.NewString("test-sns-topic-subscription-endpoint"),
						Protocol: cloudformation.NewString("test-sns-topic-subscription-protocol"),
					},
				},
			}

			expected.Resources["MyRoute53HostedZone"] = cloudformation.AWSRoute53HostedZone{
				Name: cloudformation.NewString("example.com"),
			}

			result, err := goformation.ParseJSON(template)
			It("should marshal to Go structs successfully", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			topics := result.GetAllAWSSNSTopicResources()
			It("should have one AWS::SNS::Topic resource", func() {
				Expect(topics).To(HaveLen(1))
				Expect(topics).To(HaveKey("MySNSTopic"))
			})

			topic, err := result.GetAWSSNSTopicWithName("MySNSTopic")
			It("should be able to find the AWS::SNS::Topic by name", func() {
				Expect(topic).ToNot(BeNil())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have the correct AWS::SNS::Topic values", func() {
				Expect(topic.DisplayName).To(Equal(cloudformation.NewString("test-sns-topic-display-name")))
				Expect(topic.TopicName).To(Equal(cloudformation.NewString("test-sns-topic-name")))
				Expect(topic.Subscription).To(HaveLen(1))
				Expect(topic.Subscription[0].Endpoint).To(Equal(cloudformation.NewString("test-sns-topic-subscription-endpoint")))
				Expect(topic.Subscription[0].Protocol).To(Equal(cloudformation.NewString("test-sns-topic-subscription-protocol")))
			})

			zones := result.GetAllAWSRoute53HostedZoneResources()
			It("should have one AWS::Route53::HostedZone resource", func() {
				Expect(zones).To(HaveLen(1))
				Expect(zones).To(HaveKey("MyRoute53HostedZone"))
			})

			zone, err := result.GetAWSRoute53HostedZoneWithName("MyRoute53HostedZone")
			It("should be able to find the AWS::Route53::HostedZone by name", func() {
				Expect(zone).ToNot(BeNil())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have the correct AWS::Route53::HostedZone values", func() {
				Expect(zone.Name).To(Equal(cloudformation.NewString("example.com")))
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
					Expect(err).ToNot(HaveOccurred())
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
					Expect(err).ToNot(HaveOccurred())
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
			Expect(err).ToNot(HaveOccurred())
			Expect(template).ShouldNot(PointTo(BeNil()))
		})

		function, err := template.GetAWSServerlessFunctionWithName("IntrinsicFunctionTest")
		It("should have a function named 'IntrinsicFunctionTest'", func() {
			Expect(function).To(Not(BeNil()))
			Expect(err).ToNot(HaveOccurred())
		})

		It("it should have the correct values", func() {
			Expect(function.Runtime).To(Equal(cloudformation.NewString("4.3")))
			Expect(function.Timeout).To(Equal(cloudformation.NewDouble(10)))
		})

	})

	Context("with a Serverless template containing different CodeUri formats", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-function-string-or-s3-location.yaml")
		It("should successfully parse the template", func() {
			Expect(err).ToNot(HaveOccurred())
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
			Expect(f1.CodeUri.String()).To(Equal("s3://testbucket/testkey.zip"))
		})

		f2 := functions["CodeUriWithS3LocationSpecifiedAsObject"]
		It("should parse a CodeUri property with an S3 location specified as an object", func() {
			x := &cloudformation.AWSServerlessFunction_S3Location{}
			err := f2.CodeUri.Raw().(cloudformation.AnythingMap).Convert(x)
			Expect(err).ToNot(HaveOccurred())
			Expect(x.Key.String()).To(Equal("testkey.zip"))
			Expect(x.Version.Raw()).To(Equal(cloudformation.Double(5)))
		})

		f3 := functions["CodeUriWithString"]
		It("should parse a CodeUri property with a string", func() {
			Expect(f3.CodeUri.String()).To(Equal("./testfolder"))
		})

	})

	Context("with a template defined as Go code", func() {

		template := &cloudformation.Template{
			Resources: map[string]interface{}{
				"MyLambdaFunction": cloudformation.AWSLambdaFunction{
					Handler: cloudformation.NewString("nodejs6.10"),
				},
			},
		}

		functions := template.GetAllAWSLambdaFunctionResources()
		It("should be able to retrieve all Lambda functions with GetAllAWSLambdaFunction(template)", func() {
			Expect(functions).To(HaveLen(1))
		})

		function, err := template.GetAWSLambdaFunctionWithName("MyLambdaFunction")
		It("should be able to retrieve a specific Lambda function with GetAWSLambdaFunctionWithName(template, name)", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(function).To(BeAssignableToTypeOf(cloudformation.AWSLambdaFunction{}))
		})

		It("should have the correct Handler property", func() {
			Expect(function.Handler).To(Equal(cloudformation.NewString("nodejs6.10")))
		})

	})

	Context("with a template that defines an AWS::Serverless::Function", func() {

		Context("that has a CodeUri property set as an S3 Location", func() {

			template := &cloudformation.Template{
				Resources: map[string]interface{}{
					"MySAMFunction": cloudformation.AWSServerlessFunction{
						Handler: cloudformation.NewString("nodejs6.10"),
						CodeUri: cloudformation.NewValue(
							&cloudformation.AWSServerlessFunction_S3Location{
								Bucket:  cloudformation.NewString("test-bucket"),
								Key:     cloudformation.NewString("test-key"),
								Version: cloudformation.NewInteger(100),
							},
						),
					},
				},
			}

			function, err := template.GetAWSServerlessFunctionWithName("MySAMFunction")
			It("should have an AWS::Serverless::Function called MySAMFunction", func() {
				Expect(function).ToNot(BeNil())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have the correct S3 bucket/key/version", func() {
				x := function.CodeUri.Raw().(*cloudformation.AWSServerlessFunction_S3Location)
				Expect(x.Bucket.String()).To(Equal("test-bucket"))
				Expect(x.Key.String()).To(Equal("test-key"))
				Expect(x.Version.Raw()).To(Equal(cloudformation.Integer(100)))
			})

		})

		Context("that has a CodeUri property set as a string", func() {

			codeuri := cloudformation.NewString("./some-folder")
			template := &cloudformation.Template{
				Resources: map[string]interface{}{
					"MySAMFunction": cloudformation.AWSServerlessFunction{
						Handler: cloudformation.NewString("nodejs6.10"),
						CodeUri: codeuri,
					},
				},
			}

			function, err := template.GetAWSServerlessFunctionWithName("MySAMFunction")
			It("should have an AWS::Serverless::Function called MySAMFunction", func() {
				Expect(function).ToNot(BeNil())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have the correct CodeUri", func() {
				Expect(function.CodeUri.String()).To(Equal("./some-folder"))
			})

		})

	})

	Context("with a YAML template that contains AWS::Serverless::SimpleTable resource(s)", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-simpletable.yaml")

		It("should parse the template successfully", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		table, err := template.GetAWSServerlessSimpleTableWithName("TestSimpleTable")
		It("should have a table named 'TestSimpleTable'", func() {
			Expect(table).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should have a primary key set", func() {
			Expect(table.PrimaryKey).ToNot(BeNil())
		})

		It("should have the correct value for the primary key name", func() {
			Expect(table.PrimaryKey.Name).To(Equal(cloudformation.NewString("test-primary-key-name")))
		})

		It("should have the correct value for the primary key type", func() {
			Expect(table.PrimaryKey.Type).To(Equal(cloudformation.NewString("test-primary-key-type")))
		})

		It("should have provisioned throughput set", func() {
			Expect(table.ProvisionedThroughput).ToNot(BeNil())
		})

		It("should have the correct value for ReadCapacityUnits", func() {
			Expect(table.ProvisionedThroughput.ReadCapacityUnits).To(Equal(cloudformation.NewDouble(100)))
		})

		It("should have the correct value for WriteCapacityUnits", func() {
			Expect(table.ProvisionedThroughput.WriteCapacityUnits).To(Equal(cloudformation.NewDouble(200)))
		})

		It("should have a table named 'TestSimpleTableNoProperties'", func() {
			nopropertiesTable, err := template.GetAWSServerlessSimpleTableWithName("TestSimpleTableNoProperties")
			Expect(nopropertiesTable).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

	})

	Context("with a YAML template that contains AWS::Serverless::Api resource(s)", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-api.yaml")

		It("should parse the template successfully", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		api1, err := template.GetAWSServerlessApiWithName("ServerlessApiWithDefinitionUriAsString")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionUriAsString'", func() {
			Expect(api1).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should have the correct value for Name", func() {
			Expect(api1.Name.String()).To(Equal("test-name"))
		})

		It("should have the correct value for StageName", func() {
			Expect(api1.StageName.String()).To(Equal("test-stage-name"))
		})

		It("should have the correct value for DefinitionUri", func() {
			Expect(api1.DefinitionUri.String()).To(Equal("test-definition-uri"))
		})

		It("should have the correct value for CacheClusterEnabled", func() {
			Expect(api1.CacheClusterEnabled).To(Equal(cloudformation.True()))
		})

		It("should have the correct value for CacheClusterSize", func() {
			Expect(api1.CacheClusterSize.String()).To(Equal("test-cache-cluster-size"))
		})

		It("should have the correct value for Variables", func() {
			Expect(api1.Variables).To(HaveKeyWithValue("NAME", cloudformation.NewString("VALUE")))
		})

		api2, err := template.GetAWSServerlessApiWithName("ServerlessApiWithDefinitionUriAsS3Location")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionUriAsS3Location'", func() {
			Expect(api2).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should have the correct value for DefinitionUri", func() {
			Expect(api2.DefinitionUri).ToNot(BeNil())

			x := &cloudformation.AWSServerlessFunction_S3Location{}
			err := api2.DefinitionUri.Raw().(cloudformation.AnythingMap).Convert(x)
			Expect(err).ToNot(HaveOccurred())
			Expect(x.Bucket.String()).To(Equal("test-bucket"))
			Expect(x.Key.String()).To(Equal("test-key"))
			Expect(x.Version.Raw()).To(Equal(cloudformation.Double(1)))
		})

		api3, err := template.GetAWSServerlessApiWithName("ServerlessApiWithDefinitionBodyAsJSON")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionBodyAsJSON'", func() {
			Expect(api3).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should have the correct value for DefinitionBody", func() {
			Expect(api3.DefinitionBody).To(Equal("{\n  \"DefinitionKey\": \"test-definition-value\"\n}\n"))
		})

		api4, err := template.GetAWSServerlessApiWithName("ServerlessApiWithDefinitionBodyAsYAML")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionBodyAsYAML'", func() {
			Expect(api4).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
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
			Expect(err).ToNot(HaveOccurred())
			Expect(template).ShouldNot(BeNil())
		})

		function, err := template.GetAWSServerlessFunctionWithName("IntrinsicEnvironmentVariableTestFunction")
		It("should have a function named 'IntrinsicEnvironmentVariableTestFunction'", func() {
			Expect(function).To(Not(BeNil()))
			Expect(err).ToNot(HaveOccurred())
		})

		It("it should have the correct values", func() {
			Expect(function.Environment.Variables).To(HaveKey("REF_ENV_VAR"))
			Expect(function.Environment.Variables["REF_ENV_VAR"].String()).To(Equal("SomeNewValue"))
		})
	})

	Context("with a SNS event source", func() {
		event := cloudformation.AWSServerlessFunction_Properties{
			SNSEvent: &cloudformation.AWSServerlessFunction_SNSEvent{
				Topic: cloudformation.NewString("MyTopic"),
			},
		}
		It("should marshal properties correctly", func() {
			bytes, err := event.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())
			Expect(string(bytes)).To(Equal(`{"Topic":"MyTopic"}`))
		})
	})

	Context("with an SNS event source created from JSON", func() {
		eventString := `{"Topic":"MyTopic"}`
		eventJson := []byte(eventString)
		event := &cloudformation.AWSServerlessFunction_Properties{}

		It("should marshal properties correctly", func() {
			err := event.UnmarshalJSON(eventJson)
			Expect(err).ToNot(HaveOccurred())
			bytes, err := event.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())
			Expect(string(bytes)).To(Equal(`{"Topic":"MyTopic"}`))
		})
	})

})
