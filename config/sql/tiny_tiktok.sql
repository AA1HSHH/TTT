CREATE TABLE `t_user` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(32) UNIQUE NOT NULL,
  `passwd` varchar(60) NOT NULL,
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
  `create_time` bigint NOT NULL
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

INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`,`work_count`,`background_image`,`signature`) VALUES(1,'DOU1','$2a$14$VSYjOFVVrZIiKhH3QKBXL.PjXQ2DqCj.i9Otzdpe.tLqNcaCTI49e','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiFs.img',2,'https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiFs.img','Carpe Diem.');
INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`) VALUES(2,'DOU2','$2a$14$SxKboOBj8/yUlIFnrpkADuEyGMW0h/yuRoWe9mP.4hkb7.0FSptMW','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa3q.img');
INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`) VALUES(3,'DOU3','$2a$14$uksl0e5kiYo5ZPxAabxhQOOpOoK9VoHOI2qM8eoiqRSptDLdDUjPi','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRJ.img');
INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`) VALUES(4,'DOU4','$2a$14$mdh8SU4HjC9Lm8FfLVDdiuygKAO92Tjm41YwXK9wxA8FO9ngPrBby','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdP.img');
INSERT INTO `t_user`(`id`,`name`,`passwd`,`avatar`) VALUES(5,'DOU5','$2a$14$yirD4Rubtz3hlxf2/pS6bOiGaqqDBx7be6JFqAZpaOr0FWWhwqnpi','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRI.img');

INSERT INTO `t_video`(`id`,`author_id`,`title`,`publish_time`,`play_url`,`cover_url`) VALUES (1,1,'DOU1 whale','2023-02-11 11:09:17','https://prod-streaming-video-msn-com.akamaized.net/aa5cb260-7dae-44d3-acad-3c7053983ffe/1b790558-39a2-4d2a-bcd7-61f075e87fdd.mp4','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElb0.img');
INSERT INTO `t_video`(`id`,`author_id`,`title`,`publish_time`,`play_url`,`cover_url`) VALUES (2,1,'DOU1 beach','2023-02-11 11:12:17','https://prod-streaming-video-msn-com.akamaized.net/559310a7-dbb0-461c-a863-5cb758607af5/f0474526-90d0-4d3d-aaae-dd68f3f38b28.mp4','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdS.img');
INSERT INTO `t_video`(`id`,`author_id`,`title`,`publish_time`,`play_url`,`cover_url`) VALUES (3,3,'DOU3 mountain','2023-02-11 10:09:17','https://prod-streaming-video-msn-com.akamaized.net/0b927d99-e38a-4f51-8d1a-598fd4d6ee97/3493c85c-f35a-488f-9a8f-633e747fb141.mp4','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRG.img');
INSERT INTO `t_video`(`id`,`author_id`,`title`,`publish_time`,`play_url`,`cover_url`) VALUES (4,5,'DOU5 farmland','2023-02-11 10:12:17','https://prod-streaming-video-msn-com.akamaized.net/e6f8c1b2-b1ac-4343-a8bf-bd79a4d25380/9de24622-e13e-4741-95a6-04fdc39eb2b0.mp4','https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhXX.img');