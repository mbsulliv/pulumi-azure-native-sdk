// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200207preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of a SAP monitor for the specified subscription, resource group, and resource name.
func LookupSapMonitor(ctx *pulumi.Context, args *LookupSapMonitorArgs, opts ...pulumi.InvokeOption) (*LookupSapMonitorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSapMonitorResult
	err := ctx.Invoke("azure-native:hanaonazure/v20200207preview:getSapMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSapMonitorArgs struct {
	// Name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SAP monitor resource.
	SapMonitorName string `pulumi:"sapMonitorName"`
}

// SAP monitor info on Azure (ARM properties and SAP monitor properties)
type LookupSapMonitorResult struct {
	// The value indicating whether to send analytics to Microsoft
	EnableCustomerAnalytics *bool `pulumi:"enableCustomerAnalytics"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The ARM ID of the Log Analytics Workspace that is used for monitoring
	LogAnalyticsWorkspaceArmId *string `pulumi:"logAnalyticsWorkspaceArmId"`
	// The workspace ID of the log analytics workspace to be used for monitoring
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The shared key of the log analytics workspace that is used for monitoring
	LogAnalyticsWorkspaceSharedKey *string `pulumi:"logAnalyticsWorkspaceSharedKey"`
	// The name of the resource group the SAP Monitor resources get deployed into.
	ManagedResourceGroupName string `pulumi:"managedResourceGroupName"`
	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet *string `pulumi:"monitorSubnet"`
	// The name of the resource
	Name string `pulumi:"name"`
	// State of provisioning of the HanaInstance
	ProvisioningState string `pulumi:"provisioningState"`
	// The version of the payload running in the Collector VM
	SapMonitorCollectorVersion string `pulumi:"sapMonitorCollectorVersion"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSapMonitorOutput(ctx *pulumi.Context, args LookupSapMonitorOutputArgs, opts ...pulumi.InvokeOption) LookupSapMonitorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSapMonitorResult, error) {
			args := v.(LookupSapMonitorArgs)
			r, err := LookupSapMonitor(ctx, &args, opts...)
			var s LookupSapMonitorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSapMonitorResultOutput)
}

type LookupSapMonitorOutputArgs struct {
	// Name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the SAP monitor resource.
	SapMonitorName pulumi.StringInput `pulumi:"sapMonitorName"`
}

func (LookupSapMonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSapMonitorArgs)(nil)).Elem()
}

// SAP monitor info on Azure (ARM properties and SAP monitor properties)
type LookupSapMonitorResultOutput struct{ *pulumi.OutputState }

func (LookupSapMonitorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSapMonitorResult)(nil)).Elem()
}

func (o LookupSapMonitorResultOutput) ToLookupSapMonitorResultOutput() LookupSapMonitorResultOutput {
	return o
}

func (o LookupSapMonitorResultOutput) ToLookupSapMonitorResultOutputWithContext(ctx context.Context) LookupSapMonitorResultOutput {
	return o
}

// The value indicating whether to send analytics to Microsoft
func (o LookupSapMonitorResultOutput) EnableCustomerAnalytics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *bool { return v.EnableCustomerAnalytics }).(pulumi.BoolPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSapMonitorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSapMonitorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.Location }).(pulumi.StringOutput)
}

// The ARM ID of the Log Analytics Workspace that is used for monitoring
func (o LookupSapMonitorResultOutput) LogAnalyticsWorkspaceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *string { return v.LogAnalyticsWorkspaceArmId }).(pulumi.StringPtrOutput)
}

// The workspace ID of the log analytics workspace to be used for monitoring
func (o LookupSapMonitorResultOutput) LogAnalyticsWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *string { return v.LogAnalyticsWorkspaceId }).(pulumi.StringPtrOutput)
}

// The shared key of the log analytics workspace that is used for monitoring
func (o LookupSapMonitorResultOutput) LogAnalyticsWorkspaceSharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *string { return v.LogAnalyticsWorkspaceSharedKey }).(pulumi.StringPtrOutput)
}

// The name of the resource group the SAP Monitor resources get deployed into.
func (o LookupSapMonitorResultOutput) ManagedResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.ManagedResourceGroupName }).(pulumi.StringOutput)
}

// The subnet which the SAP monitor will be deployed in
func (o LookupSapMonitorResultOutput) MonitorSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *string { return v.MonitorSubnet }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupSapMonitorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of provisioning of the HanaInstance
func (o LookupSapMonitorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The version of the payload running in the Collector VM
func (o LookupSapMonitorResultOutput) SapMonitorCollectorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.SapMonitorCollectorVersion }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupSapMonitorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSapMonitorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSapMonitorResultOutput{})
}
