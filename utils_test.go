package main

import (
	"os"
	"testing"
)

var isBodyFileTestScenarios = []struct {
	in  string
	out bool
}{
	{"<br>", false},
	{"/c/s/d.txt", false},
	{"test_files/test1.txt", true},
	{"%#a", false},
	{"test_files/dir2/test1 copy.txt", true},
	{os.Getenv("mail_log"), true},
	{os.Getenv("mail_config"), true},
	{"$", false},
}

func TestIsBodyFile(t *testing.T) {
	for _, tt := range isBodyFileTestScenarios {
		t.Run(tt.in, func(t *testing.T) {
			s := isBodyFile(tt.in)
			if s != tt.out {
				t.Errorf("got %t, want %t", s, tt.out)
			}
		})
	}
}
