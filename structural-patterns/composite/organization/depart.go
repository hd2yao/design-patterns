package main

import (
	"fmt"
	"strings"
)

// HRDOrg 人力资源部门
type HRDOrg struct {
	orgName string
	depth   int
}

func (o *HRDOrg) display() {
	if o == nil {
		return
	}
	fmt.Println(strings.Repeat("-", o.depth*2), " ", o.orgName)
}

func (o *HRDOrg) duty() {
	if o == nil {
		return
	}
	fmt.Println(o.orgName, "员工招聘培训管理")
}

// FinanceOrg 财务部门
type FinanceOrg struct {
	orgName string
	depth   int
}

func (f *FinanceOrg) display() {
	if f == nil {
		return
	}
	fmt.Println(strings.Repeat("-", f.depth*2), " ", f.orgName)
}

func (f *FinanceOrg) duty() {
	if f == nil {
		return
	}
	fmt.Println(f.orgName, "员工招聘培训管理")
}
