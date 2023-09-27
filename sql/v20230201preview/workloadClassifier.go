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

// Workload classifier operations for a data warehouse
type WorkloadClassifier struct {
	pulumi.CustomResourceState

	// The workload classifier context.
	Context pulumi.StringPtrOutput `pulumi:"context"`
	// The workload classifier end time for classification.
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// The workload classifier importance.
	Importance pulumi.StringPtrOutput `pulumi:"importance"`
	// The workload classifier label.
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// The workload classifier member name.
	MemberName pulumi.StringOutput `pulumi:"memberName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The workload classifier start time for classification.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadClassifier registers a new resource with the given unique name, arguments, and options.
func NewWorkloadClassifier(ctx *pulumi.Context,
	name string, args *WorkloadClassifierArgs, opts ...pulumi.ResourceOption) (*WorkloadClassifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.MemberName == nil {
		return nil, errors.New("invalid value for required argument 'MemberName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.WorkloadGroupName == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:WorkloadClassifier"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:WorkloadClassifier"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadClassifier
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:WorkloadClassifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadClassifier gets an existing WorkloadClassifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadClassifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadClassifierState, opts ...pulumi.ResourceOption) (*WorkloadClassifier, error) {
	var resource WorkloadClassifier
	err := ctx.ReadResource("azure-native:sql/v20230201preview:WorkloadClassifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadClassifier resources.
type workloadClassifierState struct {
}

type WorkloadClassifierState struct {
}

func (WorkloadClassifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadClassifierState)(nil)).Elem()
}

type workloadClassifierArgs struct {
	// The workload classifier context.
	Context *string `pulumi:"context"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The workload classifier end time for classification.
	EndTime *string `pulumi:"endTime"`
	// The workload classifier importance.
	Importance *string `pulumi:"importance"`
	// The workload classifier label.
	Label *string `pulumi:"label"`
	// The workload classifier member name.
	MemberName string `pulumi:"memberName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The workload classifier start time for classification.
	StartTime *string `pulumi:"startTime"`
	// The name of the workload classifier to create/update.
	WorkloadClassifierName *string `pulumi:"workloadClassifierName"`
	// The name of the workload group from which to receive the classifier from.
	WorkloadGroupName string `pulumi:"workloadGroupName"`
}

// The set of arguments for constructing a WorkloadClassifier resource.
type WorkloadClassifierArgs struct {
	// The workload classifier context.
	Context pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The workload classifier end time for classification.
	EndTime pulumi.StringPtrInput
	// The workload classifier importance.
	Importance pulumi.StringPtrInput
	// The workload classifier label.
	Label pulumi.StringPtrInput
	// The workload classifier member name.
	MemberName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The workload classifier start time for classification.
	StartTime pulumi.StringPtrInput
	// The name of the workload classifier to create/update.
	WorkloadClassifierName pulumi.StringPtrInput
	// The name of the workload group from which to receive the classifier from.
	WorkloadGroupName pulumi.StringInput
}

func (WorkloadClassifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadClassifierArgs)(nil)).Elem()
}

type WorkloadClassifierInput interface {
	pulumi.Input

	ToWorkloadClassifierOutput() WorkloadClassifierOutput
	ToWorkloadClassifierOutputWithContext(ctx context.Context) WorkloadClassifierOutput
}

func (*WorkloadClassifier) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadClassifier)(nil)).Elem()
}

func (i *WorkloadClassifier) ToWorkloadClassifierOutput() WorkloadClassifierOutput {
	return i.ToWorkloadClassifierOutputWithContext(context.Background())
}

func (i *WorkloadClassifier) ToWorkloadClassifierOutputWithContext(ctx context.Context) WorkloadClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadClassifierOutput)
}

func (i *WorkloadClassifier) ToOutput(ctx context.Context) pulumix.Output[*WorkloadClassifier] {
	return pulumix.Output[*WorkloadClassifier]{
		OutputState: i.ToWorkloadClassifierOutputWithContext(ctx).OutputState,
	}
}

type WorkloadClassifierOutput struct{ *pulumi.OutputState }

func (WorkloadClassifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadClassifier)(nil)).Elem()
}

func (o WorkloadClassifierOutput) ToWorkloadClassifierOutput() WorkloadClassifierOutput {
	return o
}

func (o WorkloadClassifierOutput) ToWorkloadClassifierOutputWithContext(ctx context.Context) WorkloadClassifierOutput {
	return o
}

func (o WorkloadClassifierOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkloadClassifier] {
	return pulumix.Output[*WorkloadClassifier]{
		OutputState: o.OutputState,
	}
}

// The workload classifier context.
func (o WorkloadClassifierOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.Context }).(pulumi.StringPtrOutput)
}

// The workload classifier end time for classification.
func (o WorkloadClassifierOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The workload classifier importance.
func (o WorkloadClassifierOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.Importance }).(pulumi.StringPtrOutput)
}

// The workload classifier label.
func (o WorkloadClassifierOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// The workload classifier member name.
func (o WorkloadClassifierOutput) MemberName() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringOutput { return v.MemberName }).(pulumi.StringOutput)
}

// Resource name.
func (o WorkloadClassifierOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The workload classifier start time for classification.
func (o WorkloadClassifierOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WorkloadClassifierOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadClassifier) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadClassifierOutput{})
}
