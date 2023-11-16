// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Outbound Rule Basic Resource for the managed network of a machine learning workspace.
type ManagedNetworkSettingsRule struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Outbound Rule for the managed network of a machine learning workspace.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedNetworkSettingsRule registers a new resource with the given unique name, arguments, and options.
func NewManagedNetworkSettingsRule(ctx *pulumi.Context,
	name string, args *ManagedNetworkSettingsRuleArgs, opts ...pulumi.ResourceOption) (*ManagedNetworkSettingsRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:ManagedNetworkSettingsRule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:ManagedNetworkSettingsRule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:ManagedNetworkSettingsRule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:ManagedNetworkSettingsRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedNetworkSettingsRule
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401preview:ManagedNetworkSettingsRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedNetworkSettingsRule gets an existing ManagedNetworkSettingsRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedNetworkSettingsRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkSettingsRuleState, opts ...pulumi.ResourceOption) (*ManagedNetworkSettingsRule, error) {
	var resource ManagedNetworkSettingsRule
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401preview:ManagedNetworkSettingsRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedNetworkSettingsRule resources.
type managedNetworkSettingsRuleState struct {
}

type ManagedNetworkSettingsRuleState struct {
}

func (ManagedNetworkSettingsRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkSettingsRuleState)(nil)).Elem()
}

type managedNetworkSettingsRuleArgs struct {
	// Outbound Rule for the managed network of a machine learning workspace.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the workspace managed network outbound rule
	RuleName *string `pulumi:"ruleName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ManagedNetworkSettingsRule resource.
type ManagedNetworkSettingsRuleArgs struct {
	// Outbound Rule for the managed network of a machine learning workspace.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of the workspace managed network outbound rule
	RuleName pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (ManagedNetworkSettingsRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkSettingsRuleArgs)(nil)).Elem()
}

type ManagedNetworkSettingsRuleInput interface {
	pulumi.Input

	ToManagedNetworkSettingsRuleOutput() ManagedNetworkSettingsRuleOutput
	ToManagedNetworkSettingsRuleOutputWithContext(ctx context.Context) ManagedNetworkSettingsRuleOutput
}

func (*ManagedNetworkSettingsRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkSettingsRule)(nil)).Elem()
}

func (i *ManagedNetworkSettingsRule) ToManagedNetworkSettingsRuleOutput() ManagedNetworkSettingsRuleOutput {
	return i.ToManagedNetworkSettingsRuleOutputWithContext(context.Background())
}

func (i *ManagedNetworkSettingsRule) ToManagedNetworkSettingsRuleOutputWithContext(ctx context.Context) ManagedNetworkSettingsRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkSettingsRuleOutput)
}

type ManagedNetworkSettingsRuleOutput struct{ *pulumi.OutputState }

func (ManagedNetworkSettingsRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkSettingsRule)(nil)).Elem()
}

func (o ManagedNetworkSettingsRuleOutput) ToManagedNetworkSettingsRuleOutput() ManagedNetworkSettingsRuleOutput {
	return o
}

func (o ManagedNetworkSettingsRuleOutput) ToManagedNetworkSettingsRuleOutputWithContext(ctx context.Context) ManagedNetworkSettingsRuleOutput {
	return o
}

// The name of the resource
func (o ManagedNetworkSettingsRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkSettingsRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Outbound Rule for the managed network of a machine learning workspace.
func (o ManagedNetworkSettingsRuleOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ManagedNetworkSettingsRule) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ManagedNetworkSettingsRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedNetworkSettingsRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ManagedNetworkSettingsRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkSettingsRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkSettingsRuleOutput{})
}
