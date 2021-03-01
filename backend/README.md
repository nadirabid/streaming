#### Download/create HLS content for local

```
sh download.sh
```


#### Gen HLS files

Option 1

```bash
sh vlod.sh filename.mkv
```

Option 2

```bash
VID_NAME=IJS9E1 && FORMAT=mkv; eval 'mkdir $VID_NAME; ffmpeg -i $VID_NAME.$FORMAT -vf scale=w=1280:h=720:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -b:a 128k -c:v h264 -profile:v main -crf 20 -g 48 -keyint_min 48 -sc_threshold 0 -b:v 2500k -maxrate 2675k -bufsize 3750k -hls_time 4 -hls_playlist_type vod -hls_segment_filename $VID_NAME/720p_%03d.ts $VID_NAME/720p.m3u8'
```

#### Create Database

```
psql -d postgres -f setup.sql
```

To access db on terminal:

```
psql -d streaming
```

Also - download Postico for Mac - super good.
Or checkout: https://www.beekeeperstudio.io

#### Migration

Using `golang-migrate`. Install CLI to generate the files:

```
brew install golang-migrate
```

To generate new migration files:

```
migrate create -ext sql -dir migrations -seq <name_of_migration>
```

For local testing - to create a db copy:

```sh
psql -d exercise_parser
CREATE DATABASE exercise_parser_test WITH TEMPLATE exercise_parser;
```

