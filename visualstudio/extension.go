// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package visualstudio

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The response to an extension resource GET request.
// Azure REST API version: 2017-11-01-preview. Prior API version in Azure Native 1.x: 2014-04-01-preview
type Extension struct {
	pulumi.CustomResourceState

	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The extension plan that was purchased.
	Plan ExtensionResourcePlanResponsePtrOutput `pulumi:"plan"`
	// Resource properties.
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExtension registers a new resource with the given unique name, arguments, and options.
func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountResourceName == nil {
		return nil, errors.New("invalid value for required argument 'AccountResourceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:visualstudio/v20140401preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:visualstudio/v20171101preview:Extension"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Extension
	err := ctx.RegisterResource("azure-native:visualstudio:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtension gets an existing Extension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:visualstudio:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extension resources.
type extensionState struct {
}

type ExtensionState struct {
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	// The name of the Visual Studio Team Services account resource.
	AccountResourceName string `pulumi:"accountResourceName"`
	// The name of the extension.
	ExtensionResourceName *string `pulumi:"extensionResourceName"`
	// The Azure region of the Visual Studio account associated with this request (i.e 'southcentralus'.)
	Location *string `pulumi:"location"`
	// Extended information about the plan being purchased for this extension resource.
	Plan *ExtensionResourcePlan `pulumi:"plan"`
	// A dictionary of extended properties. This property is currently unused.
	Properties map[string]string `pulumi:"properties"`
	// Name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A dictionary of user-defined tags to be stored with the extension resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Extension resource.
type ExtensionArgs struct {
	// The name of the Visual Studio Team Services account resource.
	AccountResourceName pulumi.StringInput
	// The name of the extension.
	ExtensionResourceName pulumi.StringPtrInput
	// The Azure region of the Visual Studio account associated with this request (i.e 'southcentralus'.)
	Location pulumi.StringPtrInput
	// Extended information about the plan being purchased for this extension resource.
	Plan ExtensionResourcePlanPtrInput
	// A dictionary of extended properties. This property is currently unused.
	Properties pulumi.StringMapInput
	// Name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// A dictionary of user-defined tags to be stored with the extension resource.
	Tags pulumi.StringMapInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput
}

func (*Extension) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

func (i *Extension) ToOutput(ctx context.Context) pulumix.Output[*Extension] {
	return pulumix.Output[*Extension]{
		OutputState: i.ToExtensionOutputWithContext(ctx).OutputState,
	}
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToOutput(ctx context.Context) pulumix.Output[*Extension] {
	return pulumix.Output[*Extension]{
		OutputState: o.OutputState,
	}
}

// Resource location.
func (o ExtensionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The extension plan that was purchased.
func (o ExtensionOutput) Plan() ExtensionResourcePlanResponsePtrOutput {
	return o.ApplyT(func(v *Extension) ExtensionResourcePlanResponsePtrOutput { return v.Plan }).(ExtensionResourcePlanResponsePtrOutput)
}

// Resource properties.
func (o ExtensionOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource tags.
func (o ExtensionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}
