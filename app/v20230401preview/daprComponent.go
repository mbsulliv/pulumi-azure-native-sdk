// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Dapr Component.
type DaprComponent struct {
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

// NewDaprComponent registers a new resource with the given unique name, arguments, and options.
func NewDaprComponent(ctx *pulumi.Context,
	name string, args *DaprComponentArgs, opts ...pulumi.ResourceOption) (*DaprComponent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.IgnoreErrors == nil {
		args.IgnoreErrors = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220101preview:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230501:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230502preview:DaprComponent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DaprComponent
	err := ctx.RegisterResource("azure-native:app/v20230401preview:DaprComponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDaprComponent gets an existing DaprComponent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDaprComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DaprComponentState, opts ...pulumi.ResourceOption) (*DaprComponent, error) {
	var resource DaprComponent
	err := ctx.ReadResource("azure-native:app/v20230401preview:DaprComponent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DaprComponent resources.
type daprComponentState struct {
}

type DaprComponentState struct {
}

func (DaprComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*daprComponentState)(nil)).Elem()
}

type daprComponentArgs struct {
	// Name of the Dapr Component.
	ComponentName *string `pulumi:"componentName"`
	// Component type
	ComponentType *string `pulumi:"componentType"`
	// Name of the Managed Environment.
	EnvironmentName string `pulumi:"environmentName"`
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

// The set of arguments for constructing a DaprComponent resource.
type DaprComponentArgs struct {
	// Name of the Dapr Component.
	ComponentName pulumi.StringPtrInput
	// Component type
	ComponentType pulumi.StringPtrInput
	// Name of the Managed Environment.
	EnvironmentName pulumi.StringInput
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

func (DaprComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*daprComponentArgs)(nil)).Elem()
}

type DaprComponentInput interface {
	pulumi.Input

	ToDaprComponentOutput() DaprComponentOutput
	ToDaprComponentOutputWithContext(ctx context.Context) DaprComponentOutput
}

func (*DaprComponent) ElementType() reflect.Type {
	return reflect.TypeOf((**DaprComponent)(nil)).Elem()
}

func (i *DaprComponent) ToDaprComponentOutput() DaprComponentOutput {
	return i.ToDaprComponentOutputWithContext(context.Background())
}

func (i *DaprComponent) ToDaprComponentOutputWithContext(ctx context.Context) DaprComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprComponentOutput)
}

func (i *DaprComponent) ToOutput(ctx context.Context) pulumix.Output[*DaprComponent] {
	return pulumix.Output[*DaprComponent]{
		OutputState: i.ToDaprComponentOutputWithContext(ctx).OutputState,
	}
}

type DaprComponentOutput struct{ *pulumi.OutputState }

func (DaprComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaprComponent)(nil)).Elem()
}

func (o DaprComponentOutput) ToDaprComponentOutput() DaprComponentOutput {
	return o
}

func (o DaprComponentOutput) ToDaprComponentOutputWithContext(ctx context.Context) DaprComponentOutput {
	return o
}

func (o DaprComponentOutput) ToOutput(ctx context.Context) pulumix.Output[*DaprComponent] {
	return pulumix.Output[*DaprComponent]{
		OutputState: o.OutputState,
	}
}

// Component type
func (o DaprComponentOutput) ComponentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringPtrOutput { return v.ComponentType }).(pulumi.StringPtrOutput)
}

// Boolean describing if the component errors are ignores
func (o DaprComponentOutput) IgnoreErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.BoolPtrOutput { return v.IgnoreErrors }).(pulumi.BoolPtrOutput)
}

// Initialization timeout
func (o DaprComponentOutput) InitTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringPtrOutput { return v.InitTimeout }).(pulumi.StringPtrOutput)
}

// Component metadata
func (o DaprComponentOutput) Metadata() DaprMetadataResponseArrayOutput {
	return o.ApplyT(func(v *DaprComponent) DaprMetadataResponseArrayOutput { return v.Metadata }).(DaprMetadataResponseArrayOutput)
}

// The name of the resource
func (o DaprComponentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Names of container apps that can use this Dapr component
func (o DaprComponentOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Name of a Dapr component to retrieve component secrets from
func (o DaprComponentOutput) SecretStoreComponent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringPtrOutput { return v.SecretStoreComponent }).(pulumi.StringPtrOutput)
}

// Collection of secrets used by a Dapr component
func (o DaprComponentOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v *DaprComponent) SecretResponseArrayOutput { return v.Secrets }).(SecretResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DaprComponentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DaprComponent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DaprComponentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Component version
func (o DaprComponentOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DaprComponentOutput{})
}
