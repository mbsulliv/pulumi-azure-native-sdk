// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The top level Workspace resource container.
type Workspace struct {
	pulumi.CustomResourceState

	// Workspace creation date.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// This is a read-only property. Represents the ID associated with the workspace.
	CustomerId pulumi.StringOutput `pulumi:"customerId"`
	// The resource ID of the default Data Collection Rule to use for this workspace. Expected format is - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dcrName}.
	DefaultDataCollectionRuleResourceId pulumi.StringPtrOutput `pulumi:"defaultDataCollectionRuleResourceId"`
	// The ETag of the workspace.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Workspace features.
	Features WorkspaceFeaturesResponsePtrOutput `pulumi:"features"`
	// Indicates whether customer managed storage is mandatory for query management.
	ForceCmkForQuery pulumi.BoolPtrOutput `pulumi:"forceCmkForQuery"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Workspace modification date.
	ModifiedDate pulumi.StringOutput `pulumi:"modifiedDate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of linked private link scope resources.
	PrivateLinkScopedResources PrivateLinkScopedResourceResponseArrayOutput `pulumi:"privateLinkScopedResources"`
	// The provisioning state of the workspace.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion pulumi.StringPtrOutput `pulumi:"publicNetworkAccessForIngestion"`
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery pulumi.StringPtrOutput `pulumi:"publicNetworkAccessForQuery"`
	// The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers documentation for details.
	RetentionInDays pulumi.IntPtrOutput `pulumi:"retentionInDays"`
	// The SKU of the workspace.
	Sku WorkspaceSkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The daily volume cap for ingestion.
	WorkspaceCapping WorkspaceCappingResponsePtrOutput `pulumi:"workspaceCapping"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20151101preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20201001:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20210601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20221001:Workspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:operationalinsights/v20211201preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:operationalinsights/v20211201preview:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
}

type WorkspaceState struct {
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The resource ID of the default Data Collection Rule to use for this workspace. Expected format is - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dcrName}.
	DefaultDataCollectionRuleResourceId *string `pulumi:"defaultDataCollectionRuleResourceId"`
	// The ETag of the workspace.
	ETag *string `pulumi:"eTag"`
	// Workspace features.
	Features *WorkspaceFeatures `pulumi:"features"`
	// Indicates whether customer managed storage is mandatory for query management.
	ForceCmkForQuery *bool `pulumi:"forceCmkForQuery"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion *string `pulumi:"publicNetworkAccessForIngestion"`
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery *string `pulumi:"publicNetworkAccessForQuery"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers documentation for details.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The SKU of the workspace.
	Sku *WorkspaceSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCapping `pulumi:"workspaceCapping"`
	// The name of the workspace.
	WorkspaceName *string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The resource ID of the default Data Collection Rule to use for this workspace. Expected format is - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dcrName}.
	DefaultDataCollectionRuleResourceId pulumi.StringPtrInput
	// The ETag of the workspace.
	ETag pulumi.StringPtrInput
	// Workspace features.
	Features WorkspaceFeaturesPtrInput
	// Indicates whether customer managed storage is mandatory for query management.
	ForceCmkForQuery pulumi.BoolPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion pulumi.StringPtrInput
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers documentation for details.
	RetentionInDays pulumi.IntPtrInput
	// The SKU of the workspace.
	Sku WorkspaceSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The daily volume cap for ingestion.
	WorkspaceCapping WorkspaceCappingPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringPtrInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

// Workspace creation date.
func (o WorkspaceOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// This is a read-only property. Represents the ID associated with the workspace.
func (o WorkspaceOutput) CustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.CustomerId }).(pulumi.StringOutput)
}

// The resource ID of the default Data Collection Rule to use for this workspace. Expected format is - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dcrName}.
func (o WorkspaceOutput) DefaultDataCollectionRuleResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.DefaultDataCollectionRuleResourceId }).(pulumi.StringPtrOutput)
}

// The ETag of the workspace.
func (o WorkspaceOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Workspace features.
func (o WorkspaceOutput) Features() WorkspaceFeaturesResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceFeaturesResponsePtrOutput { return v.Features }).(WorkspaceFeaturesResponsePtrOutput)
}

// Indicates whether customer managed storage is mandatory for query management.
func (o WorkspaceOutput) ForceCmkForQuery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.ForceCmkForQuery }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o WorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Workspace modification date.
func (o WorkspaceOutput) ModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ModifiedDate }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of linked private link scope resources.
func (o WorkspaceOutput) PrivateLinkScopedResources() PrivateLinkScopedResourceResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) PrivateLinkScopedResourceResponseArrayOutput { return v.PrivateLinkScopedResources }).(PrivateLinkScopedResourceResponseArrayOutput)
}

// The provisioning state of the workspace.
func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The network access type for accessing Log Analytics ingestion.
func (o WorkspaceOutput) PublicNetworkAccessForIngestion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccessForIngestion }).(pulumi.StringPtrOutput)
}

// The network access type for accessing Log Analytics query.
func (o WorkspaceOutput) PublicNetworkAccessForQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccessForQuery }).(pulumi.StringPtrOutput)
}

// The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers documentation for details.
func (o WorkspaceOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.IntPtrOutput { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

// The SKU of the workspace.
func (o WorkspaceOutput) Sku() WorkspaceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceSkuResponsePtrOutput { return v.Sku }).(WorkspaceSkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o WorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The daily volume cap for ingestion.
func (o WorkspaceOutput) WorkspaceCapping() WorkspaceCappingResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceCappingResponsePtrOutput { return v.WorkspaceCapping }).(WorkspaceCappingResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
