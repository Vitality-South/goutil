// Package file implements constants and helper functions for working with
// files and file permissions
package file

const (
	osRead       = 04
	osWrite      = 02
	osExecute    = 01
	osUserShift  = 6
	osGroupShift = 3
	osOtherShift = 0

	// OSUserRead is unix permission 0400.
	OSUserRead = osRead << osUserShift

	// OSUserWrite is unix permission 0200.
	OSUserWrite = osWrite << osUserShift

	// OSUserExecute is unix permission 0100.
	OSUserExecute = osExecute << osUserShift

	// OSUserReadWrite is unix permission 0600.
	OSUserReadWrite = OSUserRead | OSUserWrite

	// OSUserReadWriteExecute is unix permission 0700.
	OSUserReadWriteExecute = OSUserRead | OSUserWrite | OSUserExecute

	// OSUserReadExecute is unix permission 0500.
	OSUserReadExecute = OSUserRead | OSUserExecute

	// OSUserWriteExecute is unix permission 0300.
	OSUserWriteExecute = OSUserWrite | OSUserExecute

	// OSGroupRead is unix permission 0040.
	OSGroupRead = osRead << osGroupShift

	// OSGroupWrite is unix permission 0020.
	OSGroupWrite = osWrite << osGroupShift

	// OSGroupExecute is unix permission 0010.
	OSGroupExecute = osExecute << osGroupShift

	// OSGroupReadWrite is unix permission 0060.
	OSGroupReadWrite = OSGroupRead | OSGroupWrite

	// OSGroupReadWriteExecute is unix permission 0070.
	OSGroupReadWriteExecute = OSGroupRead | OSGroupWrite | OSGroupExecute

	// OSGroupReadExecute is unix permission 0500.
	OSGroupReadExecute = OSGroupRead | OSGroupExecute

	// OSGroupWriteExecute is unix permission 0300.
	OSGroupWriteExecute = OSGroupWrite | OSGroupExecute

	// OSOtherRead is unix permission 0004.
	OSOtherRead = osRead << osOtherShift

	// OSOtherWrite is unix permission 0002.
	OSOtherWrite = osWrite << osOtherShift

	// OSOtherExecute is unix permission 0001.
	OSOtherExecute = osExecute << osOtherShift

	// OSOtherReadWrite is unix permission 0006.
	OSOtherReadWrite = OSOtherRead | OSOtherWrite

	// OSOtherReadWriteExecute is unix permission 0007.
	OSOtherReadWriteExecute = OSOtherRead | OSOtherWrite | OSOtherExecute

	// OSOtherReadExecute is unix permission 0500.
	OSOtherReadExecute = OSOtherRead | OSOtherExecute

	// OSOtherWriteExecute is unix permission 0300.
	OSOtherWriteExecute = OSOtherWrite | OSOtherExecute

	// OSAllRead is unix permission 0444.
	OSAllRead = OSUserRead | OSGroupRead | OSOtherRead

	// OSAllWrite is unix permission 0222.
	OSAllWrite = OSUserWrite | OSGroupWrite | OSOtherWrite

	// OSAllExecute is unix permission 0111.
	OSAllExecute = OSUserExecute | OSGroupExecute | OSOtherExecute

	// OSAllReadWrite is unix permission 0666.
	OSAllReadWrite = OSAllRead | OSAllWrite

	// OSAllReadWriteExecute is unix permission 0777.
	OSAllReadWriteExecute = OSAllRead | OSAllWrite | OSAllExecute

	// OSUserRWGroupROtherR is unix permission 0644.
	OSUserRWGroupROtherR = OSUserReadWrite | OSGroupRead | OSOtherRead

	// OSUserRWXGroupRXOtherRX is unix permission 0755.
	OSUserRWXGroupRXOtherRX = OSUserReadWriteExecute |
		OSGroupReadExecute | OSOtherReadExecute
)
