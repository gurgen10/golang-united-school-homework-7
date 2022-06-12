package coverage

import (
	"os"
	"sort"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// People testing

func TestLenOk(t *testing.T) {
	var p People
	p = append(p, Person{"Gug", "Tor", time.Now()})
	p = append(p, Person{"Gug", "Tor", time.Now()})
	var len = 2

	r := p.Len()

	if r != 2 {
		t.Errorf("Expected %d got %d", len, r)
	}
}

func TestLessOk(t *testing.T) {
	var p People
	name1, name2, name3 := "aHay", "bKar", "cGug"
	p = append(p, Person{name3, "Tor", time.Now().Add(-time.Second * 10)})
	p = append(p, Person{name1, "Hayrapet", time.Now()})
	p = append(p, Person{name2, "gHak", time.Now()})
	p = append(p, Person{name2, "fHakob", time.Now()})
	p = append(p, Person{name2, "aHakob", time.Now()})
	p = append(p, Person{name3, "aHakob", time.Now()})

	sort.Sort(p)

	if p[0].firstName != name1  {
		t.Errorf("Expected %s got %s",name1, p[0].firstName)
	}
	if p[1].firstName == name2 && p[1].lastName != "aHakob"  {
		t.Errorf("Expected \"aHakob\" got %s", p[1].lastName)

	}
	if p[2].firstName == name2  && p[2].lastName != "fHakob"  {
		t.Errorf("Expected \"fHakob\" got %s",p[2].lastName)
	}
	if p[3].firstName == name2  && p[3].lastName != "gHak"  {
		t.Errorf("Expected \"gHak\" got %s",p[3].lastName)
	}
	if p[4].firstName != name3  {
		t.Errorf("Expected %s got %s",name3, p[4].firstName)
	}
	if p[5].birthDay.Unix() > p[4].birthDay.Unix()   {
		t.Errorf("Expected %s: wrong sorting", p[5].firstName)
	}
}

// Matrix testing
var mStr = `5 8 9
7 8 10
89 6 3`

func TestNewOk(t *testing.T) {
	m, err := New(mStr)

	if err != nil {
		t.Errorf("Unexpected error on creating new instance: \"%s\"", err)
	}

	if m.rows != 3 {
		t.Error("Rows must be equal 3")
	}
	if m.cols != 3 {
		t.Error("Cols must be equal 3")
	}
	if m.cols != m.rows {
		t.Error("Cols must be equal rows")
	}
}
func TestNewErr(t *testing.T) {
	str := "dsdsdsdsds"
	_, err := New(str)

	if(err == nil) {
		t.Errorf("Wrong Argument \"%s\"", str)
	}
}
func TestNewWrongMatrix(t *testing.T) {
	var wStr = `5 8 9 6
		7 8 10
		89 6 3`
	_, err := New(wStr)

	if(err == nil) {
		t.Errorf("Wrong matrix: \"%s\"", err)
	}
}
func TestRows(t *testing.T) {
	m, _ := New(mStr)
	rows := m.Rows()

	if(rows[1][2] != 10) {
		t.Errorf("Expected 10 got \"%v\"", rows[1][1])
	}
}

func TestCols(t *testing.T) {
	m, _ := New(mStr)
	cols := m.Cols()

	if(cols[1][2] != 6) {
		t.Errorf("Expected 6 got \"%v\"", cols[1][2])
	}
}


func TestSetOk(t *testing.T) {
	m, _ := New(mStr)
	g := m.Set(1,1,99)

	if g && m.data[2 + 2] != 99 {
		t.Errorf("Expected 99 got %v", m.data[2 + 2])
	}
}
func TestSetErr(t *testing.T) {
	m, _ := New(mStr)
	err1 := m.Set(3,1,99)

	if err1 {
		t.Error("Row argument cant be equal or greeter than matrix rows")
	}
	err2 := m.Set(1,3,99)

	if err2 {
		t.Error("Col argument cant be equal or greeter than matrix cols")
	}
}
