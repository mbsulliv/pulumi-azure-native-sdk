// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scom

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the details of the monitored resource.
// Azure REST API version: 2023-07-07-preview.
func LookupMonitoredResource(ctx *pulumi.Context, args *LookupMonitoredResourceArgs, opts ...pulumi.InvokeOption) (*LookupMonitoredResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMonitoredResourceResult
	err := ctx.Invoke("azure-native:scom:getMonitoredResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoredResourceArgs struct {
	// Name of the SCOM managed instance.
	InstanceName string `pulumi:"instanceName"`
	// The monitored resource name.
	MonitoredResourceName string `pulumi:"monitoredResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A monitored resource.
type LookupMonitoredResourceResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The properties of a monitored resource.
	Properties MonitoredResourcePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMonitoredResourceOutput(ctx *pulumi.Context, args LookupMonitoredResourceOutputArgs, opts ...pulumi.InvokeOption) LookupMonitoredResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitoredResourceResult, error) {
			args := v.(LookupMonitoredResourceArgs)
			r, err := LookupMonitoredResource(ctx, &args, opts...)
			var s LookupMonitoredResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitoredResourceResultOutput)
}

type LookupMonitoredResourceOutputArgs struct {
	// Name of the SCOM managed instance.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// The monitored resource name.
	MonitoredResourceName pulumi.StringInput `pulumi:"monitoredResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMonitoredResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoredResourceArgs)(nil)).Elem()
}

// A monitored resource.
type LookupMonitoredResourceResultOutput struct{ *pulumi.OutputState }

func (LookupMonitoredResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoredResourceResult)(nil)).Elem()
}

func (o LookupMonitoredResourceResultOutput) ToLookupMonitoredResourceResultOutput() LookupMonitoredResourceResultOutput {
	return o
}

func (o LookupMonitoredResourceResultOutput) ToLookupMonitoredResourceResultOutputWithContext(ctx context.Context) LookupMonitoredResourceResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupMonitoredResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoredResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMonitoredResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoredResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of a monitored resource.
func (o LookupMonitoredResourceResultOutput) Properties() MonitoredResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupMonitoredResourceResult) MonitoredResourcePropertiesResponse { return v.Properties }).(MonitoredResourcePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMonitoredResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMonitoredResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMonitoredResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoredResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitoredResourceResultOutput{})
}
