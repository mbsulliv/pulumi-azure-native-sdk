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

// Dev Tool Portal resource
type DevToolPortal struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Dev Tool Portal properties payload
	Properties DevToolPortalPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDevToolPortal registers a new resource with the given unique name, arguments, and options.
func NewDevToolPortal(ctx *pulumi.Context,
	name string, args *DevToolPortalArgs, opts ...pulumi.ResourceOption) (*DevToolPortal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToDevToolPortalPropertiesPtrOutput().ApplyT(func(v *DevToolPortalProperties) *DevToolPortalProperties { return v.Defaults() }).(DevToolPortalPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:DevToolPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:DevToolPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:DevToolPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:DevToolPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:DevToolPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:DevToolPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:DevToolPortal"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DevToolPortal
	err := ctx.RegisterResource("azure-native:appplatform/v20230501preview:DevToolPortal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevToolPortal gets an existing DevToolPortal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevToolPortal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevToolPortalState, opts ...pulumi.ResourceOption) (*DevToolPortal, error) {
	var resource DevToolPortal
	err := ctx.ReadResource("azure-native:appplatform/v20230501preview:DevToolPortal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DevToolPortal resources.
type devToolPortalState struct {
}

type DevToolPortalState struct {
}

func (DevToolPortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*devToolPortalState)(nil)).Elem()
}

type devToolPortalArgs struct {
	// The name of Dev Tool Portal.
	DevToolPortalName *string `pulumi:"devToolPortalName"`
	// Dev Tool Portal properties payload
	Properties *DevToolPortalProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a DevToolPortal resource.
type DevToolPortalArgs struct {
	// The name of Dev Tool Portal.
	DevToolPortalName pulumi.StringPtrInput
	// Dev Tool Portal properties payload
	Properties DevToolPortalPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (DevToolPortalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devToolPortalArgs)(nil)).Elem()
}

type DevToolPortalInput interface {
	pulumi.Input

	ToDevToolPortalOutput() DevToolPortalOutput
	ToDevToolPortalOutputWithContext(ctx context.Context) DevToolPortalOutput
}

func (*DevToolPortal) ElementType() reflect.Type {
	return reflect.TypeOf((**DevToolPortal)(nil)).Elem()
}

func (i *DevToolPortal) ToDevToolPortalOutput() DevToolPortalOutput {
	return i.ToDevToolPortalOutputWithContext(context.Background())
}

func (i *DevToolPortal) ToDevToolPortalOutputWithContext(ctx context.Context) DevToolPortalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevToolPortalOutput)
}

func (i *DevToolPortal) ToOutput(ctx context.Context) pulumix.Output[*DevToolPortal] {
	return pulumix.Output[*DevToolPortal]{
		OutputState: i.ToDevToolPortalOutputWithContext(ctx).OutputState,
	}
}

type DevToolPortalOutput struct{ *pulumi.OutputState }

func (DevToolPortalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevToolPortal)(nil)).Elem()
}

func (o DevToolPortalOutput) ToDevToolPortalOutput() DevToolPortalOutput {
	return o
}

func (o DevToolPortalOutput) ToDevToolPortalOutputWithContext(ctx context.Context) DevToolPortalOutput {
	return o
}

func (o DevToolPortalOutput) ToOutput(ctx context.Context) pulumix.Output[*DevToolPortal] {
	return pulumix.Output[*DevToolPortal]{
		OutputState: o.OutputState,
	}
}

// The name of the resource.
func (o DevToolPortalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevToolPortal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Dev Tool Portal properties payload
func (o DevToolPortalOutput) Properties() DevToolPortalPropertiesResponseOutput {
	return o.ApplyT(func(v *DevToolPortal) DevToolPortalPropertiesResponseOutput { return v.Properties }).(DevToolPortalPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DevToolPortalOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DevToolPortal) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o DevToolPortalOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DevToolPortal) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DevToolPortalOutput{})
}
