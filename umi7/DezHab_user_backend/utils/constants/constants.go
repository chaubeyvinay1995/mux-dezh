package constants

const (
	RequestData           = "request-data"
	RequestID             = "request-id"
	GoPath                = "GOPATH"
	Provider              = "provider"
	UploadedBy            = "uploadedBy"
	Action                = "action"
	SubTag                = "subTag"
	ClientID              = "clientId"
	Result                = "result"
	Environment           = "ENV"
	DefaultEnvironment    = "dev"
	ProductionEnvironment = "prod"
	FileExtension         = "jpeg"
	PathSeparator         = "/"
	AWSFolderName         = "assets/"
	BucketPrefix          = "users/"
	HttpPrefix            = "https://"
	S3Extension           = ".s3."
	S3Suffix              = ".amazonaws.com"
	S3Prefix              = "s3://"
	BucketName
)

// error messages
const (
	InvalidOrProcessedRequestID = "invalid request id"
	MaxStoreLimitExceededErr    = "max limit exceeded - Max %d stores are allowed at once"
	MaxStoreRecordExceededErr   = "max limit exceeded - Max %d stores and %d records are allowed at once"
	UnauthorizedAPIKey          = "unauthorized API Key"
	InvalidAction               = "invalid action provided in request"
	FilterInvalidFormat         = "applied filters are not in valid format"
	ActionRequiredErr           = "action filed is required"
)

// Success Message

const (
	OrganizationCreated     = "Organization created successfully."
	OrganizationUpdated     = "Organization updated successfully."
	ProjectCreated          = "Project created successfully."
	ProjectUpdated          = "Project updated successfully."
	ProjectDeleted          = "Project deleted successfully."
	AcademicCreated         = "Academic detail created successfully."
	AcademicDeleted         = "Academic detail deleted successfully."
	AcademicUpdated         = "Academic detail updated successfully."
	CertificateCreated      = "Certificate detail created successfully."
	CertificateUpdated      = "Certificate detail updated successfully."
	CertificateDeleted      = "Certificate detail deleted successfully."
	FileUploadResponse      = "File Uploaded successfully."
	FolderCreatedResponse   = "Folder created successfully."
	FolderUpdatedResponse   = "Folder updated successfully."
	AWSAssetsDeleteResponse = "Assets deleted successfully."
	NotAuthorized           = "You are not authorized for operation."
	USERIDAlreadyPresent    = "User id already registered."
	PersonalInfoCreated     = "Personal info created successfully."
	PersonalInfoUpdated     = "Personal info updated successfully."
	PersonalInfoDeleted     = "Personal info deleted successfully."
	ProfessionalInfoDeleted = "Professional info deleted successfully."
	ProfessionalInfoCreated = "Professional info created successfully."
	ProfessionalInfoUpdated = "Professional info updated successfully."
	FileTypeNotSupported    = "File type is not supported."
)

var SupportedFileExtensions = [...]string{".pdf", ".docx", ".doc", ".xlsx", "xls", ".pptx", ".ppt", "jpg", ".jpeg",
	".gif", ".tiff", ".svg", ".psd", ".ai", ".eps", "https:\\\\", "http:\\\\", "www.", ".dwg", ".iges",
	".stl", ".rvt", ".skp", ".3ds", ".mp4", ".mkv", ".avi"}
