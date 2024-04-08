// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
type GroupQuota struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	Properties GroupQuotasEntityBaseResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGroupQuota registers a new resource with the given unique name, arguments, and options.
func NewGroupQuota(ctx *pulumi.Context,
	name string, args *GroupQuotaArgs, opts ...pulumi.ResourceOption) (*GroupQuota, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:quota:GroupQuota"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GroupQuota
	err := ctx.RegisterResource("azure-native:quota/v20230601preview:GroupQuota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupQuota gets an existing GroupQuota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupQuotaState, opts ...pulumi.ResourceOption) (*GroupQuota, error) {
	var resource GroupQuota
	err := ctx.ReadResource("azure-native:quota/v20230601preview:GroupQuota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupQuota resources.
type groupQuotaState struct {
}

type GroupQuotaState struct {
}

func (GroupQuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupQuotaState)(nil)).Elem()
}

type groupQuotaArgs struct {
	// The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
	GroupQuotaName *string `pulumi:"groupQuotaName"`
	// Management Group Id.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	Properties *GroupQuotasEntityBase `pulumi:"properties"`
}

// The set of arguments for constructing a GroupQuota resource.
type GroupQuotaArgs struct {
	// The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
	GroupQuotaName pulumi.StringPtrInput
	// Management Group Id.
	ManagementGroupId pulumi.StringInput
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	Properties GroupQuotasEntityBasePtrInput
}

func (GroupQuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupQuotaArgs)(nil)).Elem()
}

type GroupQuotaInput interface {
	pulumi.Input

	ToGroupQuotaOutput() GroupQuotaOutput
	ToGroupQuotaOutputWithContext(ctx context.Context) GroupQuotaOutput
}

func (*GroupQuota) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupQuota)(nil)).Elem()
}

func (i *GroupQuota) ToGroupQuotaOutput() GroupQuotaOutput {
	return i.ToGroupQuotaOutputWithContext(context.Background())
}

func (i *GroupQuota) ToGroupQuotaOutputWithContext(ctx context.Context) GroupQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupQuotaOutput)
}

type GroupQuotaOutput struct{ *pulumi.OutputState }

func (GroupQuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupQuota)(nil)).Elem()
}

func (o GroupQuotaOutput) ToGroupQuotaOutput() GroupQuotaOutput {
	return o
}

func (o GroupQuotaOutput) ToGroupQuotaOutputWithContext(ctx context.Context) GroupQuotaOutput {
	return o
}

// The name of the resource
func (o GroupQuotaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupQuota) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
func (o GroupQuotaOutput) Properties() GroupQuotasEntityBaseResponseOutput {
	return o.ApplyT(func(v *GroupQuota) GroupQuotasEntityBaseResponseOutput { return v.Properties }).(GroupQuotasEntityBaseResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GroupQuotaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GroupQuota) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GroupQuotaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupQuota) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupQuotaOutput{})
}
