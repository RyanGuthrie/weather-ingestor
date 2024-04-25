package forecast_office

import (
	"fmt"
	"weather-ingestor/util/container"
)

var nameToInstance = map[string]Name{
	"ABQ": ABQ, "ABR": ABR, "AER": AER, "AFC": AFC, "AFG": AFG, "AJK": AJK, "AKQ": AKQ,
	"ALU": ALU, "ALY": ALY, "AMA": AMA, "APX": APX, "ARX": ARX, "BGM": BGM, "BIS": BIS,
	"BMX": BMX, "BOI": BOI, "BOU": BOU, "BOX": BOX, "BRO": BRO, "BTV": BTV, "BUF": BUF,
	"BYZ": BYZ, "CAE": CAE, "CAR": CAR, "CHS": CHS, "CLE": CLE, "CRP": CRP, "CTP": CTP,
	"CYS": CYS, "DDC": DDC, "DLH": DLH, "DMX": DMX, "DTX": DTX, "DVN": DVN, "EAX": EAX,
	"EKA": EKA, "EPZ": EPZ, "EWX": EWX, "FFC": FFC, "FGF": FGF, "FGZ": FGZ, "FSD": FSD,
	"FWD": FWD, "GGW": GGW, "GID": GID, "GJT": GJT, "GLD": GLD, "GRB": GRB, "GRR": GRR,
	"GSP": GSP, "GUM": GUM, "GYX": GYX, "HFO": HFO, "HGX": HGX, "HNX": HNX, "HPA": HPA,
	"HUN": HUN, "ICT": ICT, "ILM": ILM, "ILN": ILN, "ILX": ILX, "IND": IND, "IWX": IWX,
	"JAN": JAN, "JAX": JAX, "JKL": JKL, "KEY": KEY, "LBF": LBF, "LCH": LCH, "LIX": LIX,
	"LKN": LKN, "LMK": LMK, "LOT": LOT, "LOX": LOX, "LSX": LSX, "LUB": LUB, "LWX": LWX,
	"LZK": LZK, "MAF": MAF, "MEG": MEG, "MFL": MFL, "MFR": MFR, "MHX": MHX, "MKX": MKX,
	"MLB": MLB, "MOB": MOB, "MPX": MPX, "MQT": MQT, "MRX": MRX, "MSO": MSO, "MTR": MTR,
	"NH1": NH1, "NH2": NH2, "OAX": OAX, "OHX": OHX, "OKX": OKX, "ONA": ONA, "ONP": ONP,
	"OTX": OTX, "OUN": OUN, "PAH": PAH, "PBZ": PBZ, "PDT": PDT, "PHI": PHI, "PIH": PIH,
	"PPG": PPG, "PQR": PQR, "PSR": PSR, "PUB": PUB, "RAH": RAH, "REV": REV, "RIW": RIW,
	"RLX": RLX, "RNK": RNK, "SEW": SEW, "SGF": SGF, "SGX": SGX, "SHV": SHV, "SJT": SJT,
	"SJU": SJU, "SLC": SLC, "STO": STO, "STU": STU, "TAE": TAE, "TBW": TBW, "TFX": TFX,
	"TOP": TOP, "TSA": TSA, "TWC": TWC, "UNR": UNR, "VEF": VEF,
}

var Instances Offices

type Offices struct {
	StringToName map[string]Name
	All          []Name
}

func init() {
	Instances = Offices{
		StringToName: nameToInstance,
		All:          container.Values(nameToInstance),
	}
}

func (o *Offices) InstanceFromString(name string) (*Instance, error) {
	parsedName, exists := nameToInstance[name]
	if !exists {
		return nil, fmt.Errorf("invalid forecast office identifier: %s", name)
	}

	return New(parsedName)
}

func (o *Offices) FromStringLazy(name string) (Name, error) {
	parsedName, exists := nameToInstance[name]
	if !exists {
		return "", fmt.Errorf("invalid forecast office identifier: %s", name)
	}

	return parsedName, nil
}
