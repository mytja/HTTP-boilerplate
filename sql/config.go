package sql

type Config struct {
	DatabaseName       string   `json:"database_name"`
	DatabaseConfig     string   `json:"database_config"`
	Debug              bool     `json:"debug"`
	Host               string   `json:"host"`
	CommitHash         string   `json:"commit_hash"`
	RemoteRepository   string   `json:"remote_repository"`
	SchoolName         string   `json:"school_name"`
	SchoolAddress      string   `json:"school_address"`
	SchoolCity         string   `json:"school_city"`
	SchoolCountry      string   `json:"school_country"`
	SchoolPostCode     int      `json:"school_post_code"`
	ParentViewGrades   bool     `json:"parent_view_grades"`
	ParentViewAbsences bool     `json:"parent_view_absences"`
	ParentViewHomework bool     `json:"parent_view_homework"`
	ParentViewGradings bool     `json:"parent_view_gradings"`
	BlockRegistrations bool     `json:"block_registrations"`
	BlockMeals         bool     `json:"block_meals"`
	SchoolFreeDays     []string `json:"school_free_days"`
}
