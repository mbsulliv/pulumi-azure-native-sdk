// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get specified scope connection created by this Network Manager.
func LookupScopeConnection(ctx *pulumi.Context, args *LookupScopeConnectionArgs, opts ...pulumi.InvokeOption) (*LookupScopeConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupScopeConnectionResult
	err := ctx.Invoke("azure-native:network/v20230601:getScopeConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeConnectionArgs struct {
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name for the cross-tenant connection.
	ScopeConnectionName string `pulumi:"scopeConnectionName"`
}

// The Scope Connections resource
type LookupScopeConnectionResult struct {
	// A description of the scope connection.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource ID.
	ResourceId *string `pulumi:"resourceId"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tenant ID.
	TenantId *string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupScopeConnectionOutput(ctx *pulumi.Context, args LookupScopeConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupScopeConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScopeConnectionResult, error) {
			args := v.(LookupScopeConnectionArgs)
			r, err := LookupScopeConnection(ctx, &args, opts...)
			var s LookupScopeConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScopeConnectionResultOutput)
}

type LookupScopeConnectionOutputArgs struct {
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name for the cross-tenant connection.
	ScopeConnectionName pulumi.StringInput `pulumi:"scopeConnectionName"`
}

func (LookupScopeConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeConnectionArgs)(nil)).Elem()
}

// The Scope Connections resource
type LookupScopeConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupScopeConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeConnectionResult)(nil)).Elem()
}

func (o LookupScopeConnectionResultOutput) ToLookupScopeConnectionResultOutput() LookupScopeConnectionResultOutput {
	return o
}

func (o LookupScopeConnectionResultOutput) ToLookupScopeConnectionResultOutputWithContext(ctx context.Context) LookupScopeConnectionResultOutput {
	return o
}

// A description of the scope connection.
func (o LookupScopeConnectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupScopeConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupScopeConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupScopeConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupScopeConnectionResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The system metadata related to this resource.
func (o LookupScopeConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tenant ID.
func (o LookupScopeConnectionResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupScopeConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScopeConnectionResultOutput{})
}
