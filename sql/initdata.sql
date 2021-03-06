-- MySQL dump 10.13  Distrib 8.0.16, for Win64 (x86_64)
--
-- Host: 10.0.0.11    Database: cms_user
-- ------------------------------------------------------
-- Server version	5.7.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (1,'2018-04-17 10:04:59','2018-04-17 10:04:59',NULL,0,1,'管理员用户','','user','/layout/manager/user/list',1),(2,'2018-04-17 10:04:59','2018-04-17 10:04:59',NULL,0,1,'模块管理','','home','/layout/manager/module/list',1),(10,'2018-04-19 15:45:40','2018-04-19 15:45:40',NULL,1,2,'获取用户列表','','','/user/user/list',1),(11,'2018-04-19 15:46:06','2018-04-19 15:46:06',NULL,1,2,'删除用户','','','/user/user/delete',1),(12,'2018-04-19 15:46:30','2018-04-19 15:46:30',NULL,1,2,'创建用户','','','/user/user/new',1),(13,'2018-04-19 15:47:03','2018-04-19 15:47:03',NULL,1,2,'更新用户信息','目前能修改用户的权限信息.','','/user/user/update',1),(14,'2018-04-19 15:47:35','2018-04-19 15:47:35',NULL,2,2,'创建模块','','','/user/menu/new',1),(15,'2018-04-19 15:48:06','2018-04-19 15:48:06',NULL,2,2,'更新模块信息','','','/user/menu/update',1),(16,'2018-04-19 15:48:25','2018-04-19 15:48:25',NULL,2,2,'获取模块列表','','','/user/menu/list',1),(17,'2018-04-19 15:48:50','2018-04-19 15:48:50',NULL,1,2,'获取用户详细信息','','','/user/user/info',1),(67,'2019-06-27 17:00:20','2019-06-27 17:00:20',NULL,0,1,'角色管理','','meh-o','/layout/manager/role/list',1),(68,'2019-07-03 13:56:54','2019-07-03 13:56:54',NULL,67,2,'获取角色列表','','','/user/role/list',1),(69,'2019-07-03 13:57:39','2019-07-03 13:57:39',NULL,67,2,'增加角色','','','/user/role/new',1),(70,'2019-07-03 13:57:49','2019-07-03 13:57:49',NULL,67,2,'更新角色','','','/user/role/update',1),(71,'2019-07-03 13:57:58','2019-07-03 13:57:58',NULL,67,2,'获取角色信息','','','/user/role/info',1);
/*!40000 ALTER TABLE `menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `role_menus`
--

LOCK TABLES `role_menus` WRITE;
/*!40000 ALTER TABLE `role_menus` DISABLE KEYS */;
INSERT INTO `role_menus` VALUES (1,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,1,1),(2,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,2,1),(3,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,10,1),(4,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,11,1),(5,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,12,1),(6,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,13,1),(7,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,14,1),(8,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,15,1),(9,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,16,1),(10,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,17,1),(11,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,67,1),(12,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,68,1),(13,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,69,1),(14,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,70,1),(15,'2020-03-19 10:47:17','2020-03-19 10:47:17',NULL,1,71,1);
/*!40000 ALTER TABLE `role_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'2019-06-27 15:51:21','2020-03-12 11:05:37',NULL,'超级管理员',1);
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2018-04-11 21:45:45','2019-06-26 15:52:29',NULL,'admin','b11ecdab356639a591477f8153394388','5319816990632489129',1,'2100-06-26 15:52:29',1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-19 10:48:24
