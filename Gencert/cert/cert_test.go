package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-12-25")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert shoul be a valid reference. got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. excepted='GOLANG COURSE', got='%v'", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "John", "2018-12-25")
	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "fhejfbejdb jvbjvbjbvjbjedv bhbvhebrjvbjebvjbeivbjebvjebfvjbj vhebv ekhbvkne rvhj nv jher vhje "
	_, err := New(course, "John", "2018-12-25")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2018-12-25")
	if err == nil {
		t.Errorf("Error should be returned on an empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "dzeihduezbbdezjbezcneznvncjezbvcijezbvjbezjivbz ve jze vjez vjezvj zje vjz evjz ezvj zevzej"
	_, err := New("Golang", name, "2018-12-25")
	if err == nil {
		t.Errorf("Error should be returned on a too long name (name=%s)", name)
	}
}

func TestDateFormat(t *testing.T) {
	date := "2018-12-25"
	_, err := New("Golang", "John", date)
	if err != nil {
		t.Errorf("Error with the date %v", err)
	}
}
