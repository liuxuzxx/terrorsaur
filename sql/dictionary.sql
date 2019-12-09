drop table if exists dictionary_type;
create table dictionary_type
(
    id              int auto_increment
        primary key,
    dictionary_code varchar(100)  not null,
    dictionary_name varchar(100) not null
)
    comment '字典类型';

insert into dictionary_type(dictionary_code, dictionary_name) VALUES ('DICT_ERYA','尔雅');
insert into dictionary_type(dictionary_code, dictionary_name) VALUES ('DICT_SWJZ','说文解字');
insert into dictionary_type(dictionary_code, dictionary_name) VALUES ('DICT_KXZD','康熙字典');