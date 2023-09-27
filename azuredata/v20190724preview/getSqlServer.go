// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190724preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a SQL Server.
func LookupSqlServer(ctx *pulumi.Context, args *LookupSqlServerArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlServerResult
	err := ctx.Invoke("azure-native:azuredata/v20190724preview:getSqlServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerArgs struct {
	// The child resources to include in the response.
	Expand *string `pulumi:"expand"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL Server.
	SqlServerName string `pulumi:"sqlServerName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName string `pulumi:"sqlServerRegistrationName"`
}

// A SQL server.
type LookupSqlServerResult struct {
	// Cores of the Sql Server.
	Cores *int `pulumi:"cores"`
	// Sql Server Edition.
	Edition *string `pulumi:"edition"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Sql Server Json Property Bag.
	PropertyBag *string `pulumi:"propertyBag"`
	// ID for Parent Sql Server Registration.
	RegistrationID *string `pulumi:"registrationID"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// Version of the Sql Server.
	Version *string `pulumi:"version"`
}

func LookupSqlServerOutput(ctx *pulumi.Context, args LookupSqlServerOutputArgs, opts ...pulumi.InvokeOption) LookupSqlServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlServerResult, error) {
			args := v.(LookupSqlServerArgs)
			r, err := LookupSqlServer(ctx, &args, opts...)
			var s LookupSqlServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlServerResultOutput)
}

type LookupSqlServerOutputArgs struct {
	// The child resources to include in the response.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the SQL Server.
	SqlServerName pulumi.StringInput `pulumi:"sqlServerName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName pulumi.StringInput `pulumi:"sqlServerRegistrationName"`
}

func (LookupSqlServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerArgs)(nil)).Elem()
}

// A SQL server.
type LookupSqlServerResultOutput struct{ *pulumi.OutputState }

func (LookupSqlServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerResult)(nil)).Elem()
}

func (o LookupSqlServerResultOutput) ToLookupSqlServerResultOutput() LookupSqlServerResultOutput {
	return o
}

func (o LookupSqlServerResultOutput) ToLookupSqlServerResultOutputWithContext(ctx context.Context) LookupSqlServerResultOutput {
	return o
}

func (o LookupSqlServerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlServerResult] {
	return pulumix.Output[LookupSqlServerResult]{
		OutputState: o.OutputState,
	}
}

// Cores of the Sql Server.
func (o LookupSqlServerResultOutput) Cores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *int { return v.Cores }).(pulumi.IntPtrOutput)
}

// Sql Server Edition.
func (o LookupSqlServerResultOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Sql Server Json Property Bag.
func (o LookupSqlServerResultOutput) PropertyBag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *string { return v.PropertyBag }).(pulumi.StringPtrOutput)
}

// ID for Parent Sql Server Registration.
func (o LookupSqlServerResultOutput) RegistrationID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *string { return v.RegistrationID }).(pulumi.StringPtrOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupSqlServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version of the Sql Server.
func (o LookupSqlServerResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlServerResultOutput{})
}
