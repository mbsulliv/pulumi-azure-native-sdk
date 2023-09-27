// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:documentdb/v20230915preview:CassandraCluster":
		r = &CassandraCluster{}
	case "azure-native:documentdb/v20230915preview:CassandraDataCenter":
		r = &CassandraDataCenter{}
	case "azure-native:documentdb/v20230915preview:CassandraResourceCassandraKeyspace":
		r = &CassandraResourceCassandraKeyspace{}
	case "azure-native:documentdb/v20230915preview:CassandraResourceCassandraTable":
		r = &CassandraResourceCassandraTable{}
	case "azure-native:documentdb/v20230915preview:CassandraResourceCassandraView":
		r = &CassandraResourceCassandraView{}
	case "azure-native:documentdb/v20230915preview:DatabaseAccount":
		r = &DatabaseAccount{}
	case "azure-native:documentdb/v20230915preview:GraphResourceGraph":
		r = &GraphResourceGraph{}
	case "azure-native:documentdb/v20230915preview:GremlinResourceGremlinDatabase":
		r = &GremlinResourceGremlinDatabase{}
	case "azure-native:documentdb/v20230915preview:GremlinResourceGremlinGraph":
		r = &GremlinResourceGremlinGraph{}
	case "azure-native:documentdb/v20230915preview:MongoCluster":
		r = &MongoCluster{}
	case "azure-native:documentdb/v20230915preview:MongoClusterFirewallRule":
		r = &MongoClusterFirewallRule{}
	case "azure-native:documentdb/v20230915preview:MongoDBResourceMongoDBCollection":
		r = &MongoDBResourceMongoDBCollection{}
	case "azure-native:documentdb/v20230915preview:MongoDBResourceMongoDBDatabase":
		r = &MongoDBResourceMongoDBDatabase{}
	case "azure-native:documentdb/v20230915preview:MongoDBResourceMongoRoleDefinition":
		r = &MongoDBResourceMongoRoleDefinition{}
	case "azure-native:documentdb/v20230915preview:MongoDBResourceMongoUserDefinition":
		r = &MongoDBResourceMongoUserDefinition{}
	case "azure-native:documentdb/v20230915preview:NotebookWorkspace":
		r = &NotebookWorkspace{}
	case "azure-native:documentdb/v20230915preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:documentdb/v20230915preview:Service":
		r = &Service{}
	case "azure-native:documentdb/v20230915preview:SqlResourceSqlContainer":
		r = &SqlResourceSqlContainer{}
	case "azure-native:documentdb/v20230915preview:SqlResourceSqlDatabase":
		r = &SqlResourceSqlDatabase{}
	case "azure-native:documentdb/v20230915preview:SqlResourceSqlRoleAssignment":
		r = &SqlResourceSqlRoleAssignment{}
	case "azure-native:documentdb/v20230915preview:SqlResourceSqlRoleDefinition":
		r = &SqlResourceSqlRoleDefinition{}
	case "azure-native:documentdb/v20230915preview:SqlResourceSqlStoredProcedure":
		r = &SqlResourceSqlStoredProcedure{}
	case "azure-native:documentdb/v20230915preview:SqlResourceSqlTrigger":
		r = &SqlResourceSqlTrigger{}
	case "azure-native:documentdb/v20230915preview:SqlResourceSqlUserDefinedFunction":
		r = &SqlResourceSqlUserDefinedFunction{}
	case "azure-native:documentdb/v20230915preview:TableResourceTable":
		r = &TableResourceTable{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"documentdb/v20230915preview",
		&module{version},
	)
}
