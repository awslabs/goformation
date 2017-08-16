package goformation_test

import (
	"encoding/json"

	"github.com/paulmaddox/goformation"
	"github.com/paulmaddox/goformation/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goformation", func() {

	Context("with a template defined as Go code", func() {

		t1 := &resources.Template{
			Resources: map[string]interface{}{
				"MyLambdaFunction": resources.AWSLambdaFunction{
					Handler: "nodejs6.10",
				},
			},
		}

		data, err := json.Marshal(t1)
		It("should successfully marshal to JSON", func() {
			Expect(err).To(BeNil())
		})

		template, err := goformation.Parse(data)
		It("should successfully parse the JSON template", func() {
			Expect(template).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		functions := template.GetAllAWSLambdaFunctionResources()
		It("should be able to retrieve all Lambda functions with GetAllAWSLambdaFunction(template)", func() {
			Expect(functions).To(HaveLen(1))
		})

		function, err := template.GetAWSLambdaFunctionWithName("MyLambdaFunction")
		It("should be able to retrieve a specific Lambda function with GetAWSLambdaFunctionWithName(template, name)", func() {
			Expect(err).To(BeNil())
			Expect(function).To(BeAssignableToTypeOf(&resources.AWSLambdaFunction{}))
		})

		It("should have the correct Handler property", func() {
			Expect(function.Handler).To(Equal("nodejs6.10"))
		})

	})

	Context("with a template that defines an AWS::Serverless::Function", func() {

		Context("that has a CodeUri property set as an S3 Location", func() {

			t1 := &resources.Template{
				Resources: map[string]interface{}{
					"MySAMFunction": resources.AWSServerlessFunction{
						Handler: "nodejs6.10",
						CodeUri: resources.AWSServerlessFunction_StringOrS3Location{
							S3Location: resources.AWSServerlessFunction_S3Location{
								Bucket:  "test-bucket",
								Key:     "test-key",
								Version: 100,
							},
						},
					},
				},
			}

			data, err := json.Marshal(t1)
			It("should successfully marshal to JSON", func() {
				Expect(err).To(BeNil())
			})

			template, err := goformation.Parse(data)
			It("should successfully parse the JSON template", func() {
				Expect(template).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

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

			t1 := &resources.Template{
				Resources: map[string]interface{}{
					"MySAMFunction": resources.AWSServerlessFunction{
						Handler: "nodejs6.10",
						CodeUri: resources.AWSServerlessFunction_StringOrS3Location{
							String: "./some-folder",
						},
					},
				},
			}

			data, err := json.Marshal(t1)
			It("should successfully marshal to JSON", func() {
				Expect(err).To(BeNil())
			})

			template, err := goformation.Parse(data)
			It("should successfully parse the JSON template", func() {
				Expect(template).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			function, err := template.GetAWSServerlessFunctionWithName("MySAMFunction")
			It("should have an AWS::Serverless::Function called MySAMFunction", func() {
				Expect(function).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have the correct CodeUri", func() {
				Expect(function.CodeUri.String).To(Equal("./some-folder"))
			})

		})

	})

})
