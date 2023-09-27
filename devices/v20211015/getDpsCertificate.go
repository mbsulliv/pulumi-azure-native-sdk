// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the certificate from the provisioning service.
func LookupDpsCertificate(ctx *pulumi.Context, args *LookupDpsCertificateArgs, opts ...pulumi.InvokeOption) (*LookupDpsCertificateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDpsCertificateResult
	err := ctx.Invoke("azure-native:devices/v20211015:getDpsCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDpsCertificateArgs struct {
	// Name of the certificate to retrieve.
	CertificateName string `pulumi:"certificateName"`
	// Name of the provisioning service the certificate is associated with.
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	// Resource group identifier.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The X509 Certificate.
type LookupDpsCertificateResult struct {
	// The entity tag.
	Etag string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The name of the certificate.
	Name string `pulumi:"name"`
	// properties of a certificate
	Properties CertificatePropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupDpsCertificateOutput(ctx *pulumi.Context, args LookupDpsCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupDpsCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDpsCertificateResult, error) {
			args := v.(LookupDpsCertificateArgs)
			r, err := LookupDpsCertificate(ctx, &args, opts...)
			var s LookupDpsCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDpsCertificateResultOutput)
}

type LookupDpsCertificateOutputArgs struct {
	// Name of the certificate to retrieve.
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// Name of the provisioning service the certificate is associated with.
	ProvisioningServiceName pulumi.StringInput `pulumi:"provisioningServiceName"`
	// Resource group identifier.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDpsCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDpsCertificateArgs)(nil)).Elem()
}

// The X509 Certificate.
type LookupDpsCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupDpsCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDpsCertificateResult)(nil)).Elem()
}

func (o LookupDpsCertificateResultOutput) ToLookupDpsCertificateResultOutput() LookupDpsCertificateResultOutput {
	return o
}

func (o LookupDpsCertificateResultOutput) ToLookupDpsCertificateResultOutputWithContext(ctx context.Context) LookupDpsCertificateResultOutput {
	return o
}

func (o LookupDpsCertificateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDpsCertificateResult] {
	return pulumix.Output[LookupDpsCertificateResult]{
		OutputState: o.OutputState,
	}
}

// The entity tag.
func (o LookupDpsCertificateResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupDpsCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the certificate.
func (o LookupDpsCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// properties of a certificate
func (o LookupDpsCertificateResultOutput) Properties() CertificatePropertiesResponseOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) CertificatePropertiesResponse { return v.Properties }).(CertificatePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDpsCertificateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type.
func (o LookupDpsCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDpsCertificateResultOutput{})
}
