// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Base class for backup ProtectionIntent.
type ProtectionIntent struct {
	pulumi.CustomResourceState

	// Optional ETag.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ProtectionIntentResource properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProtectionIntent registers a new resource with the given unique name, arguments, and options.
func NewProtectionIntent(ctx *pulumi.Context,
	name string, args *ProtectionIntentArgs, opts ...pulumi.ResourceOption) (*ProtectionIntent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20170701:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220901preview:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220930preview:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230101:ProtectionIntent"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230201:ProtectionIntent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProtectionIntent
	err := ctx.RegisterResource("azure-native:recoveryservices/v20230401:ProtectionIntent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectionIntent gets an existing ProtectionIntent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectionIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionIntentState, opts ...pulumi.ResourceOption) (*ProtectionIntent, error) {
	var resource ProtectionIntent
	err := ctx.ReadResource("azure-native:recoveryservices/v20230401:ProtectionIntent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectionIntent resources.
type protectionIntentState struct {
}

type ProtectionIntentState struct {
}

func (ProtectionIntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionIntentState)(nil)).Elem()
}

type protectionIntentArgs struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Fabric name associated with the backup item.
	FabricName string `pulumi:"fabricName"`
	// Intent object name.
	IntentObjectName *string `pulumi:"intentObjectName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// ProtectionIntentResource properties
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a ProtectionIntent resource.
type ProtectionIntentArgs struct {
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Fabric name associated with the backup item.
	FabricName pulumi.StringInput
	// Intent object name.
	IntentObjectName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// ProtectionIntentResource properties
	Properties pulumi.Input
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the recovery services vault.
	VaultName pulumi.StringInput
}

func (ProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionIntentArgs)(nil)).Elem()
}

type ProtectionIntentInput interface {
	pulumi.Input

	ToProtectionIntentOutput() ProtectionIntentOutput
	ToProtectionIntentOutputWithContext(ctx context.Context) ProtectionIntentOutput
}

func (*ProtectionIntent) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionIntent)(nil)).Elem()
}

func (i *ProtectionIntent) ToProtectionIntentOutput() ProtectionIntentOutput {
	return i.ToProtectionIntentOutputWithContext(context.Background())
}

func (i *ProtectionIntent) ToProtectionIntentOutputWithContext(ctx context.Context) ProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionIntentOutput)
}

func (i *ProtectionIntent) ToOutput(ctx context.Context) pulumix.Output[*ProtectionIntent] {
	return pulumix.Output[*ProtectionIntent]{
		OutputState: i.ToProtectionIntentOutputWithContext(ctx).OutputState,
	}
}

type ProtectionIntentOutput struct{ *pulumi.OutputState }

func (ProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionIntent)(nil)).Elem()
}

func (o ProtectionIntentOutput) ToProtectionIntentOutput() ProtectionIntentOutput {
	return o
}

func (o ProtectionIntentOutput) ToProtectionIntentOutputWithContext(ctx context.Context) ProtectionIntentOutput {
	return o
}

func (o ProtectionIntentOutput) ToOutput(ctx context.Context) pulumix.Output[*ProtectionIntent] {
	return pulumix.Output[*ProtectionIntent]{
		OutputState: o.OutputState,
	}
}

// Optional ETag.
func (o ProtectionIntentOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ProtectionIntentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o ProtectionIntentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ProtectionIntentResource properties
func (o ProtectionIntentOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource tags.
func (o ProtectionIntentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o ProtectionIntentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionIntent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectionIntentOutput{})
}
