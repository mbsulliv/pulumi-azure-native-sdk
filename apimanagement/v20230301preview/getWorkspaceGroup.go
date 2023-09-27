// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the group specified by its identifier.
func LookupWorkspaceGroup(ctx *pulumi.Context, args *LookupWorkspaceGroupArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceGroupResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getWorkspaceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceGroupArgs struct {
	// Group identifier. Must be unique in the current API Management service instance.
	GroupId string `pulumi:"groupId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Contract details.
type LookupWorkspaceGroupResult struct {
	// true if the group is one of the three system groups (Administrators, Developers, or Guests); otherwise false.
	BuiltIn bool `pulumi:"builtIn"`
	// Group description. Can contain HTML formatting tags.
	Description *string `pulumi:"description"`
	// Group name.
	DisplayName string `pulumi:"displayName"`
	// For external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
	ExternalId *string `pulumi:"externalId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceGroupOutput(ctx *pulumi.Context, args LookupWorkspaceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceGroupResult, error) {
			args := v.(LookupWorkspaceGroupArgs)
			r, err := LookupWorkspaceGroup(ctx, &args, opts...)
			var s LookupWorkspaceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceGroupResultOutput)
}

type LookupWorkspaceGroupOutputArgs struct {
	// Group identifier. Must be unique in the current API Management service instance.
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceGroupArgs)(nil)).Elem()
}

// Contract details.
type LookupWorkspaceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceGroupResult)(nil)).Elem()
}

func (o LookupWorkspaceGroupResultOutput) ToLookupWorkspaceGroupResultOutput() LookupWorkspaceGroupResultOutput {
	return o
}

func (o LookupWorkspaceGroupResultOutput) ToLookupWorkspaceGroupResultOutputWithContext(ctx context.Context) LookupWorkspaceGroupResultOutput {
	return o
}

func (o LookupWorkspaceGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceGroupResult] {
	return pulumix.Output[LookupWorkspaceGroupResult]{
		OutputState: o.OutputState,
	}
}

// true if the group is one of the three system groups (Administrators, Developers, or Guests); otherwise false.
func (o LookupWorkspaceGroupResultOutput) BuiltIn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkspaceGroupResult) bool { return v.BuiltIn }).(pulumi.BoolOutput)
}

// Group description. Can contain HTML formatting tags.
func (o LookupWorkspaceGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Group name.
func (o LookupWorkspaceGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// For external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
func (o LookupWorkspaceGroupResultOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceGroupResult) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceGroupResultOutput{})
}
