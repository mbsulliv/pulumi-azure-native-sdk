// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation returns properties of a SyncSet.
func LookupSyncSet(ctx *pulumi.Context, args *LookupSyncSetArgs, opts ...pulumi.InvokeOption) (*LookupSyncSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSyncSetResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20230701preview:getSyncSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncSetArgs struct {
	// The name of the SyncSet resource.
	ChildResourceName string `pulumi:"childResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// SyncSet represents a SyncSet for an Azure Red Hat OpenShift Cluster.
type LookupSyncSetResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Resources represents the SyncSets configuration.
	Resources *string `pulumi:"resources"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSyncSetOutput(ctx *pulumi.Context, args LookupSyncSetOutputArgs, opts ...pulumi.InvokeOption) LookupSyncSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSyncSetResult, error) {
			args := v.(LookupSyncSetArgs)
			r, err := LookupSyncSet(ctx, &args, opts...)
			var s LookupSyncSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSyncSetResultOutput)
}

type LookupSyncSetOutputArgs struct {
	// The name of the SyncSet resource.
	ChildResourceName pulumi.StringInput `pulumi:"childResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSyncSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncSetArgs)(nil)).Elem()
}

// SyncSet represents a SyncSet for an Azure Red Hat OpenShift Cluster.
type LookupSyncSetResultOutput struct{ *pulumi.OutputState }

func (LookupSyncSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncSetResult)(nil)).Elem()
}

func (o LookupSyncSetResultOutput) ToLookupSyncSetResultOutput() LookupSyncSetResultOutput {
	return o
}

func (o LookupSyncSetResultOutput) ToLookupSyncSetResultOutputWithContext(ctx context.Context) LookupSyncSetResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSyncSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSyncSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resources represents the SyncSets configuration.
func (o LookupSyncSetResultOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncSetResult) *string { return v.Resources }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSyncSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSyncSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSyncSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncSetResultOutput{})
}
