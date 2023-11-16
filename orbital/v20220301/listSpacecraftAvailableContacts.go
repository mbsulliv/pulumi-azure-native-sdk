// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns list of available contacts. A contact is available if the spacecraft is visible from the ground station for more than the minimum viable contact duration provided in the contact profile.
func ListSpacecraftAvailableContacts(ctx *pulumi.Context, args *ListSpacecraftAvailableContactsArgs, opts ...pulumi.InvokeOption) (*ListSpacecraftAvailableContactsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSpacecraftAvailableContactsResult
	err := ctx.Invoke("azure-native:orbital/v20220301:listSpacecraftAvailableContacts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSpacecraftAvailableContactsArgs struct {
	// The reference to the contact profile resource.
	ContactProfile ContactParametersContactProfile `pulumi:"contactProfile"`
	// End time of a contact (ISO 8601 UTC standard).
	EndTime string `pulumi:"endTime"`
	// Name of Azure Ground Station.
	GroundStationName string `pulumi:"groundStationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Spacecraft ID.
	SpacecraftName string `pulumi:"spacecraftName"`
	// Start time of a contact (ISO 8601 UTC standard).
	StartTime string `pulumi:"startTime"`
}

// Response for the ListAvailableContacts API service call.
type ListSpacecraftAvailableContactsResult struct {
	// The URL to get the next set of results.
	NextLink string `pulumi:"nextLink"`
	// A list of available contacts.
	Value []AvailableContactsResponse `pulumi:"value"`
}

func ListSpacecraftAvailableContactsOutput(ctx *pulumi.Context, args ListSpacecraftAvailableContactsOutputArgs, opts ...pulumi.InvokeOption) ListSpacecraftAvailableContactsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSpacecraftAvailableContactsResult, error) {
			args := v.(ListSpacecraftAvailableContactsArgs)
			r, err := ListSpacecraftAvailableContacts(ctx, &args, opts...)
			var s ListSpacecraftAvailableContactsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSpacecraftAvailableContactsResultOutput)
}

type ListSpacecraftAvailableContactsOutputArgs struct {
	// The reference to the contact profile resource.
	ContactProfile ContactParametersContactProfileInput `pulumi:"contactProfile"`
	// End time of a contact (ISO 8601 UTC standard).
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// Name of Azure Ground Station.
	GroundStationName pulumi.StringInput `pulumi:"groundStationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Spacecraft ID.
	SpacecraftName pulumi.StringInput `pulumi:"spacecraftName"`
	// Start time of a contact (ISO 8601 UTC standard).
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (ListSpacecraftAvailableContactsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSpacecraftAvailableContactsArgs)(nil)).Elem()
}

// Response for the ListAvailableContacts API service call.
type ListSpacecraftAvailableContactsResultOutput struct{ *pulumi.OutputState }

func (ListSpacecraftAvailableContactsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSpacecraftAvailableContactsResult)(nil)).Elem()
}

func (o ListSpacecraftAvailableContactsResultOutput) ToListSpacecraftAvailableContactsResultOutput() ListSpacecraftAvailableContactsResultOutput {
	return o
}

func (o ListSpacecraftAvailableContactsResultOutput) ToListSpacecraftAvailableContactsResultOutputWithContext(ctx context.Context) ListSpacecraftAvailableContactsResultOutput {
	return o
}

// The URL to get the next set of results.
func (o ListSpacecraftAvailableContactsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListSpacecraftAvailableContactsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// A list of available contacts.
func (o ListSpacecraftAvailableContactsResultOutput) Value() AvailableContactsResponseArrayOutput {
	return o.ApplyT(func(v ListSpacecraftAvailableContactsResult) []AvailableContactsResponse { return v.Value }).(AvailableContactsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSpacecraftAvailableContactsResultOutput{})
}
