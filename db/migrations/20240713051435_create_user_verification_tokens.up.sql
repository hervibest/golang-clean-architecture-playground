create table user_verification_tokens
(
    id         varchar(100) not null,
    user_id    varchar(100) not null,
    token      varchar(100) not null,
    created_at bigint       not null,
    updated_at bigint       not null,
    primary key (id),
    foreign key fk_user_verification_tokens_user_id (user_id) references users (id)
) engine = innodb;