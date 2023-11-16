// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App Job
type Job struct {
	pulumi.CustomResourceState

	// Container Apps Job configuration properties.
	Configuration JobConfigurationResponsePtrOutput `pulumi:"configuration"`
	// Resource ID of environment.
	EnvironmentId pulumi.StringPtrOutput `pulumi:"environmentId"`
	// The endpoint of the eventstream of the container apps job.
	EventStreamEndpoint pulumi.StringOutput `pulumi:"eventStreamEndpoint"`
	// The complex type of the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Managed identities needed by a container app job to interact with other Azure services to not maintain any secrets or credentials in code.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Outbound IP Addresses of a container apps job.
	OutboundIpAddresses pulumi.StringArrayOutput `pulumi:"outboundIpAddresses"`
	// Provisioning state of the Container Apps Job.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Container Apps job definition.
	Template JobTemplateResponsePtrOutput `pulumi:"template"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Workload profile name to pin for container apps job execution.
	WorkloadProfileName pulumi.StringPtrOutput `pulumi:"workloadProfileName"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Configuration != nil {
		args.Configuration = args.Configuration.ToJobConfigurationPtrOutput().ApplyT(func(v *JobConfiguration) *JobConfiguration { return v.Defaults() }).(JobConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:Job"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230401preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230501:Job"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230502preview:Job"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Job
	err := ctx.RegisterResource("azure-native:app/v20230801preview:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:app/v20230801preview:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// Container Apps Job configuration properties.
	Configuration *JobConfiguration `pulumi:"configuration"`
	// Resource ID of environment.
	EnvironmentId *string `pulumi:"environmentId"`
	// The complex type of the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Managed identities needed by a container app job to interact with other Azure services to not maintain any secrets or credentials in code.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Job Name
	JobName *string `pulumi:"jobName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Container Apps job definition.
	Template *JobTemplate `pulumi:"template"`
	// Workload profile name to pin for container apps job execution.
	WorkloadProfileName *string `pulumi:"workloadProfileName"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// Container Apps Job configuration properties.
	Configuration JobConfigurationPtrInput
	// Resource ID of environment.
	EnvironmentId pulumi.StringPtrInput
	// The complex type of the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// Managed identities needed by a container app job to interact with other Azure services to not maintain any secrets or credentials in code.
	Identity ManagedServiceIdentityPtrInput
	// Job Name
	JobName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Container Apps job definition.
	Template JobTemplatePtrInput
	// Workload profile name to pin for container apps job execution.
	WorkloadProfileName pulumi.StringPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

// Container Apps Job configuration properties.
func (o JobOutput) Configuration() JobConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Job) JobConfigurationResponsePtrOutput { return v.Configuration }).(JobConfigurationResponsePtrOutput)
}

// Resource ID of environment.
func (o JobOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// The endpoint of the eventstream of the container apps job.
func (o JobOutput) EventStreamEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.EventStreamEndpoint }).(pulumi.StringOutput)
}

// The complex type of the extended location.
func (o JobOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Job) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Managed identities needed by a container app job to interact with other Azure services to not maintain any secrets or credentials in code.
func (o JobOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Job) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o JobOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Outbound IP Addresses of a container apps job.
func (o JobOutput) OutboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Job) pulumi.StringArrayOutput { return v.OutboundIpAddresses }).(pulumi.StringArrayOutput)
}

// Provisioning state of the Container Apps Job.
func (o JobOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o JobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Job) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o JobOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Container Apps job definition.
func (o JobOutput) Template() JobTemplateResponsePtrOutput {
	return o.ApplyT(func(v *Job) JobTemplateResponsePtrOutput { return v.Template }).(JobTemplateResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o JobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Workload profile name to pin for container apps job execution.
func (o JobOutput) WorkloadProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.WorkloadProfileName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
