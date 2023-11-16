// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get baseboard management controller key set of the provided cluster.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01.
func LookupBmcKeySet(ctx *pulumi.Context, args *LookupBmcKeySetArgs, opts ...pulumi.InvokeOption) (*LookupBmcKeySetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBmcKeySetResult
	err := ctx.Invoke("azure-native:networkcloud:getBmcKeySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBmcKeySetArgs struct {
	// The name of the baseboard management controller key set.
	BmcKeySetName string `pulumi:"bmcKeySetName"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupBmcKeySetResult struct {
	// The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
	AzureGroupId string `pulumi:"azureGroupId"`
	// The more detailed status of the key set.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The date and time after which the users in this key set will be removed from the baseboard management controllers.
	Expiration string `pulumi:"expiration"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The last time this key set was validated.
	LastValidation string `pulumi:"lastValidation"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The access level allowed for the users in this key set.
	PrivilegeLevel string `pulumi:"privilegeLevel"`
	// The provisioning state of the baseboard management controller key set.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The unique list of permitted users.
	UserList []KeySetUserResponse `pulumi:"userList"`
	// The status evaluation of each user.
	UserListStatus []KeySetUserStatusResponse `pulumi:"userListStatus"`
}

func LookupBmcKeySetOutput(ctx *pulumi.Context, args LookupBmcKeySetOutputArgs, opts ...pulumi.InvokeOption) LookupBmcKeySetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBmcKeySetResult, error) {
			args := v.(LookupBmcKeySetArgs)
			r, err := LookupBmcKeySet(ctx, &args, opts...)
			var s LookupBmcKeySetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBmcKeySetResultOutput)
}

type LookupBmcKeySetOutputArgs struct {
	// The name of the baseboard management controller key set.
	BmcKeySetName pulumi.StringInput `pulumi:"bmcKeySetName"`
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBmcKeySetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBmcKeySetArgs)(nil)).Elem()
}

type LookupBmcKeySetResultOutput struct{ *pulumi.OutputState }

func (LookupBmcKeySetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBmcKeySetResult)(nil)).Elem()
}

func (o LookupBmcKeySetResultOutput) ToLookupBmcKeySetResultOutput() LookupBmcKeySetResultOutput {
	return o
}

func (o LookupBmcKeySetResultOutput) ToLookupBmcKeySetResultOutputWithContext(ctx context.Context) LookupBmcKeySetResultOutput {
	return o
}

// The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
func (o LookupBmcKeySetResultOutput) AzureGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.AzureGroupId }).(pulumi.StringOutput)
}

// The more detailed status of the key set.
func (o LookupBmcKeySetResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupBmcKeySetResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The date and time after which the users in this key set will be removed from the baseboard management controllers.
func (o LookupBmcKeySetResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.Expiration }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupBmcKeySetResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupBmcKeySetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last time this key set was validated.
func (o LookupBmcKeySetResultOutput) LastValidation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.LastValidation }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupBmcKeySetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBmcKeySetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The access level allowed for the users in this key set.
func (o LookupBmcKeySetResultOutput) PrivilegeLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.PrivilegeLevel }).(pulumi.StringOutput)
}

// The provisioning state of the baseboard management controller key set.
func (o LookupBmcKeySetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBmcKeySetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupBmcKeySetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBmcKeySetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique list of permitted users.
func (o LookupBmcKeySetResultOutput) UserList() KeySetUserResponseArrayOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) []KeySetUserResponse { return v.UserList }).(KeySetUserResponseArrayOutput)
}

// The status evaluation of each user.
func (o LookupBmcKeySetResultOutput) UserListStatus() KeySetUserStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupBmcKeySetResult) []KeySetUserStatusResponse { return v.UserListStatus }).(KeySetUserStatusResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBmcKeySetResultOutput{})
}
