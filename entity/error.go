package entity

import "errors"

var (
	ErrNotFound           = errors.New("не найдено")
	ErrRequestFailed      = errors.New("ошибка запроса")
	ErrNameInvalid        = errors.New("некорректное имя")
	ErrStatusInvalid      = errors.New("некорректный статус")
	ErrUUIDInvalid        = errors.New("некорректный UUID")
	ErrJSONParseFailed    = errors.New("невозможно разобрать JSON")
	ErrTaskCreationFailed = errors.New("не удалось создать задачу")
	ErrTaskFetchFailed    = errors.New("не удалось получить задачи")
	ErrTaskUpdateFailed   = errors.New("не удалось обновить задачу")
	ErrTaskDeleteFailed   = errors.New("не удалось удалить задачу")
	ErrInvalidTaskID      = errors.New("некорректный ID задачи")
	ErrDatabaseConnection = errors.New("не удалось подключиться к базе данных")
)
