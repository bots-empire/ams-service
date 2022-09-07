package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"ams-service/internal/entity"
	"ams-service/internal/service"
)

func HandleRouts(mux *http.ServeMux, m *service.Manager, logger *zap.Logger) {
	accessRouts(mux, m, logger)

	usersRouts(mux, m, logger)
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
