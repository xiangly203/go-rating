CREATE TABLE voucher
(
    `id`           INT AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `created_at`   TIMESTAMP          NOT NULL COMMENT '创建时间',
    `updated_at`   TIMESTAMP          NOT NULL COMMENT '更新时间',
    `deleted_at`   TIMESTAMP COMMENT '删除时间',
    `voucher_id`   VARCHAR(255)       NOT NULL COMMENT '凭证ID',
    `shop_id`      VARCHAR(255)       NOT NULL COMMENT '商店ID',
    `title`        VARCHAR(255)       NOT NULL COMMENT '标题',
    `sub_title`    VARCHAR(255)       NOT NULL COMMENT '副标题',
    `rules`        VARCHAR(255)       NOT NULL COMMENT '规则',
    `pay_value`    DECIMAL(10, 2)     NOT NULL COMMENT '支付值',
    `actual_value` DECIMAL(10, 2)     NOT NULL COMMENT '实际值',
    `type`         VARCHAR(255)       NOT NULL COMMENT '类型',
    `status`       VARCHAR(255)       NOT NULL COMMENT '状态',
    PRIMARY KEY (`id`) COMMENT '主键',
    UNIQUE KEY `uq_voucher_id` (`voucher_id`) COMMENT '产品ID唯一键'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='优惠券表';


CREATE TABLE seckill_voucher
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    `created_at`   TIMESTAMP          NOT NULL COMMENT '创建时间',
    `updated_at`   TIMESTAMP          NOT NULL COMMENT '更新时间',
    `deleted_at`   TIMESTAMP COMMENT '删除时间',
    voucher_id VARCHAR(255) NOT NULL COMMENT '凭证ID',
    stock      INT          NOT NULL COMMENT '库存',
    begin_time DATETIME     NOT NULL COMMENT '开始时间',
    end_time   DATETIME     NOT NULL COMMENT '结束时间'
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4
 COLLATE = utf8mb4_general_ci COMMENT ='秒杀优惠券表';

CREATE TABLE `voucher_order`
(
    `id`            INT AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `created_at`    TIMESTAMP          NOT NULL COMMENT '创建时间',
    `updated_at`    TIMESTAMP          NOT NULL COMMENT '更新时间',
    `deleted_at`    TIMESTAMP COMMENT '删除时间',
    `voucher_order_id` VARCHAR(255)       NOT NULL COMMENT '订单信息ID',
    `user_id`       VARCHAR(255)       NOT NULL COMMENT '用户ID',
    `receiver_id`  DOUBLE             NOT NULL COMMENT '支付方式',
    `pay_status`    INT                NOT NULL COMMENT '支付状态 0未支付 1已支付',
    `pay_amount`    DOUBLE             NOT NULL COMMENT '支付金额',
    `status`  INT                NOT NULL COMMENT '优惠券状态 0无效，1有效，2已使用，3已退款',
    `use_time`    TIMESTAMP          NOT NULL COMMENT '核销时间',
    `refund_time`    TIMESTAMP          NOT NULL COMMENT '退款时间',
    PRIMARY KEY (`id`) COMMENT '主键',
    UNIQUE KEY `uq_order_info_id` (`voucher_order_id`) COMMENT '订单信息ID唯一键'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='优惠券订单信息表';
