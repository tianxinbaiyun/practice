package rate

import (
	"gorm.io/gorm"
	"math"
	"time"
)

// AllowN 限制器计算并保存到数据库
func AllowN(db *gorm.DB, key string, burst, rate, cost, period float64) (
	allowed, remaining, retryAfter, resetAfter float64) {

	emissionInterval := period / rate
	increment := emissionInterval * cost
	burstOffset := emissionInterval * burst
	nowTime := float64(time.Now().UnixNano()) / 1000000000

	data := GetRateLimit(db, key)

	tat := data.Value
	if tat == 0 {
		tat = nowTime
	}
	tat = math.Max(tat, nowTime)
	newTat := tat + increment
	allowAt := newTat - burstOffset
	diff := nowTime - allowAt
	remaining = diff / emissionInterval
	if remaining < 0 {
		resetAfter = tat - nowTime
		retryAfter = diff * -1
		return 0, 0, retryAfter, resetAfter
	}
	resetAfter = newTat - nowTime
	newData := &RateLimit{
		Key:        key,
		Value:      newTat,
		ExpireTime: time.Now().Add(time.Duration(resetAfter) * time.Second),
		Rate:       rate,
		Burst:      burst,
		Period:     period,
	}
	if data.Key == "" {
		rowsAffected := CreateRateLimit(db, newData)
		if rowsAffected == 0 {
			time.Sleep(time.Duration(emissionInterval) * time.Second)
			return AllowN(db, key, burst, rate, cost, period)
		}
	} else {
		rowsAffected := SetRateLimit(db, newData, data.Value)
		if rowsAffected == 0 {
			time.Sleep(time.Duration(emissionInterval) * time.Second)
			return AllowN(db, key, burst, rate, cost, period)
		}
	}
	retryAfter = -1

	return cost, remaining, retryAfter, resetAfter
}

// RateLimit 限制器数据库结构体
type RateLimit struct {
	Key        string    `gorm:"size:256;primaryKey;comment:限额key" json:"key"`
	Value      float64   `gorm:"type:decimal(18,6);comment:存储值" json:"value"`
	ExpireTime time.Time `gorm:"comment:过期时间" json:"expire_time"`
	Rate       float64   `gorm:"type:decimal(18,6);comment:速度" json:"rate"`
	Burst      float64   `gorm:"type:decimal(18,6);comment:爆破" json:"burst"`
	Period     float64   `gorm:"comment:期间" json:"period"`
}

// GetRateLimit 根据key获取数据库限额记录
func GetRateLimit(db *gorm.DB, key string) (data *RateLimit) {
	data = &RateLimit{}
	_ = db.Where("`key` = ? ", key).First(data).Error
	return
}

// SetRateLimit 根据key更新数据库限额记录
func SetRateLimit(db *gorm.DB, data *RateLimit, oldValue float64) int64 {
	return db.Where("value=?", oldValue).Updates(data).RowsAffected
}

// CreateRateLimit 创建数据库限额记录
func CreateRateLimit(db *gorm.DB, data *RateLimit) int64 {
	return db.Create(data).RowsAffected
}

// DeleteRateLimit 创建数据库限额记录
func DeleteRateLimit(db *gorm.DB, key string) error {
	return db.Where("`key` = ?", key).Delete(&RateLimit{}).Error
}
