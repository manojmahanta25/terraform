package VpcWithEc2

import (
	"cdk.tf/go/stack/Config"
	"cdk.tf/go/stack/Config/EC2Config"
	"cdk.tf/go/stack/Config/LoadBalancer"
	"cdk.tf/go/stack/Config/SecurityGroupConfig"
	"cdk.tf/go/stack/Config/VpsConfig"
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewVpcWithEc2Stack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	Config.AwsConfig(stack)
	vpcInstance := VpsConfig.NewVpc(stack)
	subnetPublic := VpsConfig.NewSubnet(stack, vpcInstance.Id(), "subnet_public", "10.0.1.0/24", "us-east-1a")
	subnetPublic2 := VpsConfig.NewSubnet(stack, vpcInstance.Id(), "subnet_public2", "10.0.3.0/24", "us-east-1b")
	VpsConfig.NewSubnet(stack, vpcInstance.Id(), "subnet_private", "10.0.2.0/24", "us-east-1a")
	VpsConfig.NewSubnet(stack, vpcInstance.Id(), "subnet_private2", "10.0.4.0/24", "us-east-1b")
	securityGroup := VpsConfig.ToBeDone(stack, vpcInstance.Id(), subnetPublic.Id(), subnetPublic2.Id())
	userData := `#!/bin/bash
echo "Hello, World 200" > index.html
python3 -m http.server 8080 &`

	//	userData2 := `#!/bin/bash
	//sudo apt update -y
	//sudo apt install apache2 -y
	//sudo systemctl start apache2
	//sudo bash -c 'echo web server2 > /var/www/html/index.html'`
	userData2 := `#!/bin/bash
echo "Hello, World 202" > index.html
python3 -m http.server 8080 &`

	instance := EC2Config.NewEC2Instance(stack, "instance 1").AddSubnetId(subnetPublic.Id()).VpcSecurityGroupIds(jsii.Strings(*securityGroup.Id())).AddUserData(userData)
	instance2 := EC2Config.NewEC2Instance(stack, "instance 2").AddSubnetId(subnetPublic.Id()).VpcSecurityGroupIds(jsii.Strings(*securityGroup.Id())).AddUserData(userData2)
	ingress := []*vpc.SecurityGroupIngress{
		{
			FromPort:   jsii.Number(80),
			ToPort:     jsii.Number(80),
			Protocol:   jsii.String("tcp"),
			CidrBlocks: jsii.Strings("0.0.0.0/0"),
		},
	}
	lbSecurity := SecurityGroupConfig.NewSecurityGroup(stack, "lb-security").AddVpcId(vpcInstance.Id()).AddIngress(ingress)
	loadBalancer := LoadBalancer.NewLoadBalancer(stack).AddSecurityGroupToAlb(jsii.Strings(*lbSecurity.Id())).AddSubnetToAlb(jsii.Strings(*subnetPublic.Id(), *subnetPublic2.Id())).Done()
	lbListener := LoadBalancer.NewALBListener(stack, loadBalancer.Arn()).Done()
	targetGroup := LoadBalancer.NewELBNewTargetGroup(stack).AddVPCId(vpcInstance.Id()).Done()
	LoadBalancer.NewElbListenerRule(stack, lbListener.Arn(), targetGroup.Arn()).Done()
	LoadBalancer.NewTargetGroupAttachment(stack, "instance-1-attachment", targetGroup.Arn(), instance.Id()).Done()
	LoadBalancer.NewTargetGroupAttachment(stack, "instance-2-attachment", targetGroup.Arn(), instance2.Id()).Done()
	cdktf.NewTerraformOutput(stack, jsii.String("vpc_id"), &cdktf.TerraformOutputConfig{
		Value: vpcInstance.Id(),
	})
	cdktf.NewTerraformOutput(stack, jsii.String("public_ip 1"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})
	cdktf.NewTerraformOutput(stack, jsii.String("public_ip 2"), &cdktf.TerraformOutputConfig{
		Value: instance2.PublicIp(),
	})
	cdktf.NewTerraformOutput(stack, jsii.String("loadBalancer"), &cdktf.TerraformOutputConfig{
		Value: loadBalancer.DnsName(),
	})
	return stack
}
