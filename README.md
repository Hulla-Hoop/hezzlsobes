# hezzlsobes
Сервис golang,Postgres,Clickhouse,Kafka

ТЗ

1) Развернуть сервис на Golang, Postgres, Clickhouse, Nats (альтернатива kafka), Redis 
2) Описать модели данных и миграций 
3) В миграциях Postgres 
 a) Проставить primary-key и индексы на указанные поля
 b) При добавлении записи в таблицу устанавливать приоритет как макс приоритет в таблице +1. Приоритеты начинаются с 1
 c) При накатке миграций добавить одну запись в Campaigns таблицу по умолчанию
    id = serial
    name = Первая запись
4) Реализовать CRUD методы на GET-POST-PATCH-DELETE данных в таблице GOODS в Postgres 
5) При редактировании данных в Postgres ставить блокировку на чтение записи и оборачивать все в транзакцию. Валидируем поля при редактировании.
6) При редактировании данных в GOODS инвалидируем данные в REDIS 

------При редактировании записи удаляются из Redis

7) Если записи нет (проверяем на PATCH-DELETE), выдаем ошибку (статус 404) 
   - code = 3
   - message = “errors.good.notFound“
   - details = {}

------Сделано

8) При GET запросе данных из Postgres кешировать данные в Redis на минуту. Пытаемся получить данные сперва из Redis, если их нет, идем в БД и кладем их в REDIS 

------Структура List кешируется не полностью(goods кэшируются в Redis, а структруа Meta нет) сделано для того чтобы при обращении к разным страницам записи не брались полностью с бд
------При отсутствие структруы в Redis она берется в базе


9) При добавлении, редактировании или удалении записи в Postgres писать лог в Clickhouse через очередь Nats (альтернатива kafka). Логи писать пачками в Clickhouse :

------Логи пишутся пачками по 5 записей.

10) При обращении в БД использовать чистый SQL:

------Используется чистый SQL