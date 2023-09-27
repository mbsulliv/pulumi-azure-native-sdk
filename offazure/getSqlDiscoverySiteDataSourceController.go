// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a SqlDiscoverySiteDataSource
// Azure REST API version: 2023-06-06.
func LookupSqlDiscoverySiteDataSourceController(ctx *pulumi.Context, args *LookupSqlDiscoverySiteDataSourceControllerArgs, opts ...pulumi.InvokeOption) (*LookupSqlDiscoverySiteDataSourceControllerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlDiscoverySiteDataSourceControllerResult
	err := ctx.Invoke("azure-native:offazure:getSqlDiscoverySiteDataSourceController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDiscoverySiteDataSourceControllerArgs struct {
	// SQL Discovery site data source name.
	DiscoverySiteDataSourceName string `pulumi:"discoverySiteDataSourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
	// SQL site name.
	SqlSiteName string `pulumi:"sqlSiteName"`
}

// A SQL discovery site data source resource.
type LookupSqlDiscoverySiteDataSourceControllerResult struct {
	// Gets or sets the discovery site Id.
	DiscoverySiteId *string `pulumi:"discoverySiteId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// provisioning state enum
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSqlDiscoverySiteDataSourceControllerOutput(ctx *pulumi.Context, args LookupSqlDiscoverySiteDataSourceControllerOutputArgs, opts ...pulumi.InvokeOption) LookupSqlDiscoverySiteDataSourceControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlDiscoverySiteDataSourceControllerResult, error) {
			args := v.(LookupSqlDiscoverySiteDataSourceControllerArgs)
			r, err := LookupSqlDiscoverySiteDataSourceController(ctx, &args, opts...)
			var s LookupSqlDiscoverySiteDataSourceControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlDiscoverySiteDataSourceControllerResultOutput)
}

type LookupSqlDiscoverySiteDataSourceControllerOutputArgs struct {
	// SQL Discovery site data source name.
	DiscoverySiteDataSourceName pulumi.StringInput `pulumi:"discoverySiteDataSourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
	// SQL site name.
	SqlSiteName pulumi.StringInput `pulumi:"sqlSiteName"`
}

func (LookupSqlDiscoverySiteDataSourceControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDiscoverySiteDataSourceControllerArgs)(nil)).Elem()
}

// A SQL discovery site data source resource.
type LookupSqlDiscoverySiteDataSourceControllerResultOutput struct{ *pulumi.OutputState }

func (LookupSqlDiscoverySiteDataSourceControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDiscoverySiteDataSourceControllerResult)(nil)).Elem()
}

func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) ToLookupSqlDiscoverySiteDataSourceControllerResultOutput() LookupSqlDiscoverySiteDataSourceControllerResultOutput {
	return o
}

func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) ToLookupSqlDiscoverySiteDataSourceControllerResultOutputWithContext(ctx context.Context) LookupSqlDiscoverySiteDataSourceControllerResultOutput {
	return o
}

func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlDiscoverySiteDataSourceControllerResult] {
	return pulumix.Output[LookupSqlDiscoverySiteDataSourceControllerResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the discovery site Id.
func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlDiscoverySiteDataSourceControllerResult) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDiscoverySiteDataSourceControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDiscoverySiteDataSourceControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// provisioning state enum
func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDiscoverySiteDataSourceControllerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlDiscoverySiteDataSourceControllerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlDiscoverySiteDataSourceControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDiscoverySiteDataSourceControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlDiscoverySiteDataSourceControllerResultOutput{})
}
