// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220330preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A SQL Migration Service.
type SqlMigrationService struct {
	pulumi.CustomResourceState

	// Current state of the Integration runtime.
	IntegrationRuntimeState pulumi.StringOutput    `pulumi:"integrationRuntimeState"`
	Location                pulumi.StringPtrOutput `pulumi:"location"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	// Provisioning state to track the async operation status.
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}

// NewSqlMigrationService registers a new resource with the given unique name, arguments, and options.
func NewSqlMigrationService(ctx *pulumi.Context,
	name string, args *SqlMigrationServiceArgs, opts ...pulumi.ResourceOption) (*SqlMigrationService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration:SqlMigrationService"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:SqlMigrationService"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:SqlMigrationService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlMigrationService
	err := ctx.RegisterResource("azure-native:datamigration/v20220330preview:SqlMigrationService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlMigrationService gets an existing SqlMigrationService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlMigrationService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlMigrationServiceState, opts ...pulumi.ResourceOption) (*SqlMigrationService, error) {
	var resource SqlMigrationService
	err := ctx.ReadResource("azure-native:datamigration/v20220330preview:SqlMigrationService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlMigrationService resources.
type sqlMigrationServiceState struct {
}

type SqlMigrationServiceState struct {
}

func (SqlMigrationServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlMigrationServiceState)(nil)).Elem()
}

type sqlMigrationServiceArgs struct {
	Location *string `pulumi:"location"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL Migration Service.
	SqlMigrationServiceName *string           `pulumi:"sqlMigrationServiceName"`
	Tags                    map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqlMigrationService resource.
type SqlMigrationServiceArgs struct {
	Location pulumi.StringPtrInput
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Name of the SQL Migration Service.
	SqlMigrationServiceName pulumi.StringPtrInput
	Tags                    pulumi.StringMapInput
}

func (SqlMigrationServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlMigrationServiceArgs)(nil)).Elem()
}

type SqlMigrationServiceInput interface {
	pulumi.Input

	ToSqlMigrationServiceOutput() SqlMigrationServiceOutput
	ToSqlMigrationServiceOutputWithContext(ctx context.Context) SqlMigrationServiceOutput
}

func (*SqlMigrationService) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlMigrationService)(nil)).Elem()
}

func (i *SqlMigrationService) ToSqlMigrationServiceOutput() SqlMigrationServiceOutput {
	return i.ToSqlMigrationServiceOutputWithContext(context.Background())
}

func (i *SqlMigrationService) ToSqlMigrationServiceOutputWithContext(ctx context.Context) SqlMigrationServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlMigrationServiceOutput)
}

func (i *SqlMigrationService) ToOutput(ctx context.Context) pulumix.Output[*SqlMigrationService] {
	return pulumix.Output[*SqlMigrationService]{
		OutputState: i.ToSqlMigrationServiceOutputWithContext(ctx).OutputState,
	}
}

type SqlMigrationServiceOutput struct{ *pulumi.OutputState }

func (SqlMigrationServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlMigrationService)(nil)).Elem()
}

func (o SqlMigrationServiceOutput) ToSqlMigrationServiceOutput() SqlMigrationServiceOutput {
	return o
}

func (o SqlMigrationServiceOutput) ToSqlMigrationServiceOutputWithContext(ctx context.Context) SqlMigrationServiceOutput {
	return o
}

func (o SqlMigrationServiceOutput) ToOutput(ctx context.Context) pulumix.Output[*SqlMigrationService] {
	return pulumix.Output[*SqlMigrationService]{
		OutputState: o.OutputState,
	}
}

// Current state of the Integration runtime.
func (o SqlMigrationServiceOutput) IntegrationRuntimeState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringOutput { return v.IntegrationRuntimeState }).(pulumi.StringOutput)
}

func (o SqlMigrationServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SqlMigrationServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state to track the async operation status.
func (o SqlMigrationServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlMigrationServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlMigrationService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlMigrationServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlMigrationServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlMigrationServiceOutput{})
}
