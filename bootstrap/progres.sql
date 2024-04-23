DROP TABLE IF EXISTS tb_blog;

CREATE TABLE tb_blog
(
    id          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    shop_id     BIGINT        NOT NULL,
    user_id     BIGINT        NOT NULL,
    title       varchar(255)  NOT NULL,
    images      varchar(2048) NOT NULL,
    content     varchar(2048) NOT NULL,
    liked       INT                    DEFAULT 0,
    comments    INT                    DEFAULT 0,
    create_time TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN tb_blog.id IS '主键';
COMMENT ON COLUMN tb_blog.shop_id IS '商户id';
COMMENT ON COLUMN tb_blog.user_id IS '用户id';
COMMENT ON COLUMN tb_blog.title IS '标题';
COMMENT ON COLUMN tb_blog.images IS '探店的照片，最多9张，多张以","隔开';
COMMENT ON COLUMN tb_blog.content IS '探店的文字描述';
COMMENT ON COLUMN tb_blog.liked IS '点赞数量';
COMMENT ON COLUMN tb_blog.comments IS '评论数量';
COMMENT ON COLUMN tb_blog.create_time IS '创建时间';
COMMENT ON COLUMN tb_blog.update_time IS '更新时间';

-- ----------------------------
-- Records of tb_blog
-- ----------------------------
INSERT INTO tb_blog(shop_id, user_id, title, images, content, liked, comments, create_time, update_time)
VALUES (4, 2, '无尽浪漫的夜晚丨在万花丛中摇晃着红酒杯🍷品战斧牛排🥩',
        '/imgs/blogs/7/14/4771fefb-1a87-4252-816c-9f7ec41ffa4a.jpg,/imgs/blogs/4/10/2f07e3c9-ddce-482d-9ea7-c21450f8d7cd.jpg,/imgs/blogs/2/6/b0756279-65da-4f2d-b62a-33f74b06454a.jpg,/imgs/blogs/10/7/7e97f47d-eb49-4dc9-a583-95faa7aed287.jpg,/imgs/blogs/1/2/4a7b496b-2a08-4af7-aa95-df2c3bd0ef97.jpg,/imgs/blogs/14/3/52b290eb-8b5d-403b-8373-ba0bb856d18e.jpg',
        '生活就是一半烟火·一半诗意<br/>手执烟火谋生活·心怀诗意以谋爱·<br/>当然<br/>\r\n男朋友给不了的浪漫要学会自己给🍒<br/>\n无法重来的一生·尽量快乐.<br/><br/>🏰「小筑里·神秘浪漫花园餐厅」🏰<br/><br/>\n💯这是一家最最最美花园的西餐厅·到处都是花餐桌上是花前台是花  美好无处不在\n品一口葡萄酒，维亚红酒马瑟兰·微醺上头工作的疲惫消失无际·生如此多娇🍃<br/><br/>📍地址:延安路200号(家乐福面)<br/><br/>🚌交通:地铁①号线定安路B口出右转过下通道右转就到啦～<br/><br/>--------------🥣菜品详情🥣---------------<br/><br/>「战斧牛排]<br/>\n超大一块战斧牛排经过火焰的炙烤发出阵阵香，外焦里嫩让人垂涎欲滴，切开牛排的那一刻，牛排的汁水顺势流了出来，分熟的牛排肉质软，简直细嫩到犯规，一刻都等不了要放入嘴里咀嚼～<br/><br/>「奶油培根意面」<br/>太太太好吃了💯<br/>我真的无法形容它的美妙，意面混合奶油香菇的香味真的太太太香了，我真的舔盘了，一丁点美味都不想浪费‼️<br/><br/><br/>「香菜汁烤鲈鱼」<br/>这个酱是辣的 真的绝好吃‼️<br/>鲈鱼本身就很嫩没什么刺，烤过之后外皮酥酥的，鱼肉蘸上酱料根本停不下来啊啊啊啊<br/>能吃辣椒的小伙伴一定要尝尝<br/><br/>非常可 好吃子🍽\n<br/>--------------🍃个人感受🍃---------------<br/><br/>【👩🏻‍🍳服务】<br/>小姐姐特别耐心的给我们介绍彩票 <br/>推荐特色菜品，拍照需要帮忙也是尽心尽力配合，太爱他们了<br/><br/>【🍃环境】<br/>比较有格调的西餐厅 整个餐厅的布局可称得上的万花丛生 有种在人间仙境的感觉🌸<br/>集美食美酒与鲜花为一体的风格店铺 令人向往<br/>烟火皆是生活 人间皆是浪漫<br/>',
        1, 104, '2021-12-28 19:50:01', '2022-03-10 14:26:34');
INSERT INTO tb_blog(shop_id, user_id, title, images, content, liked, comments, create_time, update_time)
VALUES (1, 2, '人均30💰杭州这家港式茶餐厅我疯狂打call‼️',
        '/imgs/blogs/4/7/863cc302-d150-420d-a596-b16e9232a1a6.jpg,/imgs/blogs/11/12/8b37d208-9414-4e78-b065-9199647bb3e3.jpg,/imgs/blogs/4/1/fa74a6d6-3026-4cb7-b0b6-35abb1e52d11.jpg,/imgs/blogs/9/12/ac2ce2fb-0605-4f14-82cc-c962b8c86688.jpg,/imgs/blogs/4/0/26a7cd7e-6320-432c-a0b4-1b7418f45ec7.jpg,/imgs/blogs/15/9/cea51d9b-ac15-49f6-b9f1-9cf81e9b9c85.jpg',
        '又吃到一家好吃的茶餐厅🍴环境是怀旧tvb港风📺边吃边拍照片📷几十种菜品均价都在20+💰可以是很平价了！<br>·<br>店名：九记冰厅(远洋店)<br>地址：杭州市丽水路远洋乐堤港负一楼（溜冰场旁边）<br>·<br>✔️黯然销魂饭（38💰）<br>这碗饭我吹爆！米饭上盖满了甜甜的叉烧 还有两颗溏心蛋🍳每一粒米饭都裹着浓郁的酱汁 光盘了<br>·<br>✔️铜锣湾漏奶华（28💰）<br>黄油吐司烤的脆脆的 上面洒满了可可粉🍫一刀切开 奶盖流心像瀑布一样流出来  满足<br>·<br>✔️神仙一口西多士士（16💰）<br>简简单单却超级好吃！西多士烤的很脆 黄油味浓郁 面包体超级柔软 上面淋了炼乳<br>·<br>✔️怀旧五柳炸蛋饭（28💰）<br>四个鸡蛋炸成蓬松的炸蛋！也太好吃了吧！还有大块鸡排 上淋了酸甜的酱汁 太合我胃口了！！<br>·<br>✔️烧味双拼例牌（66💰）<br>选了烧鹅➕叉烧 他家烧腊品质真的惊艳到我！据说是每日广州发货 到店现烧现卖的黑棕鹅 每口都是正宗的味道！肉质很嫩 皮超级超级酥脆！一口爆油！叉烧肉也一点都不柴 甜甜的很入味 搭配梅子酱很解腻 ！<br>·<br>✔️红烧脆皮乳鸽（18.8💰）<br>乳鸽很大只 这个价格也太划算了吧， 肉质很有嚼劲 脆皮很酥 越吃越香～<br>·<br>✔️大满足小吃拼盘（25💰）<br>翅尖➕咖喱鱼蛋➕蝴蝶虾➕盐酥鸡<br>zui喜欢里面的咖喱鱼！咖喱酱香甜浓郁！鱼蛋很q弹～<br>·<br>✔️港式熊仔丝袜奶茶（19💰）<br>小熊🐻造型的奶茶冰也太可爱了！颜值担当 很地道的丝袜奶茶 茶味特别浓郁～<br>·',
        1, 0, '2021-12-28 20:57:49', '2022-03-10 09:21:39');
INSERT INTO tb_blog(shop_id, user_id, title, images, content, liked, comments, create_time, update_time)
VALUES (10, 1, '杭州周末好去处｜💰50就可以骑马啦🐎', '/imgs/blogs/blog1.jpg', '杭州周末好去处｜💰50就可以骑马啦🐎', 1, 0,
        '2022-01-11 16:05:47', '2022-03-10 09:21:41');
INSERT INTO tb_blog(shop_id, user_id, title, images, content, liked, comments, create_time, update_time)
VALUES (10, 1, '杭州周末好去处｜💰50就可以骑马啦🐎', '/imgs/blogs/blog1.jpg', '杭州周末好去处｜💰50就可以骑马啦🐎', 1, 0,
        '2022-01-11 16:05:47', '2022-03-10 09:21:42');

-- ----------------------------
-- Table structure for tb_blog_comments
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_blog_comments;

-- Create the table
CREATE TABLE tb_blog_comments
(
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT       NOT NULL,
    blog_id     BIGINT       NOT NULL,
    parent_id   BIGINT       NOT NULL,
    answer_id   BIGINT       NOT NULL,
    content     VARCHAR(255) NOT NULL,
    liked       INT                   DEFAULT NULL,
    status      SMALLINT              DEFAULT NULL,
    create_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to each column
COMMENT ON COLUMN tb_blog_comments.id IS '主键';
COMMENT ON COLUMN tb_blog_comments.user_id IS '用户id';
COMMENT ON COLUMN tb_blog_comments.blog_id IS '探店id';
COMMENT ON COLUMN tb_blog_comments.parent_id IS '关联的1级评论id，如果是一级评论，则值为0';
COMMENT ON COLUMN tb_blog_comments.answer_id IS '回复的评论id';
COMMENT ON COLUMN tb_blog_comments.content IS '回复的内容';
COMMENT ON COLUMN tb_blog_comments.liked IS '点赞数';
COMMENT ON COLUMN tb_blog_comments.status IS '状态，0：正常，1：被举报，2：禁止查看';
COMMENT ON COLUMN tb_blog_comments.create_time IS '创建时间';
COMMENT ON COLUMN tb_blog_comments.update_time IS '更新时间';
-- ----------------------------
-- Records of tb_blog_comments
-- ----------------------------

-- ----------------------------
-- Table structure for tb_follow
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_follow;

-- Create the table
CREATE TABLE tb_follow
(
    id             BIGSERIAL PRIMARY KEY,
    user_id        BIGINT    NOT NULL,
    follow_user_id BIGINT    NOT NULL,
    create_time    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to each column
COMMENT ON COLUMN tb_follow.id IS '主键';
COMMENT ON COLUMN tb_follow.user_id IS '用户id';
COMMENT ON COLUMN tb_follow.follow_user_id IS '关联的用户id';
COMMENT ON COLUMN tb_follow.create_time IS '创建时间';

-- ----------------------------
-- Records of tb_follow
-- ----------------------------

-- ----------------------------
-- Table structure for tb_seckill_voucher
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_seckill_voucher;

-- Create the table
CREATE TABLE tb_seckill_voucher
(
    voucher_id  BIGINT    NOT NULL,
    stock       INT       NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    begin_time  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_time    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (voucher_id)
);

-- Add comments to each column and the table
COMMENT ON COLUMN tb_seckill_voucher.voucher_id IS '关联的优惠券的id';
COMMENT ON COLUMN tb_seckill_voucher.stock IS '库存';
COMMENT ON COLUMN tb_seckill_voucher.create_time IS '创建时间';
COMMENT ON COLUMN tb_seckill_voucher.begin_time IS '生效时间';
COMMENT ON COLUMN tb_seckill_voucher.end_time IS '失效时间';
COMMENT ON COLUMN tb_seckill_voucher.update_time IS '更新时间';
COMMENT ON TABLE tb_seckill_voucher IS '秒杀优惠券表，与优惠券是一对一关系';

-- ----------------------------
-- Records of tb_seckill_voucher
-- ----------------------------

-- ----------------------------
-- Table structure for tb_shop
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_shop;

-- Create the table
CREATE TABLE tb_shop
(
    id          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(128)     NOT NULL,
    type_id     BIGINT           NOT NULL,
    images      VARCHAR(1024)    NOT NULL,
    area        VARCHAR(128)                DEFAULT NULL,
    address     VARCHAR(255)     NOT NULL,
    x           DOUBLE PRECISION NOT NULL,
    y           DOUBLE PRECISION NOT NULL,
    avg_price   BIGINT                      DEFAULT NULL,
    sold        INT                         DEFAULT NULL,
    comments    INT                         DEFAULT NULL,
    score       INT              NOT NULL,
    open_hours  VARCHAR(32)                 DEFAULT NULL,
    create_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to each column
COMMENT ON COLUMN tb_shop.id IS '主键';
COMMENT ON COLUMN tb_shop.name IS '商铺名称';
COMMENT ON COLUMN tb_shop.type_id IS '商铺类型的id';
COMMENT ON COLUMN tb_shop.images IS '商铺图片，多个图片以\n隔开';
COMMENT ON COLUMN tb_shop.area IS '商圈，例如陆家嘴';
COMMENT ON COLUMN tb_shop.address IS '地址';
COMMENT ON COLUMN tb_shop.x IS '经度';
COMMENT ON COLUMN tb_shop.y IS '维度';
COMMENT ON COLUMN tb_shop.avg_price IS '均价，取整数';
COMMENT ON COLUMN tb_shop.sold IS '销量';
COMMENT ON COLUMN tb_shop.comments IS '评论数量';
COMMENT ON COLUMN tb_shop.score IS '评分，1~5分，乘10保存，避免小数';
COMMENT ON COLUMN tb_shop.open_hours IS '营业时间，例如 10:00-22:00';
COMMENT ON COLUMN tb_shop.create_time IS '创建时间';
COMMENT ON COLUMN tb_shop.update_time IS '更新时间';

CREATE INDEX foreign_key_type ON tb_shop (type_id);
-- ----------------------------
-- Records of tb_shop
-- ----------------------------
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('103茶餐厅', 1,
        'https://qcloud.dpfile.com/pc/jiclIsCKmOI2arxKN1Uf0Hx3PucIJH8q0QSz-Z8llzcN56-_QiKuOvyio1OOxsRtFoXqu0G3iT2T27qat3WhLVEuLYk00OmSS1IdNpm8K8sG4JN9RIm2mTKcbLtc2o2vfCF2ubeXzk49OsGrXt_KYDCngOyCwZK-s3fqawWswzk.jpg,https://qcloud.dpfile.com/pc/IOf6VX3qaBgFXFVgp75w-KKJmWZjFc8GXDU8g9bQC6YGCpAmG00QbfT4vCCBj7njuzFvxlbkWx5uwqY2qcjixFEuLYk00OmSS1IdNpm8K8sG4JN9RIm2mTKcbLtc2o2vmIU_8ZGOT1OjpJmLxG6urQ.jpg',
        '大关', '金华路锦昌文华苑29号', 120.149192, 30.316078, 80, 0000004215, 0000003035, 37, '10:00-22:00',
        '2021-12-22 18:10:39', '2022-01-13 17:32:19');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('蔡馬洪涛烤肉·老北京铜锅涮羊肉', 1,
        'https://p0.meituan.net/bbia/c1870d570e73accbc9fee90b48faca41195272.jpg,http://p0.meituan.net/mogu/397e40c28fc87715b3d5435710a9f88d706914.jpg,https://qcloud.dpfile.com/pc/MZTdRDqCZdbPDUO0Hk6lZENRKzpKRF7kavrkEI99OxqBZTzPfIxa5E33gBfGouhFuzFvxlbkWx5uwqY2qcjixFEuLYk00OmSS1IdNpm8K8sG4JN9RIm2mTKcbLtc2o2vmIU_8ZGOT1OjpJmLxG6urQ.jpg',
        '拱宸桥/上塘', '上塘路1035号（中国工商银行旁）', 120.151505, 30.333422, 85, 0000002160, 0000001460, 46,
        '11:30-03:00', '2021-12-22 19:00:13', '2022-01-11 16:12:26');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('新白鹿餐厅(运河上街店)', 1,
        'https://p0.meituan.net/biztone/694233_1619500156517.jpeg,https://img.meituan.net/msmerchant/876ca8983f7395556eda9ceb064e6bc51840883.png,https://img.meituan.net/msmerchant/86a76ed53c28eff709a36099aefe28b51554088.png',
        '运河上街', '台州路2号运河上街购物中心F5', 120.151954, 30.32497, 61, 0000012035, 0000008045, 47, '10:30-21:00',
        '2021-12-22 19:10:05', '2022-01-11 16:12:42');
INSERT INTO tb_shop(name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                    create_time, update_time)
VALUES ('Mamala(杭州远洋乐堤港店)', 1,
        'https://img.meituan.net/msmerchant/232f8fdf09050838bd33fb24e79f30f9606056.jpg,https://qcloud.dpfile.com/pc/rDe48Xe15nQOHCcEEkmKUp5wEKWbimt-HDeqYRWsYJseXNncvMiXbuED7x1tXqN4uzFvxlbkWx5uwqY2qcjixFEuLYk00OmSS1IdNpm8K8sG4JN9RIm2mTKcbLtc2o2vmIU_8ZGOT1OjpJmLxG6urQ.jpg',
        '拱宸桥/上塘', '丽水路66号远洋乐堤港商城2期1层B115号', 120.146659, 30.312742, 290, 0000013519, 0000009529, 49,
        '11:00-22:00', '2021-12-22 19:17:15', '2022-01-11 16:12:51');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('海底捞火锅(水晶城购物中心店）', 1,
        'https://img.meituan.net/msmerchant/054b5de0ba0b50c18a620cc37482129a45739.jpg,https://img.meituan.net/msmerchant/59b7eff9b60908d52bd4aea9ff356e6d145920.jpg,https://qcloud.dpfile.com/pc/Qe2PTEuvtJ5skpUXKKoW9OQ20qc7nIpHYEqJGBStJx0mpoyeBPQOJE4vOdYZwm9AuzFvxlbkWx5uwqY2qcjixFEuLYk00OmSS1IdNpm8K8sG4JN9RIm2mTKcbLtc2o2vmIU_8ZGOT1OjpJmLxG6urQ.jpg',
        '大关', '上塘路458号水晶城购物中心F6', 120.15778, 30.310633, 104, 0000004125, 0000002764, 49, '10:00-07:00',
        '2021-12-22 19:20:58', '2022-01-11 16:13:01');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('幸福里老北京涮锅（丝联店）', 1,
        'https://img.meituan.net/msmerchant/e71a2d0d693b3033c15522c43e03f09198239.jpg,https://img.meituan.net/msmerchant/9f8a966d60ffba00daf35458522273ca658239.jpg,https://img.meituan.net/msmerchant/ef9ca5ef6c05d381946fe4a9aa7d9808554502.jpg',
        '拱宸桥/上塘', '金华南路189号丝联166号', 120.148603, 30.318618, 130, 0000009531, 0000007324, 46,
        '11:00-13:50,17:00-20:50', '2021-12-22 19:24:53', '2022-01-11 16:13:09');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('炉鱼(拱墅万达广场店)', 1,
        'https://img.meituan.net/msmerchant/909434939a49b36f340523232924402166854.jpg,https://img.meituan.net/msmerchant/32fd2425f12e27db0160e837461c10303700032.jpg,https://img.meituan.net/msmerchant/f7022258ccb8dabef62a0514d3129562871160.jpg',
        '北部新城', '杭行路666号万达商业中心4幢2单元409室(铺位号4005)', 120.124691, 30.336819, 85, 0000002631,
        0000001320, 47, '00:00-24:00', '2021-12-22 19:40:52', '2022-01-11 16:13:19');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('浅草屋寿司（运河上街店）', 1,
        'https://img.meituan.net/msmerchant/cf3dff697bf7f6e11f4b79c4e7d989e4591290.jpg,https://img.meituan.net/msmerchant/0b463f545355c8d8f021eb2987dcd0c8567811.jpg,https://img.meituan.net/msmerchant/c3c2516939efaf36c4ccc64b0e629fad587907.jpg',
        '运河上街', '拱墅区金华路80号运河上街B1', 120.150526, 30.325231, 88, 0000002406, 0000001206, 46, ' 11:00-21:30',
        '2021-12-22 19:51:06', '2022-01-11 16:13:25');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('羊老三羊蝎子牛仔排北派炭火锅(运河上街店)', 1,
        'https://p0.meituan.net/biztone/163160492_1624251899456.jpeg,https://img.meituan.net/msmerchant/e478eb16f7e31a7f8b29b5e3bab6de205500837.jpg,https://img.meituan.net/msmerchant/6173eb1d18b9d70ace7fdb3f2dd939662884857.jpg',
        '运河上街', '台州路2号运河上街购物中心F5', 120.150598, 30.325251, 101, 0000002763, 0000001363, 44,
        '11:00-21:30', '2021-12-22 19:53:59', '2022-01-11 16:13:34');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('开乐迪KTV（运河上街店）', 2,
        'https://p0.meituan.net/joymerchant/a575fd4adb0b9099c5c410058148b307-674435191.jpg,https://p0.meituan.net/merchantpic/68f11bf850e25e437c5f67decfd694ab2541634.jpg,https://p0.meituan.net/dpdeal/cb3a12225860ba2875e4ea26c6d14fcc197016.jpg',
        '运河上街', '台州路2号运河上街购物中心F4', 120.149093, 30.324666, 67, 0000026891, 0000000902, 37, '00:00-24:00',
        '2021-12-22 20:25:16', '2021-12-22 20:25:16');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('INLOVE KTV(水晶城店)', 2,
        'https://p0.meituan.net/dpmerchantpic/53e74b200211d68988a4f02ae9912c6c1076826.jpg,https://qcloud.dpfile.com/pc/4iWtIvzLzwM2MGgyPu1PCDb4SWEaKqUeHm--YAt1EwR5tn8kypBcqNwHnjg96EvT_Gd2X_f-v9T8Yj4uLt25Gg.jpg,https://qcloud.dpfile.com/pc/WZsJWRI447x1VG2x48Ujgu7vwqksi_9WitdKI4j3jvIgX4MZOpGNaFtM93oSSizbGybIjx5eX6WNgCPvcASYAw.jpg',
        '水晶城', '上塘路458号水晶城购物中心6层', 120.15853, 30.310002, 75, 0000035977, 0000005684, 47, '11:30-06:00',
        '2021-12-22 20:29:02', '2021-12-22 20:39:00');
INSERT INTO tb_shop(name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                    create_time, update_time)
VALUES ('魅(杭州远洋乐堤港店)', 2,
        'https://p0.meituan.net/dpmerchantpic/63833f6ba0393e2e8722420ef33f3d40466664.jpg,https://p0.meituan.net/dpmerchantpic/ae3c94cc92c529c4b1d7f68cebed33fa105810.png,',
        '远洋乐堤港', '丽水路58号远洋乐堤港F4', 120.14983, 30.31211, 88, 0000006444, 0000000235, 46, '10:00-02:00',
        '2021-12-22 20:34:34', '2021-12-22 20:34:34');
INSERT INTO tb_shop(name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                    create_time, update_time)
VALUES ('讴K拉量贩KTV(北城天地店)', 2,
        'https://p1.meituan.net/merchantpic/598c83a8c0d06fe79ca01056e214d345875600.jpg,https://qcloud.dpfile.com/pc/HhvI0YyocYHRfGwJWqPQr34hRGRl4cWdvlNwn3dqghvi4WXlM2FY1te0-7pE3Wb9_Gd2X_f-v9T8Yj4uLt25Gg.jpg,https://qcloud.dpfile.com/pc/F5ZVzZaXFE27kvQzPnaL4V8O9QCpVw2nkzGrxZE8BqXgkfyTpNExfNG5CEPQX4pjGybIjx5eX6WNgCPvcASYAw.jpg',
        'D32天阳购物中心', '湖州街567号北城天地5层', 120.130453, 30.327655, 58, 0000018997, 0000001857, 41,
        '12:00-02:00', '2021-12-22 20:38:54', '2021-12-22 20:40:04');
INSERT INTO tb_shop (name, type_id, images, area, address, x, y, avg_price, sold, comments, score, open_hours,
                     create_time, update_time)
VALUES ('星聚会KTV(拱墅区万达店)', 2,
        'https://p0.meituan.net/dpmerchantpic/f4cd6d8d4eb1959c3ea826aa05a552c01840451.jpg,https://p0.meituan.net/dpmerchantpic/2efc07aed856a8ab0fc75c86f4b9b0061655777.jpg,https://qcloud.dpfile.com/pc/zWfzzIorCohKT0bFwsfAlHuayWjI6DBEMPHHncmz36EEMU9f48PuD9VxLLDAjdoU_Gd2X_f-v9T8Yj4uLt25Gg.jpg',
        '北部新城', '杭行路666号万达广场C座1-2F', 120.128958, 30.337252, 60, 0000017771, 0000000685, 47, '10:00-22:00',
        '2021-12-22 20:48:54', '2021-12-22 20:48:54');

-- ----------------------------
-- Table structure for tb_shop_type
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_shop_type;

-- Create the table
CREATE TABLE tb_shop_type
(
    id          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(32)                 DEFAULT NULL,
    icon        VARCHAR(255)                DEFAULT NULL,
    sort        INT                         DEFAULT NULL,
    create_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to each column
COMMENT ON COLUMN tb_shop_type.id IS '主键';
COMMENT ON COLUMN tb_shop_type.name IS '类型名称';
COMMENT ON COLUMN tb_shop_type.icon IS '图标';
COMMENT ON COLUMN tb_shop_type.sort IS '顺序';
COMMENT ON COLUMN tb_shop_type.create_time IS '创建时间';
COMMENT ON COLUMN tb_shop_type.update_time IS '更新时间';

-- ----------------------------
-- Records of tb_shop_type
-- ----------------------------
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('美食', '/types/ms.png', 1, '2021-12-22 20:17:47', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type (name, icon, sort, create_time, update_time)
VALUES ('KTV', '/types/KTV.png', 2, '2021-12-22 20:18:27', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('丽人·美发', '/types/lrmf.png', 3, '2021-12-22 20:18:48', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('健身运动', '/types/jsyd.png', 10, '2021-12-22 20:19:04', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('按摩·足疗', '/types/amzl.png', 5, '2021-12-22 20:19:27', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('美容SPA', '/types/spa.png', 6, '2021-12-22 20:19:35', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('亲子游乐', '/types/qzyl.png', 7, '2021-12-22 20:19:53', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('酒吧', '/types/jiuba.png', 8, '2021-12-22 20:20:02', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('轰趴馆', '/types/hpg.png', 9, '2021-12-22 20:20:08', '2021-12-23 11:24:31');
INSERT INTO tb_shop_type(name, icon, sort, create_time, update_time)
VALUES ('美睫·美甲', '/types/mjmj.png', 4, '2021-12-22 20:21:46', '2021-12-23 11:24:31');

-- ----------------------------
-- Table structure for tb_sign
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_sign;

-- Create the table
CREATE TABLE tb_sign
(
    id        BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id   BIGINT   NOT NULL,
    year      INT      NOT NULL,
    month     SMALLINT NOT NULL,
    date      DATE     NOT NULL,
    is_backup BOOLEAN DEFAULT NULL
);

-- Add comments to each column
COMMENT ON COLUMN tb_sign.id IS '主键';
COMMENT ON COLUMN tb_sign.user_id IS '用户id';
COMMENT ON COLUMN tb_sign.year IS '签到的年';
COMMENT ON COLUMN tb_sign.month IS '签到的月';
COMMENT ON COLUMN tb_sign.date IS '签到的日期';
COMMENT ON COLUMN tb_sign.is_backup IS '是否补签';


-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_user;

-- Create the table
CREATE TABLE tb_user
(
    id          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    phone       VARCHAR(11)                 NOT NULL,
    password    VARCHAR(128)                         DEFAULT '',
    nick_name   VARCHAR(32)                          DEFAULT '',
    icon        VARCHAR(255)                         DEFAULT '',
    create_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to each column
COMMENT ON COLUMN tb_user.id IS '主键';
COMMENT ON COLUMN tb_user.phone IS '手机号码';
COMMENT ON COLUMN tb_user.password IS '密码，加密存储';
COMMENT ON COLUMN tb_user.nick_name IS '昵称，默认是用户id';
COMMENT ON COLUMN tb_user.icon IS '人物头像';
COMMENT ON COLUMN tb_user.create_time IS '创建时间';
COMMENT ON COLUMN tb_user.update_time IS '更新时间';

-- Create a unique index on the phone column
CREATE UNIQUE INDEX uniqe_key_phone ON tb_user (phone);
-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO tb_user(phone, password, nick_name, icon, create_time, update_time)
VALUES ('13686869696', '', '小鱼同学', '/imgs/blogs/blog1.jpg', '2021-12-24 10:27:19', '2022-01-11 16:04:00');
INSERT INTO tb_user(phone, password, nick_name, icon, create_time, update_time)
VALUES ('13838411438', '', '可可今天不吃肉', '/imgs/icons/kkjtbcr.jpg', '2021-12-24 15:14:39',
        '2021-12-28 19:58:04');
INSERT INTO tb_user(phone, password, nick_name, icon, create_time, update_time)
VALUES ('13456789011', '', 'user_slxaxy2au9f3tanffaxr', '', '2022-01-07 12:07:53', '2022-01-07 12:07:53');
INSERT INTO tb_user(phone, password, nick_name, icon, create_time, update_time)
VALUES ('13456789001', '', '可爱多', '/imgs/icons/user5-icon.png', '2022-01-07 16:11:33', '2022-03-11 09:09:20');
INSERT INTO tb_user(phone, password, nick_name, icon, create_time, update_time)
VALUES ('13456762069', '', 'user_xn5wr3hpsv', '', '2022-02-07 17:54:10', '2022-02-07 17:54:10');

-- ----------------------------
-- Table structure for tb_user_info
-- ----------------------------
-- Drop the tb_user_info table if it exists
DROP TABLE IF EXISTS tb_user_info;

-- Create the tb_user_info table
CREATE TABLE tb_user_info
(
    user_id     BIGINT PRIMARY KEY,
    city        VARCHAR(64)                          DEFAULT '',
    introduce   VARCHAR(128)                         DEFAULT NULL,
    fans        INT                                  DEFAULT 0,
    followee    INT                                  DEFAULT 0,
    gender      SMALLINT                             DEFAULT 0,
    birthday    DATE                                 DEFAULT NULL,
    credits     INT                                  DEFAULT 0,
    level       SMALLINT                             DEFAULT 0,
    create_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to each column
COMMENT ON COLUMN tb_user_info.user_id IS '主键，用户id';
COMMENT ON COLUMN tb_user_info.city IS '城市名称';
COMMENT ON COLUMN tb_user_info.introduce IS '个人介绍，不要超过128个字符';
COMMENT ON COLUMN tb_user_info.fans IS '粉丝数量';
COMMENT ON COLUMN tb_user_info.followee IS '关注的人的数量';
COMMENT ON COLUMN tb_user_info.gender IS '性别，0：男，1：女';
COMMENT ON COLUMN tb_user_info.birthday IS '生日';
COMMENT ON COLUMN tb_user_info.credits IS '积分';
COMMENT ON COLUMN tb_user_info.level IS '会员级别，0~9级,0代表未开通会员';
COMMENT ON COLUMN tb_user_info.create_time IS '创建时间';
COMMENT ON COLUMN tb_user_info.update_time IS '更新时间';


-- ----------------------------
-- Table structure for tb_voucher
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_voucher;

-- Create the table
CREATE TABLE tb_voucher
(
    id           BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    shop_id      BIGINT                               DEFAULT NULL,
    title        VARCHAR(255)                NOT NULL,
    sub_title    VARCHAR(255)                         DEFAULT NULL,
    rules        VARCHAR(1024)                        DEFAULT NULL,
    pay_value    BIGINT                      NOT NULL,
    actual_value BIGINT                      NOT NULL,
    type         SMALLINT                    NOT NULL DEFAULT 0,
    status       SMALLINT                    NOT NULL DEFAULT 1,
    create_time  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to each column
COMMENT ON COLUMN tb_voucher.id IS '主键';
COMMENT ON COLUMN tb_voucher.shop_id IS '商铺id';
COMMENT ON COLUMN tb_voucher.title IS '代金券标题';
COMMENT ON COLUMN tb_voucher.sub_title IS '副标题';
COMMENT ON COLUMN tb_voucher.rules IS '使用规则';
COMMENT ON COLUMN tb_voucher.pay_value IS '支付金额，单位是分。例如200代表2元';
COMMENT ON COLUMN tb_voucher.actual_value IS '抵扣金额，单位是分。例如200代表2元';
COMMENT ON COLUMN tb_voucher.type IS '0,普通券；1,秒杀券';
COMMENT ON COLUMN tb_voucher.status IS '1,上架; 2,下架; 3,过期';
COMMENT ON COLUMN tb_voucher.create_time IS '创建时间';
COMMENT ON COLUMN tb_voucher.update_time IS '更新时间';

-- ----------------------------
-- Records of tb_voucher
-- ----------------------------
INSERT INTO tb_voucher(shop_id, title, sub_title, rules, pay_value, actual_value, type, status, create_time,
                       update_time)
VALUES (1, '50元代金券', '周一至周日均可使用', '全场通用\\n无需预约\\n可无限叠加\\不兑现、不找零\\n仅限堂食', 4750,
        5000, 0, 1, '2022-01-04 09:42:39', '2022-01-04 09:43:31');

-- ----------------------------
-- Table structure for tb_voucher_order
-- ----------------------------
-- Drop table if it exists
DROP TABLE IF EXISTS tb_voucher_order;

-- Create the table
CREATE TABLE tb_voucher_order
(
    id          BIGINT PRIMARY KEY,
    user_id     BIGINT                      NOT NULL,
    voucher_id  BIGINT                      NOT NULL,
    pay_type    SMALLINT                    NOT NULL DEFAULT 1,
    status      SMALLINT                    NOT NULL DEFAULT 1,
    create_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    pay_time    TIMESTAMP WITHOUT TIME ZONE          DEFAULT NULL,
    use_time    TIMESTAMP WITHOUT TIME ZONE          DEFAULT NULL,
    refund_time TIMESTAMP WITHOUT TIME ZONE          DEFAULT NULL,
    update_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to each column
COMMENT ON COLUMN tb_voucher_order.id IS '主键';
COMMENT ON COLUMN tb_voucher_order.user_id IS '下单的用户id';
COMMENT ON COLUMN tb_voucher_order.voucher_id IS '购买的代金券id';
COMMENT ON COLUMN tb_voucher_order.pay_type IS '支付方式 1：余额支付；2：支付宝；3：微信';
COMMENT ON COLUMN tb_voucher_order.status IS '订单状态，1：未支付；2：已支付；3：已核销；4：已取消；5：退款中；6：已退款';
COMMENT ON COLUMN tb_voucher_order.create_time IS '下单时间';
COMMENT ON COLUMN tb_voucher_order.pay_time IS '支付时间';
COMMENT ON COLUMN tb_voucher_order.use_time IS '核销时间';
COMMENT ON COLUMN tb_voucher_order.refund_time IS '退款时间';
COMMENT ON COLUMN tb_voucher_order.update_time IS '更新时间';



