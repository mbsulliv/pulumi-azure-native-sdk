// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the SAP Disk Configuration Layout prod/non-prod SAP System.
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2021-12-01-preview, 2022-11-01-preview, 2023-10-01-preview.
func GetSAPDiskConfigurations(ctx *pulumi.Context, args *GetSAPDiskConfigurationsArgs, opts ...pulumi.InvokeOption) (*GetSAPDiskConfigurationsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetSAPDiskConfigurationsResult
	err := ctx.Invoke("azure-native:workloads:getSAPDiskConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSAPDiskConfigurationsArgs struct {
	// The geo-location where the SAP resources will be created.
	AppLocation string `pulumi:"appLocation"`
	// The database type. Eg: HANA, DB2, etc
	DatabaseType string `pulumi:"databaseType"`
	// The VM SKU for database instance.
	DbVmSku string `pulumi:"dbVmSku"`
	// The deployment type. Eg: SingleServer/ThreeTier
	DeploymentType string `pulumi:"deploymentType"`
	// Defines the environment type - Production/Non Production.
	Environment string `pulumi:"environment"`
	// The name of Azure region.
	Location string `pulumi:"location"`
	// Defines the SAP Product type.
	SapProduct string `pulumi:"sapProduct"`
}

// The list of disk configuration for vmSku which are part of SAP deployment.
type GetSAPDiskConfigurationsResult struct {
	// The disk configuration for the db volume. For HANA, Required volumes are: ['hana/data', 'hana/log', hana/shared', 'usr/sap', 'os'], Optional volume : ['backup'].
	VolumeConfigurations map[string]SAPDiskConfigurationResponse `pulumi:"volumeConfigurations"`
}

func GetSAPDiskConfigurationsOutput(ctx *pulumi.Context, args GetSAPDiskConfigurationsOutputArgs, opts ...pulumi.InvokeOption) GetSAPDiskConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAPDiskConfigurationsResult, error) {
			args := v.(GetSAPDiskConfigurationsArgs)
			r, err := GetSAPDiskConfigurations(ctx, &args, opts...)
			var s GetSAPDiskConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAPDiskConfigurationsResultOutput)
}

type GetSAPDiskConfigurationsOutputArgs struct {
	// The geo-location where the SAP resources will be created.
	AppLocation pulumi.StringInput `pulumi:"appLocation"`
	// The database type. Eg: HANA, DB2, etc
	DatabaseType pulumi.StringInput `pulumi:"databaseType"`
	// The VM SKU for database instance.
	DbVmSku pulumi.StringInput `pulumi:"dbVmSku"`
	// The deployment type. Eg: SingleServer/ThreeTier
	DeploymentType pulumi.StringInput `pulumi:"deploymentType"`
	// Defines the environment type - Production/Non Production.
	Environment pulumi.StringInput `pulumi:"environment"`
	// The name of Azure region.
	Location pulumi.StringInput `pulumi:"location"`
	// Defines the SAP Product type.
	SapProduct pulumi.StringInput `pulumi:"sapProduct"`
}

func (GetSAPDiskConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPDiskConfigurationsArgs)(nil)).Elem()
}

// The list of disk configuration for vmSku which are part of SAP deployment.
type GetSAPDiskConfigurationsResultOutput struct{ *pulumi.OutputState }

func (GetSAPDiskConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPDiskConfigurationsResult)(nil)).Elem()
}

func (o GetSAPDiskConfigurationsResultOutput) ToGetSAPDiskConfigurationsResultOutput() GetSAPDiskConfigurationsResultOutput {
	return o
}

func (o GetSAPDiskConfigurationsResultOutput) ToGetSAPDiskConfigurationsResultOutputWithContext(ctx context.Context) GetSAPDiskConfigurationsResultOutput {
	return o
}

// The disk configuration for the db volume. For HANA, Required volumes are: ['hana/data', 'hana/log', hana/shared', 'usr/sap', 'os'], Optional volume : ['backup'].
func (o GetSAPDiskConfigurationsResultOutput) VolumeConfigurations() SAPDiskConfigurationResponseMapOutput {
	return o.ApplyT(func(v GetSAPDiskConfigurationsResult) map[string]SAPDiskConfigurationResponse {
		return v.VolumeConfigurations
	}).(SAPDiskConfigurationResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAPDiskConfigurationsResultOutput{})
}
