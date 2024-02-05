// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Application Live View resource
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-01-01-preview.
type ApplicationLiveView struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Application Live View properties payload
	Properties ApplicationLiveViewPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationLiveView registers a new resource with the given unique name, arguments, and options.
func NewApplicationLiveView(ctx *pulumi.Context,
	name string, args *ApplicationLiveViewArgs, opts ...pulumi.ResourceOption) (*ApplicationLiveView, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ApplicationLiveView"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:ApplicationLiveView"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:ApplicationLiveView"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230501preview:ApplicationLiveView"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:ApplicationLiveView"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:ApplicationLiveView"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:ApplicationLiveView"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231201:ApplicationLiveView"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20240101preview:ApplicationLiveView"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationLiveView
	err := ctx.RegisterResource("azure-native:appplatform:ApplicationLiveView", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationLiveView gets an existing ApplicationLiveView resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationLiveView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationLiveViewState, opts ...pulumi.ResourceOption) (*ApplicationLiveView, error) {
	var resource ApplicationLiveView
	err := ctx.ReadResource("azure-native:appplatform:ApplicationLiveView", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationLiveView resources.
type applicationLiveViewState struct {
}

type ApplicationLiveViewState struct {
}

func (ApplicationLiveViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationLiveViewState)(nil)).Elem()
}

type applicationLiveViewArgs struct {
	// The name of Application Live View.
	ApplicationLiveViewName *string `pulumi:"applicationLiveViewName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ApplicationLiveView resource.
type ApplicationLiveViewArgs struct {
	// The name of Application Live View.
	ApplicationLiveViewName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (ApplicationLiveViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationLiveViewArgs)(nil)).Elem()
}

type ApplicationLiveViewInput interface {
	pulumi.Input

	ToApplicationLiveViewOutput() ApplicationLiveViewOutput
	ToApplicationLiveViewOutputWithContext(ctx context.Context) ApplicationLiveViewOutput
}

func (*ApplicationLiveView) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLiveView)(nil)).Elem()
}

func (i *ApplicationLiveView) ToApplicationLiveViewOutput() ApplicationLiveViewOutput {
	return i.ToApplicationLiveViewOutputWithContext(context.Background())
}

func (i *ApplicationLiveView) ToApplicationLiveViewOutputWithContext(ctx context.Context) ApplicationLiveViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLiveViewOutput)
}

type ApplicationLiveViewOutput struct{ *pulumi.OutputState }

func (ApplicationLiveViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLiveView)(nil)).Elem()
}

func (o ApplicationLiveViewOutput) ToApplicationLiveViewOutput() ApplicationLiveViewOutput {
	return o
}

func (o ApplicationLiveViewOutput) ToApplicationLiveViewOutputWithContext(ctx context.Context) ApplicationLiveViewOutput {
	return o
}

// The name of the resource.
func (o ApplicationLiveViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationLiveView) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Application Live View properties payload
func (o ApplicationLiveViewOutput) Properties() ApplicationLiveViewPropertiesResponseOutput {
	return o.ApplyT(func(v *ApplicationLiveView) ApplicationLiveViewPropertiesResponseOutput { return v.Properties }).(ApplicationLiveViewPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ApplicationLiveViewOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationLiveView) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ApplicationLiveViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationLiveView) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationLiveViewOutput{})
}
