package models

type BaseRequest struct {
	UID    string `json:"uid"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

type HBA1CRequest struct {
	BaseRequest
	RDW  float64 `json:"rdw"`
	WBC  float64 `json:"wbc"`
	RBC  float64 `json:"rbc"`
	HGB  float64 `json:"hgb"`
	HCT  float64 `json:"hct"`
	MCV  float64 `json:"mcv"`
	MCH  float64 `json:"mch"`
	MCHC float64 `json:"mchc"`
	PLT  float64 `json:"plt"`
	NEU  float64 `json:"neu"`
	EOS  float64 `json:"eos"`
	BAS  float64 `json:"bas"`
	LYM  float64 `json:"lym"`
	MON  float64 `json:"mon"`
	SOE  float64 `json:"soe"`
	CHOL float64 `json:"chol"`
	GLU  float64 `json:"glu"`
}

type LDLRequest struct {
	BaseRequest
	RDW  float64 `json:"rdw"`
	WBC  float64 `json:"wbc"`
	RBC  float64 `json:"rbc"`
	HGB  float64 `json:"hgb"`
	HCT  float64 `json:"hct"`
	MCV  float64 `json:"mcv"`
	MCH  float64 `json:"mch"`
	MCHC float64 `json:"mchc"`
	PLT  float64 `json:"plt"`
	NEU  float64 `json:"neu"`
	EOS  float64 `json:"eos"`
	BAS  float64 `json:"bas"`
	LYM  float64 `json:"lym"`
	MON  float64 `json:"mon"`
	SOE  float64 `json:"soe"`
	CHOL float64 `json:"chol"`
	GLU  float64 `json:"glu"`
}

type LDLLRequest struct {
	BaseRequest
	CHOL float64 `json:"chol"`
	HDL  float64 `json:"hdl"`
	TG   float64 `json:"tg"`
}

type FERRRequest struct {
	BaseRequest
	RDW  float64 `json:"rdw"`
	WBC  float64 `json:"wbc"`
	RBC  float64 `json:"rbc"`
	HGB  float64 `json:"hgb"`
	HCT  float64 `json:"hct"`
	MCV  float64 `json:"mcv"`
	MCH  float64 `json:"mch"`
	MCHC float64 `json:"mchc"`
	PLT  float64 `json:"plt"`
	NEU  float64 `json:"neu"`
	EOS  float64 `json:"eos"`
	BAS  float64 `json:"bas"`
	LYM  float64 `json:"lym"`
	MON  float64 `json:"mon"`
	SOE  float64 `json:"soe"`
	CRP  float64 `json:"crp"`
}

type TGRequest struct {
	BaseRequest
	RDW  float64 `json:"rdw"`
	WBC  float64 `json:"wbc"`
	RBC  float64 `json:"rbc"`
	HGB  float64 `json:"hgb"`
	HCT  float64 `json:"hct"`
	MCV  float64 `json:"mcv"`
	MCH  float64 `json:"mch"`
	MCHC float64 `json:"mchc"`
	PLT  float64 `json:"plt"`
	NEU  float64 `json:"neu"`
	EOS  float64 `json:"eos"`
	BAS  float64 `json:"bas"`
	LYM  float64 `json:"lym"`
	MON  float64 `json:"mon"`
	SOE  float64 `json:"soe"`
	CHOL float64 `json:"chol"`
	GLU  float64 `json:"glu"`
}

type HDLRequest struct {
	BaseRequest
	RDW  float64 `json:"rdw"`
	WBC  float64 `json:"wbc"`
	RBC  float64 `json:"rbc"`
	HGB  float64 `json:"hgb"`
	HCT  float64 `json:"hct"`
	MCV  float64 `json:"mcv"`
	MCH  float64 `json:"mch"`
	MCHC float64 `json:"mchc"`
	PLT  float64 `json:"plt"`
	NEU  float64 `json:"neu"`
	EOS  float64 `json:"eos"`
	BAS  float64 `json:"bas"`
	LYM  float64 `json:"lym"`
	MON  float64 `json:"mon"`
	SOE  float64 `json:"soe"`
	CHOL float64 `json:"chol"`
	GLU  float64 `json:"glu"`
}
