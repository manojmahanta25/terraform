package VpsConfig

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewVpc(scope constructs.Construct) vpc.Vpc {
	return vpc.NewVpc(scope, jsii.String("my-vpc"), &vpc.VpcConfig{
		CidrBlock:       jsii.String("10.0.0.0/16"),
		InstanceTenancy: jsii.String("default"),
		Tags: &map[string]*string{
			"Name": jsii.String("test-vpc"),
		},
	})
}
