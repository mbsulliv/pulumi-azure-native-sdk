// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an incident.
func LookupIncident(ctx *pulumi.Context, args *LookupIncidentArgs, opts ...pulumi.InvokeOption) (*LookupIncidentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIncidentResult
	err := ctx.Invoke("azure-native:securityinsights/v20230601preview:getIncident", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentArgs struct {
	// Incident ID
	IncidentId string `pulumi:"incidentId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupIncidentResult struct {
	// Additional data on the incident
	AdditionalData IncidentAdditionalDataResponse `pulumi:"additionalData"`
	// The reason the incident was closed
	Classification *string `pulumi:"classification"`
	// Describes the reason the incident was closed
	ClassificationComment *string `pulumi:"classificationComment"`
	// The classification reason the incident was closed with
	ClassificationReason *string `pulumi:"classificationReason"`
	// The time the incident was created
	CreatedTimeUtc string `pulumi:"createdTimeUtc"`
	// The description of the incident
	Description *string `pulumi:"description"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The time of the first activity in the incident
	FirstActivityTimeUtc *string `pulumi:"firstActivityTimeUtc"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// A sequential number
	IncidentNumber int `pulumi:"incidentNumber"`
	// The deep-link url to the incident in Azure portal
	IncidentUrl string `pulumi:"incidentUrl"`
	// List of labels relevant to this incident
	Labels []IncidentLabelResponse `pulumi:"labels"`
	// The time of the last activity in the incident
	LastActivityTimeUtc *string `pulumi:"lastActivityTimeUtc"`
	// The last time the incident was updated
	LastModifiedTimeUtc string `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Describes a user that the incident is assigned to
	Owner *IncidentOwnerInfoResponse `pulumi:"owner"`
	// The incident ID assigned by the incident provider
	ProviderIncidentId string `pulumi:"providerIncidentId"`
	// The name of the source provider that generated the incident
	ProviderName string `pulumi:"providerName"`
	// List of resource ids of Analytic rules related to the incident
	RelatedAnalyticRuleIds []string `pulumi:"relatedAnalyticRuleIds"`
	// The severity of the incident
	Severity string `pulumi:"severity"`
	// The status of the incident
	Status string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Describes a team for the incident
	TeamInformation *TeamInformationResponse `pulumi:"teamInformation"`
	// The title of the incident
	Title string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIncidentOutput(ctx *pulumi.Context, args LookupIncidentOutputArgs, opts ...pulumi.InvokeOption) LookupIncidentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIncidentResult, error) {
			args := v.(LookupIncidentArgs)
			r, err := LookupIncident(ctx, &args, opts...)
			var s LookupIncidentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIncidentResultOutput)
}

type LookupIncidentOutputArgs struct {
	// Incident ID
	IncidentId pulumi.StringInput `pulumi:"incidentId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIncidentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentArgs)(nil)).Elem()
}

type LookupIncidentResultOutput struct{ *pulumi.OutputState }

func (LookupIncidentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentResult)(nil)).Elem()
}

func (o LookupIncidentResultOutput) ToLookupIncidentResultOutput() LookupIncidentResultOutput {
	return o
}

func (o LookupIncidentResultOutput) ToLookupIncidentResultOutputWithContext(ctx context.Context) LookupIncidentResultOutput {
	return o
}

func (o LookupIncidentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIncidentResult] {
	return pulumix.Output[LookupIncidentResult]{
		OutputState: o.OutputState,
	}
}

// Additional data on the incident
func (o LookupIncidentResultOutput) AdditionalData() IncidentAdditionalDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentResult) IncidentAdditionalDataResponse { return v.AdditionalData }).(IncidentAdditionalDataResponseOutput)
}

// The reason the incident was closed
func (o LookupIncidentResultOutput) Classification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.Classification }).(pulumi.StringPtrOutput)
}

// Describes the reason the incident was closed
func (o LookupIncidentResultOutput) ClassificationComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.ClassificationComment }).(pulumi.StringPtrOutput)
}

// The classification reason the incident was closed with
func (o LookupIncidentResultOutput) ClassificationReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.ClassificationReason }).(pulumi.StringPtrOutput)
}

// The time the incident was created
func (o LookupIncidentResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// The description of the incident
func (o LookupIncidentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Etag of the azure resource
func (o LookupIncidentResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The time of the first activity in the incident
func (o LookupIncidentResultOutput) FirstActivityTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.FirstActivityTimeUtc }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupIncidentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Id }).(pulumi.StringOutput)
}

// A sequential number
func (o LookupIncidentResultOutput) IncidentNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIncidentResult) int { return v.IncidentNumber }).(pulumi.IntOutput)
}

// The deep-link url to the incident in Azure portal
func (o LookupIncidentResultOutput) IncidentUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.IncidentUrl }).(pulumi.StringOutput)
}

// List of labels relevant to this incident
func (o LookupIncidentResultOutput) Labels() IncidentLabelResponseArrayOutput {
	return o.ApplyT(func(v LookupIncidentResult) []IncidentLabelResponse { return v.Labels }).(IncidentLabelResponseArrayOutput)
}

// The time of the last activity in the incident
func (o LookupIncidentResultOutput) LastActivityTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.LastActivityTimeUtc }).(pulumi.StringPtrOutput)
}

// The last time the incident was updated
func (o LookupIncidentResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIncidentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describes a user that the incident is assigned to
func (o LookupIncidentResultOutput) Owner() IncidentOwnerInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *IncidentOwnerInfoResponse { return v.Owner }).(IncidentOwnerInfoResponsePtrOutput)
}

// The incident ID assigned by the incident provider
func (o LookupIncidentResultOutput) ProviderIncidentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.ProviderIncidentId }).(pulumi.StringOutput)
}

// The name of the source provider that generated the incident
func (o LookupIncidentResultOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.ProviderName }).(pulumi.StringOutput)
}

// List of resource ids of Analytic rules related to the incident
func (o LookupIncidentResultOutput) RelatedAnalyticRuleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIncidentResult) []string { return v.RelatedAnalyticRuleIds }).(pulumi.StringArrayOutput)
}

// The severity of the incident
func (o LookupIncidentResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Severity }).(pulumi.StringOutput)
}

// The status of the incident
func (o LookupIncidentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIncidentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Describes a team for the incident
func (o LookupIncidentResultOutput) TeamInformation() TeamInformationResponsePtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *TeamInformationResponse { return v.TeamInformation }).(TeamInformationResponsePtrOutput)
}

// The title of the incident
func (o LookupIncidentResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIncidentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIncidentResultOutput{})
}
