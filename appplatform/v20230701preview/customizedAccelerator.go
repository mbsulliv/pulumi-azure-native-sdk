// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Customized accelerator resource
type CustomizedAccelerator struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Customized accelerator properties payload
	Properties CustomizedAcceleratorPropertiesResponseOutput `pulumi:"properties"`
	// Sku of the customized accelerator resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomizedAccelerator registers a new resource with the given unique name, arguments, and options.
func NewCustomizedAccelerator(ctx *pulumi.Context,
	name string, args *CustomizedAcceleratorArgs, opts ...pulumi.ResourceOption) (*CustomizedAccelerator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationAcceleratorName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationAcceleratorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:CustomizedAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:CustomizedAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:CustomizedAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:CustomizedAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230501preview:CustomizedAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:CustomizedAccelerator"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomizedAccelerator
	err := ctx.RegisterResource("azure-native:appplatform/v20230701preview:CustomizedAccelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomizedAccelerator gets an existing CustomizedAccelerator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomizedAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomizedAcceleratorState, opts ...pulumi.ResourceOption) (*CustomizedAccelerator, error) {
	var resource CustomizedAccelerator
	err := ctx.ReadResource("azure-native:appplatform/v20230701preview:CustomizedAccelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomizedAccelerator resources.
type customizedAcceleratorState struct {
}

type CustomizedAcceleratorState struct {
}

func (CustomizedAcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*customizedAcceleratorState)(nil)).Elem()
}

type customizedAcceleratorArgs struct {
	// The name of the application accelerator.
	ApplicationAcceleratorName string `pulumi:"applicationAcceleratorName"`
	// The name of the customized accelerator.
	CustomizedAcceleratorName *string `pulumi:"customizedAcceleratorName"`
	// Customized accelerator properties payload
	Properties *CustomizedAcceleratorProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// Sku of the customized accelerator resource
	Sku *Sku `pulumi:"sku"`
}

// The set of arguments for constructing a CustomizedAccelerator resource.
type CustomizedAcceleratorArgs struct {
	// The name of the application accelerator.
	ApplicationAcceleratorName pulumi.StringInput
	// The name of the customized accelerator.
	CustomizedAcceleratorName pulumi.StringPtrInput
	// Customized accelerator properties payload
	Properties CustomizedAcceleratorPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
	// Sku of the customized accelerator resource
	Sku SkuPtrInput
}

func (CustomizedAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customizedAcceleratorArgs)(nil)).Elem()
}

type CustomizedAcceleratorInput interface {
	pulumi.Input

	ToCustomizedAcceleratorOutput() CustomizedAcceleratorOutput
	ToCustomizedAcceleratorOutputWithContext(ctx context.Context) CustomizedAcceleratorOutput
}

func (*CustomizedAccelerator) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizedAccelerator)(nil)).Elem()
}

func (i *CustomizedAccelerator) ToCustomizedAcceleratorOutput() CustomizedAcceleratorOutput {
	return i.ToCustomizedAcceleratorOutputWithContext(context.Background())
}

func (i *CustomizedAccelerator) ToCustomizedAcceleratorOutputWithContext(ctx context.Context) CustomizedAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizedAcceleratorOutput)
}

func (i *CustomizedAccelerator) ToOutput(ctx context.Context) pulumix.Output[*CustomizedAccelerator] {
	return pulumix.Output[*CustomizedAccelerator]{
		OutputState: i.ToCustomizedAcceleratorOutputWithContext(ctx).OutputState,
	}
}

type CustomizedAcceleratorOutput struct{ *pulumi.OutputState }

func (CustomizedAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizedAccelerator)(nil)).Elem()
}

func (o CustomizedAcceleratorOutput) ToCustomizedAcceleratorOutput() CustomizedAcceleratorOutput {
	return o
}

func (o CustomizedAcceleratorOutput) ToCustomizedAcceleratorOutputWithContext(ctx context.Context) CustomizedAcceleratorOutput {
	return o
}

func (o CustomizedAcceleratorOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomizedAccelerator] {
	return pulumix.Output[*CustomizedAccelerator]{
		OutputState: o.OutputState,
	}
}

// The name of the resource.
func (o CustomizedAcceleratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Customized accelerator properties payload
func (o CustomizedAcceleratorOutput) Properties() CustomizedAcceleratorPropertiesResponseOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) CustomizedAcceleratorPropertiesResponseOutput { return v.Properties }).(CustomizedAcceleratorPropertiesResponseOutput)
}

// Sku of the customized accelerator resource
func (o CustomizedAcceleratorOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o CustomizedAcceleratorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o CustomizedAcceleratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedAccelerator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomizedAcceleratorOutput{})
}
