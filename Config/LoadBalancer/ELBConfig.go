package LoadBalancer

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/elb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ElbConfig struct {
	elb.Lb
}

func NewLoadBalancer(scope constructs.Construct) *ElbConfig {
	lb := elb.NewLb(scope, jsii.String("elb_test"), &elb.LbConfig{
		LoadBalancerType: jsii.String("application"),
		Name:             jsii.String("lb-for-test"),
	})

	return &ElbConfig{lb}
}

func (alb *ElbConfig) AddSecurityGroupToAlb(s *[]*string) *ElbConfig {
	alb.SetSecurityGroups(s)
	return alb
}

func (alb *ElbConfig) AddSubnetToAlb(s *[]*string) *ElbConfig {
	alb.SetSubnets(s)
	return alb
}

func (alb *ElbConfig) Done() ElbConfig {
	return *alb
}
