// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Workload group operations for a data warehouse
type WorkloadGroup struct {
	pulumi.CustomResourceState

	// The workload group importance level.
	Importance pulumi.StringPtrOutput `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent pulumi.IntOutput `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest pulumi.Float64PtrOutput `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent pulumi.IntOutput `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest pulumi.Float64Output `pulumi:"minResourcePercentPerRequest"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The workload group query execution timeout.
	QueryExecutionTimeout pulumi.IntPtrOutput `pulumi:"queryExecutionTimeout"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadGroup registers a new resource with the given unique name, arguments, and options.
func NewWorkloadGroup(ctx *pulumi.Context,
	name string, args *WorkloadGroupArgs, opts ...pulumi.ResourceOption) (*WorkloadGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.MaxResourcePercent == nil {
		return nil, errors.New("invalid value for required argument 'MaxResourcePercent'")
	}
	if args.MinResourcePercent == nil {
		return nil, errors.New("invalid value for required argument 'MinResourcePercent'")
	}
	if args.MinResourcePercentPerRequest == nil {
		return nil, errors.New("invalid value for required argument 'MinResourcePercentPerRequest'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:WorkloadGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadGroup
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:WorkloadGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadGroup gets an existing WorkloadGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadGroupState, opts ...pulumi.ResourceOption) (*WorkloadGroup, error) {
	var resource WorkloadGroup
	err := ctx.ReadResource("azure-native:sql/v20230201preview:WorkloadGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadGroup resources.
type workloadGroupState struct {
}

type WorkloadGroupState struct {
}

func (WorkloadGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupState)(nil)).Elem()
}

type workloadGroupArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The workload group importance level.
	Importance *string `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent int `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent int `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest float64 `pulumi:"minResourcePercentPerRequest"`
	// The workload group query execution timeout.
	QueryExecutionTimeout *int `pulumi:"queryExecutionTimeout"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the workload group.
	WorkloadGroupName *string `pulumi:"workloadGroupName"`
}

// The set of arguments for constructing a WorkloadGroup resource.
type WorkloadGroupArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The workload group importance level.
	Importance pulumi.StringPtrInput
	// The workload group cap percentage resource.
	MaxResourcePercent pulumi.IntInput
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest pulumi.Float64PtrInput
	// The workload group minimum percentage resource.
	MinResourcePercent pulumi.IntInput
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest pulumi.Float64Input
	// The workload group query execution timeout.
	QueryExecutionTimeout pulumi.IntPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name of the workload group.
	WorkloadGroupName pulumi.StringPtrInput
}

func (WorkloadGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupArgs)(nil)).Elem()
}

type WorkloadGroupInput interface {
	pulumi.Input

	ToWorkloadGroupOutput() WorkloadGroupOutput
	ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput
}

func (*WorkloadGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadGroup)(nil)).Elem()
}

func (i *WorkloadGroup) ToWorkloadGroupOutput() WorkloadGroupOutput {
	return i.ToWorkloadGroupOutputWithContext(context.Background())
}

func (i *WorkloadGroup) ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadGroupOutput)
}

func (i *WorkloadGroup) ToOutput(ctx context.Context) pulumix.Output[*WorkloadGroup] {
	return pulumix.Output[*WorkloadGroup]{
		OutputState: i.ToWorkloadGroupOutputWithContext(ctx).OutputState,
	}
}

type WorkloadGroupOutput struct{ *pulumi.OutputState }

func (WorkloadGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadGroup)(nil)).Elem()
}

func (o WorkloadGroupOutput) ToWorkloadGroupOutput() WorkloadGroupOutput {
	return o
}

func (o WorkloadGroupOutput) ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput {
	return o
}

func (o WorkloadGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkloadGroup] {
	return pulumix.Output[*WorkloadGroup]{
		OutputState: o.OutputState,
	}
}

// The workload group importance level.
func (o WorkloadGroupOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.StringPtrOutput { return v.Importance }).(pulumi.StringPtrOutput)
}

// The workload group cap percentage resource.
func (o WorkloadGroupOutput) MaxResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.IntOutput { return v.MaxResourcePercent }).(pulumi.IntOutput)
}

// The workload group request maximum grant percentage.
func (o WorkloadGroupOutput) MaxResourcePercentPerRequest() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.Float64PtrOutput { return v.MaxResourcePercentPerRequest }).(pulumi.Float64PtrOutput)
}

// The workload group minimum percentage resource.
func (o WorkloadGroupOutput) MinResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.IntOutput { return v.MinResourcePercent }).(pulumi.IntOutput)
}

// The workload group request minimum grant percentage.
func (o WorkloadGroupOutput) MinResourcePercentPerRequest() pulumi.Float64Output {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.Float64Output { return v.MinResourcePercentPerRequest }).(pulumi.Float64Output)
}

// Resource name.
func (o WorkloadGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The workload group query execution timeout.
func (o WorkloadGroupOutput) QueryExecutionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.IntPtrOutput { return v.QueryExecutionTimeout }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o WorkloadGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadGroupOutput{})
}
