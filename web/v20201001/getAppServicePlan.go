// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an App Service plan.
func LookupAppServicePlan(ctx *pulumi.Context, args *LookupAppServicePlanArgs, opts ...pulumi.InvokeOption) (*LookupAppServicePlanResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAppServicePlanResult
	err := ctx.Invoke("azure-native:web/v20201001:getAppServicePlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppServicePlanArgs struct {
	// Name of the App Service plan.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// App Service plan.
type LookupAppServicePlanResult struct {
	// The time when the server farm free offer expires.
	FreeOfferExpirationTime *string `pulumi:"freeOfferExpirationTime"`
	// Geographical location for the App Service plan.
	GeoRegion string `pulumi:"geoRegion"`
	// Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	// If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	HyperV *bool `pulumi:"hyperV"`
	// Resource Id.
	Id string `pulumi:"id"`
	// If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot *bool `pulumi:"isSpot"`
	// Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	IsXenon *bool `pulumi:"isXenon"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
	MaximumElasticWorkerCount *int `pulumi:"maximumElasticWorkerCount"`
	// Maximum number of instances that can be assigned to this App Service plan.
	MaximumNumberOfWorkers int `pulumi:"maximumNumberOfWorkers"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Number of apps assigned to this App Service plan.
	NumberOfSites int `pulumi:"numberOfSites"`
	// If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
	// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
	PerSiteScaling *bool `pulumi:"perSiteScaling"`
	// Provisioning state of the App Service Environment.
	ProvisioningState string `pulumi:"provisioningState"`
	// If Linux app service plan <code>true</code>, <code>false</code> otherwise.
	Reserved *bool `pulumi:"reserved"`
	// Resource group of the App Service plan.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescriptionResponse `pulumi:"sku"`
	// The time when the server farm expires. Valid only if it is a spot server farm.
	SpotExpirationTime *string `pulumi:"spotExpirationTime"`
	// App Service plan status.
	Status string `pulumi:"status"`
	// App Service plan subscription.
	Subscription string `pulumi:"subscription"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Scaling worker count.
	TargetWorkerCount *int `pulumi:"targetWorkerCount"`
	// Scaling worker size ID.
	TargetWorkerSizeId *int `pulumi:"targetWorkerSizeId"`
	// Resource type.
	Type string `pulumi:"type"`
	// Target worker tier assigned to the App Service plan.
	WorkerTierName *string `pulumi:"workerTierName"`
}

// Defaults sets the appropriate defaults for LookupAppServicePlanResult
func (val *LookupAppServicePlanResult) Defaults() *LookupAppServicePlanResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.HyperV == nil {
		hyperV_ := false
		tmp.HyperV = &hyperV_
	}
	if tmp.IsXenon == nil {
		isXenon_ := false
		tmp.IsXenon = &isXenon_
	}
	if tmp.PerSiteScaling == nil {
		perSiteScaling_ := false
		tmp.PerSiteScaling = &perSiteScaling_
	}
	if tmp.Reserved == nil {
		reserved_ := false
		tmp.Reserved = &reserved_
	}
	return &tmp
}

func LookupAppServicePlanOutput(ctx *pulumi.Context, args LookupAppServicePlanOutputArgs, opts ...pulumi.InvokeOption) LookupAppServicePlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServicePlanResult, error) {
			args := v.(LookupAppServicePlanArgs)
			r, err := LookupAppServicePlan(ctx, &args, opts...)
			var s LookupAppServicePlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServicePlanResultOutput)
}

type LookupAppServicePlanOutputArgs struct {
	// Name of the App Service plan.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServicePlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServicePlanArgs)(nil)).Elem()
}

// App Service plan.
type LookupAppServicePlanResultOutput struct{ *pulumi.OutputState }

func (LookupAppServicePlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServicePlanResult)(nil)).Elem()
}

func (o LookupAppServicePlanResultOutput) ToLookupAppServicePlanResultOutput() LookupAppServicePlanResultOutput {
	return o
}

func (o LookupAppServicePlanResultOutput) ToLookupAppServicePlanResultOutputWithContext(ctx context.Context) LookupAppServicePlanResultOutput {
	return o
}

// The time when the server farm free offer expires.
func (o LookupAppServicePlanResultOutput) FreeOfferExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *string { return v.FreeOfferExpirationTime }).(pulumi.StringPtrOutput)
}

// Geographical location for the App Service plan.
func (o LookupAppServicePlanResultOutput) GeoRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.GeoRegion }).(pulumi.StringOutput)
}

// Specification for the App Service Environment to use for the App Service plan.
func (o LookupAppServicePlanResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *HostingEnvironmentProfileResponse {
		return v.HostingEnvironmentProfile
	}).(HostingEnvironmentProfileResponsePtrOutput)
}

// If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
func (o LookupAppServicePlanResultOutput) HyperV() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.HyperV }).(pulumi.BoolPtrOutput)
}

// Resource Id.
func (o LookupAppServicePlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Id }).(pulumi.StringOutput)
}

// If <code>true</code>, this App Service Plan owns spot instances.
func (o LookupAppServicePlanResultOutput) IsSpot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.IsSpot }).(pulumi.BoolPtrOutput)
}

// Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
func (o LookupAppServicePlanResultOutput) IsXenon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.IsXenon }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o LookupAppServicePlanResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location.
func (o LookupAppServicePlanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
func (o LookupAppServicePlanResultOutput) MaximumElasticWorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *int { return v.MaximumElasticWorkerCount }).(pulumi.IntPtrOutput)
}

// Maximum number of instances that can be assigned to this App Service plan.
func (o LookupAppServicePlanResultOutput) MaximumNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) int { return v.MaximumNumberOfWorkers }).(pulumi.IntOutput)
}

// Resource Name.
func (o LookupAppServicePlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of apps assigned to this App Service plan.
func (o LookupAppServicePlanResultOutput) NumberOfSites() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) int { return v.NumberOfSites }).(pulumi.IntOutput)
}

// If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
func (o LookupAppServicePlanResultOutput) PerSiteScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.PerSiteScaling }).(pulumi.BoolPtrOutput)
}

// Provisioning state of the App Service Environment.
func (o LookupAppServicePlanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// If Linux app service plan <code>true</code>, <code>false</code> otherwise.
func (o LookupAppServicePlanResultOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.Reserved }).(pulumi.BoolPtrOutput)
}

// Resource group of the App Service plan.
func (o LookupAppServicePlanResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Description of a SKU for a scalable resource.
func (o LookupAppServicePlanResultOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *SkuDescriptionResponse { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

// The time when the server farm expires. Valid only if it is a spot server farm.
func (o LookupAppServicePlanResultOutput) SpotExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *string { return v.SpotExpirationTime }).(pulumi.StringPtrOutput)
}

// App Service plan status.
func (o LookupAppServicePlanResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Status }).(pulumi.StringOutput)
}

// App Service plan subscription.
func (o LookupAppServicePlanResultOutput) Subscription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Subscription }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o LookupAppServicePlanResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAppServicePlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Scaling worker count.
func (o LookupAppServicePlanResultOutput) TargetWorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *int { return v.TargetWorkerCount }).(pulumi.IntPtrOutput)
}

// Scaling worker size ID.
func (o LookupAppServicePlanResultOutput) TargetWorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *int { return v.TargetWorkerSizeId }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o LookupAppServicePlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Type }).(pulumi.StringOutput)
}

// Target worker tier assigned to the App Service plan.
func (o LookupAppServicePlanResultOutput) WorkerTierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *string { return v.WorkerTierName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServicePlanResultOutput{})
}
