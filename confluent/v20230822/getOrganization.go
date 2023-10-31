// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230822

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Organization resource.
func LookupOrganization(ctx *pulumi.Context, args *LookupOrganizationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationResult
	err := ctx.Invoke("azure-native:confluent/v20230822:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationArgs struct {
	// Organization resource name
	OrganizationName string `pulumi:"organizationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Organization resource.
type LookupOrganizationResult struct {
	// The creation time of the resource.
	CreatedTime string `pulumi:"createdTime"`
	// The ARM id of the resource.
	Id string `pulumi:"id"`
	// Location of Organization resource
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Confluent offer detail
	OfferDetail OfferDetailResponse `pulumi:"offerDetail"`
	// Id of the Confluent organization.
	OrganizationId string `pulumi:"organizationId"`
	// Provision states for confluent RP
	ProvisioningState string `pulumi:"provisioningState"`
	// SSO url for the Confluent organization.
	SsoUrl string `pulumi:"ssoUrl"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Organization resource tags
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Subscriber detail
	UserDetail UserDetailResponse `pulumi:"userDetail"`
}

func LookupOrganizationOutput(ctx *pulumi.Context, args LookupOrganizationOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationResult, error) {
			args := v.(LookupOrganizationArgs)
			r, err := LookupOrganization(ctx, &args, opts...)
			var s LookupOrganizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationResultOutput)
}

type LookupOrganizationOutputArgs struct {
	// Organization resource name
	OrganizationName pulumi.StringInput `pulumi:"organizationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationArgs)(nil)).Elem()
}

// Organization resource.
type LookupOrganizationResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationResult)(nil)).Elem()
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutput() LookupOrganizationResultOutput {
	return o
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutputWithContext(ctx context.Context) LookupOrganizationResultOutput {
	return o
}

func (o LookupOrganizationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOrganizationResult] {
	return pulumix.Output[LookupOrganizationResult]{
		OutputState: o.OutputState,
	}
}

// The creation time of the resource.
func (o LookupOrganizationResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// The ARM id of the resource.
func (o LookupOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of Organization resource
func (o LookupOrganizationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Confluent offer detail
func (o LookupOrganizationResultOutput) OfferDetail() OfferDetailResponseOutput {
	return o.ApplyT(func(v LookupOrganizationResult) OfferDetailResponse { return v.OfferDetail }).(OfferDetailResponseOutput)
}

// Id of the Confluent organization.
func (o LookupOrganizationResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// Provision states for confluent RP
func (o LookupOrganizationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SSO url for the Confluent organization.
func (o LookupOrganizationResultOutput) SsoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.SsoUrl }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o LookupOrganizationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOrganizationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Organization resource tags
func (o LookupOrganizationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrganizationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupOrganizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Subscriber detail
func (o LookupOrganizationResultOutput) UserDetail() UserDetailResponseOutput {
	return o.ApplyT(func(v LookupOrganizationResult) UserDetailResponse { return v.UserDetail }).(UserDetailResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationResultOutput{})
}