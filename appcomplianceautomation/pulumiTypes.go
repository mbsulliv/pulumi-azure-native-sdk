// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appcomplianceautomation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The overview of the compliance result for one report.
type OverviewStatusResponse struct {
	// The count of all failed full automation control.
	FailedCount *int `pulumi:"failedCount"`
	// The count of all manual control.
	ManualCount *int `pulumi:"manualCount"`
	// The count of all passed full automation control.
	PassedCount *int `pulumi:"passedCount"`
}

// The overview of the compliance result for one report.
type OverviewStatusResponseOutput struct{ *pulumi.OutputState }

func (OverviewStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OverviewStatusResponse)(nil)).Elem()
}

func (o OverviewStatusResponseOutput) ToOverviewStatusResponseOutput() OverviewStatusResponseOutput {
	return o
}

func (o OverviewStatusResponseOutput) ToOverviewStatusResponseOutputWithContext(ctx context.Context) OverviewStatusResponseOutput {
	return o
}

// The count of all failed full automation control.
func (o OverviewStatusResponseOutput) FailedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OverviewStatusResponse) *int { return v.FailedCount }).(pulumi.IntPtrOutput)
}

// The count of all manual control.
func (o OverviewStatusResponseOutput) ManualCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OverviewStatusResponse) *int { return v.ManualCount }).(pulumi.IntPtrOutput)
}

// The count of all passed full automation control.
func (o OverviewStatusResponseOutput) PassedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OverviewStatusResponse) *int { return v.PassedCount }).(pulumi.IntPtrOutput)
}

type OverviewStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (OverviewStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OverviewStatusResponse)(nil)).Elem()
}

func (o OverviewStatusResponsePtrOutput) ToOverviewStatusResponsePtrOutput() OverviewStatusResponsePtrOutput {
	return o
}

func (o OverviewStatusResponsePtrOutput) ToOverviewStatusResponsePtrOutputWithContext(ctx context.Context) OverviewStatusResponsePtrOutput {
	return o
}

func (o OverviewStatusResponsePtrOutput) Elem() OverviewStatusResponseOutput {
	return o.ApplyT(func(v *OverviewStatusResponse) OverviewStatusResponse {
		if v != nil {
			return *v
		}
		var ret OverviewStatusResponse
		return ret
	}).(OverviewStatusResponseOutput)
}

// The count of all failed full automation control.
func (o OverviewStatusResponsePtrOutput) FailedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OverviewStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.FailedCount
	}).(pulumi.IntPtrOutput)
}

// The count of all manual control.
func (o OverviewStatusResponsePtrOutput) ManualCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OverviewStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.ManualCount
	}).(pulumi.IntPtrOutput)
}

// The count of all passed full automation control.
func (o OverviewStatusResponsePtrOutput) PassedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OverviewStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.PassedCount
	}).(pulumi.IntPtrOutput)
}

// A list which includes all the compliance result for one report.
type ReportComplianceStatusResponse struct {
	// The Microsoft 365 certification name.
	M365 *OverviewStatusResponse `pulumi:"m365"`
}

// A list which includes all the compliance result for one report.
type ReportComplianceStatusResponseOutput struct{ *pulumi.OutputState }

func (ReportComplianceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportComplianceStatusResponse)(nil)).Elem()
}

func (o ReportComplianceStatusResponseOutput) ToReportComplianceStatusResponseOutput() ReportComplianceStatusResponseOutput {
	return o
}

func (o ReportComplianceStatusResponseOutput) ToReportComplianceStatusResponseOutputWithContext(ctx context.Context) ReportComplianceStatusResponseOutput {
	return o
}

// The Microsoft 365 certification name.
func (o ReportComplianceStatusResponseOutput) M365() OverviewStatusResponsePtrOutput {
	return o.ApplyT(func(v ReportComplianceStatusResponse) *OverviewStatusResponse { return v.M365 }).(OverviewStatusResponsePtrOutput)
}

// Report's properties.
type ReportProperties struct {
	// Report offer Guid.
	OfferGuid *string `pulumi:"offerGuid"`
	// List of resource data.
	Resources []ResourceMetadata `pulumi:"resources"`
	// Report collection trigger time's time zone, the available list can be obtained by executing "Get-TimeZone -ListAvailable" in PowerShell.
	// An example of valid timezone id is "Pacific Standard Time".
	TimeZone string `pulumi:"timeZone"`
	// Report collection trigger time.
	TriggerTime string `pulumi:"triggerTime"`
}

// ReportPropertiesInput is an input type that accepts ReportPropertiesArgs and ReportPropertiesOutput values.
// You can construct a concrete instance of `ReportPropertiesInput` via:
//
//	ReportPropertiesArgs{...}
type ReportPropertiesInput interface {
	pulumi.Input

	ToReportPropertiesOutput() ReportPropertiesOutput
	ToReportPropertiesOutputWithContext(context.Context) ReportPropertiesOutput
}

// Report's properties.
type ReportPropertiesArgs struct {
	// Report offer Guid.
	OfferGuid pulumi.StringPtrInput `pulumi:"offerGuid"`
	// List of resource data.
	Resources ResourceMetadataArrayInput `pulumi:"resources"`
	// Report collection trigger time's time zone, the available list can be obtained by executing "Get-TimeZone -ListAvailable" in PowerShell.
	// An example of valid timezone id is "Pacific Standard Time".
	TimeZone pulumi.StringInput `pulumi:"timeZone"`
	// Report collection trigger time.
	TriggerTime pulumi.StringInput `pulumi:"triggerTime"`
}

func (ReportPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportProperties)(nil)).Elem()
}

func (i ReportPropertiesArgs) ToReportPropertiesOutput() ReportPropertiesOutput {
	return i.ToReportPropertiesOutputWithContext(context.Background())
}

func (i ReportPropertiesArgs) ToReportPropertiesOutputWithContext(ctx context.Context) ReportPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportPropertiesOutput)
}

// Report's properties.
type ReportPropertiesOutput struct{ *pulumi.OutputState }

func (ReportPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportProperties)(nil)).Elem()
}

func (o ReportPropertiesOutput) ToReportPropertiesOutput() ReportPropertiesOutput {
	return o
}

func (o ReportPropertiesOutput) ToReportPropertiesOutputWithContext(ctx context.Context) ReportPropertiesOutput {
	return o
}

// Report offer Guid.
func (o ReportPropertiesOutput) OfferGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportProperties) *string { return v.OfferGuid }).(pulumi.StringPtrOutput)
}

// List of resource data.
func (o ReportPropertiesOutput) Resources() ResourceMetadataArrayOutput {
	return o.ApplyT(func(v ReportProperties) []ResourceMetadata { return v.Resources }).(ResourceMetadataArrayOutput)
}

// Report collection trigger time's time zone, the available list can be obtained by executing "Get-TimeZone -ListAvailable" in PowerShell.
// An example of valid timezone id is "Pacific Standard Time".
func (o ReportPropertiesOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v ReportProperties) string { return v.TimeZone }).(pulumi.StringOutput)
}

// Report collection trigger time.
func (o ReportPropertiesOutput) TriggerTime() pulumi.StringOutput {
	return o.ApplyT(func(v ReportProperties) string { return v.TriggerTime }).(pulumi.StringOutput)
}

// Report's properties.
type ReportPropertiesResponse struct {
	// Report compliance status.
	ComplianceStatus ReportComplianceStatusResponse `pulumi:"complianceStatus"`
	// Report id in database.
	Id string `pulumi:"id"`
	// Report last collection trigger time.
	LastTriggerTime string `pulumi:"lastTriggerTime"`
	// Report next collection trigger time.
	NextTriggerTime string `pulumi:"nextTriggerTime"`
	// Report offer Guid.
	OfferGuid *string `pulumi:"offerGuid"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// Report name.
	ReportName string `pulumi:"reportName"`
	// List of resource data.
	Resources []ResourceMetadataResponse `pulumi:"resources"`
	// Report status.
	Status string `pulumi:"status"`
	// List of subscription Ids.
	Subscriptions []string `pulumi:"subscriptions"`
	// Report's tenant id.
	TenantId string `pulumi:"tenantId"`
	// Report collection trigger time's time zone, the available list can be obtained by executing "Get-TimeZone -ListAvailable" in PowerShell.
	// An example of valid timezone id is "Pacific Standard Time".
	TimeZone string `pulumi:"timeZone"`
	// Report collection trigger time.
	TriggerTime string `pulumi:"triggerTime"`
}

// Report's properties.
type ReportPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ReportPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportPropertiesResponse)(nil)).Elem()
}

func (o ReportPropertiesResponseOutput) ToReportPropertiesResponseOutput() ReportPropertiesResponseOutput {
	return o
}

func (o ReportPropertiesResponseOutput) ToReportPropertiesResponseOutputWithContext(ctx context.Context) ReportPropertiesResponseOutput {
	return o
}

// Report compliance status.
func (o ReportPropertiesResponseOutput) ComplianceStatus() ReportComplianceStatusResponseOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) ReportComplianceStatusResponse { return v.ComplianceStatus }).(ReportComplianceStatusResponseOutput)
}

// Report id in database.
func (o ReportPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

// Report last collection trigger time.
func (o ReportPropertiesResponseOutput) LastTriggerTime() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.LastTriggerTime }).(pulumi.StringOutput)
}

// Report next collection trigger time.
func (o ReportPropertiesResponseOutput) NextTriggerTime() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.NextTriggerTime }).(pulumi.StringOutput)
}

// Report offer Guid.
func (o ReportPropertiesResponseOutput) OfferGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) *string { return v.OfferGuid }).(pulumi.StringPtrOutput)
}

// Azure lifecycle management
func (o ReportPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Report name.
func (o ReportPropertiesResponseOutput) ReportName() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.ReportName }).(pulumi.StringOutput)
}

// List of resource data.
func (o ReportPropertiesResponseOutput) Resources() ResourceMetadataResponseArrayOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) []ResourceMetadataResponse { return v.Resources }).(ResourceMetadataResponseArrayOutput)
}

// Report status.
func (o ReportPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

// List of subscription Ids.
func (o ReportPropertiesResponseOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) []string { return v.Subscriptions }).(pulumi.StringArrayOutput)
}

// Report's tenant id.
func (o ReportPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

// Report collection trigger time's time zone, the available list can be obtained by executing "Get-TimeZone -ListAvailable" in PowerShell.
// An example of valid timezone id is "Pacific Standard Time".
func (o ReportPropertiesResponseOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.TimeZone }).(pulumi.StringOutput)
}

// Report collection trigger time.
func (o ReportPropertiesResponseOutput) TriggerTime() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.TriggerTime }).(pulumi.StringOutput)
}

// Single resource Id's metadata.
type ResourceMetadata struct {
	// Resource Id - e.g. "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1".
	ResourceId string `pulumi:"resourceId"`
	// Resource kind.
	ResourceKind *string `pulumi:"resourceKind"`
	// Resource name.
	ResourceName *string `pulumi:"resourceName"`
	// Resource type.
	ResourceType *string `pulumi:"resourceType"`
	// Resource's tag type.
	Tags map[string]string `pulumi:"tags"`
}

// ResourceMetadataInput is an input type that accepts ResourceMetadataArgs and ResourceMetadataOutput values.
// You can construct a concrete instance of `ResourceMetadataInput` via:
//
//	ResourceMetadataArgs{...}
type ResourceMetadataInput interface {
	pulumi.Input

	ToResourceMetadataOutput() ResourceMetadataOutput
	ToResourceMetadataOutputWithContext(context.Context) ResourceMetadataOutput
}

// Single resource Id's metadata.
type ResourceMetadataArgs struct {
	// Resource Id - e.g. "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1".
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Resource kind.
	ResourceKind pulumi.StringPtrInput `pulumi:"resourceKind"`
	// Resource name.
	ResourceName pulumi.StringPtrInput `pulumi:"resourceName"`
	// Resource type.
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	// Resource's tag type.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (ResourceMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMetadata)(nil)).Elem()
}

func (i ResourceMetadataArgs) ToResourceMetadataOutput() ResourceMetadataOutput {
	return i.ToResourceMetadataOutputWithContext(context.Background())
}

func (i ResourceMetadataArgs) ToResourceMetadataOutputWithContext(ctx context.Context) ResourceMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMetadataOutput)
}

// ResourceMetadataArrayInput is an input type that accepts ResourceMetadataArray and ResourceMetadataArrayOutput values.
// You can construct a concrete instance of `ResourceMetadataArrayInput` via:
//
//	ResourceMetadataArray{ ResourceMetadataArgs{...} }
type ResourceMetadataArrayInput interface {
	pulumi.Input

	ToResourceMetadataArrayOutput() ResourceMetadataArrayOutput
	ToResourceMetadataArrayOutputWithContext(context.Context) ResourceMetadataArrayOutput
}

type ResourceMetadataArray []ResourceMetadataInput

func (ResourceMetadataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMetadata)(nil)).Elem()
}

func (i ResourceMetadataArray) ToResourceMetadataArrayOutput() ResourceMetadataArrayOutput {
	return i.ToResourceMetadataArrayOutputWithContext(context.Background())
}

func (i ResourceMetadataArray) ToResourceMetadataArrayOutputWithContext(ctx context.Context) ResourceMetadataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMetadataArrayOutput)
}

// Single resource Id's metadata.
type ResourceMetadataOutput struct{ *pulumi.OutputState }

func (ResourceMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMetadata)(nil)).Elem()
}

func (o ResourceMetadataOutput) ToResourceMetadataOutput() ResourceMetadataOutput {
	return o
}

func (o ResourceMetadataOutput) ToResourceMetadataOutputWithContext(ctx context.Context) ResourceMetadataOutput {
	return o
}

// Resource Id - e.g. "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1".
func (o ResourceMetadataOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceMetadata) string { return v.ResourceId }).(pulumi.StringOutput)
}

// Resource kind.
func (o ResourceMetadataOutput) ResourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadata) *string { return v.ResourceKind }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ResourceMetadataOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadata) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ResourceMetadataOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadata) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// Resource's tag type.
func (o ResourceMetadataOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceMetadata) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ResourceMetadataArrayOutput struct{ *pulumi.OutputState }

func (ResourceMetadataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMetadata)(nil)).Elem()
}

func (o ResourceMetadataArrayOutput) ToResourceMetadataArrayOutput() ResourceMetadataArrayOutput {
	return o
}

func (o ResourceMetadataArrayOutput) ToResourceMetadataArrayOutputWithContext(ctx context.Context) ResourceMetadataArrayOutput {
	return o
}

func (o ResourceMetadataArrayOutput) Index(i pulumi.IntInput) ResourceMetadataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceMetadata {
		return vs[0].([]ResourceMetadata)[vs[1].(int)]
	}).(ResourceMetadataOutput)
}

// Single resource Id's metadata.
type ResourceMetadataResponse struct {
	// Resource Id - e.g. "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1".
	ResourceId string `pulumi:"resourceId"`
	// Resource kind.
	ResourceKind *string `pulumi:"resourceKind"`
	// Resource name.
	ResourceName *string `pulumi:"resourceName"`
	// Resource type.
	ResourceType *string `pulumi:"resourceType"`
	// Resource's tag type.
	Tags map[string]string `pulumi:"tags"`
}

// Single resource Id's metadata.
type ResourceMetadataResponseOutput struct{ *pulumi.OutputState }

func (ResourceMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMetadataResponse)(nil)).Elem()
}

func (o ResourceMetadataResponseOutput) ToResourceMetadataResponseOutput() ResourceMetadataResponseOutput {
	return o
}

func (o ResourceMetadataResponseOutput) ToResourceMetadataResponseOutputWithContext(ctx context.Context) ResourceMetadataResponseOutput {
	return o
}

// Resource Id - e.g. "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1".
func (o ResourceMetadataResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

// Resource kind.
func (o ResourceMetadataResponseOutput) ResourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) *string { return v.ResourceKind }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ResourceMetadataResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ResourceMetadataResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// Resource's tag type.
func (o ResourceMetadataResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ResourceMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMetadataResponse)(nil)).Elem()
}

func (o ResourceMetadataResponseArrayOutput) ToResourceMetadataResponseArrayOutput() ResourceMetadataResponseArrayOutput {
	return o
}

func (o ResourceMetadataResponseArrayOutput) ToResourceMetadataResponseArrayOutputWithContext(ctx context.Context) ResourceMetadataResponseArrayOutput {
	return o
}

func (o ResourceMetadataResponseArrayOutput) Index(i pulumi.IntInput) ResourceMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceMetadataResponse {
		return vs[0].([]ResourceMetadataResponse)[vs[1].(int)]
	}).(ResourceMetadataResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OverviewStatusResponseOutput{})
	pulumi.RegisterOutputType(OverviewStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportComplianceStatusResponseOutput{})
	pulumi.RegisterOutputType(ReportPropertiesOutput{})
	pulumi.RegisterOutputType(ReportPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ResourceMetadataOutput{})
	pulumi.RegisterOutputType(ResourceMetadataArrayOutput{})
	pulumi.RegisterOutputType(ResourceMetadataResponseOutput{})
	pulumi.RegisterOutputType(ResourceMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}