// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a database advisor.
func LookupDatabaseAdvisor(ctx *pulumi.Context, args *LookupDatabaseAdvisorArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAdvisorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseAdvisorResult
	err := ctx.Invoke("azure-native:sql/v20230501preview:getDatabaseAdvisor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAdvisorArgs struct {
	// The name of the Database Advisor.
	AdvisorName string `pulumi:"advisorName"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Database, Server or Elastic Pool Advisor.
type LookupDatabaseAdvisorResult struct {
	// Gets the status of availability of this advisor to customers. Possible values are 'GA', 'PublicPreview', 'LimitedPublicPreview' and 'PrivatePreview'.
	AdvisorStatus string `pulumi:"advisorStatus"`
	// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
	AutoExecuteStatus string `pulumi:"autoExecuteStatus"`
	// Gets the resource from which current value of auto-execute status is inherited. Auto-execute status can be set on (and inherited from) different levels in the resource hierarchy. Possible values are 'Subscription', 'Server', 'ElasticPool', 'Database' and 'Default' (when status is not explicitly set on any level).
	AutoExecuteStatusInheritedFrom string `pulumi:"autoExecuteStatusInheritedFrom"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource kind.
	Kind string `pulumi:"kind"`
	// Gets the time when the current resource was analyzed for recommendations by this advisor.
	LastChecked string `pulumi:"lastChecked"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Gets that status of recommendations for this advisor and reason for not having any recommendations. Possible values include, but are not limited to, 'Ok' (Recommendations available),LowActivity (not enough workload to analyze), 'DbSeemsTuned' (Database is doing well), etc.
	RecommendationsStatus string `pulumi:"recommendationsStatus"`
	// Gets the recommended actions for this advisor.
	RecommendedActions []RecommendedActionResponse `pulumi:"recommendedActions"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupDatabaseAdvisorOutput(ctx *pulumi.Context, args LookupDatabaseAdvisorOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAdvisorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAdvisorResult, error) {
			args := v.(LookupDatabaseAdvisorArgs)
			r, err := LookupDatabaseAdvisor(ctx, &args, opts...)
			var s LookupDatabaseAdvisorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAdvisorResultOutput)
}

type LookupDatabaseAdvisorOutputArgs struct {
	// The name of the Database Advisor.
	AdvisorName pulumi.StringInput `pulumi:"advisorName"`
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseAdvisorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAdvisorArgs)(nil)).Elem()
}

// Database, Server or Elastic Pool Advisor.
type LookupDatabaseAdvisorResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAdvisorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAdvisorResult)(nil)).Elem()
}

func (o LookupDatabaseAdvisorResultOutput) ToLookupDatabaseAdvisorResultOutput() LookupDatabaseAdvisorResultOutput {
	return o
}

func (o LookupDatabaseAdvisorResultOutput) ToLookupDatabaseAdvisorResultOutputWithContext(ctx context.Context) LookupDatabaseAdvisorResultOutput {
	return o
}

// Gets the status of availability of this advisor to customers. Possible values are 'GA', 'PublicPreview', 'LimitedPublicPreview' and 'PrivatePreview'.
func (o LookupDatabaseAdvisorResultOutput) AdvisorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.AdvisorStatus }).(pulumi.StringOutput)
}

// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
func (o LookupDatabaseAdvisorResultOutput) AutoExecuteStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.AutoExecuteStatus }).(pulumi.StringOutput)
}

// Gets the resource from which current value of auto-execute status is inherited. Auto-execute status can be set on (and inherited from) different levels in the resource hierarchy. Possible values are 'Subscription', 'Server', 'ElasticPool', 'Database' and 'Default' (when status is not explicitly set on any level).
func (o LookupDatabaseAdvisorResultOutput) AutoExecuteStatusInheritedFrom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.AutoExecuteStatusInheritedFrom }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupDatabaseAdvisorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource kind.
func (o LookupDatabaseAdvisorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Gets the time when the current resource was analyzed for recommendations by this advisor.
func (o LookupDatabaseAdvisorResultOutput) LastChecked() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.LastChecked }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupDatabaseAdvisorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupDatabaseAdvisorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets that status of recommendations for this advisor and reason for not having any recommendations. Possible values include, but are not limited to, 'Ok' (Recommendations available),LowActivity (not enough workload to analyze), 'DbSeemsTuned' (Database is doing well), etc.
func (o LookupDatabaseAdvisorResultOutput) RecommendationsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.RecommendationsStatus }).(pulumi.StringOutput)
}

// Gets the recommended actions for this advisor.
func (o LookupDatabaseAdvisorResultOutput) RecommendedActions() RecommendedActionResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) []RecommendedActionResponse { return v.RecommendedActions }).(RecommendedActionResponseArrayOutput)
}

// Resource type.
func (o LookupDatabaseAdvisorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAdvisorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAdvisorResultOutput{})
}
