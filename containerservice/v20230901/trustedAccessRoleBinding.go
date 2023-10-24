// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Defines binding between a resource and role
type TrustedAccessRoleBinding struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The current provisioning state of trusted access role binding.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceId pulumi.StringOutput `pulumi:"sourceResourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTrustedAccessRoleBinding registers a new resource with the given unique name, arguments, and options.
func NewTrustedAccessRoleBinding(ctx *pulumi.Context,
	name string, args *TrustedAccessRoleBindingArgs, opts ...pulumi.ResourceOption) (*TrustedAccessRoleBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.SourceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221002preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221102preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230102preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230202preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230302preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230402preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230502preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230602preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230702preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230802preview:TrustedAccessRoleBinding"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TrustedAccessRoleBinding
	err := ctx.RegisterResource("azure-native:containerservice/v20230901:TrustedAccessRoleBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrustedAccessRoleBinding gets an existing TrustedAccessRoleBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrustedAccessRoleBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrustedAccessRoleBindingState, opts ...pulumi.ResourceOption) (*TrustedAccessRoleBinding, error) {
	var resource TrustedAccessRoleBinding
	err := ctx.ReadResource("azure-native:containerservice/v20230901:TrustedAccessRoleBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrustedAccessRoleBinding resources.
type trustedAccessRoleBindingState struct {
}

type TrustedAccessRoleBindingState struct {
}

func (TrustedAccessRoleBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*trustedAccessRoleBindingState)(nil)).Elem()
}

type trustedAccessRoleBindingArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
	Roles []string `pulumi:"roles"`
	// The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceId string `pulumi:"sourceResourceId"`
	// The name of trusted access role binding.
	TrustedAccessRoleBindingName *string `pulumi:"trustedAccessRoleBindingName"`
}

// The set of arguments for constructing a TrustedAccessRoleBinding resource.
type TrustedAccessRoleBindingArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput
	// A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
	Roles pulumi.StringArrayInput
	// The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceId pulumi.StringInput
	// The name of trusted access role binding.
	TrustedAccessRoleBindingName pulumi.StringPtrInput
}

func (TrustedAccessRoleBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trustedAccessRoleBindingArgs)(nil)).Elem()
}

type TrustedAccessRoleBindingInput interface {
	pulumi.Input

	ToTrustedAccessRoleBindingOutput() TrustedAccessRoleBindingOutput
	ToTrustedAccessRoleBindingOutputWithContext(ctx context.Context) TrustedAccessRoleBindingOutput
}

func (*TrustedAccessRoleBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustedAccessRoleBinding)(nil)).Elem()
}

func (i *TrustedAccessRoleBinding) ToTrustedAccessRoleBindingOutput() TrustedAccessRoleBindingOutput {
	return i.ToTrustedAccessRoleBindingOutputWithContext(context.Background())
}

func (i *TrustedAccessRoleBinding) ToTrustedAccessRoleBindingOutputWithContext(ctx context.Context) TrustedAccessRoleBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedAccessRoleBindingOutput)
}

func (i *TrustedAccessRoleBinding) ToOutput(ctx context.Context) pulumix.Output[*TrustedAccessRoleBinding] {
	return pulumix.Output[*TrustedAccessRoleBinding]{
		OutputState: i.ToTrustedAccessRoleBindingOutputWithContext(ctx).OutputState,
	}
}

type TrustedAccessRoleBindingOutput struct{ *pulumi.OutputState }

func (TrustedAccessRoleBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustedAccessRoleBinding)(nil)).Elem()
}

func (o TrustedAccessRoleBindingOutput) ToTrustedAccessRoleBindingOutput() TrustedAccessRoleBindingOutput {
	return o
}

func (o TrustedAccessRoleBindingOutput) ToTrustedAccessRoleBindingOutputWithContext(ctx context.Context) TrustedAccessRoleBindingOutput {
	return o
}

func (o TrustedAccessRoleBindingOutput) ToOutput(ctx context.Context) pulumix.Output[*TrustedAccessRoleBinding] {
	return pulumix.Output[*TrustedAccessRoleBinding]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o TrustedAccessRoleBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of trusted access role binding.
func (o TrustedAccessRoleBindingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
func (o TrustedAccessRoleBindingOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The ARM resource ID of source resource that trusted access is configured for.
func (o TrustedAccessRoleBindingOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringOutput { return v.SourceResourceId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o TrustedAccessRoleBindingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TrustedAccessRoleBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrustedAccessRoleBindingOutput{})
}
