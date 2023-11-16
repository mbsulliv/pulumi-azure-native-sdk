// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabricmesh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This type describes a value of a secret resource. The name of this resource is the version identifier corresponding to this secret value.
// Azure REST API version: 2018-09-01-preview. Prior API version in Azure Native 1.x: 2018-09-01-preview.
type SecretValue struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// The actual value of the secret.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewSecretValue registers a new resource with the given unique name, arguments, and options.
func NewSecretValue(ctx *pulumi.Context,
	name string, args *SecretValueArgs, opts ...pulumi.ResourceOption) (*SecretValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecretResourceName == nil {
		return nil, errors.New("invalid value for required argument 'SecretResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180901preview:SecretValue"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecretValue
	err := ctx.RegisterResource("azure-native:servicefabricmesh:SecretValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretValue gets an existing SecretValue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretValueState, opts ...pulumi.ResourceOption) (*SecretValue, error) {
	var resource SecretValue
	err := ctx.ReadResource("azure-native:servicefabricmesh:SecretValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretValue resources.
type secretValueState struct {
}

type SecretValueState struct {
}

func (SecretValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretValueState)(nil)).Elem()
}

type secretValueArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the secret resource.
	SecretResourceName string `pulumi:"secretResourceName"`
	// The name of the secret resource value which is typically the version identifier for the value.
	SecretValueResourceName *string `pulumi:"secretValueResourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The actual value of the secret.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a SecretValue resource.
type SecretValueArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Azure resource group name
	ResourceGroupName pulumi.StringInput
	// The name of the secret resource.
	SecretResourceName pulumi.StringInput
	// The name of the secret resource value which is typically the version identifier for the value.
	SecretValueResourceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The actual value of the secret.
	Value pulumi.StringPtrInput
}

func (SecretValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretValueArgs)(nil)).Elem()
}

type SecretValueInput interface {
	pulumi.Input

	ToSecretValueOutput() SecretValueOutput
	ToSecretValueOutputWithContext(ctx context.Context) SecretValueOutput
}

func (*SecretValue) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretValue)(nil)).Elem()
}

func (i *SecretValue) ToSecretValueOutput() SecretValueOutput {
	return i.ToSecretValueOutputWithContext(context.Background())
}

func (i *SecretValue) ToSecretValueOutputWithContext(ctx context.Context) SecretValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretValueOutput)
}

type SecretValueOutput struct{ *pulumi.OutputState }

func (SecretValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretValue)(nil)).Elem()
}

func (o SecretValueOutput) ToSecretValueOutput() SecretValueOutput {
	return o
}

func (o SecretValueOutput) ToSecretValueOutputWithContext(ctx context.Context) SecretValueOutput {
	return o
}

// The geo-location where the resource lives
func (o SecretValueOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SecretValueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of the resource.
func (o SecretValueOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o SecretValueOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o SecretValueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The actual value of the secret.
func (o SecretValueOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretValue) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretValueOutput{})
}
