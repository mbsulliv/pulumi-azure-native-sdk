// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package providerhub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Default rollout definition.
// Azure REST API version: 2021-09-01-preview. Prior API version in Azure Native 1.x: 2020-11-20
type DefaultRollout struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the rollout.
	Properties DefaultRolloutResponsePropertiesOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDefaultRollout registers a new resource with the given unique name, arguments, and options.
func NewDefaultRollout(ctx *pulumi.Context,
	name string, args *DefaultRolloutArgs, opts ...pulumi.ResourceOption) (*DefaultRollout, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:DefaultRollout"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:DefaultRollout"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210601preview:DefaultRollout"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:DefaultRollout"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DefaultRollout
	err := ctx.RegisterResource("azure-native:providerhub:DefaultRollout", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultRollout gets an existing DefaultRollout resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultRollout(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultRolloutState, opts ...pulumi.ResourceOption) (*DefaultRollout, error) {
	var resource DefaultRollout
	err := ctx.ReadResource("azure-native:providerhub:DefaultRollout", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultRollout resources.
type defaultRolloutState struct {
}

type DefaultRolloutState struct {
}

func (DefaultRolloutState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultRolloutState)(nil)).Elem()
}

type defaultRolloutArgs struct {
	// Properties of the rollout.
	Properties *DefaultRolloutProperties `pulumi:"properties"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The rollout name.
	RolloutName *string `pulumi:"rolloutName"`
}

// The set of arguments for constructing a DefaultRollout resource.
type DefaultRolloutArgs struct {
	// Properties of the rollout.
	Properties DefaultRolloutPropertiesPtrInput
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput
	// The rollout name.
	RolloutName pulumi.StringPtrInput
}

func (DefaultRolloutArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultRolloutArgs)(nil)).Elem()
}

type DefaultRolloutInput interface {
	pulumi.Input

	ToDefaultRolloutOutput() DefaultRolloutOutput
	ToDefaultRolloutOutputWithContext(ctx context.Context) DefaultRolloutOutput
}

func (*DefaultRollout) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultRollout)(nil)).Elem()
}

func (i *DefaultRollout) ToDefaultRolloutOutput() DefaultRolloutOutput {
	return i.ToDefaultRolloutOutputWithContext(context.Background())
}

func (i *DefaultRollout) ToDefaultRolloutOutputWithContext(ctx context.Context) DefaultRolloutOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultRolloutOutput)
}

func (i *DefaultRollout) ToOutput(ctx context.Context) pulumix.Output[*DefaultRollout] {
	return pulumix.Output[*DefaultRollout]{
		OutputState: i.ToDefaultRolloutOutputWithContext(ctx).OutputState,
	}
}

type DefaultRolloutOutput struct{ *pulumi.OutputState }

func (DefaultRolloutOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultRollout)(nil)).Elem()
}

func (o DefaultRolloutOutput) ToDefaultRolloutOutput() DefaultRolloutOutput {
	return o
}

func (o DefaultRolloutOutput) ToDefaultRolloutOutputWithContext(ctx context.Context) DefaultRolloutOutput {
	return o
}

func (o DefaultRolloutOutput) ToOutput(ctx context.Context) pulumix.Output[*DefaultRollout] {
	return pulumix.Output[*DefaultRollout]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o DefaultRolloutOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultRollout) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the rollout.
func (o DefaultRolloutOutput) Properties() DefaultRolloutResponsePropertiesOutput {
	return o.ApplyT(func(v *DefaultRollout) DefaultRolloutResponsePropertiesOutput { return v.Properties }).(DefaultRolloutResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DefaultRolloutOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DefaultRollout) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DefaultRolloutOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultRollout) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DefaultRolloutOutput{})
}
