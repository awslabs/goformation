package main_test

import (
	"encoding/json"

	"github.com/awslabs/goformation/v4"
	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/global"
	"github.com/awslabs/goformation/v4/cloudformation/serverless"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SAM Globals", func() {

	Context("with a resource that contains a mix of primitive, custom and polymorphic properties", func() {

		Context("described as Go structs", func() {

			Context("with a simple primitive used for a polymorphic property", func() {

				codeuri := "s3://bucket/key"
				resource := &global.Function{
					Runtime: "nodejs6.10",
					CodeUri: &serverless.Function_CodeUri{
						String: &codeuri,
					},
				}

				expected := []byte(`{"CodeUri":"s3://bucket/key","Runtime":"nodejs6.10"}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a custom type used for a polymorphic property", func() {

				resource := &global.Function{
					Runtime: "nodejs6.10",
					CodeUri: &serverless.Function_CodeUri{
						S3Location: &serverless.Function_S3Location{
							Bucket:  "test-bucket",
							Key:     "test-key",
							Version: 123,
						},
					},
				}

				expected := []byte(`{"CodeUri":{"Bucket":"test-bucket","Key":"test-key","Version":123},"Runtime":"nodejs6.10"}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})
			})

		})

	})

	Context("with a Global resource", func() {

		Context("specified as JSON", func() {

			Context("with valid fields", func() {

				property := []byte(`{"CodeUri":"s3://bucket/key","Runtime":"nodejs6.10"}`)
				codeuri := "s3://bucket/key"
				expected := &global.Function{
					Runtime: "nodejs6.10",
					CodeUri: &serverless.Function_CodeUri{
						String: &codeuri,
					},
				}

				result := &global.Function{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})
			})

			Context("with an excluded field", func() {

				property := []byte(`{"CodeUri":"s3://bucket/key","Runtime":"nodejs6.10","FunctionName":"Excluded"}`)

				result := &global.Function{}
				err := json.Unmarshal(property, result)
				It("should throw an error on unmarshal", func() {
					Expect(err).Should(HaveOccurred())
				})
			})

		})

	})

	Context("with template described as Go structs with globals", func() {

		template := cloudformation.NewTemplate()
		template.Globals["Function"] = &global.Function{
			Runtime: "nodejs12.x",
			Timeout: 180,
			Handler: "index.handler",
			Environment: &serverless.Function_FunctionEnvironment{
				Variables: map[string]string{
					"TABLE_NAME": "data-table",
				},
			},
		}

		expected := `{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Globals": {
    "Function": {
      "Environment": {
        "Variables": {
          "TABLE_NAME": "data-table"
        }
      },
      "Handler": "index.handler",
      "Runtime": "nodejs12.x",
      "Timeout": 180
    }
  }
}`

		output, err := template.JSON()

		It("it should convert to expected JSON output", func() {
			Expect(err).To(BeNil())
			Expect(output).ShouldNot(BeNil())
			Expect(string(output)).To(Equal(expected))
		})

	})

	Context("with a YAML template that contains global resources", func() {

		template, err := goformation.Open("../test/yaml/globals.yml")

		It("should successfully parse the template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		It("should have a global Function, with properties set", func() {

			expected := &global.Function{
				Runtime: "nodejs12.x",
				Timeout: 180,
				Handler: "index.handler",
				Environment: &serverless.Function_FunctionEnvironment{
					Variables: map[string]string{
						"TABLE_NAME": "data-table",
					},
				},
			}

			Expect(template.Globals["Function"]).To(Equal(expected))

		})

	})

})
