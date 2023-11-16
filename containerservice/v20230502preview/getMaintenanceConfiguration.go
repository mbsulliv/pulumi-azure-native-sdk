// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230502preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// See [planned maintenance](https://docs.microsoft.com/azure/aks/planned-maintenance) for more information about planned maintenance.
func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:containerservice/v20230502preview:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupMaintenanceConfigurationArgs struct {
	// The name of the maintenance configuration.
	ConfigName string `pulumi:"configName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// See [planned maintenance](https://docs.microsoft.com/azure/aks/planned-maintenance) for more information about planned maintenance.
type LookupMaintenanceConfigurationResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Maintenance window for the maintenance configuration.
	MaintenanceWindow *MaintenanceWindowResponse `pulumi:"maintenanceWindow"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `pulumi:"name"`
	// Time slots on which upgrade is not allowed.
	NotAllowedTime []TimeSpanResponse `pulumi:"notAllowedTime"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// If two array entries specify the same day of the week, the applied configuration is the union of times in both entries.
	TimeInWeek []TimeInWeekResponse `pulumi:"timeInWeek"`
	// Resource type
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupMaintenanceConfigurationResult
func (val *LookupMaintenanceConfigurationResult) Defaults() *LookupMaintenanceConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.MaintenanceWindow = tmp.MaintenanceWindow.Defaults()

	return &tmp
}

func LookupMaintenanceConfigurationOutput(ctx *pulumi.Context, args LookupMaintenanceConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupMaintenanceConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMaintenanceConfigurationResult, error) {
			args := v.(LookupMaintenanceConfigurationArgs)
			r, err := LookupMaintenanceConfiguration(ctx, &args, opts...)
			var s LookupMaintenanceConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMaintenanceConfigurationResultOutput)
}

type LookupMaintenanceConfigurationOutputArgs struct {
	// The name of the maintenance configuration.
	ConfigName pulumi.StringInput `pulumi:"configName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupMaintenanceConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceConfigurationArgs)(nil)).Elem()
}

// See [planned maintenance](https://docs.microsoft.com/azure/aks/planned-maintenance) for more information about planned maintenance.
type LookupMaintenanceConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupMaintenanceConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceConfigurationResult)(nil)).Elem()
}

func (o LookupMaintenanceConfigurationResultOutput) ToLookupMaintenanceConfigurationResultOutput() LookupMaintenanceConfigurationResultOutput {
	return o
}

func (o LookupMaintenanceConfigurationResultOutput) ToLookupMaintenanceConfigurationResultOutputWithContext(ctx context.Context) LookupMaintenanceConfigurationResultOutput {
	return o
}

// Resource ID.
func (o LookupMaintenanceConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Maintenance window for the maintenance configuration.
func (o LookupMaintenanceConfigurationResultOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *MaintenanceWindowResponse { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupMaintenanceConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Time slots on which upgrade is not allowed.
func (o LookupMaintenanceConfigurationResultOutput) NotAllowedTime() TimeSpanResponseArrayOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) []TimeSpanResponse { return v.NotAllowedTime }).(TimeSpanResponseArrayOutput)
}

// The system metadata relating to this resource.
func (o LookupMaintenanceConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// If two array entries specify the same day of the week, the applied configuration is the union of times in both entries.
func (o LookupMaintenanceConfigurationResultOutput) TimeInWeek() TimeInWeekResponseArrayOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) []TimeInWeekResponse { return v.TimeInWeek }).(TimeInWeekResponseArrayOutput)
}

// Resource type
func (o LookupMaintenanceConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMaintenanceConfigurationResultOutput{})
}
