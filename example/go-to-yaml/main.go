package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/awslabs/goformation/cloudformation"
)

func main() {

	// Create a new CloudFormation template
	template := cloudformation.NewTemplate()

	// Create an Amazon SNS topic, with a unique name based off the current timestamp
	template.Resources["MyTopic"] = &cloudformation.AWSSNSTopic{
		TopicName: "my-topic-" + strconv.FormatInt(time.Now().Unix(), 10),
	}

	// Create a subscription, connected to our topic, that forwards notifications to an email address
	template.Resources["MyTopicSubscription"] = &cloudformation.AWSSNSSubscription{
		TopicArn: cloudformation.Ref("MyTopic"),
		Protocol: "email",
		Endpoint: "some.email@example.com",
	}

	// Let's see the JSON
	j, err := template.JSON()
	if err != nil {
		fmt.Printf("Failed to generate JSON: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(j))
	}

	y, err := template.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(y))
	}

}
