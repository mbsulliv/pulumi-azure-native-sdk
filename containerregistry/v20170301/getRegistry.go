// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the properties of the specified container registry.
func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryResult
	err := ctx.Invoke("azure-native:containerregistry/v20170301:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a container registry.
type LookupRegistryResult struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `pulumi:"adminUserEnabled"`
	// The creation date of the container registry in ISO8601 format.
	CreationDate string `pulumi:"creationDate"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer string `pulumi:"loginServer"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning state of the container registry at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SKU of the container registry.
	Sku SkuResponse `pulumi:"sku"`
	// The properties of the storage account for the container registry.
	StorageAccount *StorageAccountPropertiesResponse `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryResult
func (val *LookupRegistryResult) Defaults() *LookupRegistryResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AdminUserEnabled == nil {
		adminUserEnabled_ := false
		tmp.AdminUserEnabled = &adminUserEnabled_
	}
	return &tmp
}

func LookupRegistryOutput(ctx *pulumi.Context, args LookupRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryResult, error) {
			args := v.(LookupRegistryArgs)
			r, err := LookupRegistry(ctx, &args, opts...)
			var s LookupRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryResultOutput)
}

type LookupRegistryOutputArgs struct {
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryArgs)(nil)).Elem()
}

// An object that represents a container registry.
type LookupRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryResult)(nil)).Elem()
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutput() LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutputWithContext(ctx context.Context) LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRegistryResult] {
	return pulumix.Output[LookupRegistryResult]{
		OutputState: o.OutputState,
	}
}

// The value that indicates whether the admin user is enabled.
func (o LookupRegistryResultOutput) AdminUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *bool { return v.AdminUserEnabled }).(pulumi.BoolPtrOutput)
}

// The creation date of the container registry in ISO8601 format.
func (o LookupRegistryResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The resource ID.
func (o LookupRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o LookupRegistryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Location }).(pulumi.StringOutput)
}

// The URL that can be used to log into the container registry.
func (o LookupRegistryResultOutput) LoginServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.LoginServer }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the container registry at the time the operation was called.
func (o LookupRegistryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the container registry.
func (o LookupRegistryResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// The properties of the storage account for the container registry.
func (o LookupRegistryResultOutput) StorageAccount() StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *StorageAccountPropertiesResponse { return v.StorageAccount }).(StorageAccountPropertiesResponsePtrOutput)
}

// The tags of the resource.
func (o LookupRegistryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegistryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryResultOutput{})
}
