package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"Task2API/internal/models"
	"Task2API/pkg/logger"
)

type TargetAPIConfig struct {
	URL           string
	Authorization string
}

type PredictionService struct {
	config TargetAPIConfig
	logger logger.Logger
}

func NewPredictionService(config TargetAPIConfig, logger logger.Logger) *PredictionService {
	return &PredictionService{
		config: config,
		logger: logger,
	}
}

func (s *PredictionService) PredictHBA1C(ctx context.Context, params url.Values) ([]byte, error) {
	request, err := s.createHBA1CRequest(params)
	if err != nil {
		return nil, err
	}
	return s.sendRequest(ctx, request, s.config.URL+"/predict/hba1c")
}

func (s *PredictionService) PredictLDL(ctx context.Context, params url.Values) ([]byte, error) {
	request, err := s.createLDLRequest(params)
	if err != nil {
		return nil, err
	}
	return s.sendRequest(ctx, request, s.config.URL+"/predict/ldl")
}

func (s *PredictionService) PredictLDLL(ctx context.Context, params url.Values) ([]byte, error) {
	request, err := s.createLDLLRequest(params)
	if err != nil {
		return nil, err
	}
	return s.sendRequest(ctx, request, s.config.URL+"/predict/ldll")
}

func (s *PredictionService) PredictFERR(ctx context.Context, params url.Values) ([]byte, error) {
	request, err := s.createFERRRequest(params)
	if err != nil {
		return nil, err
	}
	return s.sendRequest(ctx, request, s.config.URL+"/predict/ferr")
}

func (s *PredictionService) PredictTG(ctx context.Context, params url.Values) ([]byte, error) {
	request, err := s.createTGRequest(params)
	if err != nil {
		return nil, err
	}
	return s.sendRequest(ctx, request, s.config.URL+"/predict/tg")
}

func (s *PredictionService) PredictHDL(ctx context.Context, params url.Values) ([]byte, error) {
	request, err := s.createHDLRequest(params)
	if err != nil {
		return nil, err
	}
	return s.sendRequest(ctx, request, s.config.URL+"/predict/hdl")
}

func (s *PredictionService) sendRequest(ctx context.Context, request interface{}, url string) ([]byte, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	s.logger.Debug("Sending request to %s: %s", url, string(requestBody))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", s.config.Authorization)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	s.logger.Debug("Received response from %s: %s", url, string(responseBody))

	if resp.StatusCode >= 400 {
		if resp.StatusCode == 500 {
			return nil, fmt.Errorf("target API returned status %d: service temporarily unavailable", resp.StatusCode)
		}
		return nil, fmt.Errorf("target API returned status %d: %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}

func (s *PredictionService) createHBA1CRequest(params url.Values) (*models.HBA1CRequest, error) {
	request := &models.HBA1CRequest{
		BaseRequest: models.BaseRequest{
			UID:    "web-client",
			Age:    25,
			Gender: 1,
		},
	}

	var err error

	if ageStr := params.Get("age"); ageStr != "" {
		request.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			return nil, fmt.Errorf("invalid age parameter: %v", err)
		}
	}

	if genderStr := params.Get("gender"); genderStr != "" {
		request.Gender, err = strconv.Atoi(genderStr)
		if err != nil {
			return nil, fmt.Errorf("invalid gender parameter: %v", err)
		}
	}

	floatFields := map[string]*float64{
		"rdw":  &request.RDW,
		"wbc":  &request.WBC,
		"rbc":  &request.RBC,
		"hgb":  &request.HGB,
		"hct":  &request.HCT,
		"mcv":  &request.MCV,
		"mch":  &request.MCH,
		"mchc": &request.MCHC,
		"plt":  &request.PLT,
		"neu":  &request.NEU,
		"eos":  &request.EOS,
		"bas":  &request.BAS,
		"lym":  &request.LYM,
		"mon":  &request.MON,
		"soe":  &request.SOE,
		"chol": &request.CHOL,
		"glu":  &request.GLU,
	}

	for param, field := range floatFields {
		if valStr := params.Get(param); valStr != "" {
			*field, err = strconv.ParseFloat(valStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid %s parameter: %v", param, err)
			}
		}
	}

	return request, nil
}

func (s *PredictionService) createLDLRequest(params url.Values) (*models.LDLRequest, error) {
	request := &models.LDLRequest{
		BaseRequest: models.BaseRequest{
			UID:    "web-client",
			Age:    25,
			Gender: 1,
		},
	}

	var err error

	if ageStr := params.Get("age"); ageStr != "" {
		request.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			return nil, fmt.Errorf("invalid age parameter: %v", err)
		}
	}

	if genderStr := params.Get("gender"); genderStr != "" {
		request.Gender, err = strconv.Atoi(genderStr)
		if err != nil {
			return nil, fmt.Errorf("invalid gender parameter: %v", err)
		}
	}

	floatFields := map[string]*float64{
		"rdw":  &request.RDW,
		"wbc":  &request.WBC,
		"rbc":  &request.RBC,
		"hgb":  &request.HGB,
		"hct":  &request.HCT,
		"mcv":  &request.MCV,
		"mch":  &request.MCH,
		"mchc": &request.MCHC,
		"plt":  &request.PLT,
		"neu":  &request.NEU,
		"eos":  &request.EOS,
		"bas":  &request.BAS,
		"lym":  &request.LYM,
		"mon":  &request.MON,
		"soe":  &request.SOE,
		"chol": &request.CHOL,
		"glu":  &request.GLU,
	}

	for param, field := range floatFields {
		if valStr := params.Get(param); valStr != "" {
			*field, err = strconv.ParseFloat(valStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid %s parameter: %v", param, err)
			}
		}
	}

	return request, nil
}

func (s *PredictionService) createLDLLRequest(params url.Values) (*models.LDLLRequest, error) {
	request := &models.LDLLRequest{
		BaseRequest: models.BaseRequest{
			UID:    "web-client",
			Age:    25,
			Gender: 1,
		},
	}

	var err error

	if ageStr := params.Get("age"); ageStr != "" {
		request.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			return nil, fmt.Errorf("invalid age parameter: %v", err)
		}
	}

	if genderStr := params.Get("gender"); genderStr != "" {
		request.Gender, err = strconv.Atoi(genderStr)
		if err != nil {
			return nil, fmt.Errorf("invalid gender parameter: %v", err)
		}
	}

	floatFields := map[string]*float64{
		"chol": &request.CHOL,
		"hdl":  &request.HDL,
		"tg":   &request.TG,
	}

	for param, field := range floatFields {
		if valStr := params.Get(param); valStr != "" {
			*field, err = strconv.ParseFloat(valStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid %s parameter: %v", param, err)
			}
		}
	}

	return request, nil
}

func (s *PredictionService) createFERRRequest(params url.Values) (*models.FERRRequest, error) {
	request := &models.FERRRequest{
		BaseRequest: models.BaseRequest{
			UID:    "web-client",
			Age:    25,
			Gender: 1,
		},
	}

	var err error

	if ageStr := params.Get("age"); ageStr != "" {
		request.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			return nil, fmt.Errorf("invalid age parameter: %v", err)
		}
	}

	if genderStr := params.Get("gender"); genderStr != "" {
		request.Gender, err = strconv.Atoi(genderStr)
		if err != nil {
			return nil, fmt.Errorf("invalid gender parameter: %v", err)
		}
	}

	floatFields := map[string]*float64{
		"rdw":  &request.RDW,
		"wbc":  &request.WBC,
		"rbc":  &request.RBC,
		"hgb":  &request.HGB,
		"hct":  &request.HCT,
		"mcv":  &request.MCV,
		"mch":  &request.MCH,
		"mchc": &request.MCHC,
		"plt":  &request.PLT,
		"neu":  &request.NEU,
		"eos":  &request.EOS,
		"bas":  &request.BAS,
		"lym":  &request.LYM,
		"mon":  &request.MON,
		"soe":  &request.SOE,
		"crp":  &request.CRP,
	}

	for param, field := range floatFields {
		if valStr := params.Get(param); valStr != "" {
			*field, err = strconv.ParseFloat(valStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid %s parameter: %v", param, err)
			}
		}
	}

	return request, nil
}

func (s *PredictionService) createTGRequest(params url.Values) (*models.TGRequest, error) {
	request := &models.TGRequest{
		BaseRequest: models.BaseRequest{
			UID:    "web-client",
			Age:    25,
			Gender: 1,
		},
	}

	var err error

	if ageStr := params.Get("age"); ageStr != "" {
		request.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			return nil, fmt.Errorf("invalid age parameter: %v", err)
		}
	}

	if genderStr := params.Get("gender"); genderStr != "" {
		request.Gender, err = strconv.Atoi(genderStr)
		if err != nil {
			return nil, fmt.Errorf("invalid gender parameter: %v", err)
		}
	}

	floatFields := map[string]*float64{
		"rdw":  &request.RDW,
		"wbc":  &request.WBC,
		"rbc":  &request.RBC,
		"hgb":  &request.HGB,
		"hct":  &request.HCT,
		"mcv":  &request.MCV,
		"mch":  &request.MCH,
		"mchc": &request.MCHC,
		"plt":  &request.PLT,
		"neu":  &request.NEU,
		"eos":  &request.EOS,
		"bas":  &request.BAS,
		"lym":  &request.LYM,
		"mon":  &request.MON,
		"soe":  &request.SOE,
		"chol": &request.CHOL,
		"glu":  &request.GLU,
	}

	for param, field := range floatFields {
		if valStr := params.Get(param); valStr != "" {
			*field, err = strconv.ParseFloat(valStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid %s parameter: %v", param, err)
			}
		}
	}

	return request, nil
}

func (s *PredictionService) createHDLRequest(params url.Values) (*models.HDLRequest, error) {
	request := &models.HDLRequest{
		BaseRequest: models.BaseRequest{
			UID:    "web-client",
			Age:    25,
			Gender: 1,
		},
	}

	var err error

	if ageStr := params.Get("age"); ageStr != "" {
		request.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			return nil, fmt.Errorf("invalid age parameter: %v", err)
		}
	}

	if genderStr := params.Get("gender"); genderStr != "" {
		request.Gender, err = strconv.Atoi(genderStr)
		if err != nil {
			return nil, fmt.Errorf("invalid gender parameter: %v", err)
		}
	}

	floatFields := map[string]*float64{
		"rdw":  &request.RDW,
		"wbc":  &request.WBC,
		"rbc":  &request.RBC,
		"hgb":  &request.HGB,
		"hct":  &request.HCT,
		"mcv":  &request.MCV,
		"mch":  &request.MCH,
		"mchc": &request.MCHC,
		"plt":  &request.PLT,
		"neu":  &request.NEU,
		"eos":  &request.EOS,
		"bas":  &request.BAS,
		"lym":  &request.LYM,
		"mon":  &request.MON,
		"soe":  &request.SOE,
		"chol": &request.CHOL,
		"glu":  &request.GLU,
	}

	for param, field := range floatFields {
		if valStr := params.Get(param); valStr != "" {
			*field, err = strconv.ParseFloat(valStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid %s parameter: %v", param, err)
			}
		}
	}

	return request, nil
}
