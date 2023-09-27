// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// FarmBeats ARM Resource.
type FarmBeatsModel struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Uri of the FarmBeats instance.
	InstanceUri pulumi.StringOutput `pulumi:"instanceUri"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The private endpoint connection resource.
	PrivateEndpointConnections PrivateEndpointConnectionResponseOutput `pulumi:"privateEndpointConnections"`
	// FarmBeats instance provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Property to allow or block public traffic for an Azure FarmBeats resource.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Sensor integration request model.
	SensorIntegration SensorIntegrationResponsePtrOutput `pulumi:"sensorIntegration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFarmBeatsModel registers a new resource with the given unique name, arguments, and options.
func NewFarmBeatsModel(ctx *pulumi.Context,
	name string, args *FarmBeatsModelArgs, opts ...pulumi.ResourceOption) (*FarmBeatsModel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:agfoodplatform:FarmBeatsModel"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20200512preview:FarmBeatsModel"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20230601preview:FarmBeatsModel"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FarmBeatsModel
	err := ctx.RegisterResource("azure-native:agfoodplatform/v20210901preview:FarmBeatsModel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFarmBeatsModel gets an existing FarmBeatsModel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFarmBeatsModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FarmBeatsModelState, opts ...pulumi.ResourceOption) (*FarmBeatsModel, error) {
	var resource FarmBeatsModel
	err := ctx.ReadResource("azure-native:agfoodplatform/v20210901preview:FarmBeatsModel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FarmBeatsModel resources.
type farmBeatsModelState struct {
}

type FarmBeatsModelState struct {
}

func (FarmBeatsModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*farmBeatsModelState)(nil)).Elem()
}

type farmBeatsModelArgs struct {
	// FarmBeats resource name.
	FarmBeatsResourceName *string `pulumi:"farmBeatsResourceName"`
	// Identity for the resource.
	Identity *Identity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Property to allow or block public traffic for an Azure FarmBeats resource.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sensor integration request model.
	SensorIntegration *SensorIntegration `pulumi:"sensorIntegration"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FarmBeatsModel resource.
type FarmBeatsModelArgs struct {
	// FarmBeats resource name.
	FarmBeatsResourceName pulumi.StringPtrInput
	// Identity for the resource.
	Identity IdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Property to allow or block public traffic for an Azure FarmBeats resource.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sensor integration request model.
	SensorIntegration SensorIntegrationPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (FarmBeatsModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*farmBeatsModelArgs)(nil)).Elem()
}

type FarmBeatsModelInput interface {
	pulumi.Input

	ToFarmBeatsModelOutput() FarmBeatsModelOutput
	ToFarmBeatsModelOutputWithContext(ctx context.Context) FarmBeatsModelOutput
}

func (*FarmBeatsModel) ElementType() reflect.Type {
	return reflect.TypeOf((**FarmBeatsModel)(nil)).Elem()
}

func (i *FarmBeatsModel) ToFarmBeatsModelOutput() FarmBeatsModelOutput {
	return i.ToFarmBeatsModelOutputWithContext(context.Background())
}

func (i *FarmBeatsModel) ToFarmBeatsModelOutputWithContext(ctx context.Context) FarmBeatsModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FarmBeatsModelOutput)
}

func (i *FarmBeatsModel) ToOutput(ctx context.Context) pulumix.Output[*FarmBeatsModel] {
	return pulumix.Output[*FarmBeatsModel]{
		OutputState: i.ToFarmBeatsModelOutputWithContext(ctx).OutputState,
	}
}

type FarmBeatsModelOutput struct{ *pulumi.OutputState }

func (FarmBeatsModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FarmBeatsModel)(nil)).Elem()
}

func (o FarmBeatsModelOutput) ToFarmBeatsModelOutput() FarmBeatsModelOutput {
	return o
}

func (o FarmBeatsModelOutput) ToFarmBeatsModelOutputWithContext(ctx context.Context) FarmBeatsModelOutput {
	return o
}

func (o FarmBeatsModelOutput) ToOutput(ctx context.Context) pulumix.Output[*FarmBeatsModel] {
	return pulumix.Output[*FarmBeatsModel]{
		OutputState: o.OutputState,
	}
}

// Identity for the resource.
func (o FarmBeatsModelOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *FarmBeatsModel) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Uri of the FarmBeats instance.
func (o FarmBeatsModelOutput) InstanceUri() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.InstanceUri }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o FarmBeatsModelOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o FarmBeatsModelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The private endpoint connection resource.
func (o FarmBeatsModelOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v *FarmBeatsModel) PrivateEndpointConnectionResponseOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseOutput)
}

// FarmBeats instance provisioning state.
func (o FarmBeatsModelOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Property to allow or block public traffic for an Azure FarmBeats resource.
func (o FarmBeatsModelOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Sensor integration request model.
func (o FarmBeatsModelOutput) SensorIntegration() SensorIntegrationResponsePtrOutput {
	return o.ApplyT(func(v *FarmBeatsModel) SensorIntegrationResponsePtrOutput { return v.SensorIntegration }).(SensorIntegrationResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FarmBeatsModelOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FarmBeatsModel) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o FarmBeatsModelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FarmBeatsModelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FarmBeatsModel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FarmBeatsModelOutput{})
}
