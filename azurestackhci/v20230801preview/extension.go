// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Details of a particular extension in HCI Cluster.
type Extension struct {
	pulumi.CustomResourceState

	// Aggregate state of Arc Extensions across the nodes in this HCI cluster.
	AggregateState pulumi.StringOutput `pulumi:"aggregateState"`
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version available.
	EnableAutomaticUpgrade pulumi.BoolPtrOutput `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// Indicates if the extension is managed by azure or the user.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of Arc Extension in each of the nodes.
	PerNodeExtensionDetails PerNodeExtensionStateResponseArrayOutput `pulumi:"perNodeExtensionDetails"`
	// Protected settings (may contain secrets).
	ProtectedSettings pulumi.AnyOutput `pulumi:"protectedSettings"`
	// Provisioning state of the Extension proxy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrOutput `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings pulumi.AnyOutput `pulumi:"settings"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the script handler. Latest version would be used if not specified.
	TypeHandlerVersion pulumi.StringPtrOutput `pulumi:"typeHandlerVersion"`
}

// NewExtension registers a new resource with the given unique name, arguments, and options.
func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArcSettingName == nil {
		return nil, errors.New("invalid value for required argument 'ArcSettingName'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210101preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220101:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220301:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220501:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220901:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221001:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221201:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230201:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230301:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230601:Extension"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801:Extension"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Extension
	err := ctx.RegisterResource("azure-native:azurestackhci/v20230801preview:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtension gets an existing Extension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:azurestackhci/v20230801preview:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extension resources.
type extensionState struct {
}

type ExtensionState struct {
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	// The name of the proxy resource holding details of HCI ArcSetting information.
	ArcSettingName string `pulumi:"arcSettingName"`
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version available.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// The name of the machine extension.
	ExtensionName *string `pulumi:"extensionName"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// Protected settings (may contain secrets).
	ProtectedSettings interface{} `pulumi:"protectedSettings"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Json formatted public settings for the extension.
	Settings interface{} `pulumi:"settings"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler. Latest version would be used if not specified.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
}

// The set of arguments for constructing a Extension resource.
type ExtensionArgs struct {
	// The name of the proxy resource holding details of HCI ArcSetting information.
	ArcSettingName pulumi.StringInput
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version available.
	EnableAutomaticUpgrade pulumi.BoolPtrInput
	// The name of the machine extension.
	ExtensionName pulumi.StringPtrInput
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// Protected settings (may contain secrets).
	ProtectedSettings pulumi.Input
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Json formatted public settings for the extension.
	Settings pulumi.Input
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler. Latest version would be used if not specified.
	TypeHandlerVersion pulumi.StringPtrInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput
}

func (*Extension) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

// Aggregate state of Arc Extensions across the nodes in this HCI cluster.
func (o ExtensionOutput) AggregateState() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.AggregateState }).(pulumi.StringOutput)
}

// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
func (o ExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version available.
func (o ExtensionOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

// How the extension handler should be forced to update even if the extension configuration has not changed.
func (o ExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// Indicates if the extension is managed by azure or the user.
func (o ExtensionOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// The name of the resource
func (o ExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of Arc Extension in each of the nodes.
func (o ExtensionOutput) PerNodeExtensionDetails() PerNodeExtensionStateResponseArrayOutput {
	return o.ApplyT(func(v *Extension) PerNodeExtensionStateResponseArrayOutput { return v.PerNodeExtensionDetails }).(PerNodeExtensionStateResponseArrayOutput)
}

// Protected settings (may contain secrets).
func (o ExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *Extension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

// Provisioning state of the Extension proxy resource.
func (o ExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the extension handler publisher.
func (o ExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

// Json formatted public settings for the extension.
func (o ExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *Extension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Extension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the version of the script handler. Latest version would be used if not specified.
func (o ExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}