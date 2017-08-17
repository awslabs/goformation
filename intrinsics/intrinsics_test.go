package intrinsics_test

import (
	"encoding/json"

	. "github.com/awslabs/goformation/intrinsics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AWS CloudFormation intrinsic function processing", func() {

	Context("with a template that contains invalid JSON", func() {
		const template = `{`
		processed, err := Process([]byte(template), nil)
		It("should fail to process the template", func() {
			Expect(processed).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	// pmaddox@ 2017-08-17:
	// Commented out until we have support for YAML tag intrinsic functions (e.g. !Sub)
	//Context("with a YAML template that contains intrinsic functions in tag form", func() {

	// t := "AWSTemplateFormatVersion: '2010-09-09'\n"
	// t += "Transform: AWS::Serverless-2016-10-31\n"
	// t += "Description: SAM template for testing intrinsic functions with YAML tags\n"
	// t += "Resources:\n"
	// t += "  CodeUriWithS3LocationSpecifiedAsString:\n"
	// t += "    Type: AWS::Serverless::Function\n"
	// t += "    Properties:\n"
	// t += "      Runtime: !Sub test-${runtime}\n"
	// t += "      Timeout: !Ref ThisWontResolve\n"

	// data, err := yaml.YAMLToJSON([]byte(t))
	// It("should successfully convert YAML to JSON", func() {
	// 	Expect(data).ShouldNot(BeNil())
	// 	Expect(err).Should(BeNil())
	// })

	// processed, err := Process(data, nil)
	// It("should successfully process the template", func() {
	// 	Expect(processed).ShouldNot(BeNil())
	// 	Expect(err).Should(BeNil())
	// })

	//}

	Context("with a template that contains primitives, intrinsics, and nested intrinsics", func() {

		const template = `{
			"Resources": {
				"ExampleResource": {
					"Type": "AWS::Example::Resource",
					"Properties": {
						"StringProperty": "Simple string example",						
						"BooleanProperty": true,
						"NumberProperty": 123.45,
						"JoinIntrinsicProperty": { "Fn::Join": [ "some", "name" ] },				
						"JoinNestedIntrinsicProperty": { "Fn::Join": [ "some", { "Fn::Join": [ "joined", "value" ] } ] },
						"SubIntrinsicProperty": { "Fn::Sub": [ "some ${value}", { "value": "value" } ] }			
					}
				}
			}
		}`

		Context("with no processor options", func() {

			processed, err := Process([]byte(template), nil)
			It("should successfully process the template", func() {
				Expect(processed).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})

			var result interface{}
			err = json.Unmarshal(processed, &result)

			It("should be valid JSON, and marshal to a Go type", func() {
				Expect(processed).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			template := result.(map[string]interface{})
			resources := template["Resources"].(map[string]interface{})
			resource := resources["ExampleResource"].(map[string]interface{})
			properties := resource["Properties"].(map[string]interface{})

			It("should have the correct value for a primitive string property", func() {
				Expect(properties["StringProperty"]).To(Equal("Simple string example"))
			})

			It("should have the correct value for a primitive boolean property", func() {
				Expect(properties["BooleanProperty"]).To(Equal(true))
			})

			It("should have the correct value for a primitive number property", func() {
				Expect(properties["NumberProperty"]).To(Equal(123.45))
			})

			It("should have the correct value for a Fn::Join intrinsic property", func() {
				Expect(properties["JoinIntrinsicProperty"]).To(BeNil())
			})

			It("should have the correct value for a nested Fn::Join intrinsic property", func() {
				Expect(properties["JoinNestedIntrinsicProperty"]).To(BeNil())
			})

			It("should have the correct value for a Fn::Sub intrinsic property", func() {
				Expect(properties["SubIntrinsicProperty"]).To(BeNil())
			})

		})

		Context("with a processor options override for the Fn::Join function", func() {

			opts := &ProcessorOptions{
				IntrinsicHandlerOverrides: map[string]IntrinsicHandler{
					"Fn::Join": func(name string, input interface{}) interface{} {
						return "overridden"
					},
				},
			}

			processed, err := Process([]byte(template), opts)
			It("should successfully process the template", func() {
				Expect(processed).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})

			result := map[string]interface{}{}
			err = json.Unmarshal(processed, &result)

			It("should be valid JSON, and marshal to a Go type", func() {
				Expect(processed).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			resources := result["Resources"].(map[string]interface{})
			resource := resources["ExampleResource"].(map[string]interface{})
			properties := resource["Properties"].(map[string]interface{})

			It("should have the correct value for a primitive string property", func() {
				Expect(properties["StringProperty"]).To(Equal("Simple string example"))
			})

			It("should have the correct value for a primitive boolean property", func() {
				Expect(properties["BooleanProperty"]).To(Equal(true))
			})

			It("should have the correct value for a primitive number property", func() {
				Expect(properties["NumberProperty"]).To(Equal(123.45))
			})

			It("should have the correct value for an intrinsic property", func() {
				Expect(properties["JoinIntrinsicProperty"]).To(Equal("overridden"))
			})

			It("should have the correct value for a nested intrinsic property", func() {
				Expect(properties["JoinNestedIntrinsicProperty"]).To(Equal("overridden"))
			})

			It("should have the correct value for an intrinsic property that's not supposed to be overridden", func() {
				Expect(properties["SubIntrinsicProperty"]).To(BeNil())
			})

		})

	})
})
