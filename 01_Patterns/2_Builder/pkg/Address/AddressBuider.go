package Address

type addressBuilder struct {
	addres *Address
}

func NewAddressBuilder() *addressBuilder {
	return &addressBuilder{&Address{}}
}

func (a *addressBuilder) Build() *Address {
	return a.addres
}

func (a *addressBuilder) SetCountry(country string) *addressBuilder {
	a.addres.country = "Country: " + country
	return a
}
func (a *addressBuilder) SetRegion(Region string) *addressBuilder {
	a.addres.region = "Region: " + Region
	return a
}
func (a *addressBuilder) SetCity(City string) *addressBuilder {
	a.addres.city = "City: " + City
	return a
}
func (a *addressBuilder) SetPost(Post string) *addressBuilder {
	a.addres.post = "Post: " + Post
	return a
}
func (a *addressBuilder) SetHome(Home string) *addressBuilder {
	a.addres.home = "Home: " + Home
	return a
}
func (a *addressBuilder) SetUserName(UserName string) *addressBuilder {
	a.addres.userName = "UserName: " + UserName
	return a
}
func (a *addressBuilder) SetPhone(Phone string) *addressBuilder {
	a.addres.phone = "Phone: " + Phone
	return a
}
