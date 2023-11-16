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
func LookupSignalRCustomCertificate(ctx *pulumi.Context, args *LookupSignalRCustomCertificateArgs, opts ...pulumi.InvokeOption) (*LookupSignalRCustomCertificateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSignalRCustomCertificateResult
	err := ctx.Invoke("azure-native:signalrservice/v20230601preview:getSignalRCustomCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRCustomCertificateArgs struct {
	// Custom certificate name
	CertificateName string `pulumi:"certificateName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A custom certificate.
type LookupSignalRCustomCertificateResult struct {
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

func LookupSignalRCustomCertificateOutput(ctx *pulumi.Context, args LookupSignalRCustomCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRCustomCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRCustomCertificateResult, error) {
			args := v.(LookupSignalRCustomCertificateArgs)
			r, err := LookupSignalRCustomCertificate(ctx, &args, opts...)
			var s LookupSignalRCustomCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRCustomCertificateResultOutput)
}

type LookupSignalRCustomCertificateOutputArgs struct {
	// Custom certificate name
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalRCustomCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRCustomCertificateArgs)(nil)).Elem()
}

// A custom certificate.
type LookupSignalRCustomCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRCustomCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRCustomCertificateResult)(nil)).Elem()
}

func (o LookupSignalRCustomCertificateResultOutput) ToLookupSignalRCustomCertificateResultOutput() LookupSignalRCustomCertificateResultOutput {
	return o
}

func (o LookupSignalRCustomCertificateResultOutput) ToLookupSignalRCustomCertificateResultOutputWithContext(ctx context.Context) LookupSignalRCustomCertificateResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSignalRCustomCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Base uri of the KeyVault that stores certificate.
func (o LookupSignalRCustomCertificateResultOutput) KeyVaultBaseUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.KeyVaultBaseUri }).(pulumi.StringOutput)
}

// Certificate secret name.
func (o LookupSignalRCustomCertificateResultOutput) KeyVaultSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.KeyVaultSecretName }).(pulumi.StringOutput)
}

// Certificate secret version.
func (o LookupSignalRCustomCertificateResultOutput) KeyVaultSecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) *string { return v.KeyVaultSecretVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupSignalRCustomCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupSignalRCustomCertificateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSignalRCustomCertificateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSignalRCustomCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRCustomCertificateResultOutput{})
}
