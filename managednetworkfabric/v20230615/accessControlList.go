// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Access Control List resource definition.
type AccessControlList struct {
	pulumi.CustomResourceState

	// Access Control List file URL.
	AclsUrl pulumi.StringPtrOutput `pulumi:"aclsUrl"`
	// Administrative state of the resource.
	AdministrativeState pulumi.StringOutput `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// Configuration state of the resource.
	ConfigurationState pulumi.StringOutput `pulumi:"configurationState"`
	// Input method to configure Access Control List.
	ConfigurationType pulumi.StringOutput `pulumi:"configurationType"`
	// Default action that needs to be applied when no condition is matched. Example: Permit | Deny.
	DefaultAction pulumi.StringPtrOutput `pulumi:"defaultAction"`
	// List of dynamic match configurations.
	DynamicMatchConfigurations CommonDynamicMatchConfigurationResponseArrayOutput `pulumi:"dynamicMatchConfigurations"`
	// The last synced timestamp.
	LastSyncedTime pulumi.StringOutput `pulumi:"lastSyncedTime"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// List of match configurations.
	MatchConfigurations AccessControlListMatchConfigurationResponseArrayOutput `pulumi:"matchConfigurations"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccessControlList registers a new resource with the given unique name, arguments, and options.
func NewAccessControlList(ctx *pulumi.Context,
	name string, args *AccessControlListArgs, opts ...pulumi.ResourceOption) (*AccessControlList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationType == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DefaultAction == nil {
		args.DefaultAction = pulumi.StringPtr("Permit")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:AccessControlList"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230201preview:AccessControlList"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AccessControlList
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230615:AccessControlList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessControlList gets an existing AccessControlList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessControlList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessControlListState, opts ...pulumi.ResourceOption) (*AccessControlList, error) {
	var resource AccessControlList
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230615:AccessControlList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessControlList resources.
type accessControlListState struct {
}

type AccessControlListState struct {
}

func (AccessControlListState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessControlListState)(nil)).Elem()
}

type accessControlListArgs struct {
	// Name of the Access Control List.
	AccessControlListName *string `pulumi:"accessControlListName"`
	// Access Control List file URL.
	AclsUrl *string `pulumi:"aclsUrl"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Input method to configure Access Control List.
	ConfigurationType string `pulumi:"configurationType"`
	// Default action that needs to be applied when no condition is matched. Example: Permit | Deny.
	DefaultAction *string `pulumi:"defaultAction"`
	// List of dynamic match configurations.
	DynamicMatchConfigurations []CommonDynamicMatchConfiguration `pulumi:"dynamicMatchConfigurations"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// List of match configurations.
	MatchConfigurations []AccessControlListMatchConfiguration `pulumi:"matchConfigurations"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AccessControlList resource.
type AccessControlListArgs struct {
	// Name of the Access Control List.
	AccessControlListName pulumi.StringPtrInput
	// Access Control List file URL.
	AclsUrl pulumi.StringPtrInput
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Input method to configure Access Control List.
	ConfigurationType pulumi.StringInput
	// Default action that needs to be applied when no condition is matched. Example: Permit | Deny.
	DefaultAction pulumi.StringPtrInput
	// List of dynamic match configurations.
	DynamicMatchConfigurations CommonDynamicMatchConfigurationArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// List of match configurations.
	MatchConfigurations AccessControlListMatchConfigurationArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AccessControlListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessControlListArgs)(nil)).Elem()
}

type AccessControlListInput interface {
	pulumi.Input

	ToAccessControlListOutput() AccessControlListOutput
	ToAccessControlListOutputWithContext(ctx context.Context) AccessControlListOutput
}

func (*AccessControlList) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessControlList)(nil)).Elem()
}

func (i *AccessControlList) ToAccessControlListOutput() AccessControlListOutput {
	return i.ToAccessControlListOutputWithContext(context.Background())
}

func (i *AccessControlList) ToAccessControlListOutputWithContext(ctx context.Context) AccessControlListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessControlListOutput)
}

func (i *AccessControlList) ToOutput(ctx context.Context) pulumix.Output[*AccessControlList] {
	return pulumix.Output[*AccessControlList]{
		OutputState: i.ToAccessControlListOutputWithContext(ctx).OutputState,
	}
}

type AccessControlListOutput struct{ *pulumi.OutputState }

func (AccessControlListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessControlList)(nil)).Elem()
}

func (o AccessControlListOutput) ToAccessControlListOutput() AccessControlListOutput {
	return o
}

func (o AccessControlListOutput) ToAccessControlListOutputWithContext(ctx context.Context) AccessControlListOutput {
	return o
}

func (o AccessControlListOutput) ToOutput(ctx context.Context) pulumix.Output[*AccessControlList] {
	return pulumix.Output[*AccessControlList]{
		OutputState: o.OutputState,
	}
}

// Access Control List file URL.
func (o AccessControlListOutput) AclsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringPtrOutput { return v.AclsUrl }).(pulumi.StringPtrOutput)
}

// Administrative state of the resource.
func (o AccessControlListOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o AccessControlListOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Configuration state of the resource.
func (o AccessControlListOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.ConfigurationState }).(pulumi.StringOutput)
}

// Input method to configure Access Control List.
func (o AccessControlListOutput) ConfigurationType() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.ConfigurationType }).(pulumi.StringOutput)
}

// Default action that needs to be applied when no condition is matched. Example: Permit | Deny.
func (o AccessControlListOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringPtrOutput { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

// List of dynamic match configurations.
func (o AccessControlListOutput) DynamicMatchConfigurations() CommonDynamicMatchConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *AccessControlList) CommonDynamicMatchConfigurationResponseArrayOutput {
		return v.DynamicMatchConfigurations
	}).(CommonDynamicMatchConfigurationResponseArrayOutput)
}

// The last synced timestamp.
func (o AccessControlListOutput) LastSyncedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.LastSyncedTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o AccessControlListOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// List of match configurations.
func (o AccessControlListOutput) MatchConfigurations() AccessControlListMatchConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *AccessControlList) AccessControlListMatchConfigurationResponseArrayOutput {
		return v.MatchConfigurations
	}).(AccessControlListMatchConfigurationResponseArrayOutput)
}

// The name of the resource
func (o AccessControlListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o AccessControlListOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AccessControlListOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AccessControlList) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AccessControlListOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AccessControlListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessControlListOutput{})
}
