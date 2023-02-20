package test

import (
	"bytes"
	"strings"
	"testing"
	"net/http"
	"net/http/httptest"
	_ "github.com/go-sql-driver/mysql"

	"onecv_tech/controller"
)

func TestRegisterStudentsToTeacherSuccess(t *testing.T) {
	var jsonStr = []byte(`{"teacher":"teacher1@gmail.com","students":["student1@gmail.com","student2@gmail.com"]}`)

	req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.RegisterStudentsToTeacher)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestRegisterStudentsToTeacherFailure(t *testing.T) {
	var jsonStr = []byte(`{"TYPO":"teacher1@gmail.com","students":["student1@gmail.com","student2@gmail.com"]}`)

	req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.RegisterStudentsToTeacher)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestGetCommonStudentsOfTeacherSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/commonstudents?teacher=teacher1@gmail.com&teacher=teacher2@gmail.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetCommonStudentsOfTeachers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"students":["student1@gmail.com","student2@gmail.com"]}\n`
	if strings.Compare(rr.Body.String(), expected) == 1 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetCommonStudentsOfTeacherFailure(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/commonstudents", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetCommonStudentsOfTeachers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"message":"No 'teacher' in URL parameters"}\n`
	if strings.Compare(rr.Body.String(), expected) == 1 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSuspendStudentSuccess(t *testing.T) {
	var jsonStr = []byte(`{"student":"student2@gmail.com"}`)

	req, err := http.NewRequest("POST", "/api/suspend", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SuspendStudent)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestSuspendStudentFailure(t *testing.T) {
	var jsonStr = []byte(`{"TYPO":"student2@gmail.com"}`)

	req, err := http.NewRequest("POST", "/api/suspend", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SuspendStudent)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestRetrieveRecipientListrSuccess(t *testing.T) {
	var jsonStr = []byte(`{"teacher":"teacher2@gmail.com","notification":"Hellostudents! @random1@gmail.com"}\n`)
	
	req, err := http.NewRequest("POST", "/api/retrievefornotifications", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.RetrieveRecipientList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"recipients":["student1@gmail.com","student3@gmail.com","random1@gmail.com"]}\n`
	if strings.Compare(rr.Body.String(), expected) == 1 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestRetrieveRecipientListrFailure(t *testing.T) {
	var jsonStr = []byte(`{"TYPO":"teacher2@gmail.com","notification":"Hellostudents! @random1@gmail.com"}\n`)
	
	req, err := http.NewRequest("POST", "/api/retrievefornotifications", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.RetrieveRecipientList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}