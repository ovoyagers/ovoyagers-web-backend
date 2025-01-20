package formservice

import (
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/handlers/formdao"
	"github.com/petmeds24/backend/pkg/rest/src/models/formmodel"
)

type FormService struct {
	formDao *formdao.FormDao
}

func NewFormService(globalCfg *config.GlobalConfig) *FormService {
	return &FormService{formDao: formdao.NewFormDao(globalCfg)}
}

func (f *FormService) CreateForm(form *formmodel.Form) (string, error) {
	return f.formDao.CreateForm(form)
}

func (f *FormService) GetForms(limit, page, offset int) ([]map[string]interface{}, int, error) {
	return f.formDao.GetForms(limit, page, offset)
}

func (f *FormService) GetFormsByCategory(category string, limit, page, offset int) ([]map[string]interface{}, int, error) {
	return f.formDao.GetFormsByCategory(category, limit, page, offset)
}
