drop table if exists video_file;
create table video_file
(
    video_id  int auto_increment
        primary key,
    file_path varchar(4000) not null,
    file_name varchar(4000) not null,
    size      int           not null default 0
) comment '视频文件';

drop table if exists `cut_video`;
create table cut_video(
                          cut_id int(11) not null primary key auto_increment,
                          parent_id int(11) not null comment '从哪一个原始的视频剪切出来的',
                          start_time varchar(20) not null comment '之所使用字符串形式，是因为想要组成ffmpeg命令执行',
                          end_time varchar(20) not null ,
                          name varchar(100) not null default '' comment '一般使用父亲名字+开始结束时间吧',
                          status int(11) not null default 1 comment '视频的状态,1:任务建立，2:任务开始执行,3:任务执行成功，然后，视频存在，4:失败，那么就不存在了'
)comment '剪切视频片段';