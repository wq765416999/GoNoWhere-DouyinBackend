-- auto-generated definition
create table logins
(
    id         bigint unsigned auto_increment comment '自增主键'
        primary key,
    create_at  timestamp default CURRENT_TIMESTAMP not null,
    update_at  timestamp default CURRENT_TIMESTAMP not null,
    deleted_at datetime(3)                         null,
    user_id    bigint                              not null,
    name       varchar(34)                         not null,
    pass_word  varchar(65)                         not null,
    constraint idx_user_id
        unique (user_id),
    constraint idx_user_name
        unique (name),
    constraint name
        unique (name),
    constraint user_id
        unique (user_id),
        primary key (id)
)
    collate = utf8mb4_general_ci;

