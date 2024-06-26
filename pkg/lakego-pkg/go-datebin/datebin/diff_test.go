package datebin

import (
	"fmt"
	"testing"
)

func Test_Diff_Seconds(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-06-05 01:35:00",
			check: -157212,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-06-06 11:17:32",
			check: 50540,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				Seconds()

			eq(check, td.check, "failed Diff_Seconds")
		})
	}
}

func Test_Diff_SecondsAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-06-05 01:35:00",
			check: 157212,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-06-06 11:17:32",
			check: 50540,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				SecondsAbs()

			eq(check, td.check, "failed Diff SecondsAbs")
		})
	}
}

func Test_Diff_Minutes(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-06-05 01:35:00",
			check: -2620,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-06-06 11:17:32",
			check: 842,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				Minutes()

			eq(check, td.check, "failed Diff Minutes")
		})
	}
}

func Test_Diff_MinutesAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-06-05 01:35:00",
			check: 2620,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-06-06 11:17:32",
			check: 842,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				MinutesAbs()

			eq(check, td.check, "failed Diff MinutesAbs")
		})
	}
}

func Test_Diff_Hours(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-06-05 01:35:00",
			check: -43,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-06-06 12:17:32",
			check: 15,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				Hours()

			eq(check, td.check, "failed Diff MinutesAbs")
		})
	}
}

func Test_Diff_HoursAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-06-05 01:35:00",
			check: 43,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-06-06 12:17:32",
			check: 15,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				HoursAbs()

			eq(check, td.check, "failed Diff HoursAbs")
		})
	}
}

func Test_Diff_Days(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-05-05 01:35:00",
			check: -32,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-06-12 12:17:32",
			check: 6,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				Days()

			eq(check, td.check, "failed Diff Days")
		})
	}
}

func Test_Diff_DaysAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-05-05 01:35:00",
			check: 32,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-06-12 12:17:32",
			check: 6,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				DaysAbs()

			eq(check, td.check, "failed Diff DaysAbs")
		})
	}
}

func Test_Diff_Weeks(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-03-05 01:35:00",
			check: -13,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-07-12 12:17:32",
			check: 5,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				Weeks()

			eq(check, td.check, "failed Diff Weeks")
		})
	}
}

func Test_Diff_WeeksAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-03-05 01:35:00",
			check: 13,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2024-07-12 12:17:32",
			check: 5,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				WeeksAbs()

			eq(check, td.check, "failed Diff WeeksAbs")
		})
	}
}

func Test_Diff_Months(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-02-05 01:35:00",
			check: -5,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2025-07-12 12:17:32",
			check: 13,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				Months()

			eq(check, td.check, "failed Diff Months")
		})
	}
}

func Test_Diff_MonthsAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2024-02-05 01:35:00",
			check: 5,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2025-07-12 12:17:32",
			check: 13,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				MonthsAbs()

			eq(check, td.check, "failed Diff MonthsAbs")
		})
	}
}

func Test_Diff_Years(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2021-02-05 01:35:00",
			check: -3,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2026-07-12 12:17:32",
			check: 2,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				Years()

			eq(check, td.check, "failed Diff Years")
		})
	}
}

func Test_Diff_YearsAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check int64
	}{
		{
			index: "index-1",
			date1: "2024-06-06 21:15:12",
			date2: "2021-02-05 01:35:00",
			check: 3,
		},
		{
			index: "index-2",
			date1: "2024-06-05 21:15:12",
			date2: "2026-07-12 12:17:32",
			check: 2,
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				YearsAbs()

			eq(check, td.check, "failed Diff YearsAbs")
		})
	}
}

func Test_Diff_Format(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index  string
		date1  string
		date2  string
		format string
		check  string
	}{
		{
			index:  "index-1",
			date1:  "2024-06-05 21:15:12",
			date2:  "2026-07-12 12:17:32",
			format: "diff {Y} years",
			check:  "diff 2 years",
		},
		{
			index:  "index-2",
			date1:  "2024-06-05 21:15:12",
			date2:  "2024-08-12 12:17:32",
			format: "diff {m} Months",
			check:  "diff 2 Months",
		},
		{
			index:  "index-3",
			date1:  "2024-06-05 21:15:12",
			date2:  "2024-08-12 12:17:32",
			format: "diff {d} Days",
			check:  "diff 67 Days",
		},
		{
			index:  "index-4",
			date1:  "2024-06-05 21:15:12",
			date2:  "2024-06-05 12:17:32",
			format: "diff {H} Hours",
			check:  "diff -8 Hours",
		},
		{
			index:  "index-5",
			date1:  "2024-06-05 11:15:12",
			date2:  "2024-06-05 12:17:32",
			format: "diff {i} Minutes",
			check:  "diff 62 Minutes",
		},
		{
			index:  "index-6",
			date1:  "2024-06-05 11:15:12",
			date2:  "2024-06-05 12:17:32",
			format: "diff {s} Seconds",
			check:  "diff 3740 Seconds",
		},
		{
			index:  "index-7",
			date1:  "2024-06-05 21:15:12",
			date2:  "2024-08-12 12:17:32",
			format: "diff {w} Weeks",
			check:  "diff 9 Weeks",
		},
		{
			index:  "index-8",
			date1:  "2024-06-05 21:15:12",
			date2:  "2024-08-12 12:17:32",
			format: "diff {WW} Weeks {DD} days",
			check:  "diff 9 Weeks 4 days",
		},
		{
			index:  "index-9",
			date1:  "2024-06-05 21:15:12",
			date2:  "2024-08-12 12:17:32",
			format: "diff {dd} days {HH} Hour {ii} Minute {ss} Second",
			check:  "diff 67 days 15 Hour 2 Minute 20 Second",
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				Format(td.format)

			eq(check, td.check, "failed Diff Format")
		})
	}
}

func Test_Diff_DurationBetween(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check string
	}{
		{
			index: "index-1",
			date1: "2023-06-06 21:15:12",
			date2: "2023-06-05 01:35:00",
			check: "-43h40m12s",
		},
		{
			index: "index-2",
			date1: "2022-07-05 21:15:12",
			date2: "2022-07-12 12:17:32",
			check: "159h2m20s",
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				DurationBetween()

			eq(check.String(), td.check, "failed Diff DurationBetween")
		})
	}
}

func Test_Diff_DurationBetweens(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check string
	}{
		{
			index: "index-1",
			date1: "2023-06-06 21:15:12",
			date2: "2023-06-05 01:35:00",
			check: "-1 days -19 hours -40 minutes -12 seconds",
		},
		{
			index: "index-2",
			date1: "2022-07-05 21:15:12",
			date2: "2022-07-12 12:17:32",
			check: "6 days 15 hours 02 minutes 20 seconds",
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			days, hours, minutes, seconds := Parse(td.date1).
				Diff(Parse(td.date2)).
				DurationBetweens()

			eq(
				fmt.Sprintf(
					"%d days %02d hours %02d minutes %02d seconds",
					days, hours, minutes, seconds,
				),
				td.check,
				"failed Diff DurationBetweens",
			)
		})
	}
}

func Test_Diff_DurationBetweenAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check string
	}{
		{
			index: "index-1",
			date1: "2023-06-06 21:15:12",
			date2: "2023-06-05 01:35:00",
			check: "43h40m12s",
		},
		{
			index: "index-2",
			date1: "2022-07-05 21:15:12",
			date2: "2022-07-12 12:17:32",
			check: "159h2m20s",
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			check := Parse(td.date1).
				Diff(Parse(td.date2)).
				DurationBetweenAbs()

			eq(check.String(), td.check, "failed Diff DurationBetweenAbs")
		})
	}
}

func Test_Diff_DurationBetweensAbs(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index string
		date1 string
		date2 string
		check string
	}{
		{
			index: "index-1",
			date1: "2023-06-06 21:15:12",
			date2: "2023-06-05 01:35:00",
			check: "1 days 19 hours 40 minutes 12 seconds",
		},
		{
			index: "index-2",
			date1: "2022-07-05 21:15:12",
			date2: "2022-07-12 12:17:32",
			check: "6 days 15 hours 02 minutes 20 seconds",
		},
	}

	for _, td := range tests {
		t.Run(td.index, func(t *testing.T) {
			days, hours, minutes, seconds := Parse(td.date1).
				Diff(Parse(td.date2)).
				DurationBetweensAbs()

			eq(
				fmt.Sprintf(
					"%d days %02d hours %02d minutes %02d seconds",
					days, hours, minutes, seconds,
				),
				td.check,
				"failed Diff DurationBetweensAbs",
			)
		})
	}
}
