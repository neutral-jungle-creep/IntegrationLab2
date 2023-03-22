package domain

type Faculty struct {
	Abbr       string  `json:"abbr"`
	Code       *string `json:"code"`
	FacultyOid int     `json:"facultyOid"`
	Institute  string  `json:"institute"`
	Name       string  `json:"name"`
}

type Group struct {
	ChairOid        int    `json:"chairOid"`
	Course          int    `json:"course"`
	Faculty         string `json:"faculty"`
	FacultyOid      int    `json:"facultyOid"`
	FormOfEducation string `json:"formOfEducation"`
	GroupOid        int    `json:"groupOid"`
	Number          string `json:"number"`
	Speciality      string `json:"speciality"`
}

type Lesson struct {
	Auditorium    string  `json:"auditorium"`
	AuditoriumOid int     `json:"auditoriumOid"`
	BeginLesson   string  `json:"beginLesson"`
	Building      string  `json:"building"`
	Date          string  `json:"date"`
	DayOfWeek     int     `json:"dayOfWeek"`
	Discipline    string  `json:"discipline"`
	EndLesson     string  `json:"endLesson"`
	Group         string  `json:"group"`
	GroupOid      int     `json:"groupOid"`
	KindOfWork    string  `json:"kindOfWork"`
	Lecturer      string  `json:"lecturer"`
	LecturerOid   int     `json:"lecturerOid"`
	Stream        *string `json:"stream"`
	StreamOid     int     `json:"streamOid"`
	SubGroup      *string `json:"subGroup"`
	SubGroupOid   int     `json:"subGroupOid"`
}
