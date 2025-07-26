package checker

import (
	"fmt"
	"time"
)

// Result представляет результат проверки доступности веб-сайта.
type Result struct {
	URL          string    // URL проверяемого сайта
	Timestamp    time.Time // Время проведения проверки
	Success      bool      // Флаг успешности проверки
	StatusCode   int       // HTTP-статус ответа (например, 200, 404, 500 и т.д.)
	ResponseTime float64   // Время ответа сервера в миллисекундах
	Error        error     // Ошибка, если проверка завершилась неудачей
}

// String возвращает удобное для чтения строковое представление результата.
func (r Result) String() string {
	if r.Success {
		return fmt.Sprintf("[%s] Site %s responded successfully with status %d in %.2f ms.",
			r.Timestamp.Format(time.RFC3339), r.URL, r.StatusCode, r.ResponseTime)
	}
	return fmt.Sprintf("[%s] Site %s failed with error: %v.", r.Timestamp.Format(time.RFC3339), r.URL, r.Error)
}
