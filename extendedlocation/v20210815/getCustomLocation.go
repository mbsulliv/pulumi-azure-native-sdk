// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210815

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the customLocation with a specified resource group and name.
func LookupCustomLocation(ctx *pulumi.Context, args *LookupCustomLocationArgs, opts ...pulumi.InvokeOption) (*LookupCustomLocationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomLocationResult
	err := ctx.Invoke("azure-native:extendedlocation/v20210815:getCustomLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomLocationArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Custom Locations name.
	ResourceName string `pulumi:"resourceName"`
}

// Custom Locations definition.
type LookupCustomLocationResult struct {
	// This is optional input that contains the authentication that should be used to generate the namespace.
	Authentication *CustomLocationPropertiesResponseAuthentication `pulumi:"authentication"`
	// Contains the reference to the add-on that contains charts to deploy CRDs and operators.
	ClusterExtensionIds []string `pulumi:"clusterExtensionIds"`
	// Display name for the Custom Locations location.
	DisplayName *string `pulumi:"displayName"`
	// Connected Cluster or AKS Cluster. The Custom Locations RP will perform a checkAccess API for listAdminCredentials permissions.
	HostResourceId *string `pulumi:"hostResourceId"`
	// Type of host the Custom Locations is referencing (Kubernetes, etc...).
	HostType *string `pulumi:"hostType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Kubernetes namespace that will be created on the specified cluster.
	Namespace *string `pulumi:"namespace"`
	// Provisioning State for the Custom Location.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCustomLocationOutput(ctx *pulumi.Context, args LookupCustomLocationOutputArgs, opts ...pulumi.InvokeOption) LookupCustomLocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomLocationResult, error) {
			args := v.(LookupCustomLocationArgs)
			r, err := LookupCustomLocation(ctx, &args, opts...)
			var s LookupCustomLocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomLocationResultOutput)
}

type LookupCustomLocationOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Custom Locations name.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupCustomLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomLocationArgs)(nil)).Elem()
}

// Custom Locations definition.
type LookupCustomLocationResultOutput struct{ *pulumi.OutputState }

func (LookupCustomLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomLocationResult)(nil)).Elem()
}

func (o LookupCustomLocationResultOutput) ToLookupCustomLocationResultOutput() LookupCustomLocationResultOutput {
	return o
}

func (o LookupCustomLocationResultOutput) ToLookupCustomLocationResultOutputWithContext(ctx context.Context) LookupCustomLocationResultOutput {
	return o
}

// This is optional input that contains the authentication that should be used to generate the namespace.
func (o LookupCustomLocationResultOutput) Authentication() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *CustomLocationPropertiesResponseAuthentication {
		return v.Authentication
	}).(CustomLocationPropertiesResponseAuthenticationPtrOutput)
}

// Contains the reference to the add-on that contains charts to deploy CRDs and operators.
func (o LookupCustomLocationResultOutput) ClusterExtensionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) []string { return v.ClusterExtensionIds }).(pulumi.StringArrayOutput)
}

// Display name for the Custom Locations location.
func (o LookupCustomLocationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Connected Cluster or AKS Cluster. The Custom Locations RP will perform a checkAccess API for listAdminCredentials permissions.
func (o LookupCustomLocationResultOutput) HostResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.HostResourceId }).(pulumi.StringPtrOutput)
}

// Type of host the Custom Locations is referencing (Kubernetes, etc...).
func (o LookupCustomLocationResultOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.HostType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCustomLocationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupCustomLocationResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupCustomLocationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCustomLocationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Kubernetes namespace that will be created on the specified cluster.
func (o LookupCustomLocationResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Provisioning State for the Custom Location.
func (o LookupCustomLocationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o LookupCustomLocationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupCustomLocationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCustomLocationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomLocationResultOutput{})
}
