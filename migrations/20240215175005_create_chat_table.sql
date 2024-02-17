-- +goose Up
-- +goose StatementBegin
--------------------------- ЧАТ ------------------------------------
CREATE TABLE IF NOT EXISTS chat (
    id serial primary key,
    owner int
);

COMMENT ON TABLE chat IS 'Таблица с чатами';
COMMENT ON COLUMN chat.id IS 'Id чата';
COMMENT ON COLUMN chat.owner IS 'Id создателя чата';

--------------------------- ПОЛЬЗОВАТЕЛЬ ЧАТА ------------------------------------
CREATE TABLE IF NOT EXISTS chat_user (
    id serial primary key,
    chat_id int not null,
    name varchar(255) not null,
    foreign key (chat_id) references chat (id) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMENT ON TABLE chat_user IS 'Таблица с пользователями определенного чата';
COMMENT ON COLUMN chat_user.id IS 'Id пользователя';
COMMENT ON COLUMN chat_user.name IS 'Имя пользователя';
COMMENT ON COLUMN chat_user.chat_id IS 'Id чата, внешний ключ (связь с таблицей чатов)';

--------------------------- СООБЩЕНИЯ ------------------------------------

CREATE TABLE IF NOT EXISTS message (
    id serial primary key,
    timestamp timestamp not null default now(),
    chat_id int not null,
    user_id int,
    text text not null,
    foreign key (chat_id) references chat (id) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMENT ON TABLE message IS 'Таблица с сообщениями в чате';
COMMENT ON COLUMN message.id IS 'Id сообщения';
COMMENT ON COLUMN message.timestamp IS 'Время отправки сообщения';
COMMENT ON COLUMN message.chat_id IS 'Id чата, внешний ключ (связь с таблицей чатов)';
COMMENT ON COLUMN message.user_id IS 'Id пользователя отправившего сообщение';
COMMENT ON COLUMN message.text IS 'Текст сообщения';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table message;
drop table chat_user;
drop table chat;
-- +goose StatementEnd
