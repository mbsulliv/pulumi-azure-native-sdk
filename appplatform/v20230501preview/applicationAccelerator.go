// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Application accelerator resource
type ApplicationAccelerator struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Application accelerator properties payload
	Properties ApplicationAcceleratorPropertiesResponseOutput `pulumi:"properties"`
	// Sku of the application accelerator resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationAccelerator registers a new resource with the given unique name, arguments, and options.
func NewApplicationAccelerator(ctx *pulumi.Context,
	name string, args *ApplicationAcceleratorArgs, opts ...pulumi.ResourceOption) (*ApplicationAccelerator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:appplatform:ApplicationAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ApplicationAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:ApplicationAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:ApplicationAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:ApplicationAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:ApplicationAccelerator"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:ApplicationAccelerator"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationAccelerator
	err := ctx.RegisterResource("azure-native:appplatform/v20230501preview:ApplicationAccelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationAccelerator gets an existing ApplicationAccelerator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationAcceleratorState, opts ...pulumi.ResourceOption) (*ApplicationAccelerator, error) {
	var resource ApplicationAccelerator
	err := ctx.ReadResource("azure-native:appplatform/v20230501preview:ApplicationAccelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationAccelerator resources.
type applicationAcceleratorState struct {
}

type ApplicationAcceleratorState struct {
}

func (ApplicationAcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationAcceleratorState)(nil)).Elem()
}

type applicationAcceleratorArgs struct {
	// The name of the application accelerator.
	ApplicationAcceleratorName *string `pulumi:"applicationAcceleratorName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// Sku of the application accelerator resource
	Sku *Sku `pulumi:"sku"`
}

// The set of arguments for constructing a ApplicationAccelerator resource.
type ApplicationAcceleratorArgs struct {
	// The name of the application accelerator.
	ApplicationAcceleratorName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
	// Sku of the application accelerator resource
	Sku SkuPtrInput
}

func (ApplicationAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationAcceleratorArgs)(nil)).Elem()
}

type ApplicationAcceleratorInput interface {
	pulumi.Input

	ToApplicationAcceleratorOutput() ApplicationAcceleratorOutput
	ToApplicationAcceleratorOutputWithContext(ctx context.Context) ApplicationAcceleratorOutput
}

func (*ApplicationAccelerator) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationAccelerator)(nil)).Elem()
}

func (i *ApplicationAccelerator) ToApplicationAcceleratorOutput() ApplicationAcceleratorOutput {
	return i.ToApplicationAcceleratorOutputWithContext(context.Background())
}

func (i *ApplicationAccelerator) ToApplicationAcceleratorOutputWithContext(ctx context.Context) ApplicationAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAcceleratorOutput)
}

func (i *ApplicationAccelerator) ToOutput(ctx context.Context) pulumix.Output[*ApplicationAccelerator] {
	return pulumix.Output[*ApplicationAccelerator]{
		OutputState: i.ToApplicationAcceleratorOutputWithContext(ctx).OutputState,
	}
}

type ApplicationAcceleratorOutput struct{ *pulumi.OutputState }

func (ApplicationAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationAccelerator)(nil)).Elem()
}

func (o ApplicationAcceleratorOutput) ToApplicationAcceleratorOutput() ApplicationAcceleratorOutput {
	return o
}

func (o ApplicationAcceleratorOutput) ToApplicationAcceleratorOutputWithContext(ctx context.Context) ApplicationAcceleratorOutput {
	return o
}

func (o ApplicationAcceleratorOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationAccelerator] {
	return pulumix.Output[*ApplicationAccelerator]{
		OutputState: o.OutputState,
	}
}

// The name of the resource.
func (o ApplicationAcceleratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Application accelerator properties payload
func (o ApplicationAcceleratorOutput) Properties() ApplicationAcceleratorPropertiesResponseOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) ApplicationAcceleratorPropertiesResponseOutput { return v.Properties }).(ApplicationAcceleratorPropertiesResponseOutput)
}

// Sku of the application accelerator resource
func (o ApplicationAcceleratorOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ApplicationAcceleratorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ApplicationAcceleratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAccelerator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationAcceleratorOutput{})
}
