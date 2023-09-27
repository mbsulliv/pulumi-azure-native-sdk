// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetesconfiguration

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Azure Arc PrivateLinkScope definition.
// Azure REST API version: 2022-04-02-preview. Prior API version in Azure Native 1.x: 2022-04-02-preview
type PrivateLinkScope struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties that define a Azure Arc PrivateLinkScope resource.
	Properties KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkScope registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkScope(ctx *pulumi.Context,
	name string, args *PrivateLinkScopeArgs, opts ...pulumi.ResourceOption) (*PrivateLinkScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220402preview:PrivateLinkScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkScope
	err := ctx.RegisterResource("azure-native:kubernetesconfiguration:PrivateLinkScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkScope gets an existing PrivateLinkScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkScopeState, opts ...pulumi.ResourceOption) (*PrivateLinkScope, error) {
	var resource PrivateLinkScope
	err := ctx.ReadResource("azure-native:kubernetesconfiguration:PrivateLinkScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkScope resources.
type privateLinkScopeState struct {
}

type PrivateLinkScopeState struct {
}

func (PrivateLinkScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopeState)(nil)).Elem()
}

type privateLinkScopeArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Properties that define a Azure Arc PrivateLinkScope resource.
	Properties *KubernetesConfigurationPrivateLinkScopeProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName *string `pulumi:"scopeName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateLinkScope resource.
type PrivateLinkScopeArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Properties that define a Azure Arc PrivateLinkScope resource.
	Properties KubernetesConfigurationPrivateLinkScopePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateLinkScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopeArgs)(nil)).Elem()
}

type PrivateLinkScopeInput interface {
	pulumi.Input

	ToPrivateLinkScopeOutput() PrivateLinkScopeOutput
	ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput
}

func (*PrivateLinkScope) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScope)(nil)).Elem()
}

func (i *PrivateLinkScope) ToPrivateLinkScopeOutput() PrivateLinkScopeOutput {
	return i.ToPrivateLinkScopeOutputWithContext(context.Background())
}

func (i *PrivateLinkScope) ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopeOutput)
}

func (i *PrivateLinkScope) ToOutput(ctx context.Context) pulumix.Output[*PrivateLinkScope] {
	return pulumix.Output[*PrivateLinkScope]{
		OutputState: i.ToPrivateLinkScopeOutputWithContext(ctx).OutputState,
	}
}

type PrivateLinkScopeOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScope)(nil)).Elem()
}

func (o PrivateLinkScopeOutput) ToPrivateLinkScopeOutput() PrivateLinkScopeOutput {
	return o
}

func (o PrivateLinkScopeOutput) ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput {
	return o
}

func (o PrivateLinkScopeOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateLinkScope] {
	return pulumix.Output[*PrivateLinkScope]{
		OutputState: o.OutputState,
	}
}

// The geo-location where the resource lives
func (o PrivateLinkScopeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrivateLinkScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties that define a Azure Arc PrivateLinkScope resource.
func (o PrivateLinkScopeOutput) Properties() KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkScope) KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput {
		return v.Properties
	}).(KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PrivateLinkScopeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkScope) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PrivateLinkScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateLinkScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkScopeOutput{})
}
