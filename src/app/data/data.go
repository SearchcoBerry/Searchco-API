package data

type Syllabus struct {
	Title                      string   `json:"title"`
	Folder                     string   `json:"folder"`
	CourseTitle                string   `json:"courseTitle"`
	Classroom                  string   `json:"classroom"`
	Teacher                    string   `json:"teacher"`
	TargetGrade                string   `json:"targetGrade"`
	Semester                   string   `json:"semester"`
	WeekdayPeriod              string   `json:"weekdayPeriod"`
	CourseClassification       string   `json:"courseClassification"`
	CreditClassification       string   `json:"creditClassification"`
	CreditTypeOtherInformation string   `json:"creditTypeOtherInformation"`
	TypeClass                  string   `json:"typeClass"`
	Credits                    string   `json:"credits"`
	Precautions                string   `json:"precautions"`
	AdditionalInformation      string   `json:"additionalInformation"`
	ContactAboutClass          string   `json:"contactAboutClass"`
	PolicyAndRelevance         string   `json:"policyAndRelevance"`
	AimLecture                 string   `json:"aimLecture"`
	ForStudents                string   `json:"forStudents"`
	CourseGoals                string   `json:"courseGoals"`
	CourseSchedule             string   `json:"courseSchedule"`
	Theme                      []string `json:"theme"`
	Homework                   []string `json:"homework"`
	References                 string   `json:"references"`
	MeansOfLearning            string   `json:"meansOfLearning"`
	Evaluation                 string   `json:"evaluation"`
	RelatedSubjects            string   `json:"relatedSubjects"`
}

type Timetable struct {
	Gakubu     string `json:"gakubu"`
	Gakka      string `json:"gakka"`
	No         string `json:"No"`
	Niti       string `json:"niti"`
	Gen        int    `json:"gen"`
	Gakki      string `json:"gakki"`
	Kurasu     string `json:"kurasu"`
	Kamoku     string `json:"kamoku"`
	Tantou     string `json:"tantou"`
	Kyoushitsu string `json:"kyoushitsu"`
	Sou        string `json:"sou"`
	Ji         string `json:"ji"`
	Bikou      string `json:"bikou"`
}