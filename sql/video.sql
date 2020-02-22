drop table if exists video_file;
create table video_file
(
    video_id  int auto_increment
        primary key,
    file_path varchar(4000) not null,
    file_name varchar(4000) not null,
    size      int           not null default 0
) comment '视频文件';