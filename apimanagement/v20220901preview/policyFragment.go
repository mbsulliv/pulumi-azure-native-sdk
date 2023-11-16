// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Policy fragment contract details.
type PolicyFragment struct {
	pulumi.CustomResourceState

	// Policy fragment description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Format of the policy fragment content.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Contents of the policy fragment.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewPolicyFragment registers a new resource with the given unique name, arguments, and options.
func NewPolicyFragment(ctx *pulumi.Context,
	name string, args *PolicyFragmentArgs, opts ...pulumi.ResourceOption) (*PolicyFragment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.Format == nil {
		args.Format = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:PolicyFragment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:PolicyFragment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:PolicyFragment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:PolicyFragment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:PolicyFragment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PolicyFragment
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:PolicyFragment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyFragment gets an existing PolicyFragment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyFragment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyFragmentState, opts ...pulumi.ResourceOption) (*PolicyFragment, error) {
	var resource PolicyFragment
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:PolicyFragment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyFragment resources.
type policyFragmentState struct {
}

type PolicyFragmentState struct {
}

func (PolicyFragmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyFragmentState)(nil)).Elem()
}

type policyFragmentArgs struct {
	// Policy fragment description.
	Description *string `pulumi:"description"`
	// Format of the policy fragment content.
	Format *string `pulumi:"format"`
	// A resource identifier.
	Id *string `pulumi:"id"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Contents of the policy fragment.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a PolicyFragment resource.
type PolicyFragmentArgs struct {
	// Policy fragment description.
	Description pulumi.StringPtrInput
	// Format of the policy fragment content.
	Format pulumi.StringPtrInput
	// A resource identifier.
	Id pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Contents of the policy fragment.
	Value pulumi.StringInput
}

func (PolicyFragmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyFragmentArgs)(nil)).Elem()
}

type PolicyFragmentInput interface {
	pulumi.Input

	ToPolicyFragmentOutput() PolicyFragmentOutput
	ToPolicyFragmentOutputWithContext(ctx context.Context) PolicyFragmentOutput
}

func (*PolicyFragment) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyFragment)(nil)).Elem()
}

func (i *PolicyFragment) ToPolicyFragmentOutput() PolicyFragmentOutput {
	return i.ToPolicyFragmentOutputWithContext(context.Background())
}

func (i *PolicyFragment) ToPolicyFragmentOutputWithContext(ctx context.Context) PolicyFragmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyFragmentOutput)
}

type PolicyFragmentOutput struct{ *pulumi.OutputState }

func (PolicyFragmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyFragment)(nil)).Elem()
}

func (o PolicyFragmentOutput) ToPolicyFragmentOutput() PolicyFragmentOutput {
	return o
}

func (o PolicyFragmentOutput) ToPolicyFragmentOutputWithContext(ctx context.Context) PolicyFragmentOutput {
	return o
}

// Policy fragment description.
func (o PolicyFragmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Format of the policy fragment content.
func (o PolicyFragmentOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o PolicyFragmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PolicyFragmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Contents of the policy fragment.
func (o PolicyFragmentOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyFragmentOutput{})
}
