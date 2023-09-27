// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Data Manager For Agriculture ARM Resource.
type DataManagerForAgricultureResource struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Uri of the Data Manager For Agriculture instance.
	InstanceUri pulumi.StringOutput `pulumi:"instanceUri"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoints.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Data Manager For Agriculture instance provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Property to allow or block public traffic for an Azure Data Manager For Agriculture resource.
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

// NewDataManagerForAgricultureResource registers a new resource with the given unique name, arguments, and options.
func NewDataManagerForAgricultureResource(ctx *pulumi.Context,
	name string, args *DataManagerForAgricultureResourceArgs, opts ...pulumi.ResourceOption) (*DataManagerForAgricultureResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:agfoodplatform:DataManagerForAgricultureResource"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20200512preview:DataManagerForAgricultureResource"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20210901preview:DataManagerForAgricultureResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataManagerForAgricultureResource
	err := ctx.RegisterResource("azure-native:agfoodplatform/v20230601preview:DataManagerForAgricultureResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataManagerForAgricultureResource gets an existing DataManagerForAgricultureResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataManagerForAgricultureResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataManagerForAgricultureResourceState, opts ...pulumi.ResourceOption) (*DataManagerForAgricultureResource, error) {
	var resource DataManagerForAgricultureResource
	err := ctx.ReadResource("azure-native:agfoodplatform/v20230601preview:DataManagerForAgricultureResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataManagerForAgricultureResource resources.
type dataManagerForAgricultureResourceState struct {
}

type DataManagerForAgricultureResourceState struct {
}

func (DataManagerForAgricultureResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataManagerForAgricultureResourceState)(nil)).Elem()
}

type dataManagerForAgricultureResourceArgs struct {
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName *string `pulumi:"dataManagerForAgricultureResourceName"`
	// Identity for the resource.
	Identity *Identity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Property to allow or block public traffic for an Azure Data Manager For Agriculture resource.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sensor integration request model.
	SensorIntegration *SensorIntegration `pulumi:"sensorIntegration"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataManagerForAgricultureResource resource.
type DataManagerForAgricultureResourceArgs struct {
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName pulumi.StringPtrInput
	// Identity for the resource.
	Identity IdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Property to allow or block public traffic for an Azure Data Manager For Agriculture resource.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sensor integration request model.
	SensorIntegration SensorIntegrationPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DataManagerForAgricultureResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataManagerForAgricultureResourceArgs)(nil)).Elem()
}

type DataManagerForAgricultureResourceInput interface {
	pulumi.Input

	ToDataManagerForAgricultureResourceOutput() DataManagerForAgricultureResourceOutput
	ToDataManagerForAgricultureResourceOutputWithContext(ctx context.Context) DataManagerForAgricultureResourceOutput
}

func (*DataManagerForAgricultureResource) ElementType() reflect.Type {
	return reflect.TypeOf((**DataManagerForAgricultureResource)(nil)).Elem()
}

func (i *DataManagerForAgricultureResource) ToDataManagerForAgricultureResourceOutput() DataManagerForAgricultureResourceOutput {
	return i.ToDataManagerForAgricultureResourceOutputWithContext(context.Background())
}

func (i *DataManagerForAgricultureResource) ToDataManagerForAgricultureResourceOutputWithContext(ctx context.Context) DataManagerForAgricultureResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataManagerForAgricultureResourceOutput)
}

func (i *DataManagerForAgricultureResource) ToOutput(ctx context.Context) pulumix.Output[*DataManagerForAgricultureResource] {
	return pulumix.Output[*DataManagerForAgricultureResource]{
		OutputState: i.ToDataManagerForAgricultureResourceOutputWithContext(ctx).OutputState,
	}
}

type DataManagerForAgricultureResourceOutput struct{ *pulumi.OutputState }

func (DataManagerForAgricultureResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataManagerForAgricultureResource)(nil)).Elem()
}

func (o DataManagerForAgricultureResourceOutput) ToDataManagerForAgricultureResourceOutput() DataManagerForAgricultureResourceOutput {
	return o
}

func (o DataManagerForAgricultureResourceOutput) ToDataManagerForAgricultureResourceOutputWithContext(ctx context.Context) DataManagerForAgricultureResourceOutput {
	return o
}

func (o DataManagerForAgricultureResourceOutput) ToOutput(ctx context.Context) pulumix.Output[*DataManagerForAgricultureResource] {
	return pulumix.Output[*DataManagerForAgricultureResource]{
		OutputState: o.OutputState,
	}
}

// Identity for the resource.
func (o DataManagerForAgricultureResourceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Uri of the Data Manager For Agriculture instance.
func (o DataManagerForAgricultureResourceOutput) InstanceUri() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) pulumi.StringOutput { return v.InstanceUri }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o DataManagerForAgricultureResourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DataManagerForAgricultureResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoints.
func (o DataManagerForAgricultureResourceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Data Manager For Agriculture instance provisioning state.
func (o DataManagerForAgricultureResourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Property to allow or block public traffic for an Azure Data Manager For Agriculture resource.
func (o DataManagerForAgricultureResourceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Sensor integration request model.
func (o DataManagerForAgricultureResourceOutput) SensorIntegration() SensorIntegrationResponsePtrOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) SensorIntegrationResponsePtrOutput {
		return v.SensorIntegration
	}).(SensorIntegrationResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DataManagerForAgricultureResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DataManagerForAgricultureResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataManagerForAgricultureResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManagerForAgricultureResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataManagerForAgricultureResourceOutput{})
}
