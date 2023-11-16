// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the recommended SAP Availability Zone Pair Details for your region.
func GetSAPAvailabilityZoneDetails(ctx *pulumi.Context, args *GetSAPAvailabilityZoneDetailsArgs, opts ...pulumi.InvokeOption) (*GetSAPAvailabilityZoneDetailsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetSAPAvailabilityZoneDetailsResult
	err := ctx.Invoke("azure-native:workloads/v20221101preview:getSAPAvailabilityZoneDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSAPAvailabilityZoneDetailsArgs struct {
	// The geo-location where the SAP resources will be created.
	AppLocation string `pulumi:"appLocation"`
	// The database type. Eg: HANA, DB2, etc
	DatabaseType string `pulumi:"databaseType"`
	// The name of Azure region.
	Location string `pulumi:"location"`
	// Defines the SAP Product type.
	SapProduct string `pulumi:"sapProduct"`
}

// The list of supported availability zone pairs which are part of SAP HA deployment.
type GetSAPAvailabilityZoneDetailsResult struct {
	// Gets the list of availability zone pairs.
	AvailabilityZonePairs []SAPAvailabilityZonePairResponse `pulumi:"availabilityZonePairs"`
}

func GetSAPAvailabilityZoneDetailsOutput(ctx *pulumi.Context, args GetSAPAvailabilityZoneDetailsOutputArgs, opts ...pulumi.InvokeOption) GetSAPAvailabilityZoneDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAPAvailabilityZoneDetailsResult, error) {
			args := v.(GetSAPAvailabilityZoneDetailsArgs)
			r, err := GetSAPAvailabilityZoneDetails(ctx, &args, opts...)
			var s GetSAPAvailabilityZoneDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAPAvailabilityZoneDetailsResultOutput)
}

type GetSAPAvailabilityZoneDetailsOutputArgs struct {
	// The geo-location where the SAP resources will be created.
	AppLocation pulumi.StringInput `pulumi:"appLocation"`
	// The database type. Eg: HANA, DB2, etc
	DatabaseType pulumi.StringInput `pulumi:"databaseType"`
	// The name of Azure region.
	Location pulumi.StringInput `pulumi:"location"`
	// Defines the SAP Product type.
	SapProduct pulumi.StringInput `pulumi:"sapProduct"`
}

func (GetSAPAvailabilityZoneDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPAvailabilityZoneDetailsArgs)(nil)).Elem()
}

// The list of supported availability zone pairs which are part of SAP HA deployment.
type GetSAPAvailabilityZoneDetailsResultOutput struct{ *pulumi.OutputState }

func (GetSAPAvailabilityZoneDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPAvailabilityZoneDetailsResult)(nil)).Elem()
}

func (o GetSAPAvailabilityZoneDetailsResultOutput) ToGetSAPAvailabilityZoneDetailsResultOutput() GetSAPAvailabilityZoneDetailsResultOutput {
	return o
}

func (o GetSAPAvailabilityZoneDetailsResultOutput) ToGetSAPAvailabilityZoneDetailsResultOutputWithContext(ctx context.Context) GetSAPAvailabilityZoneDetailsResultOutput {
	return o
}

// Gets the list of availability zone pairs.
func (o GetSAPAvailabilityZoneDetailsResultOutput) AvailabilityZonePairs() SAPAvailabilityZonePairResponseArrayOutput {
	return o.ApplyT(func(v GetSAPAvailabilityZoneDetailsResult) []SAPAvailabilityZonePairResponse {
		return v.AvailabilityZonePairs
	}).(SAPAvailabilityZonePairResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAPAvailabilityZoneDetailsResultOutput{})
}
