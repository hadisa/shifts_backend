package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// An object with an ID
type Node interface {
	IsNode()
	GetID() string
}

// Remove user account.
//
// Requires one of the following permissions: AUTHENTICATED_USER.
type AccountDelete struct {
	Errors []*AccountError `json:"errors"`
	User   *User           `json:"user"`
}

type AccountError struct {
	// Name of a field that caused the error. A value of `null` indicates that the error isn't associated with a particular field.
	Field *string `json:"field"`
	// The error message.
	Message *string `json:"message"`
	// The error code.
	Code AccountErrorCode `json:"code"`
	// A type of address that causes the error.
	AddressType *AddressTypeEnum `json:"addressType"`
}

type AccountInput struct {
	// Given name.
	FirstName *string `json:"firstName"`
	// Family name.
	LastName *string `json:"lastName"`
	Phone    *string `json:"phone"`
	Whatsapp *string `json:"whatsapp"`
	// User language code.
	LanguageCode *LanguageCodeEnum `json:"languageCode"`
	// Billing address of the customer.
	DefaultBillingAddress *AddressInput `json:"defaultBillingAddress"`
	// Shipping address of the customer.
	DefaultShippingAddress *AddressInput `json:"defaultShippingAddress"`
}

// Register a new user.
type AccountRegister struct {
	// Informs whether users need to confirm their email address.
	RequiresConfirmation *bool           `json:"requiresConfirmation"`
	Errors               []*AccountError `json:"errors"`
	User                 *User           `json:"user"`
}

type AccountRegisterInput struct {
	// the unique identities.id form kratos user
	IdentityID string `json:"identityId"`
	// Given name.
	FirstName *string `json:"firstName"`
	// Family name.
	LastName *string `json:"lastName"`
	Phone    *string `json:"phone"`
	Whatsapp *string `json:"whatsapp"`
	// User language code.
	LanguageCode *LanguageCodeEnum `json:"languageCode"`
	// The email address of the user.
	Email string `json:"email"`
	// Password.
	Password string `json:"password"`
	// Base of frontend URL that will be needed to create confirmation URL.
	RedirectURL *string `json:"redirectUrl"`
	// Slug of a channel which will be used to notify users. Optional when only one channel exists.
	Channel *string `json:"channel"`
}

type AccountRequestDeletion struct {
	Errors []*AccountError `json:"errors"`
	User   *User           `json:"user"`
}

// Updates the account of the logged-in user.
//
// Requires one of the following permissions: AUTHENTICATED_USER.
type AccountUpdate struct {
	AccountErrors []*AccountError `json:"accountErrors"`
	Errors        []*AccountError `json:"errors"`
	User          *User           `json:"user"`
}

// Represents user address data.
type Address struct {
	ID             string `json:"id"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	CompanyName    string `json:"companyName"`
	StreetAddress1 string `json:"streetAddress1"`
	StreetAddress2 string `json:"streetAddress2"`
	City           string `json:"city"`
	CityArea       string `json:"cityArea"`
	PostalCode     string `json:"postalCode"`
	// Shop's default country.
	Country     *CountryDisplay `json:"country"`
	CountryArea string          `json:"countryArea"`
	Phone       *string         `json:"phone"`
	// Address is user's default shipping address.
	IsDefaultShippingAddress *bool `json:"isDefaultShippingAddress"`
	// Address is user's default billing address.
	IsDefaultBillingAddress *bool `json:"isDefaultBillingAddress"`
}

func (Address) IsNode()            {}
func (this Address) GetID() string { return this.ID }

type AddressInput struct {
	// Given name.
	FirstName *string `json:"firstName"`
	// Family name.
	LastName *string `json:"lastName"`
	// Company or organization.
	CompanyName *string `json:"companyName"`
	// Address.
	StreetAddress1 *string `json:"streetAddress1"`
	// Address.
	StreetAddress2 *string `json:"streetAddress2"`
	// City.
	City *string `json:"city"`
	// District.
	CityArea *string `json:"cityArea"`
	// Postal code.
	PostalCode *string `json:"postalCode"`
	// Country.
	Country *CountryCode `json:"country"`
	// State or province.
	CountryArea *string `json:"countryArea"`
	// Phone number.
	Phone *string `json:"phone"`

	
}

type CountryDisplay struct {
	// Country code.
	Code string `json:"code"`
	// Country name.
	Country string `json:"country"`
}

// The Relay compliant `PageInfo` type, containing data necessary to paginate this connection.
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor"`
}

// Represents user data.
type User struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	IsStaff   bool    `json:"isStaff"`
	IsActive  bool    `json:"isActive"`
	Phone     *string `json:"phone"`
	Whatsapp  *string `json:"whatsapp"`
	// A note about the customer.
	//
	// Requires one of the following permissions: MANAGE_USERS, MANAGE_STAFF.
	Note   *string `json:"note"`
	Avatar *string `json:"avatar"`
	// User language code.
	LanguageCode LanguageCodeEnum `json:"languageCode"`
	LastLogin    *time.Time       `json:"lastLogin"`
	DateJoined   time.Time        `json:"dateJoined"`
	UpdatedAt    time.Time        `json:"updatedAt"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }

type UserCountableEdge struct {
	// The item at the end of the edge.
	Node *User `json:"node"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

type AccountErrorCode string

const (
	AccountErrorCodeActivateOwnAccount          AccountErrorCode = "ACTIVATE_OWN_ACCOUNT"
	AccountErrorCodeActivateSuperuserAccount    AccountErrorCode = "ACTIVATE_SUPERUSER_ACCOUNT"
	AccountErrorCodeDuplicatedInputItem         AccountErrorCode = "DUPLICATED_INPUT_ITEM"
	AccountErrorCodeDeactivateOwnAccount        AccountErrorCode = "DEACTIVATE_OWN_ACCOUNT"
	AccountErrorCodeDeactivateSuperuserAccount  AccountErrorCode = "DEACTIVATE_SUPERUSER_ACCOUNT"
	AccountErrorCodeDeleteNonStaffUser          AccountErrorCode = "DELETE_NON_STAFF_USER"
	AccountErrorCodeDeleteOwnAccount            AccountErrorCode = "DELETE_OWN_ACCOUNT"
	AccountErrorCodeDeleteStaffAccount          AccountErrorCode = "DELETE_STAFF_ACCOUNT"
	AccountErrorCodeDeleteSuperuserAccount      AccountErrorCode = "DELETE_SUPERUSER_ACCOUNT"
	AccountErrorCodeGraphqlError                AccountErrorCode = "GRAPHQL_ERROR"
	AccountErrorCodeInactive                    AccountErrorCode = "INACTIVE"
	AccountErrorCodeInvalid                     AccountErrorCode = "INVALID"
	AccountErrorCodeInvalidPassword             AccountErrorCode = "INVALID_PASSWORD"
	AccountErrorCodeLeftNotManageablePermission AccountErrorCode = "LEFT_NOT_MANAGEABLE_PERMISSION"
	AccountErrorCodeInvalidCredentials          AccountErrorCode = "INVALID_CREDENTIALS"
	AccountErrorCodeNotFound                    AccountErrorCode = "NOT_FOUND"
	AccountErrorCodeOutOfScopeUser              AccountErrorCode = "OUT_OF_SCOPE_USER"
	AccountErrorCodeOutOfScopeGroup             AccountErrorCode = "OUT_OF_SCOPE_GROUP"
	AccountErrorCodeOutOfScopePermission        AccountErrorCode = "OUT_OF_SCOPE_PERMISSION"
	AccountErrorCodePasswordEntirelyNumeric     AccountErrorCode = "PASSWORD_ENTIRELY_NUMERIC"
	AccountErrorCodePasswordTooCommon           AccountErrorCode = "PASSWORD_TOO_COMMON"
	AccountErrorCodePasswordTooShort            AccountErrorCode = "PASSWORD_TOO_SHORT"
	AccountErrorCodePasswordTooSimilar          AccountErrorCode = "PASSWORD_TOO_SIMILAR"
	AccountErrorCodeRequired                    AccountErrorCode = "REQUIRED"
	AccountErrorCodeUnique                      AccountErrorCode = "UNIQUE"
	AccountErrorCodeJwtSignatureExpired         AccountErrorCode = "JWT_SIGNATURE_EXPIRED"
	AccountErrorCodeJwtInvalidToken             AccountErrorCode = "JWT_INVALID_TOKEN"
	AccountErrorCodeJwtDecodeError              AccountErrorCode = "JWT_DECODE_ERROR"
	AccountErrorCodeJwtMissingToken             AccountErrorCode = "JWT_MISSING_TOKEN"
	AccountErrorCodeJwtInvalidCsrfToken         AccountErrorCode = "JWT_INVALID_CSRF_TOKEN"
	AccountErrorCodeChannelInactive             AccountErrorCode = "CHANNEL_INACTIVE"
	AccountErrorCodeMissingChannelSlug          AccountErrorCode = "MISSING_CHANNEL_SLUG"
	AccountErrorCodeAccountNotConfirmed         AccountErrorCode = "ACCOUNT_NOT_CONFIRMED"
)

var AllAccountErrorCode = []AccountErrorCode{
	AccountErrorCodeActivateOwnAccount,
	AccountErrorCodeActivateSuperuserAccount,
	AccountErrorCodeDuplicatedInputItem,
	AccountErrorCodeDeactivateOwnAccount,
	AccountErrorCodeDeactivateSuperuserAccount,
	AccountErrorCodeDeleteNonStaffUser,
	AccountErrorCodeDeleteOwnAccount,
	AccountErrorCodeDeleteStaffAccount,
	AccountErrorCodeDeleteSuperuserAccount,
	AccountErrorCodeGraphqlError,
	AccountErrorCodeInactive,
	AccountErrorCodeInvalid,
	AccountErrorCodeInvalidPassword,
	AccountErrorCodeLeftNotManageablePermission,
	AccountErrorCodeInvalidCredentials,
	AccountErrorCodeNotFound,
	AccountErrorCodeOutOfScopeUser,
	AccountErrorCodeOutOfScopeGroup,
	AccountErrorCodeOutOfScopePermission,
	AccountErrorCodePasswordEntirelyNumeric,
	AccountErrorCodePasswordTooCommon,
	AccountErrorCodePasswordTooShort,
	AccountErrorCodePasswordTooSimilar,
	AccountErrorCodeRequired,
	AccountErrorCodeUnique,
	AccountErrorCodeJwtSignatureExpired,
	AccountErrorCodeJwtInvalidToken,
	AccountErrorCodeJwtDecodeError,
	AccountErrorCodeJwtMissingToken,
	AccountErrorCodeJwtInvalidCsrfToken,
	AccountErrorCodeChannelInactive,
	AccountErrorCodeMissingChannelSlug,
	AccountErrorCodeAccountNotConfirmed,
}

func (e AccountErrorCode) IsValid() bool {
	switch e {
	case AccountErrorCodeActivateOwnAccount, AccountErrorCodeActivateSuperuserAccount, AccountErrorCodeDuplicatedInputItem, AccountErrorCodeDeactivateOwnAccount, AccountErrorCodeDeactivateSuperuserAccount, AccountErrorCodeDeleteNonStaffUser, AccountErrorCodeDeleteOwnAccount, AccountErrorCodeDeleteStaffAccount, AccountErrorCodeDeleteSuperuserAccount, AccountErrorCodeGraphqlError, AccountErrorCodeInactive, AccountErrorCodeInvalid, AccountErrorCodeInvalidPassword, AccountErrorCodeLeftNotManageablePermission, AccountErrorCodeInvalidCredentials, AccountErrorCodeNotFound, AccountErrorCodeOutOfScopeUser, AccountErrorCodeOutOfScopeGroup, AccountErrorCodeOutOfScopePermission, AccountErrorCodePasswordEntirelyNumeric, AccountErrorCodePasswordTooCommon, AccountErrorCodePasswordTooShort, AccountErrorCodePasswordTooSimilar, AccountErrorCodeRequired, AccountErrorCodeUnique, AccountErrorCodeJwtSignatureExpired, AccountErrorCodeJwtInvalidToken, AccountErrorCodeJwtDecodeError, AccountErrorCodeJwtMissingToken, AccountErrorCodeJwtInvalidCsrfToken, AccountErrorCodeChannelInactive, AccountErrorCodeMissingChannelSlug, AccountErrorCodeAccountNotConfirmed:
		return true
	}
	return false
}

func (e AccountErrorCode) String() string {
	return string(e)
}

func (e *AccountErrorCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountErrorCode", str)
	}
	return nil
}

func (e AccountErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AddressTypeEnum string

const (
	AddressTypeEnumBilling  AddressTypeEnum = "BILLING"
	AddressTypeEnumShipping AddressTypeEnum = "SHIPPING"
)

var AllAddressTypeEnum = []AddressTypeEnum{
	AddressTypeEnumBilling,
	AddressTypeEnumShipping,
}

func (e AddressTypeEnum) IsValid() bool {
	switch e {
	case AddressTypeEnumBilling, AddressTypeEnumShipping:
		return true
	}
	return false
}

func (e AddressTypeEnum) String() string {
	return string(e)
}

func (e *AddressTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddressTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddressTypeEnum", str)
	}
	return nil
}

func (e AddressTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// An enumeration.
type CountryCode string

const (
	CountryCodeAf CountryCode = "AF"
	CountryCodeAx CountryCode = "AX"
	CountryCodeAl CountryCode = "AL"
	CountryCodeDz CountryCode = "DZ"
	CountryCodeAs CountryCode = "AS"
	CountryCodeAd CountryCode = "AD"
	CountryCodeAo CountryCode = "AO"
	CountryCodeAi CountryCode = "AI"
	CountryCodeAq CountryCode = "AQ"
	CountryCodeAg CountryCode = "AG"
	CountryCodeAr CountryCode = "AR"
	CountryCodeAm CountryCode = "AM"
	CountryCodeAw CountryCode = "AW"
	CountryCodeAu CountryCode = "AU"
	CountryCodeAt CountryCode = "AT"
	CountryCodeAz CountryCode = "AZ"
	CountryCodeBs CountryCode = "BS"
	CountryCodeBh CountryCode = "BH"
	CountryCodeBd CountryCode = "BD"
	CountryCodeBb CountryCode = "BB"
	CountryCodeBy CountryCode = "BY"
	CountryCodeBe CountryCode = "BE"
	CountryCodeBz CountryCode = "BZ"
	CountryCodeBj CountryCode = "BJ"
	CountryCodeBm CountryCode = "BM"
	CountryCodeBt CountryCode = "BT"
	CountryCodeBo CountryCode = "BO"
	CountryCodeBq CountryCode = "BQ"
	CountryCodeBa CountryCode = "BA"
	CountryCodeBw CountryCode = "BW"
	CountryCodeBv CountryCode = "BV"
	CountryCodeBr CountryCode = "BR"
	CountryCodeIo CountryCode = "IO"
	CountryCodeBn CountryCode = "BN"
	CountryCodeBg CountryCode = "BG"
	CountryCodeBf CountryCode = "BF"
	CountryCodeBi CountryCode = "BI"
	CountryCodeCv CountryCode = "CV"
	CountryCodeKh CountryCode = "KH"
	CountryCodeCm CountryCode = "CM"
	CountryCodeCa CountryCode = "CA"
	CountryCodeKy CountryCode = "KY"
	CountryCodeCf CountryCode = "CF"
	CountryCodeTd CountryCode = "TD"
	CountryCodeCl CountryCode = "CL"
	CountryCodeCn CountryCode = "CN"
	CountryCodeCx CountryCode = "CX"
	CountryCodeCc CountryCode = "CC"
	CountryCodeCo CountryCode = "CO"
	CountryCodeKm CountryCode = "KM"
	CountryCodeCg CountryCode = "CG"
	CountryCodeCd CountryCode = "CD"
	CountryCodeCk CountryCode = "CK"
	CountryCodeCr CountryCode = "CR"
	CountryCodeCi CountryCode = "CI"
	CountryCodeHr CountryCode = "HR"
	CountryCodeCu CountryCode = "CU"
	CountryCodeCw CountryCode = "CW"
	CountryCodeCy CountryCode = "CY"
	CountryCodeCz CountryCode = "CZ"
	CountryCodeDk CountryCode = "DK"
	CountryCodeDj CountryCode = "DJ"
	CountryCodeDm CountryCode = "DM"
	CountryCodeDo CountryCode = "DO"
	CountryCodeEc CountryCode = "EC"
	CountryCodeEg CountryCode = "EG"
	CountryCodeSv CountryCode = "SV"
	CountryCodeGq CountryCode = "GQ"
	CountryCodeEr CountryCode = "ER"
	CountryCodeEe CountryCode = "EE"
	CountryCodeSz CountryCode = "SZ"
	CountryCodeEt CountryCode = "ET"
	CountryCodeEu CountryCode = "EU"
	CountryCodeFk CountryCode = "FK"
	CountryCodeFo CountryCode = "FO"
	CountryCodeFj CountryCode = "FJ"
	CountryCodeFi CountryCode = "FI"
	CountryCodeFr CountryCode = "FR"
	CountryCodeGf CountryCode = "GF"
	CountryCodePf CountryCode = "PF"
	CountryCodeTf CountryCode = "TF"
	CountryCodeGa CountryCode = "GA"
	CountryCodeGm CountryCode = "GM"
	CountryCodeGe CountryCode = "GE"
	CountryCodeDe CountryCode = "DE"
	CountryCodeGh CountryCode = "GH"
	CountryCodeGi CountryCode = "GI"
	CountryCodeGr CountryCode = "GR"
	CountryCodeGl CountryCode = "GL"
	CountryCodeGd CountryCode = "GD"
	CountryCodeGp CountryCode = "GP"
	CountryCodeGu CountryCode = "GU"
	CountryCodeGt CountryCode = "GT"
	CountryCodeGg CountryCode = "GG"
	CountryCodeGn CountryCode = "GN"
	CountryCodeGw CountryCode = "GW"
	CountryCodeGy CountryCode = "GY"
	CountryCodeHt CountryCode = "HT"
	CountryCodeHm CountryCode = "HM"
	CountryCodeVa CountryCode = "VA"
	CountryCodeHn CountryCode = "HN"
	CountryCodeHk CountryCode = "HK"
	CountryCodeHu CountryCode = "HU"
	CountryCodeIs CountryCode = "IS"
	CountryCodeIn CountryCode = "IN"
	CountryCodeID CountryCode = "ID"
	CountryCodeIr CountryCode = "IR"
	CountryCodeIq CountryCode = "IQ"
	CountryCodeIe CountryCode = "IE"
	CountryCodeIm CountryCode = "IM"
	CountryCodeIl CountryCode = "IL"
	CountryCodeIt CountryCode = "IT"
	CountryCodeJm CountryCode = "JM"
	CountryCodeJp CountryCode = "JP"
	CountryCodeJe CountryCode = "JE"
	CountryCodeJo CountryCode = "JO"
	CountryCodeKz CountryCode = "KZ"
	CountryCodeKe CountryCode = "KE"
	CountryCodeKi CountryCode = "KI"
	CountryCodeKw CountryCode = "KW"
	CountryCodeKg CountryCode = "KG"
	CountryCodeLa CountryCode = "LA"
	CountryCodeLv CountryCode = "LV"
	CountryCodeLb CountryCode = "LB"
	CountryCodeLs CountryCode = "LS"
	CountryCodeLr CountryCode = "LR"
	CountryCodeLy CountryCode = "LY"
	CountryCodeLi CountryCode = "LI"
	CountryCodeLt CountryCode = "LT"
	CountryCodeLu CountryCode = "LU"
	CountryCodeMo CountryCode = "MO"
	CountryCodeMg CountryCode = "MG"
	CountryCodeMw CountryCode = "MW"
	CountryCodeMy CountryCode = "MY"
	CountryCodeMv CountryCode = "MV"
	CountryCodeMl CountryCode = "ML"
	CountryCodeMt CountryCode = "MT"
	CountryCodeMh CountryCode = "MH"
	CountryCodeMq CountryCode = "MQ"
	CountryCodeMr CountryCode = "MR"
	CountryCodeMu CountryCode = "MU"
	CountryCodeYt CountryCode = "YT"
	CountryCodeMx CountryCode = "MX"
	CountryCodeFm CountryCode = "FM"
	CountryCodeMd CountryCode = "MD"
	CountryCodeMc CountryCode = "MC"
	CountryCodeMn CountryCode = "MN"
	CountryCodeMe CountryCode = "ME"
	CountryCodeMs CountryCode = "MS"
	CountryCodeMa CountryCode = "MA"
	CountryCodeMz CountryCode = "MZ"
	CountryCodeMm CountryCode = "MM"
	CountryCodeNa CountryCode = "NA"
	CountryCodeNr CountryCode = "NR"
	CountryCodeNp CountryCode = "NP"
	CountryCodeNl CountryCode = "NL"
	CountryCodeNc CountryCode = "NC"
	CountryCodeNz CountryCode = "NZ"
	CountryCodeNi CountryCode = "NI"
	CountryCodeNe CountryCode = "NE"
	CountryCodeNg CountryCode = "NG"
	CountryCodeNu CountryCode = "NU"
	CountryCodeNf CountryCode = "NF"
	CountryCodeKp CountryCode = "KP"
	CountryCodeMk CountryCode = "MK"
	CountryCodeMp CountryCode = "MP"
	CountryCodeNo CountryCode = "NO"
	CountryCodeOm CountryCode = "OM"
	CountryCodePk CountryCode = "PK"
	CountryCodePw CountryCode = "PW"
	CountryCodePs CountryCode = "PS"
	CountryCodePa CountryCode = "PA"
	CountryCodePg CountryCode = "PG"
	CountryCodePy CountryCode = "PY"
	CountryCodePe CountryCode = "PE"
	CountryCodePh CountryCode = "PH"
	CountryCodePn CountryCode = "PN"
	CountryCodePl CountryCode = "PL"
	CountryCodePt CountryCode = "PT"
	CountryCodePr CountryCode = "PR"
	CountryCodeQa CountryCode = "QA"
	CountryCodeRe CountryCode = "RE"
	CountryCodeRo CountryCode = "RO"
	CountryCodeRu CountryCode = "RU"
	CountryCodeRw CountryCode = "RW"
	CountryCodeBl CountryCode = "BL"
	CountryCodeSh CountryCode = "SH"
	CountryCodeKn CountryCode = "KN"
	CountryCodeLc CountryCode = "LC"
	CountryCodeMf CountryCode = "MF"
	CountryCodePm CountryCode = "PM"
	CountryCodeVc CountryCode = "VC"
	CountryCodeWs CountryCode = "WS"
	CountryCodeSm CountryCode = "SM"
	CountryCodeSt CountryCode = "ST"
	CountryCodeSa CountryCode = "SA"
	CountryCodeSn CountryCode = "SN"
	CountryCodeRs CountryCode = "RS"
	CountryCodeSc CountryCode = "SC"
	CountryCodeSl CountryCode = "SL"
	CountryCodeSg CountryCode = "SG"
	CountryCodeSx CountryCode = "SX"
	CountryCodeSk CountryCode = "SK"
	CountryCodeSi CountryCode = "SI"
	CountryCodeSb CountryCode = "SB"
	CountryCodeSo CountryCode = "SO"
	CountryCodeZa CountryCode = "ZA"
	CountryCodeGs CountryCode = "GS"
	CountryCodeKr CountryCode = "KR"
	CountryCodeSs CountryCode = "SS"
	CountryCodeEs CountryCode = "ES"
	CountryCodeLk CountryCode = "LK"
	CountryCodeSd CountryCode = "SD"
	CountryCodeSr CountryCode = "SR"
	CountryCodeSj CountryCode = "SJ"
	CountryCodeSe CountryCode = "SE"
	CountryCodeCh CountryCode = "CH"
	CountryCodeSy CountryCode = "SY"
	CountryCodeTw CountryCode = "TW"
	CountryCodeTj CountryCode = "TJ"
	CountryCodeTz CountryCode = "TZ"
	CountryCodeTh CountryCode = "TH"
	CountryCodeTl CountryCode = "TL"
	CountryCodeTg CountryCode = "TG"
	CountryCodeTk CountryCode = "TK"
	CountryCodeTo CountryCode = "TO"
	CountryCodeTt CountryCode = "TT"
	CountryCodeTn CountryCode = "TN"
	CountryCodeTr CountryCode = "TR"
	CountryCodeTm CountryCode = "TM"
	CountryCodeTc CountryCode = "TC"
	CountryCodeTv CountryCode = "TV"
	CountryCodeUg CountryCode = "UG"
	CountryCodeUa CountryCode = "UA"
	CountryCodeAe CountryCode = "AE"
	CountryCodeGb CountryCode = "GB"
	CountryCodeUm CountryCode = "UM"
	CountryCodeUs CountryCode = "US"
	CountryCodeUy CountryCode = "UY"
	CountryCodeUz CountryCode = "UZ"
	CountryCodeVu CountryCode = "VU"
	CountryCodeVe CountryCode = "VE"
	CountryCodeVn CountryCode = "VN"
	CountryCodeVg CountryCode = "VG"
	CountryCodeVi CountryCode = "VI"
	CountryCodeWf CountryCode = "WF"
	CountryCodeEh CountryCode = "EH"
	CountryCodeYe CountryCode = "YE"
	CountryCodeZm CountryCode = "ZM"
	CountryCodeZw CountryCode = "ZW"
)

var AllCountryCode = []CountryCode{
	CountryCodeAf,
	CountryCodeAx,
	CountryCodeAl,
	CountryCodeDz,
	CountryCodeAs,
	CountryCodeAd,
	CountryCodeAo,
	CountryCodeAi,
	CountryCodeAq,
	CountryCodeAg,
	CountryCodeAr,
	CountryCodeAm,
	CountryCodeAw,
	CountryCodeAu,
	CountryCodeAt,
	CountryCodeAz,
	CountryCodeBs,
	CountryCodeBh,
	CountryCodeBd,
	CountryCodeBb,
	CountryCodeBy,
	CountryCodeBe,
	CountryCodeBz,
	CountryCodeBj,
	CountryCodeBm,
	CountryCodeBt,
	CountryCodeBo,
	CountryCodeBq,
	CountryCodeBa,
	CountryCodeBw,
	CountryCodeBv,
	CountryCodeBr,
	CountryCodeIo,
	CountryCodeBn,
	CountryCodeBg,
	CountryCodeBf,
	CountryCodeBi,
	CountryCodeCv,
	CountryCodeKh,
	CountryCodeCm,
	CountryCodeCa,
	CountryCodeKy,
	CountryCodeCf,
	CountryCodeTd,
	CountryCodeCl,
	CountryCodeCn,
	CountryCodeCx,
	CountryCodeCc,
	CountryCodeCo,
	CountryCodeKm,
	CountryCodeCg,
	CountryCodeCd,
	CountryCodeCk,
	CountryCodeCr,
	CountryCodeCi,
	CountryCodeHr,
	CountryCodeCu,
	CountryCodeCw,
	CountryCodeCy,
	CountryCodeCz,
	CountryCodeDk,
	CountryCodeDj,
	CountryCodeDm,
	CountryCodeDo,
	CountryCodeEc,
	CountryCodeEg,
	CountryCodeSv,
	CountryCodeGq,
	CountryCodeEr,
	CountryCodeEe,
	CountryCodeSz,
	CountryCodeEt,
	CountryCodeEu,
	CountryCodeFk,
	CountryCodeFo,
	CountryCodeFj,
	CountryCodeFi,
	CountryCodeFr,
	CountryCodeGf,
	CountryCodePf,
	CountryCodeTf,
	CountryCodeGa,
	CountryCodeGm,
	CountryCodeGe,
	CountryCodeDe,
	CountryCodeGh,
	CountryCodeGi,
	CountryCodeGr,
	CountryCodeGl,
	CountryCodeGd,
	CountryCodeGp,
	CountryCodeGu,
	CountryCodeGt,
	CountryCodeGg,
	CountryCodeGn,
	CountryCodeGw,
	CountryCodeGy,
	CountryCodeHt,
	CountryCodeHm,
	CountryCodeVa,
	CountryCodeHn,
	CountryCodeHk,
	CountryCodeHu,
	CountryCodeIs,
	CountryCodeIn,
	CountryCodeID,
	CountryCodeIr,
	CountryCodeIq,
	CountryCodeIe,
	CountryCodeIm,
	CountryCodeIl,
	CountryCodeIt,
	CountryCodeJm,
	CountryCodeJp,
	CountryCodeJe,
	CountryCodeJo,
	CountryCodeKz,
	CountryCodeKe,
	CountryCodeKi,
	CountryCodeKw,
	CountryCodeKg,
	CountryCodeLa,
	CountryCodeLv,
	CountryCodeLb,
	CountryCodeLs,
	CountryCodeLr,
	CountryCodeLy,
	CountryCodeLi,
	CountryCodeLt,
	CountryCodeLu,
	CountryCodeMo,
	CountryCodeMg,
	CountryCodeMw,
	CountryCodeMy,
	CountryCodeMv,
	CountryCodeMl,
	CountryCodeMt,
	CountryCodeMh,
	CountryCodeMq,
	CountryCodeMr,
	CountryCodeMu,
	CountryCodeYt,
	CountryCodeMx,
	CountryCodeFm,
	CountryCodeMd,
	CountryCodeMc,
	CountryCodeMn,
	CountryCodeMe,
	CountryCodeMs,
	CountryCodeMa,
	CountryCodeMz,
	CountryCodeMm,
	CountryCodeNa,
	CountryCodeNr,
	CountryCodeNp,
	CountryCodeNl,
	CountryCodeNc,
	CountryCodeNz,
	CountryCodeNi,
	CountryCodeNe,
	CountryCodeNg,
	CountryCodeNu,
	CountryCodeNf,
	CountryCodeKp,
	CountryCodeMk,
	CountryCodeMp,
	CountryCodeNo,
	CountryCodeOm,
	CountryCodePk,
	CountryCodePw,
	CountryCodePs,
	CountryCodePa,
	CountryCodePg,
	CountryCodePy,
	CountryCodePe,
	CountryCodePh,
	CountryCodePn,
	CountryCodePl,
	CountryCodePt,
	CountryCodePr,
	CountryCodeQa,
	CountryCodeRe,
	CountryCodeRo,
	CountryCodeRu,
	CountryCodeRw,
	CountryCodeBl,
	CountryCodeSh,
	CountryCodeKn,
	CountryCodeLc,
	CountryCodeMf,
	CountryCodePm,
	CountryCodeVc,
	CountryCodeWs,
	CountryCodeSm,
	CountryCodeSt,
	CountryCodeSa,
	CountryCodeSn,
	CountryCodeRs,
	CountryCodeSc,
	CountryCodeSl,
	CountryCodeSg,
	CountryCodeSx,
	CountryCodeSk,
	CountryCodeSi,
	CountryCodeSb,
	CountryCodeSo,
	CountryCodeZa,
	CountryCodeGs,
	CountryCodeKr,
	CountryCodeSs,
	CountryCodeEs,
	CountryCodeLk,
	CountryCodeSd,
	CountryCodeSr,
	CountryCodeSj,
	CountryCodeSe,
	CountryCodeCh,
	CountryCodeSy,
	CountryCodeTw,
	CountryCodeTj,
	CountryCodeTz,
	CountryCodeTh,
	CountryCodeTl,
	CountryCodeTg,
	CountryCodeTk,
	CountryCodeTo,
	CountryCodeTt,
	CountryCodeTn,
	CountryCodeTr,
	CountryCodeTm,
	CountryCodeTc,
	CountryCodeTv,
	CountryCodeUg,
	CountryCodeUa,
	CountryCodeAe,
	CountryCodeGb,
	CountryCodeUm,
	CountryCodeUs,
	CountryCodeUy,
	CountryCodeUz,
	CountryCodeVu,
	CountryCodeVe,
	CountryCodeVn,
	CountryCodeVg,
	CountryCodeVi,
	CountryCodeWf,
	CountryCodeEh,
	CountryCodeYe,
	CountryCodeZm,
	CountryCodeZw,
}

func (e CountryCode) IsValid() bool {
	switch e {
	case CountryCodeAf, CountryCodeAx, CountryCodeAl, CountryCodeDz, CountryCodeAs, CountryCodeAd, CountryCodeAo, CountryCodeAi, CountryCodeAq, CountryCodeAg, CountryCodeAr, CountryCodeAm, CountryCodeAw, CountryCodeAu, CountryCodeAt, CountryCodeAz, CountryCodeBs, CountryCodeBh, CountryCodeBd, CountryCodeBb, CountryCodeBy, CountryCodeBe, CountryCodeBz, CountryCodeBj, CountryCodeBm, CountryCodeBt, CountryCodeBo, CountryCodeBq, CountryCodeBa, CountryCodeBw, CountryCodeBv, CountryCodeBr, CountryCodeIo, CountryCodeBn, CountryCodeBg, CountryCodeBf, CountryCodeBi, CountryCodeCv, CountryCodeKh, CountryCodeCm, CountryCodeCa, CountryCodeKy, CountryCodeCf, CountryCodeTd, CountryCodeCl, CountryCodeCn, CountryCodeCx, CountryCodeCc, CountryCodeCo, CountryCodeKm, CountryCodeCg, CountryCodeCd, CountryCodeCk, CountryCodeCr, CountryCodeCi, CountryCodeHr, CountryCodeCu, CountryCodeCw, CountryCodeCy, CountryCodeCz, CountryCodeDk, CountryCodeDj, CountryCodeDm, CountryCodeDo, CountryCodeEc, CountryCodeEg, CountryCodeSv, CountryCodeGq, CountryCodeEr, CountryCodeEe, CountryCodeSz, CountryCodeEt, CountryCodeEu, CountryCodeFk, CountryCodeFo, CountryCodeFj, CountryCodeFi, CountryCodeFr, CountryCodeGf, CountryCodePf, CountryCodeTf, CountryCodeGa, CountryCodeGm, CountryCodeGe, CountryCodeDe, CountryCodeGh, CountryCodeGi, CountryCodeGr, CountryCodeGl, CountryCodeGd, CountryCodeGp, CountryCodeGu, CountryCodeGt, CountryCodeGg, CountryCodeGn, CountryCodeGw, CountryCodeGy, CountryCodeHt, CountryCodeHm, CountryCodeVa, CountryCodeHn, CountryCodeHk, CountryCodeHu, CountryCodeIs, CountryCodeIn, CountryCodeID, CountryCodeIr, CountryCodeIq, CountryCodeIe, CountryCodeIm, CountryCodeIl, CountryCodeIt, CountryCodeJm, CountryCodeJp, CountryCodeJe, CountryCodeJo, CountryCodeKz, CountryCodeKe, CountryCodeKi, CountryCodeKw, CountryCodeKg, CountryCodeLa, CountryCodeLv, CountryCodeLb, CountryCodeLs, CountryCodeLr, CountryCodeLy, CountryCodeLi, CountryCodeLt, CountryCodeLu, CountryCodeMo, CountryCodeMg, CountryCodeMw, CountryCodeMy, CountryCodeMv, CountryCodeMl, CountryCodeMt, CountryCodeMh, CountryCodeMq, CountryCodeMr, CountryCodeMu, CountryCodeYt, CountryCodeMx, CountryCodeFm, CountryCodeMd, CountryCodeMc, CountryCodeMn, CountryCodeMe, CountryCodeMs, CountryCodeMa, CountryCodeMz, CountryCodeMm, CountryCodeNa, CountryCodeNr, CountryCodeNp, CountryCodeNl, CountryCodeNc, CountryCodeNz, CountryCodeNi, CountryCodeNe, CountryCodeNg, CountryCodeNu, CountryCodeNf, CountryCodeKp, CountryCodeMk, CountryCodeMp, CountryCodeNo, CountryCodeOm, CountryCodePk, CountryCodePw, CountryCodePs, CountryCodePa, CountryCodePg, CountryCodePy, CountryCodePe, CountryCodePh, CountryCodePn, CountryCodePl, CountryCodePt, CountryCodePr, CountryCodeQa, CountryCodeRe, CountryCodeRo, CountryCodeRu, CountryCodeRw, CountryCodeBl, CountryCodeSh, CountryCodeKn, CountryCodeLc, CountryCodeMf, CountryCodePm, CountryCodeVc, CountryCodeWs, CountryCodeSm, CountryCodeSt, CountryCodeSa, CountryCodeSn, CountryCodeRs, CountryCodeSc, CountryCodeSl, CountryCodeSg, CountryCodeSx, CountryCodeSk, CountryCodeSi, CountryCodeSb, CountryCodeSo, CountryCodeZa, CountryCodeGs, CountryCodeKr, CountryCodeSs, CountryCodeEs, CountryCodeLk, CountryCodeSd, CountryCodeSr, CountryCodeSj, CountryCodeSe, CountryCodeCh, CountryCodeSy, CountryCodeTw, CountryCodeTj, CountryCodeTz, CountryCodeTh, CountryCodeTl, CountryCodeTg, CountryCodeTk, CountryCodeTo, CountryCodeTt, CountryCodeTn, CountryCodeTr, CountryCodeTm, CountryCodeTc, CountryCodeTv, CountryCodeUg, CountryCodeUa, CountryCodeAe, CountryCodeGb, CountryCodeUm, CountryCodeUs, CountryCodeUy, CountryCodeUz, CountryCodeVu, CountryCodeVe, CountryCodeVn, CountryCodeVg, CountryCodeVi, CountryCodeWf, CountryCodeEh, CountryCodeYe, CountryCodeZm, CountryCodeZw:
		return true
	}
	return false
}

func (e CountryCode) String() string {
	return string(e)
}

func (e *CountryCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CountryCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CountryCode", str)
	}
	return nil
}

func (e CountryCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LanguageCodeEnum string

const (
	LanguageCodeEnumAf           LanguageCodeEnum = "AF"
	LanguageCodeEnumAfNa         LanguageCodeEnum = "AF_NA"
	LanguageCodeEnumAfZa         LanguageCodeEnum = "AF_ZA"
	LanguageCodeEnumAgq          LanguageCodeEnum = "AGQ"
	LanguageCodeEnumAgqCm        LanguageCodeEnum = "AGQ_CM"
	LanguageCodeEnumAk           LanguageCodeEnum = "AK"
	LanguageCodeEnumAkGh         LanguageCodeEnum = "AK_GH"
	LanguageCodeEnumAm           LanguageCodeEnum = "AM"
	LanguageCodeEnumAmEt         LanguageCodeEnum = "AM_ET"
	LanguageCodeEnumAr           LanguageCodeEnum = "AR"
	LanguageCodeEnumArAe         LanguageCodeEnum = "AR_AE"
	LanguageCodeEnumArBh         LanguageCodeEnum = "AR_BH"
	LanguageCodeEnumArDj         LanguageCodeEnum = "AR_DJ"
	LanguageCodeEnumArDz         LanguageCodeEnum = "AR_DZ"
	LanguageCodeEnumArEg         LanguageCodeEnum = "AR_EG"
	LanguageCodeEnumArEh         LanguageCodeEnum = "AR_EH"
	LanguageCodeEnumArEr         LanguageCodeEnum = "AR_ER"
	LanguageCodeEnumArIl         LanguageCodeEnum = "AR_IL"
	LanguageCodeEnumArIq         LanguageCodeEnum = "AR_IQ"
	LanguageCodeEnumArJo         LanguageCodeEnum = "AR_JO"
	LanguageCodeEnumArKm         LanguageCodeEnum = "AR_KM"
	LanguageCodeEnumArKw         LanguageCodeEnum = "AR_KW"
	LanguageCodeEnumArLb         LanguageCodeEnum = "AR_LB"
	LanguageCodeEnumArLy         LanguageCodeEnum = "AR_LY"
	LanguageCodeEnumArMa         LanguageCodeEnum = "AR_MA"
	LanguageCodeEnumArMr         LanguageCodeEnum = "AR_MR"
	LanguageCodeEnumArOm         LanguageCodeEnum = "AR_OM"
	LanguageCodeEnumArPs         LanguageCodeEnum = "AR_PS"
	LanguageCodeEnumArQa         LanguageCodeEnum = "AR_QA"
	LanguageCodeEnumArSa         LanguageCodeEnum = "AR_SA"
	LanguageCodeEnumArSd         LanguageCodeEnum = "AR_SD"
	LanguageCodeEnumArSo         LanguageCodeEnum = "AR_SO"
	LanguageCodeEnumArSs         LanguageCodeEnum = "AR_SS"
	LanguageCodeEnumArSy         LanguageCodeEnum = "AR_SY"
	LanguageCodeEnumArTd         LanguageCodeEnum = "AR_TD"
	LanguageCodeEnumArTn         LanguageCodeEnum = "AR_TN"
	LanguageCodeEnumArYe         LanguageCodeEnum = "AR_YE"
	LanguageCodeEnumAs           LanguageCodeEnum = "AS"
	LanguageCodeEnumAsIn         LanguageCodeEnum = "AS_IN"
	LanguageCodeEnumAsa          LanguageCodeEnum = "ASA"
	LanguageCodeEnumAsaTz        LanguageCodeEnum = "ASA_TZ"
	LanguageCodeEnumAst          LanguageCodeEnum = "AST"
	LanguageCodeEnumAstEs        LanguageCodeEnum = "AST_ES"
	LanguageCodeEnumAz           LanguageCodeEnum = "AZ"
	LanguageCodeEnumAzCyrl       LanguageCodeEnum = "AZ_CYRL"
	LanguageCodeEnumAzCyrlAz     LanguageCodeEnum = "AZ_CYRL_AZ"
	LanguageCodeEnumAzLatn       LanguageCodeEnum = "AZ_LATN"
	LanguageCodeEnumAzLatnAz     LanguageCodeEnum = "AZ_LATN_AZ"
	LanguageCodeEnumBas          LanguageCodeEnum = "BAS"
	LanguageCodeEnumBasCm        LanguageCodeEnum = "BAS_CM"
	LanguageCodeEnumBe           LanguageCodeEnum = "BE"
	LanguageCodeEnumBeBy         LanguageCodeEnum = "BE_BY"
	LanguageCodeEnumBem          LanguageCodeEnum = "BEM"
	LanguageCodeEnumBemZm        LanguageCodeEnum = "BEM_ZM"
	LanguageCodeEnumBez          LanguageCodeEnum = "BEZ"
	LanguageCodeEnumBezTz        LanguageCodeEnum = "BEZ_TZ"
	LanguageCodeEnumBg           LanguageCodeEnum = "BG"
	LanguageCodeEnumBgBg         LanguageCodeEnum = "BG_BG"
	LanguageCodeEnumBm           LanguageCodeEnum = "BM"
	LanguageCodeEnumBmMl         LanguageCodeEnum = "BM_ML"
	LanguageCodeEnumBn           LanguageCodeEnum = "BN"
	LanguageCodeEnumBnBd         LanguageCodeEnum = "BN_BD"
	LanguageCodeEnumBnIn         LanguageCodeEnum = "BN_IN"
	LanguageCodeEnumBo           LanguageCodeEnum = "BO"
	LanguageCodeEnumBoCn         LanguageCodeEnum = "BO_CN"
	LanguageCodeEnumBoIn         LanguageCodeEnum = "BO_IN"
	LanguageCodeEnumBr           LanguageCodeEnum = "BR"
	LanguageCodeEnumBrFr         LanguageCodeEnum = "BR_FR"
	LanguageCodeEnumBrx          LanguageCodeEnum = "BRX"
	LanguageCodeEnumBrxIn        LanguageCodeEnum = "BRX_IN"
	LanguageCodeEnumBs           LanguageCodeEnum = "BS"
	LanguageCodeEnumBsCyrl       LanguageCodeEnum = "BS_CYRL"
	LanguageCodeEnumBsCyrlBa     LanguageCodeEnum = "BS_CYRL_BA"
	LanguageCodeEnumBsLatn       LanguageCodeEnum = "BS_LATN"
	LanguageCodeEnumBsLatnBa     LanguageCodeEnum = "BS_LATN_BA"
	LanguageCodeEnumCa           LanguageCodeEnum = "CA"
	LanguageCodeEnumCaAd         LanguageCodeEnum = "CA_AD"
	LanguageCodeEnumCaEs         LanguageCodeEnum = "CA_ES"
	LanguageCodeEnumCaEsValencia LanguageCodeEnum = "CA_ES_VALENCIA"
	LanguageCodeEnumCaFr         LanguageCodeEnum = "CA_FR"
	LanguageCodeEnumCaIt         LanguageCodeEnum = "CA_IT"
	LanguageCodeEnumCcp          LanguageCodeEnum = "CCP"
	LanguageCodeEnumCcpBd        LanguageCodeEnum = "CCP_BD"
	LanguageCodeEnumCcpIn        LanguageCodeEnum = "CCP_IN"
	LanguageCodeEnumCe           LanguageCodeEnum = "CE"
	LanguageCodeEnumCeRu         LanguageCodeEnum = "CE_RU"
	LanguageCodeEnumCeb          LanguageCodeEnum = "CEB"
	LanguageCodeEnumCebPh        LanguageCodeEnum = "CEB_PH"
	LanguageCodeEnumCgg          LanguageCodeEnum = "CGG"
	LanguageCodeEnumCggUg        LanguageCodeEnum = "CGG_UG"
	LanguageCodeEnumChr          LanguageCodeEnum = "CHR"
	LanguageCodeEnumChrUs        LanguageCodeEnum = "CHR_US"
	LanguageCodeEnumCkb          LanguageCodeEnum = "CKB"
	LanguageCodeEnumCkbIq        LanguageCodeEnum = "CKB_IQ"
	LanguageCodeEnumCkbIr        LanguageCodeEnum = "CKB_IR"
	LanguageCodeEnumCs           LanguageCodeEnum = "CS"
	LanguageCodeEnumCsCz         LanguageCodeEnum = "CS_CZ"
	LanguageCodeEnumCu           LanguageCodeEnum = "CU"
	LanguageCodeEnumCuRu         LanguageCodeEnum = "CU_RU"
	LanguageCodeEnumCy           LanguageCodeEnum = "CY"
	LanguageCodeEnumCyGb         LanguageCodeEnum = "CY_GB"
	LanguageCodeEnumDa           LanguageCodeEnum = "DA"
	LanguageCodeEnumDaDk         LanguageCodeEnum = "DA_DK"
	LanguageCodeEnumDaGl         LanguageCodeEnum = "DA_GL"
	LanguageCodeEnumDav          LanguageCodeEnum = "DAV"
	LanguageCodeEnumDavKe        LanguageCodeEnum = "DAV_KE"
	LanguageCodeEnumDe           LanguageCodeEnum = "DE"
	LanguageCodeEnumDeAt         LanguageCodeEnum = "DE_AT"
	LanguageCodeEnumDeBe         LanguageCodeEnum = "DE_BE"
	LanguageCodeEnumDeCh         LanguageCodeEnum = "DE_CH"
	LanguageCodeEnumDeDe         LanguageCodeEnum = "DE_DE"
	LanguageCodeEnumDeIt         LanguageCodeEnum = "DE_IT"
	LanguageCodeEnumDeLi         LanguageCodeEnum = "DE_LI"
	LanguageCodeEnumDeLu         LanguageCodeEnum = "DE_LU"
	LanguageCodeEnumDje          LanguageCodeEnum = "DJE"
	LanguageCodeEnumDjeNe        LanguageCodeEnum = "DJE_NE"
	LanguageCodeEnumDsb          LanguageCodeEnum = "DSB"
	LanguageCodeEnumDsbDe        LanguageCodeEnum = "DSB_DE"
	LanguageCodeEnumDua          LanguageCodeEnum = "DUA"
	LanguageCodeEnumDuaCm        LanguageCodeEnum = "DUA_CM"
	LanguageCodeEnumDyo          LanguageCodeEnum = "DYO"
	LanguageCodeEnumDyoSn        LanguageCodeEnum = "DYO_SN"
	LanguageCodeEnumDz           LanguageCodeEnum = "DZ"
	LanguageCodeEnumDzBt         LanguageCodeEnum = "DZ_BT"
	LanguageCodeEnumEbu          LanguageCodeEnum = "EBU"
	LanguageCodeEnumEbuKe        LanguageCodeEnum = "EBU_KE"
	LanguageCodeEnumEe           LanguageCodeEnum = "EE"
	LanguageCodeEnumEeGh         LanguageCodeEnum = "EE_GH"
	LanguageCodeEnumEeTg         LanguageCodeEnum = "EE_TG"
	LanguageCodeEnumEl           LanguageCodeEnum = "EL"
	LanguageCodeEnumElCy         LanguageCodeEnum = "EL_CY"
	LanguageCodeEnumElGr         LanguageCodeEnum = "EL_GR"
	LanguageCodeEnumEn           LanguageCodeEnum = "EN"
	LanguageCodeEnumEnAe         LanguageCodeEnum = "EN_AE"
	LanguageCodeEnumEnAg         LanguageCodeEnum = "EN_AG"
	LanguageCodeEnumEnAi         LanguageCodeEnum = "EN_AI"
	LanguageCodeEnumEnAs         LanguageCodeEnum = "EN_AS"
	LanguageCodeEnumEnAt         LanguageCodeEnum = "EN_AT"
	LanguageCodeEnumEnAu         LanguageCodeEnum = "EN_AU"
	LanguageCodeEnumEnBb         LanguageCodeEnum = "EN_BB"
	LanguageCodeEnumEnBe         LanguageCodeEnum = "EN_BE"
	LanguageCodeEnumEnBi         LanguageCodeEnum = "EN_BI"
	LanguageCodeEnumEnBm         LanguageCodeEnum = "EN_BM"
	LanguageCodeEnumEnBs         LanguageCodeEnum = "EN_BS"
	LanguageCodeEnumEnBw         LanguageCodeEnum = "EN_BW"
	LanguageCodeEnumEnBz         LanguageCodeEnum = "EN_BZ"
	LanguageCodeEnumEnCa         LanguageCodeEnum = "EN_CA"
	LanguageCodeEnumEnCc         LanguageCodeEnum = "EN_CC"
	LanguageCodeEnumEnCh         LanguageCodeEnum = "EN_CH"
	LanguageCodeEnumEnCk         LanguageCodeEnum = "EN_CK"
	LanguageCodeEnumEnCm         LanguageCodeEnum = "EN_CM"
	LanguageCodeEnumEnCx         LanguageCodeEnum = "EN_CX"
	LanguageCodeEnumEnCy         LanguageCodeEnum = "EN_CY"
	LanguageCodeEnumEnDe         LanguageCodeEnum = "EN_DE"
	LanguageCodeEnumEnDg         LanguageCodeEnum = "EN_DG"
	LanguageCodeEnumEnDk         LanguageCodeEnum = "EN_DK"
	LanguageCodeEnumEnDm         LanguageCodeEnum = "EN_DM"
	LanguageCodeEnumEnEr         LanguageCodeEnum = "EN_ER"
	LanguageCodeEnumEnFi         LanguageCodeEnum = "EN_FI"
	LanguageCodeEnumEnFj         LanguageCodeEnum = "EN_FJ"
	LanguageCodeEnumEnFk         LanguageCodeEnum = "EN_FK"
	LanguageCodeEnumEnFm         LanguageCodeEnum = "EN_FM"
	LanguageCodeEnumEnGb         LanguageCodeEnum = "EN_GB"
	LanguageCodeEnumEnGd         LanguageCodeEnum = "EN_GD"
	LanguageCodeEnumEnGg         LanguageCodeEnum = "EN_GG"
	LanguageCodeEnumEnGh         LanguageCodeEnum = "EN_GH"
	LanguageCodeEnumEnGi         LanguageCodeEnum = "EN_GI"
	LanguageCodeEnumEnGm         LanguageCodeEnum = "EN_GM"
	LanguageCodeEnumEnGu         LanguageCodeEnum = "EN_GU"
	LanguageCodeEnumEnGy         LanguageCodeEnum = "EN_GY"
	LanguageCodeEnumEnHk         LanguageCodeEnum = "EN_HK"
	LanguageCodeEnumEnIe         LanguageCodeEnum = "EN_IE"
	LanguageCodeEnumEnIl         LanguageCodeEnum = "EN_IL"
	LanguageCodeEnumEnIm         LanguageCodeEnum = "EN_IM"
	LanguageCodeEnumEnIn         LanguageCodeEnum = "EN_IN"
	LanguageCodeEnumEnIo         LanguageCodeEnum = "EN_IO"
	LanguageCodeEnumEnJe         LanguageCodeEnum = "EN_JE"
	LanguageCodeEnumEnJm         LanguageCodeEnum = "EN_JM"
	LanguageCodeEnumEnKe         LanguageCodeEnum = "EN_KE"
	LanguageCodeEnumEnKi         LanguageCodeEnum = "EN_KI"
	LanguageCodeEnumEnKn         LanguageCodeEnum = "EN_KN"
	LanguageCodeEnumEnKy         LanguageCodeEnum = "EN_KY"
	LanguageCodeEnumEnLc         LanguageCodeEnum = "EN_LC"
	LanguageCodeEnumEnLr         LanguageCodeEnum = "EN_LR"
	LanguageCodeEnumEnLs         LanguageCodeEnum = "EN_LS"
	LanguageCodeEnumEnMg         LanguageCodeEnum = "EN_MG"
	LanguageCodeEnumEnMh         LanguageCodeEnum = "EN_MH"
	LanguageCodeEnumEnMo         LanguageCodeEnum = "EN_MO"
	LanguageCodeEnumEnMp         LanguageCodeEnum = "EN_MP"
	LanguageCodeEnumEnMs         LanguageCodeEnum = "EN_MS"
	LanguageCodeEnumEnMt         LanguageCodeEnum = "EN_MT"
	LanguageCodeEnumEnMu         LanguageCodeEnum = "EN_MU"
	LanguageCodeEnumEnMw         LanguageCodeEnum = "EN_MW"
	LanguageCodeEnumEnMy         LanguageCodeEnum = "EN_MY"
	LanguageCodeEnumEnNa         LanguageCodeEnum = "EN_NA"
	LanguageCodeEnumEnNf         LanguageCodeEnum = "EN_NF"
	LanguageCodeEnumEnNg         LanguageCodeEnum = "EN_NG"
	LanguageCodeEnumEnNl         LanguageCodeEnum = "EN_NL"
	LanguageCodeEnumEnNr         LanguageCodeEnum = "EN_NR"
	LanguageCodeEnumEnNu         LanguageCodeEnum = "EN_NU"
	LanguageCodeEnumEnNz         LanguageCodeEnum = "EN_NZ"
	LanguageCodeEnumEnPg         LanguageCodeEnum = "EN_PG"
	LanguageCodeEnumEnPh         LanguageCodeEnum = "EN_PH"
	LanguageCodeEnumEnPk         LanguageCodeEnum = "EN_PK"
	LanguageCodeEnumEnPn         LanguageCodeEnum = "EN_PN"
	LanguageCodeEnumEnPr         LanguageCodeEnum = "EN_PR"
	LanguageCodeEnumEnPw         LanguageCodeEnum = "EN_PW"
	LanguageCodeEnumEnRw         LanguageCodeEnum = "EN_RW"
	LanguageCodeEnumEnSb         LanguageCodeEnum = "EN_SB"
	LanguageCodeEnumEnSc         LanguageCodeEnum = "EN_SC"
	LanguageCodeEnumEnSd         LanguageCodeEnum = "EN_SD"
	LanguageCodeEnumEnSe         LanguageCodeEnum = "EN_SE"
	LanguageCodeEnumEnSg         LanguageCodeEnum = "EN_SG"
	LanguageCodeEnumEnSh         LanguageCodeEnum = "EN_SH"
	LanguageCodeEnumEnSi         LanguageCodeEnum = "EN_SI"
	LanguageCodeEnumEnSl         LanguageCodeEnum = "EN_SL"
	LanguageCodeEnumEnSs         LanguageCodeEnum = "EN_SS"
	LanguageCodeEnumEnSx         LanguageCodeEnum = "EN_SX"
	LanguageCodeEnumEnSz         LanguageCodeEnum = "EN_SZ"
	LanguageCodeEnumEnTc         LanguageCodeEnum = "EN_TC"
	LanguageCodeEnumEnTk         LanguageCodeEnum = "EN_TK"
	LanguageCodeEnumEnTo         LanguageCodeEnum = "EN_TO"
	LanguageCodeEnumEnTt         LanguageCodeEnum = "EN_TT"
	LanguageCodeEnumEnTv         LanguageCodeEnum = "EN_TV"
	LanguageCodeEnumEnTz         LanguageCodeEnum = "EN_TZ"
	LanguageCodeEnumEnUg         LanguageCodeEnum = "EN_UG"
	LanguageCodeEnumEnUm         LanguageCodeEnum = "EN_UM"
	LanguageCodeEnumEnUs         LanguageCodeEnum = "EN_US"
	LanguageCodeEnumEnVc         LanguageCodeEnum = "EN_VC"
	LanguageCodeEnumEnVg         LanguageCodeEnum = "EN_VG"
	LanguageCodeEnumEnVi         LanguageCodeEnum = "EN_VI"
	LanguageCodeEnumEnVu         LanguageCodeEnum = "EN_VU"
	LanguageCodeEnumEnWs         LanguageCodeEnum = "EN_WS"
	LanguageCodeEnumEnZa         LanguageCodeEnum = "EN_ZA"
	LanguageCodeEnumEnZm         LanguageCodeEnum = "EN_ZM"
	LanguageCodeEnumEnZw         LanguageCodeEnum = "EN_ZW"
	LanguageCodeEnumEo           LanguageCodeEnum = "EO"
	LanguageCodeEnumEs           LanguageCodeEnum = "ES"
	LanguageCodeEnumEsAr         LanguageCodeEnum = "ES_AR"
	LanguageCodeEnumEsBo         LanguageCodeEnum = "ES_BO"
	LanguageCodeEnumEsBr         LanguageCodeEnum = "ES_BR"
	LanguageCodeEnumEsBz         LanguageCodeEnum = "ES_BZ"
	LanguageCodeEnumEsCl         LanguageCodeEnum = "ES_CL"
	LanguageCodeEnumEsCo         LanguageCodeEnum = "ES_CO"
	LanguageCodeEnumEsCr         LanguageCodeEnum = "ES_CR"
	LanguageCodeEnumEsCu         LanguageCodeEnum = "ES_CU"
	LanguageCodeEnumEsDo         LanguageCodeEnum = "ES_DO"
	LanguageCodeEnumEsEa         LanguageCodeEnum = "ES_EA"
	LanguageCodeEnumEsEc         LanguageCodeEnum = "ES_EC"
	LanguageCodeEnumEsEs         LanguageCodeEnum = "ES_ES"
	LanguageCodeEnumEsGq         LanguageCodeEnum = "ES_GQ"
	LanguageCodeEnumEsGt         LanguageCodeEnum = "ES_GT"
	LanguageCodeEnumEsHn         LanguageCodeEnum = "ES_HN"
	LanguageCodeEnumEsIc         LanguageCodeEnum = "ES_IC"
	LanguageCodeEnumEsMx         LanguageCodeEnum = "ES_MX"
	LanguageCodeEnumEsNi         LanguageCodeEnum = "ES_NI"
	LanguageCodeEnumEsPa         LanguageCodeEnum = "ES_PA"
	LanguageCodeEnumEsPe         LanguageCodeEnum = "ES_PE"
	LanguageCodeEnumEsPh         LanguageCodeEnum = "ES_PH"
	LanguageCodeEnumEsPr         LanguageCodeEnum = "ES_PR"
	LanguageCodeEnumEsPy         LanguageCodeEnum = "ES_PY"
	LanguageCodeEnumEsSv         LanguageCodeEnum = "ES_SV"
	LanguageCodeEnumEsUs         LanguageCodeEnum = "ES_US"
	LanguageCodeEnumEsUy         LanguageCodeEnum = "ES_UY"
	LanguageCodeEnumEsVe         LanguageCodeEnum = "ES_VE"
	LanguageCodeEnumEt           LanguageCodeEnum = "ET"
	LanguageCodeEnumEtEe         LanguageCodeEnum = "ET_EE"
	LanguageCodeEnumEu           LanguageCodeEnum = "EU"
	LanguageCodeEnumEuEs         LanguageCodeEnum = "EU_ES"
	LanguageCodeEnumEwo          LanguageCodeEnum = "EWO"
	LanguageCodeEnumEwoCm        LanguageCodeEnum = "EWO_CM"
	LanguageCodeEnumFa           LanguageCodeEnum = "FA"
	LanguageCodeEnumFaAf         LanguageCodeEnum = "FA_AF"
	LanguageCodeEnumFaIr         LanguageCodeEnum = "FA_IR"
	LanguageCodeEnumFf           LanguageCodeEnum = "FF"
	LanguageCodeEnumFfAdlm       LanguageCodeEnum = "FF_ADLM"
	LanguageCodeEnumFfAdlmBf     LanguageCodeEnum = "FF_ADLM_BF"
	LanguageCodeEnumFfAdlmCm     LanguageCodeEnum = "FF_ADLM_CM"
	LanguageCodeEnumFfAdlmGh     LanguageCodeEnum = "FF_ADLM_GH"
	LanguageCodeEnumFfAdlmGm     LanguageCodeEnum = "FF_ADLM_GM"
	LanguageCodeEnumFfAdlmGn     LanguageCodeEnum = "FF_ADLM_GN"
	LanguageCodeEnumFfAdlmGw     LanguageCodeEnum = "FF_ADLM_GW"
	LanguageCodeEnumFfAdlmLr     LanguageCodeEnum = "FF_ADLM_LR"
	LanguageCodeEnumFfAdlmMr     LanguageCodeEnum = "FF_ADLM_MR"
	LanguageCodeEnumFfAdlmNe     LanguageCodeEnum = "FF_ADLM_NE"
	LanguageCodeEnumFfAdlmNg     LanguageCodeEnum = "FF_ADLM_NG"
	LanguageCodeEnumFfAdlmSl     LanguageCodeEnum = "FF_ADLM_SL"
	LanguageCodeEnumFfAdlmSn     LanguageCodeEnum = "FF_ADLM_SN"
	LanguageCodeEnumFfLatn       LanguageCodeEnum = "FF_LATN"
	LanguageCodeEnumFfLatnBf     LanguageCodeEnum = "FF_LATN_BF"
	LanguageCodeEnumFfLatnCm     LanguageCodeEnum = "FF_LATN_CM"
	LanguageCodeEnumFfLatnGh     LanguageCodeEnum = "FF_LATN_GH"
	LanguageCodeEnumFfLatnGm     LanguageCodeEnum = "FF_LATN_GM"
	LanguageCodeEnumFfLatnGn     LanguageCodeEnum = "FF_LATN_GN"
	LanguageCodeEnumFfLatnGw     LanguageCodeEnum = "FF_LATN_GW"
	LanguageCodeEnumFfLatnLr     LanguageCodeEnum = "FF_LATN_LR"
	LanguageCodeEnumFfLatnMr     LanguageCodeEnum = "FF_LATN_MR"
	LanguageCodeEnumFfLatnNe     LanguageCodeEnum = "FF_LATN_NE"
	LanguageCodeEnumFfLatnNg     LanguageCodeEnum = "FF_LATN_NG"
	LanguageCodeEnumFfLatnSl     LanguageCodeEnum = "FF_LATN_SL"
	LanguageCodeEnumFfLatnSn     LanguageCodeEnum = "FF_LATN_SN"
	LanguageCodeEnumFi           LanguageCodeEnum = "FI"
	LanguageCodeEnumFiFi         LanguageCodeEnum = "FI_FI"
	LanguageCodeEnumFil          LanguageCodeEnum = "FIL"
	LanguageCodeEnumFilPh        LanguageCodeEnum = "FIL_PH"
	LanguageCodeEnumFo           LanguageCodeEnum = "FO"
	LanguageCodeEnumFoDk         LanguageCodeEnum = "FO_DK"
	LanguageCodeEnumFoFo         LanguageCodeEnum = "FO_FO"
	LanguageCodeEnumFr           LanguageCodeEnum = "FR"
	LanguageCodeEnumFrBe         LanguageCodeEnum = "FR_BE"
	LanguageCodeEnumFrBf         LanguageCodeEnum = "FR_BF"
	LanguageCodeEnumFrBi         LanguageCodeEnum = "FR_BI"
	LanguageCodeEnumFrBj         LanguageCodeEnum = "FR_BJ"
	LanguageCodeEnumFrBl         LanguageCodeEnum = "FR_BL"
	LanguageCodeEnumFrCa         LanguageCodeEnum = "FR_CA"
	LanguageCodeEnumFrCd         LanguageCodeEnum = "FR_CD"
	LanguageCodeEnumFrCf         LanguageCodeEnum = "FR_CF"
	LanguageCodeEnumFrCg         LanguageCodeEnum = "FR_CG"
	LanguageCodeEnumFrCh         LanguageCodeEnum = "FR_CH"
	LanguageCodeEnumFrCi         LanguageCodeEnum = "FR_CI"
	LanguageCodeEnumFrCm         LanguageCodeEnum = "FR_CM"
	LanguageCodeEnumFrDj         LanguageCodeEnum = "FR_DJ"
	LanguageCodeEnumFrDz         LanguageCodeEnum = "FR_DZ"
	LanguageCodeEnumFrFr         LanguageCodeEnum = "FR_FR"
	LanguageCodeEnumFrGa         LanguageCodeEnum = "FR_GA"
	LanguageCodeEnumFrGf         LanguageCodeEnum = "FR_GF"
	LanguageCodeEnumFrGn         LanguageCodeEnum = "FR_GN"
	LanguageCodeEnumFrGp         LanguageCodeEnum = "FR_GP"
	LanguageCodeEnumFrGq         LanguageCodeEnum = "FR_GQ"
	LanguageCodeEnumFrHt         LanguageCodeEnum = "FR_HT"
	LanguageCodeEnumFrKm         LanguageCodeEnum = "FR_KM"
	LanguageCodeEnumFrLu         LanguageCodeEnum = "FR_LU"
	LanguageCodeEnumFrMa         LanguageCodeEnum = "FR_MA"
	LanguageCodeEnumFrMc         LanguageCodeEnum = "FR_MC"
	LanguageCodeEnumFrMf         LanguageCodeEnum = "FR_MF"
	LanguageCodeEnumFrMg         LanguageCodeEnum = "FR_MG"
	LanguageCodeEnumFrMl         LanguageCodeEnum = "FR_ML"
	LanguageCodeEnumFrMq         LanguageCodeEnum = "FR_MQ"
	LanguageCodeEnumFrMr         LanguageCodeEnum = "FR_MR"
	LanguageCodeEnumFrMu         LanguageCodeEnum = "FR_MU"
	LanguageCodeEnumFrNc         LanguageCodeEnum = "FR_NC"
	LanguageCodeEnumFrNe         LanguageCodeEnum = "FR_NE"
	LanguageCodeEnumFrPf         LanguageCodeEnum = "FR_PF"
	LanguageCodeEnumFrPm         LanguageCodeEnum = "FR_PM"
	LanguageCodeEnumFrRe         LanguageCodeEnum = "FR_RE"
	LanguageCodeEnumFrRw         LanguageCodeEnum = "FR_RW"
	LanguageCodeEnumFrSc         LanguageCodeEnum = "FR_SC"
	LanguageCodeEnumFrSn         LanguageCodeEnum = "FR_SN"
	LanguageCodeEnumFrSy         LanguageCodeEnum = "FR_SY"
	LanguageCodeEnumFrTd         LanguageCodeEnum = "FR_TD"
	LanguageCodeEnumFrTg         LanguageCodeEnum = "FR_TG"
	LanguageCodeEnumFrTn         LanguageCodeEnum = "FR_TN"
	LanguageCodeEnumFrVu         LanguageCodeEnum = "FR_VU"
	LanguageCodeEnumFrWf         LanguageCodeEnum = "FR_WF"
	LanguageCodeEnumFrYt         LanguageCodeEnum = "FR_YT"
	LanguageCodeEnumFur          LanguageCodeEnum = "FUR"
	LanguageCodeEnumFurIt        LanguageCodeEnum = "FUR_IT"
	LanguageCodeEnumFy           LanguageCodeEnum = "FY"
	LanguageCodeEnumFyNl         LanguageCodeEnum = "FY_NL"
	LanguageCodeEnumGa           LanguageCodeEnum = "GA"
	LanguageCodeEnumGaGb         LanguageCodeEnum = "GA_GB"
	LanguageCodeEnumGaIe         LanguageCodeEnum = "GA_IE"
	LanguageCodeEnumGd           LanguageCodeEnum = "GD"
	LanguageCodeEnumGdGb         LanguageCodeEnum = "GD_GB"
	LanguageCodeEnumGl           LanguageCodeEnum = "GL"
	LanguageCodeEnumGlEs         LanguageCodeEnum = "GL_ES"
	LanguageCodeEnumGsw          LanguageCodeEnum = "GSW"
	LanguageCodeEnumGswCh        LanguageCodeEnum = "GSW_CH"
	LanguageCodeEnumGswFr        LanguageCodeEnum = "GSW_FR"
	LanguageCodeEnumGswLi        LanguageCodeEnum = "GSW_LI"
	LanguageCodeEnumGu           LanguageCodeEnum = "GU"
	LanguageCodeEnumGuIn         LanguageCodeEnum = "GU_IN"
	LanguageCodeEnumGuz          LanguageCodeEnum = "GUZ"
	LanguageCodeEnumGuzKe        LanguageCodeEnum = "GUZ_KE"
	LanguageCodeEnumGv           LanguageCodeEnum = "GV"
	LanguageCodeEnumGvIm         LanguageCodeEnum = "GV_IM"
	LanguageCodeEnumHa           LanguageCodeEnum = "HA"
	LanguageCodeEnumHaGh         LanguageCodeEnum = "HA_GH"
	LanguageCodeEnumHaNe         LanguageCodeEnum = "HA_NE"
	LanguageCodeEnumHaNg         LanguageCodeEnum = "HA_NG"
	LanguageCodeEnumHaw          LanguageCodeEnum = "HAW"
	LanguageCodeEnumHawUs        LanguageCodeEnum = "HAW_US"
	LanguageCodeEnumHe           LanguageCodeEnum = "HE"
	LanguageCodeEnumHeIl         LanguageCodeEnum = "HE_IL"
	LanguageCodeEnumHi           LanguageCodeEnum = "HI"
	LanguageCodeEnumHiIn         LanguageCodeEnum = "HI_IN"
	LanguageCodeEnumHr           LanguageCodeEnum = "HR"
	LanguageCodeEnumHrBa         LanguageCodeEnum = "HR_BA"
	LanguageCodeEnumHrHr         LanguageCodeEnum = "HR_HR"
	LanguageCodeEnumHsb          LanguageCodeEnum = "HSB"
	LanguageCodeEnumHsbDe        LanguageCodeEnum = "HSB_DE"
	LanguageCodeEnumHu           LanguageCodeEnum = "HU"
	LanguageCodeEnumHuHu         LanguageCodeEnum = "HU_HU"
	LanguageCodeEnumHy           LanguageCodeEnum = "HY"
	LanguageCodeEnumHyAm         LanguageCodeEnum = "HY_AM"
	LanguageCodeEnumIa           LanguageCodeEnum = "IA"
	LanguageCodeEnumID           LanguageCodeEnum = "ID"
	LanguageCodeEnumIDID         LanguageCodeEnum = "ID_ID"
	LanguageCodeEnumIg           LanguageCodeEnum = "IG"
	LanguageCodeEnumIgNg         LanguageCodeEnum = "IG_NG"
	LanguageCodeEnumIi           LanguageCodeEnum = "II"
	LanguageCodeEnumIiCn         LanguageCodeEnum = "II_CN"
	LanguageCodeEnumIs           LanguageCodeEnum = "IS"
	LanguageCodeEnumIsIs         LanguageCodeEnum = "IS_IS"
	LanguageCodeEnumIt           LanguageCodeEnum = "IT"
	LanguageCodeEnumItCh         LanguageCodeEnum = "IT_CH"
	LanguageCodeEnumItIt         LanguageCodeEnum = "IT_IT"
	LanguageCodeEnumItSm         LanguageCodeEnum = "IT_SM"
	LanguageCodeEnumItVa         LanguageCodeEnum = "IT_VA"
	LanguageCodeEnumJa           LanguageCodeEnum = "JA"
	LanguageCodeEnumJaJp         LanguageCodeEnum = "JA_JP"
	LanguageCodeEnumJgo          LanguageCodeEnum = "JGO"
	LanguageCodeEnumJgoCm        LanguageCodeEnum = "JGO_CM"
	LanguageCodeEnumJmc          LanguageCodeEnum = "JMC"
	LanguageCodeEnumJmcTz        LanguageCodeEnum = "JMC_TZ"
	LanguageCodeEnumJv           LanguageCodeEnum = "JV"
	LanguageCodeEnumJvID         LanguageCodeEnum = "JV_ID"
	LanguageCodeEnumKa           LanguageCodeEnum = "KA"
	LanguageCodeEnumKaGe         LanguageCodeEnum = "KA_GE"
	LanguageCodeEnumKab          LanguageCodeEnum = "KAB"
	LanguageCodeEnumKabDz        LanguageCodeEnum = "KAB_DZ"
	LanguageCodeEnumKam          LanguageCodeEnum = "KAM"
	LanguageCodeEnumKamKe        LanguageCodeEnum = "KAM_KE"
	LanguageCodeEnumKde          LanguageCodeEnum = "KDE"
	LanguageCodeEnumKdeTz        LanguageCodeEnum = "KDE_TZ"
	LanguageCodeEnumKea          LanguageCodeEnum = "KEA"
	LanguageCodeEnumKeaCv        LanguageCodeEnum = "KEA_CV"
	LanguageCodeEnumKhq          LanguageCodeEnum = "KHQ"
	LanguageCodeEnumKhqMl        LanguageCodeEnum = "KHQ_ML"
	LanguageCodeEnumKi           LanguageCodeEnum = "KI"
	LanguageCodeEnumKiKe         LanguageCodeEnum = "KI_KE"
	LanguageCodeEnumKk           LanguageCodeEnum = "KK"
	LanguageCodeEnumKkKz         LanguageCodeEnum = "KK_KZ"
	LanguageCodeEnumKkj          LanguageCodeEnum = "KKJ"
	LanguageCodeEnumKkjCm        LanguageCodeEnum = "KKJ_CM"
	LanguageCodeEnumKl           LanguageCodeEnum = "KL"
	LanguageCodeEnumKlGl         LanguageCodeEnum = "KL_GL"
	LanguageCodeEnumKln          LanguageCodeEnum = "KLN"
	LanguageCodeEnumKlnKe        LanguageCodeEnum = "KLN_KE"
	LanguageCodeEnumKm           LanguageCodeEnum = "KM"
	LanguageCodeEnumKmKh         LanguageCodeEnum = "KM_KH"
	LanguageCodeEnumKn           LanguageCodeEnum = "KN"
	LanguageCodeEnumKnIn         LanguageCodeEnum = "KN_IN"
	LanguageCodeEnumKo           LanguageCodeEnum = "KO"
	LanguageCodeEnumKoKp         LanguageCodeEnum = "KO_KP"
	LanguageCodeEnumKoKr         LanguageCodeEnum = "KO_KR"
	LanguageCodeEnumKok          LanguageCodeEnum = "KOK"
	LanguageCodeEnumKokIn        LanguageCodeEnum = "KOK_IN"
	LanguageCodeEnumKs           LanguageCodeEnum = "KS"
	LanguageCodeEnumKsArab       LanguageCodeEnum = "KS_ARAB"
	LanguageCodeEnumKsArabIn     LanguageCodeEnum = "KS_ARAB_IN"
	LanguageCodeEnumKsb          LanguageCodeEnum = "KSB"
	LanguageCodeEnumKsbTz        LanguageCodeEnum = "KSB_TZ"
	LanguageCodeEnumKsf          LanguageCodeEnum = "KSF"
	LanguageCodeEnumKsfCm        LanguageCodeEnum = "KSF_CM"
	LanguageCodeEnumKsh          LanguageCodeEnum = "KSH"
	LanguageCodeEnumKshDe        LanguageCodeEnum = "KSH_DE"
	LanguageCodeEnumKu           LanguageCodeEnum = "KU"
	LanguageCodeEnumKuTr         LanguageCodeEnum = "KU_TR"
	LanguageCodeEnumKw           LanguageCodeEnum = "KW"
	LanguageCodeEnumKwGb         LanguageCodeEnum = "KW_GB"
	LanguageCodeEnumKy           LanguageCodeEnum = "KY"
	LanguageCodeEnumKyKg         LanguageCodeEnum = "KY_KG"
	LanguageCodeEnumLag          LanguageCodeEnum = "LAG"
	LanguageCodeEnumLagTz        LanguageCodeEnum = "LAG_TZ"
	LanguageCodeEnumLb           LanguageCodeEnum = "LB"
	LanguageCodeEnumLbLu         LanguageCodeEnum = "LB_LU"
	LanguageCodeEnumLg           LanguageCodeEnum = "LG"
	LanguageCodeEnumLgUg         LanguageCodeEnum = "LG_UG"
	LanguageCodeEnumLkt          LanguageCodeEnum = "LKT"
	LanguageCodeEnumLktUs        LanguageCodeEnum = "LKT_US"
	LanguageCodeEnumLn           LanguageCodeEnum = "LN"
	LanguageCodeEnumLnAo         LanguageCodeEnum = "LN_AO"
	LanguageCodeEnumLnCd         LanguageCodeEnum = "LN_CD"
	LanguageCodeEnumLnCf         LanguageCodeEnum = "LN_CF"
	LanguageCodeEnumLnCg         LanguageCodeEnum = "LN_CG"
	LanguageCodeEnumLo           LanguageCodeEnum = "LO"
	LanguageCodeEnumLoLa         LanguageCodeEnum = "LO_LA"
	LanguageCodeEnumLrc          LanguageCodeEnum = "LRC"
	LanguageCodeEnumLrcIq        LanguageCodeEnum = "LRC_IQ"
	LanguageCodeEnumLrcIr        LanguageCodeEnum = "LRC_IR"
	LanguageCodeEnumLt           LanguageCodeEnum = "LT"
	LanguageCodeEnumLtLt         LanguageCodeEnum = "LT_LT"
	LanguageCodeEnumLu           LanguageCodeEnum = "LU"
	LanguageCodeEnumLuCd         LanguageCodeEnum = "LU_CD"
	LanguageCodeEnumLuo          LanguageCodeEnum = "LUO"
	LanguageCodeEnumLuoKe        LanguageCodeEnum = "LUO_KE"
	LanguageCodeEnumLuy          LanguageCodeEnum = "LUY"
	LanguageCodeEnumLuyKe        LanguageCodeEnum = "LUY_KE"
	LanguageCodeEnumLv           LanguageCodeEnum = "LV"
	LanguageCodeEnumLvLv         LanguageCodeEnum = "LV_LV"
	LanguageCodeEnumMai          LanguageCodeEnum = "MAI"
	LanguageCodeEnumMaiIn        LanguageCodeEnum = "MAI_IN"
	LanguageCodeEnumMas          LanguageCodeEnum = "MAS"
	LanguageCodeEnumMasKe        LanguageCodeEnum = "MAS_KE"
	LanguageCodeEnumMasTz        LanguageCodeEnum = "MAS_TZ"
	LanguageCodeEnumMer          LanguageCodeEnum = "MER"
	LanguageCodeEnumMerKe        LanguageCodeEnum = "MER_KE"
	LanguageCodeEnumMfe          LanguageCodeEnum = "MFE"
	LanguageCodeEnumMfeMu        LanguageCodeEnum = "MFE_MU"
	LanguageCodeEnumMg           LanguageCodeEnum = "MG"
	LanguageCodeEnumMgMg         LanguageCodeEnum = "MG_MG"
	LanguageCodeEnumMgh          LanguageCodeEnum = "MGH"
	LanguageCodeEnumMghMz        LanguageCodeEnum = "MGH_MZ"
	LanguageCodeEnumMgo          LanguageCodeEnum = "MGO"
	LanguageCodeEnumMgoCm        LanguageCodeEnum = "MGO_CM"
	LanguageCodeEnumMi           LanguageCodeEnum = "MI"
	LanguageCodeEnumMiNz         LanguageCodeEnum = "MI_NZ"
	LanguageCodeEnumMk           LanguageCodeEnum = "MK"
	LanguageCodeEnumMkMk         LanguageCodeEnum = "MK_MK"
	LanguageCodeEnumMl           LanguageCodeEnum = "ML"
	LanguageCodeEnumMlIn         LanguageCodeEnum = "ML_IN"
	LanguageCodeEnumMn           LanguageCodeEnum = "MN"
	LanguageCodeEnumMnMn         LanguageCodeEnum = "MN_MN"
	LanguageCodeEnumMni          LanguageCodeEnum = "MNI"
	LanguageCodeEnumMniBeng      LanguageCodeEnum = "MNI_BENG"
	LanguageCodeEnumMniBengIn    LanguageCodeEnum = "MNI_BENG_IN"
	LanguageCodeEnumMr           LanguageCodeEnum = "MR"
	LanguageCodeEnumMrIn         LanguageCodeEnum = "MR_IN"
	LanguageCodeEnumMs           LanguageCodeEnum = "MS"
	LanguageCodeEnumMsBn         LanguageCodeEnum = "MS_BN"
	LanguageCodeEnumMsID         LanguageCodeEnum = "MS_ID"
	LanguageCodeEnumMsMy         LanguageCodeEnum = "MS_MY"
	LanguageCodeEnumMsSg         LanguageCodeEnum = "MS_SG"
	LanguageCodeEnumMt           LanguageCodeEnum = "MT"
	LanguageCodeEnumMtMt         LanguageCodeEnum = "MT_MT"
	LanguageCodeEnumMua          LanguageCodeEnum = "MUA"
	LanguageCodeEnumMuaCm        LanguageCodeEnum = "MUA_CM"
	LanguageCodeEnumMy           LanguageCodeEnum = "MY"
	LanguageCodeEnumMyMm         LanguageCodeEnum = "MY_MM"
	LanguageCodeEnumMzn          LanguageCodeEnum = "MZN"
	LanguageCodeEnumMznIr        LanguageCodeEnum = "MZN_IR"
	LanguageCodeEnumNaq          LanguageCodeEnum = "NAQ"
	LanguageCodeEnumNaqNa        LanguageCodeEnum = "NAQ_NA"
	LanguageCodeEnumNb           LanguageCodeEnum = "NB"
	LanguageCodeEnumNbNo         LanguageCodeEnum = "NB_NO"
	LanguageCodeEnumNbSj         LanguageCodeEnum = "NB_SJ"
	LanguageCodeEnumNd           LanguageCodeEnum = "ND"
	LanguageCodeEnumNdZw         LanguageCodeEnum = "ND_ZW"
	LanguageCodeEnumNds          LanguageCodeEnum = "NDS"
	LanguageCodeEnumNdsDe        LanguageCodeEnum = "NDS_DE"
	LanguageCodeEnumNdsNl        LanguageCodeEnum = "NDS_NL"
	LanguageCodeEnumNe           LanguageCodeEnum = "NE"
	LanguageCodeEnumNeIn         LanguageCodeEnum = "NE_IN"
	LanguageCodeEnumNeNp         LanguageCodeEnum = "NE_NP"
	LanguageCodeEnumNl           LanguageCodeEnum = "NL"
	LanguageCodeEnumNlAw         LanguageCodeEnum = "NL_AW"
	LanguageCodeEnumNlBe         LanguageCodeEnum = "NL_BE"
	LanguageCodeEnumNlBq         LanguageCodeEnum = "NL_BQ"
	LanguageCodeEnumNlCw         LanguageCodeEnum = "NL_CW"
	LanguageCodeEnumNlNl         LanguageCodeEnum = "NL_NL"
	LanguageCodeEnumNlSr         LanguageCodeEnum = "NL_SR"
	LanguageCodeEnumNlSx         LanguageCodeEnum = "NL_SX"
	LanguageCodeEnumNmg          LanguageCodeEnum = "NMG"
	LanguageCodeEnumNmgCm        LanguageCodeEnum = "NMG_CM"
	LanguageCodeEnumNn           LanguageCodeEnum = "NN"
	LanguageCodeEnumNnNo         LanguageCodeEnum = "NN_NO"
	LanguageCodeEnumNnh          LanguageCodeEnum = "NNH"
	LanguageCodeEnumNnhCm        LanguageCodeEnum = "NNH_CM"
	LanguageCodeEnumNus          LanguageCodeEnum = "NUS"
	LanguageCodeEnumNusSs        LanguageCodeEnum = "NUS_SS"
	LanguageCodeEnumNyn          LanguageCodeEnum = "NYN"
	LanguageCodeEnumNynUg        LanguageCodeEnum = "NYN_UG"
	LanguageCodeEnumOm           LanguageCodeEnum = "OM"
	LanguageCodeEnumOmEt         LanguageCodeEnum = "OM_ET"
	LanguageCodeEnumOmKe         LanguageCodeEnum = "OM_KE"
	LanguageCodeEnumOr           LanguageCodeEnum = "OR"
	LanguageCodeEnumOrIn         LanguageCodeEnum = "OR_IN"
	LanguageCodeEnumOs           LanguageCodeEnum = "OS"
	LanguageCodeEnumOsGe         LanguageCodeEnum = "OS_GE"
	LanguageCodeEnumOsRu         LanguageCodeEnum = "OS_RU"
	LanguageCodeEnumPa           LanguageCodeEnum = "PA"
	LanguageCodeEnumPaArab       LanguageCodeEnum = "PA_ARAB"
	LanguageCodeEnumPaArabPk     LanguageCodeEnum = "PA_ARAB_PK"
	LanguageCodeEnumPaGuru       LanguageCodeEnum = "PA_GURU"
	LanguageCodeEnumPaGuruIn     LanguageCodeEnum = "PA_GURU_IN"
	LanguageCodeEnumPcm          LanguageCodeEnum = "PCM"
	LanguageCodeEnumPcmNg        LanguageCodeEnum = "PCM_NG"
	LanguageCodeEnumPl           LanguageCodeEnum = "PL"
	LanguageCodeEnumPlPl         LanguageCodeEnum = "PL_PL"
	LanguageCodeEnumPrg          LanguageCodeEnum = "PRG"
	LanguageCodeEnumPs           LanguageCodeEnum = "PS"
	LanguageCodeEnumPsAf         LanguageCodeEnum = "PS_AF"
	LanguageCodeEnumPsPk         LanguageCodeEnum = "PS_PK"
	LanguageCodeEnumPt           LanguageCodeEnum = "PT"
	LanguageCodeEnumPtAo         LanguageCodeEnum = "PT_AO"
	LanguageCodeEnumPtBr         LanguageCodeEnum = "PT_BR"
	LanguageCodeEnumPtCh         LanguageCodeEnum = "PT_CH"
	LanguageCodeEnumPtCv         LanguageCodeEnum = "PT_CV"
	LanguageCodeEnumPtGq         LanguageCodeEnum = "PT_GQ"
	LanguageCodeEnumPtGw         LanguageCodeEnum = "PT_GW"
	LanguageCodeEnumPtLu         LanguageCodeEnum = "PT_LU"
	LanguageCodeEnumPtMo         LanguageCodeEnum = "PT_MO"
	LanguageCodeEnumPtMz         LanguageCodeEnum = "PT_MZ"
	LanguageCodeEnumPtPt         LanguageCodeEnum = "PT_PT"
	LanguageCodeEnumPtSt         LanguageCodeEnum = "PT_ST"
	LanguageCodeEnumPtTl         LanguageCodeEnum = "PT_TL"
	LanguageCodeEnumQu           LanguageCodeEnum = "QU"
	LanguageCodeEnumQuBo         LanguageCodeEnum = "QU_BO"
	LanguageCodeEnumQuEc         LanguageCodeEnum = "QU_EC"
	LanguageCodeEnumQuPe         LanguageCodeEnum = "QU_PE"
	LanguageCodeEnumRm           LanguageCodeEnum = "RM"
	LanguageCodeEnumRmCh         LanguageCodeEnum = "RM_CH"
	LanguageCodeEnumRn           LanguageCodeEnum = "RN"
	LanguageCodeEnumRnBi         LanguageCodeEnum = "RN_BI"
	LanguageCodeEnumRo           LanguageCodeEnum = "RO"
	LanguageCodeEnumRoMd         LanguageCodeEnum = "RO_MD"
	LanguageCodeEnumRoRo         LanguageCodeEnum = "RO_RO"
	LanguageCodeEnumRof          LanguageCodeEnum = "ROF"
	LanguageCodeEnumRofTz        LanguageCodeEnum = "ROF_TZ"
	LanguageCodeEnumRu           LanguageCodeEnum = "RU"
	LanguageCodeEnumRuBy         LanguageCodeEnum = "RU_BY"
	LanguageCodeEnumRuKg         LanguageCodeEnum = "RU_KG"
	LanguageCodeEnumRuKz         LanguageCodeEnum = "RU_KZ"
	LanguageCodeEnumRuMd         LanguageCodeEnum = "RU_MD"
	LanguageCodeEnumRuRu         LanguageCodeEnum = "RU_RU"
	LanguageCodeEnumRuUa         LanguageCodeEnum = "RU_UA"
	LanguageCodeEnumRw           LanguageCodeEnum = "RW"
	LanguageCodeEnumRwRw         LanguageCodeEnum = "RW_RW"
	LanguageCodeEnumRwk          LanguageCodeEnum = "RWK"
	LanguageCodeEnumRwkTz        LanguageCodeEnum = "RWK_TZ"
	LanguageCodeEnumSah          LanguageCodeEnum = "SAH"
	LanguageCodeEnumSahRu        LanguageCodeEnum = "SAH_RU"
	LanguageCodeEnumSaq          LanguageCodeEnum = "SAQ"
	LanguageCodeEnumSaqKe        LanguageCodeEnum = "SAQ_KE"
	LanguageCodeEnumSat          LanguageCodeEnum = "SAT"
	LanguageCodeEnumSatOlck      LanguageCodeEnum = "SAT_OLCK"
	LanguageCodeEnumSatOlckIn    LanguageCodeEnum = "SAT_OLCK_IN"
	LanguageCodeEnumSbp          LanguageCodeEnum = "SBP"
	LanguageCodeEnumSbpTz        LanguageCodeEnum = "SBP_TZ"
	LanguageCodeEnumSd           LanguageCodeEnum = "SD"
	LanguageCodeEnumSdArab       LanguageCodeEnum = "SD_ARAB"
	LanguageCodeEnumSdArabPk     LanguageCodeEnum = "SD_ARAB_PK"
	LanguageCodeEnumSdDeva       LanguageCodeEnum = "SD_DEVA"
	LanguageCodeEnumSdDevaIn     LanguageCodeEnum = "SD_DEVA_IN"
	LanguageCodeEnumSe           LanguageCodeEnum = "SE"
	LanguageCodeEnumSeFi         LanguageCodeEnum = "SE_FI"
	LanguageCodeEnumSeNo         LanguageCodeEnum = "SE_NO"
	LanguageCodeEnumSeSe         LanguageCodeEnum = "SE_SE"
	LanguageCodeEnumSeh          LanguageCodeEnum = "SEH"
	LanguageCodeEnumSehMz        LanguageCodeEnum = "SEH_MZ"
	LanguageCodeEnumSes          LanguageCodeEnum = "SES"
	LanguageCodeEnumSesMl        LanguageCodeEnum = "SES_ML"
	LanguageCodeEnumSg           LanguageCodeEnum = "SG"
	LanguageCodeEnumSgCf         LanguageCodeEnum = "SG_CF"
	LanguageCodeEnumShi          LanguageCodeEnum = "SHI"
	LanguageCodeEnumShiLatn      LanguageCodeEnum = "SHI_LATN"
	LanguageCodeEnumShiLatnMa    LanguageCodeEnum = "SHI_LATN_MA"
	LanguageCodeEnumShiTfng      LanguageCodeEnum = "SHI_TFNG"
	LanguageCodeEnumShiTfngMa    LanguageCodeEnum = "SHI_TFNG_MA"
	LanguageCodeEnumSi           LanguageCodeEnum = "SI"
	LanguageCodeEnumSiLk         LanguageCodeEnum = "SI_LK"
	LanguageCodeEnumSk           LanguageCodeEnum = "SK"
	LanguageCodeEnumSkSk         LanguageCodeEnum = "SK_SK"
	LanguageCodeEnumSl           LanguageCodeEnum = "SL"
	LanguageCodeEnumSlSi         LanguageCodeEnum = "SL_SI"
	LanguageCodeEnumSmn          LanguageCodeEnum = "SMN"
	LanguageCodeEnumSmnFi        LanguageCodeEnum = "SMN_FI"
	LanguageCodeEnumSn           LanguageCodeEnum = "SN"
	LanguageCodeEnumSnZw         LanguageCodeEnum = "SN_ZW"
	LanguageCodeEnumSo           LanguageCodeEnum = "SO"
	LanguageCodeEnumSoDj         LanguageCodeEnum = "SO_DJ"
	LanguageCodeEnumSoEt         LanguageCodeEnum = "SO_ET"
	LanguageCodeEnumSoKe         LanguageCodeEnum = "SO_KE"
	LanguageCodeEnumSoSo         LanguageCodeEnum = "SO_SO"
	LanguageCodeEnumSq           LanguageCodeEnum = "SQ"
	LanguageCodeEnumSqAl         LanguageCodeEnum = "SQ_AL"
	LanguageCodeEnumSqMk         LanguageCodeEnum = "SQ_MK"
	LanguageCodeEnumSqXk         LanguageCodeEnum = "SQ_XK"
	LanguageCodeEnumSr           LanguageCodeEnum = "SR"
	LanguageCodeEnumSrCyrl       LanguageCodeEnum = "SR_CYRL"
	LanguageCodeEnumSrCyrlBa     LanguageCodeEnum = "SR_CYRL_BA"
	LanguageCodeEnumSrCyrlMe     LanguageCodeEnum = "SR_CYRL_ME"
	LanguageCodeEnumSrCyrlRs     LanguageCodeEnum = "SR_CYRL_RS"
	LanguageCodeEnumSrCyrlXk     LanguageCodeEnum = "SR_CYRL_XK"
	LanguageCodeEnumSrLatn       LanguageCodeEnum = "SR_LATN"
	LanguageCodeEnumSrLatnBa     LanguageCodeEnum = "SR_LATN_BA"
	LanguageCodeEnumSrLatnMe     LanguageCodeEnum = "SR_LATN_ME"
	LanguageCodeEnumSrLatnRs     LanguageCodeEnum = "SR_LATN_RS"
	LanguageCodeEnumSrLatnXk     LanguageCodeEnum = "SR_LATN_XK"
	LanguageCodeEnumSu           LanguageCodeEnum = "SU"
	LanguageCodeEnumSuLatn       LanguageCodeEnum = "SU_LATN"
	LanguageCodeEnumSuLatnID     LanguageCodeEnum = "SU_LATN_ID"
	LanguageCodeEnumSv           LanguageCodeEnum = "SV"
	LanguageCodeEnumSvAx         LanguageCodeEnum = "SV_AX"
	LanguageCodeEnumSvFi         LanguageCodeEnum = "SV_FI"
	LanguageCodeEnumSvSe         LanguageCodeEnum = "SV_SE"
	LanguageCodeEnumSw           LanguageCodeEnum = "SW"
	LanguageCodeEnumSwCd         LanguageCodeEnum = "SW_CD"
	LanguageCodeEnumSwKe         LanguageCodeEnum = "SW_KE"
	LanguageCodeEnumSwTz         LanguageCodeEnum = "SW_TZ"
	LanguageCodeEnumSwUg         LanguageCodeEnum = "SW_UG"
	LanguageCodeEnumTa           LanguageCodeEnum = "TA"
	LanguageCodeEnumTaIn         LanguageCodeEnum = "TA_IN"
	LanguageCodeEnumTaLk         LanguageCodeEnum = "TA_LK"
	LanguageCodeEnumTaMy         LanguageCodeEnum = "TA_MY"
	LanguageCodeEnumTaSg         LanguageCodeEnum = "TA_SG"
	LanguageCodeEnumTe           LanguageCodeEnum = "TE"
	LanguageCodeEnumTeIn         LanguageCodeEnum = "TE_IN"
	LanguageCodeEnumTeo          LanguageCodeEnum = "TEO"
	LanguageCodeEnumTeoKe        LanguageCodeEnum = "TEO_KE"
	LanguageCodeEnumTeoUg        LanguageCodeEnum = "TEO_UG"
	LanguageCodeEnumTg           LanguageCodeEnum = "TG"
	LanguageCodeEnumTgTj         LanguageCodeEnum = "TG_TJ"
	LanguageCodeEnumTh           LanguageCodeEnum = "TH"
	LanguageCodeEnumThTh         LanguageCodeEnum = "TH_TH"
	LanguageCodeEnumTi           LanguageCodeEnum = "TI"
	LanguageCodeEnumTiEr         LanguageCodeEnum = "TI_ER"
	LanguageCodeEnumTiEt         LanguageCodeEnum = "TI_ET"
	LanguageCodeEnumTk           LanguageCodeEnum = "TK"
	LanguageCodeEnumTkTm         LanguageCodeEnum = "TK_TM"
	LanguageCodeEnumTo           LanguageCodeEnum = "TO"
	LanguageCodeEnumToTo         LanguageCodeEnum = "TO_TO"
	LanguageCodeEnumTr           LanguageCodeEnum = "TR"
	LanguageCodeEnumTrCy         LanguageCodeEnum = "TR_CY"
	LanguageCodeEnumTrTr         LanguageCodeEnum = "TR_TR"
	LanguageCodeEnumTt           LanguageCodeEnum = "TT"
	LanguageCodeEnumTtRu         LanguageCodeEnum = "TT_RU"
	LanguageCodeEnumTwq          LanguageCodeEnum = "TWQ"
	LanguageCodeEnumTwqNe        LanguageCodeEnum = "TWQ_NE"
	LanguageCodeEnumTzm          LanguageCodeEnum = "TZM"
	LanguageCodeEnumTzmMa        LanguageCodeEnum = "TZM_MA"
	LanguageCodeEnumUg           LanguageCodeEnum = "UG"
	LanguageCodeEnumUgCn         LanguageCodeEnum = "UG_CN"
	LanguageCodeEnumUk           LanguageCodeEnum = "UK"
	LanguageCodeEnumUkUa         LanguageCodeEnum = "UK_UA"
	LanguageCodeEnumUr           LanguageCodeEnum = "UR"
	LanguageCodeEnumUrIn         LanguageCodeEnum = "UR_IN"
	LanguageCodeEnumUrPk         LanguageCodeEnum = "UR_PK"
	LanguageCodeEnumUz           LanguageCodeEnum = "UZ"
	LanguageCodeEnumUzArab       LanguageCodeEnum = "UZ_ARAB"
	LanguageCodeEnumUzArabAf     LanguageCodeEnum = "UZ_ARAB_AF"
	LanguageCodeEnumUzCyrl       LanguageCodeEnum = "UZ_CYRL"
	LanguageCodeEnumUzCyrlUz     LanguageCodeEnum = "UZ_CYRL_UZ"
	LanguageCodeEnumUzLatn       LanguageCodeEnum = "UZ_LATN"
	LanguageCodeEnumUzLatnUz     LanguageCodeEnum = "UZ_LATN_UZ"
	LanguageCodeEnumVai          LanguageCodeEnum = "VAI"
	LanguageCodeEnumVaiLatn      LanguageCodeEnum = "VAI_LATN"
	LanguageCodeEnumVaiLatnLr    LanguageCodeEnum = "VAI_LATN_LR"
	LanguageCodeEnumVaiVaii      LanguageCodeEnum = "VAI_VAII"
	LanguageCodeEnumVaiVaiiLr    LanguageCodeEnum = "VAI_VAII_LR"
	LanguageCodeEnumVi           LanguageCodeEnum = "VI"
	LanguageCodeEnumViVn         LanguageCodeEnum = "VI_VN"
	LanguageCodeEnumVo           LanguageCodeEnum = "VO"
	LanguageCodeEnumVun          LanguageCodeEnum = "VUN"
	LanguageCodeEnumVunTz        LanguageCodeEnum = "VUN_TZ"
	LanguageCodeEnumWae          LanguageCodeEnum = "WAE"
	LanguageCodeEnumWaeCh        LanguageCodeEnum = "WAE_CH"
	LanguageCodeEnumWo           LanguageCodeEnum = "WO"
	LanguageCodeEnumWoSn         LanguageCodeEnum = "WO_SN"
	LanguageCodeEnumXh           LanguageCodeEnum = "XH"
	LanguageCodeEnumXhZa         LanguageCodeEnum = "XH_ZA"
	LanguageCodeEnumXog          LanguageCodeEnum = "XOG"
	LanguageCodeEnumXogUg        LanguageCodeEnum = "XOG_UG"
	LanguageCodeEnumYav          LanguageCodeEnum = "YAV"
	LanguageCodeEnumYavCm        LanguageCodeEnum = "YAV_CM"
	LanguageCodeEnumYi           LanguageCodeEnum = "YI"
	LanguageCodeEnumYo           LanguageCodeEnum = "YO"
	LanguageCodeEnumYoBj         LanguageCodeEnum = "YO_BJ"
	LanguageCodeEnumYoNg         LanguageCodeEnum = "YO_NG"
	LanguageCodeEnumYue          LanguageCodeEnum = "YUE"
	LanguageCodeEnumYueHans      LanguageCodeEnum = "YUE_HANS"
	LanguageCodeEnumYueHansCn    LanguageCodeEnum = "YUE_HANS_CN"
	LanguageCodeEnumYueHant      LanguageCodeEnum = "YUE_HANT"
	LanguageCodeEnumYueHantHk    LanguageCodeEnum = "YUE_HANT_HK"
	LanguageCodeEnumZgh          LanguageCodeEnum = "ZGH"
	LanguageCodeEnumZghMa        LanguageCodeEnum = "ZGH_MA"
	LanguageCodeEnumZh           LanguageCodeEnum = "ZH"
	LanguageCodeEnumZhHans       LanguageCodeEnum = "ZH_HANS"
	LanguageCodeEnumZhHansCn     LanguageCodeEnum = "ZH_HANS_CN"
	LanguageCodeEnumZhHansHk     LanguageCodeEnum = "ZH_HANS_HK"
	LanguageCodeEnumZhHansMo     LanguageCodeEnum = "ZH_HANS_MO"
	LanguageCodeEnumZhHansSg     LanguageCodeEnum = "ZH_HANS_SG"
	LanguageCodeEnumZhHant       LanguageCodeEnum = "ZH_HANT"
	LanguageCodeEnumZhHantHk     LanguageCodeEnum = "ZH_HANT_HK"
	LanguageCodeEnumZhHantMo     LanguageCodeEnum = "ZH_HANT_MO"
	LanguageCodeEnumZhHantTw     LanguageCodeEnum = "ZH_HANT_TW"
	LanguageCodeEnumZu           LanguageCodeEnum = "ZU"
	LanguageCodeEnumZuZa         LanguageCodeEnum = "ZU_ZA"
)

var AllLanguageCodeEnum = []LanguageCodeEnum{
	LanguageCodeEnumAf,
	LanguageCodeEnumAfNa,
	LanguageCodeEnumAfZa,
	LanguageCodeEnumAgq,
	LanguageCodeEnumAgqCm,
	LanguageCodeEnumAk,
	LanguageCodeEnumAkGh,
	LanguageCodeEnumAm,
	LanguageCodeEnumAmEt,
	LanguageCodeEnumAr,
	LanguageCodeEnumArAe,
	LanguageCodeEnumArBh,
	LanguageCodeEnumArDj,
	LanguageCodeEnumArDz,
	LanguageCodeEnumArEg,
	LanguageCodeEnumArEh,
	LanguageCodeEnumArEr,
	LanguageCodeEnumArIl,
	LanguageCodeEnumArIq,
	LanguageCodeEnumArJo,
	LanguageCodeEnumArKm,
	LanguageCodeEnumArKw,
	LanguageCodeEnumArLb,
	LanguageCodeEnumArLy,
	LanguageCodeEnumArMa,
	LanguageCodeEnumArMr,
	LanguageCodeEnumArOm,
	LanguageCodeEnumArPs,
	LanguageCodeEnumArQa,
	LanguageCodeEnumArSa,
	LanguageCodeEnumArSd,
	LanguageCodeEnumArSo,
	LanguageCodeEnumArSs,
	LanguageCodeEnumArSy,
	LanguageCodeEnumArTd,
	LanguageCodeEnumArTn,
	LanguageCodeEnumArYe,
	LanguageCodeEnumAs,
	LanguageCodeEnumAsIn,
	LanguageCodeEnumAsa,
	LanguageCodeEnumAsaTz,
	LanguageCodeEnumAst,
	LanguageCodeEnumAstEs,
	LanguageCodeEnumAz,
	LanguageCodeEnumAzCyrl,
	LanguageCodeEnumAzCyrlAz,
	LanguageCodeEnumAzLatn,
	LanguageCodeEnumAzLatnAz,
	LanguageCodeEnumBas,
	LanguageCodeEnumBasCm,
	LanguageCodeEnumBe,
	LanguageCodeEnumBeBy,
	LanguageCodeEnumBem,
	LanguageCodeEnumBemZm,
	LanguageCodeEnumBez,
	LanguageCodeEnumBezTz,
	LanguageCodeEnumBg,
	LanguageCodeEnumBgBg,
	LanguageCodeEnumBm,
	LanguageCodeEnumBmMl,
	LanguageCodeEnumBn,
	LanguageCodeEnumBnBd,
	LanguageCodeEnumBnIn,
	LanguageCodeEnumBo,
	LanguageCodeEnumBoCn,
	LanguageCodeEnumBoIn,
	LanguageCodeEnumBr,
	LanguageCodeEnumBrFr,
	LanguageCodeEnumBrx,
	LanguageCodeEnumBrxIn,
	LanguageCodeEnumBs,
	LanguageCodeEnumBsCyrl,
	LanguageCodeEnumBsCyrlBa,
	LanguageCodeEnumBsLatn,
	LanguageCodeEnumBsLatnBa,
	LanguageCodeEnumCa,
	LanguageCodeEnumCaAd,
	LanguageCodeEnumCaEs,
	LanguageCodeEnumCaEsValencia,
	LanguageCodeEnumCaFr,
	LanguageCodeEnumCaIt,
	LanguageCodeEnumCcp,
	LanguageCodeEnumCcpBd,
	LanguageCodeEnumCcpIn,
	LanguageCodeEnumCe,
	LanguageCodeEnumCeRu,
	LanguageCodeEnumCeb,
	LanguageCodeEnumCebPh,
	LanguageCodeEnumCgg,
	LanguageCodeEnumCggUg,
	LanguageCodeEnumChr,
	LanguageCodeEnumChrUs,
	LanguageCodeEnumCkb,
	LanguageCodeEnumCkbIq,
	LanguageCodeEnumCkbIr,
	LanguageCodeEnumCs,
	LanguageCodeEnumCsCz,
	LanguageCodeEnumCu,
	LanguageCodeEnumCuRu,
	LanguageCodeEnumCy,
	LanguageCodeEnumCyGb,
	LanguageCodeEnumDa,
	LanguageCodeEnumDaDk,
	LanguageCodeEnumDaGl,
	LanguageCodeEnumDav,
	LanguageCodeEnumDavKe,
	LanguageCodeEnumDe,
	LanguageCodeEnumDeAt,
	LanguageCodeEnumDeBe,
	LanguageCodeEnumDeCh,
	LanguageCodeEnumDeDe,
	LanguageCodeEnumDeIt,
	LanguageCodeEnumDeLi,
	LanguageCodeEnumDeLu,
	LanguageCodeEnumDje,
	LanguageCodeEnumDjeNe,
	LanguageCodeEnumDsb,
	LanguageCodeEnumDsbDe,
	LanguageCodeEnumDua,
	LanguageCodeEnumDuaCm,
	LanguageCodeEnumDyo,
	LanguageCodeEnumDyoSn,
	LanguageCodeEnumDz,
	LanguageCodeEnumDzBt,
	LanguageCodeEnumEbu,
	LanguageCodeEnumEbuKe,
	LanguageCodeEnumEe,
	LanguageCodeEnumEeGh,
	LanguageCodeEnumEeTg,
	LanguageCodeEnumEl,
	LanguageCodeEnumElCy,
	LanguageCodeEnumElGr,
	LanguageCodeEnumEn,
	LanguageCodeEnumEnAe,
	LanguageCodeEnumEnAg,
	LanguageCodeEnumEnAi,
	LanguageCodeEnumEnAs,
	LanguageCodeEnumEnAt,
	LanguageCodeEnumEnAu,
	LanguageCodeEnumEnBb,
	LanguageCodeEnumEnBe,
	LanguageCodeEnumEnBi,
	LanguageCodeEnumEnBm,
	LanguageCodeEnumEnBs,
	LanguageCodeEnumEnBw,
	LanguageCodeEnumEnBz,
	LanguageCodeEnumEnCa,
	LanguageCodeEnumEnCc,
	LanguageCodeEnumEnCh,
	LanguageCodeEnumEnCk,
	LanguageCodeEnumEnCm,
	LanguageCodeEnumEnCx,
	LanguageCodeEnumEnCy,
	LanguageCodeEnumEnDe,
	LanguageCodeEnumEnDg,
	LanguageCodeEnumEnDk,
	LanguageCodeEnumEnDm,
	LanguageCodeEnumEnEr,
	LanguageCodeEnumEnFi,
	LanguageCodeEnumEnFj,
	LanguageCodeEnumEnFk,
	LanguageCodeEnumEnFm,
	LanguageCodeEnumEnGb,
	LanguageCodeEnumEnGd,
	LanguageCodeEnumEnGg,
	LanguageCodeEnumEnGh,
	LanguageCodeEnumEnGi,
	LanguageCodeEnumEnGm,
	LanguageCodeEnumEnGu,
	LanguageCodeEnumEnGy,
	LanguageCodeEnumEnHk,
	LanguageCodeEnumEnIe,
	LanguageCodeEnumEnIl,
	LanguageCodeEnumEnIm,
	LanguageCodeEnumEnIn,
	LanguageCodeEnumEnIo,
	LanguageCodeEnumEnJe,
	LanguageCodeEnumEnJm,
	LanguageCodeEnumEnKe,
	LanguageCodeEnumEnKi,
	LanguageCodeEnumEnKn,
	LanguageCodeEnumEnKy,
	LanguageCodeEnumEnLc,
	LanguageCodeEnumEnLr,
	LanguageCodeEnumEnLs,
	LanguageCodeEnumEnMg,
	LanguageCodeEnumEnMh,
	LanguageCodeEnumEnMo,
	LanguageCodeEnumEnMp,
	LanguageCodeEnumEnMs,
	LanguageCodeEnumEnMt,
	LanguageCodeEnumEnMu,
	LanguageCodeEnumEnMw,
	LanguageCodeEnumEnMy,
	LanguageCodeEnumEnNa,
	LanguageCodeEnumEnNf,
	LanguageCodeEnumEnNg,
	LanguageCodeEnumEnNl,
	LanguageCodeEnumEnNr,
	LanguageCodeEnumEnNu,
	LanguageCodeEnumEnNz,
	LanguageCodeEnumEnPg,
	LanguageCodeEnumEnPh,
	LanguageCodeEnumEnPk,
	LanguageCodeEnumEnPn,
	LanguageCodeEnumEnPr,
	LanguageCodeEnumEnPw,
	LanguageCodeEnumEnRw,
	LanguageCodeEnumEnSb,
	LanguageCodeEnumEnSc,
	LanguageCodeEnumEnSd,
	LanguageCodeEnumEnSe,
	LanguageCodeEnumEnSg,
	LanguageCodeEnumEnSh,
	LanguageCodeEnumEnSi,
	LanguageCodeEnumEnSl,
	LanguageCodeEnumEnSs,
	LanguageCodeEnumEnSx,
	LanguageCodeEnumEnSz,
	LanguageCodeEnumEnTc,
	LanguageCodeEnumEnTk,
	LanguageCodeEnumEnTo,
	LanguageCodeEnumEnTt,
	LanguageCodeEnumEnTv,
	LanguageCodeEnumEnTz,
	LanguageCodeEnumEnUg,
	LanguageCodeEnumEnUm,
	LanguageCodeEnumEnUs,
	LanguageCodeEnumEnVc,
	LanguageCodeEnumEnVg,
	LanguageCodeEnumEnVi,
	LanguageCodeEnumEnVu,
	LanguageCodeEnumEnWs,
	LanguageCodeEnumEnZa,
	LanguageCodeEnumEnZm,
	LanguageCodeEnumEnZw,
	LanguageCodeEnumEo,
	LanguageCodeEnumEs,
	LanguageCodeEnumEsAr,
	LanguageCodeEnumEsBo,
	LanguageCodeEnumEsBr,
	LanguageCodeEnumEsBz,
	LanguageCodeEnumEsCl,
	LanguageCodeEnumEsCo,
	LanguageCodeEnumEsCr,
	LanguageCodeEnumEsCu,
	LanguageCodeEnumEsDo,
	LanguageCodeEnumEsEa,
	LanguageCodeEnumEsEc,
	LanguageCodeEnumEsEs,
	LanguageCodeEnumEsGq,
	LanguageCodeEnumEsGt,
	LanguageCodeEnumEsHn,
	LanguageCodeEnumEsIc,
	LanguageCodeEnumEsMx,
	LanguageCodeEnumEsNi,
	LanguageCodeEnumEsPa,
	LanguageCodeEnumEsPe,
	LanguageCodeEnumEsPh,
	LanguageCodeEnumEsPr,
	LanguageCodeEnumEsPy,
	LanguageCodeEnumEsSv,
	LanguageCodeEnumEsUs,
	LanguageCodeEnumEsUy,
	LanguageCodeEnumEsVe,
	LanguageCodeEnumEt,
	LanguageCodeEnumEtEe,
	LanguageCodeEnumEu,
	LanguageCodeEnumEuEs,
	LanguageCodeEnumEwo,
	LanguageCodeEnumEwoCm,
	LanguageCodeEnumFa,
	LanguageCodeEnumFaAf,
	LanguageCodeEnumFaIr,
	LanguageCodeEnumFf,
	LanguageCodeEnumFfAdlm,
	LanguageCodeEnumFfAdlmBf,
	LanguageCodeEnumFfAdlmCm,
	LanguageCodeEnumFfAdlmGh,
	LanguageCodeEnumFfAdlmGm,
	LanguageCodeEnumFfAdlmGn,
	LanguageCodeEnumFfAdlmGw,
	LanguageCodeEnumFfAdlmLr,
	LanguageCodeEnumFfAdlmMr,
	LanguageCodeEnumFfAdlmNe,
	LanguageCodeEnumFfAdlmNg,
	LanguageCodeEnumFfAdlmSl,
	LanguageCodeEnumFfAdlmSn,
	LanguageCodeEnumFfLatn,
	LanguageCodeEnumFfLatnBf,
	LanguageCodeEnumFfLatnCm,
	LanguageCodeEnumFfLatnGh,
	LanguageCodeEnumFfLatnGm,
	LanguageCodeEnumFfLatnGn,
	LanguageCodeEnumFfLatnGw,
	LanguageCodeEnumFfLatnLr,
	LanguageCodeEnumFfLatnMr,
	LanguageCodeEnumFfLatnNe,
	LanguageCodeEnumFfLatnNg,
	LanguageCodeEnumFfLatnSl,
	LanguageCodeEnumFfLatnSn,
	LanguageCodeEnumFi,
	LanguageCodeEnumFiFi,
	LanguageCodeEnumFil,
	LanguageCodeEnumFilPh,
	LanguageCodeEnumFo,
	LanguageCodeEnumFoDk,
	LanguageCodeEnumFoFo,
	LanguageCodeEnumFr,
	LanguageCodeEnumFrBe,
	LanguageCodeEnumFrBf,
	LanguageCodeEnumFrBi,
	LanguageCodeEnumFrBj,
	LanguageCodeEnumFrBl,
	LanguageCodeEnumFrCa,
	LanguageCodeEnumFrCd,
	LanguageCodeEnumFrCf,
	LanguageCodeEnumFrCg,
	LanguageCodeEnumFrCh,
	LanguageCodeEnumFrCi,
	LanguageCodeEnumFrCm,
	LanguageCodeEnumFrDj,
	LanguageCodeEnumFrDz,
	LanguageCodeEnumFrFr,
	LanguageCodeEnumFrGa,
	LanguageCodeEnumFrGf,
	LanguageCodeEnumFrGn,
	LanguageCodeEnumFrGp,
	LanguageCodeEnumFrGq,
	LanguageCodeEnumFrHt,
	LanguageCodeEnumFrKm,
	LanguageCodeEnumFrLu,
	LanguageCodeEnumFrMa,
	LanguageCodeEnumFrMc,
	LanguageCodeEnumFrMf,
	LanguageCodeEnumFrMg,
	LanguageCodeEnumFrMl,
	LanguageCodeEnumFrMq,
	LanguageCodeEnumFrMr,
	LanguageCodeEnumFrMu,
	LanguageCodeEnumFrNc,
	LanguageCodeEnumFrNe,
	LanguageCodeEnumFrPf,
	LanguageCodeEnumFrPm,
	LanguageCodeEnumFrRe,
	LanguageCodeEnumFrRw,
	LanguageCodeEnumFrSc,
	LanguageCodeEnumFrSn,
	LanguageCodeEnumFrSy,
	LanguageCodeEnumFrTd,
	LanguageCodeEnumFrTg,
	LanguageCodeEnumFrTn,
	LanguageCodeEnumFrVu,
	LanguageCodeEnumFrWf,
	LanguageCodeEnumFrYt,
	LanguageCodeEnumFur,
	LanguageCodeEnumFurIt,
	LanguageCodeEnumFy,
	LanguageCodeEnumFyNl,
	LanguageCodeEnumGa,
	LanguageCodeEnumGaGb,
	LanguageCodeEnumGaIe,
	LanguageCodeEnumGd,
	LanguageCodeEnumGdGb,
	LanguageCodeEnumGl,
	LanguageCodeEnumGlEs,
	LanguageCodeEnumGsw,
	LanguageCodeEnumGswCh,
	LanguageCodeEnumGswFr,
	LanguageCodeEnumGswLi,
	LanguageCodeEnumGu,
	LanguageCodeEnumGuIn,
	LanguageCodeEnumGuz,
	LanguageCodeEnumGuzKe,
	LanguageCodeEnumGv,
	LanguageCodeEnumGvIm,
	LanguageCodeEnumHa,
	LanguageCodeEnumHaGh,
	LanguageCodeEnumHaNe,
	LanguageCodeEnumHaNg,
	LanguageCodeEnumHaw,
	LanguageCodeEnumHawUs,
	LanguageCodeEnumHe,
	LanguageCodeEnumHeIl,
	LanguageCodeEnumHi,
	LanguageCodeEnumHiIn,
	LanguageCodeEnumHr,
	LanguageCodeEnumHrBa,
	LanguageCodeEnumHrHr,
	LanguageCodeEnumHsb,
	LanguageCodeEnumHsbDe,
	LanguageCodeEnumHu,
	LanguageCodeEnumHuHu,
	LanguageCodeEnumHy,
	LanguageCodeEnumHyAm,
	LanguageCodeEnumIa,
	LanguageCodeEnumID,
	LanguageCodeEnumIDID,
	LanguageCodeEnumIg,
	LanguageCodeEnumIgNg,
	LanguageCodeEnumIi,
	LanguageCodeEnumIiCn,
	LanguageCodeEnumIs,
	LanguageCodeEnumIsIs,
	LanguageCodeEnumIt,
	LanguageCodeEnumItCh,
	LanguageCodeEnumItIt,
	LanguageCodeEnumItSm,
	LanguageCodeEnumItVa,
	LanguageCodeEnumJa,
	LanguageCodeEnumJaJp,
	LanguageCodeEnumJgo,
	LanguageCodeEnumJgoCm,
	LanguageCodeEnumJmc,
	LanguageCodeEnumJmcTz,
	LanguageCodeEnumJv,
	LanguageCodeEnumJvID,
	LanguageCodeEnumKa,
	LanguageCodeEnumKaGe,
	LanguageCodeEnumKab,
	LanguageCodeEnumKabDz,
	LanguageCodeEnumKam,
	LanguageCodeEnumKamKe,
	LanguageCodeEnumKde,
	LanguageCodeEnumKdeTz,
	LanguageCodeEnumKea,
	LanguageCodeEnumKeaCv,
	LanguageCodeEnumKhq,
	LanguageCodeEnumKhqMl,
	LanguageCodeEnumKi,
	LanguageCodeEnumKiKe,
	LanguageCodeEnumKk,
	LanguageCodeEnumKkKz,
	LanguageCodeEnumKkj,
	LanguageCodeEnumKkjCm,
	LanguageCodeEnumKl,
	LanguageCodeEnumKlGl,
	LanguageCodeEnumKln,
	LanguageCodeEnumKlnKe,
	LanguageCodeEnumKm,
	LanguageCodeEnumKmKh,
	LanguageCodeEnumKn,
	LanguageCodeEnumKnIn,
	LanguageCodeEnumKo,
	LanguageCodeEnumKoKp,
	LanguageCodeEnumKoKr,
	LanguageCodeEnumKok,
	LanguageCodeEnumKokIn,
	LanguageCodeEnumKs,
	LanguageCodeEnumKsArab,
	LanguageCodeEnumKsArabIn,
	LanguageCodeEnumKsb,
	LanguageCodeEnumKsbTz,
	LanguageCodeEnumKsf,
	LanguageCodeEnumKsfCm,
	LanguageCodeEnumKsh,
	LanguageCodeEnumKshDe,
	LanguageCodeEnumKu,
	LanguageCodeEnumKuTr,
	LanguageCodeEnumKw,
	LanguageCodeEnumKwGb,
	LanguageCodeEnumKy,
	LanguageCodeEnumKyKg,
	LanguageCodeEnumLag,
	LanguageCodeEnumLagTz,
	LanguageCodeEnumLb,
	LanguageCodeEnumLbLu,
	LanguageCodeEnumLg,
	LanguageCodeEnumLgUg,
	LanguageCodeEnumLkt,
	LanguageCodeEnumLktUs,
	LanguageCodeEnumLn,
	LanguageCodeEnumLnAo,
	LanguageCodeEnumLnCd,
	LanguageCodeEnumLnCf,
	LanguageCodeEnumLnCg,
	LanguageCodeEnumLo,
	LanguageCodeEnumLoLa,
	LanguageCodeEnumLrc,
	LanguageCodeEnumLrcIq,
	LanguageCodeEnumLrcIr,
	LanguageCodeEnumLt,
	LanguageCodeEnumLtLt,
	LanguageCodeEnumLu,
	LanguageCodeEnumLuCd,
	LanguageCodeEnumLuo,
	LanguageCodeEnumLuoKe,
	LanguageCodeEnumLuy,
	LanguageCodeEnumLuyKe,
	LanguageCodeEnumLv,
	LanguageCodeEnumLvLv,
	LanguageCodeEnumMai,
	LanguageCodeEnumMaiIn,
	LanguageCodeEnumMas,
	LanguageCodeEnumMasKe,
	LanguageCodeEnumMasTz,
	LanguageCodeEnumMer,
	LanguageCodeEnumMerKe,
	LanguageCodeEnumMfe,
	LanguageCodeEnumMfeMu,
	LanguageCodeEnumMg,
	LanguageCodeEnumMgMg,
	LanguageCodeEnumMgh,
	LanguageCodeEnumMghMz,
	LanguageCodeEnumMgo,
	LanguageCodeEnumMgoCm,
	LanguageCodeEnumMi,
	LanguageCodeEnumMiNz,
	LanguageCodeEnumMk,
	LanguageCodeEnumMkMk,
	LanguageCodeEnumMl,
	LanguageCodeEnumMlIn,
	LanguageCodeEnumMn,
	LanguageCodeEnumMnMn,
	LanguageCodeEnumMni,
	LanguageCodeEnumMniBeng,
	LanguageCodeEnumMniBengIn,
	LanguageCodeEnumMr,
	LanguageCodeEnumMrIn,
	LanguageCodeEnumMs,
	LanguageCodeEnumMsBn,
	LanguageCodeEnumMsID,
	LanguageCodeEnumMsMy,
	LanguageCodeEnumMsSg,
	LanguageCodeEnumMt,
	LanguageCodeEnumMtMt,
	LanguageCodeEnumMua,
	LanguageCodeEnumMuaCm,
	LanguageCodeEnumMy,
	LanguageCodeEnumMyMm,
	LanguageCodeEnumMzn,
	LanguageCodeEnumMznIr,
	LanguageCodeEnumNaq,
	LanguageCodeEnumNaqNa,
	LanguageCodeEnumNb,
	LanguageCodeEnumNbNo,
	LanguageCodeEnumNbSj,
	LanguageCodeEnumNd,
	LanguageCodeEnumNdZw,
	LanguageCodeEnumNds,
	LanguageCodeEnumNdsDe,
	LanguageCodeEnumNdsNl,
	LanguageCodeEnumNe,
	LanguageCodeEnumNeIn,
	LanguageCodeEnumNeNp,
	LanguageCodeEnumNl,
	LanguageCodeEnumNlAw,
	LanguageCodeEnumNlBe,
	LanguageCodeEnumNlBq,
	LanguageCodeEnumNlCw,
	LanguageCodeEnumNlNl,
	LanguageCodeEnumNlSr,
	LanguageCodeEnumNlSx,
	LanguageCodeEnumNmg,
	LanguageCodeEnumNmgCm,
	LanguageCodeEnumNn,
	LanguageCodeEnumNnNo,
	LanguageCodeEnumNnh,
	LanguageCodeEnumNnhCm,
	LanguageCodeEnumNus,
	LanguageCodeEnumNusSs,
	LanguageCodeEnumNyn,
	LanguageCodeEnumNynUg,
	LanguageCodeEnumOm,
	LanguageCodeEnumOmEt,
	LanguageCodeEnumOmKe,
	LanguageCodeEnumOr,
	LanguageCodeEnumOrIn,
	LanguageCodeEnumOs,
	LanguageCodeEnumOsGe,
	LanguageCodeEnumOsRu,
	LanguageCodeEnumPa,
	LanguageCodeEnumPaArab,
	LanguageCodeEnumPaArabPk,
	LanguageCodeEnumPaGuru,
	LanguageCodeEnumPaGuruIn,
	LanguageCodeEnumPcm,
	LanguageCodeEnumPcmNg,
	LanguageCodeEnumPl,
	LanguageCodeEnumPlPl,
	LanguageCodeEnumPrg,
	LanguageCodeEnumPs,
	LanguageCodeEnumPsAf,
	LanguageCodeEnumPsPk,
	LanguageCodeEnumPt,
	LanguageCodeEnumPtAo,
	LanguageCodeEnumPtBr,
	LanguageCodeEnumPtCh,
	LanguageCodeEnumPtCv,
	LanguageCodeEnumPtGq,
	LanguageCodeEnumPtGw,
	LanguageCodeEnumPtLu,
	LanguageCodeEnumPtMo,
	LanguageCodeEnumPtMz,
	LanguageCodeEnumPtPt,
	LanguageCodeEnumPtSt,
	LanguageCodeEnumPtTl,
	LanguageCodeEnumQu,
	LanguageCodeEnumQuBo,
	LanguageCodeEnumQuEc,
	LanguageCodeEnumQuPe,
	LanguageCodeEnumRm,
	LanguageCodeEnumRmCh,
	LanguageCodeEnumRn,
	LanguageCodeEnumRnBi,
	LanguageCodeEnumRo,
	LanguageCodeEnumRoMd,
	LanguageCodeEnumRoRo,
	LanguageCodeEnumRof,
	LanguageCodeEnumRofTz,
	LanguageCodeEnumRu,
	LanguageCodeEnumRuBy,
	LanguageCodeEnumRuKg,
	LanguageCodeEnumRuKz,
	LanguageCodeEnumRuMd,
	LanguageCodeEnumRuRu,
	LanguageCodeEnumRuUa,
	LanguageCodeEnumRw,
	LanguageCodeEnumRwRw,
	LanguageCodeEnumRwk,
	LanguageCodeEnumRwkTz,
	LanguageCodeEnumSah,
	LanguageCodeEnumSahRu,
	LanguageCodeEnumSaq,
	LanguageCodeEnumSaqKe,
	LanguageCodeEnumSat,
	LanguageCodeEnumSatOlck,
	LanguageCodeEnumSatOlckIn,
	LanguageCodeEnumSbp,
	LanguageCodeEnumSbpTz,
	LanguageCodeEnumSd,
	LanguageCodeEnumSdArab,
	LanguageCodeEnumSdArabPk,
	LanguageCodeEnumSdDeva,
	LanguageCodeEnumSdDevaIn,
	LanguageCodeEnumSe,
	LanguageCodeEnumSeFi,
	LanguageCodeEnumSeNo,
	LanguageCodeEnumSeSe,
	LanguageCodeEnumSeh,
	LanguageCodeEnumSehMz,
	LanguageCodeEnumSes,
	LanguageCodeEnumSesMl,
	LanguageCodeEnumSg,
	LanguageCodeEnumSgCf,
	LanguageCodeEnumShi,
	LanguageCodeEnumShiLatn,
	LanguageCodeEnumShiLatnMa,
	LanguageCodeEnumShiTfng,
	LanguageCodeEnumShiTfngMa,
	LanguageCodeEnumSi,
	LanguageCodeEnumSiLk,
	LanguageCodeEnumSk,
	LanguageCodeEnumSkSk,
	LanguageCodeEnumSl,
	LanguageCodeEnumSlSi,
	LanguageCodeEnumSmn,
	LanguageCodeEnumSmnFi,
	LanguageCodeEnumSn,
	LanguageCodeEnumSnZw,
	LanguageCodeEnumSo,
	LanguageCodeEnumSoDj,
	LanguageCodeEnumSoEt,
	LanguageCodeEnumSoKe,
	LanguageCodeEnumSoSo,
	LanguageCodeEnumSq,
	LanguageCodeEnumSqAl,
	LanguageCodeEnumSqMk,
	LanguageCodeEnumSqXk,
	LanguageCodeEnumSr,
	LanguageCodeEnumSrCyrl,
	LanguageCodeEnumSrCyrlBa,
	LanguageCodeEnumSrCyrlMe,
	LanguageCodeEnumSrCyrlRs,
	LanguageCodeEnumSrCyrlXk,
	LanguageCodeEnumSrLatn,
	LanguageCodeEnumSrLatnBa,
	LanguageCodeEnumSrLatnMe,
	LanguageCodeEnumSrLatnRs,
	LanguageCodeEnumSrLatnXk,
	LanguageCodeEnumSu,
	LanguageCodeEnumSuLatn,
	LanguageCodeEnumSuLatnID,
	LanguageCodeEnumSv,
	LanguageCodeEnumSvAx,
	LanguageCodeEnumSvFi,
	LanguageCodeEnumSvSe,
	LanguageCodeEnumSw,
	LanguageCodeEnumSwCd,
	LanguageCodeEnumSwKe,
	LanguageCodeEnumSwTz,
	LanguageCodeEnumSwUg,
	LanguageCodeEnumTa,
	LanguageCodeEnumTaIn,
	LanguageCodeEnumTaLk,
	LanguageCodeEnumTaMy,
	LanguageCodeEnumTaSg,
	LanguageCodeEnumTe,
	LanguageCodeEnumTeIn,
	LanguageCodeEnumTeo,
	LanguageCodeEnumTeoKe,
	LanguageCodeEnumTeoUg,
	LanguageCodeEnumTg,
	LanguageCodeEnumTgTj,
	LanguageCodeEnumTh,
	LanguageCodeEnumThTh,
	LanguageCodeEnumTi,
	LanguageCodeEnumTiEr,
	LanguageCodeEnumTiEt,
	LanguageCodeEnumTk,
	LanguageCodeEnumTkTm,
	LanguageCodeEnumTo,
	LanguageCodeEnumToTo,
	LanguageCodeEnumTr,
	LanguageCodeEnumTrCy,
	LanguageCodeEnumTrTr,
	LanguageCodeEnumTt,
	LanguageCodeEnumTtRu,
	LanguageCodeEnumTwq,
	LanguageCodeEnumTwqNe,
	LanguageCodeEnumTzm,
	LanguageCodeEnumTzmMa,
	LanguageCodeEnumUg,
	LanguageCodeEnumUgCn,
	LanguageCodeEnumUk,
	LanguageCodeEnumUkUa,
	LanguageCodeEnumUr,
	LanguageCodeEnumUrIn,
	LanguageCodeEnumUrPk,
	LanguageCodeEnumUz,
	LanguageCodeEnumUzArab,
	LanguageCodeEnumUzArabAf,
	LanguageCodeEnumUzCyrl,
	LanguageCodeEnumUzCyrlUz,
	LanguageCodeEnumUzLatn,
	LanguageCodeEnumUzLatnUz,
	LanguageCodeEnumVai,
	LanguageCodeEnumVaiLatn,
	LanguageCodeEnumVaiLatnLr,
	LanguageCodeEnumVaiVaii,
	LanguageCodeEnumVaiVaiiLr,
	LanguageCodeEnumVi,
	LanguageCodeEnumViVn,
	LanguageCodeEnumVo,
	LanguageCodeEnumVun,
	LanguageCodeEnumVunTz,
	LanguageCodeEnumWae,
	LanguageCodeEnumWaeCh,
	LanguageCodeEnumWo,
	LanguageCodeEnumWoSn,
	LanguageCodeEnumXh,
	LanguageCodeEnumXhZa,
	LanguageCodeEnumXog,
	LanguageCodeEnumXogUg,
	LanguageCodeEnumYav,
	LanguageCodeEnumYavCm,
	LanguageCodeEnumYi,
	LanguageCodeEnumYo,
	LanguageCodeEnumYoBj,
	LanguageCodeEnumYoNg,
	LanguageCodeEnumYue,
	LanguageCodeEnumYueHans,
	LanguageCodeEnumYueHansCn,
	LanguageCodeEnumYueHant,
	LanguageCodeEnumYueHantHk,
	LanguageCodeEnumZgh,
	LanguageCodeEnumZghMa,
	LanguageCodeEnumZh,
	LanguageCodeEnumZhHans,
	LanguageCodeEnumZhHansCn,
	LanguageCodeEnumZhHansHk,
	LanguageCodeEnumZhHansMo,
	LanguageCodeEnumZhHansSg,
	LanguageCodeEnumZhHant,
	LanguageCodeEnumZhHantHk,
	LanguageCodeEnumZhHantMo,
	LanguageCodeEnumZhHantTw,
	LanguageCodeEnumZu,
	LanguageCodeEnumZuZa,
}

func (e LanguageCodeEnum) IsValid() bool {
	switch e {
	case LanguageCodeEnumAf, LanguageCodeEnumAfNa, LanguageCodeEnumAfZa, LanguageCodeEnumAgq, LanguageCodeEnumAgqCm, LanguageCodeEnumAk, LanguageCodeEnumAkGh, LanguageCodeEnumAm, LanguageCodeEnumAmEt, LanguageCodeEnumAr, LanguageCodeEnumArAe, LanguageCodeEnumArBh, LanguageCodeEnumArDj, LanguageCodeEnumArDz, LanguageCodeEnumArEg, LanguageCodeEnumArEh, LanguageCodeEnumArEr, LanguageCodeEnumArIl, LanguageCodeEnumArIq, LanguageCodeEnumArJo, LanguageCodeEnumArKm, LanguageCodeEnumArKw, LanguageCodeEnumArLb, LanguageCodeEnumArLy, LanguageCodeEnumArMa, LanguageCodeEnumArMr, LanguageCodeEnumArOm, LanguageCodeEnumArPs, LanguageCodeEnumArQa, LanguageCodeEnumArSa, LanguageCodeEnumArSd, LanguageCodeEnumArSo, LanguageCodeEnumArSs, LanguageCodeEnumArSy, LanguageCodeEnumArTd, LanguageCodeEnumArTn, LanguageCodeEnumArYe, LanguageCodeEnumAs, LanguageCodeEnumAsIn, LanguageCodeEnumAsa, LanguageCodeEnumAsaTz, LanguageCodeEnumAst, LanguageCodeEnumAstEs, LanguageCodeEnumAz, LanguageCodeEnumAzCyrl, LanguageCodeEnumAzCyrlAz, LanguageCodeEnumAzLatn, LanguageCodeEnumAzLatnAz, LanguageCodeEnumBas, LanguageCodeEnumBasCm, LanguageCodeEnumBe, LanguageCodeEnumBeBy, LanguageCodeEnumBem, LanguageCodeEnumBemZm, LanguageCodeEnumBez, LanguageCodeEnumBezTz, LanguageCodeEnumBg, LanguageCodeEnumBgBg, LanguageCodeEnumBm, LanguageCodeEnumBmMl, LanguageCodeEnumBn, LanguageCodeEnumBnBd, LanguageCodeEnumBnIn, LanguageCodeEnumBo, LanguageCodeEnumBoCn, LanguageCodeEnumBoIn, LanguageCodeEnumBr, LanguageCodeEnumBrFr, LanguageCodeEnumBrx, LanguageCodeEnumBrxIn, LanguageCodeEnumBs, LanguageCodeEnumBsCyrl, LanguageCodeEnumBsCyrlBa, LanguageCodeEnumBsLatn, LanguageCodeEnumBsLatnBa, LanguageCodeEnumCa, LanguageCodeEnumCaAd, LanguageCodeEnumCaEs, LanguageCodeEnumCaEsValencia, LanguageCodeEnumCaFr, LanguageCodeEnumCaIt, LanguageCodeEnumCcp, LanguageCodeEnumCcpBd, LanguageCodeEnumCcpIn, LanguageCodeEnumCe, LanguageCodeEnumCeRu, LanguageCodeEnumCeb, LanguageCodeEnumCebPh, LanguageCodeEnumCgg, LanguageCodeEnumCggUg, LanguageCodeEnumChr, LanguageCodeEnumChrUs, LanguageCodeEnumCkb, LanguageCodeEnumCkbIq, LanguageCodeEnumCkbIr, LanguageCodeEnumCs, LanguageCodeEnumCsCz, LanguageCodeEnumCu, LanguageCodeEnumCuRu, LanguageCodeEnumCy, LanguageCodeEnumCyGb, LanguageCodeEnumDa, LanguageCodeEnumDaDk, LanguageCodeEnumDaGl, LanguageCodeEnumDav, LanguageCodeEnumDavKe, LanguageCodeEnumDe, LanguageCodeEnumDeAt, LanguageCodeEnumDeBe, LanguageCodeEnumDeCh, LanguageCodeEnumDeDe, LanguageCodeEnumDeIt, LanguageCodeEnumDeLi, LanguageCodeEnumDeLu, LanguageCodeEnumDje, LanguageCodeEnumDjeNe, LanguageCodeEnumDsb, LanguageCodeEnumDsbDe, LanguageCodeEnumDua, LanguageCodeEnumDuaCm, LanguageCodeEnumDyo, LanguageCodeEnumDyoSn, LanguageCodeEnumDz, LanguageCodeEnumDzBt, LanguageCodeEnumEbu, LanguageCodeEnumEbuKe, LanguageCodeEnumEe, LanguageCodeEnumEeGh, LanguageCodeEnumEeTg, LanguageCodeEnumEl, LanguageCodeEnumElCy, LanguageCodeEnumElGr, LanguageCodeEnumEn, LanguageCodeEnumEnAe, LanguageCodeEnumEnAg, LanguageCodeEnumEnAi, LanguageCodeEnumEnAs, LanguageCodeEnumEnAt, LanguageCodeEnumEnAu, LanguageCodeEnumEnBb, LanguageCodeEnumEnBe, LanguageCodeEnumEnBi, LanguageCodeEnumEnBm, LanguageCodeEnumEnBs, LanguageCodeEnumEnBw, LanguageCodeEnumEnBz, LanguageCodeEnumEnCa, LanguageCodeEnumEnCc, LanguageCodeEnumEnCh, LanguageCodeEnumEnCk, LanguageCodeEnumEnCm, LanguageCodeEnumEnCx, LanguageCodeEnumEnCy, LanguageCodeEnumEnDe, LanguageCodeEnumEnDg, LanguageCodeEnumEnDk, LanguageCodeEnumEnDm, LanguageCodeEnumEnEr, LanguageCodeEnumEnFi, LanguageCodeEnumEnFj, LanguageCodeEnumEnFk, LanguageCodeEnumEnFm, LanguageCodeEnumEnGb, LanguageCodeEnumEnGd, LanguageCodeEnumEnGg, LanguageCodeEnumEnGh, LanguageCodeEnumEnGi, LanguageCodeEnumEnGm, LanguageCodeEnumEnGu, LanguageCodeEnumEnGy, LanguageCodeEnumEnHk, LanguageCodeEnumEnIe, LanguageCodeEnumEnIl, LanguageCodeEnumEnIm, LanguageCodeEnumEnIn, LanguageCodeEnumEnIo, LanguageCodeEnumEnJe, LanguageCodeEnumEnJm, LanguageCodeEnumEnKe, LanguageCodeEnumEnKi, LanguageCodeEnumEnKn, LanguageCodeEnumEnKy, LanguageCodeEnumEnLc, LanguageCodeEnumEnLr, LanguageCodeEnumEnLs, LanguageCodeEnumEnMg, LanguageCodeEnumEnMh, LanguageCodeEnumEnMo, LanguageCodeEnumEnMp, LanguageCodeEnumEnMs, LanguageCodeEnumEnMt, LanguageCodeEnumEnMu, LanguageCodeEnumEnMw, LanguageCodeEnumEnMy, LanguageCodeEnumEnNa, LanguageCodeEnumEnNf, LanguageCodeEnumEnNg, LanguageCodeEnumEnNl, LanguageCodeEnumEnNr, LanguageCodeEnumEnNu, LanguageCodeEnumEnNz, LanguageCodeEnumEnPg, LanguageCodeEnumEnPh, LanguageCodeEnumEnPk, LanguageCodeEnumEnPn, LanguageCodeEnumEnPr, LanguageCodeEnumEnPw, LanguageCodeEnumEnRw, LanguageCodeEnumEnSb, LanguageCodeEnumEnSc, LanguageCodeEnumEnSd, LanguageCodeEnumEnSe, LanguageCodeEnumEnSg, LanguageCodeEnumEnSh, LanguageCodeEnumEnSi, LanguageCodeEnumEnSl, LanguageCodeEnumEnSs, LanguageCodeEnumEnSx, LanguageCodeEnumEnSz, LanguageCodeEnumEnTc, LanguageCodeEnumEnTk, LanguageCodeEnumEnTo, LanguageCodeEnumEnTt, LanguageCodeEnumEnTv, LanguageCodeEnumEnTz, LanguageCodeEnumEnUg, LanguageCodeEnumEnUm, LanguageCodeEnumEnUs, LanguageCodeEnumEnVc, LanguageCodeEnumEnVg, LanguageCodeEnumEnVi, LanguageCodeEnumEnVu, LanguageCodeEnumEnWs, LanguageCodeEnumEnZa, LanguageCodeEnumEnZm, LanguageCodeEnumEnZw, LanguageCodeEnumEo, LanguageCodeEnumEs, LanguageCodeEnumEsAr, LanguageCodeEnumEsBo, LanguageCodeEnumEsBr, LanguageCodeEnumEsBz, LanguageCodeEnumEsCl, LanguageCodeEnumEsCo, LanguageCodeEnumEsCr, LanguageCodeEnumEsCu, LanguageCodeEnumEsDo, LanguageCodeEnumEsEa, LanguageCodeEnumEsEc, LanguageCodeEnumEsEs, LanguageCodeEnumEsGq, LanguageCodeEnumEsGt, LanguageCodeEnumEsHn, LanguageCodeEnumEsIc, LanguageCodeEnumEsMx, LanguageCodeEnumEsNi, LanguageCodeEnumEsPa, LanguageCodeEnumEsPe, LanguageCodeEnumEsPh, LanguageCodeEnumEsPr, LanguageCodeEnumEsPy, LanguageCodeEnumEsSv, LanguageCodeEnumEsUs, LanguageCodeEnumEsUy, LanguageCodeEnumEsVe, LanguageCodeEnumEt, LanguageCodeEnumEtEe, LanguageCodeEnumEu, LanguageCodeEnumEuEs, LanguageCodeEnumEwo, LanguageCodeEnumEwoCm, LanguageCodeEnumFa, LanguageCodeEnumFaAf, LanguageCodeEnumFaIr, LanguageCodeEnumFf, LanguageCodeEnumFfAdlm, LanguageCodeEnumFfAdlmBf, LanguageCodeEnumFfAdlmCm, LanguageCodeEnumFfAdlmGh, LanguageCodeEnumFfAdlmGm, LanguageCodeEnumFfAdlmGn, LanguageCodeEnumFfAdlmGw, LanguageCodeEnumFfAdlmLr, LanguageCodeEnumFfAdlmMr, LanguageCodeEnumFfAdlmNe, LanguageCodeEnumFfAdlmNg, LanguageCodeEnumFfAdlmSl, LanguageCodeEnumFfAdlmSn, LanguageCodeEnumFfLatn, LanguageCodeEnumFfLatnBf, LanguageCodeEnumFfLatnCm, LanguageCodeEnumFfLatnGh, LanguageCodeEnumFfLatnGm, LanguageCodeEnumFfLatnGn, LanguageCodeEnumFfLatnGw, LanguageCodeEnumFfLatnLr, LanguageCodeEnumFfLatnMr, LanguageCodeEnumFfLatnNe, LanguageCodeEnumFfLatnNg, LanguageCodeEnumFfLatnSl, LanguageCodeEnumFfLatnSn, LanguageCodeEnumFi, LanguageCodeEnumFiFi, LanguageCodeEnumFil, LanguageCodeEnumFilPh, LanguageCodeEnumFo, LanguageCodeEnumFoDk, LanguageCodeEnumFoFo, LanguageCodeEnumFr, LanguageCodeEnumFrBe, LanguageCodeEnumFrBf, LanguageCodeEnumFrBi, LanguageCodeEnumFrBj, LanguageCodeEnumFrBl, LanguageCodeEnumFrCa, LanguageCodeEnumFrCd, LanguageCodeEnumFrCf, LanguageCodeEnumFrCg, LanguageCodeEnumFrCh, LanguageCodeEnumFrCi, LanguageCodeEnumFrCm, LanguageCodeEnumFrDj, LanguageCodeEnumFrDz, LanguageCodeEnumFrFr, LanguageCodeEnumFrGa, LanguageCodeEnumFrGf, LanguageCodeEnumFrGn, LanguageCodeEnumFrGp, LanguageCodeEnumFrGq, LanguageCodeEnumFrHt, LanguageCodeEnumFrKm, LanguageCodeEnumFrLu, LanguageCodeEnumFrMa, LanguageCodeEnumFrMc, LanguageCodeEnumFrMf, LanguageCodeEnumFrMg, LanguageCodeEnumFrMl, LanguageCodeEnumFrMq, LanguageCodeEnumFrMr, LanguageCodeEnumFrMu, LanguageCodeEnumFrNc, LanguageCodeEnumFrNe, LanguageCodeEnumFrPf, LanguageCodeEnumFrPm, LanguageCodeEnumFrRe, LanguageCodeEnumFrRw, LanguageCodeEnumFrSc, LanguageCodeEnumFrSn, LanguageCodeEnumFrSy, LanguageCodeEnumFrTd, LanguageCodeEnumFrTg, LanguageCodeEnumFrTn, LanguageCodeEnumFrVu, LanguageCodeEnumFrWf, LanguageCodeEnumFrYt, LanguageCodeEnumFur, LanguageCodeEnumFurIt, LanguageCodeEnumFy, LanguageCodeEnumFyNl, LanguageCodeEnumGa, LanguageCodeEnumGaGb, LanguageCodeEnumGaIe, LanguageCodeEnumGd, LanguageCodeEnumGdGb, LanguageCodeEnumGl, LanguageCodeEnumGlEs, LanguageCodeEnumGsw, LanguageCodeEnumGswCh, LanguageCodeEnumGswFr, LanguageCodeEnumGswLi, LanguageCodeEnumGu, LanguageCodeEnumGuIn, LanguageCodeEnumGuz, LanguageCodeEnumGuzKe, LanguageCodeEnumGv, LanguageCodeEnumGvIm, LanguageCodeEnumHa, LanguageCodeEnumHaGh, LanguageCodeEnumHaNe, LanguageCodeEnumHaNg, LanguageCodeEnumHaw, LanguageCodeEnumHawUs, LanguageCodeEnumHe, LanguageCodeEnumHeIl, LanguageCodeEnumHi, LanguageCodeEnumHiIn, LanguageCodeEnumHr, LanguageCodeEnumHrBa, LanguageCodeEnumHrHr, LanguageCodeEnumHsb, LanguageCodeEnumHsbDe, LanguageCodeEnumHu, LanguageCodeEnumHuHu, LanguageCodeEnumHy, LanguageCodeEnumHyAm, LanguageCodeEnumIa, LanguageCodeEnumID, LanguageCodeEnumIDID, LanguageCodeEnumIg, LanguageCodeEnumIgNg, LanguageCodeEnumIi, LanguageCodeEnumIiCn, LanguageCodeEnumIs, LanguageCodeEnumIsIs, LanguageCodeEnumIt, LanguageCodeEnumItCh, LanguageCodeEnumItIt, LanguageCodeEnumItSm, LanguageCodeEnumItVa, LanguageCodeEnumJa, LanguageCodeEnumJaJp, LanguageCodeEnumJgo, LanguageCodeEnumJgoCm, LanguageCodeEnumJmc, LanguageCodeEnumJmcTz, LanguageCodeEnumJv, LanguageCodeEnumJvID, LanguageCodeEnumKa, LanguageCodeEnumKaGe, LanguageCodeEnumKab, LanguageCodeEnumKabDz, LanguageCodeEnumKam, LanguageCodeEnumKamKe, LanguageCodeEnumKde, LanguageCodeEnumKdeTz, LanguageCodeEnumKea, LanguageCodeEnumKeaCv, LanguageCodeEnumKhq, LanguageCodeEnumKhqMl, LanguageCodeEnumKi, LanguageCodeEnumKiKe, LanguageCodeEnumKk, LanguageCodeEnumKkKz, LanguageCodeEnumKkj, LanguageCodeEnumKkjCm, LanguageCodeEnumKl, LanguageCodeEnumKlGl, LanguageCodeEnumKln, LanguageCodeEnumKlnKe, LanguageCodeEnumKm, LanguageCodeEnumKmKh, LanguageCodeEnumKn, LanguageCodeEnumKnIn, LanguageCodeEnumKo, LanguageCodeEnumKoKp, LanguageCodeEnumKoKr, LanguageCodeEnumKok, LanguageCodeEnumKokIn, LanguageCodeEnumKs, LanguageCodeEnumKsArab, LanguageCodeEnumKsArabIn, LanguageCodeEnumKsb, LanguageCodeEnumKsbTz, LanguageCodeEnumKsf, LanguageCodeEnumKsfCm, LanguageCodeEnumKsh, LanguageCodeEnumKshDe, LanguageCodeEnumKu, LanguageCodeEnumKuTr, LanguageCodeEnumKw, LanguageCodeEnumKwGb, LanguageCodeEnumKy, LanguageCodeEnumKyKg, LanguageCodeEnumLag, LanguageCodeEnumLagTz, LanguageCodeEnumLb, LanguageCodeEnumLbLu, LanguageCodeEnumLg, LanguageCodeEnumLgUg, LanguageCodeEnumLkt, LanguageCodeEnumLktUs, LanguageCodeEnumLn, LanguageCodeEnumLnAo, LanguageCodeEnumLnCd, LanguageCodeEnumLnCf, LanguageCodeEnumLnCg, LanguageCodeEnumLo, LanguageCodeEnumLoLa, LanguageCodeEnumLrc, LanguageCodeEnumLrcIq, LanguageCodeEnumLrcIr, LanguageCodeEnumLt, LanguageCodeEnumLtLt, LanguageCodeEnumLu, LanguageCodeEnumLuCd, LanguageCodeEnumLuo, LanguageCodeEnumLuoKe, LanguageCodeEnumLuy, LanguageCodeEnumLuyKe, LanguageCodeEnumLv, LanguageCodeEnumLvLv, LanguageCodeEnumMai, LanguageCodeEnumMaiIn, LanguageCodeEnumMas, LanguageCodeEnumMasKe, LanguageCodeEnumMasTz, LanguageCodeEnumMer, LanguageCodeEnumMerKe, LanguageCodeEnumMfe, LanguageCodeEnumMfeMu, LanguageCodeEnumMg, LanguageCodeEnumMgMg, LanguageCodeEnumMgh, LanguageCodeEnumMghMz, LanguageCodeEnumMgo, LanguageCodeEnumMgoCm, LanguageCodeEnumMi, LanguageCodeEnumMiNz, LanguageCodeEnumMk, LanguageCodeEnumMkMk, LanguageCodeEnumMl, LanguageCodeEnumMlIn, LanguageCodeEnumMn, LanguageCodeEnumMnMn, LanguageCodeEnumMni, LanguageCodeEnumMniBeng, LanguageCodeEnumMniBengIn, LanguageCodeEnumMr, LanguageCodeEnumMrIn, LanguageCodeEnumMs, LanguageCodeEnumMsBn, LanguageCodeEnumMsID, LanguageCodeEnumMsMy, LanguageCodeEnumMsSg, LanguageCodeEnumMt, LanguageCodeEnumMtMt, LanguageCodeEnumMua, LanguageCodeEnumMuaCm, LanguageCodeEnumMy, LanguageCodeEnumMyMm, LanguageCodeEnumMzn, LanguageCodeEnumMznIr, LanguageCodeEnumNaq, LanguageCodeEnumNaqNa, LanguageCodeEnumNb, LanguageCodeEnumNbNo, LanguageCodeEnumNbSj, LanguageCodeEnumNd, LanguageCodeEnumNdZw, LanguageCodeEnumNds, LanguageCodeEnumNdsDe, LanguageCodeEnumNdsNl, LanguageCodeEnumNe, LanguageCodeEnumNeIn, LanguageCodeEnumNeNp, LanguageCodeEnumNl, LanguageCodeEnumNlAw, LanguageCodeEnumNlBe, LanguageCodeEnumNlBq, LanguageCodeEnumNlCw, LanguageCodeEnumNlNl, LanguageCodeEnumNlSr, LanguageCodeEnumNlSx, LanguageCodeEnumNmg, LanguageCodeEnumNmgCm, LanguageCodeEnumNn, LanguageCodeEnumNnNo, LanguageCodeEnumNnh, LanguageCodeEnumNnhCm, LanguageCodeEnumNus, LanguageCodeEnumNusSs, LanguageCodeEnumNyn, LanguageCodeEnumNynUg, LanguageCodeEnumOm, LanguageCodeEnumOmEt, LanguageCodeEnumOmKe, LanguageCodeEnumOr, LanguageCodeEnumOrIn, LanguageCodeEnumOs, LanguageCodeEnumOsGe, LanguageCodeEnumOsRu, LanguageCodeEnumPa, LanguageCodeEnumPaArab, LanguageCodeEnumPaArabPk, LanguageCodeEnumPaGuru, LanguageCodeEnumPaGuruIn, LanguageCodeEnumPcm, LanguageCodeEnumPcmNg, LanguageCodeEnumPl, LanguageCodeEnumPlPl, LanguageCodeEnumPrg, LanguageCodeEnumPs, LanguageCodeEnumPsAf, LanguageCodeEnumPsPk, LanguageCodeEnumPt, LanguageCodeEnumPtAo, LanguageCodeEnumPtBr, LanguageCodeEnumPtCh, LanguageCodeEnumPtCv, LanguageCodeEnumPtGq, LanguageCodeEnumPtGw, LanguageCodeEnumPtLu, LanguageCodeEnumPtMo, LanguageCodeEnumPtMz, LanguageCodeEnumPtPt, LanguageCodeEnumPtSt, LanguageCodeEnumPtTl, LanguageCodeEnumQu, LanguageCodeEnumQuBo, LanguageCodeEnumQuEc, LanguageCodeEnumQuPe, LanguageCodeEnumRm, LanguageCodeEnumRmCh, LanguageCodeEnumRn, LanguageCodeEnumRnBi, LanguageCodeEnumRo, LanguageCodeEnumRoMd, LanguageCodeEnumRoRo, LanguageCodeEnumRof, LanguageCodeEnumRofTz, LanguageCodeEnumRu, LanguageCodeEnumRuBy, LanguageCodeEnumRuKg, LanguageCodeEnumRuKz, LanguageCodeEnumRuMd, LanguageCodeEnumRuRu, LanguageCodeEnumRuUa, LanguageCodeEnumRw, LanguageCodeEnumRwRw, LanguageCodeEnumRwk, LanguageCodeEnumRwkTz, LanguageCodeEnumSah, LanguageCodeEnumSahRu, LanguageCodeEnumSaq, LanguageCodeEnumSaqKe, LanguageCodeEnumSat, LanguageCodeEnumSatOlck, LanguageCodeEnumSatOlckIn, LanguageCodeEnumSbp, LanguageCodeEnumSbpTz, LanguageCodeEnumSd, LanguageCodeEnumSdArab, LanguageCodeEnumSdArabPk, LanguageCodeEnumSdDeva, LanguageCodeEnumSdDevaIn, LanguageCodeEnumSe, LanguageCodeEnumSeFi, LanguageCodeEnumSeNo, LanguageCodeEnumSeSe, LanguageCodeEnumSeh, LanguageCodeEnumSehMz, LanguageCodeEnumSes, LanguageCodeEnumSesMl, LanguageCodeEnumSg, LanguageCodeEnumSgCf, LanguageCodeEnumShi, LanguageCodeEnumShiLatn, LanguageCodeEnumShiLatnMa, LanguageCodeEnumShiTfng, LanguageCodeEnumShiTfngMa, LanguageCodeEnumSi, LanguageCodeEnumSiLk, LanguageCodeEnumSk, LanguageCodeEnumSkSk, LanguageCodeEnumSl, LanguageCodeEnumSlSi, LanguageCodeEnumSmn, LanguageCodeEnumSmnFi, LanguageCodeEnumSn, LanguageCodeEnumSnZw, LanguageCodeEnumSo, LanguageCodeEnumSoDj, LanguageCodeEnumSoEt, LanguageCodeEnumSoKe, LanguageCodeEnumSoSo, LanguageCodeEnumSq, LanguageCodeEnumSqAl, LanguageCodeEnumSqMk, LanguageCodeEnumSqXk, LanguageCodeEnumSr, LanguageCodeEnumSrCyrl, LanguageCodeEnumSrCyrlBa, LanguageCodeEnumSrCyrlMe, LanguageCodeEnumSrCyrlRs, LanguageCodeEnumSrCyrlXk, LanguageCodeEnumSrLatn, LanguageCodeEnumSrLatnBa, LanguageCodeEnumSrLatnMe, LanguageCodeEnumSrLatnRs, LanguageCodeEnumSrLatnXk, LanguageCodeEnumSu, LanguageCodeEnumSuLatn, LanguageCodeEnumSuLatnID, LanguageCodeEnumSv, LanguageCodeEnumSvAx, LanguageCodeEnumSvFi, LanguageCodeEnumSvSe, LanguageCodeEnumSw, LanguageCodeEnumSwCd, LanguageCodeEnumSwKe, LanguageCodeEnumSwTz, LanguageCodeEnumSwUg, LanguageCodeEnumTa, LanguageCodeEnumTaIn, LanguageCodeEnumTaLk, LanguageCodeEnumTaMy, LanguageCodeEnumTaSg, LanguageCodeEnumTe, LanguageCodeEnumTeIn, LanguageCodeEnumTeo, LanguageCodeEnumTeoKe, LanguageCodeEnumTeoUg, LanguageCodeEnumTg, LanguageCodeEnumTgTj, LanguageCodeEnumTh, LanguageCodeEnumThTh, LanguageCodeEnumTi, LanguageCodeEnumTiEr, LanguageCodeEnumTiEt, LanguageCodeEnumTk, LanguageCodeEnumTkTm, LanguageCodeEnumTo, LanguageCodeEnumToTo, LanguageCodeEnumTr, LanguageCodeEnumTrCy, LanguageCodeEnumTrTr, LanguageCodeEnumTt, LanguageCodeEnumTtRu, LanguageCodeEnumTwq, LanguageCodeEnumTwqNe, LanguageCodeEnumTzm, LanguageCodeEnumTzmMa, LanguageCodeEnumUg, LanguageCodeEnumUgCn, LanguageCodeEnumUk, LanguageCodeEnumUkUa, LanguageCodeEnumUr, LanguageCodeEnumUrIn, LanguageCodeEnumUrPk, LanguageCodeEnumUz, LanguageCodeEnumUzArab, LanguageCodeEnumUzArabAf, LanguageCodeEnumUzCyrl, LanguageCodeEnumUzCyrlUz, LanguageCodeEnumUzLatn, LanguageCodeEnumUzLatnUz, LanguageCodeEnumVai, LanguageCodeEnumVaiLatn, LanguageCodeEnumVaiLatnLr, LanguageCodeEnumVaiVaii, LanguageCodeEnumVaiVaiiLr, LanguageCodeEnumVi, LanguageCodeEnumViVn, LanguageCodeEnumVo, LanguageCodeEnumVun, LanguageCodeEnumVunTz, LanguageCodeEnumWae, LanguageCodeEnumWaeCh, LanguageCodeEnumWo, LanguageCodeEnumWoSn, LanguageCodeEnumXh, LanguageCodeEnumXhZa, LanguageCodeEnumXog, LanguageCodeEnumXogUg, LanguageCodeEnumYav, LanguageCodeEnumYavCm, LanguageCodeEnumYi, LanguageCodeEnumYo, LanguageCodeEnumYoBj, LanguageCodeEnumYoNg, LanguageCodeEnumYue, LanguageCodeEnumYueHans, LanguageCodeEnumYueHansCn, LanguageCodeEnumYueHant, LanguageCodeEnumYueHantHk, LanguageCodeEnumZgh, LanguageCodeEnumZghMa, LanguageCodeEnumZh, LanguageCodeEnumZhHans, LanguageCodeEnumZhHansCn, LanguageCodeEnumZhHansHk, LanguageCodeEnumZhHansMo, LanguageCodeEnumZhHansSg, LanguageCodeEnumZhHant, LanguageCodeEnumZhHantHk, LanguageCodeEnumZhHantMo, LanguageCodeEnumZhHantTw, LanguageCodeEnumZu, LanguageCodeEnumZuZa:
		return true
	}
	return false
}

func (e LanguageCodeEnum) String() string {
	return string(e)
}

func (e *LanguageCodeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LanguageCodeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageCodeEnum", str)
	}
	return nil
}

func (e LanguageCodeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderDirection string

const (
	// Specifies an ascending sort order.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending sort order.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StaffMemberStatus string

const (
	// User account has been activated.
	StaffMemberStatusActive StaffMemberStatus = "ACTIVE"
	// User account has not been activated yet.
	StaffMemberStatusDeactivated StaffMemberStatus = "DEACTIVATED"
)

var AllStaffMemberStatus = []StaffMemberStatus{
	StaffMemberStatusActive,
	StaffMemberStatusDeactivated,
}

func (e StaffMemberStatus) IsValid() bool {
	switch e {
	case StaffMemberStatusActive, StaffMemberStatusDeactivated:
		return true
	}
	return false
}

func (e StaffMemberStatus) String() string {
	return string(e)
}

func (e *StaffMemberStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StaffMemberStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StaffMemberStatus", str)
	}
	return nil
}

func (e StaffMemberStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserSortField string

const (
	// Sort users by first name.
	UserSortFieldFirstName UserSortField = "FIRST_NAME"
	// Sort users by last name.
	UserSortFieldLastName UserSortField = "LAST_NAME"
	// Sort users by email.
	UserSortFieldEmail UserSortField = "EMAIL"
	// Sort users by order count.
	UserSortFieldOrderCount UserSortField = "ORDER_COUNT"
	// Sort users by created at.
	UserSortFieldCreatedAt UserSortField = "CREATED_AT"
	// Sort users by last modified at.
	UserSortFieldLastModifiedAt UserSortField = "LAST_MODIFIED_AT"
)

var AllUserSortField = []UserSortField{
	UserSortFieldFirstName,
	UserSortFieldLastName,
	UserSortFieldEmail,
	UserSortFieldOrderCount,
	UserSortFieldCreatedAt,
	UserSortFieldLastModifiedAt,
}

func (e UserSortField) IsValid() bool {
	switch e {
	case UserSortFieldFirstName, UserSortFieldLastName, UserSortFieldEmail, UserSortFieldOrderCount, UserSortFieldCreatedAt, UserSortFieldLastModifiedAt:
		return true
	}
	return false
}

func (e UserSortField) String() string {
	return string(e)
}

func (e *UserSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserSortField", str)
	}
	return nil
}

func (e UserSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
