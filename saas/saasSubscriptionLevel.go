// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// SaaS REST API resource definition.
// Azure REST API version: 2018-03-01-beta. Prior API version in Azure Native 1.x: 2018-03-01-beta
type SaasSubscriptionLevel struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// saas properties
	Properties SaasResourceResponsePropertiesOutput `pulumi:"properties"`
	// the resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSaasSubscriptionLevel registers a new resource with the given unique name, arguments, and options.
func NewSaasSubscriptionLevel(ctx *pulumi.Context,
	name string, args *SaasSubscriptionLevelArgs, opts ...pulumi.ResourceOption) (*SaasSubscriptionLevel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:saas/v20180301beta:SaasSubscriptionLevel"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SaasSubscriptionLevel
	err := ctx.RegisterResource("azure-native:saas:SaasSubscriptionLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSaasSubscriptionLevel gets an existing SaasSubscriptionLevel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSaasSubscriptionLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SaasSubscriptionLevelState, opts ...pulumi.ResourceOption) (*SaasSubscriptionLevel, error) {
	var resource SaasSubscriptionLevel
	err := ctx.ReadResource("azure-native:saas:SaasSubscriptionLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SaasSubscriptionLevel resources.
type saasSubscriptionLevelState struct {
}

type SaasSubscriptionLevelState struct {
}

func (SaasSubscriptionLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*saasSubscriptionLevelState)(nil)).Elem()
}

type saasSubscriptionLevelArgs struct {
	// Resource location. Only value allowed for SaaS is 'global'
	Location *string `pulumi:"location"`
	// The resource name
	Name *string `pulumi:"name"`
	// Properties of the SaaS resource that are relevant for creation.
	Properties *SaasCreationProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName *string `pulumi:"resourceName"`
	// the resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SaasSubscriptionLevel resource.
type SaasSubscriptionLevelArgs struct {
	// Resource location. Only value allowed for SaaS is 'global'
	Location pulumi.StringPtrInput
	// The resource name
	Name pulumi.StringPtrInput
	// Properties of the SaaS resource that are relevant for creation.
	Properties SaasCreationPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringPtrInput
	// the resource tags.
	Tags pulumi.StringMapInput
}

func (SaasSubscriptionLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*saasSubscriptionLevelArgs)(nil)).Elem()
}

type SaasSubscriptionLevelInput interface {
	pulumi.Input

	ToSaasSubscriptionLevelOutput() SaasSubscriptionLevelOutput
	ToSaasSubscriptionLevelOutputWithContext(ctx context.Context) SaasSubscriptionLevelOutput
}

func (*SaasSubscriptionLevel) ElementType() reflect.Type {
	return reflect.TypeOf((**SaasSubscriptionLevel)(nil)).Elem()
}

func (i *SaasSubscriptionLevel) ToSaasSubscriptionLevelOutput() SaasSubscriptionLevelOutput {
	return i.ToSaasSubscriptionLevelOutputWithContext(context.Background())
}

func (i *SaasSubscriptionLevel) ToSaasSubscriptionLevelOutputWithContext(ctx context.Context) SaasSubscriptionLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SaasSubscriptionLevelOutput)
}

func (i *SaasSubscriptionLevel) ToOutput(ctx context.Context) pulumix.Output[*SaasSubscriptionLevel] {
	return pulumix.Output[*SaasSubscriptionLevel]{
		OutputState: i.ToSaasSubscriptionLevelOutputWithContext(ctx).OutputState,
	}
}

type SaasSubscriptionLevelOutput struct{ *pulumi.OutputState }

func (SaasSubscriptionLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SaasSubscriptionLevel)(nil)).Elem()
}

func (o SaasSubscriptionLevelOutput) ToSaasSubscriptionLevelOutput() SaasSubscriptionLevelOutput {
	return o
}

func (o SaasSubscriptionLevelOutput) ToSaasSubscriptionLevelOutputWithContext(ctx context.Context) SaasSubscriptionLevelOutput {
	return o
}

func (o SaasSubscriptionLevelOutput) ToOutput(ctx context.Context) pulumix.Output[*SaasSubscriptionLevel] {
	return pulumix.Output[*SaasSubscriptionLevel]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o SaasSubscriptionLevelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SaasSubscriptionLevel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// saas properties
func (o SaasSubscriptionLevelOutput) Properties() SaasResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *SaasSubscriptionLevel) SaasResourceResponsePropertiesOutput { return v.Properties }).(SaasResourceResponsePropertiesOutput)
}

// the resource tags.
func (o SaasSubscriptionLevelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SaasSubscriptionLevel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o SaasSubscriptionLevelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SaasSubscriptionLevel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SaasSubscriptionLevelOutput{})
}
