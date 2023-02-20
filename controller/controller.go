package controller
 
import (
    "encoding/json"
		"strings"
    "net/http"
 
    "onecv_tech/model"
)

var SERVER_ERROR_MSG = "Internal Server Error, please try again."

type Response struct {
	Message string `json:"message"`
}

/*
Wrapper function to return HTTP error response messages
*/
func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

/*
Wrapper function to return HTTP success response messages
*/
func JSONSuccess(w http.ResponseWriter, resp interface{}, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}


// API 1: [POST] Register one or more students to a specified teacher
func RegisterStudentsToTeacher(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	var isSuccess bool

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)

	if err != nil {
		JSONError(w, &Response{err.Error()}, http.StatusBadRequest)
		return
	}

	isSuccess = model.AddTeacherIfNotExist(req.Teacher)
	if !isSuccess {
		JSONError(w, &Response{SERVER_ERROR_MSG}, http.StatusInternalServerError)
		return
	}
	for _, student := range req.Students {
		isSuccess = model.AddStudentIfNotExist(student)
		if !isSuccess {
			JSONError(w, &Response{SERVER_ERROR_MSG}, http.StatusInternalServerError)
			return
		}
	}
	isSuccess = model.AddStudentsToTeachers(req.Teacher, req.Students)
	if !isSuccess {
		JSONError(w, &Response{SERVER_ERROR_MSG}, http.StatusInternalServerError)
		return
	}
  
	JSONSuccess(w, &Response{}, http.StatusNoContent)
	return;
}

// API 2: [GET] Retrieve a list of students common to a given list of teachers
func GetCommonStudentsOfTeachers(w http.ResponseWriter, r *http.Request) {
	var isSuccess bool

	teacherEmails := r.URL.Query()["teacher"]
	if teacherEmails == nil {
		JSONError(w, &Response{"No 'teacher' in URL parameters"}, http.StatusBadRequest)
		return
	}
	
	isSuccess, studentEmails := model.FindCommonStudentsToTeachers(teacherEmails)
	if !isSuccess {
		JSONError(w, &Response{SERVER_ERROR_MSG}, http.StatusInternalServerError)
		return
	}

	var response model.GetCommonResponse

	response.Students = studentEmails 
	JSONSuccess(w, response, http.StatusOK)
	return;
}

// API 3: [POST] Suspend a specified student
func SuspendStudent(w http.ResponseWriter, r *http.Request) {
	var req model.SuspendRequest
	var isSuccess bool

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)

	if err != nil {
		JSONError(w, &Response{err.Error()}, http.StatusBadRequest)
		return
	}

	isSuccess = model.AddStudentToSuspendList(req.Student)
	if !isSuccess {
		JSONError(w, &Response{SERVER_ERROR_MSG}, http.StatusInternalServerError)
		return
	}

	JSONSuccess(w, &Response{}, http.StatusNoContent)
	return;
}

// API 4: [POST] Retrieve a list of students who can receive a given notification
func RetrieveRecipientList(w http.ResponseWriter, r *http.Request) {
	var req model.NotificationRequest
	var isSuccess bool

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)

	if err != nil {
		JSONError(w, &Response{err.Error()}, http.StatusBadRequest)
		return
	}

	var recipients []string

	isSuccess, recipients = model.GetRecipientsOfTeacher(req.Teacher)
	if !isSuccess {
		JSONError(w, &Response{SERVER_ERROR_MSG}, http.StatusInternalServerError)
		return
	}
	
	otherRecipients := getOtherRecipients(req.Notification)
	recipients = append(recipients, otherRecipients...)

	var response model.NotificationResponse
	response.Recipients = recipients

	JSONSuccess(w, response, http.StatusOK)
	return;
}

/* 
Helper function for API4
Parse the body for emails that are mentioned

Returns a list of emails that are mentioned in a "notification"
*/
func getOtherRecipients(notification string) []string {
	var result []string

	string_arr := strings.Fields(notification)
	for _, word := range string_arr {
		if string(word[0]) == "@" {
			result = append(result, word[1:])
		}
	}
	return result
}