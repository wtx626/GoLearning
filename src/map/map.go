package main

import "fmt"

type GrantClassInfo struct {
	SchoolName string `json:"school_name"`
	GradeName  string `json:"grade_name"`
	ClassName  string `json:"class_name"`
}

func main() {
	classUsers := []GrantClassInfo{{
		SchoolName: "1",
		GradeName:  "2",
		ClassName:  "3",
	}, {
		SchoolName: "1",
		GradeName:  "2",
		ClassName:  "3",
	}, {
		SchoolName: "zhangsan",
		GradeName:  "lisi",
		ClassName:  "aa",
	}, {
		SchoolName: "zhangsan",
		GradeName:  "lisi",
		ClassName:  "aa",
	}}

	classsMap := map[GrantClassInfo]struct{}{}

	for _, classUser := range classUsers {
		classsMap[GrantClassInfo{
			SchoolName: classUser.SchoolName,
			GradeName:  classUser.GradeName,
			ClassName:  classUser.ClassName,
		}] = struct{}{}
	}

	for k := range classsMap {
		fmt.Print(k)
	}
}
