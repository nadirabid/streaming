#!/usr/bin/env bash

set -e # exit when any command fails

rm -rf summer_adrift* # clean old video files

brew install youtube-dl

youtube-dl -f 'bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best' -o summer_adrift1.mp4 'https://www.youtube.com/watch?v=W1yaBmL42tY'
youtube-dl -f 'bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best' -o summer_adrift2.mp4 'https://www.youtube.com/watch?v=SuEZ3dJ4jxc'
youtube-dl -f 'bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best' -o summer_adrift3.mp4 'https://www.youtube.com/watch?v=-j6eKmu6qLE'

for i in 1 2 3
do 
    mkdir -p summer_adrift$i/thumbnails

    # cut to 30 seconds or hls will take time
    ffmpeg -ss 00:00:00 -i summer_adrift$i.mp4 -ss 00:00:00 -t 00:00:30 summer_adrift$i/vid.mp4

    # create hls
    sh vlod.sh summer_adrift$i/vid.mp4 summer_adrift$i/hls

    # create thumbnails
    for size in 842x480 1280x720 1920x1080
    do
        ffmpeg -i summer_adrift$i/vid.mp4 -vf  "thumbnail,scale=$size" -frames:v 1 summer_adrift$i/thumbnails/$size.jpg
    done
done

rm -rf summer_adrift*.mp4
rm -rf summer_adrift*/*.mp4
