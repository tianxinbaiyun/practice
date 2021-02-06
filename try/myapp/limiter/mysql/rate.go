package rate

import (
	"fmt"
	"gorm.io/gorm"
	"math"
	"time"
)

const prefix = "rate:"

// Limit 限制结构体
type Limit struct {
	Rate   float64
	Burst  float64
	Period time.Duration
}

// String 限制字符串
func (l Limit) String() string {
	return fmt.Sprintf("%f req/%s (burst %f)", l.Rate, fmtDur(l.Period), l.Burst)
}

// IsZero 判断是否为空
func (l Limit) IsZero() bool {
	return l == Limit{}
}

// fmtDur 速率单位判断
func fmtDur(d time.Duration) string {
	switch d {
	case time.Second:
		return "s"
	case time.Minute:
		return "m"
	case time.Hour:
		return "h"
	}
	return d.String()
}

// PerSecond 每秒限制速率
func PerSecond(rate float64) Limit {
	return Limit{
		Rate:   rate,
		Period: time.Second,
		Burst:  rate,
	}
}

// PerMinute 每分钟限制速率
func PerMinute(rate float64) Limit {
	return Limit{
		Rate:   rate,
		Period: time.Minute,
		Burst:  rate,
	}
}

// PerHour 每小时限制速率
func PerHour(rate float64) Limit {
	return Limit{
		Rate:   rate,
		Period: time.Hour,
		Burst:  rate,
	}
}

//------------------------------------------------------------------------------

// Limiter 限制器结构体
type Limiter struct {
	DB *gorm.DB
}

// NewLimiter 初始化限制器
func NewLimiter(db *gorm.DB) *Limiter {
	return &Limiter{
		DB: db,
	}
}

// Allow 方法是AllowN的简化方法
func (l Limiter) Allow(key string, limit Limit) (*Result, error) {
	return l.AllowN(key, limit, 1)
}

// AllowN 返回允许执行的情况
func (l Limiter) AllowN(key string, limit Limit, n float64) (*Result, error) {
	key = prefix + key
	allowed, remaining, retryAfter, resetAfter := AllowN(l.DB, key, limit.Burst, limit.Rate, n, limit.Period.Seconds())
	res := &Result{
		Limit:      limit,
		Allowed:    int(math.Ceil(allowed)),
		Remaining:  int(math.Ceil(remaining)),
		RetryAfter: dur(retryAfter),
		ResetAfter: dur(resetAfter),
	}
	return res, nil
}

// Reset 重置限制器
func (l *Limiter) Reset(key string) error {
	return DeleteRateLimit(l.DB, key)
}

// dur 把float64类型转成时间类型
func dur(f float64) time.Duration {
	if f == -1 {
		return -1
	}
	return time.Duration(f * float64(time.Second))
}

// Result 限制器返回的结构
type Result struct {
	Limit      Limit
	Allowed    int
	Remaining  int
	RetryAfter time.Duration
	ResetAfter time.Duration
}
