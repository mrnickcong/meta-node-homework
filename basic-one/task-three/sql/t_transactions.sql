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

 Date: 12/11/2025 14:02:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_transactions
-- ----------------------------
DROP TABLE IF EXISTS `t_transactions`;
CREATE TABLE `t_transactions` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `from_account_id` int DEFAULT NULL COMMENT '转出账户ID',
  `to_account_id` int DEFAULT NULL COMMENT '转入账户ID',
  `amount` decimal(10,2) DEFAULT NULL COMMENT '转账金额',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of t_transactions
-- ----------------------------
BEGIN;
INSERT INTO `t_transactions` (`id`, `from_account_id`, `to_account_id`, `amount`) VALUES (1, 1, 2, 100.00);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
