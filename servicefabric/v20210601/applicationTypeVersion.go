// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An application type version resource for the specified application type name resource.
type ApplicationTypeVersion struct {
	pulumi.CustomResourceState

	// The URL to the application package
	AppPackageUrl pulumi.StringOutput `pulumi:"appPackageUrl"`
	// List of application type parameters that can be overridden when creating or updating the application.
	DefaultParameterList pulumi.StringMapOutput `pulumi:"defaultParameterList"`
	// Azure resource etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationTypeVersion registers a new resource with the given unique name, arguments, and options.
func NewApplicationTypeVersion(ctx *pulumi.Context,
	name string, args *ApplicationTypeVersionArgs, opts ...pulumi.ResourceOption) (*ApplicationTypeVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppPackageUrl == nil {
		return nil, errors.New("invalid value for required argument 'AppPackageUrl'")
	}
	if args.ApplicationTypeName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationTypeName'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric/v20170701preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190601preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20191101preview:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200301:ApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20201201preview:ApplicationTypeVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationTypeVersion
	err := ctx.RegisterResource("azure-native:servicefabric/v20210601:ApplicationTypeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationTypeVersion gets an existing ApplicationTypeVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationTypeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationTypeVersionState, opts ...pulumi.ResourceOption) (*ApplicationTypeVersion, error) {
	var resource ApplicationTypeVersion
	err := ctx.ReadResource("azure-native:servicefabric/v20210601:ApplicationTypeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationTypeVersion resources.
type applicationTypeVersionState struct {
}

type ApplicationTypeVersionState struct {
}

func (ApplicationTypeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeVersionState)(nil)).Elem()
}

type applicationTypeVersionArgs struct {
	// The URL to the application package
	AppPackageUrl string `pulumi:"appPackageUrl"`
	// The name of the application type name resource.
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The application type version.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a ApplicationTypeVersion resource.
type ApplicationTypeVersionArgs struct {
	// The URL to the application package
	AppPackageUrl pulumi.StringInput
	// The name of the application type name resource.
	ApplicationTypeName pulumi.StringInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// The application type version.
	Version pulumi.StringPtrInput
}

func (ApplicationTypeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeVersionArgs)(nil)).Elem()
}

type ApplicationTypeVersionInput interface {
	pulumi.Input

	ToApplicationTypeVersionOutput() ApplicationTypeVersionOutput
	ToApplicationTypeVersionOutputWithContext(ctx context.Context) ApplicationTypeVersionOutput
}

func (*ApplicationTypeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersion)(nil)).Elem()
}

func (i *ApplicationTypeVersion) ToApplicationTypeVersionOutput() ApplicationTypeVersionOutput {
	return i.ToApplicationTypeVersionOutputWithContext(context.Background())
}

func (i *ApplicationTypeVersion) ToApplicationTypeVersionOutputWithContext(ctx context.Context) ApplicationTypeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionOutput)
}

func (i *ApplicationTypeVersion) ToOutput(ctx context.Context) pulumix.Output[*ApplicationTypeVersion] {
	return pulumix.Output[*ApplicationTypeVersion]{
		OutputState: i.ToApplicationTypeVersionOutputWithContext(ctx).OutputState,
	}
}

type ApplicationTypeVersionOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersion)(nil)).Elem()
}

func (o ApplicationTypeVersionOutput) ToApplicationTypeVersionOutput() ApplicationTypeVersionOutput {
	return o
}

func (o ApplicationTypeVersionOutput) ToApplicationTypeVersionOutputWithContext(ctx context.Context) ApplicationTypeVersionOutput {
	return o
}

func (o ApplicationTypeVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationTypeVersion] {
	return pulumix.Output[*ApplicationTypeVersion]{
		OutputState: o.OutputState,
	}
}

// The URL to the application package
func (o ApplicationTypeVersionOutput) AppPackageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.AppPackageUrl }).(pulumi.StringOutput)
}

// List of application type parameters that can be overridden when creating or updating the application.
func (o ApplicationTypeVersionOutput) DefaultParameterList() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringMapOutput { return v.DefaultParameterList }).(pulumi.StringMapOutput)
}

// Azure resource etag.
func (o ApplicationTypeVersionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// It will be deprecated in New API, resource location depends on the parent resource.
func (o ApplicationTypeVersionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o ApplicationTypeVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response
func (o ApplicationTypeVersionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ApplicationTypeVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o ApplicationTypeVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o ApplicationTypeVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationTypeVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationTypeVersionOutput{})
}
