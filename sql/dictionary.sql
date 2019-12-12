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

drop table if exists dictionary;
create table dictionary(
    id int not null auto_increment primary key comment '主键自增id',
    chinese_character varchar(20) not null comment '汉字',
    explanation varchar(1000) not null comment '汉字解释',
    source varchar(2000) not null comment '来源语句',
    dictionary_type int not null comment '词典类型，来源于哪一种词典'
);