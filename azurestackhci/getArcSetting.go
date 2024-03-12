// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get ArcSetting resource details of HCI Cluster.
// Azure REST API version: 2023-03-01.
//
// Other available API versions: 2021-09-01-preview, 2022-12-15-preview, 2023-06-01, 2023-08-01, 2023-08-01-preview, 2023-11-01-preview, 2024-01-01.
func LookupArcSetting(ctx *pulumi.Context, args *LookupArcSettingArgs, opts ...pulumi.InvokeOption) (*LookupArcSettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupArcSettingResult
	err := ctx.Invoke("azure-native:azurestackhci:getArcSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArcSettingArgs struct {
	// The name of the proxy resource holding details of HCI ArcSetting information.
	ArcSettingName string `pulumi:"arcSettingName"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ArcSetting details.
type LookupArcSettingResult struct {
	// Aggregate state of Arc agent across the nodes in this HCI cluster.
	AggregateState string `pulumi:"aggregateState"`
	// App id of arc AAD identity.
	ArcApplicationClientId *string `pulumi:"arcApplicationClientId"`
	// Object id of arc AAD identity.
	ArcApplicationObjectId *string `pulumi:"arcApplicationObjectId"`
	// Tenant id of arc AAD identity.
	ArcApplicationTenantId *string `pulumi:"arcApplicationTenantId"`
	// The resource group that hosts the Arc agents, ie. Hybrid Compute Machine resources.
	ArcInstanceResourceGroup *string `pulumi:"arcInstanceResourceGroup"`
	// Object id of arc AAD service principal.
	ArcServicePrincipalObjectId *string `pulumi:"arcServicePrincipalObjectId"`
	// contains connectivity related configuration for ARC resources
	ConnectivityProperties []ArcConnectivityPropertiesResponse `pulumi:"connectivityProperties"`
	// Properties for each of the default extensions category
	DefaultExtensions []DefaultExtensionDetailsResponse `pulumi:"defaultExtensions"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// State of Arc agent in each of the nodes.
	PerNodeDetails []PerNodeStateResponse `pulumi:"perNodeDetails"`
	// Provisioning state of the ArcSetting proxy resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupArcSettingOutput(ctx *pulumi.Context, args LookupArcSettingOutputArgs, opts ...pulumi.InvokeOption) LookupArcSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArcSettingResult, error) {
			args := v.(LookupArcSettingArgs)
			r, err := LookupArcSetting(ctx, &args, opts...)
			var s LookupArcSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArcSettingResultOutput)
}

type LookupArcSettingOutputArgs struct {
	// The name of the proxy resource holding details of HCI ArcSetting information.
	ArcSettingName pulumi.StringInput `pulumi:"arcSettingName"`
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupArcSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArcSettingArgs)(nil)).Elem()
}

// ArcSetting details.
type LookupArcSettingResultOutput struct{ *pulumi.OutputState }

func (LookupArcSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArcSettingResult)(nil)).Elem()
}

func (o LookupArcSettingResultOutput) ToLookupArcSettingResultOutput() LookupArcSettingResultOutput {
	return o
}

func (o LookupArcSettingResultOutput) ToLookupArcSettingResultOutputWithContext(ctx context.Context) LookupArcSettingResultOutput {
	return o
}

// Aggregate state of Arc agent across the nodes in this HCI cluster.
func (o LookupArcSettingResultOutput) AggregateState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.AggregateState }).(pulumi.StringOutput)
}

// App id of arc AAD identity.
func (o LookupArcSettingResultOutput) ArcApplicationClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcApplicationClientId }).(pulumi.StringPtrOutput)
}

// Object id of arc AAD identity.
func (o LookupArcSettingResultOutput) ArcApplicationObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcApplicationObjectId }).(pulumi.StringPtrOutput)
}

// Tenant id of arc AAD identity.
func (o LookupArcSettingResultOutput) ArcApplicationTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcApplicationTenantId }).(pulumi.StringPtrOutput)
}

// The resource group that hosts the Arc agents, ie. Hybrid Compute Machine resources.
func (o LookupArcSettingResultOutput) ArcInstanceResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcInstanceResourceGroup }).(pulumi.StringPtrOutput)
}

// Object id of arc AAD service principal.
func (o LookupArcSettingResultOutput) ArcServicePrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcServicePrincipalObjectId }).(pulumi.StringPtrOutput)
}

// contains connectivity related configuration for ARC resources
func (o LookupArcSettingResultOutput) ConnectivityProperties() ArcConnectivityPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupArcSettingResult) []ArcConnectivityPropertiesResponse { return v.ConnectivityProperties }).(ArcConnectivityPropertiesResponseArrayOutput)
}

// Properties for each of the default extensions category
func (o LookupArcSettingResultOutput) DefaultExtensions() DefaultExtensionDetailsResponseArrayOutput {
	return o.ApplyT(func(v LookupArcSettingResult) []DefaultExtensionDetailsResponse { return v.DefaultExtensions }).(DefaultExtensionDetailsResponseArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupArcSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupArcSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of Arc agent in each of the nodes.
func (o LookupArcSettingResultOutput) PerNodeDetails() PerNodeStateResponseArrayOutput {
	return o.ApplyT(func(v LookupArcSettingResult) []PerNodeStateResponse { return v.PerNodeDetails }).(PerNodeStateResponseArrayOutput)
}

// Provisioning state of the ArcSetting proxy resource.
func (o LookupArcSettingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupArcSettingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupArcSettingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupArcSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArcSettingResultOutput{})
}
