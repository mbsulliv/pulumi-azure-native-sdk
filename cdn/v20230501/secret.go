// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Friendly Secret name mapping to the any Secret or secret related information.
type Secret struct {
	pulumi.CustomResourceState

	DeploymentStatus pulumi.StringOutput `pulumi:"deploymentStatus"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// object which contains secret parameters
	Parameters pulumi.AnyOutput `pulumi:"parameters"`
	// The name of the profile which holds the secret.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// Provisioning status
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Secret"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Secret"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Secret"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Secret"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:Secret"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230701preview:Secret"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Secret
	err := ctx.RegisterResource("azure-native:cdn/v20230501:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("azure-native:cdn/v20230501:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
}

type SecretState struct {
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	// object which contains secret parameters
	Parameters interface{} `pulumi:"parameters"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Secret under the profile.
	SecretName *string `pulumi:"secretName"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// object which contains secret parameters
	Parameters pulumi.Input
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Name of the Secret under the profile.
	SecretName pulumi.StringPtrInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}

type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(ctx context.Context) SecretOutput
}

func (*Secret) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (i *Secret) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i *Secret) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}

func (i *Secret) ToOutput(ctx context.Context) pulumix.Output[*Secret] {
	return pulumix.Output[*Secret]{
		OutputState: i.ToSecretOutputWithContext(ctx).OutputState,
	}
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

func (o SecretOutput) ToOutput(ctx context.Context) pulumix.Output[*Secret] {
	return pulumix.Output[*Secret]{
		OutputState: o.OutputState,
	}
}

func (o SecretOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Resource name.
func (o SecretOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// object which contains secret parameters
func (o SecretOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *Secret) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

// The name of the profile which holds the secret.
func (o SecretOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

// Provisioning status
func (o SecretOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o SecretOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Secret) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o SecretOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretOutput{})
}
