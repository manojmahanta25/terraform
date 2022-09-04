package EC2Config

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/ec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Ec2Instance struct {
	ec2.Instance
}

func NewEC2Instance(scope constructs.Construct, name string) *Ec2Instance {
	instance := ec2.NewInstance(scope, jsii.String(name), &ec2.InstanceConfig{
		Ami:                      jsii.String("ami-052efd3df9dad4825"),
		InstanceType:             jsii.String("t2.micro"),
		KeyName:                  jsii.String("test-key"),
		AssociatePublicIpAddress: true,
	})
	return &Ec2Instance{instance}
}

func (ec2 *Ec2Instance) AddSubnetId(subId *string) *Ec2Instance {
	ec2.SetSubnetId(subId)
	return ec2
}

func (ec2 *Ec2Instance) VpcSecurityGroupIds(vpcSecId *[]*string) *Ec2Instance {
	ec2.SetVpcSecurityGroupIds(vpcSecId)
	return ec2
}

func (ec2 *Ec2Instance) AddUserData(userdata string) Ec2Instance {
	ec2.SetUserData(jsii.String(userdata))
	return *ec2
}
