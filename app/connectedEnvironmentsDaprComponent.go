// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Dapr Component.
// Azure REST API version: 2022-10-01.
//
// Other available API versions: 2023-04-01-preview, 2023-05-01, 2023-05-02-preview.
type ConnectedEnvironmentsDaprComponent struct {
	pulumi.CustomResourceState

	// Component type
	ComponentType pulumi.StringPtrOutput `pulumi:"componentType"`
	// Boolean describing if the component errors are ignores
	IgnoreErrors pulumi.BoolPtrOutput `pulumi:"ignoreErrors"`
	// Initialization timeout
	InitTimeout pulumi.StringPtrOutput `pulumi:"initTimeout"`
	// Component metadata
	Metadata DaprMetadataResponseArrayOutput `pulumi:"metadata"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Names of container apps that can use this Dapr component
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Name of a Dapr component to retrieve component secrets from
	SecretStoreComponent pulumi.StringPtrOutput `pulumi:"secretStoreComponent"`
	// Collection of secrets used by a Dapr component
	Secrets SecretResponseArrayOutput `pulumi:"secrets"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Component version
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewConnectedEnvironmentsDaprComponent registers a new resource with the given unique name, arguments, and options.
func NewConnectedEnvironmentsDaprComponent(ctx *pulumi.Context,
	name string, args *ConnectedEnvironmentsDaprComponentArgs, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsDaprComponent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectedEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedEnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.IgnoreErrors == nil {
		args.IgnoreErrors = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ConnectedEnvironmentsDaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ConnectedEnvironmentsDaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:ConnectedEnvironmentsDaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230401preview:ConnectedEnvironmentsDaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230501:ConnectedEnvironmentsDaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230502preview:ConnectedEnvironmentsDaprComponent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectedEnvironmentsDaprComponent
	err := ctx.RegisterResource("azure-native:app:ConnectedEnvironmentsDaprComponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectedEnvironmentsDaprComponent gets an existing ConnectedEnvironmentsDaprComponent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectedEnvironmentsDaprComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedEnvironmentsDaprComponentState, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsDaprComponent, error) {
	var resource ConnectedEnvironmentsDaprComponent
	err := ctx.ReadResource("azure-native:app:ConnectedEnvironmentsDaprComponent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectedEnvironmentsDaprComponent resources.
type connectedEnvironmentsDaprComponentState struct {
}

type ConnectedEnvironmentsDaprComponentState struct {
}

func (ConnectedEnvironmentsDaprComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsDaprComponentState)(nil)).Elem()
}

type connectedEnvironmentsDaprComponentArgs struct {
	// Name of the Dapr Component.
	ComponentName *string `pulumi:"componentName"`
	// Component type
	ComponentType *string `pulumi:"componentType"`
	// Name of the connected environment.
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	// Boolean describing if the component errors are ignores
	IgnoreErrors *bool `pulumi:"ignoreErrors"`
	// Initialization timeout
	InitTimeout *string `pulumi:"initTimeout"`
	// Component metadata
	Metadata []DaprMetadata `pulumi:"metadata"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Names of container apps that can use this Dapr component
	Scopes []string `pulumi:"scopes"`
	// Name of a Dapr component to retrieve component secrets from
	SecretStoreComponent *string `pulumi:"secretStoreComponent"`
	// Collection of secrets used by a Dapr component
	Secrets []Secret `pulumi:"secrets"`
	// Component version
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a ConnectedEnvironmentsDaprComponent resource.
type ConnectedEnvironmentsDaprComponentArgs struct {
	// Name of the Dapr Component.
	ComponentName pulumi.StringPtrInput
	// Component type
	ComponentType pulumi.StringPtrInput
	// Name of the connected environment.
	ConnectedEnvironmentName pulumi.StringInput
	// Boolean describing if the component errors are ignores
	IgnoreErrors pulumi.BoolPtrInput
	// Initialization timeout
	InitTimeout pulumi.StringPtrInput
	// Component metadata
	Metadata DaprMetadataArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Names of container apps that can use this Dapr component
	Scopes pulumi.StringArrayInput
	// Name of a Dapr component to retrieve component secrets from
	SecretStoreComponent pulumi.StringPtrInput
	// Collection of secrets used by a Dapr component
	Secrets SecretArrayInput
	// Component version
	Version pulumi.StringPtrInput
}

func (ConnectedEnvironmentsDaprComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsDaprComponentArgs)(nil)).Elem()
}

type ConnectedEnvironmentsDaprComponentInput interface {
	pulumi.Input

	ToConnectedEnvironmentsDaprComponentOutput() ConnectedEnvironmentsDaprComponentOutput
	ToConnectedEnvironmentsDaprComponentOutputWithContext(ctx context.Context) ConnectedEnvironmentsDaprComponentOutput
}

func (*ConnectedEnvironmentsDaprComponent) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsDaprComponent)(nil)).Elem()
}

func (i *ConnectedEnvironmentsDaprComponent) ToConnectedEnvironmentsDaprComponentOutput() ConnectedEnvironmentsDaprComponentOutput {
	return i.ToConnectedEnvironmentsDaprComponentOutputWithContext(context.Background())
}

func (i *ConnectedEnvironmentsDaprComponent) ToConnectedEnvironmentsDaprComponentOutputWithContext(ctx context.Context) ConnectedEnvironmentsDaprComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedEnvironmentsDaprComponentOutput)
}

func (i *ConnectedEnvironmentsDaprComponent) ToOutput(ctx context.Context) pulumix.Output[*ConnectedEnvironmentsDaprComponent] {
	return pulumix.Output[*ConnectedEnvironmentsDaprComponent]{
		OutputState: i.ToConnectedEnvironmentsDaprComponentOutputWithContext(ctx).OutputState,
	}
}

type ConnectedEnvironmentsDaprComponentOutput struct{ *pulumi.OutputState }

func (ConnectedEnvironmentsDaprComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsDaprComponent)(nil)).Elem()
}

func (o ConnectedEnvironmentsDaprComponentOutput) ToConnectedEnvironmentsDaprComponentOutput() ConnectedEnvironmentsDaprComponentOutput {
	return o
}

func (o ConnectedEnvironmentsDaprComponentOutput) ToConnectedEnvironmentsDaprComponentOutputWithContext(ctx context.Context) ConnectedEnvironmentsDaprComponentOutput {
	return o
}

func (o ConnectedEnvironmentsDaprComponentOutput) ToOutput(ctx context.Context) pulumix.Output[*ConnectedEnvironmentsDaprComponent] {
	return pulumix.Output[*ConnectedEnvironmentsDaprComponent]{
		OutputState: o.OutputState,
	}
}

// Component type
func (o ConnectedEnvironmentsDaprComponentOutput) ComponentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringPtrOutput { return v.ComponentType }).(pulumi.StringPtrOutput)
}

// Boolean describing if the component errors are ignores
func (o ConnectedEnvironmentsDaprComponentOutput) IgnoreErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.BoolPtrOutput { return v.IgnoreErrors }).(pulumi.BoolPtrOutput)
}

// Initialization timeout
func (o ConnectedEnvironmentsDaprComponentOutput) InitTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringPtrOutput { return v.InitTimeout }).(pulumi.StringPtrOutput)
}

// Component metadata
func (o ConnectedEnvironmentsDaprComponentOutput) Metadata() DaprMetadataResponseArrayOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) DaprMetadataResponseArrayOutput { return v.Metadata }).(DaprMetadataResponseArrayOutput)
}

// The name of the resource
func (o ConnectedEnvironmentsDaprComponentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Names of container apps that can use this Dapr component
func (o ConnectedEnvironmentsDaprComponentOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Name of a Dapr component to retrieve component secrets from
func (o ConnectedEnvironmentsDaprComponentOutput) SecretStoreComponent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringPtrOutput { return v.SecretStoreComponent }).(pulumi.StringPtrOutput)
}

// Collection of secrets used by a Dapr component
func (o ConnectedEnvironmentsDaprComponentOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) SecretResponseArrayOutput { return v.Secrets }).(SecretResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConnectedEnvironmentsDaprComponentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConnectedEnvironmentsDaprComponentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Component version
func (o ConnectedEnvironmentsDaprComponentOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedEnvironmentsDaprComponentOutput{})
}
