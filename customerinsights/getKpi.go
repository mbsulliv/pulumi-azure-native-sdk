// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a KPI in the hub.
// Azure REST API version: 2017-04-26.
func LookupKpi(ctx *pulumi.Context, args *LookupKpiArgs, opts ...pulumi.InvokeOption) (*LookupKpiResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKpiResult
	err := ctx.Invoke("azure-native:customerinsights:getKpi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKpiArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the KPI.
	KpiName string `pulumi:"kpiName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The KPI resource format.
type LookupKpiResult struct {
	// The aliases.
	Aliases []KpiAliasResponse `pulumi:"aliases"`
	// The calculation window.
	CalculationWindow string `pulumi:"calculationWindow"`
	// Name of calculation window field.
	CalculationWindowFieldName *string `pulumi:"calculationWindowFieldName"`
	// Localized description for the KPI.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the KPI.
	DisplayName map[string]string `pulumi:"displayName"`
	// The mapping entity type.
	EntityType string `pulumi:"entityType"`
	// The mapping entity name.
	EntityTypeName string `pulumi:"entityTypeName"`
	// The computation expression for the KPI.
	Expression string `pulumi:"expression"`
	// The KPI extracts.
	Extracts []KpiExtractResponse `pulumi:"extracts"`
	// The filter expression for the KPI.
	Filter *string `pulumi:"filter"`
	// The computation function for the KPI.
	Function string `pulumi:"function"`
	// the group by properties for the KPI.
	GroupBy []string `pulumi:"groupBy"`
	// The KPI GroupByMetadata.
	GroupByMetadata []KpiGroupByMetadataResponse `pulumi:"groupByMetadata"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The KPI name.
	KpiName string `pulumi:"kpiName"`
	// Resource name.
	Name string `pulumi:"name"`
	// The participant profiles.
	ParticipantProfilesMetadata []KpiParticipantProfilesMetadataResponse `pulumi:"participantProfilesMetadata"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The hub name.
	TenantId string `pulumi:"tenantId"`
	// The KPI thresholds.
	ThresHolds *KpiThresholdsResponse `pulumi:"thresHolds"`
	// Resource type.
	Type string `pulumi:"type"`
	// The unit of measurement for the KPI.
	Unit *string `pulumi:"unit"`
}

func LookupKpiOutput(ctx *pulumi.Context, args LookupKpiOutputArgs, opts ...pulumi.InvokeOption) LookupKpiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKpiResult, error) {
			args := v.(LookupKpiArgs)
			r, err := LookupKpi(ctx, &args, opts...)
			var s LookupKpiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKpiResultOutput)
}

type LookupKpiOutputArgs struct {
	// The name of the hub.
	HubName pulumi.StringInput `pulumi:"hubName"`
	// The name of the KPI.
	KpiName pulumi.StringInput `pulumi:"kpiName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKpiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKpiArgs)(nil)).Elem()
}

// The KPI resource format.
type LookupKpiResultOutput struct{ *pulumi.OutputState }

func (LookupKpiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKpiResult)(nil)).Elem()
}

func (o LookupKpiResultOutput) ToLookupKpiResultOutput() LookupKpiResultOutput {
	return o
}

func (o LookupKpiResultOutput) ToLookupKpiResultOutputWithContext(ctx context.Context) LookupKpiResultOutput {
	return o
}

func (o LookupKpiResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupKpiResult] {
	return pulumix.Output[LookupKpiResult]{
		OutputState: o.OutputState,
	}
}

// The aliases.
func (o LookupKpiResultOutput) Aliases() KpiAliasResponseArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []KpiAliasResponse { return v.Aliases }).(KpiAliasResponseArrayOutput)
}

// The calculation window.
func (o LookupKpiResultOutput) CalculationWindow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.CalculationWindow }).(pulumi.StringOutput)
}

// Name of calculation window field.
func (o LookupKpiResultOutput) CalculationWindowFieldName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKpiResult) *string { return v.CalculationWindowFieldName }).(pulumi.StringPtrOutput)
}

// Localized description for the KPI.
func (o LookupKpiResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKpiResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

// Localized display name for the KPI.
func (o LookupKpiResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKpiResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

// The mapping entity type.
func (o LookupKpiResultOutput) EntityType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.EntityType }).(pulumi.StringOutput)
}

// The mapping entity name.
func (o LookupKpiResultOutput) EntityTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.EntityTypeName }).(pulumi.StringOutput)
}

// The computation expression for the KPI.
func (o LookupKpiResultOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Expression }).(pulumi.StringOutput)
}

// The KPI extracts.
func (o LookupKpiResultOutput) Extracts() KpiExtractResponseArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []KpiExtractResponse { return v.Extracts }).(KpiExtractResponseArrayOutput)
}

// The filter expression for the KPI.
func (o LookupKpiResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKpiResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The computation function for the KPI.
func (o LookupKpiResultOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Function }).(pulumi.StringOutput)
}

// the group by properties for the KPI.
func (o LookupKpiResultOutput) GroupBy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []string { return v.GroupBy }).(pulumi.StringArrayOutput)
}

// The KPI GroupByMetadata.
func (o LookupKpiResultOutput) GroupByMetadata() KpiGroupByMetadataResponseArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []KpiGroupByMetadataResponse { return v.GroupByMetadata }).(KpiGroupByMetadataResponseArrayOutput)
}

// Resource ID.
func (o LookupKpiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Id }).(pulumi.StringOutput)
}

// The KPI name.
func (o LookupKpiResultOutput) KpiName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.KpiName }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupKpiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Name }).(pulumi.StringOutput)
}

// The participant profiles.
func (o LookupKpiResultOutput) ParticipantProfilesMetadata() KpiParticipantProfilesMetadataResponseArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []KpiParticipantProfilesMetadataResponse { return v.ParticipantProfilesMetadata }).(KpiParticipantProfilesMetadataResponseArrayOutput)
}

// Provisioning state.
func (o LookupKpiResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The hub name.
func (o LookupKpiResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The KPI thresholds.
func (o LookupKpiResultOutput) ThresHolds() KpiThresholdsResponsePtrOutput {
	return o.ApplyT(func(v LookupKpiResult) *KpiThresholdsResponse { return v.ThresHolds }).(KpiThresholdsResponsePtrOutput)
}

// Resource type.
func (o LookupKpiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unit of measurement for the KPI.
func (o LookupKpiResultOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKpiResult) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKpiResultOutput{})
}
