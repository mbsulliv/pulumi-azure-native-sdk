// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an integration account certificate.
func LookupIntegrationAccountCertificate(ctx *pulumi.Context, args *LookupIntegrationAccountCertificateArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountCertificateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationAccountCertificateResult
	err := ctx.Invoke("azure-native:logic/v20150801preview:getIntegrationAccountCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountCertificateArgs struct {
	// The integration account certificate name.
	CertificateName string `pulumi:"certificateName"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupIntegrationAccountCertificateResult struct {
	// The changed time.
	ChangedTime string `pulumi:"changedTime"`
	// The created time.
	CreatedTime string `pulumi:"createdTime"`
	// The resource id.
	Id *string `pulumi:"id"`
	// The key details in the key vault.
	Key *KeyVaultKeyReferenceResponse `pulumi:"key"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The public certificate.
	PublicCertificate *string `pulumi:"publicCertificate"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

func LookupIntegrationAccountCertificateOutput(ctx *pulumi.Context, args LookupIntegrationAccountCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountCertificateResult, error) {
			args := v.(LookupIntegrationAccountCertificateArgs)
			r, err := LookupIntegrationAccountCertificate(ctx, &args, opts...)
			var s LookupIntegrationAccountCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountCertificateResultOutput)
}

type LookupIntegrationAccountCertificateOutputArgs struct {
	// The integration account certificate name.
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// The integration account name.
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountCertificateArgs)(nil)).Elem()
}

type LookupIntegrationAccountCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountCertificateResult)(nil)).Elem()
}

func (o LookupIntegrationAccountCertificateResultOutput) ToLookupIntegrationAccountCertificateResultOutput() LookupIntegrationAccountCertificateResultOutput {
	return o
}

func (o LookupIntegrationAccountCertificateResultOutput) ToLookupIntegrationAccountCertificateResultOutputWithContext(ctx context.Context) LookupIntegrationAccountCertificateResultOutput {
	return o
}

func (o LookupIntegrationAccountCertificateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIntegrationAccountCertificateResult] {
	return pulumix.Output[LookupIntegrationAccountCertificateResult]{
		OutputState: o.OutputState,
	}
}

// The changed time.
func (o LookupIntegrationAccountCertificateResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

// The created time.
func (o LookupIntegrationAccountCertificateResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// The resource id.
func (o LookupIntegrationAccountCertificateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The key details in the key vault.
func (o LookupIntegrationAccountCertificateResultOutput) Key() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *KeyVaultKeyReferenceResponse { return v.Key }).(KeyVaultKeyReferenceResponsePtrOutput)
}

// The resource location.
func (o LookupIntegrationAccountCertificateResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The metadata.
func (o LookupIntegrationAccountCertificateResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The resource name.
func (o LookupIntegrationAccountCertificateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The public certificate.
func (o LookupIntegrationAccountCertificateResultOutput) PublicCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *string { return v.PublicCertificate }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o LookupIntegrationAccountCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupIntegrationAccountCertificateResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountCertificateResultOutput{})
}
