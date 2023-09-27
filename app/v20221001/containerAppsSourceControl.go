// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Container App SourceControl.
type ContainerAppsSourceControl struct {
	pulumi.CustomResourceState

	// The branch which will trigger the auto deployment
	Branch pulumi.StringPtrOutput `pulumi:"branch"`
	// Container App Revision Template with all possible settings and the
	// defaults if user did not provide them. The defaults are populated
	// as they were at the creation time
	GithubActionConfiguration GithubActionConfigurationResponsePtrOutput `pulumi:"githubActionConfiguration"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Current provisioning State of the operation
	OperationState pulumi.StringOutput `pulumi:"operationState"`
	// The repo url which will be integrated to ContainerApp.
	RepoUrl pulumi.StringPtrOutput `pulumi:"repoUrl"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContainerAppsSourceControl registers a new resource with the given unique name, arguments, and options.
func NewContainerAppsSourceControl(ctx *pulumi.Context,
	name string, args *ContainerAppsSourceControlArgs, opts ...pulumi.ResourceOption) (*ContainerAppsSourceControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerAppName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerAppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ContainerAppsSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220101preview:ContainerAppsSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:ContainerAppsSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ContainerAppsSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:ContainerAppsSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230401preview:ContainerAppsSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230501:ContainerAppsSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230502preview:ContainerAppsSourceControl"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ContainerAppsSourceControl
	err := ctx.RegisterResource("azure-native:app/v20221001:ContainerAppsSourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerAppsSourceControl gets an existing ContainerAppsSourceControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerAppsSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAppsSourceControlState, opts ...pulumi.ResourceOption) (*ContainerAppsSourceControl, error) {
	var resource ContainerAppsSourceControl
	err := ctx.ReadResource("azure-native:app/v20221001:ContainerAppsSourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerAppsSourceControl resources.
type containerAppsSourceControlState struct {
}

type ContainerAppsSourceControlState struct {
}

func (ContainerAppsSourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppsSourceControlState)(nil)).Elem()
}

type containerAppsSourceControlArgs struct {
	// The branch which will trigger the auto deployment
	Branch *string `pulumi:"branch"`
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// Container App Revision Template with all possible settings and the
	// defaults if user did not provide them. The defaults are populated
	// as they were at the creation time
	GithubActionConfiguration *GithubActionConfiguration `pulumi:"githubActionConfiguration"`
	// The repo url which will be integrated to ContainerApp.
	RepoUrl *string `pulumi:"repoUrl"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Container App SourceControl.
	SourceControlName *string `pulumi:"sourceControlName"`
}

// The set of arguments for constructing a ContainerAppsSourceControl resource.
type ContainerAppsSourceControlArgs struct {
	// The branch which will trigger the auto deployment
	Branch pulumi.StringPtrInput
	// Name of the Container App.
	ContainerAppName pulumi.StringInput
	// Container App Revision Template with all possible settings and the
	// defaults if user did not provide them. The defaults are populated
	// as they were at the creation time
	GithubActionConfiguration GithubActionConfigurationPtrInput
	// The repo url which will be integrated to ContainerApp.
	RepoUrl pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of the Container App SourceControl.
	SourceControlName pulumi.StringPtrInput
}

func (ContainerAppsSourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppsSourceControlArgs)(nil)).Elem()
}

type ContainerAppsSourceControlInput interface {
	pulumi.Input

	ToContainerAppsSourceControlOutput() ContainerAppsSourceControlOutput
	ToContainerAppsSourceControlOutputWithContext(ctx context.Context) ContainerAppsSourceControlOutput
}

func (*ContainerAppsSourceControl) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsSourceControl)(nil)).Elem()
}

func (i *ContainerAppsSourceControl) ToContainerAppsSourceControlOutput() ContainerAppsSourceControlOutput {
	return i.ToContainerAppsSourceControlOutputWithContext(context.Background())
}

func (i *ContainerAppsSourceControl) ToContainerAppsSourceControlOutputWithContext(ctx context.Context) ContainerAppsSourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppsSourceControlOutput)
}

func (i *ContainerAppsSourceControl) ToOutput(ctx context.Context) pulumix.Output[*ContainerAppsSourceControl] {
	return pulumix.Output[*ContainerAppsSourceControl]{
		OutputState: i.ToContainerAppsSourceControlOutputWithContext(ctx).OutputState,
	}
}

type ContainerAppsSourceControlOutput struct{ *pulumi.OutputState }

func (ContainerAppsSourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsSourceControl)(nil)).Elem()
}

func (o ContainerAppsSourceControlOutput) ToContainerAppsSourceControlOutput() ContainerAppsSourceControlOutput {
	return o
}

func (o ContainerAppsSourceControlOutput) ToContainerAppsSourceControlOutputWithContext(ctx context.Context) ContainerAppsSourceControlOutput {
	return o
}

func (o ContainerAppsSourceControlOutput) ToOutput(ctx context.Context) pulumix.Output[*ContainerAppsSourceControl] {
	return pulumix.Output[*ContainerAppsSourceControl]{
		OutputState: o.OutputState,
	}
}

// The branch which will trigger the auto deployment
func (o ContainerAppsSourceControlOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

// Container App Revision Template with all possible settings and the
// defaults if user did not provide them. The defaults are populated
// as they were at the creation time
func (o ContainerAppsSourceControlOutput) GithubActionConfiguration() GithubActionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) GithubActionConfigurationResponsePtrOutput {
		return v.GithubActionConfiguration
	}).(GithubActionConfigurationResponsePtrOutput)
}

// The name of the resource
func (o ContainerAppsSourceControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Current provisioning State of the operation
func (o ContainerAppsSourceControlOutput) OperationState() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringOutput { return v.OperationState }).(pulumi.StringOutput)
}

// The repo url which will be integrated to ContainerApp.
func (o ContainerAppsSourceControlOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringPtrOutput { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ContainerAppsSourceControlOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ContainerAppsSourceControlOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerAppsSourceControlOutput{})
}
