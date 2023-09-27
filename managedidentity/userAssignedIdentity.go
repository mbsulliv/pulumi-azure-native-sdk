// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managedidentity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Describes an identity resource.
// Azure REST API version: 2023-01-31. Prior API version in Azure Native 1.x: 2018-11-30
type UserAssignedIdentity struct {
	pulumi.CustomResourceState

	// The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the service principal object associated with the created identity.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The id of the tenant which the identity belongs to.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUserAssignedIdentity registers a new resource with the given unique name, arguments, and options.
func NewUserAssignedIdentity(ctx *pulumi.Context,
	name string, args *UserAssignedIdentityArgs, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managedidentity/v20150831preview:UserAssignedIdentity"),
		},
		{
			Type: pulumi.String("azure-native:managedidentity/v20181130:UserAssignedIdentity"),
		},
		{
			Type: pulumi.String("azure-native:managedidentity/v20210930preview:UserAssignedIdentity"),
		},
		{
			Type: pulumi.String("azure-native:managedidentity/v20220131preview:UserAssignedIdentity"),
		},
		{
			Type: pulumi.String("azure-native:managedidentity/v20230131:UserAssignedIdentity"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource UserAssignedIdentity
	err := ctx.RegisterResource("azure-native:managedidentity:UserAssignedIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserAssignedIdentity gets an existing UserAssignedIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAssignedIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAssignedIdentityState, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	var resource UserAssignedIdentity
	err := ctx.ReadResource("azure-native:managedidentity:UserAssignedIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserAssignedIdentity resources.
type userAssignedIdentityState struct {
}

type UserAssignedIdentityState struct {
}

func (UserAssignedIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityState)(nil)).Elem()
}

type userAssignedIdentityArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the Resource Group to which the identity belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the identity resource.
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a UserAssignedIdentity resource.
type UserAssignedIdentityArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the Resource Group to which the identity belongs.
	ResourceGroupName pulumi.StringInput
	// The name of the identity resource.
	ResourceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityArgs)(nil)).Elem()
}

type UserAssignedIdentityInput interface {
	pulumi.Input

	ToUserAssignedIdentityOutput() UserAssignedIdentityOutput
	ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput
}

func (*UserAssignedIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentity)(nil)).Elem()
}

func (i *UserAssignedIdentity) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return i.ToUserAssignedIdentityOutputWithContext(context.Background())
}

func (i *UserAssignedIdentity) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityOutput)
}

func (i *UserAssignedIdentity) ToOutput(ctx context.Context) pulumix.Output[*UserAssignedIdentity] {
	return pulumix.Output[*UserAssignedIdentity]{
		OutputState: i.ToUserAssignedIdentityOutputWithContext(ctx).OutputState,
	}
}

type UserAssignedIdentityOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToOutput(ctx context.Context) pulumix.Output[*UserAssignedIdentity] {
	return pulumix.Output[*UserAssignedIdentity]{
		OutputState: o.OutputState,
	}
}

// The id of the app associated with the identity. This is a random generated UUID by MSI.
func (o UserAssignedIdentityOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o UserAssignedIdentityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o UserAssignedIdentityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the service principal object associated with the created identity.
func (o UserAssignedIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o UserAssignedIdentityOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o UserAssignedIdentityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The id of the tenant which the identity belongs to.
func (o UserAssignedIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o UserAssignedIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserAssignedIdentityOutput{})
}
