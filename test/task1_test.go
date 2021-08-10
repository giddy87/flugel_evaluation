package terraform

import  (
	"testing"
	"github.com/stretchr/testify/assert"
        "github.com/gruntwork-io/terratest/modules/terraform"
        "github.com/gruntwork-io/terratest/modules/aws"

)

func TestTask1Plan(t *testing.T) {
         terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../task1",
	})
         terraform.InitAndPlan(t, terraformOptions)
}


func TestTagsValidation(t *testing.T) {
terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
                TerraformDir: "../task1",
        })

//Clean up after Tests
 defer terraform.Destroy(t, terraformOptions)
 terraform.InitAndApply(t, terraformOptions)
awsRegion := "us-east-2"
instanceID := terraform.Output(t, terraformOptions, "instance_id")

instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)

	// website::tag::3::Check if the EC2 instance with a given tag and name is set.
	testingTag, containsFlugelNameTag := instanceTags["Name"]
	assert.True(t, containsFlugelNameTag)
	assert.Equal(t, "Flugel", testingTag)

	// Verify that our expected name tag is one of the tags
//	nameTag, containsNameTag := instanceTags["Name"]
//	assert.True(t, containsNameTag)
//	assert.Equal(t, expectedName, nameTag)instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)


// tags := terraform.OutputMap(t, terraformOptions,"tags")
// assert.Contains(t, tags, "Name: Flugel")
// assert.Contains(t, tags, "Owner: InfraTeam")



}
