package terraform
import  (
	"testing"
        "github.com/gruntwork-io/terratest/modules/terraform"
	"fmt"
	"time"
	 http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

)

func TestTask2Plan(t *testing.T) {
terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../task2",
	})
 terraform.InitAndPlan(t, terraformOptions)
}



func TestTask2ApplicationReachability(t *testing.T) {
terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
                TerraformDir: "../task2",
        })

 defer terraform.Destroy(t, terraformOptions)
 terraform.InitAndApply(t, terraformOptions)
 expectedResponse := "Name : Flugel"
 publicdns := terraform.Output(t, terraformOptions,"dns_name")
 url := fmt.Sprintf("http://%s", publicdns)
//http_helper.HttpGet(t, url, nil)
 http_helper.HttpGetWithRetry(t, url, nil, 200, expectedResponse, 5, 5*time.Second)
}
