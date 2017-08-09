/*
Navicat MySQL Data Transfer

Source Server         : localmysql
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : temp

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2017-08-01 23:27:54
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for data
-- ----------------------------
DROP TABLE IF EXISTS `data`;
CREATE TABLE `data` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `date` datetime NOT NULL,
  `value` varchar(255) NOT NULL,
  PRIMARY KEY (`id`,`date`)
) ENGINE=MyISAM AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of data
-- ----------------------------
INSERT INTO `data` VALUES ('1', '2017-08-01 20:35:45', '27.00000');
INSERT INTO `data` VALUES ('2', '2017-08-01 21:10:59', '26.00000');
INSERT INTO `data` VALUES ('3', '2017-08-01 21:27:44', '26.50000');
INSERT INTO `data` VALUES ('4', '2017-08-01 22:17:35', '28.70000');
INSERT INTO `data` VALUES ('5', '2017-08-01 22:17:56', '25.30000');
INSERT INTO `data` VALUES ('6', '2017-08-01 22:18:17', '26.00000');
INSERT INTO `data` VALUES ('7', '2017-08-01 22:18:20', '30.00000');
INSERT INTO `data` VALUES ('8', '2017-08-01 22:18:24', '30.67000');
INSERT INTO `data` VALUES ('12', '2017-08-01 15:13:16', '46.98000');
INSERT INTO `data` VALUES ('13', '2017-08-01 15:13:18', '46.98000');
INSERT INTO `data` VALUES ('14', '2017-08-01 15:14:21', '32.7');
INSERT INTO `data` VALUES ('15', '2017-08-01 15:15:45', '91.32');
INSERT INTO `data` VALUES ('16', '2017-08-01 15:21:49', '32.91');
