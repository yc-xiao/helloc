create database if not exists H default character set utf8 collate utf8_general_ci;

/*
create user 'hello'@'%' identified by '123456';
grant all on H.* to 'hello'@'%';
*/

-- 进入数据库，创建表　user, video, comment
use `H`;

drop table if exists comment;
drop table if exists video;
drop table if exists user;

-- 用户(id, nickname, account, password, email, phone, isAdmin, photoFile(头像), createdTime)(微信略)
create table if not exists `user` (
    `id` int primary key auto_increment,
    `nickname` varchar(255) not null,
    `account` varchar(255) not null unique,
    `password` varchar(255) not null,
    `email` varchar(255) default "",
    `phone` varchar(255) default "",
    `photoFile` varchar(255) default "",
    `isAdmin` tinyint(1) default 0,
    `createdTime` varchar(255) default ""
);

-- 创建用户
insert into `user` (nickname, account, password, isAdmin) values ("小明", "xiaoming", "123456", 0),
                                                                 ("小红", "xiaohong", "123456", 0),
                                                                 ("小黄", "xiaohuang", "123456", 0),
                                                                 ("管理员大大", "admin", "123456", 1);

-- 视频(id, userId, title, desc, category, label, rm, duration(时长), videoFile, createdTime, pass)
create table if not exists `video` (
    `id` int primary key auto_increment,
    `userId` int default 0,
    `title` varchar(255) not null,
    `desc` text default "",
    `category` varchar(255) default "",
    `label` varchar(255) default "",
    `rm` varchar(255) default "",
    `duration` varchar(255) default "",
    `videoFile` varchar(255) default "",
    `createdTime` varchar(255) default "",
    `pass` tinyint(1) default 0,
    foreign key(userId) references user(id)
);

-- 创建视频
insert into `video` (userId, title, pass) values (1, "黄飞鸿1", 1),
                                                (1, "黄飞鸿2", 1),
                                                (1, "黄飞鸿3", 0),
                                                (4, "黄飞鸿4", 1);

-- 评论(id, objId, model, context, createdTime)
create table if not exists `comment` (
    `id` int primary key auto_increment,
    `objId` int not null,
    `model` varchar(255) not null,
    `context` text default "",
    `createdTime` varchar(255) default "",
#     `originId` int not null,
#     `originModel` varchar(255) not null,
    `userId` int default 0,
    foreign key(userId) references user(id)
);


-- 创建评论
insert into `comment` (userId, context, model, objId) values (1, "用户1->视频1->1星", "video", 1),
                                                             (2, "用户2->视频1->5星", "video", 1),
                                                             (2, "用户2->评论1->5星", "comment", 1),
                                                             (1, "用户1回复->评论3->3星", "comment", 3),
                                                             (3, "用户3->评论1->5星", "comment", 1);

