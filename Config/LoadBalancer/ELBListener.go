package LoadBalancer

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/elb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ALBListener struct {
	elb.LbListener
}

func NewALBListener(scope constructs.Construct, loadBalancerArn *string) *ALBListener {
	albL := elb.NewLbListener(scope, jsii.String("elb-listener-test"), &elb.LbListenerConfig{
		LoadBalancerArn: loadBalancerArn,
		Port:            jsii.Number(80),
		Protocol:        jsii.String("HTTP"),
		DefaultAction: &[]*elb.LbListenerDefaultAction{
			{
				Type: jsii.String("fixed-response"),
				FixedResponse: &elb.LbListenerDefaultActionFixedResponse{
					ContentType: jsii.String("text/plain"),
					MessageBody: jsii.String("404: page not found"),
					StatusCode:  jsii.String("404"),
				},
			},
		},
	})
	return &ALBListener{albL}
}

func (albL *ALBListener) Done() ALBListener {
	return *albL
}

type ElbListenerRule struct {
	elb.LbListenerRule
}

func NewElbListenerRule(scope constructs.Construct, ListenerArn, TargetGroupArn *string) *ElbListenerRule {
	elbLR := elb.NewLbListenerRule(scope, jsii.String("testElbListenerRule"), &elb.LbListenerRuleConfig{
		ListenerArn: ListenerArn,
		Condition: &[]*elb.LbListenerRuleCondition{
			{
				PathPattern: &elb.LbListenerRuleConditionPathPattern{
					Values: jsii.Strings("*"),
				},
			},
		},
		Priority: jsii.Number(100),
		Action: &[]*elb.LbListenerRuleAction{
			{
				Type:           jsii.String("forward"),
				TargetGroupArn: TargetGroupArn,
			},
		},
	})
	return &ElbListenerRule{elbLR}
}

func (elbLR *ElbListenerRule) Done() ElbListenerRule {
	return *elbLR
}
