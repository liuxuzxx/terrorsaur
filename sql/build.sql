CREATE DATABASE IF NOT EXISTS ancient_article DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
use ancient_article;
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
    attribute_value int(11) not null comment '属性值'
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 comment '诗词的属性信息';

insert into author(name, dynasty, detail) values('广大劳动人民','中国','劳动人民总结生活中的物质，升华成了精神生活，俗语，歇后语和诗歌慢慢的演变而来');
insert into author(name, dynasty, detail) values('佚名','中国','无姓无名之人');
insert into author(name, dynasty, detail) values('屈原','中国','中国战国时期的有名辞人---楚国的最高法院院长');


drop function if exists shijing_attribute;

delimiter $$
drop procedure if exists `shijing_attribute`$$
create procedure `shijing_attribute`()
begin
    DECLARE i int default 1;
    WHILE (i < 306) DO
    insert into article_attribute(article_id, attribute_code, attribute_value) VALUES (i, 1, 1);
    insert into article_attribute(article_id, attribute_code, attribute_value) VALUES (i, 2, 1);
    SET i = i + 1;
    END WHILE;
end$$

delimiter ;
call shijing_attribute();


drop table if exists idioms;
create table idioms
(
    id      int(11) auto_increment not null comment '成语的ID' primary key ,
    term  varchar(100) not null comment '成语',
    pronunciation varchar(100) not null comment '发音',
    interpretation varchar(1000) not null comment '释义',
    source varchar(1000) comment '出处',
    example varchar(2000) comment '例子'
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8 comment '成语的信息';


drop table if exists `ancient_plate`;
create table ancient_plate(
                             plate_id int not null auto_increment primary key comment '自增的ID,板块ID',
                             plate_name varchar(20) not null default '' comment '板块名字',
                             detail varchar(2000) not null default '' comment '板块描述信息',
                             plate_order int not null default -1 comment '显示的顺序'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='古文板块类型';

insert into ancient_plate(plate_name, detail, plate_order) VALUES ('古诗词','古代的文章，包含诗词歌赋',1);
insert into ancient_plate(plate_name, detail, plate_order) VALUES ('短语','所有的短语类型',2);
