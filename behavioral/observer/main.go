package main

import "fmt"

// Observer определяет интерфейс для всех подписчиков,
// которые должны получать уведомления от издателя.
type Observer interface {
	Update(data string)
}

// Publisher (Издатель) управляет списком подписчиков и рассылает им уведомления.
type Publisher struct {
	observers []Observer // Список зарегистрированных наблюдателей
}

// Subscribe добавляет нового подписчика в список.
func (p *Publisher) Subscribe(o Observer) {
	p.observers = append(p.observers, o)
}

// Unsubscribe находит и удаляет подписчика из списка.
func (p *Publisher) Unsubscribe(o Observer) {
	for i, observer := range p.observers {
		if observer == o {
			// Вырезаем найденный элемент из среза по индексу
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

// Notify оповещает всех текущих подписчиков о произошедшем событии.
func (p *Publisher) Notify(data string) {
	for _, observer := range p.observers {
		observer.Update(data)
	}
}

// EmailListener — конкретный подписчик для отправки уведомлений по почте.
type EmailListener struct {
	Email string
}

// Update выполняет бизнес-логику отправки email при получении события.
func (e *EmailListener) Update(data string) {
	fmt.Printf("Отправка email на %s: %s\n", e.Email, data)
}

// LoggingListener — конкретный подписчик для записи событий в лог-файл.
type LoggingListener struct {
	FilePath string
}

// Update выполняет бизнес-логику записи события в файл.
func (l *LoggingListener) Update(data string) {
	fmt.Printf("Запись в лог %s: %s\n", l.FilePath, data)
}

func main() {
	// 1. Создаем объект издателя
	publisher := &Publisher{}

	// 2. Создаем экземпляры подписчиков
	emailListener := &EmailListener{Email: "user@example.com"}
	loggingListener := &LoggingListener{FilePath: "/var/log/app.log"}

	// 3. Регистрируем подписчиков у издателя
	publisher.Subscribe(emailListener)
	publisher.Subscribe(loggingListener)

	// 4. Отправляем первое событие всем зарегистрированным подписчикам
	publisher.Notify("Вышла новая статья!")

	// 5. Удаляем одного из подписчиков из списка
	publisher.Unsubscribe(emailListener)

	// 6. Отправляем второе событие — его получит только оставшийся LoggingListener
	publisher.Notify("Вышел новый видеоурок!")
}
