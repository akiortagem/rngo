package rngo

import "testing"

func TestRollNDice(t *testing.T) {
	//test to see if the function returns the correct number of dice
	//and that the dice are within the correct range
	rolls := rollNDice(5, 6) //5d6
	if len(rolls) != 5 {
		t.Errorf("Expected 5 dice, got %d", len(rolls))
	}
	for _, roll := range rolls {
		if roll < 1 || roll > 6 {
			t.Errorf("Expected roll between 1 and 6, got %d", roll)
		}
	}
}

func TestSelectAdvantage(t *testing.T) {
	//test to see of the function returns the correct number of rolls
	//and that they select the correct rolls
	rolls := []int{1, 2, 3, 4, 5, 6}
	_, adv := selectAdvantage(rolls, 2)

	if len(adv) != 2 {
		t.Errorf("Expected 2 dice, got %d", len(adv))
	}

	//asserts that the returned rolls are the highest rolls
	if adv[0] != 6 || adv[1] != 5 {
		t.Errorf("Expected [6, 5], got %v", adv)
	}

	//test for disadvantage
	_, dis := selectAdvantage(rolls, -2)

	if len(dis) != 2 {
		t.Errorf("Expected 2 dice, got %d", len(dis))
	}

	//asserts that the returned rolls are the lowest rolls, not necessarily in order

	if (dis[0] + dis[1]) != 3 {
		t.Errorf("Expected [1, 2], got %v", dis)
	}

	// test early return
	_, adv = selectAdvantage(rolls, 0)

	if len(adv) != 6 {
		t.Errorf("Expected 6 dice, got %d", len(adv))
	}

	//asserts that the returned rolls are the original rolls, by summing it
	sum := 0
	for _, roll := range adv {
		sum += roll
	}
	if sum != 21 {
		t.Errorf("Expected [1, 2, 3, 4, 5, 6], got %v", adv)
	}
}

func TestExplode(t *testing.T) {
	rolls := []int{4, 6, 3}

	original, result := explode(rolls, 6, 6)

	if len(result) < 1 {
		t.Errorf("Expected at least 1 rolls but got %d", len(result))
	}

	if original[0] != 4 || original[1] != 6 || original[2] != 3 {
		t.Errorf("Expected to retain original rolls but got %v", original)
	}
}
