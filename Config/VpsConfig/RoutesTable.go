package VpsConfig

import (
	"cdk.tf/go/stack/Config/SecurityGroupConfig"
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func ToBeDone(scope constructs.Construct, vpcId *string, subnetPubId ...*string) vpc.SecurityGroup {
	internetGateway := vpc.NewInternetGateway(scope, jsii.String("my-Internet-Gateway"), &vpc.InternetGatewayConfig{
		VpcId: vpcId,
		Tags: &map[string]*string{
			"Name": jsii.String("test-internet-gateway"),
		},
	})
	routes := []*vpc.RouteTableRoute{{
		CidrBlock: jsii.String("0.0.0.0/0"),
		GatewayId: internetGateway.Id(),
	}, {
		Ipv6CidrBlock: jsii.String("::/0"),
		GatewayId:     internetGateway.Id(),
	}}
	publicRouteTable := vpc.NewRouteTable(scope, jsii.String("my-public-route-table"), &vpc.RouteTableConfig{
		VpcId: vpcId,
		Tags: &map[string]*string{
			"Name": jsii.String("test-public-route-table"),
		},
		Route: &routes,
	})
	for i, subNet := range subnetPubId {
		vpc.NewRouteTableAssociation(scope, jsii.String("my-public-route-table-association-"+string(rune(i))), &vpc.RouteTableAssociationConfig{
			SubnetId:     subNet,
			RouteTableId: publicRouteTable.Id(),
		})
	}
	//vpc.NewRouteTableAssociation(scope, jsii.String("my-public-route-table-association"), &vpc.RouteTableAssociationConfig{
	//	SubnetId:     subnetPubId,
	//	RouteTableId: publicRouteTable.Id(),
	//})
	ingress := []*vpc.SecurityGroupIngress{
		{
			FromPort:   jsii.Number(8080),
			ToPort:     jsii.Number(8080),
			Protocol:   jsii.String("tcp"),
			CidrBlocks: jsii.Strings("0.0.0.0/0"),
		},
		{
			FromPort:   jsii.Number(22),
			ToPort:     jsii.Number(22),
			Protocol:   jsii.String("tcp"),
			CidrBlocks: jsii.Strings("0.0.0.0/0"),
		},
	}
	egress := []*vpc.SecurityGroupEgress{{
		CidrBlocks: jsii.Strings("0.0.0.0/0"),
		ToPort:     jsii.Number(0),
		FromPort:   jsii.Number(0),
		Protocol:   jsii.String("-1"),
	},
	}

	securityGroup := SecurityGroupConfig.NewSecurityGroup(scope, "my-vpc-security-group").AddVpcId(vpcId).AddIngress(ingress).AddEgress(egress).Done()
	return securityGroup
}
