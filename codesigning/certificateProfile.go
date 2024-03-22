// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codesigning

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Certificate profile resource.
// Azure REST API version: 2024-02-05-preview.
type CertificateProfile struct {
	pulumi.CustomResourceState

	// Used as L in the certificate subject name.
	City pulumi.StringOutput `pulumi:"city"`
	// Used as CN in the certificate subject name.
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// Used as C in the certificate subject name.
	Country pulumi.StringOutput `pulumi:"country"`
	// Enhanced key usage of the certificate.
	EnhancedKeyUsage pulumi.StringOutput `pulumi:"enhancedKeyUsage"`
	// Identity validation id used for the certificate subject name.
	IdentityValidationId pulumi.StringPtrOutput `pulumi:"identityValidationId"`
	// Whether to include L in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCity pulumi.BoolPtrOutput `pulumi:"includeCity"`
	// Whether to include C in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCountry pulumi.BoolPtrOutput `pulumi:"includeCountry"`
	// Whether to include PC in the certificate subject name.
	IncludePostalCode pulumi.BoolPtrOutput `pulumi:"includePostalCode"`
	// Whether to include S in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeState pulumi.BoolPtrOutput `pulumi:"includeState"`
	// Whether to include STREET in the certificate subject name.
	IncludeStreetAddress pulumi.BoolPtrOutput `pulumi:"includeStreetAddress"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Used as O in the certificate subject name.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// Used as OU in the private trust certificate subject name.
	OrganizationUnit pulumi.StringOutput `pulumi:"organizationUnit"`
	// Used as PC in the certificate subject name.
	PostalCode pulumi.StringOutput `pulumi:"postalCode"`
	// Profile type of the certificate.
	ProfileType pulumi.StringOutput `pulumi:"profileType"`
	// Status of the current operation on certificate profile.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Used as S in the certificate subject name.
	State pulumi.StringOutput `pulumi:"state"`
	// Status of the certificate profile.
	Status pulumi.StringOutput `pulumi:"status"`
	// Used as STREET in the certificate subject name.
	StreetAddress pulumi.StringOutput `pulumi:"streetAddress"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCertificateProfile registers a new resource with the given unique name, arguments, and options.
func NewCertificateProfile(ctx *pulumi.Context,
	name string, args *CertificateProfileArgs, opts ...pulumi.ResourceOption) (*CertificateProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ProfileType == nil {
		return nil, errors.New("invalid value for required argument 'ProfileType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.IncludeCity == nil {
		args.IncludeCity = pulumi.BoolPtr(false)
	}
	if args.IncludeCountry == nil {
		args.IncludeCountry = pulumi.BoolPtr(false)
	}
	if args.IncludePostalCode == nil {
		args.IncludePostalCode = pulumi.BoolPtr(false)
	}
	if args.IncludeState == nil {
		args.IncludeState = pulumi.BoolPtr(false)
	}
	if args.IncludeStreetAddress == nil {
		args.IncludeStreetAddress = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:codesigning/v20240205preview:CertificateProfile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CertificateProfile
	err := ctx.RegisterResource("azure-native:codesigning:CertificateProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateProfile gets an existing CertificateProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateProfileState, opts ...pulumi.ResourceOption) (*CertificateProfile, error) {
	var resource CertificateProfile
	err := ctx.ReadResource("azure-native:codesigning:CertificateProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateProfile resources.
type certificateProfileState struct {
}

type CertificateProfileState struct {
}

func (CertificateProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateProfileState)(nil)).Elem()
}

type certificateProfileArgs struct {
	// Trusted Signing account name.
	AccountName string `pulumi:"accountName"`
	// Identity validation id used for the certificate subject name.
	IdentityValidationId *string `pulumi:"identityValidationId"`
	// Whether to include L in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCity *bool `pulumi:"includeCity"`
	// Whether to include C in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCountry *bool `pulumi:"includeCountry"`
	// Whether to include PC in the certificate subject name.
	IncludePostalCode *bool `pulumi:"includePostalCode"`
	// Whether to include S in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeState *bool `pulumi:"includeState"`
	// Whether to include STREET in the certificate subject name.
	IncludeStreetAddress *bool `pulumi:"includeStreetAddress"`
	// Certificate profile name.
	ProfileName *string `pulumi:"profileName"`
	// Profile type of the certificate.
	ProfileType string `pulumi:"profileType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a CertificateProfile resource.
type CertificateProfileArgs struct {
	// Trusted Signing account name.
	AccountName pulumi.StringInput
	// Identity validation id used for the certificate subject name.
	IdentityValidationId pulumi.StringPtrInput
	// Whether to include L in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCity pulumi.BoolPtrInput
	// Whether to include C in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCountry pulumi.BoolPtrInput
	// Whether to include PC in the certificate subject name.
	IncludePostalCode pulumi.BoolPtrInput
	// Whether to include S in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeState pulumi.BoolPtrInput
	// Whether to include STREET in the certificate subject name.
	IncludeStreetAddress pulumi.BoolPtrInput
	// Certificate profile name.
	ProfileName pulumi.StringPtrInput
	// Profile type of the certificate.
	ProfileType pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (CertificateProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateProfileArgs)(nil)).Elem()
}

type CertificateProfileInput interface {
	pulumi.Input

	ToCertificateProfileOutput() CertificateProfileOutput
	ToCertificateProfileOutputWithContext(ctx context.Context) CertificateProfileOutput
}

func (*CertificateProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProfile)(nil)).Elem()
}

func (i *CertificateProfile) ToCertificateProfileOutput() CertificateProfileOutput {
	return i.ToCertificateProfileOutputWithContext(context.Background())
}

func (i *CertificateProfile) ToCertificateProfileOutputWithContext(ctx context.Context) CertificateProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateProfileOutput)
}

type CertificateProfileOutput struct{ *pulumi.OutputState }

func (CertificateProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProfile)(nil)).Elem()
}

func (o CertificateProfileOutput) ToCertificateProfileOutput() CertificateProfileOutput {
	return o
}

func (o CertificateProfileOutput) ToCertificateProfileOutputWithContext(ctx context.Context) CertificateProfileOutput {
	return o
}

// Used as L in the certificate subject name.
func (o CertificateProfileOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.City }).(pulumi.StringOutput)
}

// Used as CN in the certificate subject name.
func (o CertificateProfileOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.CommonName }).(pulumi.StringOutput)
}

// Used as C in the certificate subject name.
func (o CertificateProfileOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// Enhanced key usage of the certificate.
func (o CertificateProfileOutput) EnhancedKeyUsage() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.EnhancedKeyUsage }).(pulumi.StringOutput)
}

// Identity validation id used for the certificate subject name.
func (o CertificateProfileOutput) IdentityValidationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringPtrOutput { return v.IdentityValidationId }).(pulumi.StringPtrOutput)
}

// Whether to include L in the certificate subject name. Applicable only for private trust, private trust ci profile types
func (o CertificateProfileOutput) IncludeCity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.BoolPtrOutput { return v.IncludeCity }).(pulumi.BoolPtrOutput)
}

// Whether to include C in the certificate subject name. Applicable only for private trust, private trust ci profile types
func (o CertificateProfileOutput) IncludeCountry() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.BoolPtrOutput { return v.IncludeCountry }).(pulumi.BoolPtrOutput)
}

// Whether to include PC in the certificate subject name.
func (o CertificateProfileOutput) IncludePostalCode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.BoolPtrOutput { return v.IncludePostalCode }).(pulumi.BoolPtrOutput)
}

// Whether to include S in the certificate subject name. Applicable only for private trust, private trust ci profile types
func (o CertificateProfileOutput) IncludeState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.BoolPtrOutput { return v.IncludeState }).(pulumi.BoolPtrOutput)
}

// Whether to include STREET in the certificate subject name.
func (o CertificateProfileOutput) IncludeStreetAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.BoolPtrOutput { return v.IncludeStreetAddress }).(pulumi.BoolPtrOutput)
}

// The name of the resource
func (o CertificateProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Used as O in the certificate subject name.
func (o CertificateProfileOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// Used as OU in the private trust certificate subject name.
func (o CertificateProfileOutput) OrganizationUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.OrganizationUnit }).(pulumi.StringOutput)
}

// Used as PC in the certificate subject name.
func (o CertificateProfileOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.PostalCode }).(pulumi.StringOutput)
}

// Profile type of the certificate.
func (o CertificateProfileOutput) ProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.ProfileType }).(pulumi.StringOutput)
}

// Status of the current operation on certificate profile.
func (o CertificateProfileOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Used as S in the certificate subject name.
func (o CertificateProfileOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Status of the certificate profile.
func (o CertificateProfileOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Used as STREET in the certificate subject name.
func (o CertificateProfileOutput) StreetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.StreetAddress }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CertificateProfileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CertificateProfile) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CertificateProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateProfileOutput{})
}
