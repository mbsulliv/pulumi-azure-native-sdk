// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The application resource.
type ManagedClusterApplication struct {
	pulumi.CustomResourceState

	// Describes the managed identities for an Azure resource.
	Identity ManagedIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// List of user assigned identities for the application, each mapped to a friendly name.
	ManagedIdentities ApplicationUserAssignedIdentityResponseArrayOutput `pulumi:"managedIdentities"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of application parameters with overridden values from their default values specified in the application manifest.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Describes the policy for a monitored application upgrade.
	UpgradePolicy ApplicationUpgradePolicyResponsePtrOutput `pulumi:"upgradePolicy"`
	// The version of the application type as defined in the application manifest.
	// This name must be the full Arm Resource ID for the referenced application type version.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewManagedClusterApplication registers a new resource with the given unique name, arguments, and options.
func NewManagedClusterApplication(ctx *pulumi.Context,
	name string, args *ManagedClusterApplicationArgs, opts ...pulumi.ResourceOption) (*ManagedClusterApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210701preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220101:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220801preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20221001preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230201preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230301preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230701preview:ManagedClusterApplication"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20231101preview:ManagedClusterApplication"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedClusterApplication
	err := ctx.RegisterResource("azure-native:servicefabric/v20230901preview:ManagedClusterApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedClusterApplication gets an existing ManagedClusterApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedClusterApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterApplicationState, opts ...pulumi.ResourceOption) (*ManagedClusterApplication, error) {
	var resource ManagedClusterApplication
	err := ctx.ReadResource("azure-native:servicefabric/v20230901preview:ManagedClusterApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedClusterApplication resources.
type managedClusterApplicationState struct {
}

type ManagedClusterApplicationState struct {
}

func (ManagedClusterApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterApplicationState)(nil)).Elem()
}

type managedClusterApplicationArgs struct {
	// The name of the application resource.
	ApplicationName *string `pulumi:"applicationName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Describes the managed identities for an Azure resource.
	Identity *ManagedIdentity `pulumi:"identity"`
	// Resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// List of user assigned identities for the application, each mapped to a friendly name.
	ManagedIdentities []ApplicationUserAssignedIdentity `pulumi:"managedIdentities"`
	// List of application parameters with overridden values from their default values specified in the application manifest.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Describes the policy for a monitored application upgrade.
	UpgradePolicy *ApplicationUpgradePolicy `pulumi:"upgradePolicy"`
	// The version of the application type as defined in the application manifest.
	// This name must be the full Arm Resource ID for the referenced application type version.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a ManagedClusterApplication resource.
type ManagedClusterApplicationArgs struct {
	// The name of the application resource.
	ApplicationName pulumi.StringPtrInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Describes the managed identities for an Azure resource.
	Identity ManagedIdentityPtrInput
	// Resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// List of user assigned identities for the application, each mapped to a friendly name.
	ManagedIdentities ApplicationUserAssignedIdentityArrayInput
	// List of application parameters with overridden values from their default values specified in the application manifest.
	Parameters pulumi.StringMapInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// Describes the policy for a monitored application upgrade.
	UpgradePolicy ApplicationUpgradePolicyPtrInput
	// The version of the application type as defined in the application manifest.
	// This name must be the full Arm Resource ID for the referenced application type version.
	Version pulumi.StringPtrInput
}

func (ManagedClusterApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterApplicationArgs)(nil)).Elem()
}

type ManagedClusterApplicationInput interface {
	pulumi.Input

	ToManagedClusterApplicationOutput() ManagedClusterApplicationOutput
	ToManagedClusterApplicationOutputWithContext(ctx context.Context) ManagedClusterApplicationOutput
}

func (*ManagedClusterApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterApplication)(nil)).Elem()
}

func (i *ManagedClusterApplication) ToManagedClusterApplicationOutput() ManagedClusterApplicationOutput {
	return i.ToManagedClusterApplicationOutputWithContext(context.Background())
}

func (i *ManagedClusterApplication) ToManagedClusterApplicationOutputWithContext(ctx context.Context) ManagedClusterApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterApplicationOutput)
}

type ManagedClusterApplicationOutput struct{ *pulumi.OutputState }

func (ManagedClusterApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterApplication)(nil)).Elem()
}

func (o ManagedClusterApplicationOutput) ToManagedClusterApplicationOutput() ManagedClusterApplicationOutput {
	return o
}

func (o ManagedClusterApplicationOutput) ToManagedClusterApplicationOutputWithContext(ctx context.Context) ManagedClusterApplicationOutput {
	return o
}

// Describes the managed identities for an Azure resource.
func (o ManagedClusterApplicationOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// Resource location depends on the parent resource.
func (o ManagedClusterApplicationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// List of user assigned identities for the application, each mapped to a friendly name.
func (o ManagedClusterApplicationOutput) ManagedIdentities() ApplicationUserAssignedIdentityResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) ApplicationUserAssignedIdentityResponseArrayOutput {
		return v.ManagedIdentities
	}).(ApplicationUserAssignedIdentityResponseArrayOutput)
}

// Azure resource name.
func (o ManagedClusterApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of application parameters with overridden values from their default values specified in the application manifest.
func (o ManagedClusterApplicationOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// The current deployment or provisioning state, which only appears in the response
func (o ManagedClusterApplicationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ManagedClusterApplicationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o ManagedClusterApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o ManagedClusterApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Describes the policy for a monitored application upgrade.
func (o ManagedClusterApplicationOutput) UpgradePolicy() ApplicationUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) ApplicationUpgradePolicyResponsePtrOutput { return v.UpgradePolicy }).(ApplicationUpgradePolicyResponsePtrOutput)
}

// The version of the application type as defined in the application manifest.
// This name must be the full Arm Resource ID for the referenced application type version.
func (o ManagedClusterApplicationOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterApplication) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterApplicationOutput{})
}
