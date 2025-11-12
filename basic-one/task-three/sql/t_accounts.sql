/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80100 (8.1.0)
 Source Host           : localhost:3306
 Source Schema         : ethan_blog

 Target Server Type    : MySQL
 Target Server Version : 80100 (8.1.0)
 File Encoding         : 65001

 Date: 12/11/2025 14:02:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_accounts
-- ----------------------------
DROP TABLE IF EXISTS `t_accounts`;
CREATE TABLE `t_accounts` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `balance` decimal(10,2) DEFAULT NULL COMMENT '账户余额',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of t_accounts
-- ----------------------------
BEGIN;
INSERT INTO `t_accounts` (`id`, `balance`) VALUES (1, 900.00);
INSERT INTO `t_accounts` (`id`, `balance`) VALUES (2, 1100.00);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
