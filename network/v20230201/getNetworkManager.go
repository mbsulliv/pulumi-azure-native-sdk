// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified Network Manager.
func LookupNetworkManager(ctx *pulumi.Context, args *LookupNetworkManagerArgs, opts ...pulumi.InvokeOption) (*LookupNetworkManagerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkManagerResult
	err := ctx.Invoke("azure-native:network/v20230201:getNetworkManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkManagerArgs struct {
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Managed Network resource
type LookupNetworkManagerResult struct {
	// A description of the network manager.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Scope Access.
	NetworkManagerScopeAccesses []string `pulumi:"networkManagerScopeAccesses"`
	// Scope of Network Manager.
	NetworkManagerScopes NetworkManagerPropertiesResponseNetworkManagerScopes `pulumi:"networkManagerScopes"`
	// The provisioning state of the network manager resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupNetworkManagerOutput(ctx *pulumi.Context, args LookupNetworkManagerOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkManagerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkManagerResult, error) {
			args := v.(LookupNetworkManagerArgs)
			r, err := LookupNetworkManager(ctx, &args, opts...)
			var s LookupNetworkManagerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkManagerResultOutput)
}

type LookupNetworkManagerOutputArgs struct {
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkManagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkManagerArgs)(nil)).Elem()
}

// The Managed Network resource
type LookupNetworkManagerResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkManagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkManagerResult)(nil)).Elem()
}

func (o LookupNetworkManagerResultOutput) ToLookupNetworkManagerResultOutput() LookupNetworkManagerResultOutput {
	return o
}

func (o LookupNetworkManagerResultOutput) ToLookupNetworkManagerResultOutputWithContext(ctx context.Context) LookupNetworkManagerResultOutput {
	return o
}

func (o LookupNetworkManagerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNetworkManagerResult] {
	return pulumix.Output[LookupNetworkManagerResult]{
		OutputState: o.OutputState,
	}
}

// A description of the network manager.
func (o LookupNetworkManagerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkManagerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNetworkManagerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupNetworkManagerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNetworkManagerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Scope Access.
func (o LookupNetworkManagerResultOutput) NetworkManagerScopeAccesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) []string { return v.NetworkManagerScopeAccesses }).(pulumi.StringArrayOutput)
}

// Scope of Network Manager.
func (o LookupNetworkManagerResultOutput) NetworkManagerScopes() NetworkManagerPropertiesResponseNetworkManagerScopesOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) NetworkManagerPropertiesResponseNetworkManagerScopes {
		return v.NetworkManagerScopes
	}).(NetworkManagerPropertiesResponseNetworkManagerScopesOutput)
}

// The provisioning state of the network manager resource.
func (o LookupNetworkManagerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o LookupNetworkManagerResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupNetworkManagerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkManagerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNetworkManagerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkManagerResultOutput{})
}
