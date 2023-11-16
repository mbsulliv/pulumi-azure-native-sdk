// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azureactivedirectory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Guest Usages resource for the Microsoft.AzureActiveDirectory resource provider
// Azure REST API version: 2021-04-01.
//
// Other available API versions: 2023-01-18-preview, 2023-05-17-preview.
func LookupGuestUsage(ctx *pulumi.Context, args *LookupGuestUsageArgs, opts ...pulumi.InvokeOption) (*LookupGuestUsageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGuestUsageResult
	err := ctx.Invoke("azure-native:azureactivedirectory:getGuestUsage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestUsageArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The initial domain name of the Azure AD B2C tenant.
	ResourceName string `pulumi:"resourceName"`
}

// Guest Usages Resource
type LookupGuestUsageResult struct {
	// An identifier that represents the Guest Usages resource.
	Id string `pulumi:"id"`
	// Location of the Guest Usages resource.
	Location *string `pulumi:"location"`
	// The name of the Guest Usages resource.
	Name string `pulumi:"name"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
	// An identifier for the tenant for which the resource is being created
	TenantId *string `pulumi:"tenantId"`
	// The type of the Guest Usages resource.
	Type string `pulumi:"type"`
}

func LookupGuestUsageOutput(ctx *pulumi.Context, args LookupGuestUsageOutputArgs, opts ...pulumi.InvokeOption) LookupGuestUsageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestUsageResult, error) {
			args := v.(LookupGuestUsageArgs)
			r, err := LookupGuestUsage(ctx, &args, opts...)
			var s LookupGuestUsageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestUsageResultOutput)
}

type LookupGuestUsageOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The initial domain name of the Azure AD B2C tenant.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupGuestUsageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestUsageArgs)(nil)).Elem()
}

// Guest Usages Resource
type LookupGuestUsageResultOutput struct{ *pulumi.OutputState }

func (LookupGuestUsageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestUsageResult)(nil)).Elem()
}

func (o LookupGuestUsageResultOutput) ToLookupGuestUsageResultOutput() LookupGuestUsageResultOutput {
	return o
}

func (o LookupGuestUsageResultOutput) ToLookupGuestUsageResultOutputWithContext(ctx context.Context) LookupGuestUsageResultOutput {
	return o
}

// An identifier that represents the Guest Usages resource.
func (o LookupGuestUsageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the Guest Usages resource.
func (o LookupGuestUsageResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the Guest Usages resource.
func (o LookupGuestUsageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupGuestUsageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Key-value pairs of additional resource provisioning properties.
func (o LookupGuestUsageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// An identifier for the tenant for which the resource is being created
func (o LookupGuestUsageResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the Guest Usages resource.
func (o LookupGuestUsageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestUsageResultOutput{})
}
