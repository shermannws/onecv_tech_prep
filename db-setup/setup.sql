CREATE DATABASE onecv_tech;

USE onecv_tech;

CREATE TABLE teacher (
`email` VARCHAR(50) NOT NULL PRIMARY KEY
);

CREATE TABLE student (
`email` VARCHAR(50) NOT NULL PRIMARY KEY
);

CREATE TABLE teacher_student (
`teacher_email` VARCHAR(50) NOT NULL,
`student_email` VARCHAR(50) NOT NULL
);

CREATE TABLE student_suspend_list (
`student_email` VARCHAR(50) NOT NULL
);