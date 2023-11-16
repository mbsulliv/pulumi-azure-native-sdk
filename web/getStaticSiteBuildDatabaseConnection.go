// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Static Site Database Connection resource.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2023-01-01.
func LookupStaticSiteBuildDatabaseConnection(ctx *pulumi.Context, args *LookupStaticSiteBuildDatabaseConnectionArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteBuildDatabaseConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticSiteBuildDatabaseConnectionResult
	err := ctx.Invoke("azure-native:web:getStaticSiteBuildDatabaseConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteBuildDatabaseConnectionArgs struct {
	// Name of the database connection.
	DatabaseConnectionName string `pulumi:"databaseConnectionName"`
	// The stage site identifier.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of the static site
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Static Site Database Connection resource.
type LookupStaticSiteBuildDatabaseConnectionResult struct {
	// A list of configuration files associated with this database connection.
	ConfigurationFiles []StaticSiteDatabaseConnectionConfigurationFileOverviewResponse `pulumi:"configurationFiles"`
	// If present, the identity is used in conjunction with connection string to connect to the database. Use of the system-assigned managed identity is indicated with the string 'SystemAssigned', while use of a user-assigned managed identity is indicated with the resource id of the managed identity resource.
	ConnectionIdentity *string `pulumi:"connectionIdentity"`
	// The connection string to use to connect to the database.
	ConnectionString *string `pulumi:"connectionString"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// The region of the database resource.
	Region string `pulumi:"region"`
	// The resource id of the database.
	ResourceId string `pulumi:"resourceId"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupStaticSiteBuildDatabaseConnectionOutput(ctx *pulumi.Context, args LookupStaticSiteBuildDatabaseConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteBuildDatabaseConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteBuildDatabaseConnectionResult, error) {
			args := v.(LookupStaticSiteBuildDatabaseConnectionArgs)
			r, err := LookupStaticSiteBuildDatabaseConnection(ctx, &args, opts...)
			var s LookupStaticSiteBuildDatabaseConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteBuildDatabaseConnectionResultOutput)
}

type LookupStaticSiteBuildDatabaseConnectionOutputArgs struct {
	// Name of the database connection.
	DatabaseConnectionName pulumi.StringInput `pulumi:"databaseConnectionName"`
	// The stage site identifier.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// Name of the static site
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteBuildDatabaseConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteBuildDatabaseConnectionArgs)(nil)).Elem()
}

// Static Site Database Connection resource.
type LookupStaticSiteBuildDatabaseConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteBuildDatabaseConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteBuildDatabaseConnectionResult)(nil)).Elem()
}

func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) ToLookupStaticSiteBuildDatabaseConnectionResultOutput() LookupStaticSiteBuildDatabaseConnectionResultOutput {
	return o
}

func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) ToLookupStaticSiteBuildDatabaseConnectionResultOutputWithContext(ctx context.Context) LookupStaticSiteBuildDatabaseConnectionResultOutput {
	return o
}

// A list of configuration files associated with this database connection.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) ConfigurationFiles() StaticSiteDatabaseConnectionConfigurationFileOverviewResponseArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) []StaticSiteDatabaseConnectionConfigurationFileOverviewResponse {
		return v.ConfigurationFiles
	}).(StaticSiteDatabaseConnectionConfigurationFileOverviewResponseArrayOutput)
}

// If present, the identity is used in conjunction with connection string to connect to the database. Use of the system-assigned managed identity is indicated with the string 'SystemAssigned', while use of a user-assigned managed identity is indicated with the resource id of the managed identity resource.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) ConnectionIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) *string { return v.ConnectionIdentity }).(pulumi.StringPtrOutput)
}

// The connection string to use to connect to the database.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The region of the database resource.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) string { return v.Region }).(pulumi.StringOutput)
}

// The resource id of the database.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupStaticSiteBuildDatabaseConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteBuildDatabaseConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteBuildDatabaseConnectionResultOutput{})
}
