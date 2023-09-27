// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the details of the management group.
func LookupManagementGroup(ctx *pulumi.Context, args *LookupManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagementGroupResult
	err := ctx.Invoke("azure-native:management/v20210401:getManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupArgs struct {
	// The $expand=children query string parameter allows clients to request inclusion of children in the response payload.  $expand=path includes the path from the root group to the current group.  $expand=ancestors includes the ancestor Ids of the current group.
	Expand *string `pulumi:"expand"`
	// A filter which allows the exclusion of subscriptions from results (i.e. '$filter=children.childType ne Subscription')
	Filter *string `pulumi:"filter"`
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
	// The $recurse=true query string parameter allows clients to request inclusion of entire hierarchy in the response payload. Note that  $expand=children must be passed up if $recurse is set to true.
	Recurse *bool `pulumi:"recurse"`
}

// The management group details.
type LookupManagementGroupResult struct {
	// The list of children.
	Children []ManagementGroupChildInfoResponse `pulumi:"children"`
	// The details of a management group.
	Details *ManagementGroupDetailsResponse `pulumi:"details"`
	// The friendly name of the management group.
	DisplayName *string `pulumi:"displayName"`
	// The fully qualified ID for the management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id string `pulumi:"id"`
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name string `pulumi:"name"`
	// The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups
	Type string `pulumi:"type"`
}

func LookupManagementGroupOutput(ctx *pulumi.Context, args LookupManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupResult, error) {
			args := v.(LookupManagementGroupArgs)
			r, err := LookupManagementGroup(ctx, &args, opts...)
			var s LookupManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementGroupResultOutput)
}

type LookupManagementGroupOutputArgs struct {
	// The $expand=children query string parameter allows clients to request inclusion of children in the response payload.  $expand=path includes the path from the root group to the current group.  $expand=ancestors includes the ancestor Ids of the current group.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// A filter which allows the exclusion of subscriptions from results (i.e. '$filter=children.childType ne Subscription')
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Management Group ID.
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// The $recurse=true query string parameter allows clients to request inclusion of entire hierarchy in the response payload. Note that  $expand=children must be passed up if $recurse is set to true.
	Recurse pulumi.BoolPtrInput `pulumi:"recurse"`
}

func (LookupManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupArgs)(nil)).Elem()
}

// The management group details.
type LookupManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupResult)(nil)).Elem()
}

func (o LookupManagementGroupResultOutput) ToLookupManagementGroupResultOutput() LookupManagementGroupResultOutput {
	return o
}

func (o LookupManagementGroupResultOutput) ToLookupManagementGroupResultOutputWithContext(ctx context.Context) LookupManagementGroupResultOutput {
	return o
}

func (o LookupManagementGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagementGroupResult] {
	return pulumix.Output[LookupManagementGroupResult]{
		OutputState: o.OutputState,
	}
}

// The list of children.
func (o LookupManagementGroupResultOutput) Children() ManagementGroupChildInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) []ManagementGroupChildInfoResponse { return v.Children }).(ManagementGroupChildInfoResponseArrayOutput)
}

// The details of a management group.
func (o LookupManagementGroupResultOutput) Details() ManagementGroupDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) *ManagementGroupDetailsResponse { return v.Details }).(ManagementGroupDetailsResponsePtrOutput)
}

// The friendly name of the management group.
func (o LookupManagementGroupResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The fully qualified ID for the management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o LookupManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
func (o LookupManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
func (o LookupManagementGroupResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource.  For example, Microsoft.Management/managementGroups
func (o LookupManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupResultOutput{})
}
