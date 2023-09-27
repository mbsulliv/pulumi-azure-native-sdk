// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a catalog
func LookupCatalog(ctx *pulumi.Context, args *LookupCatalogArgs, opts ...pulumi.InvokeOption) (*LookupCatalogResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCatalogResult
	err := ctx.Invoke("azure-native:devcenter/v20230801preview:getCatalog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCatalogArgs struct {
	// The name of the Catalog.
	CatalogName string `pulumi:"catalogName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a catalog.
type LookupCatalogResult struct {
	// Properties for an Azure DevOps catalog type.
	AdoGit *GitCatalogResponse `pulumi:"adoGit"`
	// The connection state of the catalog.
	ConnectionState string `pulumi:"connectionState"`
	// Properties for a GitHub catalog type.
	GitHub *GitCatalogResponse `pulumi:"gitHub"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// When the catalog was last connected.
	LastConnectionTime string `pulumi:"lastConnectionTime"`
	// Stats of the latest synchronization.
	LastSyncStats SyncStatsResponse `pulumi:"lastSyncStats"`
	// When the catalog was last synced.
	LastSyncTime string `pulumi:"lastSyncTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The synchronization state of the catalog.
	SyncState string `pulumi:"syncState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCatalogOutput(ctx *pulumi.Context, args LookupCatalogOutputArgs, opts ...pulumi.InvokeOption) LookupCatalogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCatalogResult, error) {
			args := v.(LookupCatalogArgs)
			r, err := LookupCatalog(ctx, &args, opts...)
			var s LookupCatalogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCatalogResultOutput)
}

type LookupCatalogOutputArgs struct {
	// The name of the Catalog.
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// The name of the devcenter.
	DevCenterName pulumi.StringInput `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCatalogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogArgs)(nil)).Elem()
}

// Represents a catalog.
type LookupCatalogResultOutput struct{ *pulumi.OutputState }

func (LookupCatalogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogResult)(nil)).Elem()
}

func (o LookupCatalogResultOutput) ToLookupCatalogResultOutput() LookupCatalogResultOutput {
	return o
}

func (o LookupCatalogResultOutput) ToLookupCatalogResultOutputWithContext(ctx context.Context) LookupCatalogResultOutput {
	return o
}

func (o LookupCatalogResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCatalogResult] {
	return pulumix.Output[LookupCatalogResult]{
		OutputState: o.OutputState,
	}
}

// Properties for an Azure DevOps catalog type.
func (o LookupCatalogResultOutput) AdoGit() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v LookupCatalogResult) *GitCatalogResponse { return v.AdoGit }).(GitCatalogResponsePtrOutput)
}

// The connection state of the catalog.
func (o LookupCatalogResultOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.ConnectionState }).(pulumi.StringOutput)
}

// Properties for a GitHub catalog type.
func (o LookupCatalogResultOutput) GitHub() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v LookupCatalogResult) *GitCatalogResponse { return v.GitHub }).(GitCatalogResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCatalogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.Id }).(pulumi.StringOutput)
}

// When the catalog was last connected.
func (o LookupCatalogResultOutput) LastConnectionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.LastConnectionTime }).(pulumi.StringOutput)
}

// Stats of the latest synchronization.
func (o LookupCatalogResultOutput) LastSyncStats() SyncStatsResponseOutput {
	return o.ApplyT(func(v LookupCatalogResult) SyncStatsResponse { return v.LastSyncStats }).(SyncStatsResponseOutput)
}

// When the catalog was last synced.
func (o LookupCatalogResultOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCatalogResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupCatalogResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The synchronization state of the catalog.
func (o LookupCatalogResultOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.SyncState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCatalogResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCatalogResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCatalogResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCatalogResultOutput{})
}
