// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information related to a specific migrate project. Returns a json object of type 'migrateProject' as specified in the models section.
func LookupMigrateProjectsControllerMigrateProject(ctx *pulumi.Context, args *LookupMigrateProjectsControllerMigrateProjectArgs, opts ...pulumi.InvokeOption) (*LookupMigrateProjectsControllerMigrateProjectResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMigrateProjectsControllerMigrateProjectResult
	err := ctx.Invoke("azure-native:migrate/v20230101:getMigrateProjectsControllerMigrateProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrateProjectsControllerMigrateProjectArgs struct {
	// Name of the Azure Migrate project.
	MigrateProjectName string `pulumi:"migrateProjectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Migrate project.
type LookupMigrateProjectsControllerMigrateProjectResult struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Path reference to this project /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{projectName}
	Id string `pulumi:"id"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Name of the project.
	Name string `pulumi:"name"`
	// Properties of a migrate project.
	Properties MigrateProjectPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the object = [Microsoft.Migrate/migrateProjects].
	Type string `pulumi:"type"`
}

func LookupMigrateProjectsControllerMigrateProjectOutput(ctx *pulumi.Context, args LookupMigrateProjectsControllerMigrateProjectOutputArgs, opts ...pulumi.InvokeOption) LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMigrateProjectsControllerMigrateProjectResult, error) {
			args := v.(LookupMigrateProjectsControllerMigrateProjectArgs)
			r, err := LookupMigrateProjectsControllerMigrateProject(ctx, &args, opts...)
			var s LookupMigrateProjectsControllerMigrateProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMigrateProjectsControllerMigrateProjectResultOutput)
}

type LookupMigrateProjectsControllerMigrateProjectOutputArgs struct {
	// Name of the Azure Migrate project.
	MigrateProjectName pulumi.StringInput `pulumi:"migrateProjectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMigrateProjectsControllerMigrateProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectsControllerMigrateProjectArgs)(nil)).Elem()
}

// Migrate project.
type LookupMigrateProjectsControllerMigrateProjectResultOutput struct{ *pulumi.OutputState }

func (LookupMigrateProjectsControllerMigrateProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectsControllerMigrateProjectResult)(nil)).Elem()
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ToLookupMigrateProjectsControllerMigrateProjectResultOutput() LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return o
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ToLookupMigrateProjectsControllerMigrateProjectResultOutputWithContext(ctx context.Context) LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return o
}

// For optimistic concurrency control.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Path reference to this project /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{projectName}
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure location in which project is created.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the project.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of a migrate project.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Properties() MigrateProjectPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) MigrateProjectPropertiesResponse {
		return v.Properties
	}).(MigrateProjectPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the object = [Microsoft.Migrate/migrateProjects].
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrateProjectsControllerMigrateProjectResultOutput{})
}
