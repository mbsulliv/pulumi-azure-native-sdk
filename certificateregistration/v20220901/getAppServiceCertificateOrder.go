// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Get a certificate order.
func LookupAppServiceCertificateOrder(ctx *pulumi.Context, args *LookupAppServiceCertificateOrderArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceCertificateOrderResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAppServiceCertificateOrderResult
	err := ctx.Invoke("azure-native:certificateregistration/v20220901:getAppServiceCertificateOrder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppServiceCertificateOrderArgs struct {
	// Name of the certificate order..
	CertificateOrderName string `pulumi:"certificateOrderName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// SSL certificate purchase order.
type LookupAppServiceCertificateOrderResult struct {
	// Reasons why App Service Certificate is not renewable at the current moment.
	AppServiceCertificateNotRenewableReasons []string `pulumi:"appServiceCertificateNotRenewableReasons"`
	// <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
	AutoRenew *bool `pulumi:"autoRenew"`
	// State of the Key Vault secret.
	Certificates map[string]AppServiceCertificateResponse `pulumi:"certificates"`
	// Contact info
	Contact CertificateOrderContactResponse `pulumi:"contact"`
	// Last CSR that was created for this order.
	Csr *string `pulumi:"csr"`
	// Certificate distinguished name.
	DistinguishedName *string `pulumi:"distinguishedName"`
	// Domain verification token.
	DomainVerificationToken string `pulumi:"domainVerificationToken"`
	// Certificate expiration time.
	ExpirationTime string `pulumi:"expirationTime"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Intermediate certificate.
	Intermediate CertificateDetailsResponse `pulumi:"intermediate"`
	// <code>true</code> if private key is external; otherwise, <code>false</code>.
	IsPrivateKeyExternal bool `pulumi:"isPrivateKeyExternal"`
	// Certificate key size.
	KeySize *int `pulumi:"keySize"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Certificate last issuance time.
	LastCertificateIssuanceTime string `pulumi:"lastCertificateIssuanceTime"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Time stamp when the certificate would be auto renewed next
	NextAutoRenewalTimeStamp string `pulumi:"nextAutoRenewalTimeStamp"`
	// Certificate product type.
	ProductType string `pulumi:"productType"`
	// Status of certificate order.
	ProvisioningState string `pulumi:"provisioningState"`
	// Root certificate.
	Root CertificateDetailsResponse `pulumi:"root"`
	// Current serial number of the certificate.
	SerialNumber string `pulumi:"serialNumber"`
	// Signed certificate.
	SignedCertificate CertificateDetailsResponse `pulumi:"signedCertificate"`
	// Current order status.
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// Duration in years (must be 1).
	ValidityInYears *int `pulumi:"validityInYears"`
}

// Defaults sets the appropriate defaults for LookupAppServiceCertificateOrderResult
func (val *LookupAppServiceCertificateOrderResult) Defaults() *LookupAppServiceCertificateOrderResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AutoRenew == nil {
		autoRenew_ := true
		tmp.AutoRenew = &autoRenew_
	}
	if tmp.KeySize == nil {
		keySize_ := 2048
		tmp.KeySize = &keySize_
	}
	if tmp.ValidityInYears == nil {
		validityInYears_ := 1
		tmp.ValidityInYears = &validityInYears_
	}
	return &tmp
}

func LookupAppServiceCertificateOrderOutput(ctx *pulumi.Context, args LookupAppServiceCertificateOrderOutputArgs, opts ...pulumi.InvokeOption) LookupAppServiceCertificateOrderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServiceCertificateOrderResult, error) {
			args := v.(LookupAppServiceCertificateOrderArgs)
			r, err := LookupAppServiceCertificateOrder(ctx, &args, opts...)
			var s LookupAppServiceCertificateOrderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServiceCertificateOrderResultOutput)
}

type LookupAppServiceCertificateOrderOutputArgs struct {
	// Name of the certificate order..
	CertificateOrderName pulumi.StringInput `pulumi:"certificateOrderName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServiceCertificateOrderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceCertificateOrderArgs)(nil)).Elem()
}

// SSL certificate purchase order.
type LookupAppServiceCertificateOrderResultOutput struct{ *pulumi.OutputState }

func (LookupAppServiceCertificateOrderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceCertificateOrderResult)(nil)).Elem()
}

func (o LookupAppServiceCertificateOrderResultOutput) ToLookupAppServiceCertificateOrderResultOutput() LookupAppServiceCertificateOrderResultOutput {
	return o
}

func (o LookupAppServiceCertificateOrderResultOutput) ToLookupAppServiceCertificateOrderResultOutputWithContext(ctx context.Context) LookupAppServiceCertificateOrderResultOutput {
	return o
}

func (o LookupAppServiceCertificateOrderResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAppServiceCertificateOrderResult] {
	return pulumix.Output[LookupAppServiceCertificateOrderResult]{
		OutputState: o.OutputState,
	}
}

// Reasons why App Service Certificate is not renewable at the current moment.
func (o LookupAppServiceCertificateOrderResultOutput) AppServiceCertificateNotRenewableReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) []string {
		return v.AppServiceCertificateNotRenewableReasons
	}).(pulumi.StringArrayOutput)
}

// <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
func (o LookupAppServiceCertificateOrderResultOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// State of the Key Vault secret.
func (o LookupAppServiceCertificateOrderResultOutput) Certificates() AppServiceCertificateResponseMapOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) map[string]AppServiceCertificateResponse {
		return v.Certificates
	}).(AppServiceCertificateResponseMapOutput)
}

// Contact info
func (o LookupAppServiceCertificateOrderResultOutput) Contact() CertificateOrderContactResponseOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) CertificateOrderContactResponse { return v.Contact }).(CertificateOrderContactResponseOutput)
}

// Last CSR that was created for this order.
func (o LookupAppServiceCertificateOrderResultOutput) Csr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *string { return v.Csr }).(pulumi.StringPtrOutput)
}

// Certificate distinguished name.
func (o LookupAppServiceCertificateOrderResultOutput) DistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *string { return v.DistinguishedName }).(pulumi.StringPtrOutput)
}

// Domain verification token.
func (o LookupAppServiceCertificateOrderResultOutput) DomainVerificationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.DomainVerificationToken }).(pulumi.StringOutput)
}

// Certificate expiration time.
func (o LookupAppServiceCertificateOrderResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

// Resource Id.
func (o LookupAppServiceCertificateOrderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Id }).(pulumi.StringOutput)
}

// Intermediate certificate.
func (o LookupAppServiceCertificateOrderResultOutput) Intermediate() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) CertificateDetailsResponse { return v.Intermediate }).(CertificateDetailsResponseOutput)
}

// <code>true</code> if private key is external; otherwise, <code>false</code>.
func (o LookupAppServiceCertificateOrderResultOutput) IsPrivateKeyExternal() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) bool { return v.IsPrivateKeyExternal }).(pulumi.BoolOutput)
}

// Certificate key size.
func (o LookupAppServiceCertificateOrderResultOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *int { return v.KeySize }).(pulumi.IntPtrOutput)
}

// Kind of resource.
func (o LookupAppServiceCertificateOrderResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Certificate last issuance time.
func (o LookupAppServiceCertificateOrderResultOutput) LastCertificateIssuanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.LastCertificateIssuanceTime }).(pulumi.StringOutput)
}

// Resource Location.
func (o LookupAppServiceCertificateOrderResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name.
func (o LookupAppServiceCertificateOrderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Name }).(pulumi.StringOutput)
}

// Time stamp when the certificate would be auto renewed next
func (o LookupAppServiceCertificateOrderResultOutput) NextAutoRenewalTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.NextAutoRenewalTimeStamp }).(pulumi.StringOutput)
}

// Certificate product type.
func (o LookupAppServiceCertificateOrderResultOutput) ProductType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.ProductType }).(pulumi.StringOutput)
}

// Status of certificate order.
func (o LookupAppServiceCertificateOrderResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Root certificate.
func (o LookupAppServiceCertificateOrderResultOutput) Root() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) CertificateDetailsResponse { return v.Root }).(CertificateDetailsResponseOutput)
}

// Current serial number of the certificate.
func (o LookupAppServiceCertificateOrderResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

// Signed certificate.
func (o LookupAppServiceCertificateOrderResultOutput) SignedCertificate() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) CertificateDetailsResponse { return v.SignedCertificate }).(CertificateDetailsResponseOutput)
}

// Current order status.
func (o LookupAppServiceCertificateOrderResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupAppServiceCertificateOrderResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupAppServiceCertificateOrderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Type }).(pulumi.StringOutput)
}

// Duration in years (must be 1).
func (o LookupAppServiceCertificateOrderResultOutput) ValidityInYears() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *int { return v.ValidityInYears }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServiceCertificateOrderResultOutput{})
}
