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