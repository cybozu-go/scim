package schema

import (
	"time"

	"github.com/lestrrat-go/sketch/schema"
)

type scimSchemaBase struct{}

func (scimSchemaBase) GetSchemaURI() string { return "" }

var metatyp *schema.TypeInfo
var schemastyp *schema.TypeInfo

func init() {
	metatyp = schema.Type(`*Meta`)
	schemastyp = schema.Type(`schemas`).
		ImplementsAccept(true).
		ImplementsGet(true).
		InitializerArgumentStyle(schema.InitializerArgumentAsSlice).
		Element(`string`).
		UserFacingType(`[]string`)
}

type Group struct {
	schema.Base
	scimSchemaBase
}

func (Group) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:core:2.0:Group"
}

func (Group) Fields() []*schema.Field {
	groupmembertyp := schema.Type(`[]*GroupMember`)
	return []*schema.Field{
		schema.String(`DisplayName`),
		schema.String(`ExternalID`).
			Unexported(`externalId`),
		schema.String(`ID`),
		schema.NewField(`Members`, groupmembertyp),
		schema.NewField(`Schemas`, schemastyp),
		schema.NewField(`Meta`, metatyp),
	}
}

type EnterpriseUser struct {
	schema.Base
	scimSchemaBase
}

func (EnterpriseUser) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"
}

func (EnterpriseUser) Fields() []*schema.Field {
	emtyp := schema.Type(`*EnterpriseManager`)
	return []*schema.Field{
		schema.String(`CostCenter`),
		schema.String(`Department`),
		schema.String(`Division`),
		schema.String(`EmployeeNumber`),
		schema.NewField(`Manager`, emtyp),
		schema.String(`Organization`),
		schema.NewField(`Schemas`, schemastyp),
	}
}

type ListResponse struct {
	schema.Base
	scimSchemaBase
}

func (ListResponse) GenerateMethod(name string) bool {
	return name != `object.UnmarshalJSON`
}

func (ListResponse) GetSchemaURI() string {
	return "urn:ietf:params:scim:api:messages:2.0:ListResponse"
}

func (ListResponse) Fields() []*schema.Field {
	return []*schema.Field{
		schema.Int(`ItemsPerPage`),
		schema.NewField(`Resources`, []interface{}(nil)),
		schema.Int(`StartIndex`),
		schema.Int(`TotalResults`),
		schema.NewField(`Schemas`, schemastyp),
	}
}

type Meta struct {
	schema.Base
	scimSchemaBase
}

func (Meta) Comment() string {
	return "represents the `meta` field included in SCIM responses. See https://datatracker.ietf.org/doc/html/rfc7643#section-3.1 for details"
}

func (Meta) Fields() []*schema.Field {
	return []*schema.Field{
		schema.String(`ResourceType`),
		schema.String(`Location`),
		schema.String(`Version`),
		schema.NewField(`Created`, time.Time{}),
		schema.NewField(`LastModified`, time.Time{}),
	}
}

type PatchRequest struct {
	schema.Base
	scimSchemaBase
}

func (PatchRequest) GetSchemaURI() string {
	return "urn:ietf:params:scim:api:messages:2.0:PatchOp"
}

func (PatchRequest) Fields() []*schema.Field {
	patchoptyp := schema.Type(`[]*PatchOperation`)
	return []*schema.Field{
		schema.NewField(`Operations`, patchoptyp),
		schema.NewField(`Schemas`, schemastyp),
	}
}

type ResourceType struct {
	schema.Base
	scimSchemaBase
}

func (ResourceType) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:core:2.0:ResourceType"
}

func (ResourceType) Fields() []*schema.Field {
	schemaexttyp := schema.Type(`[]*SchemaExtension`)
	return []*schema.Field{
		schema.String(`Description`),
		schema.String(`Endpoint`),
		schema.String(`ID`),
		schema.String(`Name`),
		schema.String(`Schema`),
		schema.NewField(`SchemaExtension`, schemaexttyp),
		schema.NewField(`Schemas`, schemastyp),
	}
}

type Schema struct {
	schema.Base
	scimSchemaBase
}

func (Schema) Comment() string {
	return "represents a Schema resource as defined in the SCIM RFC"
}

func (Schema) Fields() []*schema.Field {
	satyp := schema.Type(`[]*SchemaAttribute`)

	return []*schema.Field{
		schema.NewField("Attributes", satyp),
		schema.String("Description"),
		schema.String("ID"),
		schema.String("Name"),
	}
}

type User struct {
	schema.Base
	scimSchemaBase
}

func (User) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:core:2.0:User"
}

func (User) Comment() string {
	return "represents a User resource as defined in the SCIM RFC"
}

func (User) Fields() []*schema.Field {
	addrtyp := schema.Type(`[]*Address`)
	entitlementtyp := schema.Type(`[]*Entitlement`)
	emailtyp := schema.Type(`[]*Email`)
	grpmembertyp := schema.Type(`[]*GroupMember`)
	imstyp := schema.Type(`[]*IMS`)
	namestyp := schema.Type(`*Names`)
	phonenumbertyp := schema.Type(`[]*PhoneNumber`)
	phototyp := schema.Type(`[]*Photo`)
	roletyp := schema.Type(`[]*Role`)
	certtyp := schema.Type(`[]*X509Certificate`)
	return []*schema.Field{
		schema.Bool(`Active`),
		schema.NewField(`Addresses`, addrtyp),
		schema.String(`DisplayName`),
		schema.NewField(`Emails`, emailtyp),
		schema.NewField(`Entitlements`, entitlementtyp),
		schema.String(`ExternalID`).
			Unexported(`externalId`),
		schema.NewField(`Groups`, grpmembertyp),
		schema.String(`ID`),
		schema.NewField(`IMS`, imstyp).Unexported(`ims`),
		schema.String(`Locale`),
		schema.NewField(`Meta`, metatyp),
		schema.NewField(`Name`, namestyp),
		schema.String(`NickName`),
		schema.String(`Password`),
		schema.NewField(`PhoneNumbers`, phonenumbertyp),
		schema.NewField(`Photos`, phototyp),
		schema.String(`PreferredLanguage`),
		schema.String(`ProfileURL`).
			Unexported(`profileUrl`),
		schema.NewField(`Roles`, roletyp),
		schema.NewField(`Schemas`, schemastyp),
		schema.String(`Timezone`),
		schema.String(`Title`),
		schema.String(`UserName`),
		schema.String(`UserType`),
		schema.NewField(`X509Certificates`, certtyp),
	}
}

type SearchRequest struct {
	schema.Base
	scimSchemaBase
}

func (SearchRequest) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:core:2.0:SearchRequest"
}

func (SearchRequest) Fields() []*schema.Field {
	return []*schema.Field{
		schema.NewField(`Attributes`, []string(nil)),
		schema.Int(`Count`),
		schema.NewField(`ExcludedAttributes`, []string(nil)),
		schema.String(`Filter`),
		schema.String(`Schema`),
		schema.NewField(`Schemas`, schemastyp),
		schema.String(`SortBy`),
		schema.String(`SortOrder`),
		schema.Int(`StartIndex`),
	}
}

type ServiceProviderConfig struct {
	schema.Base
	scimSchemaBase
}

func (ServiceProviderConfig) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"
}

func (ServiceProviderConfig) Fields() []*schema.Field {
	authschemestyp := schema.Type(`[]*AuthenticationScheme`)
	gensupporttyp := schema.Type(`*GenericSupport`)
	return []*schema.Field{
		schema.NewField(`AuthenticationSchemes`, authschemestyp),
		schema.NewField(`Bulk`, schema.Type(`*BulkSupport`)),
		schema.NewField(`ChangePassword`, gensupporttyp),
		schema.String(`DocumentationURI`).
			Unexported(`documentationUri`),
		schema.NewField(`ETag`, gensupporttyp).Unexported(`etag`),
		schema.NewField(`Filter`, schema.Type(`*FilterSupport`)),
		schema.NewField(`Patch`, gensupporttyp),
		schema.NewField(`Schemas`, schemastyp),
		schema.NewField(`Sort`, gensupporttyp),
	}
}
