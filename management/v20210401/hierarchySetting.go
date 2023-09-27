// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Settings defined at the Management Group scope.
type HierarchySetting struct {
	pulumi.CustomResourceState

	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup pulumi.StringPtrOutput `pulumi:"defaultManagementGroup"`
	// The name of the object. In this case, default.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation pulumi.BoolPtrOutput `pulumi:"requireAuthorizationForGroupCreation"`
	// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/settings.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHierarchySetting registers a new resource with the given unique name, arguments, and options.
func NewHierarchySetting(ctx *pulumi.Context,
	name string, args *HierarchySettingArgs, opts ...pulumi.ResourceOption) (*HierarchySetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:management:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-native:management/v20200201:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-native:management/v20200501:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-native:management/v20201001:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-native:management/v20230401:HierarchySetting"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HierarchySetting
	err := ctx.RegisterResource("azure-native:management/v20210401:HierarchySetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHierarchySetting gets an existing HierarchySetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHierarchySetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HierarchySettingState, opts ...pulumi.ResourceOption) (*HierarchySetting, error) {
	var resource HierarchySetting
	err := ctx.ReadResource("azure-native:management/v20210401:HierarchySetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HierarchySetting resources.
type hierarchySettingState struct {
}

type HierarchySettingState struct {
}

func (HierarchySettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*hierarchySettingState)(nil)).Elem()
}

type hierarchySettingArgs struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup *string `pulumi:"defaultManagementGroup"`
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation *bool `pulumi:"requireAuthorizationForGroupCreation"`
}

// The set of arguments for constructing a HierarchySetting resource.
type HierarchySettingArgs struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup pulumi.StringPtrInput
	// Management Group ID.
	GroupId pulumi.StringInput
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation pulumi.BoolPtrInput
}

func (HierarchySettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hierarchySettingArgs)(nil)).Elem()
}

type HierarchySettingInput interface {
	pulumi.Input

	ToHierarchySettingOutput() HierarchySettingOutput
	ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput
}

func (*HierarchySetting) ElementType() reflect.Type {
	return reflect.TypeOf((**HierarchySetting)(nil)).Elem()
}

func (i *HierarchySetting) ToHierarchySettingOutput() HierarchySettingOutput {
	return i.ToHierarchySettingOutputWithContext(context.Background())
}

func (i *HierarchySetting) ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchySettingOutput)
}

func (i *HierarchySetting) ToOutput(ctx context.Context) pulumix.Output[*HierarchySetting] {
	return pulumix.Output[*HierarchySetting]{
		OutputState: i.ToHierarchySettingOutputWithContext(ctx).OutputState,
	}
}

type HierarchySettingOutput struct{ *pulumi.OutputState }

func (HierarchySettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HierarchySetting)(nil)).Elem()
}

func (o HierarchySettingOutput) ToHierarchySettingOutput() HierarchySettingOutput {
	return o
}

func (o HierarchySettingOutput) ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput {
	return o
}

func (o HierarchySettingOutput) ToOutput(ctx context.Context) pulumix.Output[*HierarchySetting] {
	return pulumix.Output[*HierarchySetting]{
		OutputState: o.OutputState,
	}
}

// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
func (o HierarchySettingOutput) DefaultManagementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchySetting) pulumi.StringPtrOutput { return v.DefaultManagementGroup }).(pulumi.StringPtrOutput)
}

// The name of the object. In this case, default.
func (o HierarchySettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HierarchySetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
func (o HierarchySettingOutput) RequireAuthorizationForGroupCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HierarchySetting) pulumi.BoolPtrOutput { return v.RequireAuthorizationForGroupCreation }).(pulumi.BoolPtrOutput)
}

// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
func (o HierarchySettingOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HierarchySetting) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource.  For example, Microsoft.Management/managementGroups/settings.
func (o HierarchySettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HierarchySetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HierarchySettingOutput{})
}
