// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an incident in Azure Security Insights.
type Incident struct {
	pulumi.CustomResourceState

	// Additional data on the incident
	AdditionalData IncidentAdditionalDataResponseOutput `pulumi:"additionalData"`
	// The reason the incident was closed
	Classification pulumi.StringPtrOutput `pulumi:"classification"`
	// Describes the reason the incident was closed
	ClassificationComment pulumi.StringPtrOutput `pulumi:"classificationComment"`
	// The classification reason the incident was closed with
	ClassificationReason pulumi.StringPtrOutput `pulumi:"classificationReason"`
	// The time the incident was created
	CreatedTimeUtc pulumi.StringOutput `pulumi:"createdTimeUtc"`
	// The description of the incident
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The time of the first activity in the incident
	FirstActivityTimeUtc pulumi.StringPtrOutput `pulumi:"firstActivityTimeUtc"`
	// A sequential number
	IncidentNumber pulumi.IntOutput `pulumi:"incidentNumber"`
	// The deep-link url to the incident in Azure portal
	IncidentUrl pulumi.StringOutput `pulumi:"incidentUrl"`
	// List of labels relevant to this incident
	Labels IncidentLabelResponseArrayOutput `pulumi:"labels"`
	// The time of the last activity in the incident
	LastActivityTimeUtc pulumi.StringPtrOutput `pulumi:"lastActivityTimeUtc"`
	// The last time the incident was updated
	LastModifiedTimeUtc pulumi.StringOutput `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes a user that the incident is assigned to
	Owner IncidentOwnerInfoResponsePtrOutput `pulumi:"owner"`
	// The incident ID assigned by the incident provider
	ProviderIncidentId pulumi.StringPtrOutput `pulumi:"providerIncidentId"`
	// The name of the source provider that generated the incident
	ProviderName pulumi.StringPtrOutput `pulumi:"providerName"`
	// List of resource ids of Analytic rules related to the incident
	RelatedAnalyticRuleIds pulumi.StringArrayOutput `pulumi:"relatedAnalyticRuleIds"`
	// The severity of the incident
	Severity pulumi.StringOutput `pulumi:"severity"`
	// The status of the incident
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Describes a team for the incident
	TeamInformation TeamInformationResponsePtrOutput `pulumi:"teamInformation"`
	// The title of the incident
	Title pulumi.StringOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIncident registers a new resource with the given unique name, arguments, and options.
func NewIncident(ctx *pulumi.Context,
	name string, args *IncidentArgs, opts ...pulumi.ResourceOption) (*Incident, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Incident"),
		},
	})
	opts = append(opts, aliases)
	var resource Incident
	err := ctx.RegisterResource("azure-native:securityinsights/v20220701preview:Incident", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIncident gets an existing Incident resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIncident(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentState, opts ...pulumi.ResourceOption) (*Incident, error) {
	var resource Incident
	err := ctx.ReadResource("azure-native:securityinsights/v20220701preview:Incident", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Incident resources.
type incidentState struct {
}

type IncidentState struct {
}

func (IncidentState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentState)(nil)).Elem()
}

type incidentArgs struct {
	// The reason the incident was closed
	Classification *string `pulumi:"classification"`
	// Describes the reason the incident was closed
	ClassificationComment *string `pulumi:"classificationComment"`
	// The classification reason the incident was closed with
	ClassificationReason *string `pulumi:"classificationReason"`
	// The description of the incident
	Description *string `pulumi:"description"`
	// The time of the first activity in the incident
	FirstActivityTimeUtc *string `pulumi:"firstActivityTimeUtc"`
	// Incident ID
	IncidentId *string `pulumi:"incidentId"`
	// List of labels relevant to this incident
	Labels []IncidentLabel `pulumi:"labels"`
	// The time of the last activity in the incident
	LastActivityTimeUtc *string `pulumi:"lastActivityTimeUtc"`
	// Describes a user that the incident is assigned to
	Owner *IncidentOwnerInfo `pulumi:"owner"`
	// The incident ID assigned by the incident provider
	ProviderIncidentId *string `pulumi:"providerIncidentId"`
	// The name of the source provider that generated the incident
	ProviderName *string `pulumi:"providerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The severity of the incident
	Severity string `pulumi:"severity"`
	// The status of the incident
	Status string `pulumi:"status"`
	// The title of the incident
	Title string `pulumi:"title"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Incident resource.
type IncidentArgs struct {
	// The reason the incident was closed
	Classification pulumi.StringPtrInput
	// Describes the reason the incident was closed
	ClassificationComment pulumi.StringPtrInput
	// The classification reason the incident was closed with
	ClassificationReason pulumi.StringPtrInput
	// The description of the incident
	Description pulumi.StringPtrInput
	// The time of the first activity in the incident
	FirstActivityTimeUtc pulumi.StringPtrInput
	// Incident ID
	IncidentId pulumi.StringPtrInput
	// List of labels relevant to this incident
	Labels IncidentLabelArrayInput
	// The time of the last activity in the incident
	LastActivityTimeUtc pulumi.StringPtrInput
	// Describes a user that the incident is assigned to
	Owner IncidentOwnerInfoPtrInput
	// The incident ID assigned by the incident provider
	ProviderIncidentId pulumi.StringPtrInput
	// The name of the source provider that generated the incident
	ProviderName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The severity of the incident
	Severity pulumi.StringInput
	// The status of the incident
	Status pulumi.StringInput
	// The title of the incident
	Title pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IncidentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentArgs)(nil)).Elem()
}

type IncidentInput interface {
	pulumi.Input

	ToIncidentOutput() IncidentOutput
	ToIncidentOutputWithContext(ctx context.Context) IncidentOutput
}

func (*Incident) ElementType() reflect.Type {
	return reflect.TypeOf((**Incident)(nil)).Elem()
}

func (i *Incident) ToIncidentOutput() IncidentOutput {
	return i.ToIncidentOutputWithContext(context.Background())
}

func (i *Incident) ToIncidentOutputWithContext(ctx context.Context) IncidentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOutput)
}

type IncidentOutput struct{ *pulumi.OutputState }

func (IncidentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Incident)(nil)).Elem()
}

func (o IncidentOutput) ToIncidentOutput() IncidentOutput {
	return o
}

func (o IncidentOutput) ToIncidentOutputWithContext(ctx context.Context) IncidentOutput {
	return o
}

// Additional data on the incident
func (o IncidentOutput) AdditionalData() IncidentAdditionalDataResponseOutput {
	return o.ApplyT(func(v *Incident) IncidentAdditionalDataResponseOutput { return v.AdditionalData }).(IncidentAdditionalDataResponseOutput)
}

// The reason the incident was closed
func (o IncidentOutput) Classification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.Classification }).(pulumi.StringPtrOutput)
}

// Describes the reason the incident was closed
func (o IncidentOutput) ClassificationComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.ClassificationComment }).(pulumi.StringPtrOutput)
}

// The classification reason the incident was closed with
func (o IncidentOutput) ClassificationReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.ClassificationReason }).(pulumi.StringPtrOutput)
}

// The time the incident was created
func (o IncidentOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// The description of the incident
func (o IncidentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Etag of the azure resource
func (o IncidentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The time of the first activity in the incident
func (o IncidentOutput) FirstActivityTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.FirstActivityTimeUtc }).(pulumi.StringPtrOutput)
}

// A sequential number
func (o IncidentOutput) IncidentNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Incident) pulumi.IntOutput { return v.IncidentNumber }).(pulumi.IntOutput)
}

// The deep-link url to the incident in Azure portal
func (o IncidentOutput) IncidentUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.IncidentUrl }).(pulumi.StringOutput)
}

// List of labels relevant to this incident
func (o IncidentOutput) Labels() IncidentLabelResponseArrayOutput {
	return o.ApplyT(func(v *Incident) IncidentLabelResponseArrayOutput { return v.Labels }).(IncidentLabelResponseArrayOutput)
}

// The time of the last activity in the incident
func (o IncidentOutput) LastActivityTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.LastActivityTimeUtc }).(pulumi.StringPtrOutput)
}

// The last time the incident was updated
func (o IncidentOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o IncidentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Describes a user that the incident is assigned to
func (o IncidentOutput) Owner() IncidentOwnerInfoResponsePtrOutput {
	return o.ApplyT(func(v *Incident) IncidentOwnerInfoResponsePtrOutput { return v.Owner }).(IncidentOwnerInfoResponsePtrOutput)
}

// The incident ID assigned by the incident provider
func (o IncidentOutput) ProviderIncidentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.ProviderIncidentId }).(pulumi.StringPtrOutput)
}

// The name of the source provider that generated the incident
func (o IncidentOutput) ProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.ProviderName }).(pulumi.StringPtrOutput)
}

// List of resource ids of Analytic rules related to the incident
func (o IncidentOutput) RelatedAnalyticRuleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringArrayOutput { return v.RelatedAnalyticRuleIds }).(pulumi.StringArrayOutput)
}

// The severity of the incident
func (o IncidentOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// The status of the incident
func (o IncidentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IncidentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Incident) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Describes a team for the incident
func (o IncidentOutput) TeamInformation() TeamInformationResponsePtrOutput {
	return o.ApplyT(func(v *Incident) TeamInformationResponsePtrOutput { return v.TeamInformation }).(TeamInformationResponsePtrOutput)
}

// The title of the incident
func (o IncidentOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IncidentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentOutput{})
}
