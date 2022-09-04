package Config

import (
	"cdk.tf/go/stack/generated/hashicorp/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func AwsConfig(scope constructs.Construct) aws.AwsProvider {
	return aws.NewAwsProvider(scope, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region:    jsii.String("us-east-1"),
		AccessKey: jsii.String("AKIAWB2KIRNM5AXN4774"),                     //Access Key
		SecretKey: jsii.String("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"), //Secret Key
	})
}
