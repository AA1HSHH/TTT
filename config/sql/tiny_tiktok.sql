CREATE TABLE `t_user` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(32) UNIQUE NOT NULL,
  `passwd` varchar(32) NOT NULL,
  `follow_count` bigint NOT NULL DEFAULT 0,
  `follower_count` bigint NOT NULL DEFAULT 0
);

CREATE TABLE `t_follow` (
  `user_id` bigint,
  `fans_id` bigint,
  PRIMARY KEY (`user_id`, `fans_id`)
);

CREATE TABLE `t_video` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `author_id` bigint,
  `title` varchar(255) NOT NULL,
  `publish_time` datetime NOT NULL,
  `play_url` varchar(255) NOT NULL,
  `cover_url` varchar(255) NOT NULL,
  `favorite_count` bigint NOT NULL DEFAULT 0,
  `comment_count` bigint NOT NULL DEFAULT 0
);

CREATE TABLE `t_video_comment` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `video_id` bigint,
  `writer_id` bigint,
  `content` varchar(255),
  `create_date` datetime,
  `is_delete` datetime
);

CREATE TABLE `t_video_favorite` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `video_id` bigint,
  `liker_id` bigint,
  `is_delete` datetime
);

CREATE INDEX `t_user_index_0` ON `t_user` (`name`);

CREATE INDEX `t_video_comment_index_1` ON `t_video_comment` (`video_id`);

CREATE INDEX `t_video_favorite_index_2` ON `t_video_favorite` (`video_id`);

CREATE INDEX `t_video_favorite_index_3` ON `t_video_favorite` (`liker_id`);

ALTER TABLE `t_follow` ADD FOREIGN KEY (`user_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_follow` ADD FOREIGN KEY (`fans_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_video` ADD FOREIGN KEY (`author_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_video_comment` ADD FOREIGN KEY (`video_id`) REFERENCES `t_video` (`id`);

ALTER TABLE `t_video_comment` ADD FOREIGN KEY (`writer_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_video_favorite` ADD FOREIGN KEY (`video_id`) REFERENCES `t_video` (`id`);

ALTER TABLE `t_video_favorite` ADD FOREIGN KEY (`liker_id`) REFERENCES `t_user` (`id`);
