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

