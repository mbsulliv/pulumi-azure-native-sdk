// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devspaces

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2019-04-01. Prior API version in Azure Native 1.x: 2019-04-01.
type Controller struct {
	pulumi.CustomResourceState

	// DNS name for accessing DataPlane services
	DataPlaneFqdn pulumi.StringOutput `pulumi:"dataPlaneFqdn"`
	// DNS suffix for public endpoints running in the Azure Dev Spaces Controller.
	HostSuffix pulumi.StringOutput `pulumi:"hostSuffix"`
	// Region where the Azure resource is located.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Azure Dev Spaces Controller.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Tags for the Azure resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// DNS of the target container host's API server
	TargetContainerHostApiServerFqdn pulumi.StringOutput `pulumi:"targetContainerHostApiServerFqdn"`
	// Resource ID of the target container host
	TargetContainerHostResourceId pulumi.StringOutput `pulumi:"targetContainerHostResourceId"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewController registers a new resource with the given unique name, arguments, and options.
func NewController(ctx *pulumi.Context,
	name string, args *ControllerArgs, opts ...pulumi.ResourceOption) (*Controller, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.TargetContainerHostCredentialsBase64 == nil {
		return nil, errors.New("invalid value for required argument 'TargetContainerHostCredentialsBase64'")
	}
	if args.TargetContainerHostResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetContainerHostResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devspaces/v20190401:Controller"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Controller
	err := ctx.RegisterResource("azure-native:devspaces:Controller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetController gets an existing Controller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControllerState, opts ...pulumi.ResourceOption) (*Controller, error) {
	var resource Controller
	err := ctx.ReadResource("azure-native:devspaces:Controller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Controller resources.
type controllerState struct {
}

type ControllerState struct {
}

func (ControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerState)(nil)).Elem()
}

type controllerArgs struct {
	// Region where the Azure resource is located.
	Location *string `pulumi:"location"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku Sku `pulumi:"sku"`
	// Tags for the Azure resource.
	Tags map[string]string `pulumi:"tags"`
	// Credentials of the target container host (base64).
	TargetContainerHostCredentialsBase64 string `pulumi:"targetContainerHostCredentialsBase64"`
	// Resource ID of the target container host
	TargetContainerHostResourceId string `pulumi:"targetContainerHostResourceId"`
}

// The set of arguments for constructing a Controller resource.
type ControllerArgs struct {
	// Region where the Azure resource is located.
	Location pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku SkuInput
	// Tags for the Azure resource.
	Tags pulumi.StringMapInput
	// Credentials of the target container host (base64).
	TargetContainerHostCredentialsBase64 pulumi.StringInput
	// Resource ID of the target container host
	TargetContainerHostResourceId pulumi.StringInput
}

func (ControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerArgs)(nil)).Elem()
}

type ControllerInput interface {
	pulumi.Input

	ToControllerOutput() ControllerOutput
	ToControllerOutputWithContext(ctx context.Context) ControllerOutput
}

func (*Controller) ElementType() reflect.Type {
	return reflect.TypeOf((**Controller)(nil)).Elem()
}

func (i *Controller) ToControllerOutput() ControllerOutput {
	return i.ToControllerOutputWithContext(context.Background())
}

func (i *Controller) ToControllerOutputWithContext(ctx context.Context) ControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerOutput)
}

func (i *Controller) ToOutput(ctx context.Context) pulumix.Output[*Controller] {
	return pulumix.Output[*Controller]{
		OutputState: i.ToControllerOutputWithContext(ctx).OutputState,
	}
}

type ControllerOutput struct{ *pulumi.OutputState }

func (ControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Controller)(nil)).Elem()
}

func (o ControllerOutput) ToControllerOutput() ControllerOutput {
	return o
}

func (o ControllerOutput) ToControllerOutputWithContext(ctx context.Context) ControllerOutput {
	return o
}

func (o ControllerOutput) ToOutput(ctx context.Context) pulumix.Output[*Controller] {
	return pulumix.Output[*Controller]{
		OutputState: o.OutputState,
	}
}

// DNS name for accessing DataPlane services
func (o ControllerOutput) DataPlaneFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringOutput { return v.DataPlaneFqdn }).(pulumi.StringOutput)
}

// DNS suffix for public endpoints running in the Azure Dev Spaces Controller.
func (o ControllerOutput) HostSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringOutput { return v.HostSuffix }).(pulumi.StringOutput)
}

// Region where the Azure resource is located.
func (o ControllerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o ControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Azure Dev Spaces Controller.
func (o ControllerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Model representing SKU for Azure Dev Spaces Controller.
func (o ControllerOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Controller) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Tags for the Azure resource.
func (o ControllerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// DNS of the target container host's API server
func (o ControllerOutput) TargetContainerHostApiServerFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringOutput { return v.TargetContainerHostApiServerFqdn }).(pulumi.StringOutput)
}

// Resource ID of the target container host
func (o ControllerOutput) TargetContainerHostResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringOutput { return v.TargetContainerHostResourceId }).(pulumi.StringOutput)
}

// The type of the resource.
func (o ControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Controller) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ControllerOutput{})
}
