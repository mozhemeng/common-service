CREATE TABLE IF NOT EXISTS `role` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) UNIQUE NOT NULL COMMENT '名称',
    `description` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '描述',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色';

INSERT IGNORE INTO `role` (id, name, description) VALUES (1, 'root', 'this is root');


CREATE TABLE IF NOT EXISTS `user` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(20) UNIQUE NOT NULL COMMENT '用户名',
    `password_hashed` VARCHAR(100) NOT NULL COMMENT '密码',
    `nickname` VARCHAR(20) NOT NULL COMMENT '昵称',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态',
    `role_id` INT UNSIGNED NOT NULL COMMENT '角色id',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';


CREATE TABLE IF NOT EXISTS `casbin_rule` (
    `p_type` varchar(100) NOT NULL COMMENT '规则类型',
    `v0` varchar(100) DEFAULT '' COMMENT '角色名',
    `v1` varchar(100) DEFAULT '' COMMENT 'api路径',
    `v2` varchar(100) DEFAULT '' COMMENT 'http方法',
    `v3` varchar(100) DEFAULT '',
    `v4` varchar(100) DEFAULT '',
    `v5` varchar(100) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限规则表';
