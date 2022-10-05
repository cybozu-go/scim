package schema

import (
	"time"

	"github.com/lestrrat-go/sketch/schema"
)

type scimSchemaBase struct{}

func (scimSchemaBase) GetSchemaURI() string { return "" }

var metatyp *schema.TypeSpec
var schemastyp *schema.TypeSpec

func init() {
	metatyp = schema.TypeName(`*Meta`)
	schemastyp = schema.TypeName(`schemas`).
		AcceptValue(true).
		GetValue(true).
		InitializerArgumentStyle(schema.InitializerArgumentAsSlice).
		Element(`string`).
		ApparentType(`[]string`)
}

type Address struct {
	schema.Base
	scimSchemaBase
}

func (Address) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Country`),
		schema.String(`Formatted`),
		schema.String(`Locality`),
		schema.String(`PostalCode`),
		schema.String(`Region`),
		schema.String(`StreetAddress`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
	}
}

type AssociatedGroup struct {
	schema.Base
	scimSchemaBase
}

func (AssociatedGroup) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Display`),
		schema.String(`Reference`).Unexported(`ref`).JSON(`$ref`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.String(`Value`),
	}
}

type AuthenticationScheme struct {
	schema.Base
	scimSchemaBase
}

func (AuthenticationScheme) Fields() []*schema.FieldSpec {
	authschemetyp := schema.TypeName(`AuthenticationSchemeType`).
		ZeroVal(`InvalidAuthenticationScheme`)
	return []*schema.FieldSpec{
		schema.String(`Description`).Required(true),
		schema.String(`DocumentationURI`).
			Unexported(`documentationURI`).
			JSON(`documentationUri`),
		schema.String(`Name`).Required(true),
		schema.String(`SpecURI`).
			Unexported(`specURI`).
			JSON(`specUri`),
		schema.Field(`Type`, authschemetyp).
			Unexported(`typ`).
			JSON(`type`),
	}
}

type BulkSupport struct {
	schema.Base
	scimSchemaBase
}

func (BulkSupport) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.Int(`MaxOperations`).Required(true),
		schema.Int(`MaxPayloadSize`).Required(true),
		schema.Bool(`Supported`).Required(true),
	}
}

type Email struct {
	schema.Base
	scimSchemaBase
}

func (Email) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.String(`Value`),
	}
}

type EnterpriseManager struct {
	schema.Base
	scimSchemaBase
}

func (EnterpriseManager) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`DisplayName`),
		schema.String(`ID`),
		schema.String(`Reference`).
			Unexported(`ref`).
			JSON(`$ref`),
	}
}

type EnterpriseUser struct {
	schema.Base
	scimSchemaBase
}

func (EnterpriseUser) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"
}

func (EnterpriseUser) Fields() []*schema.FieldSpec {
	emtyp := schema.TypeName(`*EnterpriseManager`)
	return []*schema.FieldSpec{
		schema.String(`CostCenter`),
		schema.String(`Department`),
		schema.String(`Division`),
		schema.String(`EmployeeNumber`),
		schema.Field(`Manager`, emtyp),
		schema.String(`Organization`),
		schema.Field(`Schemas`, schemastyp),
	}
}

type Entitlement struct {
	schema.Base
	scimSchemaBase
}

func (Entitlement) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.String(`Value`),
	}
}

type Error struct {
	schema.Base
	scimSchemaBase
}

func (Error) Fields() []*schema.FieldSpec {
	errtyp := schema.TypeName(`ErrorType`).ZeroVal(`""`)
	return []*schema.FieldSpec{
		schema.String(`Detail`),
		schema.Field(`SCIMType`, errtyp).Unexported(`scimType`),
		schema.Int(`Status`),
	}
}

type FilterSupport struct {
	schema.Base
	scimSchemaBase
}

func (FilterSupport) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.Int(`MaxResults`),
		schema.Bool(`Supported`),
	}
}

type GenericSupport struct {
	schema.Base
	scimSchemaBase
}

func (GenericSupport) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.Bool(`Supported`).Required(true),
	}
}

type Group struct {
	schema.Base
	scimSchemaBase
}

func (Group) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:core:2.0:Group"
}

func (Group) Fields() []*schema.FieldSpec {
	groupmembertyp := schema.TypeName(`[]*GroupMember`)
	return []*schema.FieldSpec{
		schema.String(`DisplayName`),
		schema.String(`ExternalID`).
			Unexported(`externalID`).
			JSON(`externalId`),
		schema.String(`ID`),
		schema.Field(`Members`, groupmembertyp),
		schema.Field(`Schemas`, schemastyp),
		schema.Field(`Meta`, metatyp),
	}
}

type GroupMember struct {
	schema.Base
	scimSchemaBase
}

func (GroupMember) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Value`).Required(true),
		schema.String(`Reference`).Unexported(`ref`).JSON(`$ref`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
	}
}

type IMS struct {
	schema.Base
	scimSchemaBase
}

func (IMS) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.String(`Value`),
	}
}

type ListResponse struct {
	schema.Base
	scimSchemaBase
}

func (ListResponse) GenerateSymbol(name string) bool {
	return name != `object.method.UnmarshalJSON` &&
		name != `object.method.decodeExtraField`
}

func (ListResponse) GetSchemaURI() string {
	return "urn:ietf:params:scim:api:messages:2.0:ListResponse"
}

func (ListResponse) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.Int(`ItemsPerPage`),
		schema.Field(`Resources`, []interface{}(nil)),
		schema.Int(`StartIndex`),
		schema.Int(`TotalResults`),
		schema.Field(`Schemas`, schemastyp),
	}
}

type Meta struct {
	schema.Base
	scimSchemaBase
}

func (Meta) Comment() string {
	return "represents the `meta` field included in SCIM responses. See https://datatracker.ietf.org/doc/html/rfc7643#section-3.1 for details"
}

func (Meta) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`ResourceType`),
		schema.String(`Location`),
		schema.String(`Version`),
		schema.Field(`Created`, time.Time{}),
		schema.Field(`LastModified`, time.Time{}),
	}
}

type Names struct {
	schema.Base
	scimSchemaBase
}

func (Names) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`FamilyName`),
		schema.String(`Formatted`),
		schema.String(`GivenName`),
		schema.String(`HonorificPrefix`),
		schema.String(`HonorificSuffix`),
		schema.String(`MiddleName`),
	}
}

type PartialResourceRepresentationRequest struct {
	schema.Base
	scimSchemaBase
}

func (PartialResourceRepresentationRequest) Fields() []*schema.FieldSpec {
	sslicetyp := schema.TypeName(`[]string`)
	return []*schema.FieldSpec{
		schema.Field(`Attributes`, sslicetyp),
		schema.Field(`ExcludedAttributes`, sslicetyp),
	}
}

type PatchOperation struct {
	schema.Base
	scimSchemaBase
}

func (PatchOperation) Fields() []*schema.FieldSpec {
	potype := schema.TypeName(`PatchOperationValue`).
		GetValue(true).
		AcceptValue(true).
		ApparentType(`interface{}`)
	return []*schema.FieldSpec{
		schema.String(`ExternalID`).
			Unexported(`externalID`).
			JSON(`externalId`),
		schema.String(`ID`),
		schema.Field(`Meta`, metatyp),
		schema.Field(`Op`, schema.TypeName(`PatchOperationType`).ZeroVal(`""`)),
		schema.String(`Path`),
		schema.Field(`Value`, potype),
	}
}

type PatchRequest struct {
	schema.Base
	scimSchemaBase
}

func (PatchRequest) GetSchemaURI() string {
	return "urn:ietf:params:scim:api:messages:2.0:PatchOp"
}

func (PatchRequest) Fields() []*schema.FieldSpec {
	patchoptyp := schema.TypeName(`[]*PatchOperation`)
	return []*schema.FieldSpec{
		schema.Field(`Operations`, patchoptyp),
		schema.Field(`Schemas`, schemastyp),
	}
}

type PhoneNumber struct {
	schema.Base
	scimSchemaBase
}

func (PhoneNumber) Fields() []*schema.FieldSpec {
	pntype := schema.TypeName(`PhoneNumberValue`).
		ZeroVal(`""`).
		ApparentType(`string`).
		GetValue(true).
		AcceptValue(true)
	return []*schema.FieldSpec{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.Field(`Value`, pntype),
	}
}

type Photo struct {
	schema.Base
	scimSchemaBase
}

func (Photo) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.String(`Value`),
	}
}

type ResourceType struct {
	schema.Base
	scimSchemaBase
}

func (ResourceType) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:core:2.0:ResourceType"
}

func (ResourceType) Fields() []*schema.FieldSpec {
	schemaexttyp := schema.TypeName(`[]*SchemaExtension`)
	return []*schema.FieldSpec{
		schema.String(`Description`),
		schema.String(`Endpoint`).Required(true),
		schema.String(`ID`),
		schema.String(`Name`).Required(true),
		schema.String(`Schema`).Required(true),
		schema.Field(`SchemaExtensions`, schemaexttyp),
		schema.Field(`Schemas`, schemastyp),
	}
}

type Role struct {
	schema.Base
	scimSchemaBase
}

func (Role) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.String(`Value`),
	}
}

type Schema struct {
	schema.Base
	scimSchemaBase
}

func (Schema) Comment() string {
	return "represents a Schema resource as defined in the SCIM RFC"
}

func (Schema) Fields() []*schema.FieldSpec {
	satyp := schema.TypeName(`[]*SchemaAttribute`)

	return []*schema.FieldSpec{
		schema.Field("Attributes", satyp).Required(true),
		schema.String("Description"),
		schema.String("ID").Required(true),
		schema.String("Name").Required(true),
	}
}

type SchemaAttribute struct {
	schema.Base
	scimSchemaBase
}

func (SchemaAttribute) Fields() []*schema.FieldSpec {
	muttype := schema.TypeName(`Mutability`).ZeroVal(`MutReadOnly`)
	rettype := schema.TypeName(`Returned`).ZeroVal(`ReturnedNever`)
	subattrtype := schema.TypeName(`[]*SchemaAttribute`)
	dttype := schema.TypeName(`DataType`).ZeroVal(`InvalidDataType`)
	uniqtype := schema.TypeName(`Uniqueness`).ZeroVal(`UniqNone`)
	return []*schema.FieldSpec{
		schema.Field(`CanonicalValues`, []interface{}(nil)),
		schema.Bool(`CaseExact`),
		schema.String(`Description`),
		schema.Bool(`MultiValued`).Required(true),
		schema.Field(`Mutability`, muttype),
		schema.String(`Name`),
		schema.Field(`ReferenceTypes`, []string(nil)),
		schema.Bool(`Required`),
		schema.Field(`Returned`, rettype),
		schema.Field(`SubAttributes`, subattrtype),
		schema.Field(`Type`, dttype).Unexported(`typ`).JSON(`type`).Required(true),
		schema.Field(`Uniqueness`, uniqtype),
		schema.String("GoAccessorName").IsExtension(true).
			Comment("returns the exported method name to retrieve the particular attribute. For example, attribute that // has the JSON field name `externalId` might return `ExternalID`, `$ref` might return `Reference`, etc."),
	}
}

type SchemaExtension struct {
	schema.Base
	scimSchemaBase
}

func (SchemaExtension) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Schema`).Required(true),
		schema.Bool(`Required`).Required(true),
	}
}

type SearchRequest struct {
	schema.Base
	scimSchemaBase
}

func (SearchRequest) GetSchemaURI() string {
	return "urn:ietf:params:scim:schemas:core:2.0:SearchRequest"
}

func (SearchRequest) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.Field(`Attributes`, []string(nil)),
		schema.Int(`Count`),
		schema.Field(`ExcludedAttributes`, []string(nil)),
		schema.String(`Filter`),
		schema.String(`Schema`),
		schema.Field(`Schemas`, schemastyp),
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

func (ServiceProviderConfig) Fields() []*schema.FieldSpec {
	authschemestyp := schema.TypeName(`[]*AuthenticationScheme`)
	gensupporttyp := schema.TypeName(`*GenericSupport`)
	return []*schema.FieldSpec{
		schema.Field(`AuthenticationSchemes`, authschemestyp).Required(true),
		schema.Field(`Bulk`, schema.TypeName(`*BulkSupport`)).Required(true),
		schema.Field(`ChangePassword`, gensupporttyp).Required(true),
		schema.String(`DocumentationURI`).
			Unexported(`documentationURI`).
			JSON(`documentationUri`),
		schema.Field(`ETag`, gensupporttyp).Unexported(`etag`),
		schema.Field(`Filter`, schema.TypeName(`*FilterSupport`)).Required(true),
		schema.Field(`Patch`, gensupporttyp).Required(true).Required(true),
		schema.Field(`Schemas`, schemastyp),
		schema.Field(`Sort`, gensupporttyp).Required(true),
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

func (User) Fields() []*schema.FieldSpec {
	addrtyp := schema.TypeName(`[]*Address`)
	entitlementtyp := schema.TypeName(`[]*Entitlement`)
	emailtyp := schema.TypeName(`[]*Email`)
	grpmembertyp := schema.TypeName(`[]*AssociatedGroup`)
	imstyp := schema.TypeName(`[]*IMS`)
	namestyp := schema.TypeName(`*Names`)
	phonenumbertyp := schema.TypeName(`[]*PhoneNumber`)
	phototyp := schema.TypeName(`[]*Photo`)
	roletyp := schema.TypeName(`[]*Role`)
	certtyp := schema.TypeName(`[]*X509Certificate`)
	return []*schema.FieldSpec{
		schema.Bool(`Active`),
		schema.Field(`Addresses`, addrtyp),
		schema.String(`DisplayName`),
		schema.Field(`Emails`, emailtyp),
		schema.Field(`Entitlements`, entitlementtyp),
		schema.String(`ExternalID`).
			Unexported(`externalID`).
			JSON(`externalId`),
		schema.Field(`Groups`, grpmembertyp),
		schema.String(`ID`),
		schema.Field(`IMS`, imstyp).Unexported(`ims`),
		schema.String(`Locale`),
		schema.Field(`Meta`, metatyp),
		schema.Field(`Name`, namestyp),
		schema.String(`NickName`),
		schema.String(`Password`),
		schema.Field(`PhoneNumbers`, phonenumbertyp),
		schema.Field(`Photos`, phototyp),
		schema.String(`PreferredLanguage`),
		schema.String(`ProfileURL`).
			Unexported(`profileURL`).
			JSON(`profileUrl`),
		schema.Field(`Roles`, roletyp),
		schema.Field(`Schemas`, schemastyp),
		schema.String(`Timezone`),
		schema.String(`Title`),
		schema.String(`UserName`).Required(true),
		schema.String(`UserType`),
		schema.Field(`X509Certificates`, certtyp),
	}
}

type X509Certificate struct {
	schema.Base
	scimSchemaBase
}

func (X509Certificate) Fields() []*schema.FieldSpec {
	return []*schema.FieldSpec{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.String(`Value`),
	}
}
