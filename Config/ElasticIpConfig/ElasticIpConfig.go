package ElasticIpConfig

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/ec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewElasticIp(scope constructs.Construct) {
	ec2.NewEip(scope, jsii.String("elastic-ip-test"), &ec2.EipConfig{
		Vpc: nil,
	})
}

/*
TODO
1. LoadBalancer
2. Lister
3. target group
4. target group attachment
5. listener rule

*/
