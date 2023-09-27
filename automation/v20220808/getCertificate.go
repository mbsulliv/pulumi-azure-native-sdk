// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieve the certificate identified by certificate name.
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:automation/v20220808:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of certificate.
	CertificateName string `pulumi:"certificateName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the certificate.
type LookupCertificateResult struct {
	// Gets the creation time.
	CreationTime string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets the expiry time of the certificate.
	ExpiryTime string `pulumi:"expiryTime"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// Gets the is exportable flag of the certificate.
	IsExportable bool `pulumi:"isExportable"`
	// Gets the last modified time.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the thumbprint of the certificate.
	Thumbprint string `pulumi:"thumbprint"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			var s LookupCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateResultOutput)
}

type LookupCertificateOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The name of certificate.
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

// Definition of the certificate.
type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCertificateResult] {
	return pulumix.Output[LookupCertificateResult]{
		OutputState: o.OutputState,
	}
}

// Gets the creation time.
func (o LookupCertificateResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets or sets the description.
func (o LookupCertificateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets the expiry time of the certificate.
func (o LookupCertificateResultOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ExpiryTime }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource
func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets the is exportable flag of the certificate.
func (o LookupCertificateResultOutput) IsExportable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCertificateResult) bool { return v.IsExportable }).(pulumi.BoolOutput)
}

// Gets the last modified time.
func (o LookupCertificateResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the thumbprint of the certificate.
func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

// The type of the resource.
func (o LookupCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
