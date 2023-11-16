// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managementpartner

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the management partner using the partnerId, objectId and tenantId.
// Azure REST API version: 2018-02-01.
func LookupPartner(ctx *pulumi.Context, args *LookupPartnerArgs, opts ...pulumi.InvokeOption) (*LookupPartnerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPartnerResult
	err := ctx.Invoke("azure-native:managementpartner:getPartner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerArgs struct {
	// Id of the Partner
	PartnerId string `pulumi:"partnerId"`
}

// this is the management partner operations response
type LookupPartnerResult struct {
	// This is the DateTime when the partner was created.
	CreatedTime *string `pulumi:"createdTime"`
	// Type of the partner
	Etag *int `pulumi:"etag"`
	// Identifier of the partner
	Id string `pulumi:"id"`
	// Name of the partner
	Name string `pulumi:"name"`
	// This is the object id.
	ObjectId *string `pulumi:"objectId"`
	// This is the partner id
	PartnerId *string `pulumi:"partnerId"`
	// This is the partner name
	PartnerName *string `pulumi:"partnerName"`
	// This is the tenant id.
	TenantId *string `pulumi:"tenantId"`
	// Type of resource. "Microsoft.ManagementPartner/partners"
	Type string `pulumi:"type"`
	// This is the DateTime when the partner was updated.
	UpdatedTime *string `pulumi:"updatedTime"`
	// This is the version.
	Version *int `pulumi:"version"`
}

func LookupPartnerOutput(ctx *pulumi.Context, args LookupPartnerOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerResult, error) {
			args := v.(LookupPartnerArgs)
			r, err := LookupPartner(ctx, &args, opts...)
			var s LookupPartnerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerResultOutput)
}

type LookupPartnerOutputArgs struct {
	// Id of the Partner
	PartnerId pulumi.StringInput `pulumi:"partnerId"`
}

func (LookupPartnerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerArgs)(nil)).Elem()
}

// this is the management partner operations response
type LookupPartnerResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerResult)(nil)).Elem()
}

func (o LookupPartnerResultOutput) ToLookupPartnerResultOutput() LookupPartnerResultOutput {
	return o
}

func (o LookupPartnerResultOutput) ToLookupPartnerResultOutputWithContext(ctx context.Context) LookupPartnerResultOutput {
	return o
}

// This is the DateTime when the partner was created.
func (o LookupPartnerResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// Type of the partner
func (o LookupPartnerResultOutput) Etag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *int { return v.Etag }).(pulumi.IntPtrOutput)
}

// Identifier of the partner
func (o LookupPartnerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the partner
func (o LookupPartnerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Name }).(pulumi.StringOutput)
}

// This is the object id.
func (o LookupPartnerResultOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

// This is the partner id
func (o LookupPartnerResultOutput) PartnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.PartnerId }).(pulumi.StringPtrOutput)
}

// This is the partner name
func (o LookupPartnerResultOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.PartnerName }).(pulumi.StringPtrOutput)
}

// This is the tenant id.
func (o LookupPartnerResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Type of resource. "Microsoft.ManagementPartner/partners"
func (o LookupPartnerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Type }).(pulumi.StringOutput)
}

// This is the DateTime when the partner was updated.
func (o LookupPartnerResultOutput) UpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.UpdatedTime }).(pulumi.StringPtrOutput)
}

// This is the version.
func (o LookupPartnerResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerResultOutput{})
}
