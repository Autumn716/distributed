package portal

import "html/template"

// 加载student 和 students 两个HTML

var rootTemplate *template.Template

func ImportTemplates() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"../../portal/students.html",
		"../../portal/student.html",
	)
	if err != nil {
		return err
	}
	return nil
}
