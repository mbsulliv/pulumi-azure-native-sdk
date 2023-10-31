// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Maintenance configuration record type
func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:maintenance/v20230901preview:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupMaintenanceConfigurationArgs struct {
	// Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Maintenance Configuration Name
	ResourceName string `pulumi:"resourceName"`
}

// Maintenance configuration record type
type LookupMaintenanceConfigurationResult struct {
	// Duration of the maintenance window in HH:mm format. If not provided, default value will be used based on maintenance scope provided. Example: 05:00.
	Duration *string `pulumi:"duration"`
	// Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone. Expiration date must be set to a future date. If not provided, it will be set to the maximum datetime 9999-12-31 23:59:59.
	ExpirationDateTime *string `pulumi:"expirationDateTime"`
	// Gets or sets extensionProperties of the maintenanceConfiguration
	ExtensionProperties map[string]string `pulumi:"extensionProperties"`
	// Fully qualified identifier of the resource
	Id string `pulumi:"id"`
	// The input parameters to be passed to the patch run operation.
	InstallPatches *InputPatchConfigurationResponse `pulumi:"installPatches"`
	// Gets or sets location of the resource
	Location *string `pulumi:"location"`
	// Gets or sets maintenanceScope of the configuration
	MaintenanceScope *string `pulumi:"maintenanceScope"`
	// Name of the resource
	Name string `pulumi:"name"`
	// Gets or sets namespace of the resource
	Namespace *string `pulumi:"namespace"`
	// Rate at which a Maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules. Daily schedule are formatted as recurEvery: [Frequency as integer]['Day(s)']. If no frequency is provided, the default frequency is 1. Daily schedule examples are recurEvery: Day, recurEvery: 3Days.  Weekly schedule are formatted as recurEvery: [Frequency as integer]['Week(s)'] [Optional comma separated list of weekdays Monday-Sunday]. Weekly schedule examples are recurEvery: 3Weeks, recurEvery: Week Saturday,Sunday. Monthly schedules are formatted as [Frequency as integer]['Month(s)'] [Comma separated list of month days] or [Frequency as integer]['Month(s)'] [Week of Month (First, Second, Third, Fourth, Last)] [Weekday Monday-Sunday] [Optional Offset(No. of days)]. Offset value must be between -6 to 6 inclusive. Monthly schedule examples are recurEvery: Month, recurEvery: 2Months, recurEvery: Month day23,day24, recurEvery: Month Last Sunday, recurEvery: Month Fourth Monday, recurEvery: Month Last Sunday Offset-3, recurEvery: Month Third Sunday Offset6.
	RecurEvery *string `pulumi:"recurEvery"`
	// Effective start date of the maintenance window in YYYY-MM-DD hh:mm format. The start date can be set to either the current date or future date. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone.
	StartDateTime *string `pulumi:"startDateTime"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Gets or sets tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Name of the timezone. List of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell. Example: Pacific Standard Time, UTC, W. Europe Standard Time, Korea Standard Time, Cen. Australia Standard Time.
	TimeZone *string `pulumi:"timeZone"`
	// Type of the resource
	Type string `pulumi:"type"`
	// Gets or sets the visibility of the configuration. The default value is 'Custom'
	Visibility *string `pulumi:"visibility"`
}

// Defaults sets the appropriate defaults for LookupMaintenanceConfigurationResult
func (val *LookupMaintenanceConfigurationResult) Defaults() *LookupMaintenanceConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.InstallPatches = tmp.InstallPatches.Defaults()

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
	// Resource Group Name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Maintenance Configuration Name
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupMaintenanceConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceConfigurationArgs)(nil)).Elem()
}

// Maintenance configuration record type
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

func (o LookupMaintenanceConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMaintenanceConfigurationResult] {
	return pulumix.Output[LookupMaintenanceConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Duration of the maintenance window in HH:mm format. If not provided, default value will be used based on maintenance scope provided. Example: 05:00.
func (o LookupMaintenanceConfigurationResultOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

// Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone. Expiration date must be set to a future date. If not provided, it will be set to the maximum datetime 9999-12-31 23:59:59.
func (o LookupMaintenanceConfigurationResultOutput) ExpirationDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.ExpirationDateTime }).(pulumi.StringPtrOutput)
}

// Gets or sets extensionProperties of the maintenanceConfiguration
func (o LookupMaintenanceConfigurationResultOutput) ExtensionProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) map[string]string { return v.ExtensionProperties }).(pulumi.StringMapOutput)
}

// Fully qualified identifier of the resource
func (o LookupMaintenanceConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The input parameters to be passed to the patch run operation.
func (o LookupMaintenanceConfigurationResultOutput) InstallPatches() InputPatchConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *InputPatchConfigurationResponse { return v.InstallPatches }).(InputPatchConfigurationResponsePtrOutput)
}

// Gets or sets location of the resource
func (o LookupMaintenanceConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets maintenanceScope of the configuration
func (o LookupMaintenanceConfigurationResultOutput) MaintenanceScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.MaintenanceScope }).(pulumi.StringPtrOutput)
}

// Name of the resource
func (o LookupMaintenanceConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets namespace of the resource
func (o LookupMaintenanceConfigurationResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Rate at which a Maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules. Daily schedule are formatted as recurEvery: [Frequency as integer]['Day(s)']. If no frequency is provided, the default frequency is 1. Daily schedule examples are recurEvery: Day, recurEvery: 3Days.  Weekly schedule are formatted as recurEvery: [Frequency as integer]['Week(s)'] [Optional comma separated list of weekdays Monday-Sunday]. Weekly schedule examples are recurEvery: 3Weeks, recurEvery: Week Saturday,Sunday. Monthly schedules are formatted as [Frequency as integer]['Month(s)'] [Comma separated list of month days] or [Frequency as integer]['Month(s)'] [Week of Month (First, Second, Third, Fourth, Last)] [Weekday Monday-Sunday] [Optional Offset(No. of days)]. Offset value must be between -6 to 6 inclusive. Monthly schedule examples are recurEvery: Month, recurEvery: 2Months, recurEvery: Month day23,day24, recurEvery: Month Last Sunday, recurEvery: Month Fourth Monday, recurEvery: Month Last Sunday Offset-3, recurEvery: Month Third Sunday Offset6.
func (o LookupMaintenanceConfigurationResultOutput) RecurEvery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.RecurEvery }).(pulumi.StringPtrOutput)
}

// Effective start date of the maintenance window in YYYY-MM-DD hh:mm format. The start date can be set to either the current date or future date. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone.
func (o LookupMaintenanceConfigurationResultOutput) StartDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.StartDateTime }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMaintenanceConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets tags of the resource
func (o LookupMaintenanceConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Name of the timezone. List of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell. Example: Pacific Standard Time, UTC, W. Europe Standard Time, Korea Standard Time, Cen. Australia Standard Time.
func (o LookupMaintenanceConfigurationResultOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

// Type of the resource
func (o LookupMaintenanceConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the visibility of the configuration. The default value is 'Custom'
func (o LookupMaintenanceConfigurationResultOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMaintenanceConfigurationResultOutput{})
}