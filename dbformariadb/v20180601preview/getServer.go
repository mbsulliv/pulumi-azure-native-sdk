// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a server.
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbformariadb/v20180601preview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a server.
type LookupServerResult struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `pulumi:"fullyQualifiedDomainName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The master server id of a replica server.
	MasterServerId *string `pulumi:"masterServerId"`
	// Enforce a minimal Tls version for the server.
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The maximum number of replicas that a master server can have.
	ReplicaCapacity *int `pulumi:"replicaCapacity"`
	// The replication role of the server.
	ReplicationRole *string `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku *SkuResponse `pulumi:"sku"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Storage profile of a server.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// A state of a server that is visible to user.
	UserVisibleState *string `pulumi:"userVisibleState"`
	// Server version.
	Version *string `pulumi:"version"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			var s LookupServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerResultOutput)
}

type LookupServerOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

// Represents a server.
type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
func (o LookupServerResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// Earliest restore point creation time (ISO8601 format)
func (o LookupServerResultOutput) EarliestRestoreDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.EarliestRestoreDate }).(pulumi.StringPtrOutput)
}

// The fully qualified domain name of a server.
func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.FullyQualifiedDomainName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Active Directory identity of the server.
func (o LookupServerResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The master server id of a replica server.
func (o LookupServerResultOutput) MasterServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.MasterServerId }).(pulumi.StringPtrOutput)
}

// Enforce a minimal Tls version for the server.
func (o LookupServerResultOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The maximum number of replicas that a master server can have.
func (o LookupServerResultOutput) ReplicaCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *int { return v.ReplicaCapacity }).(pulumi.IntPtrOutput)
}

// The replication role of the server.
func (o LookupServerResultOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

// The SKU (pricing tier) of the server.
func (o LookupServerResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Enable ssl enforcement or not when connect to server.
func (o LookupServerResultOutput) SslEnforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SslEnforcement }).(pulumi.StringPtrOutput)
}

// Storage profile of a server.
func (o LookupServerResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Resource tags.
func (o LookupServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Type }).(pulumi.StringOutput)
}

// A state of a server that is visible to user.
func (o LookupServerResultOutput) UserVisibleState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.UserVisibleState }).(pulumi.StringPtrOutput)
}

// Server version.
func (o LookupServerResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
