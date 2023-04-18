// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings with single toggle.
type EyesOn struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Determines whether the setting is enable or disabled.
	IsEnabled pulumi.BoolOutput `pulumi:"isEnabled"`
	// The kind of the setting
	// Expected value is 'EyesOn'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEyesOn registers a new resource with the given unique name, arguments, and options.
func NewEyesOn(ctx *pulumi.Context,
	name string, args *EyesOnArgs, opts ...pulumi.ResourceOption) (*EyesOn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("EyesOn")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:EyesOn"),
		},
	})
	opts = append(opts, aliases)
	var resource EyesOn
	err := ctx.RegisterResource("azure-native:securityinsights/v20220701preview:EyesOn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEyesOn gets an existing EyesOn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEyesOn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EyesOnState, opts ...pulumi.ResourceOption) (*EyesOn, error) {
	var resource EyesOn
	err := ctx.ReadResource("azure-native:securityinsights/v20220701preview:EyesOn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EyesOn resources.
type eyesOnState struct {
}

type EyesOnState struct {
}

func (EyesOnState) ElementType() reflect.Type {
	return reflect.TypeOf((*eyesOnState)(nil)).Elem()
}

type eyesOnArgs struct {
	// The kind of the setting
	// Expected value is 'EyesOn'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName *string `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a EyesOn resource.
type EyesOnArgs struct {
	// The kind of the setting
	// Expected value is 'EyesOn'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (EyesOnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eyesOnArgs)(nil)).Elem()
}

type EyesOnInput interface {
	pulumi.Input

	ToEyesOnOutput() EyesOnOutput
	ToEyesOnOutputWithContext(ctx context.Context) EyesOnOutput
}

func (*EyesOn) ElementType() reflect.Type {
	return reflect.TypeOf((**EyesOn)(nil)).Elem()
}

func (i *EyesOn) ToEyesOnOutput() EyesOnOutput {
	return i.ToEyesOnOutputWithContext(context.Background())
}

func (i *EyesOn) ToEyesOnOutputWithContext(ctx context.Context) EyesOnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EyesOnOutput)
}

type EyesOnOutput struct{ *pulumi.OutputState }

func (EyesOnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EyesOn)(nil)).Elem()
}

func (o EyesOnOutput) ToEyesOnOutput() EyesOnOutput {
	return o
}

func (o EyesOnOutput) ToEyesOnOutputWithContext(ctx context.Context) EyesOnOutput {
	return o
}

// Etag of the azure resource
func (o EyesOnOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Determines whether the setting is enable or disabled.
func (o EyesOnOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

// The kind of the setting
// Expected value is 'EyesOn'.
func (o EyesOnOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o EyesOnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EyesOnOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EyesOn) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EyesOnOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EyesOnOutput{})
}
