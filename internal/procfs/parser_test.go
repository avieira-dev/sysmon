package procfs

import "testing"

func TestParseCPUStats(t *testing.T) {
	t.Run("valid_input", func(t *testing.T) {
		// Simplified line from /proc/stat
		input := "cpu 100 20 30 450 5 2 3 1 0 0"

		result, err := parseCPUStats(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.User != 100 {
			t.Errorf("expected User to be 100, got %d", result.User)
		}

		if result.Nice != 20 {
			t.Errorf("expected Nice to be 20, got %d", result.Nice)
		}

		if result.System != 30 {
			t.Errorf("expected System to be 30, got %d", result.System)
		}

		if result.Idle != 450 {
			t.Errorf("expected Idle to be 450, got %d", result.Idle)
		}

		if result.IOWait != 5 {
			t.Errorf("expected IOWait to be 5, got %d", result.IOWait)
		}

		if result.IRQ != 2 {
			t.Errorf("expected IRQ to be 2, got %d", result.IRQ)
		}

		if result.SoftIRQ != 3 {
			t.Errorf("expected SoftIRQ to be 3, got %d", result.SoftIRQ)
		}

		if result.Steal != 1 {
			t.Errorf("expected Steal to be 1, got %d", result.Steal)
		}

		if result.Guest != 0 {
			t.Errorf("expected Guest to be 0, got %d", result.Guest)
		}

		if result.GuestNice != 0 {
			t.Errorf("expected GuestNice to be 0, got %d", result.GuestNice)
		}
	})

	t.Run("invalid_identifier", func(t *testing.T) {
		input := "processor 100 20 30 450 5 2 3 1 0 0"

		_, err := parseCPUStats(input)
		if err == nil {
			t.Fatalf("expected an error for an invalid CPU identifier, but got nil")
		}
	})

	t.Run("insufficient_fields", func(t *testing.T) {
		input := "cpu 100 20 30 450 0"

		_, err := parseCPUStats(input)
		if err == nil {
			t.Fatalf("expected an error for insufficient CPU fields, but got nil")
		}
	})

	t.Run("non_numeric_value", func(t *testing.T) {
		input := "cpu 100 20 abc 450 5 2 3 1 0 0"

		_, err := parseCPUStats(input)
		if err == nil {
			t.Fatalf("expected an error for a non-numeric CPU field, but got nil")
		}
	})

	t.Run("value_overflow", func(t *testing.T) {
		input := "cpu 18446744073709551616 20 30 450 5 2 3 1 0 0"

		_, err := parseCPUStats(input)
		if err == nil {
			t.Fatalf("expected an error for a CPU field exceeding the uint64 limit, but got nil")
		}
	})
}
