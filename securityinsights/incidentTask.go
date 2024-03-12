// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure REST API version: 2023-06-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview.
type IncidentTask struct {
	pulumi.CustomResourceState

	// Information on the client (user or application) that made some action
	CreatedBy ClientInfoResponsePtrOutput `pulumi:"createdBy"`
	// The time the task was created
	CreatedTimeUtc pulumi.StringOutput `pulumi:"createdTimeUtc"`
	// The description of the task
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Information on the client (user or application) that made some action
	LastModifiedBy ClientInfoResponsePtrOutput `pulumi:"lastModifiedBy"`
	// The last time the task was updated
	LastModifiedTimeUtc pulumi.StringOutput `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name   pulumi.StringOutput `pulumi:"name"`
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The title of the task
	Title pulumi.StringOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIncidentTask registers a new resource with the given unique name, arguments, and options.
func NewIncidentTask(ctx *pulumi.Context,
	name string, args *IncidentTaskArgs, opts ...pulumi.ResourceOption) (*IncidentTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncidentId == nil {
		return nil, errors.New("invalid value for required argument 'IncidentId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:IncidentTask"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:IncidentTask"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IncidentTask
	err := ctx.RegisterResource("azure-native:securityinsights:IncidentTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIncidentTask gets an existing IncidentTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIncidentTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentTaskState, opts ...pulumi.ResourceOption) (*IncidentTask, error) {
	var resource IncidentTask
	err := ctx.ReadResource("azure-native:securityinsights:IncidentTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IncidentTask resources.
type incidentTaskState struct {
}

type IncidentTaskState struct {
}

func (IncidentTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentTaskState)(nil)).Elem()
}

type incidentTaskArgs struct {
	// Information on the client (user or application) that made some action
	CreatedBy *ClientInfo `pulumi:"createdBy"`
	// The description of the task
	Description *string `pulumi:"description"`
	// Incident ID
	IncidentId string `pulumi:"incidentId"`
	// Incident task ID
	IncidentTaskId *string `pulumi:"incidentTaskId"`
	// Information on the client (user or application) that made some action
	LastModifiedBy *ClientInfo `pulumi:"lastModifiedBy"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Status            string `pulumi:"status"`
	// The title of the task
	Title string `pulumi:"title"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IncidentTask resource.
type IncidentTaskArgs struct {
	// Information on the client (user or application) that made some action
	CreatedBy ClientInfoPtrInput
	// The description of the task
	Description pulumi.StringPtrInput
	// Incident ID
	IncidentId pulumi.StringInput
	// Incident task ID
	IncidentTaskId pulumi.StringPtrInput
	// Information on the client (user or application) that made some action
	LastModifiedBy ClientInfoPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	Status            pulumi.StringInput
	// The title of the task
	Title pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IncidentTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentTaskArgs)(nil)).Elem()
}

type IncidentTaskInput interface {
	pulumi.Input

	ToIncidentTaskOutput() IncidentTaskOutput
	ToIncidentTaskOutputWithContext(ctx context.Context) IncidentTaskOutput
}

func (*IncidentTask) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentTask)(nil)).Elem()
}

func (i *IncidentTask) ToIncidentTaskOutput() IncidentTaskOutput {
	return i.ToIncidentTaskOutputWithContext(context.Background())
}

func (i *IncidentTask) ToIncidentTaskOutputWithContext(ctx context.Context) IncidentTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentTaskOutput)
}

type IncidentTaskOutput struct{ *pulumi.OutputState }

func (IncidentTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentTask)(nil)).Elem()
}

func (o IncidentTaskOutput) ToIncidentTaskOutput() IncidentTaskOutput {
	return o
}

func (o IncidentTaskOutput) ToIncidentTaskOutputWithContext(ctx context.Context) IncidentTaskOutput {
	return o
}

// Information on the client (user or application) that made some action
func (o IncidentTaskOutput) CreatedBy() ClientInfoResponsePtrOutput {
	return o.ApplyT(func(v *IncidentTask) ClientInfoResponsePtrOutput { return v.CreatedBy }).(ClientInfoResponsePtrOutput)
}

// The time the task was created
func (o IncidentTaskOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// The description of the task
func (o IncidentTaskOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Etag of the azure resource
func (o IncidentTaskOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Information on the client (user or application) that made some action
func (o IncidentTaskOutput) LastModifiedBy() ClientInfoResponsePtrOutput {
	return o.ApplyT(func(v *IncidentTask) ClientInfoResponsePtrOutput { return v.LastModifiedBy }).(ClientInfoResponsePtrOutput)
}

// The last time the task was updated
func (o IncidentTaskOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o IncidentTaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IncidentTaskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IncidentTaskOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IncidentTask) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The title of the task
func (o IncidentTaskOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IncidentTaskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentTask) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentTaskOutput{})
}
