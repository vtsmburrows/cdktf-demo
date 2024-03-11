package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v19/provider"
	vpc "github.com/cdktf/cdktf-provider-aws-go/aws/v19/vpc"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	awsprovider.NewAwsProvider(stack, jsii.String("aws"), &awsprovider.AwsProviderConfig{
		Region: jsii.String("us-west-2"),
	})

	vpc.NewVpc(stack, jsii.String("MyVpc"), &vpc.VpcConfig{
		CidrBlock: jsii.String("60.0.0.0/16"),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "go-cdk")

	app.Synth()
}
