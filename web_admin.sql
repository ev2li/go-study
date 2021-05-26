DROP TABLE IF EXISTS `tbl_app_info`;
create table if not exists `tbl_app_info` (
    `app_id`  INT AUTO_INCREMENT primary key,
    `app_name` VARCHAR(1024) not null,
    `app_type` VARCHAR(64) not null,
    `create_time` timestamp  default CURRENT_TIMESTAMP ,
    `develop_path` VARCHAR(256) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `tbl_app_ip`;
create table if not exists `tbl_app_ip` (
    `app_id`  INT,
    `ip` VARCHAR(64),
    key app_id_ip_index (app_id, ip)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `tbl_log_info`;
create table if not exists `tbl_log_info` (
    `log_id`  INT AUTO_INCREMENT primary key,
    `app_id`  VARCHAR(1024)  not null,
    `log_path` VARCHAR(64) not null,
    `create_time` timestamp  default CURRENT_TIMESTAMP,
    `status` tinyint default  1,
    `topic` VARCHAR(255) not null
)ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

