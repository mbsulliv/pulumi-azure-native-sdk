// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of the DigitalTwins service.
// API Version: 2020-12-01.
type DigitalTwin struct {
	pulumi.CustomResourceState

	// Time when DigitalTwinsInstance was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Api endpoint to work with DigitalTwinsInstance.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The managed identity for the DigitalTwinsInstance.
	Identity DigitalTwinsIdentityResponsePtrOutput `pulumi:"identity"`
	// Time when DigitalTwinsInstance was updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Public network access for the DigitalTwinsInstance.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDigitalTwin registers a new resource with the given unique name, arguments, and options.
func NewDigitalTwin(ctx *pulumi.Context,
	name string, args *DigitalTwinArgs, opts ...pulumi.ResourceOption) (*DigitalTwin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:digitaltwins/v20200301preview:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20201031:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20201201:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20210630preview:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20220531:DigitalTwin"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20221031:DigitalTwin"),
		},
	})
	opts = append(opts, aliases)
	var resource DigitalTwin
	err := ctx.RegisterResource("azure-native:digitaltwins:DigitalTwin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDigitalTwin gets an existing DigitalTwin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDigitalTwin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DigitalTwinState, opts ...pulumi.ResourceOption) (*DigitalTwin, error) {
	var resource DigitalTwin
	err := ctx.ReadResource("azure-native:digitaltwins:DigitalTwin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DigitalTwin resources.
type digitalTwinState struct {
}

type DigitalTwinState struct {
}

func (DigitalTwinState) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinState)(nil)).Elem()
}

type digitalTwinArgs struct {
	// The managed identity for the DigitalTwinsInstance.
	Identity *DigitalTwinsIdentity `pulumi:"identity"`
	// The resource location.
	Location                   *string                         `pulumi:"location"`
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	// Public network access for the DigitalTwinsInstance.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName *string `pulumi:"resourceName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DigitalTwin resource.
type DigitalTwinArgs struct {
	// The managed identity for the DigitalTwinsInstance.
	Identity DigitalTwinsIdentityPtrInput
	// The resource location.
	Location                   pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	// Public network access for the DigitalTwinsInstance.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName pulumi.StringInput
	// The name of the DigitalTwinsInstance.
	ResourceName pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (DigitalTwinArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinArgs)(nil)).Elem()
}

type DigitalTwinInput interface {
	pulumi.Input

	ToDigitalTwinOutput() DigitalTwinOutput
	ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput
}

func (*DigitalTwin) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwin)(nil)).Elem()
}

func (i *DigitalTwin) ToDigitalTwinOutput() DigitalTwinOutput {
	return i.ToDigitalTwinOutputWithContext(context.Background())
}

func (i *DigitalTwin) ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinOutput)
}

type DigitalTwinOutput struct{ *pulumi.OutputState }

func (DigitalTwinOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwin)(nil)).Elem()
}

func (o DigitalTwinOutput) ToDigitalTwinOutput() DigitalTwinOutput {
	return o
}

func (o DigitalTwinOutput) ToDigitalTwinOutputWithContext(ctx context.Context) DigitalTwinOutput {
	return o
}

// Time when DigitalTwinsInstance was created.
func (o DigitalTwinOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Api endpoint to work with DigitalTwinsInstance.
func (o DigitalTwinOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The managed identity for the DigitalTwinsInstance.
func (o DigitalTwinOutput) Identity() DigitalTwinsIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DigitalTwin) DigitalTwinsIdentityResponsePtrOutput { return v.Identity }).(DigitalTwinsIdentityResponsePtrOutput)
}

// Time when DigitalTwinsInstance was updated.
func (o DigitalTwinOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// The resource location.
func (o DigitalTwinOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o DigitalTwinOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DigitalTwinOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *DigitalTwin) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state.
func (o DigitalTwinOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Public network access for the DigitalTwinsInstance.
func (o DigitalTwinOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o DigitalTwinOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o DigitalTwinOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DigitalTwin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DigitalTwinOutput{})
}
