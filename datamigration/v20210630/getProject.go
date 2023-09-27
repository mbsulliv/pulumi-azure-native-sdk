// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210630

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The project resource is a nested resource representing a stored migration project. The GET method retrieves information about a project.
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:datamigration/v20210630:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Name of the project
	ProjectName string `pulumi:"projectName"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
}

// A project resource
type LookupProjectResult struct {
	// UTC Date and time when project was created
	CreationTime string `pulumi:"creationTime"`
	// List of DatabaseInfo
	DatabasesInfo []DatabaseInfoResponse `pulumi:"databasesInfo"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The project's provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// Information for connecting to source
	SourceConnectionInfo interface{} `pulumi:"sourceConnectionInfo"`
	// Source platform for the project
	SourcePlatform string `pulumi:"sourcePlatform"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Information for connecting to target
	TargetConnectionInfo interface{} `pulumi:"targetConnectionInfo"`
	// Target platform for the project
	TargetPlatform string `pulumi:"targetPlatform"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	// Name of the resource group
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// Name of the project
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// Name of the service
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A project resource
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupProjectResult] {
	return pulumix.Output[LookupProjectResult]{
		OutputState: o.OutputState,
	}
}

// UTC Date and time when project was created
func (o LookupProjectResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// List of DatabaseInfo
func (o LookupProjectResultOutput) DatabasesInfo() DatabaseInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []DatabaseInfoResponse { return v.DatabasesInfo }).(DatabaseInfoResponseArrayOutput)
}

// Resource ID.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupProjectResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The project's provisioning state
func (o LookupProjectResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Information for connecting to source
func (o LookupProjectResultOutput) SourceConnectionInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProjectResult) interface{} { return v.SourceConnectionInfo }).(pulumi.AnyOutput)
}

// Source platform for the project
func (o LookupProjectResultOutput) SourcePlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SourcePlatform }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupProjectResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProjectResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupProjectResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProjectResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Information for connecting to target
func (o LookupProjectResultOutput) TargetConnectionInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProjectResult) interface{} { return v.TargetConnectionInfo }).(pulumi.AnyOutput)
}

// Target platform for the project
func (o LookupProjectResultOutput) TargetPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.TargetPlatform }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
