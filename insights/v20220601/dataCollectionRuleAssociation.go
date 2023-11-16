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

// Definition of generic ARM proxy resource.
type DataCollectionRuleAssociation struct {
	pulumi.CustomResourceState

	// The resource ID of the data collection endpoint that is to be associated.
	DataCollectionEndpointId pulumi.StringPtrOutput `pulumi:"dataCollectionEndpointId"`
	// The resource ID of the data collection rule that is to be associated.
	DataCollectionRuleId pulumi.StringPtrOutput `pulumi:"dataCollectionRuleId"`
	// Description of the association.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource entity tag (ETag).
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Metadata about the resource
	Metadata DataCollectionRuleAssociationResponseMetadataOutput `pulumi:"metadata"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataCollectionRuleAssociation registers a new resource with the given unique name, arguments, and options.
func NewDataCollectionRuleAssociation(ctx *pulumi.Context,
	name string, args *DataCollectionRuleAssociationArgs, opts ...pulumi.ResourceOption) (*DataCollectionRuleAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:DataCollectionRuleAssociation"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20191101preview:DataCollectionRuleAssociation"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210401:DataCollectionRuleAssociation"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210901preview:DataCollectionRuleAssociation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataCollectionRuleAssociation
	err := ctx.RegisterResource("azure-native:insights/v20220601:DataCollectionRuleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCollectionRuleAssociation gets an existing DataCollectionRuleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCollectionRuleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCollectionRuleAssociationState, opts ...pulumi.ResourceOption) (*DataCollectionRuleAssociation, error) {
	var resource DataCollectionRuleAssociation
	err := ctx.ReadResource("azure-native:insights/v20220601:DataCollectionRuleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCollectionRuleAssociation resources.
type dataCollectionRuleAssociationState struct {
}

type DataCollectionRuleAssociationState struct {
}

func (DataCollectionRuleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleAssociationState)(nil)).Elem()
}

type dataCollectionRuleAssociationArgs struct {
	// The name of the association. The name is case insensitive.
	AssociationName *string `pulumi:"associationName"`
	// The resource ID of the data collection endpoint that is to be associated.
	DataCollectionEndpointId *string `pulumi:"dataCollectionEndpointId"`
	// The resource ID of the data collection rule that is to be associated.
	DataCollectionRuleId *string `pulumi:"dataCollectionRuleId"`
	// Description of the association.
	Description *string `pulumi:"description"`
	// The identifier of the resource.
	ResourceUri string `pulumi:"resourceUri"`
}

// The set of arguments for constructing a DataCollectionRuleAssociation resource.
type DataCollectionRuleAssociationArgs struct {
	// The name of the association. The name is case insensitive.
	AssociationName pulumi.StringPtrInput
	// The resource ID of the data collection endpoint that is to be associated.
	DataCollectionEndpointId pulumi.StringPtrInput
	// The resource ID of the data collection rule that is to be associated.
	DataCollectionRuleId pulumi.StringPtrInput
	// Description of the association.
	Description pulumi.StringPtrInput
	// The identifier of the resource.
	ResourceUri pulumi.StringInput
}

func (DataCollectionRuleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleAssociationArgs)(nil)).Elem()
}

type DataCollectionRuleAssociationInput interface {
	pulumi.Input

	ToDataCollectionRuleAssociationOutput() DataCollectionRuleAssociationOutput
	ToDataCollectionRuleAssociationOutputWithContext(ctx context.Context) DataCollectionRuleAssociationOutput
}

func (*DataCollectionRuleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleAssociation)(nil)).Elem()
}

func (i *DataCollectionRuleAssociation) ToDataCollectionRuleAssociationOutput() DataCollectionRuleAssociationOutput {
	return i.ToDataCollectionRuleAssociationOutputWithContext(context.Background())
}

func (i *DataCollectionRuleAssociation) ToDataCollectionRuleAssociationOutputWithContext(ctx context.Context) DataCollectionRuleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleAssociationOutput)
}

type DataCollectionRuleAssociationOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleAssociation)(nil)).Elem()
}

func (o DataCollectionRuleAssociationOutput) ToDataCollectionRuleAssociationOutput() DataCollectionRuleAssociationOutput {
	return o
}

func (o DataCollectionRuleAssociationOutput) ToDataCollectionRuleAssociationOutputWithContext(ctx context.Context) DataCollectionRuleAssociationOutput {
	return o
}

// The resource ID of the data collection endpoint that is to be associated.
func (o DataCollectionRuleAssociationOutput) DataCollectionEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringPtrOutput { return v.DataCollectionEndpointId }).(pulumi.StringPtrOutput)
}

// The resource ID of the data collection rule that is to be associated.
func (o DataCollectionRuleAssociationOutput) DataCollectionRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringPtrOutput { return v.DataCollectionRuleId }).(pulumi.StringPtrOutput)
}

// Description of the association.
func (o DataCollectionRuleAssociationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource entity tag (ETag).
func (o DataCollectionRuleAssociationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Metadata about the resource
func (o DataCollectionRuleAssociationOutput) Metadata() DataCollectionRuleAssociationResponseMetadataOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) DataCollectionRuleAssociationResponseMetadataOutput {
		return v.Metadata
	}).(DataCollectionRuleAssociationResponseMetadataOutput)
}

// The name of the resource.
func (o DataCollectionRuleAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource provisioning state.
func (o DataCollectionRuleAssociationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DataCollectionRuleAssociationOutput) SystemData() DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput {
		return v.SystemData
	}).(DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput)
}

// The type of the resource.
func (o DataCollectionRuleAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataCollectionRuleAssociationOutput{})
}
