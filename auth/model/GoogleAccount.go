package model

type GoogleAccountModel struct {
	Id int64
	IdToken string

	Email string

	FullName string
	GivenName string
	FamilyName string
	ImageUrl string
}
