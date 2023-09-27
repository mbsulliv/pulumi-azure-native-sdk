// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// SyncSet represents a SyncSet for an Azure Red Hat OpenShift Cluster.
type SyncSet struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resources represents the SyncSets configuration.
	Resources pulumi.StringPtrOutput `pulumi:"resources"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSyncSet registers a new resource with the given unique name, arguments, and options.
func NewSyncSet(ctx *pulumi.Context,
	name string, args *SyncSetArgs, opts ...pulumi.ResourceOption) (*SyncSet, error) {
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
			Type: pulumi.String("azure-native:redhatopenshift:SyncSet"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20220904:SyncSet"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20230701preview:SyncSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SyncSet
	err := ctx.RegisterResource("azure-native:redhatopenshift/v20230401:SyncSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncSet gets an existing SyncSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncSetState, opts ...pulumi.ResourceOption) (*SyncSet, error) {
	var resource SyncSet
	err := ctx.ReadResource("azure-native:redhatopenshift/v20230401:SyncSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncSet resources.
type syncSetState struct {
}

type SyncSetState struct {
}

func (SyncSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncSetState)(nil)).Elem()
}

type syncSetArgs struct {
	// The name of the SyncSet resource.
	ChildResourceName *string `pulumi:"childResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// Resources represents the SyncSets configuration.
	Resources *string `pulumi:"resources"`
}

// The set of arguments for constructing a SyncSet resource.
type SyncSetArgs struct {
	// The name of the SyncSet resource.
	ChildResourceName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the OpenShift cluster resource.
	ResourceName pulumi.StringInput
	// Resources represents the SyncSets configuration.
	Resources pulumi.StringPtrInput
}

func (SyncSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncSetArgs)(nil)).Elem()
}

type SyncSetInput interface {
	pulumi.Input

	ToSyncSetOutput() SyncSetOutput
	ToSyncSetOutputWithContext(ctx context.Context) SyncSetOutput
}

func (*SyncSet) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncSet)(nil)).Elem()
}

func (i *SyncSet) ToSyncSetOutput() SyncSetOutput {
	return i.ToSyncSetOutputWithContext(context.Background())
}

func (i *SyncSet) ToSyncSetOutputWithContext(ctx context.Context) SyncSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncSetOutput)
}

func (i *SyncSet) ToOutput(ctx context.Context) pulumix.Output[*SyncSet] {
	return pulumix.Output[*SyncSet]{
		OutputState: i.ToSyncSetOutputWithContext(ctx).OutputState,
	}
}

type SyncSetOutput struct{ *pulumi.OutputState }

func (SyncSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncSet)(nil)).Elem()
}

func (o SyncSetOutput) ToSyncSetOutput() SyncSetOutput {
	return o
}

func (o SyncSetOutput) ToSyncSetOutputWithContext(ctx context.Context) SyncSetOutput {
	return o
}

func (o SyncSetOutput) ToOutput(ctx context.Context) pulumix.Output[*SyncSet] {
	return pulumix.Output[*SyncSet]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o SyncSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resources represents the SyncSets configuration.
func (o SyncSetOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncSet) pulumi.StringPtrOutput { return v.Resources }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SyncSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SyncSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SyncSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncSetOutput{})
}
