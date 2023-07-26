package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gitlab.com/umi7/DezHab_user_backend/api"
	"gitlab.com/umi7/DezHab_user_backend/api/handler"
	"gitlab.com/umi7/DezHab_user_backend/api/rest"
	"gitlab.com/umi7/DezHab_user_backend/config"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	"gitlab.com/umi7/DezHab_user_backend/utils"
	"net/http"
	"os"
	"time"
)

const SwaggerPath = "/src/gitlab.com/umi7/DezHab_user_backend/resources/swaggerui/"

// Server represents a server mux
type Server struct {
	*mux.Router
	Address string
}

// New setups & returns a server
func New() *Server {
	router := mux.NewRouter()
	address := config.AppConfig.Server.IP + ":" + config.AppConfig.Server.Port
	s := Server{Router: router, Address: address}
	s.SetupRouter()
	return &s
}

// SetupRouter configures the route for the server
func (s Server) SetupRouter() {
	s.HandleFunc("/status", api.HealthCheck).Methods(http.MethodGet)
	// For swagger
	dirPath := utils.GoPath + SwaggerPath
	swaggerFileHandler := http.FileServer(http.Dir(dirPath))
	swaggerHandler := http.StripPrefix("/swaggerui/", swaggerFileHandler)
	s.PathPrefix("/swaggerui/").Handler(swaggerHandler)
	apiMux := s.PathPrefix("/api").Subrouter()
	v1Router := apiMux.PathPrefix("/v1").Subrouter()
	v2Router := apiMux.PathPrefix("/v2").Subrouter()
	s.registerV1Route(v1Router)
	s.registerV2Route(v2Router)
}

// registerV1Route used to handle the route starting with v1.
func (s Server) registerV1Route(router *mux.Router) {
	router.HandleFunc("/status", handler.MakeRest(dto.GetBookKeepingRequest{},
		rest.GetBookKeeping)).Methods(http.MethodGet)
	router.HandleFunc("/status", handler.MakeRest(dto.UpdateBookKeepingRequest{},
		rest.UpdateBookKeeping)).Methods(http.MethodPut)

	// Organization Route
	router.HandleFunc("/organization", handler.MakeRest(dto.GetOrganizationInfoRequest{},
		rest.GetOrganizationInfo)).Methods(http.MethodGet)
	router.HandleFunc("/organization", handler.MakeRest(dto.CreateOrganizationInfoRequest{},
		rest.CreateOrganizationInfo)).Methods(http.MethodPost)
	router.HandleFunc("/organization", handler.MakeRest(dto.UpdateOrganizationInfoRequest{},
		rest.UpdateOrganizationInfo)).Methods(http.MethodPatch)

	// Project Route
	router.HandleFunc("/project", handler.MakeRest(dto.GetProjectRequest{},
		rest.GetProject)).Methods(http.MethodGet)
	router.HandleFunc("/project", handler.MakeRest(dto.CreateProjectRequest{},
		rest.CreateProject)).Methods(http.MethodPost)
	router.HandleFunc("/project", handler.MakeRest(dto.UpdateProjectRequest{},
		rest.UpdateProject)).Methods(http.MethodPut)
	router.HandleFunc("/project", handler.MakeRest(dto.DeleteProjectRequest{},
		rest.DeleteProject)).Methods(http.MethodDelete)

	// Academic Route
	router.HandleFunc("/academic", handler.MakeRest(dto.GetAcademicRequest{},
		rest.GetAcademic)).Methods(http.MethodGet)
	router.HandleFunc("/academic", handler.MakeRest(dto.CreateAcademicRequest{},
		rest.CreateAcademic)).Methods(http.MethodPost)
	router.HandleFunc("/academic", handler.MakeRest(dto.UpdateAcademicDetailRequest{},
		rest.UpdateAcademic)).Methods(http.MethodPut)
	router.HandleFunc("/academic", handler.MakeRest(dto.DeleteAcademicRequest{},
		rest.DeleteAcademic)).Methods(http.MethodDelete)

	// Certificate Route
	router.HandleFunc("/certificate", handler.MakeRest(dto.GetCertificateRequest{},
		rest.GetCertificate)).Methods(http.MethodGet)
	router.HandleFunc("/certificate", handler.MakeRest(dto.CreateCertificateRequest{},
		rest.CreateCertificate)).Methods(http.MethodPost)
	router.HandleFunc("/certificate", handler.MakeRest(dto.UpdateCertificateRequest{},
		rest.UpdateCertificate)).Methods(http.MethodPut)
	router.HandleFunc("/certificate", handler.MakeRest(dto.DeleteCertificateRequest{},
		rest.DeleteCertificate)).Methods(http.MethodDelete)

	// s3 related routes

	// zip route need to implement
	router.HandleFunc("/zip-s3-files", handler.MakeRest(dto.CopyS3FolderNameRequest{},
		rest.ZipS3Folder)).Methods(http.MethodPut)

	// New Routes for Sprint 1
	// Personal Info Route
	router.HandleFunc("/personal-info", handler.MakeRest(dto.GetUserPersonalInfoRequest{},
		rest.GetPersonalInfo)).Methods(http.MethodGet)
	router.HandleFunc("/personal-info", handler.MakeRest(dto.CreatePersonalInfoRequest{},
		rest.CreatePersonalInfo)).Methods(http.MethodPost)
	router.HandleFunc("/personal-info", handler.MakeRest(dto.DeletePersonalInfoRequest{},
		rest.DeletePersonalInfo)).Methods(http.MethodDelete)
	router.HandleFunc("/personal-info", handler.MakeRest(dto.UpdatePersonalInfoRequest{},
		rest.UpdatePersonalInfo)).Methods(http.MethodPut)

	// ProfessionalInfo Routes
	router.HandleFunc("/professional-info", handler.MakeRest(dto.GetUserProfessionalInfoRequest{},
		rest.GetProfessionalInfo)).Methods(http.MethodGet)
	router.HandleFunc("/professional-info", handler.MakeRest(dto.DeleteProfessionalInfoRequest{},
		rest.DeleteProfessionalInfo)).Methods(http.MethodDelete)
	router.HandleFunc("/professional-info", handler.MakeRest(dto.CreateProfessionalInfoRequest{},
		rest.CreateProfessionalInfo)).Methods(http.MethodPost)
	router.HandleFunc("/professional-info", handler.MakeRest(dto.UpdateProfessionalInfoRequest{},
		rest.UpdateProfessionalInfo)).Methods(http.MethodPut)

	// Routes for M3
	router.HandleFunc("/upload-asset", handler.Make(rest.FileUpload)).Methods(http.MethodPut)

	router.HandleFunc("/certificates-files", handler.MakeRest(dto.GetCertificateRequest{},
		rest.GetS3Assets)).Methods(http.MethodGet)

	router.HandleFunc("/create-s3-folder", handler.MakeRest(dto.CreateS3FolderRequest{},
		rest.CreateS3Folder)).Methods(http.MethodPost)

	router.HandleFunc("/delete-assets", handler.MakeRest(dto.DeleteS3AssetsRequest{},
		rest.DeleteS3Assets)).Methods(http.MethodDelete)

	router.HandleFunc("/copy-s3-files", handler.MakeRest(dto.CopyS3FolderNameRequest{},
		rest.CopyS3Files)).Methods(http.MethodPut)

	router.HandleFunc("/rename-s3-folder", handler.MakeRest(dto.RenameS3FolderNameRequest{},
		rest.RenameS3Folder)).Methods(http.MethodPut)

	router.HandleFunc("/move-s3-assets", handler.MakeRest(dto.UpdateS3FolderRequest{},
		rest.MoveS3Assets)).Methods(http.MethodPut)

	//router.HandleFunc("/move-s3-files", handler.MakeRest(dto.CopyS3FolderNameRequest{},
	//	rest.MoveS3Files)).Methods(http.MethodPut)

}

// registerV2Route used to handle the route starting with v1.
func (s Server) registerV2Route(router *mux.Router) {
}

func (s Server) ServeHTTP() {
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.Router)
	srv := &http.Server{
		Handler: loggedRouter,
		Addr:    s.Address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}
	logrus.Fatal(srv.ListenAndServe())
}
