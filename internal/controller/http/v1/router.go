package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/bots-empire/ams-service/internal/entity"
	"github.com/bots-empire/ams-service/internal/service"
)

func HandleRouts(mux *http.ServeMux, m *service.Manager, logger *zap.Logger) {
	accessRouts(mux, m, logger)

	usersRouts(mux, m, logger)

	incomeInfoRouts(mux, m, logger)
}

func accessRouts(mux *http.ServeMux, m *service.Manager, logger *zap.Logger) {
	mux.HandleFunc("/v1/accesses/check", func(w http.ResponseWriter, req *http.Request) {
		acs, err := accessesFromRequest(req)
		if err != nil {
			logger.Warn("error parse entity", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed to parse entity: %v", err), http.StatusUnprocessableEntity)
			return
		}

		if err = acs.Validate(false); err != nil {
			logger.Warn("error validate access", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("error in validate access: %v", err), http.StatusUnprocessableEntity)
			return
		}

		ok, err := m.CheckAccess(context.Background(), acs)
		if err != nil {
			logger.Warn("error check access", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed check access: %v", err), http.StatusInternalServerError)
			return
		}

		if !ok {
			logger.Warn("forbidden response status")
			w.WriteHeader(http.StatusForbidden)
			return
		}
	})
	logger.Sugar().Info("handle rout: /v1/accesses/check")

	mux.HandleFunc("/v1/accesses/add", func(w http.ResponseWriter, req *http.Request) {
		acs, err := accessesFromRequest(req)
		if err != nil {
			logger.Warn("error parse entity", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed to parse entity: %v", err), http.StatusUnprocessableEntity)
			return
		}

		if err = acs.Validate(true); err != nil {
			logger.Warn("error validate access", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("error in validate access: %v", err), http.StatusUnprocessableEntity)
			return
		}

		err = m.AddAccess(context.Background(), acs)
		if err != nil {
			logger.Warn("error add access", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed add access: %v", err), http.StatusInternalServerError)
			return
		}
	})
	logger.Sugar().Info("handle rout: /v1/accesses/add")

	mux.HandleFunc("/v1/accesses/deprive", func(w http.ResponseWriter, req *http.Request) {
		acs, err := accessesFromRequest(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse access: %v", err), http.StatusUnprocessableEntity)
			return
		}

		err = m.DepriveAccess(context.Background(), acs)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed check access: %v", err), http.StatusInternalServerError)
			return
		}
	})
	logger.Sugar().Info("handle rout: /v1/accesses/deprive")
}

func usersRouts(mux *http.ServeMux, m *service.Manager, logger *zap.Logger) {
	mux.HandleFunc("/v1/admins/get", func(w http.ResponseWriter, req *http.Request) {
		query, err := adminQueryFromRequest(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse query: %v", err), http.StatusUnprocessableEntity)
			return
		}

		if err = query.Validate(); err != nil {
			logger.Warn("error validate query", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("error in validate query: %v", err), http.StatusUnprocessableEntity)
			return
		}

		ids, err := m.GetAdminsID(context.Background(), query)
		if err != nil {
			logger.Warn("error get admins", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed check access: %v", err), http.StatusInternalServerError)
			return
		}

		w.Write(marshalResponse(ids))
	})
	logger.Sugar().Info("handle rout: /v1/admins/get")

	mux.HandleFunc("/debug/admins/get-all", func(w http.ResponseWriter, req *http.Request) {
		admins, err := m.GetAllAdmins(context.Background())
		if err != nil {
			logger.Warn("error get all admins", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed check access: %v", err), http.StatusInternalServerError)
			return
		}

		w.Write(marshalResponse(admins))
	})
	logger.Sugar().Info("handle rout: /debug/admins/get-all")
}

func incomeInfoRouts(mux *http.ServeMux, m *service.Manager, logger *zap.Logger) {
	mux.HandleFunc("/v1/income-info/add", func(w http.ResponseWriter, req *http.Request) {
		incInfo, err := incomeInfoFromRequest(req)
		if req.Body == nil {
			logger.Warn("error nil body")
			http.Error(w, fmt.Sprintf("error nil body: %v", err), http.StatusUnprocessableEntity)
			return
		}

		//req.Close = true
		if err != nil {
			logger.Warn("error parse entity", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed to parse entity: %v", err), http.StatusUnprocessableEntity)
			return
		}

		if err = incInfo.ValidateAdd(); err != nil {
			logger.Warn("error validate income info", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("error in validate income info: %v", err), http.StatusUnprocessableEntity)
			return
		}

		err = m.AddIncomeInfo(context.Background(), incInfo)
		if err != nil {
			logger.Warn("error add income info", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed add income info: %v", err), http.StatusInternalServerError)
			return
		}
	})
	logger.Sugar().Info("handle rout: /v1/income-info/add")

	mux.HandleFunc("/v1/income-info/get", func(w http.ResponseWriter, req *http.Request) {
		query, err := incomeInfoFromRequest(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse query: %v", err), http.StatusUnprocessableEntity)
			return
		}

		if err = query.ValidateGet(); err != nil {
			logger.Warn("error validate query", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("error in validate query: %v", err), http.StatusUnprocessableEntity)
			return
		}

		ids, err := m.GetIncomeInfo(context.Background(), query.UserID, query.TypeBot)
		if err != nil {
			logger.Warn("error get income info", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed check income info: %v", err), http.StatusInternalServerError)
			return
		}

		w.Write(marshalResponse(ids))
	})
	logger.Sugar().Info("handle rout: /v1/income-info/get")

	mux.HandleFunc("/debug/income-info/get-all", func(w http.ResponseWriter, req *http.Request) {
		incInfo, err := m.GetAllIncomeInfo(context.Background())
		if err != nil {
			logger.Warn("error get all income info", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed check income info: %v", err), http.StatusInternalServerError)
			return
		}

		w.Write((marshalResponse(incInfo)))
	})
	logger.Sugar().Info("handle rout: /debug/income-info/get-all")

	mux.HandleFunc("/v1/income-info/delete", func(w http.ResponseWriter, req *http.Request) {
		incInfo, err := incomeInfoFromRequest(req)
		if err != nil {
			logger.Warn("error parse entity", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed to parse entity: %v", err), http.StatusUnprocessableEntity)
			return
		}

		if err = incInfo.ValidateAdd(); err != nil {
			logger.Warn("error validate income info", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("error in validate income info: %v", err), http.StatusUnprocessableEntity)
			return
		}

		err = m.DeleteIncomeInfo(context.Background(), incInfo.UserID)
		if err != nil {
			logger.Warn("error delete income info", zap.Any("err", err))
			http.Error(w, fmt.Sprintf("failed deletet income info: %v", err), http.StatusInternalServerError)
			return
		}
	})
	logger.Sugar().Info("handle rout: /v1/income-info/delete")
}

func incomeInfoFromRequest(req *http.Request) (*entity.IncomeInfo, error) {
	decoder := json.NewDecoder(req.Body)
	var i entity.IncomeInfo
	if err := decoder.Decode(&i); err != nil {
		return nil, err
	}
	return &i, nil
}

func accessesFromRequest(req *http.Request) (*entity.Access, error) {
	decoder := json.NewDecoder(req.Body)
	var a entity.Access
	err := decoder.Decode(&a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func adminQueryFromRequest(req *http.Request) (*entity.AdminsQuery, error) {
	decoder := json.NewDecoder(req.Body)
	var a entity.AdminsQuery
	err := decoder.Decode(&a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func marshalResponse(data interface{}) []byte {
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return resp
}
