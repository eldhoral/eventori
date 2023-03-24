package applicationConstants

const (
	MigrateUp = iota
	MigrateDown
)

const (
	CHANNEL_TYPE        = "channeling"
	WORDING_PASSED_HTML = "<b class=\"text-success\">Passed</b>"
	WORDING_FAILED_HTML = "<b class=\"text-danger\">Failed</b>"
)

// pefindo
const (
	PEFINDO_SUCCESS_STATUS            = "success"
	PEFINDO_IN_PROGRESS_STATUS        = "inprogress"
	PEFINDO_STATUS_DONE               = "done"
	PEFINDO_BIGGEST_GRADE             = "A1"
	PEFINDO_NOT_SPECIFIED             = "NotSpecified"
	PEFINDO_NO_NEGATIVE_STATUS        = "NoNegativeStatus"
	PEFINDO_MONITORED                 = "Monitored"
	PEFINDO_SUBSTANDARD               = "Substandard"
	PEFINDO_DOUBTFUL                  = "Doubtful"
	PEFINDO_DELINQUENT                = "Delinquent"
	PEFINDO_DEFAULT_LIMIT             = 10
	IS_PASS_PEFINDO_ERROR             = 4
	PEFINDO_CONTRACT_TYPE_CREDIT_CARD = "CreditCard"
)

// rekonsiliasi
const (
	REKONSILIASI_PAID = iota
	REKONSILIASI_UNPAID
	REKONSILIASI_UNDERPAYMENT
	REKONSILIASI_OVERPAYMENT
	REKONSILIASI_NOT_SPECIFIED
	REKONSILIASI_FAILED
	REKONSILIASI_PEMBAYARAN_DIPERCEPAT
	REKONSILIASI_PEMBAYARAN_DIPERCEPAT_KURANG_BAYAR
	REKONSILIASI_PEMBAYARAN_DIPERCEPAT_LEBIH_BAYAR
	REKONSILIASI_PEMBAYARAN_TERLAMBAT
	REKONSILIASI_PEMBAYARAN_TERLAMBAT_KURANG_BAYAR
	REKONSILIASI_PEMBAYARAN_TERLAMBAT_LEBIH_BAYAR

	// text file generation
	TEXTFILE_UNPROCESSED
	TEXTFILE_GENERATED_SUCCESS
	TEXTFILE_GENERATED_FAILED
	TEXTFILE_ERROR
	TEXTFILE_INPROGRESS
	TEXTFILE_REJECT
	TEXTFILE_SUCCESS
)

var StatusRekonsiliasi = map[int]string{
	REKONSILIASI_PAID:                               "Sudah Bayar",
	REKONSILIASI_UNPAID:                             "Belum Bayar",
	REKONSILIASI_UNDERPAYMENT:                       "Kurang Bayar",
	REKONSILIASI_OVERPAYMENT:                        "Kelebihan Bayar",
	REKONSILIASI_NOT_SPECIFIED:                      "Dalam Proses",
	REKONSILIASI_FAILED:                             "Gagal",
	REKONSILIASI_PEMBAYARAN_DIPERCEPAT:              "Pembayaran Dipercepat",
	REKONSILIASI_PEMBAYARAN_DIPERCEPAT_KURANG_BAYAR: "Pembayaran Dipercepat, Kurang Bayar",
	REKONSILIASI_PEMBAYARAN_DIPERCEPAT_LEBIH_BAYAR:  "Pembayaran Dipercepat, Lebih Bayar",
	REKONSILIASI_PEMBAYARAN_TERLAMBAT:               "Pembayaran Terlambat",
	REKONSILIASI_PEMBAYARAN_TERLAMBAT_KURANG_BAYAR:  "Pembayaran Terlambat, Kurang Bayar",
	REKONSILIASI_PEMBAYARAN_TERLAMBAT_LEBIH_BAYAR:   "Pembayaran Terlambat, Lebih Bayar",

	// text file generation
	TEXTFILE_UNPROCESSED:       "Textfile Belum Diproses",
	TEXTFILE_GENERATED_SUCCESS: "Textfile Berhasil Dibuat",
	TEXTFILE_GENERATED_FAILED:  "Textfile Gagal Dibuat",
	TEXTFILE_ERROR:             "Textfile Error Diproses",
	TEXTFILE_INPROGRESS:        "Textfile Sedang Diproses",
	TEXTFILE_REJECT:            "Textfile Ditolak",
	TEXTFILE_SUCCESS:           "Textfile Berhasil Diproses",
}

const ()

var StatusTextFile = map[int]string{}

const (
	INTERNAL_BANK_NAME        = "NOBU"
	CONTRACT_TYPE_CREDIT_CARD = "CREDITCARD"
	CONTRACT_TYPE_PINJAMAN    = "PINJAMAN"
)

var ALL_NEGATIVE_STATUS = []string{PEFINDO_NOT_SPECIFIED, PEFINDO_NO_NEGATIVE_STATUS, PEFINDO_MONITORED, PEFINDO_SUBSTANDARD, PEFINDO_DOUBTFUL, PEFINDO_DELINQUENT}

const (
	FAILED_STATUS = iota
	SUCCESS_STATUS
	NOT_AVAILABLE_STATUS
	IN_PROGRESS_STATUS
	REJECT_STATUS
)

var StatusText = map[int]string{
	FAILED_STATUS:        "FAILED",
	SUCCESS_STATUS:       "SUCCESS",
	NOT_AVAILABLE_STATUS: "NOT AVAILABLE",
	IN_PROGRESS_STATUS:   "IN PROGRESS",
	REJECT_STATUS:        "REJECT",
}

const (
	MALE_GENDER_CODE     = 1
	MALE_GENDER_STRING   = "L"
	MALE_GENDER_STATUS   = "Laki-Laki"
	FEMALE_GENDER_CODE   = 2
	FEMALE_GENDER_STRING = "P"
	FEMALE_GENDER_STATUS = "Perempuan"
)

const (
	ResultFail = iota
	ResultPass
)

const (
	BVSystemDownResponseCode = "0"
	BVSuccessResponseCode    = "00"
)

var (
	ProspectNotSync int64 = 0
	ProspectIsSync  int64 = 1
)

var (
	ProspectIsSyncInAnalystVerification   int64 = 2
	ProspectIsSyncInComplianceCheck       int64 = 3
	ProspectIsSyncInAnalystRecomendation  int64 = 4
	ProspectIsSyncInApproverRecomendation int64 = 5
)

const (
	SLIKStatusUnProcessed = iota
	SLIKStatusOnProcess
	SLIKStatusOnWaiting
	SLIKStatusReady
	SLIKStatusChecked
)

const (
	DukcapilVerificationStatusFailed = iota
	DukcapilVerificationStatusPassed
	DukcapilVerificationStatusUnProcessed
	DukcapilVerificationStatusOnCheck
	DukcapilVerificationStatusError
)

const (
	DhnVerificationStatusFailed = iota
	DhnVerificationStatusPassed
	DhnVerificationStatusUnProcessed
	DhnVerificationStatusOnCheck
	DhnVerificationStatusError
)

const (
	PepVerificationStatusFailed = iota
	PepVerificationStatusPassed
	PepVerificationStatusUnProcessed
	PepVerificationStatusOnCheck
	PepVerificationStatusError
)

const (
	SlikVerificationStatusFailed = iota
	SlikVerificationStatusPassed
	SlikVerificationStatusUnProcessed
	SlikVerificationStatusOnCheck
	SlikVerificationStatusError
)

const (
	RacVerificationStatusFailed = iota
	RacVerificationStatusPassed
	RacVerificationStatusUnProcessed
	RacVerificationStatusOnCheck
	RacVerificationStatusError
)

const (
	EligibleVerificationStatusFailed = iota
	EligibleVerificationStatusPassed
	EligibleVerificationStatusUnProcessed
	EligibleVerificationStatusOnCheck
	EligibleVerificationStatusError
)

const (
	AssessmentTypePefindo               = "PEFINDO"
	AssessmentTypeSlik                  = "SLIK"
	AssessmentTypeSlikSpouse            = "SLIK SPOUSE"
	AssessmentTypeAsli                  = "ASLI"
	AssessmentTypeRac                   = "RAC"
	AssessmentTypeDukcapil              = "DUKCAPIL"
	AssessmentTypeDHN                   = "DHN"
	AssessmentTypeDHNSpouse             = "DHN SPOUSE"
	AssessmentTypePEP                   = "PEP"
	AssessmentTypePEPSpouse             = "PEP SPOUSE"
	InvalidFullName                     = "Nama tidak sesuai"
	InvalidNIK                          = "Nomor KTP tidak sesuai"
	InvalidMotherName                   = "Nama ibu kandung tidak sesuai"
	InvalidPOB                          = "Tempat lahir tidak sesuai"
	InvalidDOB                          = "Tanggal lahir tidak sesuai"
	InvalidAddress                      = "Alamat tidak sesuai"
	ValidFullName                       = "Nama sesuai"
	ValidNIK                            = "Nomor KTP sesuai"
	ValidMotherName                     = "Nama ibu kandung sesuai"
	ValidPOB                            = "Tempat lahir sesuai"
	ValidDOB                            = "Tanggal lahir sesuai"
	ValidAddress                        = "Alamat sesuai"
	ReasonKtpTidakSesuai                = "Data tidak sesuai dengan foto KTP"
	ReasonComplianceFailed              = "Hasil compliance check DHN / PEP gagal"
	ReasonSlikFailed                    = "Hasil pengecekan SLIK gagal"
	ReasonPefindoFailed                 = "Hasil pengecekan pefindo gagal"
	ReasonRacFailed                     = "Hasil pengecekan RAC gagal"
	ReasonDukcapilFailed                = "Hasil pengecekan Dukcapil gagal"
	SLIKRejectReason                    = "Reject Kolektibilitas - SLIK Failed"
	RACRejectReason                     = "Reject RAC"
	DukcapilRejectReason                = "Reject Dukcapil"
	PefindoRejectReason                 = "Reject Pefindo"
	MaxProspectAge                      = 65
	MinProspectAge                      = 21
	DATIIIReasonNotFound                = "Kode DATI II tidak ditemukan/Format NIK salah"
	DATIIIReasonPenanggungJawabNotFound = "Kode DATI II Penanggung jawab tidak ditemukan/Format NIK salah"
	SPOUSE_NIK_INVALID                  = "NIK Penanggung Jawab tidak valid"
	ProfessionCodeNotFound              = "Kode profesi tidak ditemukan atau tidak valid"
	DukcapilMatch                       = "Sesuai"
	DukcapilNotMatch                    = "Tidak Sesuai"
	DukcapilPassScore                   = 90
	SystemDown                          = "System Down"
	DDMMYYYY                            = "02012006"
	YYYYMMDD                            = "20060102"
)

const (
	URLConfigAuthURL                            = 1
	URLConfigApprovalServiceInitTrx             = 2
	URLConfigApprovalServiceGetApprovalTrx      = 3
	URLConfigApprovalServiceApproveTrx          = 4
	URLConfigApprovalServiceRejectTrx           = 5
	URLConfigApprovalServiceGetDeviationHistory = 7
	URLConfigApprovalServiceGetTiers            = 8
	URLConfigPartnerByName                      = 14
	URLConfigDatiByKotaName                     = 15
)

const (
	BVTypeCustomerInquiry = "CUSTOMER INQUIRY"
	BVTypeAccountCIFCheck = "ACCOUNT CIF CHECK"
)

const (
	EmptyCifNumber = "000000000000000"
)

var (
	// TODO : if use SLIK use ALL_ASSESSMENT_TYPE_WITH_SLIK
	ALL_ASSESSMENT_TYPE           = []string{AssessmentTypeDukcapil, AssessmentTypeDHN, AssessmentTypeDHNSpouse, AssessmentTypePEP, AssessmentTypePEPSpouse, AssessmentTypePefindo, AssessmentTypeRac}
	ALL_ASSESSMENT_TYPE_WITH_SLIK = []string{AssessmentTypeDukcapil, AssessmentTypeDHN, AssessmentTypeDHNSpouse, AssessmentTypePEP, AssessmentTypePEPSpouse, AssessmentTypeSlik, AssessmentTypeRac}
	JenisPenggunaanForCheckSpouse = []string{"Modal Kerja", "Investasi"}
)

const (
	DukcapilDemografiEndpoint   = "/v1/validate/demografi"
	DukcapilSuccessCodeResponse = "00"
	DukcapilUnauthorized        = "1002"
)

const (
	//path /image to store images and /textfile to store textfile in GCS
	FilePath          = "channeling/"
	ImagePath         = "image/"
	TextfilePath      = "/repayment/textfile/"
	TextfilePathLocal = "./repayment/textfile/"
	TextFileExtension = ".txt"
	SuccessStatusText = "Success"
)

// Cache naming
const (
	CacheKeySektorEkonomi                 = "dlo_channeling:sektor_ekonomi:code:"
	CacheKeyTextFileGeneratedNumberPerDay = "dlo_channeling:text_file:generated_number"
)

const (
	StatusCodeDown = 0
)
