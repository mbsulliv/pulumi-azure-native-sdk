// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package management

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the hierarchy settings defined at the Management Group level. Settings can only be set on the root Management Group of the hierarchy.
//
// Azure REST API version: 2021-04-01.
//
// Other available API versions: 2023-04-01.
func LookupHierarchySetting(ctx *pulumi.Context, args *LookupHierarchySettingArgs, opts ...pulumi.InvokeOption) (*LookupHierarchySettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHierarchySettingResult
	err := ctx.Invoke("azure-native:management:getHierarchySetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHierarchySettingArgs struct {
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
}

// Settings defined at the Management Group scope.
type LookupHierarchySettingResult struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup *string `pulumi:"defaultManagementGroup"`
	// The fully qualified ID for the settings object.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/settings/default.
	Id string `pulumi:"id"`
	// The name of the object. In this case, default.
	Name string `pulumi:"name"`
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation *bool `pulumi:"requireAuthorizationForGroupCreation"`
	// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/settings.
	Type string `pulumi:"type"`
}

func LookupHierarchySettingOutput(ctx *pulumi.Context, args LookupHierarchySettingOutputArgs, opts ...pulumi.InvokeOption) LookupHierarchySettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHierarchySettingResult, error) {
			args := v.(LookupHierarchySettingArgs)
			r, err := LookupHierarchySetting(ctx, &args, opts...)
			var s LookupHierarchySettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHierarchySettingResultOutput)
}

type LookupHierarchySettingOutputArgs struct {
	// Management Group ID.
	GroupId pulumi.StringInput `pulumi:"groupId"`
}

func (LookupHierarchySettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHierarchySettingArgs)(nil)).Elem()
}

// Settings defined at the Management Group scope.
type LookupHierarchySettingResultOutput struct{ *pulumi.OutputState }

func (LookupHierarchySettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHierarchySettingResult)(nil)).Elem()
}

func (o LookupHierarchySettingResultOutput) ToLookupHierarchySettingResultOutput() LookupHierarchySettingResultOutput {
	return o
}

func (o LookupHierarchySettingResultOutput) ToLookupHierarchySettingResultOutputWithContext(ctx context.Context) LookupHierarchySettingResultOutput {
	return o
}

// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
func (o LookupHierarchySettingResultOutput) DefaultManagementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) *string { return v.DefaultManagementGroup }).(pulumi.StringPtrOutput)
}

// The fully qualified ID for the settings object.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/settings/default.
func (o LookupHierarchySettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the object. In this case, default.
func (o LookupHierarchySettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
func (o LookupHierarchySettingResultOutput) RequireAuthorizationForGroupCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) *bool { return v.RequireAuthorizationForGroupCreation }).(pulumi.BoolPtrOutput)
}

// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
func (o LookupHierarchySettingResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource.  For example, Microsoft.Management/managementGroups/settings.
func (o LookupHierarchySettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHierarchySettingResultOutput{})
}
