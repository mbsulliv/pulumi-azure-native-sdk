// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vmwarecloudsimple

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Dedicated cloud service model
// Azure REST API version: 2019-04-01. Prior API version in Azure Native 1.x: 2019-04-01.
type DedicatedCloudService struct {
	pulumi.CustomResourceState

	// gateway Subnet for the account. It will collect the subnet address and always treat it as /28
	GatewaySubnet pulumi.StringOutput `pulumi:"gatewaySubnet"`
	// indicates whether account onboarded or not in a given region
	IsAccountOnboarded pulumi.StringOutput `pulumi:"isAccountOnboarded"`
	// Azure region
	Location pulumi.StringOutput `pulumi:"location"`
	// {dedicatedCloudServiceName}
	Name pulumi.StringOutput `pulumi:"name"`
	// total nodes purchased
	Nodes pulumi.IntOutput `pulumi:"nodes"`
	// link to a service management web portal
	ServiceURL pulumi.StringOutput `pulumi:"serviceURL"`
	// The list of tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// {resourceProviderNamespace}/{resourceType}
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDedicatedCloudService registers a new resource with the given unique name, arguments, and options.
func NewDedicatedCloudService(ctx *pulumi.Context,
	name string, args *DedicatedCloudServiceArgs, opts ...pulumi.ResourceOption) (*DedicatedCloudService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewaySubnet == nil {
		return nil, errors.New("invalid value for required argument 'GatewaySubnet'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:vmwarecloudsimple/v20190401:DedicatedCloudService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DedicatedCloudService
	err := ctx.RegisterResource("azure-native:vmwarecloudsimple:DedicatedCloudService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedCloudService gets an existing DedicatedCloudService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedCloudService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedCloudServiceState, opts ...pulumi.ResourceOption) (*DedicatedCloudService, error) {
	var resource DedicatedCloudService
	err := ctx.ReadResource("azure-native:vmwarecloudsimple:DedicatedCloudService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedCloudService resources.
type dedicatedCloudServiceState struct {
}

type DedicatedCloudServiceState struct {
}

func (DedicatedCloudServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCloudServiceState)(nil)).Elem()
}

type dedicatedCloudServiceArgs struct {
	// dedicated cloud Service name
	DedicatedCloudServiceName *string `pulumi:"dedicatedCloudServiceName"`
	// gateway Subnet for the account. It will collect the subnet address and always treat it as /28
	GatewaySubnet string `pulumi:"gatewaySubnet"`
	// Azure region
	Location *string `pulumi:"location"`
	// The name of the resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DedicatedCloudService resource.
type DedicatedCloudServiceArgs struct {
	// dedicated cloud Service name
	DedicatedCloudServiceName pulumi.StringPtrInput
	// gateway Subnet for the account. It will collect the subnet address and always treat it as /28
	GatewaySubnet pulumi.StringInput
	// Azure region
	Location pulumi.StringPtrInput
	// The name of the resource group
	ResourceGroupName pulumi.StringInput
	// The list of tags
	Tags pulumi.StringMapInput
}

func (DedicatedCloudServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCloudServiceArgs)(nil)).Elem()
}

type DedicatedCloudServiceInput interface {
	pulumi.Input

	ToDedicatedCloudServiceOutput() DedicatedCloudServiceOutput
	ToDedicatedCloudServiceOutputWithContext(ctx context.Context) DedicatedCloudServiceOutput
}

func (*DedicatedCloudService) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCloudService)(nil)).Elem()
}

func (i *DedicatedCloudService) ToDedicatedCloudServiceOutput() DedicatedCloudServiceOutput {
	return i.ToDedicatedCloudServiceOutputWithContext(context.Background())
}

func (i *DedicatedCloudService) ToDedicatedCloudServiceOutputWithContext(ctx context.Context) DedicatedCloudServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCloudServiceOutput)
}

func (i *DedicatedCloudService) ToOutput(ctx context.Context) pulumix.Output[*DedicatedCloudService] {
	return pulumix.Output[*DedicatedCloudService]{
		OutputState: i.ToDedicatedCloudServiceOutputWithContext(ctx).OutputState,
	}
}

type DedicatedCloudServiceOutput struct{ *pulumi.OutputState }

func (DedicatedCloudServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCloudService)(nil)).Elem()
}

func (o DedicatedCloudServiceOutput) ToDedicatedCloudServiceOutput() DedicatedCloudServiceOutput {
	return o
}

func (o DedicatedCloudServiceOutput) ToDedicatedCloudServiceOutputWithContext(ctx context.Context) DedicatedCloudServiceOutput {
	return o
}

func (o DedicatedCloudServiceOutput) ToOutput(ctx context.Context) pulumix.Output[*DedicatedCloudService] {
	return pulumix.Output[*DedicatedCloudService]{
		OutputState: o.OutputState,
	}
}

// gateway Subnet for the account. It will collect the subnet address and always treat it as /28
func (o DedicatedCloudServiceOutput) GatewaySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.GatewaySubnet }).(pulumi.StringOutput)
}

// indicates whether account onboarded or not in a given region
func (o DedicatedCloudServiceOutput) IsAccountOnboarded() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.IsAccountOnboarded }).(pulumi.StringOutput)
}

// Azure region
func (o DedicatedCloudServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// {dedicatedCloudServiceName}
func (o DedicatedCloudServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// total nodes purchased
func (o DedicatedCloudServiceOutput) Nodes() pulumi.IntOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.IntOutput { return v.Nodes }).(pulumi.IntOutput)
}

// link to a service management web portal
func (o DedicatedCloudServiceOutput) ServiceURL() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.ServiceURL }).(pulumi.StringOutput)
}

// The list of tags
func (o DedicatedCloudServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// {resourceProviderNamespace}/{resourceType}
func (o DedicatedCloudServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedCloudServiceOutput{})
}
