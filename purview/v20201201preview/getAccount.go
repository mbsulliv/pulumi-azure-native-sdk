// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get an account
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:purview/v20201201preview:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAccountArgs struct {
	// The name of the account.
	AccountName string `pulumi:"accountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Account resource
type LookupAccountResult struct {
	// Cloud connectors.
	// External cloud identifier used as part of scanning configuration.
	CloudConnectors *CloudConnectorsResponse `pulumi:"cloudConnectors"`
	// Gets the time at which the entity was created.
	CreatedAt string `pulumi:"createdAt"`
	// Gets the creator of the entity.
	CreatedBy string `pulumi:"createdBy"`
	// Gets the creators of the entity's object id.
	CreatedByObjectId string `pulumi:"createdByObjectId"`
	// The URIs that are the public endpoints of the account.
	Endpoints AccountPropertiesResponseEndpoints `pulumi:"endpoints"`
	// Gets or sets the friendly name.
	FriendlyName string `pulumi:"friendlyName"`
	// Gets or sets the identifier.
	Id string `pulumi:"id"`
	// Identity Info on the tracked resource
	Identity *IdentityResponse `pulumi:"identity"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Gets or sets the managed resource group name
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Gets the resource identifiers of the managed resources.
	ManagedResources AccountPropertiesResponseManagedResources `pulumi:"managedResources"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Gets the private endpoint connections information.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Gets or sets the state of the provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets or sets the public network access.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Gets or sets the Sku.
	Sku AccountResponseSku `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData AccountPropertiesResponseSystemData `pulumi:"systemData"`
	// Tags on the azure resource.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAccountResult
func (val *LookupAccountResult) Defaults() *LookupAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	// The name of the account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

// Account resource
type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAccountResult] {
	return pulumix.Output[LookupAccountResult]{
		OutputState: o.OutputState,
	}
}

// Cloud connectors.
// External cloud identifier used as part of scanning configuration.
func (o LookupAccountResultOutput) CloudConnectors() CloudConnectorsResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *CloudConnectorsResponse { return v.CloudConnectors }).(CloudConnectorsResponsePtrOutput)
}

// Gets the time at which the entity was created.
func (o LookupAccountResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Gets the creator of the entity.
func (o LookupAccountResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// Gets the creators of the entity's object id.
func (o LookupAccountResultOutput) CreatedByObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CreatedByObjectId }).(pulumi.StringOutput)
}

// The URIs that are the public endpoints of the account.
func (o LookupAccountResultOutput) Endpoints() AccountPropertiesResponseEndpointsOutput {
	return o.ApplyT(func(v LookupAccountResult) AccountPropertiesResponseEndpoints { return v.Endpoints }).(AccountPropertiesResponseEndpointsOutput)
}

// Gets or sets the friendly name.
func (o LookupAccountResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

// Gets or sets the identifier.
func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity Info on the tracked resource
func (o LookupAccountResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Gets or sets the location.
func (o LookupAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets the managed resource group name
func (o LookupAccountResultOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

// Gets the resource identifiers of the managed resources.
func (o LookupAccountResultOutput) ManagedResources() AccountPropertiesResponseManagedResourcesOutput {
	return o.ApplyT(func(v LookupAccountResult) AccountPropertiesResponseManagedResources { return v.ManagedResources }).(AccountPropertiesResponseManagedResourcesOutput)
}

// Gets or sets the name.
func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the private endpoint connections information.
func (o LookupAccountResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Gets or sets the state of the provisioning.
func (o LookupAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the public network access.
func (o LookupAccountResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Gets or sets the Sku.
func (o LookupAccountResultOutput) Sku() AccountResponseSkuOutput {
	return o.ApplyT(func(v LookupAccountResult) AccountResponseSku { return v.Sku }).(AccountResponseSkuOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupAccountResultOutput) SystemData() AccountPropertiesResponseSystemDataOutput {
	return o.ApplyT(func(v LookupAccountResult) AccountPropertiesResponseSystemData { return v.SystemData }).(AccountPropertiesResponseSystemDataOutput)
}

// Tags on the azure resource.
func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type.
func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
