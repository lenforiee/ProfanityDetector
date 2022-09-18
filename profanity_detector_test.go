package profanitydetector_test

import (
	"testing"

	profanitydetector "github.com/lenforiee/ProfanityDetector"
)

func TestNormalList(t *testing.T) {
	filter := profanitydetector.NewProfanityFilter()
	got := filter.IsProfanity("dickhead")
	want := true

	if got != want {
		t.Errorf("ProfanityDetectorTest (TestNormalList): got %t want %t", got, want)
	}
}

func TestCleanList(t *testing.T) {
	filter := profanitydetector.NewCleanProfanityFilter()
	got := filter.IsProfanity("dickhead")
	want := false
	if got != want {
		t.Errorf("ProfanityDetectorTest (TestCleanList): got %t want %t", got, want)
	}
}

func TestAddOne(t *testing.T) {
	filter := profanitydetector.NewProfanityFilter()
	filter.AddOne("test_string")

	got := filter.IsProfanity("test_string")
	want := true

	if got != want {
		t.Errorf("ProfanityDetectorTest (TestAddOne): got %t want %t", got, want)
	}
}

func TestAddList(t *testing.T) {
	filter := profanitydetector.NewProfanityFilter()
	filter.AddList([]string{"test_string", "test_string2"})

	got := filter.IsProfanity("test_string")
	got2 := filter.IsProfanity("test_string")
	want := true
	want2 := true

	if got != want {
		t.Errorf("ProfanityDetectorTest (TestAddList): got %t, %t want %t, %t", got, got2, want, want2)
	}
}

func TestRemoveOne(t *testing.T) {
	filter := profanitydetector.NewProfanityFilter()
	filter.RemoveOne("dickhead")

	got := filter.IsProfanity("dickhead")
	want := false

	if got != want {
		t.Errorf("ProfanityDetectorTest (TestRemoveOne): got %t want %t", got, want)
	}
}

func TestRemoveList(t *testing.T) {
	filter := profanitydetector.NewProfanityFilter()
	filter.RemoveList([]string{"dickhead", "dickface"})

	got := filter.IsProfanity("dickhead")
	got2 := filter.IsProfanity("dickface")
	want := false
	want2 := false

	if got != want {
		t.Errorf("ProfanityDetectorTest (TestRemoveList): got %t, %t want %t, %t", got, got2, want, want2)
	}
}
