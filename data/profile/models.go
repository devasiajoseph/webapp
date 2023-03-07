package profile

type Profile struct {
	ProfileID  int    `json:"profile_id" db:"profile_id"`
	FullName   string `json:"full_name" db:"full_name"`
	CountryID  int    `json:"country_id" db:"country_id"`
	About      string `json:"about" db:"about"`
	ProfilePic string `json:"profile_pic" db:"profile_pic"`
}

func (p *Profile) Create() {

}
