/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50553
Source Host           : 127.0.0.1:3306
Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2021-01-04 17:56:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `blog_list`
-- ----------------------------
DROP TABLE IF EXISTS `blog_list`;
CREATE TABLE `blog_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `title` varchar(255) NOT NULL COMMENT '文章标题',
  `introduction` varchar(255) NOT NULL COMMENT '文章简介',
  `category_id` int(11) NOT NULL COMMENT '文章分类',
  `createtime` int(11) NOT NULL COMMENT '创建时间',
  `tag_id` int(11) NOT NULL COMMENT '标签id',
  `content` text NOT NULL COMMENT '文章内容',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=34 DEFAULT CHARSET=utf8 COMMENT='博客列表';

-- ----------------------------
-- Records of blog_list
-- ----------------------------
INSERT INTO `blog_list` VALUES ('31', 'test', 'asdf', '1', '1609745853', '1', '<h3>关于 Editor.md</h3>\n\n<p><strong>Editor.md</strong> 是一款开源的、可嵌入的 Markdown 在线编辑器（组件），基于 CodeMirror、jQuery 和 Marked 构建。</p>\n');
INSERT INTO `blog_list` VALUES ('32', 'test', 'asdf', '1', '1609745905', '1', '<h3>关于 Editor.md</h3>\n\n<p><strong>Editor.md</strong> 是一款开源的、可嵌入的 Markdown 在线编辑器（组件），基于 CodeMirror、jQuery 和 Marked 构建。</p>\n\n<pre><code>sdfasdfa\nsadfasfasd\n</code></pre>\n');
INSERT INTO `blog_list` VALUES ('33', 'test', 'asdf', '1', '1609745948', '1', '<h3>关于 Editor.md</h3>\n\n<p><strong>Editor.md</strong> 是一款开源的、可嵌入的 Markdown 在线编辑器（组件），基于 CodeMirror、jQuery 和 Marked 构建。</p>\n\n<pre><code>sdfasdfa\nsadfasfasd\nthis.Data[&#34;title&#34;] = beego.AppConfig.String(&#34;title&#34;)\n	this.TplName = &#34;backend/add.html&#34;\n</code></pre>\n');
