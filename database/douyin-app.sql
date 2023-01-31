
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
   `id` varchar(24) NOT NULL,
   `author_id` varchar(24) NOT NULL COMMENT '评论的视频是哪个作者（vloger）的关联id',
   `comment_user_id` varchar(24) NOT NULL COMMENT '发布留言的用户id',
   `content` varchar(128) NOT NULL COMMENT '留言内容',
   `create_time` datetime NOT NULL COMMENT '留言时间',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';


-- ----------------------------
-- Table structure for fans
-- ----------------------------
DROP TABLE IF EXISTS `fans`;
CREATE TABLE `fans` (
    `id` varchar(24) NOT NULL,
    `author_id` varchar(24) NOT NULL COMMENT '作者用户id',
    `fan_id` varchar(24) NOT NULL COMMENT '粉丝用户id',
    `is_fan_friend_of_mine` int(1) NOT NULL COMMENT '粉丝是否是vloger的朋友，如果成为朋友，则本表的双方此字段都需要设置为1，如果有一人取关，则两边都需要设置为0',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `writer_id` (`author_id`,`fan_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='粉丝表';



-- ----------------------------
-- Table structure for my_liked_video
-- ----------------------------
DROP TABLE IF EXISTS `my_liked_vlog`;
CREATE TABLE `my_liked_vlog` (
     `id` varchar(24) NOT NULL,
     `user_id` varchar(24) NOT NULL COMMENT '用户id',
     `video_id` varchar(24) NOT NULL COMMENT '喜欢的短视频id',
     PRIMARY KEY (`id`) USING BTREE,
     UNIQUE KEY `writer_id` (`user_id`,`video_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='点赞短视频关联表';



-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
     `id` varchar(24) NOT NULL,
     `user_name` varchar(16) NOT NULL COMMENT '用户名',
     `password` varchar(32) NOT NULL COMMENT '密码',
     `token` varchar(32) NOT NULL COMMENT 'token',
     `follow_count` int(12) NOT NULL COMMENT '关注总数',
     `follower_count` int(12) NOT NULL COMMENT '粉丝总数',
     PRIMARY KEY (`id`),
     UNIQUE KEY `user_name` (`user_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';



-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
      `id` varchar(24) NOT NULL,
      `user_id` varchar(24) NOT NULL COMMENT '视频作者',
      `play_url` varchar(255) NOT NULL COMMENT '视频播放地址',
      `cover_url` varchar(255) NOT NULL COMMENT '视频封面地址',
      `title` varchar(255) NOT NULL COMMENT '视频标题',
      `favorite_count` int(12) NOT NULL COMMENT '点赞总数',
      `comment_count` int(12) NOT NULL COMMENT '评论总数',
      `created_time` datetime NOT NULL COMMENT '创建时间',
      `updated_time` datetime NOT NULL COMMENT '更新时间',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';



SET FOREIGN_KEY_CHECKS = 1;
