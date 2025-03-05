// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.Product{})
		DB.AutoMigrate( //nolint:errcheck
			&model.Product{},
			&model.Category{},
		)
		if needDemoData {
			// DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','Sticker','Sticker')")
			// DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
			// DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")

			// 1.插入产品类别数据
			DB.Exec("INSERT INTO `product`.`category` (id, created_at, updated_at, name, description) VALUES (1, '2023-12-06 15:05:06', '2023-12-06 15:05:06', '服饰', 'Fashion and clothing items'), (2, '2023-12-06 15:05:06', '2023-12-06 15:05:06', '图书', 'Books and educational materials'), (3, '2023-12-06 15:05:06', '2023-12-06 15:05:06', '电脑', 'Computers and accessories'), (4, '2023-12-06 15:05:06', '2023-12-06 15:05:06', '母婴', 'Maternity and baby products'), (5, '2023-12-06 15:05:06', '2023-12-06 15:05:06', '美食', 'Food and beverages');")
			// 2.创建产品数据
			// 2.1 服饰类产品
			DB.Exec("INSERT INTO `product`.`product` (id, created_at, updated_at, name, description, picture, price) VALUES (1, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '【成毅同款】雪中飞冬季鹅绒工装男女羽绒服貉子毛领情侣装外套', '【成毅同款】雪中飞冬季鹅绒工装男女羽绒服貉子毛领情侣装外套', '/static/image/clothes/XZF_PARKA.jpg', 629.00), (2, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '优衣库男装女装UT Keith x coke卫衣长袖印花运动衫新款477982', '优衣库男装女装UT Keith x coke卫衣长袖印花运动衫新款477982', '/static/image/clothes/UNIQLO_hoodie.jpg', 129.00), (3, '2023-12-06 15:26:19', '2023-12-06 15:26:19', 'Nike耐克官方WINDRUNNER风行者系列男子摇粒绒里料梭织夹克HV1065', 'Nike耐克官方WINDRUNNER风行者系列男子摇粒绒里料梭织夹克HV1065', '/static/image/clothes/Nike_WINDRUNNER.jpg', 559.00), (4, '2023-12-06 15:26:19', '2023-12-06 15:26:19', 'Nike耐克官方PRIMARY DRI-FIT男子速干连帽衫防晒衣春季FZ0968', 'Nike耐克官方PRIMARY DRI-FIT男子速干连帽衫防晒衣春季FZ0968', '/static/image/clothes/Nike_PRIMARY-DRI-FIT.jpg', 329.00), (5, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '优衣库男装女装宽松直筒牛仔裤水洗产品长裤常规款475551 470542', '优衣库男装女装宽松直筒牛仔裤水洗产品长裤常规款475551 470542', '/static/image/clothes/UNIQLO_jeans.jpg', 299.00), (6, '2023-12-06 15:26:19', '2023-12-06 15:26:19', 'it Awayout男女同款阔腿裤工装裤24秋冬潮酷街头纯色直筒裤', 'it Awayout男女同款阔腿裤工装裤24秋冬潮酷街头纯色直筒裤', '/static/image/clothes/it_Awayout_Cargo_pants.jpg', 232.75);")
			// 2.2 图书类产品
			DB.Exec("INSERT INTO `product`.`product` (id, created_at, updated_at, name, description, picture, price) VALUES (11, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '我与地坛(纪念版) 史铁生散文集灵魂代表之作', '我与地坛(纪念版) 史铁生散文集灵魂代表之作', '/static/image/books/woyuditan.jpg', 14.25), (12, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '【官方正版新书】 算法竞赛入门经典 第2版第二版刘汝佳 清华大学出版', '【官方正版新书】 算法竞赛入门经典 第2版第二版刘汝佳 清华大学出版', '/static/image/books/algorithm_competition.jpg', 46.28), (13, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '【当当网正版】C Primer Plus 第六6版中文版', '【当当网正版】C Primer Plus 第六6版中文版', '/static/image/books/C++.jpg', 24.5), (14, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '病隙碎笔2021纪念版史铁生一部特殊的人生笔记我与地坛鲁迅文学奖', '病隙碎笔2021纪念版史铁生一部特殊的人生笔记我与地坛鲁迅文学奖', '/static/image/books/bingxisuibi.jpg', 24.5), (15, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '人生设计课 比尔博内特等著 中信出版社图书 正版书籍', '人生设计课 比尔博内特等著 中信出版社图书 正版书籍', '/static/image/books/Stanford.jpg', 32.3);")

			// 2.3 电脑相关
			DB.Exec("INSERT INTO `product`.`product` (id, created_at, updated_at, name, description, picture, price) VALUES (21, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '七彩虹 AOC27英寸高清2K 144Hz 180Hz IPS台式电脑组装机显示器', '七彩虹 AOC27英寸高清2K 144Hz 180Hz IPS台式电脑组装机显示器', '/static/image/computers/Colorful_aoc27.jpg', 999.00), (22, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '华硕小钢炮VG249Q3A 259电脑IPS显示器144Hz165HZ电竞官方旗舰店', '华硕小钢炮VG249Q3A 259电脑IPS显示器144Hz165HZ电竞官方旗舰店', '/static/image/computers/ASUS_vg249q3a.jpg', 799.00), (23, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '七彩虹RTX5080/5090D/5070TI火神UltraAD台式游戏电竞显卡16G/32G', '七彩虹RTX5080/5090D/5070TI火神UltraAD台式游戏电竞显卡16G/32G', '/static/image/computers/Colorful_RTX5070TI.jpg', 7599.00), (24, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '电竞叛客RTX 4060/Ti 8G白色台式电脑主机黑神话悟空游戏独立显卡', '电竞叛客RTX 4060/Ti 8G白色台式电脑主机黑神话悟空游戏独立显卡', '/static/image/computers/DJPK_RTX4060TI.jpg', 2599.00), (25, '2023-12-06 15:26:19', '2023-12-06 15:26:19', 'NIZ宁芝 68Pro RT电竞版 有线三模 动态触点游戏普拉姆静电容键盘', 'NIZ宁芝 68Pro RT电竞版 有线三模 动态触点游戏普拉姆静电容键盘', '/static/image/computers/niz68.jpg', 703.51);")

			// 2.4 母婴
			DB.Exec("INSERT INTO `product`.`product` (id, created_at, updated_at, name, description, picture, price) VALUES (31, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '【政府补贴】追觅0缠毛躺平助力洗地机母婴级吸拖扫一体机T40Pro', '【政府补贴】追觅0缠毛躺平助力洗地机母婴级吸拖扫一体机T40Pro', '/static/image/motherAndBaby/xichenqi.jpg', 1638.93), (32, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '【丁禹兮同款】超值试用全棉时代纯棉悬挂式壁挂式加厚棉柔洗脸巾', '【丁禹兮同款】超值试用全棉时代纯棉悬挂式壁挂式加厚棉柔洗脸巾', '/static/image/motherAndBaby/xilianjin.jpg', 9.90), (33, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '全棉时代新生儿见面礼宝宝周岁满月礼出生礼物初生套装婴儿礼盒', '全棉时代新生儿见面礼宝宝周岁满月礼出生礼物初生套装婴儿礼盒', '/static/image/motherAndBaby/baby_clothe.jpg', 222.00), (34, '2023-12-06 15:26:19', '2023-12-06 15:26:19', 'babycare手摇铃礼盒新生婴儿见面礼物玩具安抚巾宝宝布书满月牙胶', 'babycare手摇铃礼盒新生婴儿见面礼物玩具安抚巾宝宝布书满月牙胶', '/static/image/motherAndBaby/wanju.jpg', 269.00), (35, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '苏泊尔紫外线婴儿奶瓶消毒柜带烘干二合一恒温壶一体器泡奶BW016', '苏泊尔紫外线婴儿奶瓶消毒柜带烘干二合一恒温壶一体器泡奶BW016', '/static/image/motherAndBaby/milk_machine.jpg', 455.05);")

			// 2.5 美食
			DB.Exec("INSERT INTO `product`.`product` (id, created_at, updated_at, name, description, picture, price) VALUES (41, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '三只松鼠_全麦欧包1000g 0蔗糖粗粮夹心面包代餐零食早餐', '三只松鼠_全麦欧包1000g 0蔗糖粗粮夹心面包代餐零食早餐', '/static/image/food/3songshu.jpg', 9.31), (42, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '海南特产南国椰香薄饼160gX4盒代餐食品土特产零食薄饼脆饼椰子饼', '海南特产南国椰香薄饼160gX4盒代餐食品土特产零食薄饼脆饼椰子饼', '/static/image/food/yezibing.jpg', 15.34), (43, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '山姆超市休闲零食海盐苏打饼干1.5kg', '山姆超市休闲零食海盐苏打饼干1.5kg', '/static/image/food/shanmu.jpg', 39.80), (44, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '盼盼悠点咸味苏打饼干早餐休闲零食酥脆饼干下午茶饱腹充饥小吃', '盼盼悠点咸味苏打饼干早餐休闲零食酥脆饼干下午茶饱腹充饥小吃', '/static/image/food/panpan.jpg', 8.8), (45, '2023-12-06 15:26:19', '2023-12-06 15:26:19', '吴京代言 双汇王中王火腿肠44g*10支*1袋香肠零食泡面搭档肉肠官', '吴京代言 双汇王中王火腿肠44g*10支*1袋香肠零食泡面搭档肉肠官', '/static/image/food/shuanghui.jpg', 13.9);")

			// 3. product_category 数据
			DB.Exec("INSERT INTO `product`.`product_category` (product_id, category_id) VALUES (1, 1), (2, 1), (3, 1), (4, 1), (5, 1), (6, 1);")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id, category_id) VALUES (11, 2), (12, 2), (13, 2), (14, 2), (15, 2);")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id, category_id) VALUES (21, 3), (22, 3), (23, 3), (24, 3), (25, 3);")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id, category_id) VALUES (31, 4), (32, 4), (33, 4), (34, 4), (35, 4);")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id, category_id) VALUES (41, 5), (42, 5), (43, 5), (44, 5), (45, 5);")
		}
	}
}
