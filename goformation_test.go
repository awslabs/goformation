package goformation_test

import (
	"encoding/json"

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

		data, err := json.MarshalIndent(t1, "", "    ")
		It("should successfully marshal to JSON", func() {
			Expect(err).To(BeNil())
		})

		t2 := &resources.Template{}
		err = json.Unmarshal(data, t2)
		It("should then unmarshal back to Go", func() {
			Expect(err).To(BeNil())
			Expect(t2.Resources).To(HaveKey("MyLambdaFunction"))
		})

		functions := resources.GetAllAWSLambdaFunctionResources(t2)
		It("should be able to retrieve all Lambda functions with GetAllAWSLambdaFunction(template)", func() {
			Expect(functions).To(HaveLen(1))
		})

		function, err := resources.GetAWSLambdaFunctionWithName("MyLambdaFunction", t2)
		It("should be able to retrieve a specific Lambda function with GetAWSLambdaFunctionWithName(template, name)", func() {
			Expect(err).To(BeNil())
			Expect(function).To(BeAssignableToTypeOf(&resources.AWSLambdaFunction{}))
		})

		It("should have the correct Handler property", func() {
			Expect(function.Handler).To(Equal("nodejs6.10"))
		})

	})

})
