package model

import (
	"fmt"
	"strconv"

	"onecv_tech/config"
)

type RegisterRequest struct {
	Teacher string `json:"teacher"`
	Students []string `json"students"`
}

type SuspendRequest struct {
	Student string `json:"student"`
}

type GetCommonResponse struct {
	Students []string `json:"students"`
}

type NotificationRequest struct {
	Teacher string `json:"teacher"`
	Notification string `json:"notification"`
}

type NotificationResponse struct {
	Recipients []string `json:"recipients"`
}

/*
ORM Method for adding "teacherEmail" into "teacher" table
*/
func AddTeacherIfNotExist(teacherEmail string) bool {
	db := config.Connect()
	defer db.Close()

	sql := "INSERT IGNORE INTO teacher(email) VALUES (?);"
	_, err := db.Exec(sql, teacherEmail)

	if err != nil {
			fmt.Println(err.Error())
			return false
	}
	return true
}

/*
ORM Method for adding "studentEmail" into "student" table
*/
func AddStudentIfNotExist(studentEmail string) bool {
	db := config.Connect()
	defer db.Close()

	sql := "INSERT IGNORE INTO student(email) VALUES (?);"
	_, err := db.Exec(sql, studentEmail)

	if err != nil {
			fmt.Println(err.Error())
			return false
	}
	return true
}

/*
ORM Method for adding the student emails in "studentEmails"
into "teacher_student" table
*/
func AddStudentsToTeachers(teacherEmail string, studentEmails []string) bool {
	db := config.Connect()
	defer db.Close()

	sql := "INSERT IGNORE INTO teacher_student(teacher_email, student_email) VALUES (?,?);"
	
	for _, studentEmail := range(studentEmails) {
		_, err := db.Exec(sql, teacherEmail, studentEmail)

		if err != nil {
				fmt.Println(err.Error())
				return false
		}
	}
	return true
}

/*
ORM Method to query for all common students among all teachers specified
in "teacherEmails"

Returns a slice of strings containing the student emails
*/
func FindCommonStudentsToTeachers(teacherEmails []string) (bool, []string) {
	db := config.Connect()
	defer db.Close()

	sql := "select student_email from teacher_student where teacher_email='"+teacherEmails[0]+"'";
	for index, teacherEmail := range teacherEmails {
		if index == 0 {
			continue
		}
		sql += " or teacher_email='"+teacherEmail+"'"
	}
	sql += " group by student_email having count(distinct teacher_email) = "+strconv.Itoa(len(teacherEmails))+";"
	rows, err := db.Query(sql)

	if err != nil {
			fmt.Println(err.Error())
			return false, []string{}
	}

	defer rows.Close()

	var studentEmails []string
	var studentEmail string

	for rows.Next() {
		err := rows.Scan(&studentEmail)
		if err != nil {
			fmt.Println(err)
			return false, []string{}
		}
		studentEmails = append(studentEmails, studentEmail)
	}

	return true, studentEmails
}


/*
ORM Method for adding "student" into "student_suspend_list" table
*/
func AddStudentToSuspendList(student string) bool {
	db := config.Connect()
	defer db.Close()

	sql := "INSERT IGNORE INTO student_suspend_list(student_email) VALUES (?);"
	_, err := db.Exec(sql, student)

	if err != nil {
			fmt.Println(err.Error())
			return false
	}

	return true
}

/*
ORM Method for getting all student emails tied to "teacher"
excluding those in "student_suspend_list"

Return a slice of student emails that are recipients
*/
func GetRecipientsOfTeacher(teacher string) (bool, []string) {
	db := config.Connect()
	defer db.Close()

	// check teacher students
	sql :=  "select distinct student_email from teacher_student where teacher_email=? and student_email not in (select student_email from student_suspend_list)";
	rows, err := db.Query(sql, teacher)

	if err != nil {
		fmt.Println(err.Error())
		return false, []string{}
	}

	defer rows.Close()

	var studentEmails []string
  var studentEmail string

	for rows.Next() {
		err := rows.Scan(&studentEmail)
		if err != nil {
			fmt.Println(err)
			return false, []string{}
		}
		studentEmails = append(studentEmails, studentEmail)
	}

	return true, studentEmails
}