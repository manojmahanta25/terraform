package main

import (
	"cdk.tf/go/stack/Stack/VpcWithEc2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	//Ec2Stack.Ec2Stack(app, "aws-terraform")
	VpcWithEc2.NewVpcWithEc2Stack(app, "aws-terraform")
	app.Synth()

}
