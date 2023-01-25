CREATE TABLE `t_user` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `passwd` varchar(32) NOT NULL
);

CREATE TABLE `t_video` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `owner_id` int,
  `title` varchar(255),
  `play_url` varchar(255),
  `cover_url` varchar(255)
);

CREATE TABLE `t_video_comment` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `video_id` int,
  `writer_id` int,
  `cotent` varchar(255)
);

CREATE TABLE `t_follow` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int,
  `fans_id` int
);

ALTER TABLE `t_video` ADD FOREIGN KEY (`owner_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_video_comment` ADD FOREIGN KEY (`video_id`) REFERENCES `t_video` (`id`);

ALTER TABLE `t_video_comment` ADD FOREIGN KEY (`writer_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_follow` ADD FOREIGN KEY (`user_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_follow` ADD FOREIGN KEY (`fans_id`) REFERENCES `t_user` (`id`);
