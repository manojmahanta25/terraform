package LoadBalancer

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/elb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ElbTargetGroup struct {
	elb.LbTargetGroup
}

func NewELBNewTargetGroup(scope constructs.Construct) *ElbTargetGroup {
	tg := elb.NewLbTargetGroup(scope, jsii.String("targetGroupTest"), &elb.LbTargetGroupConfig{
		Name:     jsii.String("TargetGroupTest"),
		Protocol: jsii.String("HTTP"),
		Port:     jsii.Number(8080),
		HealthCheck: &elb.LbTargetGroupHealthCheck{
			HealthyThreshold:   jsii.Number(2),
			Interval:           jsii.Number(15),
			Matcher:            jsii.String("200"),
			Path:               jsii.String("/"),
			Protocol:           jsii.String("HTTP"),
			Port:               jsii.String("8080"),
			Timeout:            jsii.Number(3),
			UnhealthyThreshold: jsii.Number(2),
		},
	})
	return &ElbTargetGroup{tg}
}

func (eTg *ElbTargetGroup) AddVPCId(vpcId *string) *ElbTargetGroup {
	eTg.SetVpcId(vpcId)
	return eTg
}

func (eTg *ElbTargetGroup) AddPortNumber(port float64) *ElbTargetGroup {
	eTg.SetPort(jsii.Number(port))
	return eTg
}

func (eTg *ElbTargetGroup) Done() ElbTargetGroup {
	return *eTg
}

type ElbTargetGroupAttachment struct {
	elb.LbTargetGroupAttachment
}

func NewTargetGroupAttachment(scope constructs.Construct, name string, targetGroupArn, targetId *string) *ElbTargetGroupAttachment {
	tgAg := elb.NewLbTargetGroupAttachment(scope, jsii.String(name), &elb.LbTargetGroupAttachmentConfig{
		Port:           jsii.Number(8080),
		TargetGroupArn: targetGroupArn,
		TargetId:       targetId,
	})
	return &ElbTargetGroupAttachment{tgAg}
}

func (eTgA *ElbTargetGroupAttachment) Done() ElbTargetGroupAttachment {
	return *eTgA
}
