CREATE TABLE `order_infos`
(
    `id`            INT AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `created_at`    TIMESTAMP          NOT NULL COMMENT '创建时间',
    `updated_at`    TIMESTAMP          NOT NULL COMMENT '更新时间',
    `deleted_at`    TIMESTAMP COMMENT '删除时间',
    `order_info_id` VARCHAR(255)       NOT NULL COMMENT '订单信息ID',
    `user_id`       VARCHAR(255)       NOT NULL COMMENT '用户ID',
    `shop_id`       VARCHAR(255)       NOT NULL COMMENT '商店ID',
    `receiver_id`   VARCHAR(255)       NOT NULL COMMENT '支付方式ID',
    `total_amount`  DOUBLE             NOT NULL COMMENT '总金额',
    `pay_status`    INT                NOT NULL COMMENT '支付状态',
    `pay_amount`    DOUBLE             NOT NULL COMMENT '支付金额',
    `order_status`  INT                NOT NULL COMMENT '订单状态',
    `description`   TEXT               NOT NULL COMMENT '描述',
    PRIMARY KEY (`id`) COMMENT '主键',
    UNIQUE KEY `uq_order_info_id` (`order_info_id`) COMMENT '订单信息ID唯一键'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='订单信息表';

CREATE TABLE `order_items`
(
    `id`            INT AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `created_at`    TIMESTAMP          NOT NULL COMMENT '创建时间',
    `updated_at`    TIMESTAMP          NULL COMMENT '更新时间',
    `deleted_at`    TIMESTAMP COMMENT '删除时间',
    `order_item_id` VARCHAR(255)       NOT NULL COMMENT '订单项ID',
    `order_info_id` VARCHAR(255)       NOT NULL COMMENT '订单信息ID',
    `order_amount`  BIGINT             NOT NULL COMMENT '订单数量',
    `orgin_price`   DOUBLE             NOT NULL COMMENT '原价',
    `real_price`    DOUBLE             NOT NULL COMMENT '实际价格',
    PRIMARY KEY (`id`) COMMENT '主键'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='订单项表';

