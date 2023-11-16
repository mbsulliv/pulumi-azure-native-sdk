// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a custom certificate.
func LookupWebPubSubCustomCertificate(ctx *pulumi.Context, args *LookupWebPubSubCustomCertificateArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubCustomCertificateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebPubSubCustomCertificateResult
	err := ctx.Invoke("azure-native:webpubsub/v20230601preview:getWebPubSubCustomCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubCustomCertificateArgs struct {
	// Custom certificate name
	CertificateName string `pulumi:"certificateName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A custom certificate.
type LookupWebPubSubCustomCertificateResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Base uri of the KeyVault that stores certificate.
	KeyVaultBaseUri string `pulumi:"keyVaultBaseUri"`
	// Certificate secret name.
	KeyVaultSecretName string `pulumi:"keyVaultSecretName"`
	// Certificate secret version.
	KeyVaultSecretVersion *string `pulumi:"keyVaultSecretVersion"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWebPubSubCustomCertificateOutput(ctx *pulumi.Context, args LookupWebPubSubCustomCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubCustomCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubCustomCertificateResult, error) {
			args := v.(LookupWebPubSubCustomCertificateArgs)
			r, err := LookupWebPubSubCustomCertificate(ctx, &args, opts...)
			var s LookupWebPubSubCustomCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubCustomCertificateResultOutput)
}

type LookupWebPubSubCustomCertificateOutputArgs struct {
	// Custom certificate name
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubCustomCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubCustomCertificateArgs)(nil)).Elem()
}

// A custom certificate.
type LookupWebPubSubCustomCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubCustomCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubCustomCertificateResult)(nil)).Elem()
}

func (o LookupWebPubSubCustomCertificateResultOutput) ToLookupWebPubSubCustomCertificateResultOutput() LookupWebPubSubCustomCertificateResultOutput {
	return o
}

func (o LookupWebPubSubCustomCertificateResultOutput) ToLookupWebPubSubCustomCertificateResultOutputWithContext(ctx context.Context) LookupWebPubSubCustomCertificateResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupWebPubSubCustomCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Base uri of the KeyVault that stores certificate.
func (o LookupWebPubSubCustomCertificateResultOutput) KeyVaultBaseUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.KeyVaultBaseUri }).(pulumi.StringOutput)
}

// Certificate secret name.
func (o LookupWebPubSubCustomCertificateResultOutput) KeyVaultSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.KeyVaultSecretName }).(pulumi.StringOutput)
}

// Certificate secret version.
func (o LookupWebPubSubCustomCertificateResultOutput) KeyVaultSecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) *string { return v.KeyVaultSecretVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupWebPubSubCustomCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupWebPubSubCustomCertificateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWebPubSubCustomCertificateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWebPubSubCustomCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubCustomCertificateResultOutput{})
}
