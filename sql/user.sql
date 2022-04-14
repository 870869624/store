create table if not exists USERS(
    ID int(11) unsigned not null auto_increment primary key,
    NICKNAME varchar(12) COMMENT '昵称',
    USERNAME varchar(64) not null COMMENT '用户名',
    PHONE varchar(16) not null COMMENT '电话号码',
    PASSWORD varchar(128) not null COMMENT '用户密码',
    GENDER TINYINT UNSIGNED not null DEFAULT 0 COMMENT '性别',
    AVATAR varchar(128) COMMENT '头像',
    STATUS TINYINT(1) NOT NULL DEFAULT 0 COMMENT '0禁用1可用',
    CREATED_AT TIMESTAMP DEFAULT NOW() COMMENT '注册时间'
)