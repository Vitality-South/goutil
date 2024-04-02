package file

import (
	"testing"
)

func TestConstants(t *testing.T) {
	t.Parallel()

	c := map[int]int{
		OSUserRead:                         0o400,
		OSUserWrite:                        0o200,
		OSUserExecute:                      0o100,
		OSUserReadWrite:                    0o600,
		OSUserReadWriteExecute:             0o700,
		OSUserReadExecute:                  0o500,
		OSUserWriteExecute:                 0o300,
		OSGroupRead:                        0o040,
		OSGroupWrite:                       0o020,
		OSGroupExecute:                     0o010,
		OSGroupReadWrite:                   0o060,
		OSGroupReadWriteExecute:            0o070,
		OSGroupReadExecute:                 0o050,
		OSGroupWriteExecute:                0o030,
		OSOtherRead:                        0o004,
		OSOtherWrite:                       0o002,
		OSOtherExecute:                     0o001,
		OSOtherReadWrite:                   0o006,
		OSOtherReadWriteExecute:            0o007,
		OSOtherReadExecute:                 0o005,
		OSOtherWriteExecute:                0o003,
		OSAllRead:                          0o444,
		OSAllWrite:                         0o222,
		OSAllExecute:                       0o111,
		OSAllReadWrite:                     0o666,
		OSAllReadWriteExecute:              0o777,
		OSUserRWGroupROtherR:               0o644,
		OSUserRWXGroupRXOtherRX:            0o755,
		OSUserReadWrite | OSGroupReadWrite: 0o660,
	}

	for k, v := range c {
		if k != v {
			t.Errorf("expected %d, got %d", v, k)
		}
	}
}
