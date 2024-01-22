# TTT
## deploy mysql
```shell
sudo docker compose up -d
sudo docker compose down -v
```
## install ffmpeg in ubuntu
```shell
sudo apt update
sudo apt install ffmpeg
```
## run

```shell
go build && ./TTT
```


### 功能说明

* 视频上传后会保存到本地 public 目录中，访问时用 ip:8080/static/video_name 即可

## doc
https://za0i9tolqu.feishu.cn/docx/CvsWdhOEfooBa0xxsEocJk4Zni9


