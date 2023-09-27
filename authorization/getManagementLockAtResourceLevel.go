// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the management lock of a resource or any level below resource.
// Azure REST API version: 2020-05-01.
func LookupManagementLockAtResourceLevel(ctx *pulumi.Context, args *LookupManagementLockAtResourceLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtResourceLevelResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagementLockAtResourceLevelResult
	err := ctx.Invoke("azure-native:authorization:getManagementLockAtResourceLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtResourceLevelArgs struct {
	// The name of lock.
	LockName string `pulumi:"lockName"`
	// An extra path parameter needed in some services, like SQL Databases.
	ParentResourcePath string `pulumi:"parentResourcePath"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
	// The namespace of the resource provider.
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	// The type of the resource.
	ResourceType string `pulumi:"resourceType"`
}

// The lock information.
type LookupManagementLockAtResourceLevelResult struct {
	// The resource ID of the lock.
	Id string `pulumi:"id"`
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level string `pulumi:"level"`
	// The name of the lock.
	Name string `pulumi:"name"`
	// Notes about the lock. Maximum of 512 characters.
	Notes *string `pulumi:"notes"`
	// The owners of the lock.
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type string `pulumi:"type"`
}

func LookupManagementLockAtResourceLevelOutput(ctx *pulumi.Context, args LookupManagementLockAtResourceLevelOutputArgs, opts ...pulumi.InvokeOption) LookupManagementLockAtResourceLevelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementLockAtResourceLevelResult, error) {
			args := v.(LookupManagementLockAtResourceLevelArgs)
			r, err := LookupManagementLockAtResourceLevel(ctx, &args, opts...)
			var s LookupManagementLockAtResourceLevelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementLockAtResourceLevelResultOutput)
}

type LookupManagementLockAtResourceLevelOutputArgs struct {
	// The name of lock.
	LockName pulumi.StringInput `pulumi:"lockName"`
	// An extra path parameter needed in some services, like SQL Databases.
	ParentResourcePath pulumi.StringInput `pulumi:"parentResourcePath"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The namespace of the resource provider.
	ResourceProviderNamespace pulumi.StringInput `pulumi:"resourceProviderNamespace"`
	// The type of the resource.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupManagementLockAtResourceLevelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockAtResourceLevelArgs)(nil)).Elem()
}

// The lock information.
type LookupManagementLockAtResourceLevelResultOutput struct{ *pulumi.OutputState }

func (LookupManagementLockAtResourceLevelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockAtResourceLevelResult)(nil)).Elem()
}

func (o LookupManagementLockAtResourceLevelResultOutput) ToLookupManagementLockAtResourceLevelResultOutput() LookupManagementLockAtResourceLevelResultOutput {
	return o
}

func (o LookupManagementLockAtResourceLevelResultOutput) ToLookupManagementLockAtResourceLevelResultOutputWithContext(ctx context.Context) LookupManagementLockAtResourceLevelResultOutput {
	return o
}

func (o LookupManagementLockAtResourceLevelResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagementLockAtResourceLevelResult] {
	return pulumix.Output[LookupManagementLockAtResourceLevelResult]{
		OutputState: o.OutputState,
	}
}

// The resource ID of the lock.
func (o LookupManagementLockAtResourceLevelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) string { return v.Id }).(pulumi.StringOutput)
}

// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
func (o LookupManagementLockAtResourceLevelResultOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) string { return v.Level }).(pulumi.StringOutput)
}

// The name of the lock.
func (o LookupManagementLockAtResourceLevelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) string { return v.Name }).(pulumi.StringOutput)
}

// Notes about the lock. Maximum of 512 characters.
func (o LookupManagementLockAtResourceLevelResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The owners of the lock.
func (o LookupManagementLockAtResourceLevelResultOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) []ManagementLockOwnerResponse { return v.Owners }).(ManagementLockOwnerResponseArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupManagementLockAtResourceLevelResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type of the lock - Microsoft.Authorization/locks.
func (o LookupManagementLockAtResourceLevelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementLockAtResourceLevelResultOutput{})
}
