// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An application package which represents a particular version of an application.
// Azure REST API version: 2023-05-01. Prior API version in Azure Native 1.x: 2021-01-01.
//
// Other available API versions: 2017-09-01.
type ApplicationPackage struct {
	pulumi.CustomResourceState

	// The ETag of the resource, used for concurrency statements.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The format of the application package, if the package is active.
	Format pulumi.StringOutput `pulumi:"format"`
	// The time at which the package was last activated, if the package is active.
	LastActivationTime pulumi.StringOutput `pulumi:"lastActivationTime"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current state of the application package.
	State pulumi.StringOutput `pulumi:"state"`
	// The URL for the application package in Azure Storage.
	StorageUrl pulumi.StringOutput `pulumi:"storageUrl"`
	// The UTC time at which the Azure Storage URL will expire.
	StorageUrlExpiry pulumi.StringOutput `pulumi:"storageUrlExpiry"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationPackage registers a new resource with the given unique name, arguments, and options.
func NewApplicationPackage(ctx *pulumi.Context,
	name string, args *ApplicationPackageArgs, opts ...pulumi.ResourceOption) (*ApplicationPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:batch/v20151201:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170501:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220101:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220601:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20221001:ApplicationPackage"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20230501:ApplicationPackage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationPackage
	err := ctx.RegisterResource("azure-native:batch:ApplicationPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationPackage gets an existing ApplicationPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationPackageState, opts ...pulumi.ResourceOption) (*ApplicationPackage, error) {
	var resource ApplicationPackage
	err := ctx.ReadResource("azure-native:batch:ApplicationPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationPackage resources.
type applicationPackageState struct {
}

type ApplicationPackageState struct {
}

func (ApplicationPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPackageState)(nil)).Elem()
}

type applicationPackageArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The name of the application. This must be unique within the account.
	ApplicationName string `pulumi:"applicationName"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The version of the application.
	VersionName *string `pulumi:"versionName"`
}

// The set of arguments for constructing a ApplicationPackage resource.
type ApplicationPackageArgs struct {
	// The name of the Batch account.
	AccountName pulumi.StringInput
	// The name of the application. This must be unique within the account.
	ApplicationName pulumi.StringInput
	// The name of the resource group that contains the Batch account.
	ResourceGroupName pulumi.StringInput
	// The version of the application.
	VersionName pulumi.StringPtrInput
}

func (ApplicationPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPackageArgs)(nil)).Elem()
}

type ApplicationPackageInput interface {
	pulumi.Input

	ToApplicationPackageOutput() ApplicationPackageOutput
	ToApplicationPackageOutputWithContext(ctx context.Context) ApplicationPackageOutput
}

func (*ApplicationPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackage)(nil)).Elem()
}

func (i *ApplicationPackage) ToApplicationPackageOutput() ApplicationPackageOutput {
	return i.ToApplicationPackageOutputWithContext(context.Background())
}

func (i *ApplicationPackage) ToApplicationPackageOutputWithContext(ctx context.Context) ApplicationPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPackageOutput)
}

func (i *ApplicationPackage) ToOutput(ctx context.Context) pulumix.Output[*ApplicationPackage] {
	return pulumix.Output[*ApplicationPackage]{
		OutputState: i.ToApplicationPackageOutputWithContext(ctx).OutputState,
	}
}

type ApplicationPackageOutput struct{ *pulumi.OutputState }

func (ApplicationPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPackage)(nil)).Elem()
}

func (o ApplicationPackageOutput) ToApplicationPackageOutput() ApplicationPackageOutput {
	return o
}

func (o ApplicationPackageOutput) ToApplicationPackageOutputWithContext(ctx context.Context) ApplicationPackageOutput {
	return o
}

func (o ApplicationPackageOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationPackage] {
	return pulumix.Output[*ApplicationPackage]{
		OutputState: o.OutputState,
	}
}

// The ETag of the resource, used for concurrency statements.
func (o ApplicationPackageOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPackage) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The format of the application package, if the package is active.
func (o ApplicationPackageOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPackage) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// The time at which the package was last activated, if the package is active.
func (o ApplicationPackageOutput) LastActivationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPackage) pulumi.StringOutput { return v.LastActivationTime }).(pulumi.StringOutput)
}

// The name of the resource.
func (o ApplicationPackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPackage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current state of the application package.
func (o ApplicationPackageOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPackage) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The URL for the application package in Azure Storage.
func (o ApplicationPackageOutput) StorageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPackage) pulumi.StringOutput { return v.StorageUrl }).(pulumi.StringOutput)
}

// The UTC time at which the Azure Storage URL will expire.
func (o ApplicationPackageOutput) StorageUrlExpiry() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPackage) pulumi.StringOutput { return v.StorageUrlExpiry }).(pulumi.StringOutput)
}

// The type of the resource.
func (o ApplicationPackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPackage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationPackageOutput{})
}
