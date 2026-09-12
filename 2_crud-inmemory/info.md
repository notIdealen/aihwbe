ctrl+shift+v - просмотр маркдауна или редактировать

## Endpoints
```text
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PUT    /tasks/{id}
DELETE /tasks/{id}
```

Задача:
```json
{
  "id": "uuid",
  "title": "Купить молоко",
  "description": "В магазине",
  "status": "todo",
  "created_at": "2026-08-15T12:00:00Z",
  "updated_at": "2026-08-15T12:00:00Z"
}
```

3. Нельзя возвращать внутренние ошибки напрямую пользователю.
4. Валидируй `title`.
5. Возвращай правильные HTTP-статусы:
   - `201 Created`;
   - `200 OK`;
   - `204 No Content`;
   - `400 Bad Request`;
   - `404 Not Found`;
   - `500 Internal Server Error`.
