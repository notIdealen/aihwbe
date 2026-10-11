package domain

type ErrorCode string

const (
	ErrNotFound       ErrorCode = "404 not found"
	ErrBadRequest     ErrorCode = "400 bad request"
	ErrInternalServer ErrorCode = "500 internal server"
)

// 400 Bad Request — некорректный или повреждённый запрос.
// 401 Unauthorized — отсутствие valid authentication credentials (не пройдена авторизация).
// 403 Forbidden — доступ к ресурсу запрещён.
// 404 Not Found — запрашиваемая страница не найдена (неверный URL).
// 405 Method Not Allowed — метод запроса не поддерживается для данного ресурса.
// 429 Too Many Requests — превышено ограничение по количеству запросов (слишком много запросов от одного клиента).

// 500 Internal Server Error — внутренняя ошибка сервера (например, из-за проблем в коде или недоступности базы данных).
// 502 Bad Gateway — сервер, выступая в роли шлюза или прокси, получил некорректный ответ от вышестоящего сервера.
// 503 Service Unavailable — сервер временно не может обработать запрос (перегрузка или техническое обслуживание).
// 504 Gateway Timeout — сервер не получил ответ от вышестоящего сервера в установленное время.

// 400 Bad Request («неправильный, некорректный запрос»)[4][5][8];
// 401 Unauthorized («не авторизован (не представился)»)[4][5];
// 402 Payment Required («необходима оплата»)[4][5];
// 403 Forbidden («запрещено (не уполномочен)»)[4][5];
// 404 Not Found («не найдено»)[4][5];
// 405 Method Not Allowed («метод не поддерживается»)[4][5];
// 406 Not Acceptable («неприемлемо»)[4][5];
// 407 Proxy Authentication Required («необходима аутентификация прокси»)[4][5];
// 408 Request Timeout («истекло время ожидания»)[4][5];
// 409 Conflict («конфликт»)[4][5][8];
// 410 Gone («удалён»)[4][5];
// 411 Length Required («необходима длина»)[4][5];
// 412 Precondition Failed («условие ложно»)[4][5][13];
// 413 Content Too Large («содержимое слишком велико», ранее — Payload Too Large)[4][5][1];
// 414 URI Too Long («URI слишком длинный»)[4][5][1];
// 415 Unsupported Media Type («неподдерживаемый тип данных»)[4][5];
// 416 Range Not Satisfiable («диапазон не достижим»)[5];
// 417 Expectation Failed («ожидание не удалось»)[5];
// 418 I’m a teapot («я — чайник»);
// 419 Authentication Timeout (not in RFC 2616) («обычно ошибка проверки CSRF»);
// 421 Misdirected Request [14];
// 422 Unprocessable Content («необрабатываемое содержимое», ранее — Unprocessable Entity)[1];
// 423 Locked («заблокировано»);
// 424 Failed Dependency («невыполненная зависимость»);
// 425 Too Early («слишком рано»);
// 426 Upgrade Required («необходимо обновление»);
// 428 Precondition Required («необходимо предусловие»)[15];
// 429 Too Many Requests («слишком много запросов»)[15];
// 431 Request Header Fields Too Large («поля заголовка запроса слишком большие»)[15];
// 449 Retry With («повторить с»)[3];
// 451 Unavailable For Legal Reasons («недоступно по юридическим причинам»)[16].
// 499 Client Closed Request (клиент закрыл соединение);
//
// 5xx: Server Error (ошибка сервера):
// 500 Internal Server Error («внутренняя ошибка сервера»)[4][5];
// 501 Not Implemented («не реализовано»)[4][5];
// 502 Bad Gateway («плохой, ошибочный шлюз»)[4][5];
// 503 Service Unavailable («сервис недоступен»)[4][5];
// 504 Gateway Timeout («шлюз не отвечает»)[4][5];
// 505 HTTP Version Not Supported («версия HTTP не поддерживается»)[4][5];
// 506 Variant Also Negotiates («вариант тоже проводит согласование»)[17];
// 507 Insufficient Storage («переполнение хранилища»);
// 508 Loop Detected («обнаружено бесконечное перенаправление»)[18];
// 509 Bandwidth Limit Exceeded («исчерпана пропускная ширина канала»);
// 510 Not Extended («не расширено»);
// 511 Network Authentication Required («требуется сетевая аутентификация»)[15];
// 520 Unknown Error («неизвестная ошибка»)[19];
// 521 Web Server Is Down («веб-сервер не работает»)[19];
// 522 Connection Timed Out («соединение не отвечает»)[19];
// 523 Origin Is Unreachable («источник недоступен»)[19];
// 524 A Timeout Occurred («время ожидания истекло»)[19];
// 525 SSL Handshake Failed («квитирование SSL не удалось»)[19];
// 526 Invalid SSL Certificate («недействительный сертификат SSL»)[19].
