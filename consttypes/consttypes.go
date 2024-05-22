package consttypes

import "os"

type TShiftInfoMerc struct {
	IsOpen      bool `json:"isOpen"`
	Is24Expired bool `json:"is24Expired"`
	Num         int  `json:"num"`
}

type TCheckInfoMerc struct {
	IsOpen bool `json:"isOpen"`
	Num    int  `json:"num"`
}

type TFnInfoMerc struct {
	Status int    `json:"status"`
	FnNum  string `json:"fnNum"`
}

type TCorrectedDataMerc struct {
	McType         int    `json:"mcType"`
	McGoodsID      string `json:"mcGoodsID"`
	ProcessingMode int    `json:"processingMode"`
}

type TOnlineCheckMerc struct {
	Result                   int                 `json:"result"`
	Description              string              `json:"description"`
	ProcessingResult         int                 `json:"processingResult"`
	McCheckResult            bool                `json:"mcCheckResult"`
	PlannedStatusCheckResult int                 `json:"plannedStatusCheckResult"`
	McCheckResultRaw         int                 `json:"mcCheckResultRaw"`
	CorrectedData            *TCorrectedDataMerc `json:"correctedData,omitempty"`
}

type TMercPayments struct {
	Cash          int `json:"cash"`
	Ecash         int `json:"ecash"`
	Prepayment    int `json:"prepayment"`
	Credit        int `json:"credit"`
	Consideration int `json:"consideration"`
}
type TCashierInfoMerc struct {
	CashierName string `json:"cashierName"`
	CashierINN  string `json:"cashierID,omitempty"`
}

type TBuyerInfoMerc struct {
	BuyerName string `json:"buyerName,omitempty"`
	BuyerINN  string `json:"buyerID,omitempty"`
}

type TCorrectionInfoMerc struct {
	CorrectionType int    `json:"correctionType"`
	CauseName      string `json:"causeName,omitempty"`
	CauseDocDate   string `json:"causeDocDate,omitempty"`
	CauseDocNum    string `json:"causeDocNum"`
}

type TUserAttribute struct {
	AttrName  string `json:"attrName"`
	AttrValue string `json:"attrValue"`
}

type TMercOpenCheck struct {
	SessionKey      string              `json:"sessionKey"`
	Command         string              `json:"command"`
	CheckType       int                 `json:"checkType"`
	TaxSystem       int                 `json:"taxSystem"`
	PrintDoc        bool                `json:"printDoc"`
	AdditionalProps string              `json:"additionalProps"`
	CashierInfo     TCashierInfoMerc    `json:"cashierInfo"`
	BuyerInfo       *TBuyerInfoMerc     `json:"buyerInfo,omitempty"`
	CorrectionInfo  TCorrectionInfoMerc `json:"correctionInfo"`
	UserAttribute   TUserAttribute      `json:"userAttribute"`
}

type TPartMerc struct {
	Numerator   int `json:"Numerator"`
	Denominator int `json:"denominator"`
}

type TMcInfoMerc struct {
	Mc             string     `json:"mc"`
	Ean            string     `json:"ean,omitempty"`
	ProcessingMode int        `json:"processingMode"`
	PlannedStatus  int        `json:"plannedStatus"`
	Part           *TPartMerc `json:"part,omitempty"`
}

type TAgentMerc struct {
	Code          int      `json:"code"`
	PayingOp      string   `json:"payingOp,omitempty"`
	PayingPhone   []string `json:"payingPhone,omitempty"`
	TransfName    string   `json:"transfName,omitempty"`
	TransfINN     string   `json:"transfINN,omitempty"`
	TransfAddress string   `json:"transfAddress,omitempty"`
	TransfPhone   string   `json:"transfPhone,omitempty"`
	OperatorPhone []string `json:"operatorPhone,omitempty"`
	SupplierPhone []string `json:"supplierPhone,omitempty"`
	SupplierINN   string   `json:"supplierINN,omitempty"`
	SupplierName  string   `json:"supplierName,omitempty"`
}

type TMercPosition struct {
	SessionKey      string       `json:"sessionKey"`
	Command         string       `json:"command"`
	MarkingCode     string       `json:"markingCode,omitempty"`
	McInfo          *TMcInfoMerc `json:"mcInfo,omitempty"`
	ProductName     string       `json:"productName"`
	Qty             int          `json:"qty"`
	MeasureUnit     int          `json:"measureUnit"`
	TaxCode         int          `json:"taxCode"`
	PaymentFormCode int          `json:"paymentFormCode"`
	ProductTypeCode int          `json:"productTypeCode"`
	Price           int          `json:"price"`
	Sum             int          `json:"sum,omitempty"`
	Agent           *TAgentMerc  `json:"agent,omitempty"`
}

type TMercCloseCheck struct {
	SessionKey  string        `json:"sessionKey"`
	Command     string        `json:"command"`
	SendCheckTo string        `json:"sendCheckTo,omitempty"`
	Payment     TMercPayments `json:"payment"`
}

type TMercKKTInfo struct {
	RegNum string `json:"regNum"`
}

type TMercRegistrationInfo struct {
	Kkt       TMercKKTInfo `json:"kkt"`
	TaxSystem []int        `json:"taxSystem"`
}

type TAnswerMercur struct {
	Result           int                    `json:"result"`
	Description      string                 `json:"description"`
	SessionKey       string                 `json:"sessionKey,omitempty"`
	ProtocolVer      string                 `json:"protocolVer,omitempty"`
	FnNum            string                 `json:"fnNum,omitempty"`
	KktNum           string                 `json:"kktNum,omitempty"`
	Model            string                 `json:"model,omitempty"`
	ShiftInfo        *TShiftInfoMerc        `json:"shiftInfo,omitempty"`
	CheckInfo        *TCheckInfoMerc        `json:"checkInfo,omitempty"`
	FnInfo           *TFnInfoMerc           `json:"fnInfo,omitempty"`
	IsCompleted      bool                   `json:"isCompleted,omitempty"`
	McCheckResultRaw int                    `json:"mcCheckResultRaw,omitempty"`
	OnlineCheck      *TOnlineCheckMerc      `json:"onlineCheck,omitempty"`
	GoodsNum         int                    `json:"goodsNum,omitempty"`
	ShiftNum         int                    `json:"shiftNum,omitempty"`
	CheckNum         int                    `json:"checkNum,omitempty"`
	FiscalDocNum     int                    `json:"fiscalDocNum,omitempty"`
	FiscalSign       string                 `json:"fiscalSign,omitempty"`
	DriverVer        string                 `json:"driverVer,omitempty"`
	DriverBaseVer    string                 `json:"driverBaseVer,omitempty"`
	RegistrationInfo *TMercRegistrationInfo `json:"registrationInfo,omitempty"`
}

type TClientInfo struct {
	EmailOrPhone string `json:"emailOrPhone"`
	Vatin        string `json:"vatin,omitempty"`
	Name         string `json:"name,omitempty"`
}

type TTaxNDS struct {
	Type string `json:"type,omitempty"`
}
type TProductCodesAtol struct {
	Undefined    string `json:"undefined,omitempty"` //32 символа только
	Code_EAN_8   string `json:"ean8,omitempty"`
	Code_EAN_13  string `json:"ean13,omitempty"`
	Code_ITF_14  string `json:"itf14,omitempty"`
	Code_GS_1    string `json:"gs10,omitempty"`
	Tag1305      string `json:"gs1m,omitempty"`
	Code_KMK     string `json:"short,omitempty"`
	Code_MI      string `json:"furs,omitempty"`
	Code_EGAIS_2 string `json:"egais20,omitempty"`
	Code_EGAIS_3 string `json:"egais30,omitempty"`
	Code_F_1     string `json:"f1,omitempty"`
	Code_F_2     string `json:"f2,omitempty"`
	Code_F_3     string `json:"f3,omitempty"`
	Code_F_4     string `json:"f4,omitempty"`
	Code_F_5     string `json:"f5,omitempty"`
	Code_F_6     string `json:"f6,omitempty"`
}

type TPayment struct {
	Type string  `json:"type"`
	Sum  float64 `json:"sum"`
}

type TGenearaPosAndTag11921191 struct {
	Type string `json:"type"`
}

type TAgentInfo struct {
	Agents []string `json:"agents"`
}
type TSupplierInfo struct {
	Vatin  string   `json:"vatin"`
	Name   string   `json:"name,omitempty"`
	Phones []string `json:"phones,omitempty"`
}

type TPosition struct {
	Type            string   `json:"type"`
	Name            string   `json:"name"`
	Price           float64  `json:"price"`
	Quantity        float64  `json:"quantity"`
	Amount          float64  `json:"amount"`
	MeasurementUnit string   `json:"measurementUnit"`
	PaymentMethod   string   `json:"paymentMethod"`
	PaymentObject   string   `json:"paymentObject"`
	Tax             *TTaxNDS `json:"tax,omitempty"`
	//fot type tag1192 //AdditionalAttribute
	Value        string             `json:"value,omitempty"`
	Print        bool               `json:"print,omitempty"`
	ProductCodes *TProductCodesAtol `json:"productCodes,omitempty"`
	ImcParams    *TImcParams        `json:"imcParams,omitempty"`
	//Mark         string             `json:"mark,omitempty"`
	AgentInfo    *TAgentInfo    `json:"agentInfo,omitempty"`
	SupplierInfo *TSupplierInfo `json:"supplierInfo,omitempty"`
}

type TAnsweChekcMark struct {
	Ready bool `json:"ready"`
}

type TBeginTaskMarkCheck struct {
	Type   string     `json:"type"`
	Params TImcParams `json:"params"`
}

type TItemInfoCheckResult struct {
	ImcCheckFlag              bool `json:"imcCheckFlag"`
	ImcCheckResult            bool `json:"imcCheckResult"`
	ImcStatusInfo             bool `json:"imcStatusInfo"`
	ImcEstimatedStatusCorrect bool `json:"imcEstimatedStatusCorrect"`
	EcrStandAloneFlag         bool `json:"ecrStandAloneFlag"`
}

type TImcParams struct {
	ImcType             string                `json:"imcType"`
	Imc                 string                `json:"imc"`
	ItemEstimatedStatus string                `json:"itemEstimatedStatus"`
	ImcModeProcessing   int                   `json:"imcModeProcessing"`
	ImcBarcode          string                `json:"imcBarcode,omitempty"`
	ItemInfoCheckResult *TItemInfoCheckResult `json:"itemInfoCheckResult,omitempty"`
	ItemQuantity        float64               `json:"itemQuantity,omitempty"`
	ItemUnits           string                `json:"itemUnits,omitempty"`
	NotSendToServer     bool                  `json:"notSendToServer,omitempty"`
}

type TAnswerGetStatusOfShift struct {
	ShiftStatus TShiftStatus `json:"shiftStatus"`
}
type TShiftStatus struct {
	DocumentsCount int    `json:"documentsCount"`
	ExpiredTime    string `json:"expiredTime"`
	Number         int    `json:"number"`
	State          string `json:"state"`
}

type TTag1192_91 struct {
	Type  string `json:"type"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Print bool   `json:"print,omitempty"`
}

type TOperator struct {
	Name  string `json:"name"`
	Vatin string `json:"vatin,omitempty"`
}

// При работе по ФФД ≥ 1.1 чеки коррекции имеют вид, аналогичный обычным чекам, но с
// добавлением информации о коррекции: тип, описание, дата документа основания и
// номер документа основания.
type TCorrectionCheck struct {
	Type string `json:"type"` //sellCorrection - чек коррекции прихода
	//buyCorrection - чек коррекции расхода
	//sellReturnCorrection - чек коррекции возврата прихода (ФФД ≥ 1.1)
	//buyReturnCorrection - чек коррекции возврата расхода
	Electronically       bool         `json:"electronically"`
	TaxationType         string       `json:"taxationType,omitempty"`
	ClientInfo           *TClientInfo `json:"clientInfo"`
	CorrectionType       string       `json:"correctionType"` //
	CorrectionBaseDate   string       `json:"correctionBaseDate"`
	CorrectionBaseNumber string       `json:"correctionBaseNumber"`
	Operator             TOperator    `json:"operator"`
	//Items                []TPosition `json:"items"`
	Items    []interface{} `json:"items"` //либо TTag1192_91, либо TPosition
	Payments []TPayment    `json:"payments"`
	Total    float64       `json:"total,omitempty"`
}

type TItemInfoCheckResultObject struct {
	ItemInfoCheckResult TItemInfoCheckResult `json:"itemInfoCheckResult"`
}

var DIROFJSONS = ".\\jsons\\works\\"
var LOGSDIR = "./logs/"

const LOGINFO = "info"
const LOGINFO_WITHSTD = "info_std"
const LOGERROR = "error"
const LOGSKIP_LINES = "skip_line"
const LOGOTHER = "other"
const LOG_PREFIX = "TASKS"

const FILE_NAME_PRINTED_CHECKS = "printed.txt"
const FILE_NAME_CONNECTION = "connection.txt"

func DoesFileExist(fullFileName string) (found bool, err error) {
	found = false
	if _, err = os.Stat(fullFileName); err == nil {
		// path/to/whatever exists
		found = true
	}
	return
}
