CREATE TABLE `products`
(
    `id`               INT AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `created_at`       TIMESTAMP          NOT NULL COMMENT '创建时间',
    `updated_at`       TIMESTAMP          NULL COMMENT '更新时间',
    `deleted_at`       TIMESTAMP COMMENT '删除时间',
    `product_name`     VARCHAR(255)       NOT NULL COMMENT '产品名称',
    `product_id`       VARCHAR(255)       NOT NULL COMMENT '产品ID',
    `shop_id`          VARCHAR(255)       NOT NULL COMMENT '商店ID',
    `stock_count`      BIGINT             NOT NULL COMMENT '库存数量',
    `sale_status`      BOOLEAN            NOT NULL COMMENT '销售状态 1上架,0下架',
    `orgin_price`      DOUBLE             NOT NULL COMMENT '原价',
    `discount_price`   DOUBLE             NOT NULL COMMENT '折扣价',
    `sale_stock_count` BIGINT             NOT NULL COMMENT '销售库存数量',
    `image_url`        TEXT               NOT NULL COMMENT '图片URL',
    `description`      TEXT               NOT NULL COMMENT '描述',
    PRIMARY KEY (`id`) COMMENT '主键',
    UNIQUE KEY `uq_product_id` (`product_id`) COMMENT '产品ID唯一键'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='SPU';
