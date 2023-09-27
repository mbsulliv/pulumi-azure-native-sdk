// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the detailed information about a linked server of a redis cache (requires Premium SKU).
func LookupLinkedServer(ctx *pulumi.Context, args *LookupLinkedServerArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLinkedServerResult
	err := ctx.Invoke("azure-native:cache/v20230401:getLinkedServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServerArgs struct {
	// The name of the linked server.
	LinkedServerName string `pulumi:"linkedServerName"`
	// The name of the redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response to put/get linked server (with properties) for Redis cache.
type LookupLinkedServerResult struct {
	// The unchanging DNS name which will always point to current geo-primary cache among the linked redis caches for seamless Geo Failover experience.
	GeoReplicatedPrimaryHostName string `pulumi:"geoReplicatedPrimaryHostName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId string `pulumi:"linkedRedisCacheId"`
	// Location of the linked redis cache.
	LinkedRedisCacheLocation string `pulumi:"linkedRedisCacheLocation"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The changing DNS name that resolves to the current geo-primary cache among the linked redis caches before or after the Geo Failover.
	PrimaryHostName string `pulumi:"primaryHostName"`
	// Terminal state of the link between primary and secondary redis cache.
	ProvisioningState string `pulumi:"provisioningState"`
	// Role of the linked server.
	ServerRole string `pulumi:"serverRole"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupLinkedServerOutput(ctx *pulumi.Context, args LookupLinkedServerOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedServerResult, error) {
			args := v.(LookupLinkedServerArgs)
			r, err := LookupLinkedServer(ctx, &args, opts...)
			var s LookupLinkedServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedServerResultOutput)
}

type LookupLinkedServerOutputArgs struct {
	// The name of the linked server.
	LinkedServerName pulumi.StringInput `pulumi:"linkedServerName"`
	// The name of the redis cache.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLinkedServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServerArgs)(nil)).Elem()
}

// Response to put/get linked server (with properties) for Redis cache.
type LookupLinkedServerResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServerResult)(nil)).Elem()
}

func (o LookupLinkedServerResultOutput) ToLookupLinkedServerResultOutput() LookupLinkedServerResultOutput {
	return o
}

func (o LookupLinkedServerResultOutput) ToLookupLinkedServerResultOutputWithContext(ctx context.Context) LookupLinkedServerResultOutput {
	return o
}

func (o LookupLinkedServerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLinkedServerResult] {
	return pulumix.Output[LookupLinkedServerResult]{
		OutputState: o.OutputState,
	}
}

// The unchanging DNS name which will always point to current geo-primary cache among the linked redis caches for seamless Geo Failover experience.
func (o LookupLinkedServerResultOutput) GeoReplicatedPrimaryHostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.GeoReplicatedPrimaryHostName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLinkedServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Fully qualified resourceId of the linked redis cache.
func (o LookupLinkedServerResultOutput) LinkedRedisCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.LinkedRedisCacheId }).(pulumi.StringOutput)
}

// Location of the linked redis cache.
func (o LookupLinkedServerResultOutput) LinkedRedisCacheLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.LinkedRedisCacheLocation }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLinkedServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The changing DNS name that resolves to the current geo-primary cache among the linked redis caches before or after the Geo Failover.
func (o LookupLinkedServerResultOutput) PrimaryHostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.PrimaryHostName }).(pulumi.StringOutput)
}

// Terminal state of the link between primary and secondary redis cache.
func (o LookupLinkedServerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Role of the linked server.
func (o LookupLinkedServerResultOutput) ServerRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.ServerRole }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLinkedServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedServerResultOutput{})
}
