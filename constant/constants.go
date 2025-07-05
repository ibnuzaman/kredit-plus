package constant

const (
	MsgInternalServerError = "Internal Server Error"
	MsgBadRequest          = "Bad Request"
	MsgUnauthorized        = "Unauthorized"
	MsgUnprocessableEntity = "Unprocessable Entity"
	MsgErrorValidation     = "Validation Error"
)

const (
	KeyLocalsAuthUser = "auth_user"
	DefaultPerPage    = 10
)

const (
	//? 2% interest per month (assumptions is flat)
	InterestPercentage = 0.02

	//? assumptions for admin fee 5% of the otr amount
	AdminFeePercentage = 0.05

	//? assumptions for insurance fee 1% of the otr amount
	InstallmentFeePercentage = 0.01
)
