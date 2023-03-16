/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.22
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OidcUserInfo OpenID Connect Userinfo
type OidcUserInfo struct {
	// End-User's birthday, represented as an ISO 8601:2004 [ISO8601‑2004] YYYY-MM-DD format. The year MAY be 0000, indicating that it is omitted. To represent only the year, YYYY format is allowed. Note that depending on the underlying platform's date related function, providing just year can result in varying month and day, so the implementers need to take this factor into account to correctly process the dates.
	Birthdate *string `json:"birthdate,omitempty"`
	// End-User's preferred e-mail address. Its value MUST conform to the RFC 5322 [RFC5322] addr-spec syntax. The RP MUST NOT rely upon this value being unique, as discussed in Section 5.7.
	Email *string `json:"email,omitempty"`
	// True if the End-User's e-mail address has been verified; otherwise false. When this Claim Value is true, this means that the OP took affirmative steps to ensure that this e-mail address was controlled by the End-User at the time the verification was performed. The means by which an e-mail address is verified is context-specific, and dependent upon the trust framework or contractual agreements within which the parties are operating.
	EmailVerified *bool `json:"email_verified,omitempty"`
	// Surname(s) or last name(s) of the End-User. Note that in some cultures, people can have multiple family names or no family name; all can be present, with the names being separated by space characters.
	FamilyName *string `json:"family_name,omitempty"`
	// End-User's gender. Values defined by this specification are female and male. Other values MAY be used when neither of the defined values are applicable.
	Gender *string `json:"gender,omitempty"`
	// Given name(s) or first name(s) of the End-User. Note that in some cultures, people can have multiple given names; all can be present, with the names being separated by space characters.
	GivenName *string `json:"given_name,omitempty"`
	// End-User's locale, represented as a BCP47 [RFC5646] language tag. This is typically an ISO 639-1 Alpha-2 [ISO639‑1] language code in lowercase and an ISO 3166-1 Alpha-2 [ISO3166‑1] country code in uppercase, separated by a dash. For example, en-US or fr-CA. As a compatibility note, some implementations have used an underscore as the separator rather than a dash, for example, en_US; Relying Parties MAY choose to accept this locale syntax as well.
	Locale *string `json:"locale,omitempty"`
	// Middle name(s) of the End-User. Note that in some cultures, people can have multiple middle names; all can be present, with the names being separated by space characters. Also note that in some cultures, middle names are not used.
	MiddleName *string `json:"middle_name,omitempty"`
	// End-User's full name in displayable form including all name parts, possibly including titles and suffixes, ordered according to the End-User's locale and preferences.
	Name *string `json:"name,omitempty"`
	// Casual name of the End-User that may or may not be the same as the given_name. For instance, a nickname value of Mike might be returned alongside a given_name value of Michael.
	Nickname *string `json:"nickname,omitempty"`
	// End-User's preferred telephone number. E.164 [E.164] is RECOMMENDED as the format of this Claim, for example, +1 (425) 555-1212 or +56 (2) 687 2400. If the phone number contains an extension, it is RECOMMENDED that the extension be represented using the RFC 3966 [RFC3966] extension syntax, for example, +1 (604) 555-1234;ext=5678.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// True if the End-User's phone number has been verified; otherwise false. When this Claim Value is true, this means that the OP took affirmative steps to ensure that this phone number was controlled by the End-User at the time the verification was performed. The means by which a phone number is verified is context-specific, and dependent upon the trust framework or contractual agreements within which the parties are operating. When true, the phone_number Claim MUST be in E.164 format and any extensions MUST be represented in RFC 3966 format.
	PhoneNumberVerified *bool `json:"phone_number_verified,omitempty"`
	// URL of the End-User's profile picture. This URL MUST refer to an image file (for example, a PNG, JPEG, or GIF image file), rather than to a Web page containing an image. Note that this URL SHOULD specifically reference a profile photo of the End-User suitable for displaying when describing the End-User, rather than an arbitrary photo taken by the End-User.
	Picture *string `json:"picture,omitempty"`
	// Non-unique shorthand name by which the End-User wishes to be referred to at the RP, such as janedoe or j.doe. This value MAY be any valid JSON string including special characters such as @, /, or whitespace.
	PreferredUsername *string `json:"preferred_username,omitempty"`
	// URL of the End-User's profile page. The contents of this Web page SHOULD be about the End-User.
	Profile *string `json:"profile,omitempty"`
	// Subject - Identifier for the End-User at the IssuerURL.
	Sub *string `json:"sub,omitempty"`
	// Time the End-User's information was last updated. Its value is a JSON number representing the number of seconds from 1970-01-01T0:0:0Z as measured in UTC until the date/time.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// URL of the End-User's Web page or blog. This Web page SHOULD contain information published by the End-User or an organization that the End-User is affiliated with.
	Website *string `json:"website,omitempty"`
	// String from zoneinfo [zoneinfo] time zone database representing the End-User's time zone. For example, Europe/Paris or America/Los_Angeles.
	Zoneinfo *string `json:"zoneinfo,omitempty"`
}

// NewOidcUserInfo instantiates a new OidcUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcUserInfo() *OidcUserInfo {
	this := OidcUserInfo{}
	return &this
}

// NewOidcUserInfoWithDefaults instantiates a new OidcUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcUserInfoWithDefaults() *OidcUserInfo {
	this := OidcUserInfo{}
	return &this
}

// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *OidcUserInfo) GetBirthdate() string {
	if o == nil || o.Birthdate == nil {
		var ret string
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetBirthdateOk() (*string, bool) {
	if o == nil || o.Birthdate == nil {
		return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *OidcUserInfo) HasBirthdate() bool {
	if o != nil && o.Birthdate != nil {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given string and assigns it to the Birthdate field.
func (o *OidcUserInfo) SetBirthdate(v string) {
	o.Birthdate = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OidcUserInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OidcUserInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OidcUserInfo) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *OidcUserInfo) GetEmailVerified() bool {
	if o == nil || o.EmailVerified == nil {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || o.EmailVerified == nil {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *OidcUserInfo) HasEmailVerified() bool {
	if o != nil && o.EmailVerified != nil {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *OidcUserInfo) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *OidcUserInfo) GetFamilyName() string {
	if o == nil || o.FamilyName == nil {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetFamilyNameOk() (*string, bool) {
	if o == nil || o.FamilyName == nil {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *OidcUserInfo) HasFamilyName() bool {
	if o != nil && o.FamilyName != nil {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *OidcUserInfo) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *OidcUserInfo) GetGender() string {
	if o == nil || o.Gender == nil {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetGenderOk() (*string, bool) {
	if o == nil || o.Gender == nil {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *OidcUserInfo) HasGender() bool {
	if o != nil && o.Gender != nil {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *OidcUserInfo) SetGender(v string) {
	o.Gender = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *OidcUserInfo) GetGivenName() string {
	if o == nil || o.GivenName == nil {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetGivenNameOk() (*string, bool) {
	if o == nil || o.GivenName == nil {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *OidcUserInfo) HasGivenName() bool {
	if o != nil && o.GivenName != nil {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *OidcUserInfo) SetGivenName(v string) {
	o.GivenName = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *OidcUserInfo) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *OidcUserInfo) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *OidcUserInfo) SetLocale(v string) {
	o.Locale = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *OidcUserInfo) GetMiddleName() string {
	if o == nil || o.MiddleName == nil {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetMiddleNameOk() (*string, bool) {
	if o == nil || o.MiddleName == nil {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *OidcUserInfo) HasMiddleName() bool {
	if o != nil && o.MiddleName != nil {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *OidcUserInfo) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OidcUserInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OidcUserInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OidcUserInfo) SetName(v string) {
	o.Name = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *OidcUserInfo) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *OidcUserInfo) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *OidcUserInfo) SetNickname(v string) {
	o.Nickname = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OidcUserInfo) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OidcUserInfo) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *OidcUserInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneNumberVerified returns the PhoneNumberVerified field value if set, zero value otherwise.
func (o *OidcUserInfo) GetPhoneNumberVerified() bool {
	if o == nil || o.PhoneNumberVerified == nil {
		var ret bool
		return ret
	}
	return *o.PhoneNumberVerified
}

// GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetPhoneNumberVerifiedOk() (*bool, bool) {
	if o == nil || o.PhoneNumberVerified == nil {
		return nil, false
	}
	return o.PhoneNumberVerified, true
}

// HasPhoneNumberVerified returns a boolean if a field has been set.
func (o *OidcUserInfo) HasPhoneNumberVerified() bool {
	if o != nil && o.PhoneNumberVerified != nil {
		return true
	}

	return false
}

// SetPhoneNumberVerified gets a reference to the given bool and assigns it to the PhoneNumberVerified field.
func (o *OidcUserInfo) SetPhoneNumberVerified(v bool) {
	o.PhoneNumberVerified = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *OidcUserInfo) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *OidcUserInfo) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *OidcUserInfo) SetPicture(v string) {
	o.Picture = &v
}

// GetPreferredUsername returns the PreferredUsername field value if set, zero value otherwise.
func (o *OidcUserInfo) GetPreferredUsername() string {
	if o == nil || o.PreferredUsername == nil {
		var ret string
		return ret
	}
	return *o.PreferredUsername
}

// GetPreferredUsernameOk returns a tuple with the PreferredUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetPreferredUsernameOk() (*string, bool) {
	if o == nil || o.PreferredUsername == nil {
		return nil, false
	}
	return o.PreferredUsername, true
}

// HasPreferredUsername returns a boolean if a field has been set.
func (o *OidcUserInfo) HasPreferredUsername() bool {
	if o != nil && o.PreferredUsername != nil {
		return true
	}

	return false
}

// SetPreferredUsername gets a reference to the given string and assigns it to the PreferredUsername field.
func (o *OidcUserInfo) SetPreferredUsername(v string) {
	o.PreferredUsername = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *OidcUserInfo) GetProfile() string {
	if o == nil || o.Profile == nil {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetProfileOk() (*string, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *OidcUserInfo) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *OidcUserInfo) SetProfile(v string) {
	o.Profile = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *OidcUserInfo) GetSub() string {
	if o == nil || o.Sub == nil {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetSubOk() (*string, bool) {
	if o == nil || o.Sub == nil {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *OidcUserInfo) HasSub() bool {
	if o != nil && o.Sub != nil {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *OidcUserInfo) SetSub(v string) {
	o.Sub = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OidcUserInfo) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OidcUserInfo) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *OidcUserInfo) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *OidcUserInfo) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *OidcUserInfo) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *OidcUserInfo) SetWebsite(v string) {
	o.Website = &v
}

// GetZoneinfo returns the Zoneinfo field value if set, zero value otherwise.
func (o *OidcUserInfo) GetZoneinfo() string {
	if o == nil || o.Zoneinfo == nil {
		var ret string
		return ret
	}
	return *o.Zoneinfo
}

// GetZoneinfoOk returns a tuple with the Zoneinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcUserInfo) GetZoneinfoOk() (*string, bool) {
	if o == nil || o.Zoneinfo == nil {
		return nil, false
	}
	return o.Zoneinfo, true
}

// HasZoneinfo returns a boolean if a field has been set.
func (o *OidcUserInfo) HasZoneinfo() bool {
	if o != nil && o.Zoneinfo != nil {
		return true
	}

	return false
}

// SetZoneinfo gets a reference to the given string and assigns it to the Zoneinfo field.
func (o *OidcUserInfo) SetZoneinfo(v string) {
	o.Zoneinfo = &v
}

func (o OidcUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Birthdate != nil {
		toSerialize["birthdate"] = o.Birthdate
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EmailVerified != nil {
		toSerialize["email_verified"] = o.EmailVerified
	}
	if o.FamilyName != nil {
		toSerialize["family_name"] = o.FamilyName
	}
	if o.Gender != nil {
		toSerialize["gender"] = o.Gender
	}
	if o.GivenName != nil {
		toSerialize["given_name"] = o.GivenName
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.MiddleName != nil {
		toSerialize["middle_name"] = o.MiddleName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.PhoneNumberVerified != nil {
		toSerialize["phone_number_verified"] = o.PhoneNumberVerified
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.PreferredUsername != nil {
		toSerialize["preferred_username"] = o.PreferredUsername
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Sub != nil {
		toSerialize["sub"] = o.Sub
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Zoneinfo != nil {
		toSerialize["zoneinfo"] = o.Zoneinfo
	}
	return json.Marshal(toSerialize)
}

type NullableOidcUserInfo struct {
	value *OidcUserInfo
	isSet bool
}

func (v NullableOidcUserInfo) Get() *OidcUserInfo {
	return v.value
}

func (v *NullableOidcUserInfo) Set(val *OidcUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcUserInfo(val *OidcUserInfo) *NullableOidcUserInfo {
	return &NullableOidcUserInfo{value: val, isSet: true}
}

func (v NullableOidcUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


