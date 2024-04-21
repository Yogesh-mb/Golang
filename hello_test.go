package hello

import "testing"

func TestSayHello(t *testing.T) {
	subTests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Yogesh"},
			result: "Hello, Yogesh!",
		},
		{
			items:  []string{"Yogesh", "Badagi"},
			result: "Hello, Yogesh, Badagi!",
		},
	}

	for _, st := range subTests {
		if s := Hey(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}

	want := "Hello, test!"
	got := Say("test!")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}

	newGot := Hey([]string{"test"})

	if want != newGot {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
