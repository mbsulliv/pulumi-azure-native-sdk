// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AccessControlList resource definition.
type AccessControlList struct {
	pulumi.CustomResourceState

	// IP address family. Example: ipv4 | ipv6.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// Access Control List conditions.
	Conditions AccessControlListConditionPropertiesResponseArrayOutput `pulumi:"conditions"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state of the resource.
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

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.Conditions == nil {
		return nil, errors.New("invalid value for required argument 'Conditions'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:AccessControlList"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230615:AccessControlList"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AccessControlList
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230201preview:AccessControlList", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230201preview:AccessControlList", name, id, state, &resource, opts...)
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
	// Name of the Access Control List
	AccessControlListName *string `pulumi:"accessControlListName"`
	// IP address family. Example: ipv4 | ipv6.
	AddressFamily string `pulumi:"addressFamily"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Access Control List conditions.
	Conditions []AccessControlListConditionProperties `pulumi:"conditions"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AccessControlList resource.
type AccessControlListArgs struct {
	// Name of the Access Control List
	AccessControlListName pulumi.StringPtrInput
	// IP address family. Example: ipv4 | ipv6.
	AddressFamily pulumi.StringInput
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Access Control List conditions.
	Conditions AccessControlListConditionPropertiesArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
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

// IP address family. Example: ipv4 | ipv6.
func (o AccessControlListOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o AccessControlListOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Access Control List conditions.
func (o AccessControlListOutput) Conditions() AccessControlListConditionPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *AccessControlList) AccessControlListConditionPropertiesResponseArrayOutput {
		return v.Conditions
	}).(AccessControlListConditionPropertiesResponseArrayOutput)
}

// The geo-location where the resource lives
func (o AccessControlListOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AccessControlListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessControlList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
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
