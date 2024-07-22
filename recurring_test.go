package recurring

import (
	"testing"
	"time"
)

func parseTime(value string) time.Time {
	t, _ := time.Parse("2006-01-02", value)
	return t
}

func TestRuleSetGenerator(t *testing.T) {
	tests := []struct {
		name   string
		input  RecurringInput
		expect string
	}{
		{
			name: "Recurring Daily; Repeat every 1 day(s); Run until it reaches 36 occurrences;",
			input: RecurringInput{
				Frequency: "daily",
				Count:     36,
				Interval:  1,
			},
			expect: "RRULE:FREQ=DAILY;INTERVAL=1;COUNT=36",
		},
		{
			name: "Recurring Weekly; Repeat every 1 week(s) on FR ; Never stop;",
			input: RecurringInput{
				Frequency: "weekly",
				WeekDays:  []string{"FR"},
				Interval:  1,
			},
			expect: "RRULE:FREQ=WEEKLY;INTERVAL=1;BYDAY=FR",
		},
		{
			name: "Recurring Monthly; Repeat every 1 month(s) on the First; Run until: 2050-07-21;",
			input: RecurringInput{
				Frequency: "monthly",
				Until:     parseTime("2050-07-21"),
				Interval:  1,
			},
			expect: "RRULE:FREQ=MONTHLY;INTERVAL=1;BYMONTHDAY=1;UNTIL=2050-07-21",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set, _ := RuleSetGenerator(tt.input)
			if tt.expect != set.String() {
				t.Errorf("expect %s got %s", tt.expect, set.String())
				t.Errorf("Test case %s failed", tt.name)
			}
		})
	}
}
