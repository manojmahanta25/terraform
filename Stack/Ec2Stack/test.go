package Ec2Stack

import (
	"cdk.tf/go/stack/Config"
	"cdk.tf/go/stack/generated/hashicorp/aws/ec2"
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func Ec2Stack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	Config.AwsConfig(stack)
	//filter := []*ec2.DataAwsEc2InstanceTypeOfferingsFilter{
	//	{
	//		Name:   jsii.String("instance-type"),
	//		Values: jsii.Strings("t2.micro", "t3.micro"),
	//	},
	//}
	data := ec2.NewDataAwsEc2InstanceTypeOfferings(stack, jsii.String("test"), &ec2.DataAwsEc2InstanceTypeOfferingsConfig{})
	values := data.InstanceTypes()
	for i := 0; i >= 10; i++ {
		x := *values
		y := x[i]
		fmt.Println(*y)
	}
	cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
		Value: "done",
	})
	return stack
}
