package goformation_test

import (
	"encoding/json"
	"fmt"

	"github.com/sanathkr/yaml"

	"github.com/awslabs/goformation/v5"
	"github.com/awslabs/goformation/v5/cloudformation"
	"github.com/awslabs/goformation/v5/cloudformation/ec2"
	"github.com/awslabs/goformation/v5/cloudformation/global"
	"github.com/awslabs/goformation/v5/cloudformation/lambda"
	"github.com/awslabs/goformation/v5/cloudformation/policies"
	"github.com/awslabs/goformation/v5/cloudformation/route53"
	"github.com/awslabs/goformation/v5/cloudformation/s3"
	"github.com/awslabs/goformation/v5/cloudformation/serverless"
	"github.com/awslabs/goformation/v5/cloudformation/sns"
	"github.com/awslabs/goformation/v5/intrinsics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

func Example_to_json() {

	// Create a new CloudFormation template
	template := cloudformation.NewTemplate()

	// Create an Amazon SNS topic, with a unique name based off the current timestamp
	template.Resources["MyTopic"] = &sns.Topic{
		TopicName: "my-topic-1575143839",
	}

	// Create a subscription, connected to our topic, that forwards notifications to an email address
	template.Resources["MyTopicSubscription"] = &sns.Subscription{
		TopicArn: cloudformation.Ref("MyTopic"),
		Protocol: "email",
		Endpoint: "some.email@example.com",
	}

	// Let's see the JSON AWS CloudFormation template
	j, err := template.JSON()
	if err != nil {
		fmt.Printf("Failed to generate JSON: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(j))
	}

	// Output:
	// {
	//   "AWSTemplateFormatVersion": "2010-09-09",
	//   "Resources": {
	//     "MyTopic": {
	//       "Properties": {
	//         "TopicName": "my-topic-1575143839"
	//       },
	//       "Type": "AWS::SNS::Topic"
	//     },
	//     "MyTopicSubscription": {
	//       "Properties": {
	//         "Endpoint": "some.email@example.com",
	//         "Protocol": "email",
	//         "TopicArn": {
	//           "Ref": "MyTopic"
	//         }
	//       },
	//       "Type": "AWS::SNS::Subscription"
	//     }
	//   }
	// }
}

func Example_to_yaml() {

	// Create a new CloudFormation template
	template := cloudformation.NewTemplate()

	// Create an Amazon SNS topic, with a unique name based off the current timestamp
	template.Resources["MyTopic"] = &sns.Topic{
		TopicName: "my-topic-1575143970",
	}

	// Create a subscription, connected to our topic, that forwards notifications to an email address
	template.Resources["MyTopicSubscription"] = &sns.Subscription{
		TopicArn: cloudformation.Ref("MyTopic"),
		Protocol: "email",
		Endpoint: "some.email@example.com",
	}

	// Let's see the YAML AWS CloudFormation template
	y, err := template.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(y))
	}

	// Output:
	// AWSTemplateFormatVersion: 2010-09-09
	// Resources:
	//   MyTopic:
	//     Properties:
	//       TopicName: my-topic-1575143970
	//     Type: AWS::SNS::Topic
	//   MyTopicSubscription:
	//     Properties:
	//       Endpoint: some.email@example.com
	//       Protocol: email
	//       TopicArn:
	//         Ref: MyTopic
	//     Type: AWS::SNS::Subscription

}

func Example_to_go() {

	// Open a template from file (can be JSON or YAML)
	template, err := goformation.Open("example/yaml-to-go/template.yaml")
	if err != nil {
		fmt.Printf("There was an error processing the template: %s", err)
		return
	}

	// You can extract all resources of a certain type
	// Each AWS CloudFormation resource is a strongly typed struct
	topics := template.GetAllSNSTopicResources()
	for name, topic := range topics {

		// E.g. Found a AWS::SNS::Topic with Logical ID ExampleTopic and TopicName 'example'
		fmt.Printf("Found a %s with Logical ID %s and TopicName %s\n", topic.AWSCloudFormationType(), name, topic.TopicName)

	}

	// You can also search for specific resources by their logicalId
	search := "ExampleTopic"
	topic, err := template.GetSNSTopicWithName(search)
	if err != nil {
		fmt.Printf("SNS topic with logical ID %s not found", search)
		return
	}

	// E.g. Found a AWS::Serverless::Function named GetHelloWorld (runtime: nodejs6.10)
	fmt.Printf("Found a %s with Logical ID %s and TopicName %s\n", topic.AWSCloudFormationType(), search, topic.TopicName)

	// Output:
	// Found a AWS::SNS::Topic with Logical ID ExampleTopic and TopicName example
	// Found a AWS::SNS::Topic with Logical ID ExampleTopic and TopicName example

}

var _ = Describe("Goformation", func() {

	Context("with a Serverless function matching 2016-10-31 specification", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-function-2016-10-31.yaml")
		It("should successfully validate the SAM template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		functions := template.GetAllServerlessFunctionResources()

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
			Expect((*f.Policies.SAMPolicyTemplateArray)[0].DynamoDBCrudPolicy.TableName).To(Equal("table_arn"))
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

		It("should correctly parse all of the function S3 event source", func() {
			Expect(f.Events).ToNot(BeNil())
			Expect(f.Events).To(HaveKey("TestS3"))
			Expect(f.Events["TestS3"].Type).To(Equal("S3"))
			Expect(f.Events["TestS3"].Properties.S3Event).ToNot(BeNil())

			event := f.Events["TestS3"].Properties.S3Event
			Expect(event.Bucket).To(Equal("my-photo-bucket"))
			Expect(event.Events.String).To(PointTo(Equal("s3:ObjectCreated:*")))
			Expect(event.Filter.S3Key.Rules).To(HaveLen(1))
			Expect(event.Filter.S3Key.Rules[0].Name).To(Equal("prefix|suffix"))
			Expect(event.Filter.S3Key.Rules[0].Value).To(Equal("my-prefix|my-suffix"))

		})

	})

	Context("with a JSON template that contains a resource with tags", func() {

		template, err := goformation.Open("test/json/resource-with-tags.json")
		It("should successfully validate the template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		resources := template.GetAllAutoScalingAutoScalingGroupResources()
		It("should have exactly one resource", func() {
			Expect(resources).To(HaveLen(1))
			Expect(resources).To(HaveKey("EcsClusterDefaultAutoScalingGroupASGC1A785DB"))
		})

		asg := resources["EcsClusterDefaultAutoScalingGroupASGC1A785DB"]
		It("should have exactly one tag defined", func() {
			Expect(asg.Tags).To(HaveLen(1))
		})

		It("should have the correct tag properties set", func() {
			Expect(asg.Tags[0].PropagateAtLaunch).To(Equal(true))
			Expect(asg.Tags[0].Key).To(Equal("Name"))
			Expect(asg.Tags[0].Value).To(Equal("aws-ecs-integ-ecs/EcsCluster/DefaultAutoScalingGroup"))
		})

	})

	Context("with a Custom Resource template", func() {

		template, err := goformation.Open("test/yaml/custom-resource.yaml")
		It("should successfully validate the template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		It("should correctly Marshal the custom resource", func() {
			data, err := template.JSON()
			Expect(err).To(BeNil())

			var result map[string]interface{}
			if err := json.Unmarshal(data, &result); err != nil {
				Fail(err.Error())
			}

			resources, ok := result["Resources"].(map[string]interface{})
			Expect(ok).To(BeTrue())
			Expect(resources).To(HaveLen(1))
			Expect(resources).To(HaveKey("MyCustomResource"))

			mcr := resources["MyCustomResource"].(map[string]interface{})
			Expect(mcr["Properties"]).To(HaveKey("CustomProperty"))
		})
	})

	Context("with an AWS CloudFormation template that contains multiple resources", func() {

		Context("described as Go structs", func() {

			template := cloudformation.NewTemplate()

			template.Resources["MySNSTopic"] = &sns.Topic{
				DisplayName: "test-sns-topic-display-name",
				TopicName:   "test-sns-topic-name",
				Subscription: []sns.Topic_Subscription{
					sns.Topic_Subscription{
						Endpoint: "test-sns-topic-subscription-endpoint",
						Protocol: "test-sns-topic-subscription-protocol",
					},
				},
			}

			template.Resources["MyRoute53HostedZone"] = &route53.HostedZone{
				Name: "example.com",
			}

			topics := template.GetAllSNSTopicResources()
			It("should have one AWS::SNS::Topic resource", func() {
				Expect(topics).To(HaveLen(1))
				Expect(topics).To(HaveKey("MySNSTopic"))
			})

			topic, err := template.GetSNSTopicWithName("MySNSTopic")
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

			zones := template.GetAllRoute53HostedZoneResources()
			It("should have one AWS::Route53::HostedZone resource", func() {
				Expect(zones).To(HaveLen(1))
				Expect(zones).To(HaveKey("MyRoute53HostedZone"))
			})

			zone, err := template.GetRoute53HostedZoneWithName("MyRoute53HostedZone")
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

			expected.Resources["MySNSTopic"] = &sns.Topic{
				DisplayName: "test-sns-topic-display-name",
				TopicName:   "test-sns-topic-name",
				Subscription: []sns.Topic_Subscription{
					sns.Topic_Subscription{
						Endpoint: "test-sns-topic-subscription-endpoint",
						Protocol: "test-sns-topic-subscription-protocol",
					},
				},
			}

			expected.Resources["MyRoute53HostedZone"] = &route53.HostedZone{
				Name: "example.com",
			}

			result, err := goformation.ParseJSON(template)
			It("should marshal to Go structs successfully", func() {
				Expect(err).To(BeNil())
			})

			topics := result.GetAllSNSTopicResources()
			It("should have one AWS::SNS::Topic resource", func() {
				Expect(topics).To(HaveLen(1))
				Expect(topics).To(HaveKey("MySNSTopic"))
			})

			topic, err := result.GetSNSTopicWithName("MySNSTopic")
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

			zones := result.GetAllRoute53HostedZoneResources()
			It("should have one AWS::Route53::HostedZone resource", func() {
				Expect(zones).To(HaveLen(1))
				Expect(zones).To(HaveKey("MyRoute53HostedZone"))
			})

			zone, err := result.GetRoute53HostedZoneWithName("MyRoute53HostedZone")
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

		function, err := template.GetServerlessFunctionWithName("IntrinsicFunctionTest")
		It("should have a function named 'IntrinsicFunctionTest'", func() {
			Expect(function).To(Not(BeNil()))
			Expect(err).To(BeNil())
		})

		It("it should have the correct values", func() {
			Expect(function.Runtime).To(Equal("4.3"))
			Expect(function.Timeout).To(Equal(10))
		})

	})

	Context("with a Serverless template containing different CORS configuration formats", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-api-string-or-cors-configuration.yaml")
		It("should successfully parse the template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		apis := template.GetAllServerlessApiResources()

		It("should have exactly two APIs", func() {
			Expect(apis).To(HaveLen(2))
			Expect(apis).To(HaveKey("RestApiWithCorsConfiguration"))
			Expect(apis).To(HaveKey("RestApiWithCorsString"))
		})

		api1 := apis["RestApiWithCorsConfiguration"]
		It("should parse a Cors configuration object", func() {
			Expect(api1.Cors.CorsConfiguration.AllowHeaders).To(Equal("'Authorization,authorization'"))
			Expect(api1.Cors.CorsConfiguration.AllowOrigin).To(Equal("'*'"))
		})

		api2 := apis["RestApiWithCorsString"]
		It("should parse a Cors string", func() {
			Expect(api2.Cors.String).To(PointTo(Equal("'www.example.com'")))
		})

	})

	Context("with a Serverless template containing different CodeUri formats", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-function-string-or-s3-location.yaml")
		It("should successfully parse the template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		functions := template.GetAllServerlessFunctionResources()

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
			Resources: cloudformation.Resources{
				"MyLambdaFunction": &lambda.Function{
					Handler: "nodejs6.10",
				},
			},
		}

		functions := template.GetAllLambdaFunctionResources()
		It("should be able to retrieve all Lambda functions with GetAllLambdaFunction(template)", func() {
			Expect(functions).To(HaveLen(1))
		})

		function, err := template.GetLambdaFunctionWithName("MyLambdaFunction")
		It("should be able to retrieve a specific Lambda function with GetLambdaFunctionWithName(template, name)", func() {
			Expect(err).To(BeNil())
			Expect(function).To(BeAssignableToTypeOf(&lambda.Function{}))
		})

		It("should have the correct Handler property", func() {
			Expect(function.Handler).To(Equal("nodejs6.10"))
		})

	})

	Context("with a template that defines an AWS::Serverless::Function", func() {

		Context("that has a CodeUri property set as an S3 Location", func() {

			template := &cloudformation.Template{
				Resources: cloudformation.Resources{
					"MySAMFunction": &serverless.Function{
						Handler: "nodejs6.10",
						CodeUri: &serverless.Function_CodeUri{
							S3Location: &serverless.Function_S3Location{
								Bucket:  "test-bucket",
								Key:     "test-key",
								Version: 100,
							},
						},
					},
				},
			}

			function, err := template.GetServerlessFunctionWithName("MySAMFunction")
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
				Resources: cloudformation.Resources{
					"MySAMFunction": &serverless.Function{
						Handler: "nodejs6.10",
						CodeUri: &serverless.Function_CodeUri{
							String: &codeuri,
						},
					},
				},
			}

			function, err := template.GetServerlessFunctionWithName("MySAMFunction")
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

		table, err := template.GetServerlessSimpleTableWithName("TestSimpleTable")
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
			nopropertiesTable, err := template.GetServerlessSimpleTableWithName("TestSimpleTableNoProperties")
			Expect(nopropertiesTable).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct DeletionPolicy", func() {
			Expect(table.AWSCloudFormationDeletionPolicy).To(Equal(policies.DeletionPolicy("Retain")))
		})
	})

	Context("with a YAML template that contains AWS::Serverless::Api resource(s)", func() {

		template, err := goformation.Open("test/yaml/aws-serverless-api.yaml")

		It("should parse the template successfully", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		api1, err := template.GetServerlessApiWithName("ServerlessApiWithDefinitionUriAsString")
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

		api2, err := template.GetServerlessApiWithName("ServerlessApiWithDefinitionUriAsS3Location")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionUriAsS3Location'", func() {
			Expect(api2).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct value for DefinitionUri", func() {
			Expect(api2.DefinitionUri.S3Location.Bucket).To(Equal("test-bucket"))
			Expect(api2.DefinitionUri.S3Location.Key).To(Equal("test-key"))
			Expect(api2.DefinitionUri.S3Location.Version).To(Equal(1))
		})

		api3, err := template.GetServerlessApiWithName("ServerlessApiWithDefinitionBodyAsJSON")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithDefinitionBodyAsJSON'", func() {
			Expect(api3).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct value for DefinitionBody", func() {
			Expect(api3.DefinitionBody).To(Equal("{\n  \"DefinitionKey\": \"test-definition-value\"\n}\n"))
		})

		api4, err := template.GetServerlessApiWithName("ServerlessApiWithDefinitionBodyAsYAML")
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

		api5, err := template.GetServerlessApiWithName("ServerlessApiWithAccessLogSettingAsYAML")
		It("should have an AWS::Serverless::Api named 'ServerlessApiWithAccessLogSettingAsYAML'", func() {
			Expect(api5).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct value for AccessLogSetting", func() {
			Expect(api5.AccessLogSetting.DestinationArn).To(Equal("arn:test"))
			Expect(api5.AccessLogSetting.Format).To(Equal("{customKey: $context.Key}"))
		})

	})

	Context("with a default SAM CLI-created YAML template", func() {
		template, err := goformation.Open("test/yaml/aws-serverless-sam-cli-default.yaml")

		It("should parse the template successfully", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		global, err := template.GetServerlessGlobalFunction()
		It("should have a Global Function definition containing a timeout", func() {
			Expect(global.Timeout).To(Equal(5))
			Expect(err).To(BeNil())
		})

		function, err := template.GetServerlessFunctionWithName("HelloWorldFunction")
		It("should have an AWS::Serverless::Function called HelloWorldFunction", func() {
			Expect(function).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should have the correct CodeUri", func() {
			Expect(function.CodeUri.String).To(PointTo(Equal("hello-world/")))
		})
	})

	Context("with a YAML template with single transform macro", func() {
		template, err := goformation.Open("test/yaml/transform-single.yaml")

		It("should parse the template successfully", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should parse transform macro into String field", func() {
			Expect(*template.Transform.String).To(Equal("MyTranformMacro"))
		})

		It("should StringArray remain nil", func() {
			Expect(template.Transform.StringArray).To(BeNil())
		})

	})

	Context("with a YAML template with multiple transform macros", func() {
		template, err := goformation.Open("test/yaml/transform-multiple.yaml")

		It("should parse the template successfully", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("should parse transform macro into StringArray field", func() {
			Expect(*template.Transform.StringArray).To(Equal([]string{"FirstMacro", "SecondMacro"}))
		})

		It("should String remain nil", func() {
			Expect(template.Transform.String).To(BeNil())
		})

	})

	Context("with a YAML template with parameter overrides", func() {

		template, err := goformation.OpenWithOptions("test/yaml/aws-serverless-function-env-vars.yaml", &intrinsics.ProcessorOptions{
			ParameterOverrides: map[string]interface{}{"ExampleParameter": "SomeNewValue"},
		})

		It("should successfully validate the SAM template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		function, err := template.GetServerlessFunctionWithName("IntrinsicEnvironmentVariableTestFunction")
		It("should have a function named 'IntrinsicEnvironmentVariableTestFunction'", func() {
			Expect(function).To(Not(BeNil()))
			Expect(err).To(BeNil())
		})

		It("it should have the correct values", func() {
			Expect(function.Environment.Variables).To(HaveKeyWithValue("REF_ENV_VAR", "SomeNewValue"))
		})
	})

	Context("with a SNS event source", func() {
		event := serverless.Function_Properties{
			SNSEvent: &serverless.Function_SNSEvent{
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
		event := serverless.Function_Properties{}
		event.UnmarshalJSON(eventJson)

		It("should marshal properties correctly", func() {
			bytes, err := event.MarshalJSON()
			Expect(err).To(BeNil())
			Expect(string(bytes)).To(Equal(eventString))
		})
	})

	Context("with an API event source", func() {
		event := serverless.Function_Properties{
			ApiEvent: &serverless.Function_ApiEvent{
				Auth: &serverless.Function_Auth{
					ApiKeyRequired:      true,
					AuthorizationScopes: []string{"scope1", "scope2"},
					Authorizer:          "aws_iam",
					ResourcePolicy: &serverless.Function_AuthResourcePolicy{
						CustomStatements: []interface{}{
							map[string]interface{}{
								"Effect":   "Allow",
								"Action":   "execute-api:*",
								"Resource": "*",
							},
						},
						AwsAccountBlacklist:    []string{"AwsAccountBlacklistValue"},
						AwsAccountWhitelist:    []string{"AwsAccountWhitelistValue"},
						IntrinsicVpcBlacklist:  []string{"IntrinsicVpcBlacklistValue"},
						IntrinsicVpcWhitelist:  []string{"IntrinsicVpcWhitelistValue"},
						IntrinsicVpceBlacklist: []string{"IntrinsicVpceBlacklistValue"},
						IntrinsicVpceWhitelist: []string{"IntrinsicVpceWhitelistValue"},
						IpRangeBlacklist:       []string{"IpRangeBlacklistValue"},
						IpRangeWhitelist:       []string{"IpRangeWhitelistValue"},
						SourceVpcBlacklist:     []string{"SourceVpcBlacklistValue"},
						SourceVpcWhitelist:     []string{"SourceVpcWhitelistValue"},
					},
				},
				Method:    "MethodValue",
				Path:      "PathValue",
				RestApiId: "RestApiIdValue",
			},
		}

		It("should marshal properties correctly", func() {
			expectedString := `{"Auth":{"ApiKeyRequired":true,"AuthorizationScopes":["scope1","scope2"],"Authorizer":"aws_iam","ResourcePolicy":{"AwsAccountBlacklist":["AwsAccountBlacklistValue"],"AwsAccountWhitelist":["AwsAccountWhitelistValue"],"CustomStatements":[{"Action":"execute-api:*","Effect":"Allow","Resource":"*"}],"IntrinsicVpcBlacklist":["IntrinsicVpcBlacklistValue"],"IntrinsicVpcWhitelist":["IntrinsicVpcWhitelistValue"],"IntrinsicVpceBlacklist":["IntrinsicVpceBlacklistValue"],"IntrinsicVpceWhitelist":["IntrinsicVpceWhitelistValue"],"IpRangeBlacklist":["IpRangeBlacklistValue"],"IpRangeWhitelist":["IpRangeWhitelistValue"],"SourceVpcBlacklist":["SourceVpcBlacklistValue"],"SourceVpcWhitelist":["SourceVpcWhitelistValue"]}},"Method":"MethodValue","Path":"PathValue","RestApiId":"RestApiIdValue"}`
			bytes, err := event.MarshalJSON()
			Expect(err).To(BeNil())
			Expect(string(bytes)).To(Equal(expectedString))
		})
	})

	Context("with a template that contains a reference to another resource within the template", func() {

		template := &cloudformation.Template{
			Resources: cloudformation.Resources{
				"TestBucket": &s3.Bucket{
					BucketName: "test-bucket",
				},
				"TestBucketPolicy": &s3.BucketPolicy{
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
				Name:  "Fn::SubVars",
				Input: cloudformation.SubVars("test-sub", map[string]interface{}{"foo": "bar"}),
				Expected: map[string]interface{}{
					"Fn::Sub": []interface{}{
						"test-sub",
						map[string]interface{}{
							"foo": "bar",
						},
					},
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
					Description: test.Input,
				}

				data, _ := template.JSON()
				var result map[string]interface{}
				json.Unmarshal(data, &result)

				desc, ok := result["Description"].(map[string]interface{})
				Expect(ok).To(BeTrue())
				Expect(desc).To(HaveLen(1))
				Expect(desc).To(BeEquivalentTo(test.Expected))

			})
		}

	})

	Context("with a template that contains nested intrinsics", func() {

		template := &cloudformation.Template{
			Resources: cloudformation.Resources{
				"TestBucket": &s3.Bucket{
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
			Resources: cloudformation.Resources{
				"TestBucket": &s3.Bucket{
					BucketName: "test-bucket",
				},
				"TestBucketPolicy": &s3.BucketPolicy{
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

	Context("with a template that contains conditional resources", func() {

		template := &cloudformation.Template{
			Conditions: map[string]interface{}{
				"MyCondition": cloudformation.Equals(cloudformation.Ref("MyParam"), "test"),
			},
			Resources: cloudformation.Resources{
				"MySNSTopic": &sns.Topic{
					TopicName:                  "test-sns-topic-name",
					AWSCloudFormationCondition: "MyCondition",
				},
			},
		}

		It("should correctly keep conditional intrinsics", func() {
			data, err := template.JSON()
			Expect(err).To(BeNil())

			var result map[string]interface{}
			if err := json.Unmarshal(data, &result); err != nil {
				Fail(err.Error())
			}

			conditions, ok := result["Conditions"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			condition, ok := conditions["MyCondition"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			equal, ok := condition["Fn::Equals"].([]interface{})
			Expect(ok).To(BeTrue())
			Expect(equal).To(HaveLen(2))

			refMap, ok := equal[0].(map[string]interface{})
			Expect(ok).To(BeTrue())

			ref, ok := refMap["Ref"].(interface{})
			Expect(ok).To(BeTrue())
			Expect(ref).To(Equal("MyParam"))

			Expect(equal[1]).To(Equal("test"))

			resources, ok := result["Resources"].(map[string]interface{})
			Expect(ok).To(BeTrue())

			topic, ok := resources["MySNSTopic"].(map[string]interface{})
			Expect(ok).To(BeTrue())
			Expect(topic).Should(HaveLen(3))

			resCondition, ok := topic["Condition"].(interface{})
			Expect(ok).To(BeTrue())
			Expect(resCondition).To(Equal("MyCondition"))

		})

	})

	Context("with a template that contains outputs with no exports", func() {

		Context("described as Go structs", func() {

			template := cloudformation.NewTemplate()

			template.Resources["Vpc"] = &ec2.VPC{
				CidrBlock: "192.168.0.0/20",
			}

			template.Outputs["VpcId"] = cloudformation.Output{
				Value: cloudformation.Ref("Vpc"),
			}

			expected := `{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Outputs": {
    "VpcId": {
      "Value": {
        "Ref": "Vpc"
      }
    }
  },
  "Resources": {
    "Vpc": {
      "Properties": {
        "CidrBlock": "192.168.0.0/20"
      },
      "Type": "AWS::EC2::VPC"
    }
  }
}`

			got, err := template.JSON()
			It("should marshal template successfully", func() {
				Expect(err).To(BeNil())
			})

			It("should be equal to expected output", func() {
				Expect(string(got)).To(Equal(expected))
			})
		})
	})

	Context("with a SAM template", func() {

		Context("described as Go structs", func() {
			template := cloudformation.NewTemplate()
			transform := "AWS::Serverless-2016-10-31"
			transformStruct := &cloudformation.Transform{
				String: &transform,
			}
			template.Transform = transformStruct

			codeUri := "hello-world/"
			template.Resources["HelloWorldFunction"] = &serverless.Function{
				CodeUri: &serverless.Function_CodeUri{
					String: &codeUri,
				},
				Handler: "hello-world",
				Runtime: "go1.x",
			}

			globals := cloudformation.Globals{}
			globals["Function"] = &global.Function{
				Timeout: 123,
			}
			template.Globals = globals

			expected := `{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Globals": {
    "Function": {
      "Timeout": 123
    }
  },
  "Resources": {
    "HelloWorldFunction": {
      "Properties": {
        "CodeUri": "hello-world/",
        "Handler": "hello-world",
        "Runtime": "go1.x"
      },
      "Type": "AWS::Serverless::Function"
    }
  },
  "Transform": "AWS::Serverless-2016-10-31"
}`

			got, err := template.JSON()
			It("should marshal template successfully", func() {
				Expect(err).To(BeNil())
			})

			It("should be equal to expected output", func() {
				Expect(string(got)).To(Equal(expected))
			})
		})

		Context("that has an image in arm64 architecture", func() {
			template := cloudformation.NewTemplate()
			transform := "AWS::Serverless-2016-10-31"
			template.Transform = &cloudformation.Transform{
				String: &transform,
			}
			template.Resources["TestFunction"] = &serverless.Function{
				Architectures: []string{"arm64"},
				ImageUri:      "image:latest-arm64",
			}

			expected := `{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "TestFunction": {
      "Properties": {
        "Architectures": [
          "arm64"
        ],
        "ImageUri": "image:latest-arm64"
      },
      "Type": "AWS::Serverless::Function"
    }
  },
  "Transform": "AWS::Serverless-2016-10-31"
}`

			got, err := template.JSON()
			It("should marshal template successfully", func() {
				Expect(err).To(BeNil())
			})

			It("should be equal to expected output", func() {
				Expect(string(got)).To(Equal(expected))
			})
		})
	})
})
