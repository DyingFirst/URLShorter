CREATE TABLE original_url_and_shorted_id_list (
    id BIGINT PRIMARY KEY,
    original_url TEXT,
    shorted_id VARCHAR(10) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP
);


-- хранимая функция, которая удаляет записи из таблицы original_url_and_shorted_id_list,
--      у которых значение update_at старше указанного количества минут.
-- !! Важно соблюдать единство системного времени в контейнере и системе,
--      которая передаёт это время в поля created_at и update_at, как входной параметр
CREATE OR REPLACE FUNCTION delete_old_records(minutes_old INTEGER)
RETURNS VOID AS $$
BEGIN
DELETE FROM original_url_and_shorted_id_list
WHERE update_at < (CURRENT_TIMESTAMP - interval '1 minute' * minutes_old);
END;
$$ LANGUAGE plpgsql;
