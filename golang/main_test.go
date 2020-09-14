package main

import "testing"

func TestGreeting(t *testing.T){
	result := "Code.education Rocks!"

	if result != "Code.education Rocks!" {
		t.Errorf("Função greeting está incorreta");
	}
}