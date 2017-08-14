package goformation_test

import (
	"encoding/json"

	. "github.com/paulmaddox/goformation"
	"github.com/paulmaddox/goformation/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goformation", func() {

	Context("with a template defined as Go code", func() {

		t1 := &Template{
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

		t2 := &Template{}
		err = json.Unmarshal(data, t2)
		It("should then unmarshal back to Go", func() {
			Expect(err).To(BeNil())
			Expect(t2.Resources).To(HaveKey("MyLambdaFunction"))
		})

	})

})
