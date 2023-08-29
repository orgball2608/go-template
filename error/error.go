package error

type Code string

var (

	// Authentication
	CodeAuthenticationTokenExpired     Code = "ERROR3_01"
	CodeAuthenticationTokenInvalid     Code = "ERROR3_02"
	CodeAuthenticationTokenSfdcExpired Code = "ERROR3_10 "

	// General
	CodeGeneralSomethingWentWrong Code = "ERROR2_001"
	CodeGeneralForbidden          Code = "ERROR2_101"
	CodeGeneralInvalidIp          Code = "ERROR2_100"
	// Logic
	CodeLogicLoginFailed                            Code = "ERROR2_200"
	CodeLogicLoginFailedManyTimes                   Code = "ERROR2_201"
	CodeLogicRecordNotFound                         Code = "ERROR2_202"
	CodeLogicRequestResetPasswordExisted            Code = "ERROR2_203"
	CodeLogicCommentsAreIncorrect                   Code = "ERROR2_205"
	CodeLogicRecordAlreadyExists                    Code = "ERROR2_207"
	CodeLogicPasswordIncorrect                      Code = "ERROR2_209"
	CodeLogicUseCurrentPasswordToChangePassword     Code = "ERROR2_210"
	CodeLogicWriteValuesAreNotConsistent            Code = "ERROR2_212"
	CodeLogicFormatRequestParam                     Code = "ERROR2_213"
	CodeLogicTokenExpired                           Code = "ERROR2_214"
	CodeLogicTokenInvalid                           Code = "ERROR2_215"
	CodeLogicUploadFileFail                         Code = "ERROR2_216"
	CodeLogicDownloadFileFailed                     Code = "ERROR2_217"
	CodeLogicRefreshTokenInvalid                    Code = "ERROR2_218"
	CodeGroupExisted                                Code = "ERROR2_220"
	CodeUserExisted                                 Code = "ERROR2_221"
	CodeLogicInputIncorrectAccordingRole            Code = "ERROR2_219"
	CodeLogicKeySearchHasCharacterFullWidthKatakana Code = "ERROR2_220"
	CodeLogicWrongTime                              Code = "ERROR2_221"
	CodeContractNotExist                            Code = "ERROR2_223"
	CodeCreateFaqBookmark                           Code = "ERROR2_222"
	CodeLogicSyncFileFaqFailed                      Code = "ERROR2_224"
)
