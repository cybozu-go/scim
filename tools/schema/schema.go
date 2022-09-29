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

type Address struct {
	schema.Base
	scimSchemaBase
}

func (Address) Fields() []*schema.Field {
	return []*schema.Field{
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

func (AssociatedGroup) Fields() []*schema.Field {
	return []*schema.Field{
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

func (AuthenticationScheme) Fields() []*schema.Field {
	authschemetyp := schema.Type(`AuthenticationSchemeType`).
		ZeroVal(`InvalidAuthenticationScheme`)
	return []*schema.Field{
		schema.String(`Description`).Required(true),
		schema.String(`DocumentationURI`).
			Unexported(`documentationUri`),
		schema.String(`Name`).Required(true),
		schema.String(`SpecURI`).
			Unexported(`specUri`),
		schema.NewField(`Type`, authschemetyp).
			Unexported(`typ`).
			JSON(`type`),
	}
}

type BulkSupport struct {
	schema.Base
	scimSchemaBase
}

func (BulkSupport) Fields() []*schema.Field {
	return []*schema.Field{
		schema.Int(`MaxOperations`).Required(true),
		schema.Int(`MaxPayloadSize`).Required(true),
		schema.Bool(`Supported`).Required(true),
	}
}

type Email struct {
	schema.Base
	scimSchemaBase
}

func (Email) Fields() []*schema.Field {
	return []*schema.Field{
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

func (EnterpriseManager) Fields() []*schema.Field {
	return []*schema.Field{
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

type Entitlement struct {
	schema.Base
	scimSchemaBase
}

func (Entitlement) Fields() []*schema.Field {
	return []*schema.Field{
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

func (Error) Fields() []*schema.Field {
	errtyp := schema.Type(`ErrorType`).ZeroVal(`""`)
	return []*schema.Field{
		schema.String(`Detail`),
		schema.NewField(`SCIMType`, errtyp).Unexported(`scimType`),
		schema.Int(`Status`),
	}
}

type FilterSupport struct {
	schema.Base
	scimSchemaBase
}

func (FilterSupport) Fields() []*schema.Field {
	return []*schema.Field{
		schema.Int(`MaxResults`),
		schema.Bool(`Supported`),
	}
}

type GenericSupport struct {
	schema.Base
	scimSchemaBase
}

func (GenericSupport) Fields() []*schema.Field {
	return []*schema.Field{
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

func (Group) Fields() []*schema.Field {
	groupmembertyp := schema.Type(`[]*GroupMember`)
	return []*schema.Field{
		schema.String(`DisplayName`),
		schema.String(`ExternalID`).
			Unexported(`externalID`).
			JSON(`externalId`),
		schema.String(`ID`),
		schema.NewField(`Members`, groupmembertyp),
		schema.NewField(`Schemas`, schemastyp),
		schema.NewField(`Meta`, metatyp),
	}
}

type GroupMember struct {
	schema.Base
	scimSchemaBase
}

func (GroupMember) Fields() []*schema.Field {
	return []*schema.Field{
		schema.String(`Value`).Required(true),
		schema.String(`Reference`).Unexported(`ref`).JSON(`$ref`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
	}
}

type IMS struct {
	schema.Base
	scimSchemaBase
}

func (IMS) Fields() []*schema.Field {
	return []*schema.Field{
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

type Names struct {
	schema.Base
	scimSchemaBase
}

func (Names) Fields() []*schema.Field {
	return []*schema.Field{
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

func (PartialResourceRepresentationRequest) Fields() []*schema.Field {
	sslicetyp := schema.Type(`[]string`)
	return []*schema.Field{
		schema.NewField(`Attributes`, sslicetyp),
		schema.NewField(`ExcludedAttributes`, sslicetyp),
	}
}

type PatchOperation struct {
	schema.Base
	scimSchemaBase
}

func (PatchOperation) Fields() []*schema.Field {
	potype := schema.Type(`PatchOperationValue`).
		ImplementsGet(true).
		ImplementsAccept(true).
		UserFacingType(`interface{}`)
	return []*schema.Field{
		schema.String(`ExternalID`).
			Unexported(`externalID`).
			JSON(`externalId`),
		schema.String(`ID`),
		schema.NewField(`Meta`, metatyp),
		schema.NewField(`Op`, schema.Type(`PatchOperationType`).ZeroVal(`""`)),
		schema.String(`Path`),
		schema.NewField(`Value`, potype),
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

type PhoneNumber struct {
	schema.Base
	scimSchemaBase
}

func (PhoneNumber) Fields() []*schema.Field {
	pntype := schema.Type(`PhoneNumberValue`).
		ZeroVal(`""`).
		UserFacingType(`string`).
		ImplementsGet(true).
		ImplementsAccept(true)
	return []*schema.Field{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.NewField(`Value`, pntype),
	}
}

type Photo struct {
	schema.Base
	scimSchemaBase
}

func (Photo) Fields() []*schema.Field {
	return []*schema.Field{
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

func (ResourceType) Fields() []*schema.Field {
	schemaexttyp := schema.Type(`[]*SchemaExtension`)
	return []*schema.Field{
		schema.String(`Description`),
		schema.String(`Endpoint`).Required(true),
		schema.String(`ID`),
		schema.String(`Name`).Required(true),
		schema.String(`Schema`).Required(true),
		schema.NewField(`SchemaExtensions`, schemaexttyp),
		schema.NewField(`Schemas`, schemastyp),
	}
}

type Role struct {
	schema.Base
	scimSchemaBase
}

func (Role) Fields() []*schema.Field {
	return []*schema.Field{
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

func (Schema) Fields() []*schema.Field {
	satyp := schema.Type(`[]*SchemaAttribute`)

	return []*schema.Field{
		schema.NewField("Attributes", satyp).Required(true),
		schema.String("Description"),
		schema.String("ID").Required(true),
		schema.String("Name").Required(true),
	}
}

type SchemaAttribute struct {
	schema.Base
	scimSchemaBase
}

func (SchemaAttribute) Fields() []*schema.Field {
	muttype := schema.Type(`Mutability`).ZeroVal(`MutReadOnly`)
	rettype := schema.Type(`Returned`).ZeroVal(`ReturnedNever`)
	subattrtype := schema.Type(`[]*SchemaAttribute`)
	dttype := schema.Type(`DataType`).ZeroVal(`InvalidDataType`)
	uniqtype := schema.Type(`Uniqueness`).ZeroVal(`UniqNone`)
	return []*schema.Field{
		schema.NewField(`CanonicalValues`, []interface{}(nil)),
		schema.Bool(`CaseExact`),
		schema.String(`Description`),
		schema.Bool(`MultiValued`).Required(true),
		schema.NewField(`Mutability`, muttype),
		schema.String(`Name`),
		schema.NewField(`ReferenceTypes`, []string(nil)),
		schema.Bool(`Required`),
		schema.NewField(`Returned`, rettype),
		schema.NewField(`SubAttributes`, subattrtype),
		schema.NewField(`Type`, dttype).Unexported(`typ`).JSON(`type`).Required(true),
		schema.NewField(`Uniqueness`, uniqtype),
		schema.String("GoAccessorName").IsExtension(true).
			Comment("returns the exported method name to retrieve the particular attribute. For example, attribute that // has the JSON field name `externalId` might return `ExternalID`, `$ref` might return `Reference`, etc."),
	}
}

type SchemaExtension struct {
	schema.Base
	scimSchemaBase
}

func (SchemaExtension) Fields() []*schema.Field {
	return []*schema.Field{
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
		schema.NewField(`AuthenticationSchemes`, authschemestyp).Required(true),
		schema.NewField(`Bulk`, schema.Type(`*BulkSupport`)).Required(true),
		schema.NewField(`ChangePassword`, gensupporttyp).Required(true),
		schema.String(`DocumentationURI`).
			Unexported(`documentationUri`),
		schema.NewField(`ETag`, gensupporttyp).Unexported(`etag`),
		schema.NewField(`Filter`, schema.Type(`*FilterSupport`)).Required(true),
		schema.NewField(`Patch`, gensupporttyp).Required(true).Required(true),
		schema.NewField(`Schemas`, schemastyp),
		schema.NewField(`Sort`, gensupporttyp).Required(true),
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
	grpmembertyp := schema.Type(`[]*AssociatedGroup`)
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
			Unexported(`externalID`).
			JSON(`externalId`),
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
			Unexported(`profileURL`).
			JSON(`profileUrl`),
		schema.NewField(`Roles`, roletyp),
		schema.NewField(`Schemas`, schemastyp),
		schema.String(`Timezone`),
		schema.String(`Title`),
		schema.String(`UserName`).Required(true),
		schema.String(`UserType`),
		schema.NewField(`X509Certificates`, certtyp),
	}
}

type X509Certificate struct {
	schema.Base
	scimSchemaBase
}

func (X509Certificate) Fields() []*schema.Field {
	return []*schema.Field{
		schema.String(`Display`),
		schema.Bool(`Primary`),
		schema.String(`Type`).Unexported(`typ`).JSON(`type`),
		schema.String(`Value`),
	}
}
