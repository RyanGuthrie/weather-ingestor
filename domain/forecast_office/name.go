package forecast_office

type Name string

func (n Name) Office() (*Instance, error) {
	return New(n)
}

const (
	ABQ Name = "ABQ"
	ABR Name = "ABR"
	AER Name = "AER"
	AFC Name = "AFC"
	AFG Name = "AFG"
	AJK Name = "AJK"
	AKQ Name = "AKQ"
	ALU Name = "ALU"
	ALY Name = "ALY"
	AMA Name = "AMA"
	APX Name = "APX"
	ARX Name = "ARX"
	BGM Name = "BGM"
	BIS Name = "BIS"
	BMX Name = "BMX"
	BOI Name = "BOI"
	BOU Name = "BOU"
	BOX Name = "BOX"
	BRO Name = "BRO"
	BTV Name = "BTV"
	BUF Name = "BUF"
	BYZ Name = "BYZ"
	CAE Name = "CAE"
	CAR Name = "CAR"
	CHS Name = "CHS"
	CLE Name = "CLE"
	CRP Name = "CRP"
	CTP Name = "CTP"
	CYS Name = "CYS"
	DDC Name = "DDC"
	DLH Name = "DLH"
	DMX Name = "DMX"
	DTX Name = "DTX"
	DVN Name = "DVN"
	EAX Name = "EAX"
	EKA Name = "EKA"
	EPZ Name = "EPZ"
	EWX Name = "EWX"
	FFC Name = "FFC"
	FGF Name = "FGF"
	FGZ Name = "FGZ"
	FSD Name = "FSD"
	FWD Name = "FWD"
	GGW Name = "GGW"
	GID Name = "GID"
	GJT Name = "GJT"
	GLD Name = "GLD"
	GRB Name = "GRB"
	GRR Name = "GRR"
	GSP Name = "GSP"
	GUM Name = "GUM"
	GYX Name = "GYX"
	HFO Name = "HFO"
	HGX Name = "HGX"
	HNX Name = "HNX"
	HPA Name = "HPA"
	HUN Name = "HUN"
	ICT Name = "ICT"
	ILM Name = "ILM"
	ILN Name = "ILN"
	ILX Name = "ILX"
	IND Name = "IND"
	IWX Name = "IWX"
	JAN Name = "JAN"
	JAX Name = "JAX"
	JKL Name = "JKL"
	KEY Name = "KEY"
	LBF Name = "LBF"
	LCH Name = "LCH"
	LIX Name = "LIX"
	LKN Name = "LKN"
	LMK Name = "LMK"
	LOT Name = "LOT"
	LOX Name = "LOX"
	LSX Name = "LSX"
	LUB Name = "LUB"
	LWX Name = "LWX"
	LZK Name = "LZK"
	MAF Name = "MAF"
	MEG Name = "MEG"
	MFL Name = "MFL"
	MFR Name = "MFR"
	MHX Name = "MHX"
	MKX Name = "MKX"
	MLB Name = "MLB"
	MOB Name = "MOB"
	MPX Name = "MPX"
	MQT Name = "MQT"
	MRX Name = "MRX"
	MSO Name = "MSO"
	MTR Name = "MTR"
	NH1 Name = "NH1"
	NH2 Name = "NH2"
	OAX Name = "OAX"
	OHX Name = "OHX"
	OKX Name = "OKX"
	ONA Name = "ONA"
	ONP Name = "ONP"
	OTX Name = "OTX"
	OUN Name = "OUN"
	PAH Name = "PAH"
	PBZ Name = "PBZ"
	PDT Name = "PDT"
	PHI Name = "PHI"
	PIH Name = "PIH"
	PPG Name = "PPG"
	PQR Name = "PQR"
	PSR Name = "PSR"
	PUB Name = "PUB"
	RAH Name = "RAH"
	REV Name = "REV"
	RIW Name = "RIW"
	RLX Name = "RLX"
	RNK Name = "RNK"
	SEW Name = "SEW"
	SGF Name = "SGF"
	SGX Name = "SGX"
	SHV Name = "SHV"
	SJT Name = "SJT"
	SJU Name = "SJU"
	SLC Name = "SLC"
	STO Name = "STO"
	STU Name = "STU"
	TAE Name = "TAE"
	TBW Name = "TBW"
	TFX Name = "TFX"
	TOP Name = "TOP"
	TSA Name = "TSA"
	TWC Name = "TWC"
	UNR Name = "UNR"
	VEF Name = "VEF"
)