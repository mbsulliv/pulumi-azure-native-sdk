// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An object that represents a credential set resource for a container registry.
type CredentialSet struct {
	pulumi.CustomResourceState

	// List of authentication credentials stored for an upstream.
	// Usually consists of a primary and an optional secondary credential.
	AuthCredentials AuthCredentialResponseArrayOutput `pulumi:"authCredentials"`
	// The creation date of credential store resource.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Identities associated with the resource. This is used to access the KeyVault secrets.
	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	// The credentials are stored for this upstream or login server.
	LoginServer pulumi.StringPtrOutput `pulumi:"loginServer"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCredentialSet registers a new resource with the given unique name, arguments, and options.
func NewCredentialSet(ctx *pulumi.Context,
	name string, args *CredentialSetArgs, opts ...pulumi.ResourceOption) (*CredentialSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:CredentialSet"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230101preview:CredentialSet"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:CredentialSet"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230801preview:CredentialSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CredentialSet
	err := ctx.RegisterResource("azure-native:containerregistry/v20230701:CredentialSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCredentialSet gets an existing CredentialSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCredentialSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialSetState, opts ...pulumi.ResourceOption) (*CredentialSet, error) {
	var resource CredentialSet
	err := ctx.ReadResource("azure-native:containerregistry/v20230701:CredentialSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CredentialSet resources.
type credentialSetState struct {
}

type CredentialSetState struct {
}

func (CredentialSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialSetState)(nil)).Elem()
}

type credentialSetArgs struct {
	// List of authentication credentials stored for an upstream.
	// Usually consists of a primary and an optional secondary credential.
	AuthCredentials []AuthCredential `pulumi:"authCredentials"`
	// The name of the credential set.
	CredentialSetName *string `pulumi:"credentialSetName"`
	// Identities associated with the resource. This is used to access the KeyVault secrets.
	Identity *IdentityProperties `pulumi:"identity"`
	// The credentials are stored for this upstream or login server.
	LoginServer *string `pulumi:"loginServer"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a CredentialSet resource.
type CredentialSetArgs struct {
	// List of authentication credentials stored for an upstream.
	// Usually consists of a primary and an optional secondary credential.
	AuthCredentials AuthCredentialArrayInput
	// The name of the credential set.
	CredentialSetName pulumi.StringPtrInput
	// Identities associated with the resource. This is used to access the KeyVault secrets.
	Identity IdentityPropertiesPtrInput
	// The credentials are stored for this upstream or login server.
	LoginServer pulumi.StringPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (CredentialSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialSetArgs)(nil)).Elem()
}

type CredentialSetInput interface {
	pulumi.Input

	ToCredentialSetOutput() CredentialSetOutput
	ToCredentialSetOutputWithContext(ctx context.Context) CredentialSetOutput
}

func (*CredentialSet) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialSet)(nil)).Elem()
}

func (i *CredentialSet) ToCredentialSetOutput() CredentialSetOutput {
	return i.ToCredentialSetOutputWithContext(context.Background())
}

func (i *CredentialSet) ToCredentialSetOutputWithContext(ctx context.Context) CredentialSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialSetOutput)
}

func (i *CredentialSet) ToOutput(ctx context.Context) pulumix.Output[*CredentialSet] {
	return pulumix.Output[*CredentialSet]{
		OutputState: i.ToCredentialSetOutputWithContext(ctx).OutputState,
	}
}

type CredentialSetOutput struct{ *pulumi.OutputState }

func (CredentialSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialSet)(nil)).Elem()
}

func (o CredentialSetOutput) ToCredentialSetOutput() CredentialSetOutput {
	return o
}

func (o CredentialSetOutput) ToCredentialSetOutputWithContext(ctx context.Context) CredentialSetOutput {
	return o
}

func (o CredentialSetOutput) ToOutput(ctx context.Context) pulumix.Output[*CredentialSet] {
	return pulumix.Output[*CredentialSet]{
		OutputState: o.OutputState,
	}
}

// List of authentication credentials stored for an upstream.
// Usually consists of a primary and an optional secondary credential.
func (o CredentialSetOutput) AuthCredentials() AuthCredentialResponseArrayOutput {
	return o.ApplyT(func(v *CredentialSet) AuthCredentialResponseArrayOutput { return v.AuthCredentials }).(AuthCredentialResponseArrayOutput)
}

// The creation date of credential store resource.
func (o CredentialSetOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Identities associated with the resource. This is used to access the KeyVault secrets.
func (o CredentialSetOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CredentialSet) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

// The credentials are stored for this upstream or login server.
func (o CredentialSetOutput) LoginServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringPtrOutput { return v.LoginServer }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o CredentialSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o CredentialSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o CredentialSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CredentialSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o CredentialSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CredentialSetOutput{})
}
