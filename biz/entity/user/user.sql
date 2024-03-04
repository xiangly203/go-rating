CREATE TABLE `users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'User name',
    `email`      varchar(128) NOT NULL DEFAULT '' COMMENT 'User email',
    `phone_number` varchar(128) NOT NULL DEFAULT '' COMMENT 'User Phone number',
    `is_admin` bool NOT NULL DEFAULT false COMMENT 'User is admin',
    `is_delete` bool NOT NULL DEFAULT false COMMENT 'User is delete',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User information create time',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User information update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User information delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_email` (`email`) COMMENT 'User email index',
    UNIQUE KEY `idx_user_name` (`user_name`) COMMENT 'User name index',
    UNIQUE KEY `idx_phone_number` (`phone_number`) COMMENT 'User Phone number index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User information table';
