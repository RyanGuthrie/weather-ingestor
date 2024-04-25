package uom

// uom = Unit of Measurement (weather.gov terminology)

type UOM string

const (
	DegC        UOM = "wmoUnit:degC"
	DegreeAngle UOM = "wmoUnit:degree_(angle)"
	KMH         UOM = "wmoUnit:km_h-1"
	M           UOM = "wmoUnit:m"
	MM          UOM = "wmoUnit:mm"
	PERCENT     UOM = "wmoUnit:percent"
)
