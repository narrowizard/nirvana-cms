-- MySQL dump 10.13  Distrib 5.7.12, for Win64 (x86_64)
--
-- Host: 10.0.0.11    Database: cms_user
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `menus`
--

DROP TABLE IF EXISTS `menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `is_menu` int(11) DEFAULT NULL,
  `remarks` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (1,'2018-04-17 10:04:59','2018-04-17 10:04:59',NULL,'用户管理','user','/layout/manager/user/list',1,0,1,NULL),(2,'2018-04-17 10:04:59','2018-04-17 10:04:59',NULL,'模块管理','home','/layout/manager/module/list',1,0,1,NULL),(10,'2018-04-19 15:45:40','2018-04-19 15:45:40',NULL,'获取用户列表','','/user/user/list',1,1,2,''),(11,'2018-04-19 15:46:06','2018-04-19 15:46:06',NULL,'删除用户','','/user/user/delete',1,1,2,''),(12,'2018-04-19 15:46:30','2018-04-19 15:46:30',NULL,'创建用户','','/user/user/new',1,1,2,''),(13,'2018-04-19 15:47:03','2018-04-19 15:47:03',NULL,'更新用户信息','','/user/user/update',1,1,2,'目前能修改用户的权限信息.'),(14,'2018-04-19 15:47:35','2018-04-19 15:47:35',NULL,'创建模块','','/user/menu/new',1,2,2,''),(15,'2018-04-19 15:48:06','2018-04-19 15:48:06',NULL,'更新模块信息','','/user/menu/update',1,2,2,''),(16,'2018-04-19 15:48:25','2018-04-19 15:48:25',NULL,'获取模块列表','','/user/menu/list',1,2,2,''),(17,'2018-04-19 15:48:50','2018-04-19 15:48:50',NULL,'获取用户详细信息','','/user/user/info',1,1,2,'');
/*!40000 ALTER TABLE `menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_menus`
--

DROP TABLE IF EXISTS `user_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `menu_id` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_menus`
--

LOCK TABLES `user_menus` WRITE;
/*!40000 ALTER TABLE `user_menus` DISABLE KEYS */;
INSERT INTO `user_menus` VALUES (1,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,10,1),(2,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,11,1),(3,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,12,1),(4,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,13,1),(5,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,17,1),(6,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,1,1),(7,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,14,1),(8,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,15,1),(9,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,16,1),(10,'2018-04-19 15:51:01','2018-04-19 15:51:01',NULL,1,2,1);
/*!40000 ALTER TABLE `user_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `account` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `salt` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2018-04-11 13:45:45','2018-04-11 13:45:45',NULL,'admin','178601f9dd552c33200a7e2961ba62a1','7972777030560804550',1);
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

-- Dump completed on 2018-04-19 15:52:00
