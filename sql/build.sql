drop table if exists `article_type`;
create table article_type(
    type_id int not null auto_increment primary key comment '自增的ID',
    type_name varchar(20) not null default '' comment '类型名字',
    detail varchar(2000) not null default '' comment '描述信息',
    type_order int not null default -1 comment '显示的顺序'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='古文类型';

insert into article_type(type_name, detail, type_order)
VALUES ('诗经','汉代、两晋及南北朝的乐府诗歌',1),
       ('楚辞','战国时期楚国的特有的一种骈文',2),
       ('汉赋','大汉之风的骈文',3),
       ('唐诗','鼎盛的唐朝定义了诗歌的发展之最',4),
       ('宋词','经济第一的宋朝别有风味',5),
       ('元曲','格调的有一次提升，电视剧的前身',6),
       ('明清','封建社会进入了风烛残年的时代',7);


drop table if exists author;
create table author(
    author_id int not null auto_increment primary key comment '作家的ID',
    name char(10) not null comment '作家名字',
    dynasty char(10) not null comment '朝代',
    detail varchar(4000) not null comment '作家的详细信息'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '作家的信息';

drop table if exists article;
create table article(
    article_id int not null auto_increment primary key comment '文章自增的ID',
    title varchar(200) not null comment '文章题目',
    content text(20000) not null comment '文章内容',
    comment text(20000) not null comment '注解',
    create_by varchar(100) not null comment '创建者',
    create_time datetime not null default now() comment '创建时间',
    update_by varchar(100) not null comment '修改者',
    update_time datetime not null default now() comment '修改时间'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '古文信息' ;

drop table if exists article_attribute;
create table article_attribute(
    article_id int(11) not null comment '文章的ID',
    attribute_code int(11)  not null comment '属性code 1:类型　2:作者　3:朝代　4:形式',
    attribute_value varchar(100) not null comment '属性值'
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 comment '诗词的属性信息';
