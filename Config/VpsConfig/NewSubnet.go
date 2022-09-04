package VpsConfig

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewSubnet(scope constructs.Construct, vpcId *string, name, cidrBlock, availabilityZone string) vpc.Subnet {
	return vpc.NewSubnet(scope, jsii.String(name), &vpc.SubnetConfig{
		VpcId:            vpcId,
		CidrBlock:        jsii.String(cidrBlock),
		AvailabilityZone: jsii.String(availabilityZone),
		Tags: &map[string]*string{
			"Name": jsii.String("test-" + name),
		},
	})
}
