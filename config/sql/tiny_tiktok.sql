CREATE TABLE `t_user` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(32) UNIQUE NOT NULL,
  `passwd` varchar(32) NOT NULL,
  `follow_count` bigint NOT NULL DEFAULT 0,
  `follower_count` bigint NOT NULL DEFAULT 0,
  `avatar` varchar(255) DEFAULT 'https://cdn-icons-png.flaticon.com/512/149/149071.png',
  `background_image` varchar(255) NOT NULL DEFAULT '',
  `signature` varchar(255) NOT NULL DEFAULT '',
  `total_favorited` bigint NOT NULL DEFAULT 0,
  `work_count` bigint NOT NULL DEFAULT 0,
  `favorite_count` bigint NOT NULL DEFAULT 0
);

CREATE TABLE `t_follow` (
  `user_id` bigint,
  `fans_id` bigint,
  PRIMARY KEY (`user_id`, `fans_id`)
);

CREATE TABLE `t_message` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `from` bigint,
  `to` bigint,
  `content` varchar(255) NOT NULL,
  `create_time` datetime NOT NULL
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

ALTER TABLE `t_message` ADD FOREIGN KEY (`from`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_message` ADD FOREIGN KEY (`to`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_video` ADD FOREIGN KEY (`author_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_video_comment` ADD FOREIGN KEY (`video_id`) REFERENCES `t_video` (`id`);

ALTER TABLE `t_video_comment` ADD FOREIGN KEY (`writer_id`) REFERENCES `t_user` (`id`);

ALTER TABLE `t_video_favorite` ADD FOREIGN KEY (`video_id`) REFERENCES `t_video` (`id`);

ALTER TABLE `t_video_favorite` ADD FOREIGN KEY (`liker_id`) REFERENCES `t_user` (`id`);

INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`,`work_count`) VALUES(1,'DOU1','1234561','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiFs.img',2);
INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`) VALUES(2,'DOU2','1234562','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa3q.img');
INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`) VALUES(3,'DOU3','1234563','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRJ.img');
INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`) VALUES(4,'DOU4','1234564','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdP.img');
INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`) VALUES(5,'DOU5','1234565','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRI.img');

INSERT INTO `t_video`(`id`,`author_id`,`title`,`publish_time`,`play_url`,`cover_url`) VALUES (1,1,'DOU1 whale','2023-02-11 11:09:17','https://prod-streaming-video-msn-com.akamaized.net/aa5cb260-7dae-44d3-acad-3c7053983ffe/1b790558-39a2-4d2a-bcd7-61f075e87fdd.mp4','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElb0.img');
INSERT INTO `t_video`(`id`,`author_id`,`title`,`publish_time`,`play_url`,`cover_url`) VALUES (2,1,'DOU1 beach','2023-02-11 11:12:17','https://prod-streaming-video-msn-com.akamaized.net/559310a7-dbb0-461c-a863-5cb758607af5/f0474526-90d0-4d3d-aaae-dd68f3f38b28.mp4','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdS.img');
