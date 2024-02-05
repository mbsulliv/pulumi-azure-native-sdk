// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a NrfDeploymentResource
func LookupNrfDeployment(ctx *pulumi.Context, args *LookupNrfDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupNrfDeploymentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNrfDeploymentResult
	err := ctx.Invoke("azure-native:mobilepacketcore/v20231015preview:getNrfDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNrfDeploymentArgs struct {
	// The name of the NrfDeployment
	NrfDeploymentName string `pulumi:"nrfDeploymentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure for Operators 5G Core Network Repository Function (NRF) Deployment Resource
type LookupNrfDeploymentResult struct {
	// Reference to cluster where the Network Function is deployed
	ClusterService string `pulumi:"clusterService"`
	// Azure for Operators 5G Core NRF component parameters
	ComponentParameters *string `pulumi:"componentParameters"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Operational status
	OperationalStatus OperationalStatusResponse `pulumi:"operationalStatus"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Release version. This is inherited from the cluster
	ReleaseVersion string `pulumi:"releaseVersion"`
	// Azure for Operators 5G Core NRF secrets parameters
	SecretsParameters *string `pulumi:"secretsParameters"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNrfDeploymentOutput(ctx *pulumi.Context, args LookupNrfDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupNrfDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNrfDeploymentResult, error) {
			args := v.(LookupNrfDeploymentArgs)
			r, err := LookupNrfDeployment(ctx, &args, opts...)
			var s LookupNrfDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNrfDeploymentResultOutput)
}

type LookupNrfDeploymentOutputArgs struct {
	// The name of the NrfDeployment
	NrfDeploymentName pulumi.StringInput `pulumi:"nrfDeploymentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNrfDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNrfDeploymentArgs)(nil)).Elem()
}

// Azure for Operators 5G Core Network Repository Function (NRF) Deployment Resource
type LookupNrfDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupNrfDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNrfDeploymentResult)(nil)).Elem()
}

func (o LookupNrfDeploymentResultOutput) ToLookupNrfDeploymentResultOutput() LookupNrfDeploymentResultOutput {
	return o
}

func (o LookupNrfDeploymentResultOutput) ToLookupNrfDeploymentResultOutputWithContext(ctx context.Context) LookupNrfDeploymentResultOutput {
	return o
}

// Reference to cluster where the Network Function is deployed
func (o LookupNrfDeploymentResultOutput) ClusterService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) string { return v.ClusterService }).(pulumi.StringOutput)
}

// Azure for Operators 5G Core NRF component parameters
func (o LookupNrfDeploymentResultOutput) ComponentParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) *string { return v.ComponentParameters }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNrfDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNrfDeploymentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNrfDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operational status
func (o LookupNrfDeploymentResultOutput) OperationalStatus() OperationalStatusResponseOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) OperationalStatusResponse { return v.OperationalStatus }).(OperationalStatusResponseOutput)
}

// The status of the last operation.
func (o LookupNrfDeploymentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Release version. This is inherited from the cluster
func (o LookupNrfDeploymentResultOutput) ReleaseVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) string { return v.ReleaseVersion }).(pulumi.StringOutput)
}

// Azure for Operators 5G Core NRF secrets parameters
func (o LookupNrfDeploymentResultOutput) SecretsParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) *string { return v.SecretsParameters }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNrfDeploymentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNrfDeploymentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNrfDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrfDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNrfDeploymentResultOutput{})
}
