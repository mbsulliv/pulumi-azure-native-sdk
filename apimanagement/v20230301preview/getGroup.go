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
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	// Group identifier. Must be unique in the current API Management service instance.
	GroupId string `pulumi:"groupId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Contract details.
type LookupGroupResult struct {
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

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

type LookupGroupOutputArgs struct {
	// Group identifier. Must be unique in the current API Management service instance.
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// Contract details.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGroupResult] {
	return pulumix.Output[LookupGroupResult]{
		OutputState: o.OutputState,
	}
}

// true if the group is one of the three system groups (Administrators, Developers, or Guests); otherwise false.
func (o LookupGroupResultOutput) BuiltIn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.BuiltIn }).(pulumi.BoolOutput)
}

// Group description. Can contain HTML formatting tags.
func (o LookupGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Group name.
func (o LookupGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// For external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
func (o LookupGroupResultOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
