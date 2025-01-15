package recordservices

import recordmodel "github.com/petmeds24/backend/pkg/rest/src/models/record-model"

func (rs *RecordService) InsertMedicalRecordDao(medicalRecord recordmodel.RecordData) (map[string]interface{}, error) {
	return rs.recordDao.InsertMedicalRecordDao(medicalRecord)
}

func (rs *RecordService) GetMedicalRecordsByPetId(petId string) ([]map[string]interface{}, error) {
	return rs.recordDao.GetMedicalRecordsByPetId(petId)
}

func (rs *RecordService) GetMedicalRecordByRecordId(recordId string) (map[string]interface{}, error) {
	return rs.recordDao.GetMedicalRecordByRecordId(recordId)
}
