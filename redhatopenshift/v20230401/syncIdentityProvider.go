// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SyncIdentityProvider represents a SyncIdentityProvider
type SyncIdentityProvider struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name      pulumi.StringOutput    `pulumi:"name"`
	Resources pulumi.StringPtrOutput `pulumi:"resources"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSyncIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewSyncIdentityProvider(ctx *pulumi.Context,
	name string, args *SyncIdentityProviderArgs, opts ...pulumi.ResourceOption) (*SyncIdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:redhatopenshift:SyncIdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20220904:SyncIdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20230701preview:SyncIdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20230904:SyncIdentityProvider"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SyncIdentityProvider
	err := ctx.RegisterResource("azure-native:redhatopenshift/v20230401:SyncIdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncIdentityProvider gets an existing SyncIdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncIdentityProviderState, opts ...pulumi.ResourceOption) (*SyncIdentityProvider, error) {
	var resource SyncIdentityProvider
	err := ctx.ReadResource("azure-native:redhatopenshift/v20230401:SyncIdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncIdentityProvider resources.
type syncIdentityProviderState struct {
}

type SyncIdentityProviderState struct {
}

func (SyncIdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncIdentityProviderState)(nil)).Elem()
}

type syncIdentityProviderArgs struct {
	// The name of the SyncIdentityProvider resource.
	ChildResourceName *string `pulumi:"childResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName string  `pulumi:"resourceName"`
	Resources    *string `pulumi:"resources"`
}

// The set of arguments for constructing a SyncIdentityProvider resource.
type SyncIdentityProviderArgs struct {
	// The name of the SyncIdentityProvider resource.
	ChildResourceName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the OpenShift cluster resource.
	ResourceName pulumi.StringInput
	Resources    pulumi.StringPtrInput
}

func (SyncIdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncIdentityProviderArgs)(nil)).Elem()
}

type SyncIdentityProviderInput interface {
	pulumi.Input

	ToSyncIdentityProviderOutput() SyncIdentityProviderOutput
	ToSyncIdentityProviderOutputWithContext(ctx context.Context) SyncIdentityProviderOutput
}

func (*SyncIdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncIdentityProvider)(nil)).Elem()
}

func (i *SyncIdentityProvider) ToSyncIdentityProviderOutput() SyncIdentityProviderOutput {
	return i.ToSyncIdentityProviderOutputWithContext(context.Background())
}

func (i *SyncIdentityProvider) ToSyncIdentityProviderOutputWithContext(ctx context.Context) SyncIdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncIdentityProviderOutput)
}

type SyncIdentityProviderOutput struct{ *pulumi.OutputState }

func (SyncIdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncIdentityProvider)(nil)).Elem()
}

func (o SyncIdentityProviderOutput) ToSyncIdentityProviderOutput() SyncIdentityProviderOutput {
	return o
}

func (o SyncIdentityProviderOutput) ToSyncIdentityProviderOutputWithContext(ctx context.Context) SyncIdentityProviderOutput {
	return o
}

// The name of the resource
func (o SyncIdentityProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncIdentityProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SyncIdentityProviderOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncIdentityProvider) pulumi.StringPtrOutput { return v.Resources }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SyncIdentityProviderOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SyncIdentityProvider) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SyncIdentityProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncIdentityProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncIdentityProviderOutput{})
}
