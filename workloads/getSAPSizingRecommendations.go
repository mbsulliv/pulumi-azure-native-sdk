// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get SAP sizing recommendations by providing input SAPS for application tier and memory required for database tier
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2021-12-01-preview, 2022-11-01-preview, 2023-10-01-preview.
func GetSAPSizingRecommendations(ctx *pulumi.Context, args *GetSAPSizingRecommendationsArgs, opts ...pulumi.InvokeOption) (*GetSAPSizingRecommendationsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetSAPSizingRecommendationsResult
	err := ctx.Invoke("azure-native:workloads:getSAPSizingRecommendations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSAPSizingRecommendationsArgs struct {
	// The geo-location where the resource is to be created.
	AppLocation string `pulumi:"appLocation"`
	// The database type.
	DatabaseType string `pulumi:"databaseType"`
	// The database memory configuration.
	DbMemory float64 `pulumi:"dbMemory"`
	// The DB scale method.
	DbScaleMethod *string `pulumi:"dbScaleMethod"`
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
	// The SAP Application Performance Standard measurement.
	Saps float64 `pulumi:"saps"`
}

// The SAP sizing recommendation result.
type GetSAPSizingRecommendationsResult struct {
	// The type of SAP deployment, single server or Three tier.
	DeploymentType string `pulumi:"deploymentType"`
}

func GetSAPSizingRecommendationsOutput(ctx *pulumi.Context, args GetSAPSizingRecommendationsOutputArgs, opts ...pulumi.InvokeOption) GetSAPSizingRecommendationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAPSizingRecommendationsResult, error) {
			args := v.(GetSAPSizingRecommendationsArgs)
			r, err := GetSAPSizingRecommendations(ctx, &args, opts...)
			var s GetSAPSizingRecommendationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAPSizingRecommendationsResultOutput)
}

type GetSAPSizingRecommendationsOutputArgs struct {
	// The geo-location where the resource is to be created.
	AppLocation pulumi.StringInput `pulumi:"appLocation"`
	// The database type.
	DatabaseType pulumi.StringInput `pulumi:"databaseType"`
	// The database memory configuration.
	DbMemory pulumi.Float64Input `pulumi:"dbMemory"`
	// The DB scale method.
	DbScaleMethod pulumi.StringPtrInput `pulumi:"dbScaleMethod"`
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
	// The SAP Application Performance Standard measurement.
	Saps pulumi.Float64Input `pulumi:"saps"`
}

func (GetSAPSizingRecommendationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPSizingRecommendationsArgs)(nil)).Elem()
}

// The SAP sizing recommendation result.
type GetSAPSizingRecommendationsResultOutput struct{ *pulumi.OutputState }

func (GetSAPSizingRecommendationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPSizingRecommendationsResult)(nil)).Elem()
}

func (o GetSAPSizingRecommendationsResultOutput) ToGetSAPSizingRecommendationsResultOutput() GetSAPSizingRecommendationsResultOutput {
	return o
}

func (o GetSAPSizingRecommendationsResultOutput) ToGetSAPSizingRecommendationsResultOutputWithContext(ctx context.Context) GetSAPSizingRecommendationsResultOutput {
	return o
}

func (o GetSAPSizingRecommendationsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSAPSizingRecommendationsResult] {
	return pulumix.Output[GetSAPSizingRecommendationsResult]{
		OutputState: o.OutputState,
	}
}

// The type of SAP deployment, single server or Three tier.
func (o GetSAPSizingRecommendationsResultOutput) DeploymentType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAPSizingRecommendationsResult) string { return v.DeploymentType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAPSizingRecommendationsResultOutput{})
}
