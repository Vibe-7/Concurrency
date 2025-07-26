Concurrency
**Concurrency** — на Go для мониторинга доступности сайтов с асинхронной отправкой уведомлений через разные каналы: консоль, файл, Telegram и e-mail.

ОСНОВАЯ ЛОГИКА

Загрузка конфигурации из config.json:

Список URL для проверки (urls).

Таймаут для HTTP-запросов (timeout).

**Параллельная проверка** каждого URL с помощью горутин и sync.WaitGroup:

HEAD-запросы через checker.CheckURL с вложенным context.Context.

Измерение времени отклика и анализ StatusCode.

**Сбор уведомлений**:

Ошибочные ответы (код ≥ 400 или сетевые ошибки) шлются в канал notifyCh.

Успешные ответы логируются через logger.Info.

Отправка уведомлений в отдельной горутине:

Подписка на канал notifyCh.

Рассылка через все подключённые Notifier'ы.

 Используемые технологии

**Go** — язык программирования.

**goroutines** и **channels** — для конкурентной обработки.

**context** — для отмены и таймаутов.

**sync.WaitGroup** — ожидание завершения всех проверок.

**flag** — для чтения параметров командной строки.

**encoding/json** — для парсинга config.json.

**godotenv** — загрузка переменных окружения из .env.

**net/http** — выполнение HEAD-запросов.

  **Нотификаторы** (Notifier)

Интерфейс Notifier описывает метод SendNotification(msg string). Реализованы:

ConsoleNotifier — вывод в консоль через log.Printf.

FileNotifier — запись в файл alerts.log.

TelegramNotifier — отправка сообщения через Telegram Bot API.

EmailNotifier — отправка письма по SMTP.
