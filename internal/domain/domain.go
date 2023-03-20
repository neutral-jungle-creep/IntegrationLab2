package domain

type Faculty struct {
	Abbr       string  `json:"abbr"`
	Code       *string `json:"code"`
	FacultyOid int     `json:"facultyOid"`
	Institute  string  `json:"institute"`
	Name       string  `json:"name"`
}
