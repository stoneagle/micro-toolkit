package services

import (
	"time"
	"toolkit/backend/models"

	"github.com/go-xorm/xorm"
)

type DeviceScan struct {
	engine *xorm.Engine
}

func NewDeviceScan(engine *xorm.Engine) *DeviceScan {
	return &DeviceScan{
		engine: engine,
	}
}

func (s *DeviceScan) One(id int) (scan models.DeviceScan, err error) {
	scan = models.DeviceScan{}
	_, err = s.engine.Where("id = ?", id).Get(&scan)
	return
}

func (s *DeviceScan) List(appId string) (scans []models.DeviceScan, err error) {
	scans = make([]models.DeviceScan, 0)
	err = s.engine.Where("app_id = ?", appId).Find(&scans)
	return
}

func (s *DeviceScan) Delete(id int) (err error) {
	var scan models.DeviceScan
	_, err = s.engine.Id(id).Get(&scan)
	if err == nil {
		_, err = s.engine.Id(id).Delete(&scan)
	}
	return
}

func (s *DeviceScan) Add(scan *models.DeviceScan) (err error) {
	scan.UpdatedAt = int(time.Now().Unix())
	scan.CreatedAt = int(time.Now().Unix())
	_, err = s.engine.Insert(scan)
	return
}

func (s *DeviceScan) UpdateByMap(id int, scan map[string]interface{}) (err error) {
	scan["updated_at"] = time.Now().Unix()
	_, err = s.engine.Table(new(models.DeviceScan)).Id(id).Update(&scan)
	return
}

func (s *DeviceScan) Update(id int, scan *models.DeviceScan) (err error) {
	_, err = s.engine.Id(id).Update(scan)
	return
}
