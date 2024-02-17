-- +goose Up
-- +goose StatementBegin
--------------------------- ЧАТ ------------------------------------
create table if not exists chat (
    id serial primary key,
    owner int
);

comment on table chat is 'Таблица с чатами';
comment on column chat.id is 'Id чата';
comment on column chat.owner is 'Id создателя чата';

--------------------------- ПОЛЬЗОВАТЕЛЬ ЧАТА ------------------------------------
create table if not exists chat_user (
    id serial primary key,
    chat_id int not null,
    name text not null,
    foreign key (chat_id) references chat (id) on delete cascade on update cascade
);

comment on table chat_user is 'Таблица с пользователями определенного чата';
comment on column chat_user.id is 'Id пользователя';
comment on column chat_user.name is 'Имя пользователя';
comment on column chat_user.chat_id is 'Id чата, внешний ключ (связь с таблицей чатов)';

--------------------------- СООБЩЕНИЯ ------------------------------------

create table if not exists message (
    id serial primary key,
    sent_at timestamp not null default now(),
    chat_id int not null,
    user_id int,
    content text not null,
    foreign key (chat_id) references chat (id) on delete cascade on update cascade
);

comment on table message is 'Таблица с сообщениями в чате';
comment on column message.id is 'Id сообщения';
comment on column message.sent_at is 'Время отправки сообщения';
comment on column message.chat_id is 'Id чата, внешний ключ (связь с таблицей чатов)';
comment on column message.user_id is 'Id пользователя отправившего сообщение';
comment on column message.content is 'Текст сообщения';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table message;
drop table chat_user;
drop table chat;
-- +goose StatementEnd
