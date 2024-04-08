// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of ARM tracked top level resource.
type DataCollectionRule struct {
	pulumi.CustomResourceState

	// The resource ID of the data collection endpoint that this rule can be used with.
	DataCollectionEndpointId pulumi.StringPtrOutput `pulumi:"dataCollectionEndpointId"`
	// The specification of data flows.
	DataFlows DataFlowResponseArrayOutput `pulumi:"dataFlows"`
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources DataCollectionRuleResponseDataSourcesPtrOutput `pulumi:"dataSources"`
	// Description of the data collection rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The specification of destinations.
	Destinations DataCollectionRuleResponseDestinationsPtrOutput `pulumi:"destinations"`
	// Resource entity tag (ETag).
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Managed service identity of the resource.
	Identity DataCollectionRuleResourceResponseIdentityPtrOutput `pulumi:"identity"`
	// The immutable ID of this data collection rule. This property is READ-ONLY.
	ImmutableId pulumi.StringOutput `pulumi:"immutableId"`
	// The kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives.
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata about the resource
	Metadata DataCollectionRuleResponseMetadataOutput `pulumi:"metadata"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Declaration of custom streams used in this rule.
	StreamDeclarations StreamDeclarationResponseMapOutput `pulumi:"streamDeclarations"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData DataCollectionRuleResourceResponseSystemDataOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataCollectionRule registers a new resource with the given unique name, arguments, and options.
func NewDataCollectionRule(ctx *pulumi.Context,
	name string, args *DataCollectionRuleArgs, opts ...pulumi.ResourceOption) (*DataCollectionRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:DataCollectionRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20191101preview:DataCollectionRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210401:DataCollectionRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210901preview:DataCollectionRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20230311:DataCollectionRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataCollectionRule
	err := ctx.RegisterResource("azure-native:insights/v20220601:DataCollectionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCollectionRule gets an existing DataCollectionRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCollectionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCollectionRuleState, opts ...pulumi.ResourceOption) (*DataCollectionRule, error) {
	var resource DataCollectionRule
	err := ctx.ReadResource("azure-native:insights/v20220601:DataCollectionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCollectionRule resources.
type dataCollectionRuleState struct {
}

type DataCollectionRuleState struct {
}

func (DataCollectionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleState)(nil)).Elem()
}

type dataCollectionRuleArgs struct {
	// The resource ID of the data collection endpoint that this rule can be used with.
	DataCollectionEndpointId *string `pulumi:"dataCollectionEndpointId"`
	// The name of the data collection rule. The name is case insensitive.
	DataCollectionRuleName *string `pulumi:"dataCollectionRuleName"`
	// The specification of data flows.
	DataFlows []DataFlow `pulumi:"dataFlows"`
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources *DataCollectionRuleDataSources `pulumi:"dataSources"`
	// Description of the data collection rule.
	Description *string `pulumi:"description"`
	// The specification of destinations.
	Destinations *DataCollectionRuleDestinations `pulumi:"destinations"`
	// Managed service identity of the resource.
	Identity *DataCollectionRuleResourceIdentity `pulumi:"identity"`
	// The kind of the resource.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Declaration of custom streams used in this rule.
	StreamDeclarations map[string]StreamDeclaration `pulumi:"streamDeclarations"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataCollectionRule resource.
type DataCollectionRuleArgs struct {
	// The resource ID of the data collection endpoint that this rule can be used with.
	DataCollectionEndpointId pulumi.StringPtrInput
	// The name of the data collection rule. The name is case insensitive.
	DataCollectionRuleName pulumi.StringPtrInput
	// The specification of data flows.
	DataFlows DataFlowArrayInput
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources DataCollectionRuleDataSourcesPtrInput
	// Description of the data collection rule.
	Description pulumi.StringPtrInput
	// The specification of destinations.
	Destinations DataCollectionRuleDestinationsPtrInput
	// Managed service identity of the resource.
	Identity DataCollectionRuleResourceIdentityPtrInput
	// The kind of the resource.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Declaration of custom streams used in this rule.
	StreamDeclarations StreamDeclarationMapInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DataCollectionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleArgs)(nil)).Elem()
}

type DataCollectionRuleInput interface {
	pulumi.Input

	ToDataCollectionRuleOutput() DataCollectionRuleOutput
	ToDataCollectionRuleOutputWithContext(ctx context.Context) DataCollectionRuleOutput
}

func (*DataCollectionRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRule)(nil)).Elem()
}

func (i *DataCollectionRule) ToDataCollectionRuleOutput() DataCollectionRuleOutput {
	return i.ToDataCollectionRuleOutputWithContext(context.Background())
}

func (i *DataCollectionRule) ToDataCollectionRuleOutputWithContext(ctx context.Context) DataCollectionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleOutput)
}

type DataCollectionRuleOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRule)(nil)).Elem()
}

func (o DataCollectionRuleOutput) ToDataCollectionRuleOutput() DataCollectionRuleOutput {
	return o
}

func (o DataCollectionRuleOutput) ToDataCollectionRuleOutputWithContext(ctx context.Context) DataCollectionRuleOutput {
	return o
}

// The resource ID of the data collection endpoint that this rule can be used with.
func (o DataCollectionRuleOutput) DataCollectionEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringPtrOutput { return v.DataCollectionEndpointId }).(pulumi.StringPtrOutput)
}

// The specification of data flows.
func (o DataCollectionRuleOutput) DataFlows() DataFlowResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataFlowResponseArrayOutput { return v.DataFlows }).(DataFlowResponseArrayOutput)
}

// The specification of data sources.
// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
func (o DataCollectionRuleOutput) DataSources() DataCollectionRuleResponseDataSourcesPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataCollectionRuleResponseDataSourcesPtrOutput { return v.DataSources }).(DataCollectionRuleResponseDataSourcesPtrOutput)
}

// Description of the data collection rule.
func (o DataCollectionRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The specification of destinations.
func (o DataCollectionRuleOutput) Destinations() DataCollectionRuleResponseDestinationsPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataCollectionRuleResponseDestinationsPtrOutput { return v.Destinations }).(DataCollectionRuleResponseDestinationsPtrOutput)
}

// Resource entity tag (ETag).
func (o DataCollectionRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Managed service identity of the resource.
func (o DataCollectionRuleOutput) Identity() DataCollectionRuleResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataCollectionRuleResourceResponseIdentityPtrOutput { return v.Identity }).(DataCollectionRuleResourceResponseIdentityPtrOutput)
}

// The immutable ID of this data collection rule. This property is READ-ONLY.
func (o DataCollectionRuleOutput) ImmutableId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.ImmutableId }).(pulumi.StringOutput)
}

// The kind of the resource.
func (o DataCollectionRuleOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives.
func (o DataCollectionRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Metadata about the resource
func (o DataCollectionRuleOutput) Metadata() DataCollectionRuleResponseMetadataOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataCollectionRuleResponseMetadataOutput { return v.Metadata }).(DataCollectionRuleResponseMetadataOutput)
}

// The name of the resource.
func (o DataCollectionRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource provisioning state.
func (o DataCollectionRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Declaration of custom streams used in this rule.
func (o DataCollectionRuleOutput) StreamDeclarations() StreamDeclarationResponseMapOutput {
	return o.ApplyT(func(v *DataCollectionRule) StreamDeclarationResponseMapOutput { return v.StreamDeclarations }).(StreamDeclarationResponseMapOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DataCollectionRuleOutput) SystemData() DataCollectionRuleResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataCollectionRuleResourceResponseSystemDataOutput { return v.SystemData }).(DataCollectionRuleResourceResponseSystemDataOutput)
}

// Resource tags.
func (o DataCollectionRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o DataCollectionRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataCollectionRuleOutput{})
}
