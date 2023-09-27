// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type BmcKeySet struct {
	pulumi.CustomResourceState

	// The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
	AzureGroupId pulumi.StringOutput `pulumi:"azureGroupId"`
	// The more detailed status of the key set.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The date and time after which the users in this key set will be removed from the baseboard management controllers.
	Expiration pulumi.StringOutput `pulumi:"expiration"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The last time this key set was validated.
	LastValidation pulumi.StringOutput `pulumi:"lastValidation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The access level allowed for the users in this key set.
	PrivilegeLevel pulumi.StringOutput `pulumi:"privilegeLevel"`
	// The provisioning state of the baseboard management controller key set.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique list of permitted users.
	UserList KeySetUserResponseArrayOutput `pulumi:"userList"`
	// The status evaluation of each user.
	UserListStatus KeySetUserStatusResponseArrayOutput `pulumi:"userListStatus"`
}

// NewBmcKeySet registers a new resource with the given unique name, arguments, and options.
func NewBmcKeySet(ctx *pulumi.Context,
	name string, args *BmcKeySetArgs, opts ...pulumi.ResourceOption) (*BmcKeySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureGroupId == nil {
		return nil, errors.New("invalid value for required argument 'AzureGroupId'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Expiration == nil {
		return nil, errors.New("invalid value for required argument 'Expiration'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.PrivilegeLevel == nil {
		return nil, errors.New("invalid value for required argument 'PrivilegeLevel'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserList == nil {
		return nil, errors.New("invalid value for required argument 'UserList'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud:BmcKeySet"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:BmcKeySet"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230501preview:BmcKeySet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BmcKeySet
	err := ctx.RegisterResource("azure-native:networkcloud/v20230701:BmcKeySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBmcKeySet gets an existing BmcKeySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBmcKeySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BmcKeySetState, opts ...pulumi.ResourceOption) (*BmcKeySet, error) {
	var resource BmcKeySet
	err := ctx.ReadResource("azure-native:networkcloud/v20230701:BmcKeySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BmcKeySet resources.
type bmcKeySetState struct {
}

type BmcKeySetState struct {
}

func (BmcKeySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*bmcKeySetState)(nil)).Elem()
}

type bmcKeySetArgs struct {
	// The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
	AzureGroupId string `pulumi:"azureGroupId"`
	// The name of the baseboard management controller key set.
	BmcKeySetName *string `pulumi:"bmcKeySetName"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The date and time after which the users in this key set will be removed from the baseboard management controllers.
	Expiration string `pulumi:"expiration"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The access level allowed for the users in this key set.
	PrivilegeLevel string `pulumi:"privilegeLevel"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The unique list of permitted users.
	UserList []KeySetUser `pulumi:"userList"`
}

// The set of arguments for constructing a BmcKeySet resource.
type BmcKeySetArgs struct {
	// The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
	AzureGroupId pulumi.StringInput
	// The name of the baseboard management controller key set.
	BmcKeySetName pulumi.StringPtrInput
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// The date and time after which the users in this key set will be removed from the baseboard management controllers.
	Expiration pulumi.StringInput
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The access level allowed for the users in this key set.
	PrivilegeLevel pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The unique list of permitted users.
	UserList KeySetUserArrayInput
}

func (BmcKeySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bmcKeySetArgs)(nil)).Elem()
}

type BmcKeySetInput interface {
	pulumi.Input

	ToBmcKeySetOutput() BmcKeySetOutput
	ToBmcKeySetOutputWithContext(ctx context.Context) BmcKeySetOutput
}

func (*BmcKeySet) ElementType() reflect.Type {
	return reflect.TypeOf((**BmcKeySet)(nil)).Elem()
}

func (i *BmcKeySet) ToBmcKeySetOutput() BmcKeySetOutput {
	return i.ToBmcKeySetOutputWithContext(context.Background())
}

func (i *BmcKeySet) ToBmcKeySetOutputWithContext(ctx context.Context) BmcKeySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BmcKeySetOutput)
}

func (i *BmcKeySet) ToOutput(ctx context.Context) pulumix.Output[*BmcKeySet] {
	return pulumix.Output[*BmcKeySet]{
		OutputState: i.ToBmcKeySetOutputWithContext(ctx).OutputState,
	}
}

type BmcKeySetOutput struct{ *pulumi.OutputState }

func (BmcKeySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BmcKeySet)(nil)).Elem()
}

func (o BmcKeySetOutput) ToBmcKeySetOutput() BmcKeySetOutput {
	return o
}

func (o BmcKeySetOutput) ToBmcKeySetOutputWithContext(ctx context.Context) BmcKeySetOutput {
	return o
}

func (o BmcKeySetOutput) ToOutput(ctx context.Context) pulumix.Output[*BmcKeySet] {
	return pulumix.Output[*BmcKeySet]{
		OutputState: o.OutputState,
	}
}

// The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
func (o BmcKeySetOutput) AzureGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.AzureGroupId }).(pulumi.StringOutput)
}

// The more detailed status of the key set.
func (o BmcKeySetOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o BmcKeySetOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The date and time after which the users in this key set will be removed from the baseboard management controllers.
func (o BmcKeySetOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.Expiration }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o BmcKeySetOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *BmcKeySet) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The last time this key set was validated.
func (o BmcKeySetOutput) LastValidation() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.LastValidation }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o BmcKeySetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o BmcKeySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The access level allowed for the users in this key set.
func (o BmcKeySetOutput) PrivilegeLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.PrivilegeLevel }).(pulumi.StringOutput)
}

// The provisioning state of the baseboard management controller key set.
func (o BmcKeySetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BmcKeySetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BmcKeySet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o BmcKeySetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BmcKeySetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BmcKeySet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique list of permitted users.
func (o BmcKeySetOutput) UserList() KeySetUserResponseArrayOutput {
	return o.ApplyT(func(v *BmcKeySet) KeySetUserResponseArrayOutput { return v.UserList }).(KeySetUserResponseArrayOutput)
}

// The status evaluation of each user.
func (o BmcKeySetOutput) UserListStatus() KeySetUserStatusResponseArrayOutput {
	return o.ApplyT(func(v *BmcKeySet) KeySetUserStatusResponseArrayOutput { return v.UserListStatus }).(KeySetUserStatusResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(BmcKeySetOutput{})
}
