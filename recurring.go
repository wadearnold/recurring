package recurring

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/teambition/rrule-go"
)

type RecurringInput struct {
	// Frequency of the recurring event. Possible values are daily, weekly, monthly, yearly
	Frequency string    `json:"frequency"`
	Until     time.Time `json:"until,omitempty"`
	// Count number of occurrences to be generated
	Count    int      `json:"count,omitempty"`
	WeekDays []string `json:"weekDays,omitempty"`
	// Interval between each frequency iteration. For example, if the interval is 2, the frequency of monthly will be every 2 months
	Interval int `json:"interval,omitempty"`
	Month    int `json:"month,omitempty"`
	Pos      int `json:"pos,omitempty"`
	Day      int `json:"day,omitempty"`
}

func (ri RecurringInput) MarshalJSON() ([]byte, error) {
	type Alias RecurringInput
	return json.Marshal(&struct {
		Until string `json:"until,omitempty"`
		*Alias
	}{
		Until: ri.Until.Format("2006-01-02"),
		Alias: (*Alias)(&ri),
	})
}

// RuleGenerator generates a RRule based on RecurringInput
func RuleSetGenerator(input RecurringInput) (rrule.Set, error) {
	set := rrule.Set{}
	ruleOption := rrule.ROption{}

	// Frequency of the recurring event
	switch input.Frequency {
	case "daily":
		ruleOption.Freq = rrule.DAILY
	case "weekly":
		ruleOption.Freq = rrule.WEEKLY
	case "monthly":
		ruleOption.Freq = rrule.MONTHLY
	case "yearly":
		ruleOption.Freq = rrule.YEARLY
	default:
		// todo: handle error properly
		log.Println("Invalid frequency:" + input.Frequency)
	}

	// Count of the recurring event
	if input.Count > 0 {
		ruleOption.Count = input.Count
	}

	// Interval between each frequency iteration
	if input.Interval > 0 {
		ruleOption.Interval = input.Interval
	}

	// Until date of the recurring event
	if !input.Until.IsZero() {
		ruleOption.Until = input.Until
	}

	// Weekdays of the recurring event
	if len(input.WeekDays) > 0 {
		ruleOption.Byweekday = make([]rrule.Weekday, len(input.WeekDays))
		for i, day := range input.WeekDays {
			switch day {
			case "MO":
				ruleOption.Byweekday[i] = rrule.MO
			case "TU":
				ruleOption.Byweekday[i] = rrule.TU
			case "WE":
				ruleOption.Byweekday[i] = rrule.WE
			case "TH":
				ruleOption.Byweekday[i] = rrule.TH
			case "FR":
				ruleOption.Byweekday[i] = rrule.FR
			case "SA":
				ruleOption.Byweekday[i] = rrule.SA
			case "SU":
				ruleOption.Byweekday[i] = rrule.SU
			default:
				log.Println("Invalid weekday: " + day)
			}
		}
	}

	// Month of the recurring event
	if input.Month > 0 {
		ruleOption.Bymonth = []int{input.Month}
	}

	// Position of the recurring event
	if input.Pos > 0 {
		ruleOption.Bysetpos = []int{input.Pos}
	}

	// Day of the recurring event
	if input.Day > 0 {
		ruleOption.Bymonthday = []int{input.Day}
	}

	r, err := rrule.NewRRule(ruleOption)
	if err != nil {
		log.Println("Error creating RRule: " + err.Error())
	}
	set.RRule(r)

	return set, nil
}

type RecurringOutput struct {
	RRule       string         `json:"rRule"`
	Occurrences []time.Time    `json:"occurrences"`
	Recurring   RecurringInput `json:"recurring"`
}

func RecurringJSON(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body
	var input RecurringInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Process the JSON input (you can add your logic here)
	log.Printf("Received JSON: %+v\n", input)

	// Create a new RRule Builder
	set, _ := RuleSetGenerator(input)

	recurringOutput := RecurringOutput{
		RRule:       set.String(),
		Occurrences: set.All(),
		Recurring:   input}

	// Send a response (optional)
	//response := map[string]string{"message": "JSON received successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&recurringOutput)
}
