// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A logical database transparent data encryption state.
type TransparentDataEncryption struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the state of the transparent data encryption.
	State pulumi.StringOutput `pulumi:"state"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTransparentDataEncryption registers a new resource with the given unique name, arguments, and options.
func NewTransparentDataEncryption(ctx *pulumi.Context,
	name string, args *TransparentDataEncryptionArgs, opts ...pulumi.ResourceOption) (*TransparentDataEncryption, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:TransparentDataEncryption"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TransparentDataEncryption
	err := ctx.RegisterResource("azure-native:sql/v20211101:TransparentDataEncryption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransparentDataEncryption gets an existing TransparentDataEncryption resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransparentDataEncryption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransparentDataEncryptionState, opts ...pulumi.ResourceOption) (*TransparentDataEncryption, error) {
	var resource TransparentDataEncryption
	err := ctx.ReadResource("azure-native:sql/v20211101:TransparentDataEncryption", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransparentDataEncryption resources.
type transparentDataEncryptionState struct {
}

type TransparentDataEncryptionState struct {
}

func (TransparentDataEncryptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*transparentDataEncryptionState)(nil)).Elem()
}

type transparentDataEncryptionArgs struct {
	// The name of the logical database for which the security alert policy is defined.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// Specifies the state of the transparent data encryption.
	State TransparentDataEncryptionStateEnum `pulumi:"state"`
	// The name of the transparent data encryption configuration.
	TdeName *string `pulumi:"tdeName"`
}

// The set of arguments for constructing a TransparentDataEncryption resource.
type TransparentDataEncryptionArgs struct {
	// The name of the logical database for which the security alert policy is defined.
	DatabaseName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// Specifies the state of the transparent data encryption.
	State TransparentDataEncryptionStateEnumInput
	// The name of the transparent data encryption configuration.
	TdeName pulumi.StringPtrInput
}

func (TransparentDataEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transparentDataEncryptionArgs)(nil)).Elem()
}

type TransparentDataEncryptionInput interface {
	pulumi.Input

	ToTransparentDataEncryptionOutput() TransparentDataEncryptionOutput
	ToTransparentDataEncryptionOutputWithContext(ctx context.Context) TransparentDataEncryptionOutput
}

func (*TransparentDataEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((**TransparentDataEncryption)(nil)).Elem()
}

func (i *TransparentDataEncryption) ToTransparentDataEncryptionOutput() TransparentDataEncryptionOutput {
	return i.ToTransparentDataEncryptionOutputWithContext(context.Background())
}

func (i *TransparentDataEncryption) ToTransparentDataEncryptionOutputWithContext(ctx context.Context) TransparentDataEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransparentDataEncryptionOutput)
}

func (i *TransparentDataEncryption) ToOutput(ctx context.Context) pulumix.Output[*TransparentDataEncryption] {
	return pulumix.Output[*TransparentDataEncryption]{
		OutputState: i.ToTransparentDataEncryptionOutputWithContext(ctx).OutputState,
	}
}

type TransparentDataEncryptionOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransparentDataEncryption)(nil)).Elem()
}

func (o TransparentDataEncryptionOutput) ToTransparentDataEncryptionOutput() TransparentDataEncryptionOutput {
	return o
}

func (o TransparentDataEncryptionOutput) ToTransparentDataEncryptionOutputWithContext(ctx context.Context) TransparentDataEncryptionOutput {
	return o
}

func (o TransparentDataEncryptionOutput) ToOutput(ctx context.Context) pulumix.Output[*TransparentDataEncryption] {
	return pulumix.Output[*TransparentDataEncryption]{
		OutputState: o.OutputState,
	}
}

// Resource name.
func (o TransparentDataEncryptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TransparentDataEncryption) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the state of the transparent data encryption.
func (o TransparentDataEncryptionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *TransparentDataEncryption) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Resource type.
func (o TransparentDataEncryptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TransparentDataEncryption) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TransparentDataEncryptionOutput{})
}
