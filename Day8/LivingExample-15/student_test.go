package main

import "testing"

func TestSave(t *testing.T) {
	stu := &student{
		Name: "stu01",
		Sex:  "man",
		Age:  10,
	}
	err := stu.Save()
	if err != nil {
		t.Fatalf("save student %v failed", err)
	}
}

func TestLoad(t *testing.T) {
	stu := &student{
		Name: "stu01",
		Sex:  "man",
		Age:  10,
	}
	err := stu.Save()
	if err != nil {
		t.Fatalf("load student %v failed", err)
	}

	stu2 := &student{}
	err = stu2.Load()
	if err != nil {
		t.Fatalf("load student2 %v failed", err)
	}

	if stu.Name != stu2.Name {
		t.Fatalf("stu1 != stu2")
	}
}
