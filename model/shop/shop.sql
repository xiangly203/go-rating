CREATE TABLE `shops`
(
    `id`          INT AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `created_at`  TIMESTAMP          NOT NULL COMMENT '创建时间',
    `updated_at`  TIMESTAMP          NOT NULL COMMENT '更新时间',
    `deleted_at`  TIMESTAMP COMMENT '删除时间',
    `shop_id`     VARCHAR(255)       NOT NULL COMMENT '商店ID',
    `shop_name`   VARCHAR(255)       NOT NULL COMMENT '商店名称',
    `description` TEXT               NOT NULL COMMENT '描述',
    `status`      INT                NOT NULL COMMENT '状态',
    PRIMARY KEY (`id`) COMMENT '主键',
    UNIQUE KEY `uq_shop_id` (`shop_id`) COMMENT '产品ID唯一键'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='商店表';
