// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Binding resource payload
type Binding struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Binding resource
	Properties BindingResourcePropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBinding registers a new resource with the given unique name, arguments, and options.
func NewBinding(ctx *pulumi.Context,
	name string, args *BindingArgs, opts ...pulumi.ResourceOption) (*Binding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:Binding"),
		},
	})
	opts = append(opts, aliases)
	var resource Binding
	err := ctx.RegisterResource("azure-native:appplatform/v20210601preview:Binding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBinding gets an existing Binding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingState, opts ...pulumi.ResourceOption) (*Binding, error) {
	var resource Binding
	err := ctx.ReadResource("azure-native:appplatform/v20210601preview:Binding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Binding resources.
type bindingState struct {
}

type BindingState struct {
}

func (BindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingState)(nil)).Elem()
}

type bindingArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the Binding resource.
	BindingName *string `pulumi:"bindingName"`
	// Properties of the Binding resource
	Properties *BindingResourceProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Binding resource.
type BindingArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput
	// The name of the Binding resource.
	BindingName pulumi.StringPtrInput
	// Properties of the Binding resource
	Properties BindingResourcePropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (BindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingArgs)(nil)).Elem()
}

type BindingInput interface {
	pulumi.Input

	ToBindingOutput() BindingOutput
	ToBindingOutputWithContext(ctx context.Context) BindingOutput
}

func (*Binding) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (i *Binding) ToBindingOutput() BindingOutput {
	return i.ToBindingOutputWithContext(context.Background())
}

func (i *Binding) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingOutput)
}

type BindingOutput struct{ *pulumi.OutputState }

func (BindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (o BindingOutput) ToBindingOutput() BindingOutput {
	return o
}

func (o BindingOutput) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return o
}

// The name of the resource.
func (o BindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the Binding resource
func (o BindingOutput) Properties() BindingResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *Binding) BindingResourcePropertiesResponseOutput { return v.Properties }).(BindingResourcePropertiesResponseOutput)
}

// The type of the resource.
func (o BindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BindingOutput{})
}
