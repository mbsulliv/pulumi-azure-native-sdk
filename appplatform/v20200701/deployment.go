// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment resource payload
type Deployment struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Deployment resource
	Properties DeploymentResourcePropertiesResponseOutput `pulumi:"properties"`
	// Sku of the Deployment resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToDeploymentResourcePropertiesPtrOutput().ApplyT(func(v *DeploymentResourceProperties) *DeploymentResourceProperties { return v.Defaults() }).(DeploymentResourcePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:Deployment"),
		},
	})
	opts = append(opts, aliases)
	var resource Deployment
	err := ctx.RegisterResource("azure-native:appplatform/v20200701:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("azure-native:appplatform/v20200701:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
}

type DeploymentState struct {
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the Deployment resource.
	DeploymentName *string `pulumi:"deploymentName"`
	// Properties of the Deployment resource
	Properties *DeploymentResourceProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// Sku of the Deployment resource
	Sku *Sku `pulumi:"sku"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput
	// The name of the Deployment resource.
	DeploymentName pulumi.StringPtrInput
	// Properties of the Deployment resource
	Properties DeploymentResourcePropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
	// Sku of the Deployment resource
	Sku SkuPtrInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

// The name of the resource.
func (o DeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the Deployment resource
func (o DeploymentOutput) Properties() DeploymentResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *Deployment) DeploymentResourcePropertiesResponseOutput { return v.Properties }).(DeploymentResourcePropertiesResponseOutput)
}

// Sku of the Deployment resource
func (o DeploymentOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Deployment) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The type of the resource.
func (o DeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentOutput{})
}
