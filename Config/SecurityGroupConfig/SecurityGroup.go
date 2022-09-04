package SecurityGroupConfig

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SecurityGroup struct {
	vpc.SecurityGroup
}

func NewSecurityGroup(scope constructs.Construct, name string) *SecurityGroup {

	egress := []*vpc.SecurityGroupEgress{{
		CidrBlocks: jsii.Strings("0.0.0.0/0"),
		ToPort:     jsii.Number(0),
		FromPort:   jsii.Number(0),
		Protocol:   jsii.String("-1"),
	},
	}
	security := vpc.NewSecurityGroup(scope, jsii.String(name), &vpc.SecurityGroupConfig{
		Name: jsii.String(name),
		//Ingress: ingress,
		Egress: egress,
	})

	//vpc.NewSecurityGroupRule(scope, jsii.String("allow_alb_http_inbound"), &vpc.SecurityGroupRuleConfig{
	//	Type:            jsii.String("ingress"),
	//	SecurityGroupId: security.Id(),
	//	FromPort:        jsii.Number(80),
	//	ToPort:          jsii.Number(80),
	//	Protocol:        jsii.String("tcp"),
	//	CidrBlocks:      jsii.Strings("0.0.0.0/0"),
	//})
	return &SecurityGroup{security}
}

func (sG *SecurityGroup) Done() SecurityGroup {
	return *sG
}

func (sG *SecurityGroup) AddIngress(ingress interface{}) *SecurityGroup {
	sG.PutIngress(ingress)
	return sG
}

func (sG *SecurityGroup) AddEgress(egress interface{}) *SecurityGroup {
	sG.PutEgress(egress)
	return sG
}

func (sG *SecurityGroup) AddVpcId(s *string) *SecurityGroup {
	sG.SetVpcId(s)
	return sG
}
