package rave

import (
	"fmt"
	"testing"
)

var s = Subscription{r}

func TestSubscription_Activate(t *testing.T) {
	ids := [2]int{1264, 1152}

	for _, id := range ids {
		err, response := s.Activate(id)
		fmt.Println(response)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Plan Activation Wasn't Successful %v", response)
		}
	}
}

func TestSubscription_List(t *testing.T) {
	err, response := s.List()
	fmt.Println(response)
	if err != nil {
		t.Fatalf("An error occurred while testing single transfer: %v", err)
	}
	if response["status"] != "success" {
		t.Fatalf("Could not get subscriptions %v", response)
	}
}

func TestSubscription_Fetch(t *testing.T) {
	subscriptionIds := [2]string{"1264", "1152"}
	for _, id := range subscriptionIds {
		err, response := s.Fetch(id)
		fmt.Println(response)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Could not fetch subscription %v", response)
		}
	}
}

func TestSubscription_Cancel(t *testing.T) {
	subscriptionIds := [2]int{1264, 1152}
	for _, id := range subscriptionIds {
		err, response := s.Cancel(id)
		fmt.Println(response)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Could not cancel subscription%v", response)
		}
	}
}
