// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// a dryrun job resource
type LinkerDryrun struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// the preview of the operations for creation
	OperationPreviews DryrunOperationPreviewResponseArrayOutput `pulumi:"operationPreviews"`
	// The parameters of the dryrun
	Parameters CreateOrUpdateDryrunParametersResponsePtrOutput `pulumi:"parameters"`
	// the result of the dryrun
	PrerequisiteResults pulumi.ArrayOutput `pulumi:"prerequisiteResults"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLinkerDryrun registers a new resource with the given unique name, arguments, and options.
func NewLinkerDryrun(ctx *pulumi.Context,
	name string, args *LinkerDryrunArgs, opts ...pulumi.ResourceOption) (*LinkerDryrun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicelinker:LinkerDryrun"),
		},
		{
			Type: pulumi.String("azure-native:servicelinker/v20230401preview:LinkerDryrun"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LinkerDryrun
	err := ctx.RegisterResource("azure-native:servicelinker/v20221101preview:LinkerDryrun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkerDryrun gets an existing LinkerDryrun resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkerDryrun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkerDryrunState, opts ...pulumi.ResourceOption) (*LinkerDryrun, error) {
	var resource LinkerDryrun
	err := ctx.ReadResource("azure-native:servicelinker/v20221101preview:LinkerDryrun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkerDryrun resources.
type linkerDryrunState struct {
}

type LinkerDryrunState struct {
}

func (LinkerDryrunState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkerDryrunState)(nil)).Elem()
}

type linkerDryrunArgs struct {
	// The name of dryrun.
	DryrunName *string `pulumi:"dryrunName"`
	// The parameters of the dryrun
	Parameters *CreateOrUpdateDryrunParameters `pulumi:"parameters"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri string `pulumi:"resourceUri"`
}

// The set of arguments for constructing a LinkerDryrun resource.
type LinkerDryrunArgs struct {
	// The name of dryrun.
	DryrunName pulumi.StringPtrInput
	// The parameters of the dryrun
	Parameters CreateOrUpdateDryrunParametersPtrInput
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri pulumi.StringInput
}

func (LinkerDryrunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkerDryrunArgs)(nil)).Elem()
}

type LinkerDryrunInput interface {
	pulumi.Input

	ToLinkerDryrunOutput() LinkerDryrunOutput
	ToLinkerDryrunOutputWithContext(ctx context.Context) LinkerDryrunOutput
}

func (*LinkerDryrun) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkerDryrun)(nil)).Elem()
}

func (i *LinkerDryrun) ToLinkerDryrunOutput() LinkerDryrunOutput {
	return i.ToLinkerDryrunOutputWithContext(context.Background())
}

func (i *LinkerDryrun) ToLinkerDryrunOutputWithContext(ctx context.Context) LinkerDryrunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkerDryrunOutput)
}

type LinkerDryrunOutput struct{ *pulumi.OutputState }

func (LinkerDryrunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkerDryrun)(nil)).Elem()
}

func (o LinkerDryrunOutput) ToLinkerDryrunOutput() LinkerDryrunOutput {
	return o
}

func (o LinkerDryrunOutput) ToLinkerDryrunOutputWithContext(ctx context.Context) LinkerDryrunOutput {
	return o
}

// The name of the resource
func (o LinkerDryrunOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkerDryrun) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the preview of the operations for creation
func (o LinkerDryrunOutput) OperationPreviews() DryrunOperationPreviewResponseArrayOutput {
	return o.ApplyT(func(v *LinkerDryrun) DryrunOperationPreviewResponseArrayOutput { return v.OperationPreviews }).(DryrunOperationPreviewResponseArrayOutput)
}

// The parameters of the dryrun
func (o LinkerDryrunOutput) Parameters() CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o.ApplyT(func(v *LinkerDryrun) CreateOrUpdateDryrunParametersResponsePtrOutput { return v.Parameters }).(CreateOrUpdateDryrunParametersResponsePtrOutput)
}

// the result of the dryrun
func (o LinkerDryrunOutput) PrerequisiteResults() pulumi.ArrayOutput {
	return o.ApplyT(func(v *LinkerDryrun) pulumi.ArrayOutput { return v.PrerequisiteResults }).(pulumi.ArrayOutput)
}

// The provisioning state.
func (o LinkerDryrunOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkerDryrun) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LinkerDryrunOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LinkerDryrun) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LinkerDryrunOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkerDryrun) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkerDryrunOutput{})
}
