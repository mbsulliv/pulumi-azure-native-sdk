// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings with single toggle.
type Ueba struct {
	pulumi.CustomResourceState

	// The relevant data sources that enriched by ueba
	DataSources pulumi.StringArrayOutput `pulumi:"dataSources"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the setting
	// Expected value is 'Ueba'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUeba registers a new resource with the given unique name, arguments, and options.
func NewUeba(ctx *pulumi.Context,
	name string, args *UebaArgs, opts ...pulumi.ResourceOption) (*Ueba, error) {
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
	args.Kind = pulumi.String("Ueba")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:Ueba"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Ueba
	err := ctx.RegisterResource("azure-native:securityinsights/v20230601preview:Ueba", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUeba gets an existing Ueba resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUeba(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UebaState, opts ...pulumi.ResourceOption) (*Ueba, error) {
	var resource Ueba
	err := ctx.ReadResource("azure-native:securityinsights/v20230601preview:Ueba", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ueba resources.
type uebaState struct {
}

type UebaState struct {
}

func (UebaState) ElementType() reflect.Type {
	return reflect.TypeOf((*uebaState)(nil)).Elem()
}

type uebaArgs struct {
	// The relevant data sources that enriched by ueba
	DataSources []string `pulumi:"dataSources"`
	// The kind of the setting
	// Expected value is 'Ueba'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName *string `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Ueba resource.
type UebaArgs struct {
	// The relevant data sources that enriched by ueba
	DataSources pulumi.StringArrayInput
	// The kind of the setting
	// Expected value is 'Ueba'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (UebaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*uebaArgs)(nil)).Elem()
}

type UebaInput interface {
	pulumi.Input

	ToUebaOutput() UebaOutput
	ToUebaOutputWithContext(ctx context.Context) UebaOutput
}

func (*Ueba) ElementType() reflect.Type {
	return reflect.TypeOf((**Ueba)(nil)).Elem()
}

func (i *Ueba) ToUebaOutput() UebaOutput {
	return i.ToUebaOutputWithContext(context.Background())
}

func (i *Ueba) ToUebaOutputWithContext(ctx context.Context) UebaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UebaOutput)
}

type UebaOutput struct{ *pulumi.OutputState }

func (UebaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ueba)(nil)).Elem()
}

func (o UebaOutput) ToUebaOutput() UebaOutput {
	return o
}

func (o UebaOutput) ToUebaOutputWithContext(ctx context.Context) UebaOutput {
	return o
}

// The relevant data sources that enriched by ueba
func (o UebaOutput) DataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringArrayOutput { return v.DataSources }).(pulumi.StringArrayOutput)
}

// Etag of the azure resource
func (o UebaOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the setting
// Expected value is 'Ueba'.
func (o UebaOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o UebaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o UebaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Ueba) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o UebaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UebaOutput{})
}
