// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a list of SAP supported SKUs for ASCS, Application and Database tier.
func GetSAPSupportedSku(ctx *pulumi.Context, args *GetSAPSupportedSkuArgs, opts ...pulumi.InvokeOption) (*GetSAPSupportedSkuResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetSAPSupportedSkuResult
	err := ctx.Invoke("azure-native:workloads/v20211201preview:getSAPSupportedSku", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSAPSupportedSkuArgs struct {
	// The geo-location where the resource is to be created.
	AppLocation string `pulumi:"appLocation"`
	// The database type. Eg: HANA, DB2, etc
	DatabaseType string `pulumi:"databaseType"`
	// The deployment type. Eg: SingleServer/ThreeTier
	DeploymentType string `pulumi:"deploymentType"`
	// Defines the environment type - Production/Non Production.
	Environment string `pulumi:"environment"`
	// The high availability type.
	HighAvailabilityType *string `pulumi:"highAvailabilityType"`
	// The name of Azure region.
	Location string `pulumi:"location"`
	// Defines the SAP Product type.
	SapProduct string `pulumi:"sapProduct"`
}

// The list of supported SKUs for different resources which are part of SAP deployment.
type GetSAPSupportedSkuResult struct {
	// Gets the list of SAP supported SKUs.
	SupportedSkus []SAPSupportedSkuResponse `pulumi:"supportedSkus"`
}

func GetSAPSupportedSkuOutput(ctx *pulumi.Context, args GetSAPSupportedSkuOutputArgs, opts ...pulumi.InvokeOption) GetSAPSupportedSkuResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAPSupportedSkuResult, error) {
			args := v.(GetSAPSupportedSkuArgs)
			r, err := GetSAPSupportedSku(ctx, &args, opts...)
			var s GetSAPSupportedSkuResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAPSupportedSkuResultOutput)
}

type GetSAPSupportedSkuOutputArgs struct {
	// The geo-location where the resource is to be created.
	AppLocation pulumi.StringInput `pulumi:"appLocation"`
	// The database type. Eg: HANA, DB2, etc
	DatabaseType pulumi.StringInput `pulumi:"databaseType"`
	// The deployment type. Eg: SingleServer/ThreeTier
	DeploymentType pulumi.StringInput `pulumi:"deploymentType"`
	// Defines the environment type - Production/Non Production.
	Environment pulumi.StringInput `pulumi:"environment"`
	// The high availability type.
	HighAvailabilityType pulumi.StringPtrInput `pulumi:"highAvailabilityType"`
	// The name of Azure region.
	Location pulumi.StringInput `pulumi:"location"`
	// Defines the SAP Product type.
	SapProduct pulumi.StringInput `pulumi:"sapProduct"`
}

func (GetSAPSupportedSkuOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPSupportedSkuArgs)(nil)).Elem()
}

// The list of supported SKUs for different resources which are part of SAP deployment.
type GetSAPSupportedSkuResultOutput struct{ *pulumi.OutputState }

func (GetSAPSupportedSkuResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPSupportedSkuResult)(nil)).Elem()
}

func (o GetSAPSupportedSkuResultOutput) ToGetSAPSupportedSkuResultOutput() GetSAPSupportedSkuResultOutput {
	return o
}

func (o GetSAPSupportedSkuResultOutput) ToGetSAPSupportedSkuResultOutputWithContext(ctx context.Context) GetSAPSupportedSkuResultOutput {
	return o
}

func (o GetSAPSupportedSkuResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSAPSupportedSkuResult] {
	return pulumix.Output[GetSAPSupportedSkuResult]{
		OutputState: o.OutputState,
	}
}

// Gets the list of SAP supported SKUs.
func (o GetSAPSupportedSkuResultOutput) SupportedSkus() SAPSupportedSkuResponseArrayOutput {
	return o.ApplyT(func(v GetSAPSupportedSkuResult) []SAPSupportedSkuResponse { return v.SupportedSkus }).(SAPSupportedSkuResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAPSupportedSkuResultOutput{})
}
