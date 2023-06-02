//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package download

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"make_data_set_so-vits-svc/config"
	"make_data_set_so-vits-svc/model"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var birthday, _ = time.ParseInLocation("2006-01-02", config.ConfigModel.AuthorBirthday, time.Local)

const dir = "./data/source/kugou/"
const api = "https://wwwapi.kugou.com/yy/index.php?r=play/getdata&callback=jQuery191015294419033165485_1674051666168&dfid=4XSnWz14ZQos2PYFIl2MiDLH&appid=1014&mid=8a6709b0f4f0674f12dabeb3a710313a&platid=4&album_audio_id=%s&_=1674051666169"

type linkData struct {
	Status  int `json:"status"`
	ErrCode int `json:"err_code"`
	Data    struct {
		Hash       string `json:"hash"`
		Timelength int    `json:"timelength"`
		Filesize   int    `json:"filesize"`
		AudioName  string `json:"audio_name"`
		HaveAlbum  int    `json:"have_album"`
		AlbumName  string `json:"album_name"`
		//AlbumId    string `json:"album_id"` json: cannot unmarshal number into Go struct field .data.album_id of type string
		Img        string `json:"img"`
		HaveMv     int    `json:"have_mv"`
		VideoId    int    `json:"video_id"`
		AuthorName string `json:"author_name"`
		SongName   string `json:"song_name"`
		Lyrics     string `json:"lyrics"`
		AuthorId   string `json:"author_id"`
		Privilege  int    `json:"privilege"`
		Privilege2 string `json:"privilege2"`
		PlayUrl    string `json:"play_url"`
		Authors    []struct {
			AuthorId      string `json:"author_id"`
			AuthorName    string `json:"author_name"`
			IsPublish     string `json:"is_publish"`
			SizableAvatar string `json:"sizable_avatar"`
			EAuthorId     string `json:"e_author_id"`
			Avatar        string `json:"avatar"`
		} `json:"authors"`
		IsFreePart         int    `json:"is_free_part"`
		Bitrate            int    `json:"bitrate"`
		RecommendAlbumId   int    `json:"recommend_album_id"`
		StoreType          string `json:"store_type"`
		AlbumAudioId       int    `json:"album_audio_id"`
		IsPublish          int    `json:"is_publish"`
		EAuthorId          string `json:"e_author_id"`
		AudioId            string `json:"audio_id"`
		HasPrivilege       bool   `json:"has_privilege"`
		PlayBackupUrl      string `json:"play_backup_url"`
		SmallLibrarySong   int    `json:"small_library_song"`
		EncodeAlbumId      string `json:"encode_album_id"`
		EncodeAlbumAudioId string `json:"encode_album_audio_id"`
		EVideoId           string `json:"e_video_id"`
	} `json:"data"`
}

func getLink(albumAudioId string) (link string, err error) {
	res, err := http.Get(fmt.Sprintf(api, albumAudioId))
	if err != nil {
		return
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var data linkData
	err = json.Unmarshal([]byte(removeGarbage(string(bytes))), &data)
	if err != nil {
		return
	}
	link = data.Data.PlayUrl
	return
}

func getBackUpLink(albumAudioId string) (link string, err error) {
	res, err := http.Get(fmt.Sprintf(api, albumAudioId))
	if err != nil {
		return
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var data linkData
	err = json.Unmarshal([]byte(removeGarbage(string(bytes))), &data)
	if err != nil {
		return
	}
	link = data.Data.PlayBackupUrl
	return
}

func download(link string, filename string) (err error) {
	res, err := http.Get(link)
	if err != nil {
		return
	}
	defer res.Body.Close()

	file, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return
	}
	return
}

func removeGarbage(str string) (newStr string) {
	//查找第一个(的位置
	start := strings.Index(str, "(")
	//查找最后一个)的位置
	end := strings.LastIndex(str, ")")
	//截取字符串
	newStr = str[start+1 : end]
	return
}

func getAge(t string) int {
	publishDate, err := time.ParseInLocation("2006-01-02", t, time.Local)
	if err != nil {
		panic(err)
	}
	age := int(math.Floor(publishDate.Sub(birthday).Hours() / 24 / 365))
	if age < 0 {
		age = 0
	}
	return age
}

func Download(kugouModel []*model.KuGouModel) (err error) {
	num := 1
	for _, v := range kugouModel {
		for _, vv := range v.Data {
			link, err2 := getLink(fmt.Sprint(vv.AlbumAudioId))
			err = err2
			if err != nil || link == "" {
				link, err = getBackUpLink(fmt.Sprint(vv.AlbumAudioId))
				if err != nil || link == "" {
					log.Println("下载链接错误", vv.AudioName, err)
					continue
				}
			}
			err = download(link, fmt.Sprintf("%d-%d-%s-%s", num, getAge(vv.PublishDate), vv.PublishDate, fmt.Sprint(vv.AudioName, filepath.Ext(link))))
			if err != nil {
				log.Println("下载错误", err)
				continue
			}
			num++
		}
	}
	return
}
