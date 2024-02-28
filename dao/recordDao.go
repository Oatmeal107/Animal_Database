/*
* @Author: Oatmeal107
* @Date:   2023/6/19 15:09
 */

package dao

import (
	"Animal_database/model"
	"strings"
)

// GetRecord  分页获取记录
func GetRecord(pageNum int, pageSize int) (records *[]model.Record, err error) {
	err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&records).Error

	return records, err
}

func UploadRecord(records *[]model.Record) (returnErr error, existRecordsIndex []string) {
	for _, record := range *records {
		if err := DB.Model(model.Record{}).Create(&record).Error; err != nil {
			returnErr = err
			data := strings.Split(err.Error(), " ")[5]
			existRecordsIndex = append(existRecordsIndex, data)
		}
	}
	return returnErr, existRecordsIndex

	//return DB.Model(model.Record{}).Create(&record).Error
}

// GetRecordByArea 根据地区获取记录 todo 现在是把所有数据都查出来, 后续可能优化成只查出数量, 目前时间可以接受
func GetRecordByArea(province string, city string, county string) (records *[]model.Record, err error) {
	query := DB.Model(model.Record{}) // note 动态构建查询条件
	if province != "" {
		query = query.Where("investigation_province = ?", province)
	}
	if city != "" {
		query = query.Where("investigation_city = ?", city)
	}
	if county != "" {
		query = query.Where("investigation_county = ?", county)
	}
	err = query.Find(&records).Error
	return records, err
}
