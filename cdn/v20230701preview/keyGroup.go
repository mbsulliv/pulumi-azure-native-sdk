// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains a list of references of UrlSigningKey type secret objects.
type KeyGroup struct {
	pulumi.CustomResourceState

	DeploymentStatus pulumi.StringOutput `pulumi:"deploymentStatus"`
	// Names of UrlSigningKey type secret objects
	KeyReferences ResourceReferenceResponseArrayOutput `pulumi:"keyReferences"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning status
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKeyGroup registers a new resource with the given unique name, arguments, and options.
func NewKeyGroup(ctx *pulumi.Context,
	name string, args *KeyGroupArgs, opts ...pulumi.ResourceOption) (*KeyGroup, error) {
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
			Type: pulumi.String("azure-native:cdn:KeyGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KeyGroup
	err := ctx.RegisterResource("azure-native:cdn/v20230701preview:KeyGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyGroup gets an existing KeyGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyGroupState, opts ...pulumi.ResourceOption) (*KeyGroup, error) {
	var resource KeyGroup
	err := ctx.ReadResource("azure-native:cdn/v20230701preview:KeyGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyGroup resources.
type keyGroupState struct {
}

type KeyGroupState struct {
}

func (KeyGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyGroupState)(nil)).Elem()
}

type keyGroupArgs struct {
	// Name of the KeyGroup under the profile.
	KeyGroupName *string `pulumi:"keyGroupName"`
	// Names of UrlSigningKey type secret objects
	KeyReferences []ResourceReference `pulumi:"keyReferences"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a KeyGroup resource.
type KeyGroupArgs struct {
	// Name of the KeyGroup under the profile.
	KeyGroupName pulumi.StringPtrInput
	// Names of UrlSigningKey type secret objects
	KeyReferences ResourceReferenceArrayInput
	// Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (KeyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyGroupArgs)(nil)).Elem()
}

type KeyGroupInput interface {
	pulumi.Input

	ToKeyGroupOutput() KeyGroupOutput
	ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput
}

func (*KeyGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyGroup)(nil)).Elem()
}

func (i *KeyGroup) ToKeyGroupOutput() KeyGroupOutput {
	return i.ToKeyGroupOutputWithContext(context.Background())
}

func (i *KeyGroup) ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyGroupOutput)
}

type KeyGroupOutput struct{ *pulumi.OutputState }

func (KeyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyGroup)(nil)).Elem()
}

func (o KeyGroupOutput) ToKeyGroupOutput() KeyGroupOutput {
	return o
}

func (o KeyGroupOutput) ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput {
	return o
}

func (o KeyGroupOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyGroup) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Names of UrlSigningKey type secret objects
func (o KeyGroupOutput) KeyReferences() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *KeyGroup) ResourceReferenceResponseArrayOutput { return v.KeyReferences }).(ResourceReferenceResponseArrayOutput)
}

// Resource name.
func (o KeyGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning status
func (o KeyGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o KeyGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KeyGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o KeyGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyGroupOutput{})
}
