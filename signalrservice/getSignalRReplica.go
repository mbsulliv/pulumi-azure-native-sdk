// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalrservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the replica and its properties.
// Azure REST API version: 2023-03-01-preview.
//
// Other available API versions: 2023-06-01-preview, 2023-08-01-preview.
func LookupSignalRReplica(ctx *pulumi.Context, args *LookupSignalRReplicaArgs, opts ...pulumi.InvokeOption) (*LookupSignalRReplicaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSignalRReplicaResult
	err := ctx.Invoke("azure-native:signalrservice:getSignalRReplica", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRReplicaArgs struct {
	// The name of the replica.
	ReplicaName string `pulumi:"replicaName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A class represent a replica resource.
type LookupSignalRReplicaResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The billing information of the resource.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSignalRReplicaOutput(ctx *pulumi.Context, args LookupSignalRReplicaOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRReplicaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRReplicaResult, error) {
			args := v.(LookupSignalRReplicaArgs)
			r, err := LookupSignalRReplica(ctx, &args, opts...)
			var s LookupSignalRReplicaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRReplicaResultOutput)
}

type LookupSignalRReplicaOutputArgs struct {
	// The name of the replica.
	ReplicaName pulumi.StringInput `pulumi:"replicaName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalRReplicaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRReplicaArgs)(nil)).Elem()
}

// A class represent a replica resource.
type LookupSignalRReplicaResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRReplicaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRReplicaResult)(nil)).Elem()
}

func (o LookupSignalRReplicaResultOutput) ToLookupSignalRReplicaResultOutput() LookupSignalRReplicaResultOutput {
	return o
}

func (o LookupSignalRReplicaResultOutput) ToLookupSignalRReplicaResultOutputWithContext(ctx context.Context) LookupSignalRReplicaResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSignalRReplicaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRReplicaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSignalRReplicaResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRReplicaResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSignalRReplicaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRReplicaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupSignalRReplicaResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRReplicaResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The billing information of the resource.
func (o LookupSignalRReplicaResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRReplicaResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSignalRReplicaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRReplicaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSignalRReplicaResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSignalRReplicaResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSignalRReplicaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRReplicaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRReplicaResultOutput{})
}
