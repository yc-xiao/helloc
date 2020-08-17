/* 创建库，初始化账号 */
create database H if not exists;
create user 'hello'@'%' identity by '123456' if not exists;
grant all on H.* to 'hello'@'%';
