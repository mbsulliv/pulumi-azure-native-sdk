// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Lists impacted resources in the tenant by an event (Security Advisory).
func ListSecurityAdvisoryImpactedResourceByTenantIdAndEventId(ctx *pulumi.Context, args *ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdArgs, opts ...pulumi.InvokeOption) (*ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult
	err := ctx.Invoke("azure-native:resourcehealth/v20230701preview:listSecurityAdvisoryImpactedResourceByTenantIdAndEventId", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdArgs struct {
	// Event Id which uniquely identifies ServiceHealth event.
	EventTrackingId string `pulumi:"eventTrackingId"`
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string `pulumi:"filter"`
}

// The List of eventImpactedResources operation response.
type ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult struct {
	// The URI to fetch the next page of events. Call ListNext() with this URI to fetch the next page of impacted resource.
	NextLink *string `pulumi:"nextLink"`
	// The list of eventImpactedResources.
	Value []EventImpactedResourceResponse `pulumi:"value"`
}

func ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdOutput(ctx *pulumi.Context, args ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdOutputArgs, opts ...pulumi.InvokeOption) ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult, error) {
			args := v.(ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdArgs)
			r, err := ListSecurityAdvisoryImpactedResourceByTenantIdAndEventId(ctx, &args, opts...)
			var s ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput)
}

type ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdOutputArgs struct {
	// Event Id which uniquely identifies ServiceHealth event.
	EventTrackingId pulumi.StringInput `pulumi:"eventTrackingId"`
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter pulumi.StringPtrInput `pulumi:"filter"`
}

func (ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdArgs)(nil)).Elem()
}

// The List of eventImpactedResources operation response.
type ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput struct{ *pulumi.OutputState }

func (ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult)(nil)).Elem()
}

func (o ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput) ToListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput() ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput {
	return o
}

func (o ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput) ToListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutputWithContext(ctx context.Context) ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput {
	return o
}

func (o ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult] {
	return pulumix.Output[ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult]{
		OutputState: o.OutputState,
	}
}

// The URI to fetch the next page of events. Call ListNext() with this URI to fetch the next page of impacted resource.
func (o ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// The list of eventImpactedResources.
func (o ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput) Value() EventImpactedResourceResponseArrayOutput {
	return o.ApplyT(func(v ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult) []EventImpactedResourceResponse {
		return v.Value
	}).(EventImpactedResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResultOutput{})
}
