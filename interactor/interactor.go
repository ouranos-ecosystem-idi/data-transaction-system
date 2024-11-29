package interactor

import (
	"data-spaces-backend/infrastructure/auth"
	auth_client "data-spaces-backend/infrastructure/auth/client"
	"data-spaces-backend/infrastructure/persistence/datastore"
	"data-spaces-backend/infrastructure/traceabilityapi"
	"data-spaces-backend/infrastructure/traceabilityapi/client"
	"data-spaces-backend/presentation/http/echo/handler"
	"data-spaces-backend/usecase"

	firebase "firebase.google.com/go/v4"
	"gorm.io/gorm"
)

type (
	Interactor interface {
		NewAppHandler() handler.AppHandler
	}

	interactor struct {
		db                     *gorm.DB
		firebaseConfig         *firebase.Config
		host                   string
		isTraceabilityAccess   bool
		TraceabilityBaseURL    string
		TraceabilityAPIVersion string
		TraceabilityAPIKey     string
		AuthenticaterUrl       string
		DataSpaceApikey        string
	}
)

// NewInteractor
// Summary: This is function which creates new Interactor.
// input: db(*gorm.DB) DB
// input: fc(*firebase.Config) Firebase config
// input: host(string) host
// input: isTraceabilityAccess(bool) is traceability access
// input: traceabilityBaseURL(string) traceability base URL
// input: traceabilityAPIVersion(string) traceability API version
// input: traceabilityAPIKey(string) traceability API key
// input: authenticaterURL(string) authenticater URL
// input: dataSpaceAPIKey(string) data space API key
// output: (Interactor) Interactor object
func NewInteractor(
	db *gorm.DB,
	fc *firebase.Config,
	host string,
	isTraceabilityAccess bool,
	traceabilityBaseURL string,
	traceabilityAPIVersion string,
	traceabilityAPIKey string,
	authenticaterURL string,
	dataSpaceAPIKey string,
) Interactor {
	return &interactor{
		db,
		fc,
		host,
		isTraceabilityAccess,
		traceabilityBaseURL,
		traceabilityAPIVersion,
		traceabilityAPIKey,
		authenticaterURL,
		dataSpaceAPIKey,
	}
}

// appHandler
// Summary: This is structure which defines appHandler.
type appHandler struct {
	handler.AuthHandler
	handler.OuranosHandler
	handler.HealthCheckHandler
}

// NewAppHandler
// Summary: This is function which creates new AppHandler.
// output: (handler.AppHandler) AppHandler object
func (i *interactor) NewAppHandler() handler.AppHandler {
	var cfpHandler handler.ICfpHandler
	var cfpCertificationHandler handler.ICfpCertificationHandler
	var partsHandler handler.IPartsHandler
	var partsStructureHandler handler.IPartsStructureHandler
	var tradeHandler handler.ITradeHandler
	var statusHandler handler.IStatusHandler

	traceabilityCli := client.NewClient(i.TraceabilityAPIKey, i.TraceabilityAPIVersion, i.TraceabilityBaseURL)
	authCli := auth_client.NewClient(i.DataSpaceApikey, i.AuthenticaterUrl)

	// repository DI
	ouranosRepository := datastore.NewOuranosRepository(i.db)
	authAPIRepository := auth.NewAuthAPIRepository(authCli)
	traceabilityRepository := traceabilityapi.NewTraceabilityRepository(traceabilityCli)
	userRequestUsecase := usecase.NewVerifyUsecase(authAPIRepository)

	if i.isTraceabilityAccess {
		// TraceabilityAPI DI

		// usecase DI
		partsUsecase := usecase.NewPartsTraceabilityUsecase(traceabilityRepository)
		partsStructureTraceabilityUsecase := usecase.NewPartsStructureTraceabilityUsecase(traceabilityRepository)
		tradeTraceabilityUsecase := usecase.NewTradeTraceabilityUsecase(traceabilityRepository)
		statusUsecase := usecase.NewStatusTraceabilityUsecase(traceabilityRepository)
		cfpUsecase := usecase.NewCfpTraceabilityUsecase(traceabilityRepository)
		cfpCertificationUsecase := usecase.NewCfpCertificationTraceabilityUsecase(traceabilityRepository)

		// handler DI
		cfpHandler = handler.NewCfpHandler(cfpUsecase)
		cfpCertificationHandler = handler.NewCfpCertificationHandler(cfpCertificationUsecase)
		partsHandler = handler.NewPartsHandler(partsUsecase, partsStructureTraceabilityUsecase, i.host)
		partsStructureHandler = handler.NewPartsStructureHandler(partsStructureTraceabilityUsecase)
		tradeHandler = handler.NewTradeHandler(tradeTraceabilityUsecase, i.host)
		statusHandler = handler.NewStatusHandler(statusUsecase, i.host)
	} else {
		// DB DI

		// usecase DI
		cfpUsecase := usecase.NewCfpUsecase(ouranosRepository)
		cfpCertificationUsecase := usecase.NewCfpCertificationUsecase(ouranosRepository)
		partsDatastoreUsecase := usecase.NewPartsUsecase(ouranosRepository)
		partsStructureDatastoreUsecase := usecase.NewPartsStructureDatastoreUsecase(ouranosRepository)
		tradeUsecase := usecase.NewTradeUsecase(ouranosRepository)
		statusUsecase := usecase.NewStatusUsecase(ouranosRepository)

		// handler DI
		cfpHandler = handler.NewCfpHandler(cfpUsecase)
		cfpCertificationHandler = handler.NewCfpCertificationHandler(cfpCertificationUsecase)
		partsHandler = handler.NewPartsHandler(partsDatastoreUsecase, partsStructureDatastoreUsecase, i.host)
		partsStructureHandler = handler.NewPartsStructureHandler(partsStructureDatastoreUsecase)
		tradeHandler = handler.NewTradeHandler(tradeUsecase, i.host)
		statusHandler = handler.NewStatusHandler(statusUsecase, i.host)
	}
	healthCheckHandler := handler.NewHealthCheckHandler()

	// handler DI
	authHandler := handler.NewAuthHandler(
		userRequestUsecase,
	)
	ouranosHandler := handler.NewOuranosHandler(
		cfpHandler,
		cfpCertificationHandler,
		partsHandler,
		partsStructureHandler,
		tradeHandler,
		statusHandler,
	)

	// appHandler DI
	appHandler := &appHandler{
		AuthHandler:        authHandler,
		OuranosHandler:     ouranosHandler,
		HealthCheckHandler: healthCheckHandler,
	}
	return appHandler
}
