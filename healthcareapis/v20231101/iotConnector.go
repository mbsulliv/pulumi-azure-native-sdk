// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IoT Connector definition.
type IotConnector struct {
	pulumi.CustomResourceState

	// Device Mappings.
	DeviceMapping IotMappingPropertiesResponsePtrOutput `pulumi:"deviceMapping"`
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServiceManagedIdentityResponseIdentityPtrOutput `pulumi:"identity"`
	// Source configuration.
	IngestionEndpointConfiguration IotEventHubIngestionEndpointConfigurationResponsePtrOutput `pulumi:"ingestionEndpointConfiguration"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotConnector registers a new resource with the given unique name, arguments, and options.
func NewIotConnector(ctx *pulumi.Context,
	name string, args *IotConnectorArgs, opts ...pulumi.ResourceOption) (*IotConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210601preview:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221001preview:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221201:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20230228:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20230906:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20231201:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20240301:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20240331:IotConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IotConnector
	err := ctx.RegisterResource("azure-native:healthcareapis/v20231101:IotConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotConnector gets an existing IotConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotConnectorState, opts ...pulumi.ResourceOption) (*IotConnector, error) {
	var resource IotConnector
	err := ctx.ReadResource("azure-native:healthcareapis/v20231101:IotConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotConnector resources.
type iotConnectorState struct {
}

type IotConnectorState struct {
}

func (IotConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotConnectorState)(nil)).Elem()
}

type iotConnectorArgs struct {
	// Device Mappings.
	DeviceMapping *IotMappingProperties `pulumi:"deviceMapping"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity `pulumi:"identity"`
	// Source configuration.
	IngestionEndpointConfiguration *IotEventHubIngestionEndpointConfiguration `pulumi:"ingestionEndpointConfiguration"`
	// The name of IoT Connector resource.
	IotConnectorName *string `pulumi:"iotConnectorName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IotConnector resource.
type IotConnectorArgs struct {
	// Device Mappings.
	DeviceMapping IotMappingPropertiesPtrInput
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServiceManagedIdentityIdentityPtrInput
	// Source configuration.
	IngestionEndpointConfiguration IotEventHubIngestionEndpointConfigurationPtrInput
	// The name of IoT Connector resource.
	IotConnectorName pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput
}

func (IotConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotConnectorArgs)(nil)).Elem()
}

type IotConnectorInput interface {
	pulumi.Input

	ToIotConnectorOutput() IotConnectorOutput
	ToIotConnectorOutputWithContext(ctx context.Context) IotConnectorOutput
}

func (*IotConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**IotConnector)(nil)).Elem()
}

func (i *IotConnector) ToIotConnectorOutput() IotConnectorOutput {
	return i.ToIotConnectorOutputWithContext(context.Background())
}

func (i *IotConnector) ToIotConnectorOutputWithContext(ctx context.Context) IotConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotConnectorOutput)
}

type IotConnectorOutput struct{ *pulumi.OutputState }

func (IotConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotConnector)(nil)).Elem()
}

func (o IotConnectorOutput) ToIotConnectorOutput() IotConnectorOutput {
	return o
}

func (o IotConnectorOutput) ToIotConnectorOutputWithContext(ctx context.Context) IotConnectorOutput {
	return o
}

// Device Mappings.
func (o IotConnectorOutput) DeviceMapping() IotMappingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *IotConnector) IotMappingPropertiesResponsePtrOutput { return v.DeviceMapping }).(IotMappingPropertiesResponsePtrOutput)
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o IotConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o IotConnectorOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v *IotConnector) ServiceManagedIdentityResponseIdentityPtrOutput { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

// Source configuration.
func (o IotConnectorOutput) IngestionEndpointConfiguration() IotEventHubIngestionEndpointConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *IotConnector) IotEventHubIngestionEndpointConfigurationResponsePtrOutput {
		return v.IngestionEndpointConfiguration
	}).(IotEventHubIngestionEndpointConfigurationResponsePtrOutput)
}

// The resource location.
func (o IotConnectorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o IotConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o IotConnectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o IotConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IotConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o IotConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o IotConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IotConnectorOutput{})
}
