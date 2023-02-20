# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.7.4-MariaDB-1:10.7.4+maria~focal-log)
# Database: admin
# Generation Time: 2023-02-20 07:05:39 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table article
# ------------------------------------------------------------

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `title` varchar(128) DEFAULT NULL COMMENT '標題',
  `author` varchar(128) DEFAULT NULL COMMENT '作者',
  `content` varchar(255) DEFAULT NULL COMMENT '内容',
  `status` int(1) DEFAULT NULL COMMENT '狀態',
  `publish_at` timestamp NULL DEFAULT NULL COMMENT '發佈時間',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_article_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章';

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;

INSERT INTO `article` (`id`, `title`, `author`, `content`, `status`, `publish_at`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`)
VALUES
	(1,'sales1','sales1','sales1',1,'2023-02-20 12:19:15','2023-02-20 12:19:20','2023-02-20 12:19:20',NULL,3,0),
	(2,'sales1-1','sales1-1','sales1-1',1,'2023-02-20 12:20:23','2023-02-20 12:20:24','2023-02-20 12:20:24',NULL,6,0),
	(3,'sales2-1','sales2-1','sales2-1',1,'2023-02-20 12:46:49','2023-02-20 12:46:50','2023-02-20 12:46:50',NULL,5,0);

/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_api
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_api`;

CREATE TABLE `sys_api` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `handle` varchar(128) DEFAULT NULL COMMENT 'handle',
  `title` varchar(128) DEFAULT NULL COMMENT '標题',
  `path` varchar(128) DEFAULT NULL COMMENT '地址',
  `type` varchar(16) DEFAULT NULL COMMENT 'API類型',
  `action` varchar(16) DEFAULT NULL COMMENT '請求類型',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_deleted_at` (`deleted_at`),
  KEY `idx_sys_api_create_by` (`create_by`),
  KEY `idx_sys_api_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_api` WRITE;
/*!40000 ALTER TABLE `sys_api` DISABLE KEYS */;

INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `type`, `action`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`)
VALUES
	(5,'go-admin/app/admin/apis.SysLoginLog.Get-fm','登入Log透過id獲取','/api/v1/sys-login-log/:id','BUS','GET','2021-05-13 19:59:00.728','2021-06-17 11:37:12.331',NULL,0,0),
	(6,'go-admin/app/admin/apis.SysOperaLog.GetPage-fm','操作Log列表','/api/v1/sys-opera-log','BUS','GET','2021-05-13 19:59:00.778','2021-06-17 11:48:40.732',NULL,0,0),
	(7,'go-admin/app/admin/apis.SysOperaLog.Get-fm','操作Log透過id獲取','/api/v1/sys-opera-log/:id','BUS','GET','2021-05-13 19:59:00.821','2021-06-16 21:49:48.598',NULL,0,0),
	(8,'go-admin/common/actions.IndexAction.func1','分類列表','/api/v1/syscategory','BUS','GET','2021-05-13 19:59:00.870','2021-06-13 20:53:47.883',NULL,0,0),
	(9,'go-admin/common/actions.ViewAction.func1','分類透過id獲取','/api/v1/syscategory/:id','BUS','GET','2021-05-13 19:59:00.945','2021-06-13 20:53:47.926',NULL,0,0),
	(10,'go-admin/common/actions.IndexAction.func1','内容列表','/api/v1/syscontent','BUS','GET','2021-05-13 19:59:01.005','2021-06-13 20:53:47.966',NULL,0,0),
	(11,'go-admin/common/actions.ViewAction.func1','内容透過id獲取','/api/v1/syscontent/:id','BUS','GET','2021-05-13 19:59:01.056','2021-06-13 20:53:48.005',NULL,0,0),
	(15,'go-admin/common/actions.IndexAction.func1','job列表','/api/v1/sysjob','BUS','GET','2021-05-13 19:59:01.248','2021-06-13 20:53:48.169',NULL,0,0),
	(16,'go-admin/common/actions.ViewAction.func1','job透過id獲取','/api/v1/sysjob/:id','BUS','GET','2021-05-13 19:59:01.298','2021-06-13 20:53:48.214',NULL,0,0),
	(21,'go-admin/app/admin/apis.SysDictType.GetPage-fm','字典類型列表','/api/v1/dict/type','BUS','GET','2021-05-13 19:59:01.525','2021-06-17 11:48:40.732',NULL,0,0),
	(22,'go-admin/app/admin/apis.SysDictType.GetAll-fm','字典類型查詢【代碼生成】','/api/v1/dict/type-option-select','SYS','GET','2021-05-13 19:59:01.582','2021-06-13 20:53:48.388',NULL,0,0),
	(23,'go-admin/app/admin/apis.SysDictType.Get-fm','字典類型透過id獲取','/api/v1/dict/type/:id','BUS','GET','2021-05-13 19:59:01.632','2021-06-17 11:48:40.732',NULL,0,0),
	(24,'go-admin/app/admin/apis.SysDictData.GetPage-fm','字典資料列表','/api/v1/dict/data','BUS','GET','2021-05-13 19:59:01.684','2021-06-17 11:48:40.732',NULL,0,0),
	(25,'go-admin/app/admin/apis.SysDictData.Get-fm','字典資料透過code獲取','/api/v1/dict/data/:dictCode','BUS','GET','2021-05-13 19:59:01.732','2021-06-17 11:48:40.732',NULL,0,0),
	(26,'go-admin/app/admin/apis.SysDictData.GetSysDictDataAll-fm','資料字典根據key獲取','/api/v1/dict-data/option-select','SYS','GET','2021-05-13 19:59:01.832','2021-06-17 11:48:40.732',NULL,0,0),
	(27,'go-admin/app/admin/apis.SysDept.GetPage-fm','部門列表','/api/v1/dept','BUS','GET','2021-05-13 19:59:01.940','2021-06-17 11:48:40.732',NULL,0,0),
	(28,'go-admin/app/admin/apis.SysDept.Get-fm','部門透過id獲取','/api/v1/dept/:id','BUS','GET','2021-05-13 19:59:02.009','2021-06-17 11:48:40.732',NULL,0,0),
	(29,'go-admin/app/admin/apis.SysDept.Get2Tree-fm','查詢部門下拉樹【角色權限-自定權限】','/api/v1/deptTree','SYS','GET','2021-05-13 19:59:02.050','2021-06-17 11:48:40.732',NULL,0,0),
	(30,'go-admin/app/admin/apis/tools.(*Gen).GetDBTableList-fm','資料庫表列表','/api/v1/db/tables/page','SYS','GET','2021-05-13 19:59:02.098','2021-06-13 20:53:48.730',NULL,0,0),
	(31,'go-admin/app/admin/apis/tools.(*Gen).GetDBColumnList-fm','資料表列列表','/api/v1/db/columns/page','SYS','GET','2021-05-13 19:59:02.140','2021-06-13 20:53:48.771',NULL,0,0),
	(32,'go-admin/app/admin/apis/tools.Gen.GenCode-fm','資料庫表生成到項目','/api/v1/gen/toproject/:tableId','SYS','GET','2021-05-13 19:59:02.183','2021-06-13 20:53:48.812',NULL,0,0),
	(33,'go-admin/app/admin/apis/tools.Gen.GenMenuAndApi-fm','資料庫表生成到DB','/api/v1/gen/todb/:tableId','SYS','GET','2021-05-13 19:59:02.227','2021-06-13 20:53:48.853',NULL,0,0),
	(34,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTablesTree-fm','關係表資料【代碼生成】','/api/v1/gen/tabletree','SYS','GET','2021-05-13 19:59:02.271','2021-06-13 20:53:48.893',NULL,0,0),
	(35,'go-admin/app/admin/apis/tools.Gen.Preview-fm','生成預覽透過id獲取','/api/v1/gen/preview/:tableId','SYS','GET','2021-05-13 19:59:02.315','2021-06-13 20:53:48.935',NULL,0,0),
	(36,'go-admin/app/admin/apis/tools.Gen.GenApiToFile-fm','生成api帶文件','/api/v1/gen/apitofile/:tableId','SYS','GET','2021-05-13 19:59:02.357','2021-06-13 20:53:48.977',NULL,0,0),
	(37,'go-admin/app/admin/apis.System.GenerateCaptchaHandler-fm','驗證碼獲取','/api/v1/getCaptcha','SYS','GET','2021-05-13 19:59:02.405','2021-06-13 20:53:49.020',NULL,0,0),
	(38,'go-admin/app/admin/apis.SysUser.GetInfo-fm','User資訊獲取','/api/v1/getinfo','SYS','GET','2021-05-13 19:59:02.447','2021-06-13 20:53:49.065',NULL,0,0),
	(39,'go-admin/app/admin/apis.SysMenu.GetPage-fm','選單列表','/api/v1/menu','BUS','GET','2021-05-13 19:59:02.497','2021-06-17 11:48:40.732',NULL,0,0),
	(40,'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm','查詢選單下拉樹結構【廢棄】','/api/v1/menuTreeselect','SYS','GET','2021-05-13 19:59:02.542','2021-06-03 22:37:21.857',NULL,0,0),
	(41,'go-admin/app/admin/apis.SysMenu.Get-fm','選單透過id獲取','/api/v1/menu/:id','BUS','GET','2021-05-13 19:59:02.584','2021-06-17 11:48:40.732',NULL,0,0),
	(42,'go-admin/app/admin/apis.SysMenu.GetMenuRole-fm','角色選單【頂部左侧選單】','/api/v1/menurole','SYS','GET','2021-05-13 19:59:02.630','2021-06-13 20:53:49.574',NULL,0,0),
	(43,'go-admin/app/admin/apis.SysMenu.GetMenuIDS-fm','獲取角色對應的選單id資料【廢棄】','/api/v1/menuids','SYS','GET','2021-05-13 19:59:02.675','2021-06-03 22:39:52.500',NULL,0,0),
	(44,'go-admin/app/admin/apis.SysRole.GetPage-fm','角色列表','/api/v1/role','BUS','GET','2021-05-13 19:59:02.720','2021-06-17 11:48:40.732',NULL,0,0),
	(45,'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm','選單權限列表【角色配選單使用】','/api/v1/roleMenuTreeselect/:roleId','SYS','GET','2021-05-13 19:59:02.762','2021-06-17 11:48:40.732',NULL,0,0),
	(46,'go-admin/app/admin/apis.SysDept.GetDeptTreeRoleSelect-fm','角色部門結構樹【自定義資料權限】','/api/v1/roleDeptTreeselect/:roleId','SYS','GET','2021-05-13 19:59:02.809','2021-06-17 11:48:40.732',NULL,0,0),
	(47,'go-admin/app/admin/apis.SysRole.Get-fm','角色透過id獲取','/api/v1/role/:id','BUS','GET','2021-05-13 19:59:02.850','2021-06-17 11:48:40.732',NULL,0,0),
	(48,'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).RefreshHandler-fm','刷新token','/api/v1/refresh_token','SYS','GET','2021-05-13 19:59:02.892','2021-06-13 20:53:49.278',NULL,0,0),
	(53,'go-admin/app/admin/apis.SysConfig.GetPage-fm','參數列表','/api/v1/config','BUS','GET','2021-05-13 19:59:03.116','2021-06-17 11:48:40.732',NULL,0,0),
	(54,'go-admin/app/admin/apis.SysConfig.Get-fm','參數透過id獲取','/api/v1/config/:id','BUS','GET','2021-05-13 19:59:03.157','2021-06-17 11:48:40.732',NULL,0,0),
	(55,'go-admin/app/admin/apis.SysConfig.GetSysConfigByKEYForService-fm','參數透過鍵名搜索【基礎預設配置】','/api/v1/configKey/:configKey','SYS','GET','2021-05-13 19:59:03.198','2021-06-13 20:53:49.745',NULL,0,0),
	(57,'go-admin/app/jobs/apis.SysJob.RemoveJobForService-fm','job移除','/api/v1/job/remove/:id','BUS','GET','2021-05-13 19:59:03.295','2021-06-13 20:53:49.786',NULL,0,0),
	(58,'go-admin/app/jobs/apis.SysJob.StartJobForService-fm','job啟動','/api/v1/job/start/:id','BUS','GET','2021-05-13 19:59:03.339','2021-06-13 20:53:49.829',NULL,0,0),
	(59,'go-admin/app/admin/apis.SysPost.GetPage-fm','職位列表','/api/v1/post','BUS','GET','2021-05-13 19:59:03.381','2021-06-17 11:48:40.732',NULL,0,0),
	(60,'go-admin/app/admin/apis.SysPost.Get-fm','職位透過id獲取','/api/v1/post/:id','BUS','GET','2021-05-13 19:59:03.433','2021-06-17 11:48:40.732',NULL,0,0),
	(62,'go-admin/app/admin/apis.SysConfig.GetSysConfigBySysApp-fm','系統前端參數','/api/v1/app-config','SYS','GET','2021-05-13 19:59:03.526','2021-06-13 20:53:49.994',NULL,0,0),
	(63,'go-admin/app/admin/apis.SysUser.GetProfile-fm','*User資訊獲取','/api/v1/user/profile','SYS','GET','2021-05-13 19:59:03.567','2021-06-13 20:53:50.038',NULL,0,0),
	(66,'github.com/go-admin-team/go-admin-core/sdk/pkg/ws.(*Manager).WsClient-fm','連接ws【joblog】','/ws/:id/:channel','BUS','GET','2021-05-13 19:59:03.705','2021-06-13 20:53:50.167',NULL,0,0),
	(67,'github.com/go-admin-team/go-admin-core/sdk/pkg/ws.(*Manager).UnWsClient-fm','登出ws【joblog】','/wslogout/:id/:channel','BUS','GET','2021-05-13 19:59:03.756','2021-06-13 20:53:50.209',NULL,0,0),
	(68,'go-admin/common/middleware/handler.Ping','*User基本資訊','/info','SYS','GET','2021-05-13 19:59:03.800','2021-06-13 20:53:50.251',NULL,0,0),
	(72,'go-admin/common/actions.CreateAction.func1','分類建立','/api/v1/syscategory','BUS','POST','2021-05-13 19:59:03.982','2021-06-13 20:53:50.336',NULL,0,0),
	(73,'go-admin/common/actions.CreateAction.func1','内容建立','/api/v1/syscontent','BUS','POST','2021-05-13 19:59:04.027','2021-06-13 20:53:50.375',NULL,0,0),
	(76,'go-admin/common/actions.CreateAction.func1','job建立','/api/v1/sysjob','BUS','POST','2021-05-13 19:59:04.164','2021-06-13 20:53:50.500',NULL,0,0),
	(80,'go-admin/app/admin/apis.SysDictData.Insert-fm','字典資料建立','/api/v1/dict/data','BUS','POST','2021-05-13 19:59:04.347','2021-06-17 11:48:40.732',NULL,0,0),
	(81,'go-admin/app/admin/apis.SysDictType.Insert-fm','字典類型建立','/api/v1/dict/type','BUS','POST','2021-05-13 19:59:04.391','2021-06-17 11:48:40.732',NULL,0,0),
	(82,'go-admin/app/admin/apis.SysDept.Insert-fm','部門建立','/api/v1/dept','BUS','POST','2021-05-13 19:59:04.435','2021-06-17 11:48:40.732',NULL,0,0),
	(85,'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).LoginHandler-fm','*登入','/api/v1/login','SYS','POST','2021-05-13 19:59:04.597','2021-06-13 20:53:50.784',NULL,0,0),
	(86,'go-admin/common/middleware/handler.LogOut','*登出','/api/v1/logout','SYS','POST','2021-05-13 19:59:04.642','2021-06-13 20:53:50.824',NULL,0,0),
	(87,'go-admin/app/admin/apis.SysConfig.Insert-fm','參數建立','/api/v1/config','BUS','POST','2021-05-13 19:59:04.685','2021-06-17 11:48:40.732',NULL,0,0),
	(88,'go-admin/app/admin/apis.SysMenu.Insert-fm','選單建立','/api/v1/menu','BUS','POST','2021-05-13 19:59:04.777','2021-06-17 11:48:40.732',NULL,0,0),
	(89,'go-admin/app/admin/apis.SysPost.Insert-fm','職位建立','/api/v1/post','BUS','POST','2021-05-13 19:59:04.886','2021-06-17 11:48:40.732',NULL,0,0),
	(90,'go-admin/app/admin/apis.SysRole.Insert-fm','角色建立','/api/v1/role','BUS','POST','2021-05-13 19:59:04.975','2021-06-17 11:48:40.732',NULL,0,0),
	(91,'go-admin/app/admin/apis.SysUser.InsetAvatar-fm','*User頭像編輯','/api/v1/user/avatar','SYS','POST','2021-05-13 19:59:05.058','2021-06-13 20:53:51.079',NULL,0,0),
	(92,'go-admin/app/admin/apis.SysApi.Update-fm','API編輯','/api/v1/sys-api/:id','BUS','PUT','2021-05-13 19:59:05.122','2023-02-20 12:20:05.454',NULL,0,0),
	(95,'go-admin/common/actions.UpdateAction.func1','分類編輯','/api/v1/syscategory/:id','BUS','PUT','2021-05-13 19:59:05.255','2021-06-13 20:53:51.247',NULL,0,0),
	(96,'go-admin/common/actions.UpdateAction.func1','内容編輯','/api/v1/syscontent/:id','BUS','PUT','2021-05-13 19:59:05.299','2021-06-13 20:53:51.289',NULL,0,0),
	(97,'go-admin/common/actions.UpdateAction.func1','job編輯','/api/v1/sysjob','BUS','PUT','2021-05-13 19:59:05.343','2021-06-13 20:53:51.331',NULL,0,0),
	(101,'go-admin/app/admin/apis.SysDictData.Update-fm','字典資料編輯','/api/v1/dict/data/:dictCode','BUS','PUT','2021-05-13 19:59:05.519','2021-06-17 11:48:40.732',NULL,0,0),
	(102,'go-admin/app/admin/apis.SysDictType.Update-fm','字典類型編輯','/api/v1/dict/type/:id','BUS','PUT','2021-05-13 19:59:05.569','2021-06-17 11:48:40.732',NULL,0,0),
	(103,'go-admin/app/admin/apis.SysDept.Update-fm','部門編輯','/api/v1/dept/:id','BUS','PUT','2021-05-13 19:59:05.613','2021-06-17 11:48:40.732',NULL,0,0),
	(104,'go-admin/app/other/apis.SysFileDir.Update-fm','文件夾編輯','/api/v1/file-dir/:id','BUS','PUT','2021-05-13 19:59:05.662','2021-06-13 20:53:51.847',NULL,0,0),
	(105,'go-admin/app/other/apis.SysFileInfo.Update-fm','文件編輯','/api/v1/file-info/:id','BUS','PUT','2021-05-13 19:59:05.709','2021-06-13 20:53:51.892',NULL,0,0),
	(106,'go-admin/app/admin/apis.SysRole.Update-fm','角色編輯','/api/v1/role/:id','BUS','PUT','2021-05-13 19:59:05.752','2021-06-17 11:48:40.732',NULL,0,0),
	(107,'go-admin/app/admin/apis.SysRole.Update2DataScope-fm','角色資料權限更新','/api/v1/roledatascope','BUS','PUT','2021-05-13 19:59:05.803','2021-06-17 11:48:40.732',NULL,0,0),
	(108,'go-admin/app/admin/apis.SysConfig.Update-fm','參數編輯','/api/v1/config/:id','BUS','PUT','2021-05-13 19:59:05.848','2021-06-17 11:48:40.732',NULL,0,0),
	(109,'go-admin/app/admin/apis.SysMenu.Update-fm','編輯選單','/api/v1/menu/:id','BUS','PUT','2021-05-13 19:59:05.891','2021-06-17 11:48:40.732',NULL,0,0),
	(110,'go-admin/app/admin/apis.SysPost.Update-fm','職位編輯','/api/v1/post/:id','BUS','PUT','2021-05-13 19:59:05.934','2021-06-17 11:48:40.732',NULL,0,0),
	(111,'go-admin/app/admin/apis.SysUser.UpdatePwd-fm','*User更新密碼','/api/v1/user/pwd','SYS','PUT','2021-05-13 19:59:05.987','2021-06-13 20:53:51.724',NULL,0,0),
	(112,'go-admin/common/actions.DeleteAction.func1','分類刪除','/api/v1/syscategory','BUS','DELETE','2021-05-13 19:59:06.030','2021-06-13 20:53:52.237',NULL,0,0),
	(113,'go-admin/common/actions.DeleteAction.func1','内容刪除','/api/v1/syscontent','BUS','DELETE','2021-05-13 19:59:06.076','2021-06-13 20:53:52.278',NULL,0,0),
	(114,'go-admin/app/admin/apis.SysLoginLog.Delete-fm','登入Log刪除','/api/v1/sys-login-log','BUS','DELETE','2021-05-13 19:59:06.118','2021-06-17 11:48:40.732',NULL,0,0),
	(115,'go-admin/app/admin/apis.SysOperaLog.Delete-fm','操作Log刪除','/api/v1/sys-opera-log','BUS','DELETE','2021-05-13 19:59:06.162','2021-06-17 11:48:40.732',NULL,0,0),
	(116,'go-admin/common/actions.DeleteAction.func1','job刪除','/api/v1/sysjob','BUS','DELETE','2021-05-13 19:59:06.206','2021-06-13 20:53:52.323',NULL,0,0),
	(117,'go-admin/app/other/apis.SysChinaAreaData.Delete-fm','行政區刪除','/api/v1/sys-area-data','BUS','DELETE','2021-05-13 19:59:06.249','2021-06-13 20:53:52.061',NULL,0,0),
	(120,'go-admin/app/admin/apis.SysDictData.Delete-fm','字典資料刪除','/api/v1/dict/data','BUS','DELETE','2021-05-13 19:59:06.387','2021-06-17 11:48:40.732',NULL,0,0),
	(121,'go-admin/app/admin/apis.SysDictType.Delete-fm','字典類型刪除','/api/v1/dict/type','BUS','DELETE','2021-05-13 19:59:06.432','2021-06-17 11:48:40.732',NULL,0,0),
	(122,'go-admin/app/admin/apis.SysDept.Delete-fm','部門刪除','/api/v1/dept/:id','BUS','DELETE','2021-05-13 19:59:06.475','2021-06-17 11:48:40.732',NULL,0,0),
	(123,'go-admin/app/other/apis.SysFileDir.Delete-fm','文件夾刪除','/api/v1/file-dir/:id','BUS','DELETE','2021-05-13 19:59:06.520','2021-06-13 20:53:52.539',NULL,0,0),
	(124,'go-admin/app/other/apis.SysFileInfo.Delete-fm','文件刪除','/api/v1/file-info/:id','BUS','DELETE','2021-05-13 19:59:06.566','2021-06-13 20:53:52.580',NULL,0,0),
	(125,'go-admin/app/admin/apis.SysConfig.Delete-fm','參數刪除','/api/v1/config','BUS','DELETE','2021-05-13 19:59:06.612','2021-06-17 11:48:40.732',NULL,0,0),
	(126,'go-admin/app/admin/apis.SysMenu.Delete-fm','刪除選單','/api/v1/menu','BUS','DELETE','2021-05-13 19:59:06.654','2021-06-17 11:48:40.732',NULL,0,0),
	(127,'go-admin/app/admin/apis.SysPost.Delete-fm','職位刪除','/api/v1/post/:id','BUS','DELETE','2021-05-13 19:59:06.702','2021-06-17 11:48:40.732',NULL,0,0),
	(128,'go-admin/app/admin/apis.SysRole.Delete-fm','角色刪除','/api/v1/role','BUS','DELETE','2021-05-13 19:59:06.746','2021-06-17 11:48:40.732',NULL,0,0),
	(131,'github.com/go-admin-team/go-admin-core/tools/transfer.Handler.func1','系統指標','/api/v1/metrics','SYS','GET','2021-05-17 22:13:55.933','2021-06-13 20:53:49.614',NULL,0,0),
	(132,'go-admin/app/other/router.registerMonitorRouter.func1','健康狀態','/api/v1/health','SYS','GET','2021-05-17 22:13:56.285','2021-06-13 20:53:49.951',NULL,0,0),
	(133,'go-admin/app/admin/apis.HelloWorld','項目預設API','/','SYS','GET','2021-05-24 10:30:44.553','2021-06-13 20:53:47.406',NULL,0,0),
	(134,'go-admin/app/other/apis.ServerMonitor.ServerInfo-fm','服務器基本狀態','/api/v1/server-monitor','SYS','GET','2021-05-24 10:30:44.937','2021-06-13 20:53:48.255',NULL,0,0),
	(135,'go-admin/app/admin/apis.SysApi.GetPage-fm','API列表','/api/v1/sys-api','BUS','GET','2021-05-24 11:37:53.303','2023-02-20 12:48:41.919',NULL,0,0),
	(136,'go-admin/app/admin/apis.SysApi.Get-fm','API透過id獲取','/api/v1/sys-api/:id','BUS','GET','2021-05-24 11:37:53.359','2023-02-20 12:20:05.454',NULL,0,0),
	(137,'go-admin/app/admin/apis.SysLoginLog.GetPage-fm','登入Log列表','/api/v1/sys-login-log','BUS','GET','2021-05-24 11:47:30.397','2021-06-17 11:48:40.732',NULL,0,0),
	(138,'go-admin/app/other/apis.File.UploadFile-fm','文件上傳','/api/v1/public/uploadFile','SYS','POST','2021-05-25 19:16:18.493','2021-06-13 20:53:50.866',NULL,0,0),
	(139,'go-admin/app/admin/apis.SysConfig.Update2Set-fm','參數資訊更新【參數配置】','/api/v1/set-config','BUS','PUT','2021-05-27 09:45:14.853','2021-06-17 11:48:40.732',NULL,0,0),
	(140,'go-admin/app/admin/apis.SysConfig.Get2Set-fm','參數獲取全部【配置使用】','/api/v1/set-config','BUS','GET','2021-05-27 11:54:14.384','2021-06-17 11:48:40.732',NULL,0,0),
	(141,'go-admin/app/admin/apis.SysUser.GetPage-fm','管理員列表','/api/v1/sys-user','BUS','GET','2021-06-13 19:24:57.111','2021-06-17 20:31:14.318',NULL,0,0),
	(142,'go-admin/app/admin/apis.SysUser.Get-fm','管理員透過id獲取','/api/v1/sys-user/:id','BUS','GET','2021-06-13 19:24:57.188','2021-06-17 20:31:14.318',NULL,0,0),
	(143,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTablesInfo-fm','','/api/v1/sys/tables/info','','GET','2021-06-13 19:24:57.437','2021-06-13 20:53:48.047',NULL,0,0),
	(144,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTables-fm','','/api/v1/sys/tables/info/:tableId','','GET','2021-06-13 19:24:57.510','2021-06-13 20:53:48.088',NULL,0,0),
	(145,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTableList-fm','','/api/v1/sys/tables/page','','GET','2021-06-13 19:24:57.582','2021-06-13 20:53:48.128',NULL,0,0),
	(146,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/static/*filepath','','GET','2021-06-13 19:24:59.641','2021-06-13 20:53:50.081',NULL,0,0),
	(147,'github.com/swaggo/gin-swagger.CustomWrapHandler.func1','','/swagger/*any','','GET','2021-06-13 19:24:59.713','2021-06-13 20:53:50.123',NULL,0,0),
	(148,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/form-generator/*filepath','','GET','2021-06-13 19:24:59.914','2021-06-13 20:53:50.295',NULL,0,0),
	(149,'go-admin/app/admin/apis/tools.(*SysTable).InsertSysTable-fm','','/api/v1/sys/tables/info','','POST','2021-06-13 19:25:00.163','2021-06-13 20:53:50.539',NULL,0,0),
	(150,'go-admin/app/admin/apis.SysUser.Insert-fm','管理員建立','/api/v1/sys-user','BUS','POST','2021-06-13 19:25:00.233','2021-06-17 20:31:14.318',NULL,0,0),
	(151,'go-admin/app/admin/apis.SysUser.Update-fm','管理員編輯','/api/v1/sys-user','BUS','PUT','2021-06-13 19:25:00.986','2021-06-17 20:31:14.318',NULL,0,0),
	(152,'go-admin/app/admin/apis/tools.(*SysTable).UpdateSysTable-fm','','/api/v1/sys/tables/info','','PUT','2021-06-13 19:25:01.149','2021-06-13 20:53:51.375',NULL,0,0),
	(153,'go-admin/app/admin/apis.SysRole.Update2Status-fm','','/api/v1/role-status','','PUT','2021-06-13 19:25:01.446','2021-06-13 20:53:51.636',NULL,0,0),
	(154,'go-admin/app/admin/apis.SysUser.ResetPwd-fm','','/api/v1/user/pwd/reset','','PUT','2021-06-13 19:25:01.601','2021-06-13 20:53:51.764',NULL,0,0),
	(155,'go-admin/app/admin/apis.SysUser.UpdateStatus-fm','','/api/v1/user/status','','PUT','2021-06-13 19:25:01.671','2021-06-13 20:53:51.806',NULL,0,0),
	(156,'go-admin/app/admin/apis.SysUser.Delete-fm','管理員刪除','/api/v1/sys-user','BUS','DELETE','2021-06-13 19:25:02.043','2021-06-17 20:31:14.318',NULL,0,0),
	(157,'go-admin/app/admin/apis/tools.(*SysTable).DeleteSysTables-fm','','/api/v1/sys/tables/info/:tableId','','DELETE','2021-06-13 19:25:02.283','2021-06-13 20:53:52.367',NULL,0,0),
	(158,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/static/*filepath','','HEAD','2021-06-13 19:25:02.734','2021-06-13 20:53:52.791',NULL,0,0),
	(159,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/form-generator/*filepath','','HEAD','2021-06-13 19:25:02.808','2021-06-13 20:53:52.838',NULL,0,0),
	(160,'go-admin/app/admin/apis.System.GenerateCaptchaHandler-fm','','/api/v1/captcha','','GET','2023-02-20 12:08:58.894','2023-02-20 12:08:58.894',NULL,0,0),
	(161,'go-admin/app/admin/apis.Article.GetPage-fm','取得文章列表','/api/v1/article','','GET','2023-02-20 12:08:58.908','2023-02-20 12:20:05.454',NULL,0,0),
	(162,'go-admin/app/admin/apis.Article.Get-fm','取得單一文章','/api/v1/article/:id','','GET','2023-02-20 12:08:58.912','2023-02-20 12:20:05.454',NULL,0,0),
	(163,'github.com/swaggo/gin-swagger.CustomWrapHandler.func1','','/swagger/admin/*any','','GET','2023-02-20 12:08:58.929','2023-02-20 12:08:58.929',NULL,0,0),
	(164,'go-admin/app/admin/apis.Article.Insert-fm','新增文章','/api/v1/article','','POST','2023-02-20 12:08:58.954','2023-02-20 12:20:05.454',NULL,0,0),
	(165,'go-admin/app/admin/apis.SysUser.UpdatePwd-fm','','/api/v1/user/pwd/set','','PUT','2023-02-20 12:08:58.980','2023-02-20 12:08:58.980',NULL,0,0),
	(166,'go-admin/app/admin/apis.Article.Update-fm','更新文章','/api/v1/article/:id','','PUT','2023-02-20 12:08:58.988','2023-02-20 12:20:05.454',NULL,0,0),
	(167,'go-admin/app/admin/apis.SysDept.Delete-fm','','/api/v1/dept','','DELETE','2023-02-20 12:08:59.007','2023-02-20 12:08:59.007',NULL,0,0),
	(168,'go-admin/app/admin/apis.Article.Delete-fm','刪除文章','/api/v1/article','','DELETE','2023-02-20 12:08:59.012','2023-02-20 12:20:05.454',NULL,0,0),
	(169,'go-admin/app/admin/apis.SysPost.Delete-fm','','/api/v1/post','','DELETE','2023-02-20 12:08:59.019','2023-02-20 12:08:59.019',NULL,0,0);

/*!40000 ALTER TABLE `sys_api` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_casbin_rule
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_casbin_rule`;

CREATE TABLE `sys_casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v6` varchar(25) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v7` varchar(25) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sys_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_casbin_rule` WRITE;
/*!40000 ALTER TABLE `sys_casbin_rule` DISABLE KEYS */;

INSERT INTO `sys_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`)
VALUES
	(8,'p','sales-team-lead','/api/v1/article','DELETE','','','','',''),
	(4,'p','sales-team-lead','/api/v1/article','GET','','','','',''),
	(6,'p','sales-team-lead','/api/v1/article','POST','','','','',''),
	(5,'p','sales-team-lead','/api/v1/article/:id','GET','','','','',''),
	(7,'p','sales-team-lead','/api/v1/article/:id','PUT','','','','',''),
	(1,'p','sales-team-lead','/api/v1/sys-api','GET','','','','',''),
	(3,'p','sales-team-lead','/api/v1/sys-api/:id','GET','','','','',''),
	(2,'p','sales-team-lead','/api/v1/sys-api/:id','PUT','','','','',''),
	(16,'p','sales-team-member','/api/v1/article','DELETE','','','','',''),
	(12,'p','sales-team-member','/api/v1/article','GET','','','','',''),
	(14,'p','sales-team-member','/api/v1/article','POST','','','','',''),
	(13,'p','sales-team-member','/api/v1/article/:id','GET','','','','',''),
	(15,'p','sales-team-member','/api/v1/article/:id','PUT','','','','',''),
	(9,'p','sales-team-member','/api/v1/sys-api','GET','','','','',''),
	(11,'p','sales-team-member','/api/v1/sys-api/:id','GET','','','','',''),
	(10,'p','sales-team-member','/api/v1/sys-api/:id','PUT','','','','','');

/*!40000 ALTER TABLE `sys_casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_columns
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_columns`;

CREATE TABLE `sys_columns` (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `table_id` bigint(20) DEFAULT NULL,
  `column_name` varchar(128) DEFAULT NULL,
  `column_comment` varchar(128) DEFAULT NULL,
  `column_type` varchar(128) DEFAULT NULL,
  `go_type` varchar(128) DEFAULT NULL,
  `go_field` varchar(128) DEFAULT NULL,
  `json_field` varchar(128) DEFAULT NULL,
  `is_pk` varchar(4) DEFAULT NULL,
  `is_increment` varchar(4) DEFAULT NULL,
  `is_required` varchar(4) DEFAULT NULL,
  `is_insert` varchar(4) DEFAULT NULL,
  `is_edit` varchar(4) DEFAULT NULL,
  `is_list` varchar(4) DEFAULT NULL,
  `is_query` varchar(4) DEFAULT NULL,
  `query_type` varchar(128) DEFAULT NULL,
  `html_type` varchar(128) DEFAULT NULL,
  `dict_type` varchar(128) DEFAULT NULL,
  `sort` bigint(20) DEFAULT NULL,
  `list` varchar(1) DEFAULT NULL,
  `pk` tinyint(1) DEFAULT NULL,
  `required` tinyint(1) DEFAULT NULL,
  `super_column` tinyint(1) DEFAULT NULL,
  `usable_column` tinyint(1) DEFAULT NULL,
  `increment` tinyint(1) DEFAULT NULL,
  `insert` tinyint(1) DEFAULT NULL,
  `edit` tinyint(1) DEFAULT NULL,
  `query` tinyint(1) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `fk_table_name` longtext DEFAULT NULL,
  `fk_table_name_class` longtext DEFAULT NULL,
  `fk_table_name_package` longtext DEFAULT NULL,
  `fk_label_id` longtext DEFAULT NULL,
  `fk_label_name` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`column_id`),
  KEY `idx_sys_columns_deleted_at` (`deleted_at`),
  KEY `idx_sys_columns_create_by` (`create_by`),
  KEY `idx_sys_columns_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_columns` WRITE;
/*!40000 ALTER TABLE `sys_columns` DISABLE KEYS */;

INSERT INTO `sys_columns` (`column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `list`, `pk`, `required`, `super_column`, `usable_column`, `increment`, `insert`, `edit`, `query`, `remark`, `fk_table_name`, `fk_table_name_class`, `fk_table_name_package`, `fk_label_id`, `fk_label_name`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`)
VALUES
	(1,1,'id','流水號','int(11) unsigned','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.504','2023-02-20 12:01:17.504',NULL,0,0),
	(2,1,'title','標題','varchar(128)','string','Title','title','0','','1','1','','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.514','2023-02-20 12:03:25.044',NULL,0,0),
	(3,1,'author','作者','varchar(128)','string','Author','author','0','','1','1','','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.524','2023-02-20 12:03:25.058',NULL,0,0),
	(4,1,'content','内容','varchar(255)','string','Content','content','0','','1','1','','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.530','2023-02-20 12:03:25.067',NULL,0,0),
	(5,1,'status','狀態','int(1)','string','Status','status','0','','1','1','','1','1','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.536','2023-02-20 12:03:25.076',NULL,0,0),
	(6,1,'publish_at','發佈時間','timestamp','time.Time','PublishAt','publishAt','0','','1','1','','1','','EQ','datetime','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.541','2023-02-20 12:03:25.082',NULL,0,0),
	(7,1,'created_at','','timestamp','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.545','2023-02-20 12:01:17.545',NULL,0,0),
	(8,1,'updated_at','','timestamp','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.549','2023-02-20 12:01:17.549',NULL,0,0),
	(9,1,'deleted_at','','timestamp','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.553','2023-02-20 12:01:17.553',NULL,0,0),
	(10,1,'create_by','','int(11) unsigned','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.556','2023-02-20 12:01:17.556',NULL,0,0),
	(11,1,'update_by','','int(11) unsigned','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2023-02-20 12:01:17.559','2023-02-20 12:01:17.559',NULL,0,0);

/*!40000 ALTER TABLE `sys_columns` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_config
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_config`;

CREATE TABLE `sys_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `config_name` varchar(128) DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) DEFAULT NULL COMMENT 'ConfigType',
  `is_frontend` varchar(64) DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(128) DEFAULT NULL COMMENT 'Remark',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  PRIMARY KEY (`id`),
  KEY `idx_sys_config_create_by` (`create_by`),
  KEY `idx_sys_config_update_by` (`update_by`),
  KEY `idx_sys_config_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_config` WRITE;
/*!40000 ALTER TABLE `sys_config` DISABLE KEYS */;

INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'佈景主題','sys_index_skinName','skin-green','Y','0','主頁-預設佈景主題名稱:藍色 skin-blue、綠色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow',1,1,'2021-05-13 19:56:37.913','2021-06-05 13:50:13.123',NULL),
	(2,'初始密碼','sys_user_initPassword','123456','Y','0','User管理-帳號初始密碼:123456',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),
	(3,'侧欄主题','sys_index_sideTheme','theme-dark','Y','0','主頁-側邊欄主题:深色主题theme-dark，淺色主题theme-light',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),
	(4,'系統名稱','sys_app_name','蓋亞業績管理系統','Y','1','',1,0,'2021-03-17 08:52:06.067','2023-02-20 12:00:52.581',NULL),
	(5,'系統logo','sys_app_logo','http://localhost:8000/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','Y','1','',1,0,'2021-03-17 08:53:19.462','2023-02-20 11:22:08.192',NULL);

/*!40000 ALTER TABLE `sys_config` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dept
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dept`;

CREATE TABLE `sys_dept` (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) DEFAULT NULL,
  `dept_path` varchar(255) DEFAULT NULL,
  `dept_name` varchar(128) DEFAULT NULL,
  `sort` tinyint(4) DEFAULT NULL,
  `leader` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  PRIMARY KEY (`dept_id`),
  KEY `idx_sys_dept_create_by` (`create_by`),
  KEY `idx_sys_dept_update_by` (`update_by`),
  KEY `idx_sys_dept_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_dept` WRITE;
/*!40000 ALTER TABLE `sys_dept` DISABLE KEYS */;

INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,0,'/0/1/','蓋亞',0,'o','0912345678','atuo@test.com',2,1,1,'2021-05-13 19:56:37.913','2021-06-05 17:06:44.960',NULL),
	(7,1,'/0/1/7/','研發部',1,'o','0912345678','atuo@test.com',2,1,1,'2021-05-13 19:56:37.913','2021-06-16 21:35:00.109',NULL),
	(8,1,'/0/1/8/','維運部',0,'o','0912345678','atuo@test.com',2,1,1,'2021-05-13 19:56:37.913','2021-06-16 21:41:39.747','2023-02-20 11:32:41.949'),
	(9,1,'/0/1/9/','業務處',0,'修哥','','atuo@test.com',2,1,1,'2021-05-13 19:56:37.913','2023-02-20 11:36:13.520',NULL),
	(10,1,'/0/1/10/','人力資源部',3,'o','0912345678','atuo@test.com',1,1,1,'2021-05-13 19:56:37.913','2021-06-05 17:07:08.503','2023-02-20 11:32:35.747'),
	(11,9,'/0/1/9/11/','業務組A',10,'組長A','','',2,0,0,'2023-02-20 11:36:55.219','2023-02-20 14:20:51.359',NULL),
	(12,9,'/0/1/9/12/','業務組B',10,'組長B','','',2,0,0,'2023-02-20 11:37:45.320','2023-02-20 11:37:45.330',NULL);

/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dict_data
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dict_data`;

CREATE TABLE `sys_dict_data` (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `dict_sort` bigint(20) DEFAULT NULL COMMENT '顯示順序',
  `dict_label` varchar(128) DEFAULT NULL COMMENT 'label',
  `dict_value` varchar(255) DEFAULT NULL COMMENT 'key',
  `dict_type` varchar(64) DEFAULT NULL COMMENT 'type',
  `css_class` varchar(128) DEFAULT NULL,
  `list_class` varchar(128) DEFAULT NULL,
  `is_default` varchar(8) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL COMMENT '狀態',
  `default` varchar(8) DEFAULT NULL COMMENT '預設',
  `remark` varchar(255) DEFAULT NULL COMMENT '備註',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  PRIMARY KEY (`dict_code`),
  KEY `idx_sys_dict_data_create_by` (`create_by`),
  KEY `idx_sys_dict_data_update_by` (`update_by`),
  KEY `idx_sys_dict_data_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_dict_data` WRITE;
/*!40000 ALTER TABLE `sys_dict_data` DISABLE KEYS */;

INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,0,'正常','2','sys_normal_disable','','','',2,'','系統正常',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:40.168',NULL),
	(2,0,'停用','1','sys_normal_disable','','','',2,'','系統停用',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(3,0,'男','0','sys_user_sex','','','',2,'','性别男',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(4,0,'女','1','sys_user_sex','','','',2,'','性别女',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(5,0,'未知','2','sys_user_sex','','','',2,'','性别未知',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(6,0,'顯示','0','sys_show_hide','','','',2,'','顯示選單',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(7,0,'隱藏','1','sys_show_hide','','','',2,'','隱藏選單',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(8,0,'是','Y','sys_yes_no','','','',2,'','系統預設是',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(9,0,'否','N','sys_yes_no','','','',2,'','系統預設否',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(10,0,'正常','2','sys_job_status','','','',2,'','正常狀態',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(11,0,'停用','1','sys_job_status','','','',2,'','停用狀態',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(12,0,'預設','DEFAULT','sys_job_group','','','',2,'','預設分组',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(13,0,'系統','SYSTEM','sys_job_group','','','',2,'','系統分组',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(14,0,'通知','1','sys_notice_type','','','',2,'','通知',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(15,0,'公告','2','sys_notice_type','','','',2,'','公告',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(16,0,'正常','2','sys_common_status','','','',2,'','正常狀態',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(17,0,'關閉','1','sys_common_status','','','',2,'','關閉狀態',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(18,0,'新增','1','sys_oper_type','','','',2,'','新增操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(19,0,'更新','2','sys_oper_type','','','',2,'','更新操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(20,0,'刪除','3','sys_oper_type','','','',2,'','刪除操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(21,0,'授權','4','sys_oper_type','','','',2,'','授權操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(22,0,'匯出','5','sys_oper_type','','','',2,'','匯出操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(23,0,'匯入','6','sys_oper_type','','','',2,'','匯入操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(24,0,'强退','7','sys_oper_type','','','',2,'','强退操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(25,0,'生成代碼','8','sys_oper_type','','','',2,'','生成操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(26,0,'清空資料','9','sys_oper_type','','','',2,'','清空操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(27,0,'成功','0','sys_notice_status','','','',2,'','成功狀態',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(28,0,'失敗','1','sys_notice_status','','','',2,'','失敗狀態',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(29,0,'登入','10','sys_oper_type','','','',2,'','登入操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(30,0,'登出','11','sys_oper_type','','','',2,'','',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(31,0,'取得驗證碼','12','sys_oper_type','','','',2,'','取得驗證碼',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(32,0,'正常','1','sys_content_status','','','',1,'','',1,1,'2021-05-13 19:56:40.845','2021-05-13 19:56:40.845',NULL),
	(33,1,'禁用','2','sys_content_status','','','',1,'','',1,1,'2021-05-13 19:56:40.845','2021-05-13 19:56:40.845',NULL);

/*!40000 ALTER TABLE `sys_dict_data` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dict_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dict_type`;

CREATE TABLE `sys_dict_type` (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `dict_name` varchar(128) DEFAULT NULL COMMENT '字典名稱',
  `dict_type` varchar(128) DEFAULT NULL COMMENT '字典類型',
  `status` tinyint(4) DEFAULT NULL COMMENT '狀態',
  `remark` varchar(255) DEFAULT NULL COMMENT '備註',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  PRIMARY KEY (`dict_id`),
  KEY `idx_sys_dict_type_create_by` (`create_by`),
  KEY `idx_sys_dict_type_update_by` (`update_by`),
  KEY `idx_sys_dict_type_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_dict_type` WRITE;
/*!40000 ALTER TABLE `sys_dict_type` DISABLE KEYS */;

INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'系統開關','sys_normal_disable',2,'系統開關列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(2,'User性别','sys_user_sex',2,'User性别列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(3,'選單狀態','sys_show_hide',2,'選單狀態列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(4,'系統是否','sys_yes_no',2,'系統是否列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(5,'任務狀態','sys_job_status',2,'任務狀態列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(6,'任務分组','sys_job_group',2,'任務分组列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(7,'通知類型','sys_notice_type',2,'通知類型列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(8,'系統狀態','sys_common_status',2,'登入狀態列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(9,'操作類型','sys_oper_type',2,'操作類型列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(10,'通知狀態','sys_notice_status',2,'通知狀態列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
	(11,'内容狀態','sys_content_status',2,'',1,1,'2021-05-13 19:56:40.813','2021-05-13 19:56:40.813',NULL);

/*!40000 ALTER TABLE `sys_dict_type` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_login_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_login_log`;

CREATE TABLE `sys_login_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `username` varchar(128) DEFAULT NULL COMMENT '使用者',
  `status` varchar(4) DEFAULT NULL COMMENT '狀態',
  `ipaddr` varchar(255) DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) DEFAULT NULL COMMENT '歸屬地',
  `browser` varchar(255) DEFAULT NULL COMMENT '瀏覽器',
  `os` varchar(255) DEFAULT NULL COMMENT '系統',
  `platform` varchar(255) DEFAULT NULL COMMENT '平台',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登入時間',
  `remark` varchar(255) DEFAULT NULL COMMENT '備註',
  `msg` varchar(255) DEFAULT NULL COMMENT '訊息',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_login_log_create_by` (`create_by`),
  KEY `idx_sys_login_log_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_login_log` WRITE;
/*!40000 ALTER TABLE `sys_login_log` DISABLE KEYS */;

INSERT INTO `sys_login_log` (`id`, `username`, `status`, `ipaddr`, `login_location`, `browser`, `os`, `platform`, `login_time`, `remark`, `msg`, `created_at`, `updated_at`, `create_by`, `update_by`)
VALUES
	(1,'admin','2','::1','','Chrome 110.0.0.0','Intel Mac OS X 10_15_7','Macintosh','2023-02-20 14:45:10','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','登出成功','2023-02-20 14:45:10.316','2023-02-20 14:45:10.316',0,0),
	(2,'admin','2','::1','','Chrome 110.0.0.0','Intel Mac OS X 10_15_7','Macintosh','2023-02-20 14:45:16','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','登入成功','2023-02-20 14:45:16.530','2023-02-20 14:45:16.530',0,0);

/*!40000 ALTER TABLE `sys_login_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) DEFAULT NULL,
  `title` varchar(128) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `paths` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `permission` varchar(255) DEFAULT NULL,
  `parent_id` smallint(6) DEFAULT NULL,
  `no_cache` tinyint(1) DEFAULT NULL,
  `breadcrumb` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `sort` tinyint(4) DEFAULT NULL,
  `visible` varchar(1) DEFAULT NULL,
  `is_frame` varchar(1) DEFAULT '0',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  PRIMARY KEY (`menu_id`),
  KEY `idx_sys_menu_create_by` (`create_by`),
  KEY `idx_sys_menu_update_by` (`update_by`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;

INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `title`, `icon`, `path`, `paths`, `menu_type`, `action`, `permission`, `parent_id`, `no_cache`, `breadcrumb`, `component`, `sort`, `visible`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(2,'Admin','系統管理','api-server','/admin','/0/2','M','無','',0,1,'','Layout',10,'0','1',0,1,'2021-05-20 21:58:45.679','2023-02-20 11:08:32.291',NULL),
	(3,'SysUserManage','User管理','user','/admin/sys-user','/0/2/3','C','無','admin:sysUser:list',2,0,'','/admin/sys-user/index',10,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.323',NULL),
	(43,'','新增管理員','app-group-fill','','/0/2/3/43','F','POST','admin:sysUser:add',3,0,'','',10,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.472',NULL),
	(44,'','查詢管理員','app-group-fill','','/0/2/3/44','F','GET','admin:sysUser:query',3,0,'','',40,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.491',NULL),
	(45,'','更新管理員','app-group-fill','','/0/2/3/45','F','PUT','admin:sysUser:edit',3,0,'','',30,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.509',NULL),
	(46,'','刪除管理員','app-group-fill','','/0/2/3/46','F','DELETE','admin:sysUser:remove',3,0,'','',20,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.520',NULL),
	(51,'SysMenuManage','選單管理','tree-table','/admin/sys-menu','/0/2/51','C','無','admin:sysMenu:list',2,1,'','/admin/sys-menu/index',30,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.336',NULL),
	(52,'SysRoleManage','角色管理','peoples','/admin/sys-role','/0/2/52','C','無','admin:sysRole:list',2,1,'','/admin/sys-role/index',20,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.347',NULL),
	(56,'SysDeptManage','部門管理','tree','/admin/sys-dept','/0/2/56','C','無','admin:sysDept:list',2,0,'','/admin/sys-dept/index',40,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.358',NULL),
	(57,'SysPostManage','職位管理','pass','/admin/sys-post','/0/2/57','C','無','admin:sysPost:list',2,0,'','/admin/sys-post/index',50,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.370',NULL),
	(58,'Dict','字典管理','education','/admin/dict','/0/2/58','C','無','admin:sysDictType:list',2,0,'','/admin/dict/index',60,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.389',NULL),
	(59,'SysDictDataManage','字典資料','education','/admin/dict/data/:dictId','/0/2/59','C','無','admin:sysDictData:list',2,0,'','/admin/dict/data',100,'1','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.402',NULL),
	(60,'Tools','開發工具','dev-tools','/dev-tools','/0/60','M','無','',0,0,'','Layout',40,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.303',NULL),
	(61,'Swagger','系統API','guide','/dev-tools/swagger','/0/60/61','C','無','',60,0,'','/dev-tools/swagger/index',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.832',NULL),
	(62,'SysConfigManage','系統參數','swagger','/admin/sys-config','/0/2/62','C','無','admin:sysConfig:list',2,0,'','/admin/sys-config/index',70,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.419',NULL),
	(211,'Log','Log管理','log','/log','/0/2/211','M','','',2,0,'','/log/index',80,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.430',NULL),
	(212,'SysLoginLogManage','登入Log','logininfor','/admin/sys-login-log','/0/2/211/212','C','','admin:sysLoginLog:list',211,0,'','/admin/sys-login-log/index',1,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.894',NULL),
	(216,'OperLog','操作Log','skill','/admin/sys-oper-log','/0/2/211/216','C','','admin:sysOperLog:list',211,0,'','/admin/sys-oper-log/index',1,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.903',NULL),
	(220,'','新增選單','app-group-fill','','/0/2/51/220','F','','admin:sysMenu:add',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.555',NULL),
	(221,'','更新選單','app-group-fill','','/0/2/51/221','F','','admin:sysMenu:edit',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.574',NULL),
	(222,'','查詢選單','app-group-fill','','/0/2/51/222','F','','admin:sysMenu:query',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.587',NULL),
	(223,'','刪除選單','app-group-fill','','/0/2/51/223','F','','admin:sysMenu:remove',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.604',NULL),
	(224,'','新增角色','app-group-fill','','/0/2/52/224','F','','admin:sysRole:add',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.616',NULL),
	(225,'','查詢角色','app-group-fill','','/0/2/52/225','F','','admin:sysRole:query',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.628',NULL),
	(226,'','更新角色','app-group-fill','','/0/2/52/226','F','','admin:sysRole:update',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.645',NULL),
	(227,'','刪除角色','app-group-fill','','/0/2/52/227','F','','admin:sysRole:remove',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.656',NULL),
	(228,'','查詢部門','app-group-fill','','/0/2/56/228','F','','admin:sysDept:query',56,0,'','',40,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.678',NULL),
	(229,'','新增部門','app-group-fill','','/0/2/56/229','F','','admin:sysDept:add',56,0,'','',10,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.693',NULL),
	(230,'','更新部門','app-group-fill','','/0/2/56/230','F','','admin:sysDept:edit',56,0,'','',30,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.706',NULL),
	(231,'','刪除部門','app-group-fill','','/0/2/56/231','F','','admin:sysDept:remove',56,0,'','',20,'0','1',0,1,'2021-05-20 22:08:44.526','2023-02-20 11:08:32.715',NULL),
	(232,'','查詢職位','app-group-fill','','/0/2/57/232','F','','admin:sysPost:query',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.725',NULL),
	(233,'','新增職位','app-group-fill','','/0/2/57/233','F','','admin:sysPost:add',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.734',NULL),
	(234,'','更新職位','app-group-fill','','/0/2/57/234','F','','admin:sysPost:edit',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.741',NULL),
	(235,'','刪除職位','app-group-fill','','/0/2/57/235','F','','admin:sysPost:remove',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.750',NULL),
	(236,'','查詢字典','app-group-fill','','/0/2/58/236','F','','admin:sysDictType:query',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.757',NULL),
	(237,'','新增類型','app-group-fill','','/0/2/58/237','F','','admin:sysDictType:add',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.765',NULL),
	(238,'','更新類型','app-group-fill','','/0/2/58/238','F','','admin:sysDictType:edit',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.774',NULL),
	(239,'','刪除類型','app-group-fill','','/0/2/58/239','F','','system:sysdicttype:remove',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.782',NULL),
	(240,'','查詢資料','app-group-fill','','/0/2/59/240','F','','admin:sysDictData:query',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.790',NULL),
	(241,'','新增資料','app-group-fill','','/0/2/59/241','F','','admin:sysDictData:add',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.799',NULL),
	(242,'','更新資料','app-group-fill','','/0/2/59/242','F','','admin:sysDictData:edit',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.806',NULL),
	(243,'','刪除資料','app-group-fill','','/0/2/59/243','F','','admin:sysDictData:remove',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.821',NULL),
	(244,'','查詢參數','app-group-fill','','/0/2/62/244','F','','admin:sysConfig:query',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.865',NULL),
	(245,'','新增參數','app-group-fill','','/0/2/62/245','F','','admin:sysConfig:add',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.872',NULL),
	(246,'','更新參數','app-group-fill','','/0/2/62/246','F','','admin:sysConfig:edit',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.880',NULL),
	(247,'','刪除參數','app-group-fill','','/0/2/62/247','F','','admin:sysConfig:remove',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.887',NULL),
	(248,'','查詢登入Log','app-group-fill','','/0/2/211/212/248','F','','admin:sysLoginLog:query',212,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.909',NULL),
	(249,'','刪除登入Log','app-group-fill','','/0/2/211/212/249','F','','admin:sysLoginLog:remove',212,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.917',NULL),
	(250,'','查詢操作Log','app-group-fill','','/0/2/211/216/250','F','','admin:sysOperLog:query',216,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.925',NULL),
	(251,'','刪除操作Log','app-group-fill','','/0/2/211/216/251','F','','admin:sysOperLog:remove',216,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.933',NULL),
	(261,'Gen','代碼生成','code','/dev-tools/gen','/0/60/261','C','','',60,0,'','/dev-tools/gen/index',2,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.840',NULL),
	(262,'EditTable','代碼生成更新','build','/dev-tools/editTable','/0/60/262','C','','',60,0,'','/dev-tools/gen/editTable',100,'1','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.847',NULL),
	(264,'Build','表單建立','build','/dev-tools/build','/0/60/264','C','','',60,0,'','/dev-tools/build/index',1,'0','1',1,1,'2020-04-11 15:52:48.000','2023-02-20 11:08:32.857',NULL),
	(528,'SysApiManage','API管理','api-doc','/admin/sys-api','/0/2/528','C','無','admin:sysApi:list',2,0,'','/admin/sys-api/index',0,'1','0',0,1,'2021-05-20 22:08:44.526','2023-02-20 12:48:41.934',NULL),
	(529,'','查詢API','app-group-fill','','/0/2/528/529','F','無','admin:sysApi:query',528,0,'','',40,'0','0',0,1,'2021-05-20 22:08:44.526','2023-02-20 12:48:41.937',NULL),
	(531,'','更新API','app-group-fill','','/0/2/528/531','F','無','admin:sysApi:edit',528,0,'','',30,'0','0',0,1,'2021-05-20 22:08:44.526','2023-02-20 12:48:41.941',NULL),
	(537,'SysTools','系統工具','system-tools','/sys-tools','/0/537','M','','',0,0,'','Layout',30,'1','1',1,1,'2021-05-21 11:13:32.166','2023-02-20 11:21:44.801',NULL),
	(540,'SysConfigSet','系統設定','system-tools','/admin/sys-config/set','/0/2/540','C','','admin:sysConfigSet:list',2,0,'','/admin/sys-config/set',0,'0','1',1,1,'2021-05-25 16:06:52.560','2023-02-20 11:08:32.459',NULL),
	(542,'','更新','upload','','/0/2/540/542','F','','admin:sysConfigSet:update',540,0,'','',0,'0','1',1,1,'2021-06-13 11:45:48.670','2023-02-20 11:08:32.955',NULL),
	(543,'','內容','pass','/article','/0/543','M','無','',0,0,'','Layout',0,'0','0',1,0,'2023-02-20 12:10:58.738','2023-02-20 12:20:05.449',NULL),
	(544,'ArticleManage','文章管理','pass','/admin/article','/0/543/544','C','無','admin:article:list',543,0,'','/admin/article/index',0,'0','0',1,1,'2023-02-20 12:10:58.762','2023-02-20 12:20:05.449',NULL),
	(545,'','分頁取得內容','','article','/0/543/544/545','F','無','admin:article:query',544,0,'','',0,'0','0',1,1,'2023-02-20 12:10:58.778','2023-02-20 12:20:05.449',NULL),
	(546,'','建立內容','','article','/0/543/544/546','F','無','admin:article:add',544,0,'','',0,'0','0',1,1,'2023-02-20 12:10:58.790','2023-02-20 12:20:05.449',NULL),
	(547,'','更新內容','','article','/0/543/544/547','F','無','admin:article:edit',544,0,'','',0,'0','0',1,1,'2023-02-20 12:10:58.800','2023-02-20 12:20:05.449',NULL),
	(548,'','刪除內容','','article','/0/543/544/548','F','無','admin:article:remove',544,0,'','',0,'0','0',1,1,'2023-02-20 12:10:58.807','2023-02-20 12:20:05.449',NULL);

/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_menu_api_rule
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_menu_api_rule`;

CREATE TABLE `sys_menu_api_rule` (
  `sys_menu_menu_id` bigint(20) NOT NULL,
  `sys_api_id` bigint(20) NOT NULL COMMENT '流水號',
  PRIMARY KEY (`sys_menu_menu_id`,`sys_api_id`),
  KEY `fk_sys_menu_api_rule_sys_api` (`sys_api_id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_menu_api_rule` WRITE;
/*!40000 ALTER TABLE `sys_menu_api_rule` DISABLE KEYS */;

INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`)
VALUES
	(3,141),
	(43,150),
	(44,141),
	(45,142),
	(45,151),
	(46,156),
	(51,39),
	(51,135),
	(52,44),
	(56,27),
	(57,59),
	(58,21),
	(58,26),
	(59,24),
	(62,53),
	(212,137),
	(216,6),
	(220,88),
	(221,41),
	(221,109),
	(222,39),
	(223,126),
	(224,90),
	(225,44),
	(226,29),
	(226,45),
	(226,46),
	(226,47),
	(226,106),
	(226,107),
	(227,128),
	(228,27),
	(229,82),
	(230,28),
	(230,103),
	(231,122),
	(232,59),
	(233,89),
	(234,60),
	(234,110),
	(235,127),
	(236,21),
	(236,26),
	(237,81),
	(238,23),
	(238,102),
	(239,121),
	(240,24),
	(241,80),
	(242,25),
	(242,101),
	(243,120),
	(244,53),
	(245,87),
	(246,54),
	(246,108),
	(247,125),
	(248,137),
	(249,114),
	(250,6),
	(251,115),
	(528,135),
	(529,135),
	(531,92),
	(531,136),
	(540,140),
	(542,139),
	(544,161),
	(544,162),
	(544,164),
	(544,166),
	(544,168);

/*!40000 ALTER TABLE `sys_menu_api_rule` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_migration
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_migration`;

CREATE TABLE `sys_migration` (
  `version` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `apply_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_migration` WRITE;
/*!40000 ALTER TABLE `sys_migration` DISABLE KEYS */;

INSERT INTO `sys_migration` (`version`, `apply_time`)
VALUES
	('1599190683659','2023-02-20 11:08:32.265'),
	('1653638869132','2023-02-20 11:08:32.959');

/*!40000 ALTER TABLE `sys_migration` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_opera_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_opera_log`;

CREATE TABLE `sys_opera_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `title` varchar(255) DEFAULT NULL COMMENT '操作模組',
  `business_type` varchar(128) DEFAULT NULL COMMENT '操作類型',
  `business_types` varchar(128) DEFAULT NULL COMMENT 'BusinessTypes',
  `method` varchar(128) DEFAULT NULL COMMENT '函數',
  `request_method` varchar(128) DEFAULT NULL COMMENT '請求方式: GET POST PUT DELETE',
  `operator_type` varchar(128) DEFAULT NULL COMMENT '操作類型',
  `oper_name` varchar(128) DEFAULT NULL COMMENT '操作者',
  `dept_name` varchar(128) DEFAULT NULL COMMENT '部門名稱',
  `oper_url` varchar(255) DEFAULT NULL COMMENT '訪問位置',
  `oper_ip` varchar(128) DEFAULT NULL COMMENT 'client ip',
  `oper_location` varchar(128) DEFAULT NULL COMMENT '訪問位置',
  `oper_param` text DEFAULT NULL COMMENT '請求參數',
  `status` varchar(4) DEFAULT NULL COMMENT '操作狀態 1:正常 2:關閉',
  `oper_time` timestamp NULL DEFAULT NULL COMMENT '操作時間',
  `json_result` varchar(255) DEFAULT NULL COMMENT 'response',
  `remark` varchar(255) DEFAULT NULL COMMENT '備註',
  `latency_time` varchar(128) DEFAULT NULL COMMENT 'response time',
  `user_agent` varchar(255) DEFAULT NULL COMMENT 'ua',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_opera_log_create_by` (`create_by`),
  KEY `idx_sys_opera_log_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_opera_log` WRITE;
/*!40000 ALTER TABLE `sys_opera_log` DISABLE KEYS */;

INSERT INTO `sys_opera_log` (`id`, `title`, `business_type`, `business_types`, `method`, `request_method`, `operator_type`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`)
VALUES
	(1,'','','','','GET','','','','/api/v1/captcha','::1','','','2','2023-02-20 14:45:10','{\"code\":200,\"data\":\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAPAAAABQCAMAAAAQlwhOAAAA81BMVEUAAA','','3.402417ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:10.800','2023-02-20 14:45:10.800',0,0),
	(2,'','','','','GET','','','','/api/v1/app-config','::1','','','1','2023-02-20 14:45:10','{\"requestId\":\"c7c8958c-0de1-4f1a-99d5-ae15606e66b5\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"sys_app','','11.273792ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:10.819','2023-02-20 14:45:10.819',0,0),
	(3,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:45:10','','','510.958µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:10.865','2023-02-20 14:45:10.865',0,0),
	(4,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:45:16','{\"requestId\":\"18c912f9-3946-4b60-93e1-cb1e13c2251c\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','5.276833ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:16.543','2023-02-20 14:45:16.543',1,1),
	(5,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:45:16','{\"requestId\":\"e7987ee9-1203-4c0c-98e5-63697242bce7\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','4.143292ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:16.551','2023-02-20 14:45:16.551',1,1),
	(6,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:45:16','','','260.959µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:16.732','2023-02-20 14:45:16.732',0,0),
	(7,'','','','','GET','','admin','','/api/v1/roleMenuTreeselect/0','::1','','','1','2023-02-20 14:45:16','{\"requestId\":\"ca935130-c4ed-42ea-a7ac-d891950c42f7\",\"code\":200,\"msg\":\"取得成功\",\"data\":{\"checked','','6.65025ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:16.810','2023-02-20 14:45:16.810',1,1),
	(8,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_normal_disable','::1','','','1','2023-02-20 14:45:16','{\"requestId\":\"f82fedbb-ad28-4952-b771-178bcbba4f9c\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','13.671958ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:16.824','2023-02-20 14:45:16.824',1,1),
	(9,'','','','','GET','','admin','','/api/v1/role?pageIndex=1&pageSize=10','::1','','','1','2023-02-20 14:45:16','{\"requestId\":\"c0da1e08-ffcd-4670-8fdb-48dcac575784\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','31.3655ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:16.831','2023-02-20 14:45:16.831',1,1),
	(10,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:45:19','{\"requestId\":\"f8bae015-8656-413e-8cca-a64b36797351\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','4.962583ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:19.503','2023-02-20 14:45:19.503',1,1),
	(11,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:45:19','{\"requestId\":\"e3bb60f4-f28a-4323-9bc1-b3ae658e89ea\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','7.672917ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:19.523','2023-02-20 14:45:19.523',1,1),
	(12,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:45:19','','','378.25µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:19.720','2023-02-20 14:45:19.720',0,0),
	(13,'','','','','GET','','admin','','/api/v1/roleMenuTreeselect/0','::1','','','1','2023-02-20 14:45:19','{\"requestId\":\"adb96321-f6eb-4412-ab1f-26d45a221682\",\"code\":200,\"msg\":\"取得成功\",\"data\":{\"checked','','31.774959ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:19.749','2023-02-20 14:45:19.749',1,1),
	(14,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_normal_disable','::1','','','1','2023-02-20 14:45:19','{\"requestId\":\"83991815-7bba-4a10-aa40-26ffe8b54ce7\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','36.14875ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:19.761','2023-02-20 14:45:19.761',1,1),
	(15,'','','','','GET','','admin','','/api/v1/role?pageIndex=1&pageSize=10','::1','','','1','2023-02-20 14:45:19','{\"requestId\":\"1dede46e-64ef-4c23-8a5b-2571097cecda\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','36.85425ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:19.772','2023-02-20 14:45:19.772',1,1),
	(16,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:45:42','{\"requestId\":\"c04e9905-0a45-4d21-8045-564f7787c300\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','10.98775ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:42.763','2023-02-20 14:45:42.763',1,1),
	(17,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:45:47','{\"requestId\":\"753b99b8-4b1b-4e25-a482-ee4795a6be18\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','5.888792ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:47.037','2023-02-20 14:45:47.037',1,1),
	(18,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:45:47','{\"requestId\":\"48cba042-751d-45a0-8d84-1abbc458381a\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','8.24025ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:45:47.046','2023-02-20 14:45:47.046',1,1),
	(19,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:47:35','{\"requestId\":\"f089cb4f-0fc1-4f0a-824e-8c08c5107752\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','14.038958ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:47:35.036','2023-02-20 14:47:35.036',1,1),
	(20,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:47:40','{\"requestId\":\"3bf743a8-5254-4de5-ba6f-52c73210f6bc\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','7.422458ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:47:40.432','2023-02-20 14:47:40.432',1,1),
	(21,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:47:40','{\"requestId\":\"1161f7ae-b04a-4fc7-8024-de4df777d24f\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','11.210083ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:47:40.443','2023-02-20 14:47:40.443',1,1),
	(22,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:47:51','{\"requestId\":\"43c4d32e-26db-49ec-9302-4e6c1602444a\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','12.015958ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:47:51.280','2023-02-20 14:47:51.280',1,1),
	(23,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:48:43','{\"requestId\":\"62f27c89-b149-4555-bac6-d698ee68a86b\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','11.42875ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:48:43.899','2023-02-20 14:48:43.899',1,1),
	(24,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:48:43','{\"requestId\":\"d075a0f8-d501-42f7-8954-92250ab4cede\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','4.698875ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:48:43.951','2023-02-20 14:48:43.951',1,1),
	(25,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:48:44','','','1.100417ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:48:44.213','2023-02-20 14:48:44.213',0,0),
	(26,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:48:44','{\"requestId\":\"bf36bc70-ec00-42f6-adcf-c4f9b4cab3e2\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','17.551291ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:48:44.233','2023-02-20 14:48:44.233',1,1),
	(27,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:49:14','{\"requestId\":\"bdf32a9b-7a8a-4618-9103-f7b5e0b0073c\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','11.559375ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:49:14.882','2023-02-20 14:49:14.882',1,1),
	(28,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:49:34','{\"requestId\":\"67f24a50-eaf7-4b7a-8bb3-3f9a4e41d50f\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','14.925792ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:49:34.332','2023-02-20 14:49:34.332',1,1),
	(29,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:49:34','{\"requestId\":\"63633a67-829b-4bf1-b667-7beb56efa2b2\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','5.491333ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:49:34.353','2023-02-20 14:49:34.353',1,1),
	(30,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:49:34','','','463.458µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:49:34.670','2023-02-20 14:49:34.670',0,0),
	(31,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:49:34','{\"requestId\":\"8581337c-ce0d-4626-abc9-38310ba9d56e\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','21.12925ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:49:34.692','2023-02-20 14:49:34.692',1,1),
	(32,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:50:01','{\"requestId\":\"e9927014-1aba-4183-af99-7b8668ef05dc\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','5.588875ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:01.134','2023-02-20 14:50:01.134',1,1),
	(33,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:50:01','{\"requestId\":\"2121d393-4a59-4b60-b9c6-f38a25a5bb0e\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','7.139667ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:01.153','2023-02-20 14:50:01.153',1,1),
	(34,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:50:01','','','502.916µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:01.361','2023-02-20 14:50:01.361',0,0),
	(35,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:50:01','{\"requestId\":\"db60d75f-5800-4b8a-b317-b50822e3afa7\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','9.726208ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:01.437','2023-02-20 14:50:01.437',1,1),
	(36,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:50:25','{\"requestId\":\"03657f74-6fc8-4b77-81fb-226e1d1a3810\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','41.950917ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:25.248','2023-02-20 14:50:25.248',1,1),
	(37,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:50:25','{\"requestId\":\"20b6fdc0-d477-4e1b-89b6-ca0bb462990a\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','3.917791ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:25.302','2023-02-20 14:50:25.302',1,1),
	(38,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:50:25','','','343.208µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:25.521','2023-02-20 14:50:25.521',0,0),
	(39,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:50:25','{\"requestId\":\"014ad131-c057-4954-a881-3119c23af9a8\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','25.629291ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:25.550','2023-02-20 14:50:25.550',1,1),
	(40,'','','','','GET','','admin','','/api/v1/config?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:50:53','{\"requestId\":\"6559f6cb-febb-4064-a8e2-1775c7c83ce7\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','14.815083ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:53.385','2023-02-20 14:50:53.385',1,1),
	(41,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_yes_no','::1','','','1','2023-02-20 14:50:53','{\"requestId\":\"4ed6f423-7cda-4f17-bd8a-958b5ac55742\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','18.515125ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:53.398','2023-02-20 14:50:53.398',1,1),
	(42,'','','','','GET','','admin','','/api/v1/dict/type?pageIndex=1&pageSize=10','::1','','','1','2023-02-20 14:50:55','{\"requestId\":\"d9ac3b10-bed5-4004-9fde-f7b2f195775d\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','5.829709ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:55.678','2023-02-20 14:50:55.678',1,1),
	(43,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_normal_disable','::1','','','1','2023-02-20 14:50:55','{\"requestId\":\"36c058f1-e8dd-4cf3-b0f8-1530f2d5de3f\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','6.937084ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:55.688','2023-02-20 14:50:55.688',1,1),
	(44,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:50:58','{\"requestId\":\"d52a8aa0-48bb-46fa-9ee7-47a1510ad3ce\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','6.075208ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:58.685','2023-02-20 14:50:58.685',1,1),
	(45,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:50:58','{\"requestId\":\"4d1441de-3445-40fb-9591-4ee0b3140e52\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','7.180333ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:50:58.695','2023-02-20 14:50:58.695',1,1),
	(46,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:52:04','{\"requestId\":\"5fc0d620-1696-47c5-afcf-f2b83f1bda3c\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','45.320125ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:52:04.909','2023-02-20 14:52:04.909',1,1),
	(47,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:52:04','{\"requestId\":\"8c406e33-9013-48e5-a84e-b0f8850b4942\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','45.830042ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:52:04.924','2023-02-20 14:52:04.924',1,1),
	(48,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:52:17','{\"requestId\":\"d10db5a4-e8bf-486a-9c74-83666c451b08\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','8.4715ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:52:17.061','2023-02-20 14:52:17.061',1,1),
	(49,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:52:17','{\"requestId\":\"c28b10d4-d7ab-4faf-aca8-7a09b8987bba\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','10.613333ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:52:17.073','2023-02-20 14:52:17.073',1,1),
	(50,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:52:30','{\"requestId\":\"c270abde-4ee1-4311-bb07-40dd3dc6ca41\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','18.974625ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:52:30.901','2023-02-20 14:52:30.901',1,1),
	(51,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:52:30','{\"requestId\":\"f9d84f33-bd51-4b94-8d0c-ee3942d760f5\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','26.58625ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:52:30.919','2023-02-20 14:52:30.919',1,1),
	(52,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:53:32','{\"requestId\":\"ad18a430-018e-4b4c-a140-768fad76a33a\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','9.209333ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:32.178','2023-02-20 14:53:32.178',1,1),
	(53,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:53:32','{\"requestId\":\"661026f5-9a3e-4a78-a14a-e51b32d023a0\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','5.522291ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:32.198','2023-02-20 14:53:32.198',1,1),
	(54,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:53:32','','','402.75µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:32.425','2023-02-20 14:53:32.425',0,0),
	(55,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:53:32','{\"requestId\":\"7c8695bd-c26d-4c79-9880-d7e4f4c9d2a1\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','12.715375ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:32.450','2023-02-20 14:53:32.450',1,1),
	(56,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:53:32','{\"requestId\":\"2daadb00-01e4-4cf5-ac38-0b4ff97bbb30\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','21.088208ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:32.460','2023-02-20 14:53:32.460',1,1),
	(57,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:53:52','{\"requestId\":\"3c5cb48a-b507-4480-b7b2-8186b5f5070c\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','7.522292ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:52.153','2023-02-20 14:53:52.153',1,1),
	(58,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:53:52','{\"requestId\":\"ac53a0be-3423-4f02-b722-0d02260f19d6\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','6.029917ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:52.172','2023-02-20 14:53:52.172',1,1),
	(59,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:53:52','','','472.542µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:52.421','2023-02-20 14:53:52.421',0,0),
	(60,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:53:52','{\"requestId\":\"9aa87929-8b0a-46c5-8d81-4883ce1a4d63\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','11.170125ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:52.441','2023-02-20 14:53:52.441',1,1),
	(61,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:53:52','{\"requestId\":\"608180f0-afb2-4ccb-81f9-9531bc179258\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','24.00925ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:53:52.453','2023-02-20 14:53:52.453',1,1),
	(62,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:54:02','{\"requestId\":\"2e63f80f-4e55-4e4a-bbea-103a91354078\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','7.824667ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:54:02.081','2023-02-20 14:54:02.081',1,1),
	(63,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:54:02','{\"requestId\":\"250d517d-6c2b-4732-ad09-8fffdde214fc\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','13.036958ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:54:02.114','2023-02-20 14:54:02.114',1,1),
	(64,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:54:02','','','382.5µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:54:02.310','2023-02-20 14:54:02.310',0,0),
	(65,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:54:02','{\"requestId\":\"37e18af7-cf50-40f1-9a75-f682c66841c4\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','11.075125ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:54:02.328','2023-02-20 14:54:02.328',1,1),
	(66,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:54:02','{\"requestId\":\"10248b0f-ed3d-4587-854d-5037e8278925\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','20.664542ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:54:02.340','2023-02-20 14:54:02.340',1,1),
	(67,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:54:20','{\"requestId\":\"6b4abce4-b58e-4908-aa95-e123865f50d5\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','17.279625ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:54:20.551','2023-02-20 14:54:20.551',1,1),
	(68,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:54:20','{\"requestId\":\"c0175705-40f3-4fe2-894f-a3c33fba7e1f\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','22.491583ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:54:20.557','2023-02-20 14:54:20.557',1,1),
	(69,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:55:44','{\"requestId\":\"6cab62c7-96b2-411b-b95d-fa89c0519cea\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','24.677083ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:44.885','2023-02-20 14:55:44.885',1,1),
	(70,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:55:44','{\"requestId\":\"4a31c13a-1967-449a-91ed-4d91aa241be5\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','6.867583ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:44.904','2023-02-20 14:55:44.904',1,1),
	(71,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:55:45','','','685µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:45.205','2023-02-20 14:55:45.205',0,0),
	(72,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:55:45','{\"requestId\":\"ffe687b0-0e02-4751-b13d-9e19384ea65a\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','8.748833ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:45.276','2023-02-20 14:55:45.276',1,1),
	(73,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:55:45','{\"requestId\":\"27025ba1-b0ed-445a-b010-5433f372a1b9\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','10.174125ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:45.300','2023-02-20 14:55:45.300',1,1),
	(74,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 14:55:58','{\"requestId\":\"8d2f0878-e343-40c6-824f-3c1f38b5e6cc\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','14.438042ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:58.976','2023-02-20 14:55:58.976',1,1),
	(75,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 14:55:59','{\"requestId\":\"68c4132c-46b8-4714-938c-e85211c96125\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','3.478708ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:59.004','2023-02-20 14:55:59.004',1,1),
	(76,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 14:55:59','','','433.542µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:59.159','2023-02-20 14:55:59.159',0,0),
	(77,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 14:55:59','{\"requestId\":\"2c803b23-dc72-48e8-92d5-ba031a5f2637\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','9.540792ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:59.224','2023-02-20 14:55:59.224',1,1),
	(78,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 14:55:59','{\"requestId\":\"430628c0-edee-4e3b-9579-dff13a05da30\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','17.457208ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 14:55:59.236','2023-02-20 14:55:59.236',1,1),
	(79,'','','','','GET','','admin','','/api/v1/getinfo','::1','','','1','2023-02-20 15:00:16','{\"requestId\":\"58006703-9084-4941-95ab-c23b0b3bf1f0\",\"code\":200,\"data\":{\"avatar\":\"https://wpimg.walls','','38.031709ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 15:00:16.624','2023-02-20 15:00:16.624',1,1),
	(80,'','','','','GET','','admin','','/api/v1/menurole','::1','','','1','2023-02-20 15:00:16','{\"requestId\":\"82e9a265-eca4-429a-9720-752e7033fb2b\",\"code\":200,\"data\":[{\"menuId\":543,\"menuName\":\"\",\"','','3.864209ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 15:00:16.655','2023-02-20 15:00:16.655',1,1),
	(81,'','','','','GET','','','','/static/uploadfile/89f18216-0325-4246-a9e3-fab66e454c4d.png','::1','','','2','2023-02-20 15:00:16','','','529.792µs','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 15:00:16.911','2023-02-20 15:00:16.911',0,0),
	(82,'','','','','GET','','admin','','/api/v1/sys-opera-log?pageIndex=1&pageSize=10&createdAtOrder=desc','::1','','','1','2023-02-20 15:00:16','{\"requestId\":\"976521db-9334-476d-8fbd-a537c105f3f6\",\"code\":200,\"msg\":\"查詢成功\",\"data\":{\"count\":','','14.823666ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 15:00:16.919','2023-02-20 15:00:16.919',1,1),
	(83,'','','','','GET','','admin','','/api/v1/dict-data/option-select?dictType=sys_common_status','::1','','','1','2023-02-20 15:00:16','{\"requestId\":\"c9ded827-d85b-4ede-92b7-6a930bb917d4\",\"code\":200,\"msg\":\"查詢成功\",\"data\":[{\"label\"','','24.05175ms','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36','2023-02-20 15:00:16.932','2023-02-20 15:00:16.932',1,1);

/*!40000 ALTER TABLE `sys_opera_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_post`;

CREATE TABLE `sys_post` (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) DEFAULT NULL,
  `post_code` varchar(128) DEFAULT NULL,
  `sort` tinyint(4) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  PRIMARY KEY (`post_id`),
  KEY `idx_sys_post_create_by` (`create_by`),
  KEY `idx_sys_post_update_by` (`update_by`),
  KEY `idx_sys_post_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_post` WRITE;
/*!40000 ALTER TABLE `sys_post` DISABLE KEYS */;

INSERT INTO `sys_post` (`post_id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'業務組長','ST',2,2,'',1,1,'2021-05-13 19:56:37.913','2023-02-20 11:31:46.094',NULL),
	(2,'業務助理','SA',3,2,'',1,1,'2021-05-13 19:56:37.913','2023-02-20 11:31:55.960',NULL),
	(3,'業務組員','SM',4,2,'',1,1,'2021-05-13 19:56:37.913','2023-02-20 11:31:59.942',NULL),
	(4,'業務處長','SD',1,2,'',1,1,'2023-02-20 11:31:14.390','2023-02-20 11:31:33.959',NULL);

/*!40000 ALTER TABLE `sys_post` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色流水號',
  `role_name` varchar(128) DEFAULT NULL COMMENT '角色名稱',
  `status` varchar(4) DEFAULT NULL,
  `role_key` varchar(128) DEFAULT NULL COMMENT '角色代碼',
  `role_sort` bigint(20) DEFAULT NULL COMMENT '角色排序',
  `flag` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL COMMENT '備註',
  `admin` tinyint(1) DEFAULT NULL,
  `data_scope` varchar(128) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  PRIMARY KEY (`role_id`),
  KEY `idx_sys_role_create_by` (`create_by`),
  KEY `idx_sys_role_update_by` (`update_by`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;

INSERT INTO `sys_role` (`role_id`, `role_name`, `status`, `role_key`, `role_sort`, `flag`, `remark`, `admin`, `data_scope`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'系統管理員','2','admin',1,'','',1,'1',1,1,'2021-05-13 19:56:37.913','2023-02-20 11:42:25.964',NULL),
	(2,'業務組長','2','sales-team-lead',0,'','',0,'4',0,0,'2023-02-20 11:43:12.706','2023-02-20 12:18:23.372',NULL),
	(3,'系統組員','2','sales-team-member',0,'','',0,'5',0,0,'2023-02-20 11:50:09.758','2023-02-20 12:20:05.446',NULL),
	(4,'業務處長','2','sales-director',0,'','',0,'1',0,0,'2023-02-20 11:51:31.663','2023-02-20 12:13:09.999',NULL),
	(5,'業務助理','2','sales-assistant',0,'','',0,'3',0,0,'2023-02-20 11:52:18.151','2023-02-20 12:13:13.667',NULL),
	(6,'業務特殊角色','2','SS',0,'','',0,'2',0,0,'2023-02-20 14:20:35.071','2023-02-20 14:20:51.355',NULL);

/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role_dept
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role_dept`;

CREATE TABLE `sys_role_dept` (
  `role_id` smallint(6) NOT NULL,
  `dept_id` smallint(6) NOT NULL,
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_role_dept` WRITE;
/*!40000 ALTER TABLE `sys_role_dept` DISABLE KEYS */;

INSERT INTO `sys_role_dept` (`role_id`, `dept_id`)
VALUES
	(6,11);

/*!40000 ALTER TABLE `sys_role_dept` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role_menu`;

CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL COMMENT '角色流水號',
  `menu_id` bigint(20) NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`),
  KEY `fk_sys_role_menu_sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;

INSERT INTO `sys_role_menu` (`role_id`, `menu_id`)
VALUES
	(2,528),
	(2,529),
	(2,531),
	(2,543),
	(2,544),
	(2,545),
	(2,546),
	(2,547),
	(2,548),
	(3,528),
	(3,529),
	(3,531),
	(3,543),
	(3,544),
	(3,545),
	(3,546),
	(3,547),
	(3,548),
	(4,543),
	(4,544),
	(4,545),
	(4,546),
	(4,547),
	(4,548),
	(5,543),
	(5,544),
	(5,545),
	(5,546),
	(5,547),
	(5,548);

/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_tables
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_tables`;

CREATE TABLE `sys_tables` (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '表流水號',
  `table_name` varchar(255) DEFAULT NULL COMMENT '表名稱',
  `table_comment` varchar(255) DEFAULT NULL COMMENT '表備註',
  `class_name` varchar(255) DEFAULT NULL COMMENT '類名',
  `tpl_category` varchar(255) DEFAULT NULL COMMENT '包名',
  `package_name` varchar(255) DEFAULT NULL,
  `module_name` varchar(255) DEFAULT NULL COMMENT 'go文件名',
  `module_front_name` varchar(255) DEFAULT NULL COMMENT '前端文件名',
  `business_name` varchar(255) DEFAULT NULL COMMENT '業務邏輯名稱',
  `function_name` varchar(255) DEFAULT NULL COMMENT '功能名稱',
  `function_author` varchar(255) DEFAULT NULL COMMENT '功能作者',
  `pk_column` varchar(255) DEFAULT NULL,
  `pk_go_field` varchar(255) DEFAULT NULL,
  `pk_json_field` varchar(255) DEFAULT NULL,
  `options` varchar(255) DEFAULT NULL,
  `tree_code` varchar(255) DEFAULT NULL,
  `tree_parent_code` varchar(255) DEFAULT NULL,
  `tree_name` varchar(255) DEFAULT NULL,
  `tree` tinyint(1) DEFAULT 0,
  `crud` tinyint(1) DEFAULT 1,
  `remark` varchar(255) DEFAULT NULL,
  `is_data_scope` tinyint(4) DEFAULT NULL,
  `is_actions` tinyint(4) DEFAULT NULL,
  `is_auth` tinyint(4) DEFAULT NULL,
  `is_logical_delete` varchar(1) DEFAULT NULL,
  `logical_delete` tinyint(1) DEFAULT NULL,
  `logical_delete_column` varchar(128) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`table_id`),
  KEY `idx_sys_tables_create_by` (`create_by`),
  KEY `idx_sys_tables_update_by` (`update_by`),
  KEY `idx_sys_tables_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_tables` WRITE;
/*!40000 ALTER TABLE `sys_tables` DISABLE KEYS */;

INSERT INTO `sys_tables` (`table_id`, `table_name`, `table_comment`, `class_name`, `tpl_category`, `package_name`, `module_name`, `module_front_name`, `business_name`, `function_name`, `function_author`, `pk_column`, `pk_go_field`, `pk_json_field`, `options`, `tree_code`, `tree_parent_code`, `tree_name`, `tree`, `crud`, `remark`, `is_data_scope`, `is_actions`, `is_auth`, `is_logical_delete`, `logical_delete`, `logical_delete_column`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`)
VALUES
	(1,'article','內容系統','Article','crud','admin','article','','article','內容','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2023-02-20 12:01:17.492','2023-02-20 12:03:25.025',NULL,0,0);

/*!40000 ALTER TABLE `sys_tables` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `username` varchar(64) DEFAULT NULL COMMENT '使用者',
  `password` varchar(128) DEFAULT NULL COMMENT '密碼',
  `nick_name` varchar(128) DEFAULT NULL COMMENT '暱稱',
  `phone` varchar(11) DEFAULT NULL COMMENT '手機號碼',
  `role_id` bigint(20) DEFAULT NULL COMMENT '角色ID',
  `salt` varchar(255) DEFAULT NULL COMMENT 'slat',
  `avatar` varchar(255) DEFAULT NULL COMMENT '頭貼',
  `sex` varchar(255) DEFAULT NULL COMMENT '性别',
  `email` varchar(128) DEFAULT NULL COMMENT '信箱',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部門',
  `post_id` bigint(20) DEFAULT NULL COMMENT '職稱',
  `remark` varchar(255) DEFAULT NULL COMMENT '備註',
  `status` varchar(4) DEFAULT NULL COMMENT '狀態',
  `create_by` bigint(20) DEFAULT NULL COMMENT '建立者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '建立時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最後更新時間',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '刪除時間',
  PRIMARY KEY (`user_id`),
  KEY `idx_sys_user_create_by` (`create_by`),
  KEY `idx_sys_user_update_by` (`update_by`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;

INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'admin','$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.','andy','0912345678',1,'','','0','andy.wang@gaia.net',1,1,'','2',1,1,'2021-05-13 19:56:37.914','2023-02-20 11:29:38.918',NULL),
	(2,'andy','$2a$10$okQ8tSdcFqY9YrcBUJBRnePakqeCD04Fn1dN08hJuaAnPWlRfNPTW','andy','0912339728',0,'','','','test@test.com',1,0,'','2',1,0,'2023-02-20 11:29:14.582','2023-02-20 11:29:14.582',NULL),
	(3,'sales1','$2a$10$s1xI7B/eq9S3.lMVaGf4HuXXOzezspCr79uE5F/otakmWrit82a9e','sales1','0912345678',2,'','','','sales1@test.com',11,1,'','2',1,0,'2023-02-20 11:48:54.245','2023-02-20 11:48:54.245',NULL),
	(4,'sales2','$2a$10$Ltz/xxTFNILfxqPpwBlYi.pl4tQn7oYnK7jy.v29jTQHTRbqzLF1y','sales2','0912345678',2,'','','','sales2@test.com',12,1,'','2',1,0,'2023-02-20 11:53:22.463','2023-02-20 11:53:22.463',NULL),
	(5,'sales2-1','$2a$10$WPbXOtuiSyVKM7Dm0bnDkO/Jj.tc4frmJLhB6ouTSN3ulLmKWyjzG','sales2-1','0912345678',3,'','','','sales2-1@test.com',12,3,'','2',1,0,'2023-02-20 11:54:15.749','2023-02-20 11:54:15.749',NULL),
	(6,'sales1-1','$2a$10$LrTDODWy08V0Jd1t3vdij.jRBMc7mLl54ccn0qvMQit6twF9gEiNe','sales1-1','0912345678',3,'','','','sales1-1@test.com',11,3,'','2',1,0,'2023-02-20 11:54:55.383','2023-02-20 11:54:55.383',NULL);

/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
