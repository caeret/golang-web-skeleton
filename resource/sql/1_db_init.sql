-- +migrate Up
CREATE TABLE `user` (
                         `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                         `name` VARCHAR(32) NOT NULL,
                         `password_hash` VARCHAR(512) NOT NULL,
                         `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down
DROP TABLE `user`;