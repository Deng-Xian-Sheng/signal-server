//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package model

// KuGouModel 酷狗Model，接口路由 /kmr/v1/audio_group/author
// 请求此路由稍微复杂，需要构造很多参数
// 有一个简便的方式，开启抓包软件，打开酷狗音乐app，打开我的主页，点击我的作品，不断下拉刷新
// 将接口返回值保存到json文件
type KuGouModel struct {
	Extra struct {
		PageTotal int `json:"page_total"` // 总页数
		Group     int `json:"group"`      // 分组数量
	} `json:"extra"`
	Data []struct {
		RpType128        string `json:"rp_type_128"`        // 128kbps 文件的 rp 类型
		IsOriginal       int    `json:"is_original"`        // 是否为原创
		Filesize128      int    `json:"filesize_128"`       // 128kbps 文件大小
		AuthorName       string `json:"author_name"`        // 作者姓名
		FailProcess128   int    `json:"fail_process_128"`   // 处理失败的 128kbps 文件数量
		PayTypeFlac      int    `json:"pay_type_flac"`      // FLAC 文件的付费类型
		OldHideFlac      int    `json:"old_hide_flac"`      // 已隐藏的 FLAC 文件数
		Filesize320      int    `json:"filesize_320"`       // 320kbps 文件大小
		Price            int    `json:"price"`              // 价格
		TopicUrlSuper    string `json:"topic_url_super"`    // 超高品质文件的主题 URL
		OldCpy320        int    `json:"old_cpy_320"`        // 320kbps 文件的旧副本数量
		MixsongType      int    `json:"mixsong_type"`       // 混合歌曲类型
		Extname          string `json:"extname"`            // 扩展名
		Filesize         int    `json:"filesize"`           // 文件大小
		TopicUrlFlac     string `json:"topic_url_flac"`     // FLAC 文件的主题 URL
		Price320         int    `json:"price_320"`          // 320kbps 文件价格
		PkgPriceSuper    int    `json:"pkg_price_super"`    // 超高品质文件的套餐价格
		RpTypeSuper      string `json:"rp_type_super"`      // 超高品质文件的 rp 类型
		AudioName        string `json:"audio_name"`         // 音频名称
		RpId             int    `json:"rp_id"`              // rp 编号
		FailProcessSuper int    `json:"fail_process_super"` // 处理失败的超高品质文件数量
		FailProcessFlac  int    `json:"fail_process_flac"`  // 处理失败的 FLAC 文件数量
		TransParam       struct {
			Qualitymap struct {
				Attr0 int `json:"attr0"` // 属性0
			} `json:"qualitymap"`
			CpyAttr0         int `json:"cpy_attr0"`           // 拷贝属性0
			CpyGrade         int `json:"cpy_grade,omitempty"` // 拷贝等级
			PayBlockTpl      int `json:"pay_block_tpl"`       // 支付阻止模板
			MusicpackAdvance int `json:"musicpack_advance"`   // 音乐包预付费
			Display          int `json:"display"`             // 显示
			CpyLevel         int `json:"cpy_level,omitempty"` // 拷贝级别
			Classmap         struct {
				Attr0 int `json:"attr0"` // 属性0
			} `json:"classmap"`
			DisplayRate    int    `json:"display_rate"`              // 显示比例
			Cid            int    `json:"cid"`                       // CID
			HashMultitrack string `json:"hash_multitrack,omitempty"` // 多音轨哈希
			HashOffset     struct {
				EndMs      int    `json:"end_ms"`      // 结束毫秒
				ClipHash   string `json:"clip_hash"`   // 剪辑哈希
				FileType   int    `json:"file_type"`   // 文件类型
				EndByte    int    `json:"end_byte"`    // 结束字节
				OffsetHash string `json:"offset_hash"` // 偏移哈希
				StartByte  int    `json:"start_byte"`  // 开始字节
				StartMs    int    `json:"start_ms"`    // 开始毫秒
			} `json:"hash_offset,omitempty"`
			CpyCover       int    `json:"cpy_cover,omitempty"`        // 拷贝封面
			AppidBlock     string `json:"appid_block,omitempty"`      // 应用程序 ID 块
			AllQualityFree int    `json:"all_quality_free,omitempty"` // 全部质量免费
		} `json:"trans_param"`
		AudioId          int    `json:"audio_id"`           // 音频 ID
		StatusSuper      int    `json:"status_super"`       // 超高品质文件的状态
		Version          int    `json:"version"`            // 版本
		AlbumAudioRemark string `json:"album_audio_remark"` // 专辑音频备注
		RpTypeHigh       string `json:"rp_type_high"`       // 高品质文件的 rp 类型
		PayType128       int    `json:"pay_type_128"`       // 128kbps 文件的付费类型
		Musical          struct {
			PublishType     int    `json:"publish_type,omitempty"`     // 发布类型
			Uploader        string `json:"uploader,omitempty"`         // 上传者
			PublishTime     string `json:"publish_time,omitempty"`     // 发布时间
			UploaderContent string `json:"uploader_content,omitempty"` // 上传者内容
		} `json:"musical"`
		BitrateFlac     int    `json:"bitrate_flac"         ` // FLAC 文件的比特率
		OldHide         int    `json:"old_hide"             ` // 已隐藏的文件数量
		RpType          string `json:"rp_type"              ` // 文件的 rp 类型
		FailProcess     int    `json:"fail_process"         ` // 处理失败的文件数量
		RpType320       string `json:"rp_type_320"          ` // 320kbps 文件的 rp 类型
		FailProcessHigh int    `json:"fail_process_high"    ` // 处理失败的高品质文件数量
		StatusFlac      int    `json:"status_flac"          ` // FLAC 文件的状态
		PayType320      int    `json:"pay_type_320"         ` // 320kbps 文件的付费类型
		PkgPrice128     int    `json:"pkg_price_128"        ` // 128kbps 文件的套餐价格
		Status320       int    `json:"status_320"           ` // 320kbps 文件的状态
		Price128        int    `json:"price_128"            ` // 128kbps 文件的价格
		VideoHash       string `json:"video_hash"           ` // 视频哈希值
		FilesizeHigh    int    `json:"filesize_high"        ` // 高品质文件大小
		PkgPriceHigh    int    `json:"pkg_price_high"       ` // 高品质文件的套餐价格
		TimelengthHigh  int    `json:"timelength_high"      ` // 高品质文件时长
		Identity        int    `json:"identity"             ` // 身份标识
		OldCpy          int    `json:"old_cpy"              ` // 旧副本数量
		CdUrl           string `json:"cd_url"               ` // CD URL
		PrivilegeHigh   int    `json:"privilege_high"       ` // 高品质文件的权限
		Status128       int    `json:"status_128"           ` // 128kbps 文件的状态
		PublishDate     string `json:"publish_date"         ` // 发布日期
		Privilege320    int    `json:"privilege_320"        ` // 320kbps 文件的权限
		PrivilegeSuper  int    `json:"privilege_super"      ` // 超高品质文件的权限
		Cid             int    `json:"cid"                  ` // 内容标识符
		PriceHigh       int    `json:"price_high"           ` // 高品质文件的价格
		TopicRemark     string `json:"topic_remark"         ` // 主题备注
		Bitrate         int    `json:"bitrate"              ` // 比特率
		Hash320         string `json:"hash_320"             ` // 320kbps 文件的哈希值
		OldCpyFlac      int    `json:"old_cpy_flac"         ` // FLAC 文件的旧副本数量
		PayType         int    `json:"pay_type"             ` // 付费类型
		Remarks         []struct {
			Remark          string `json:"remark"`                    // 备注
			RemarkType      int    `json:"remark_type"`               // 备注类型
			RelAlbumAudioId int    `json:"rel_album_audio_id"`        // 关联专辑音频ID
			IsDefautAlias   string `json:"is_defaut_alias,omitempty"` // 是否默认别名
		} `json:"remarks" ` // 备注信息
		Timelength320  int    `json:"timelength_320"`  // 320码率时长
		TimelengthFlac int    `json:"timelength_flac"` // FLAC码率时长
		TopicUrl320    string `json:"topic_url_320"`   // 320码率专题地址
		PayTypeSuper   int    `json:"pay_type_super"`  // 超级会员付费类型
		HashFlac       string `json:"hash_flac"`       // FLAC哈希值
		OldHideSuper   int    `json:"old_hide_super"`  // 超级会员是否隐藏
		FilesizeFlac   int    `json:"filesize_flac"`   // FLAC文件大小
		AlbumName      string `json:"album_name"`      // 专辑名
		OldHide128     int    `json:"old_hide_128"`    // 128码率是否隐藏
		HashSuper      string `json:"hash_super"`      // 超级会员哈希值
		Timelength     int    `json:"timelength"`      // 歌曲时长
		AlbumId        int    `json:"album_id"`        // 专辑ID
		OldCpySuper    int    `json:"old_cpy_super"`   // 超级会员是否拷贝
		StatusHigh     int    `json:"status_high"`     // 高品质状态值
		TopicUrlHigh   string `json:"topic_url_high"`  // 高品质专题地址
		Mvdata         []struct {
			Trk  string `json:"trk" `  // MV歌词
			Hash string `json:"hash" ` // MV哈希值
			Id   string `json:"id"   ` // MV ID
			Typ  int    `json:"typ" `  // MV类型
		} `json:"mvdata"` // MV相关信息
		OldHide320      int    `json:"old_hide_320"`     // 320码率是否隐藏
		Playcount       int    `json:"playcount"`        // 播放次数
		VideoFilesize   int    `json:"video_filesize"`   // 视频文件大小
		Privilege       int    `json:"privilege"`        // 用户权限
		VideoFileHead   int    `json:"video_file_head"`  // 视频文件头
		FailProcess320  int    `json:"fail_process_320"` // 320码率失败处理
		BitrateHigh     int    `json:"bitrate_high"`     // 高品质码率
		TmpPrivilege    int    `json:"tmp_privilege"`    // 临时权限
		PriceFlac       int    `json:"price_flac"`       // FLAC价格
		VideoId         int    `json:"video_id"`         // 视频ID
		RpTypeFlac      string `json:"rp_type_flac"`     // FLAC专辑类型
		OldCpy128       int    `json:"old_cpy_128"`      // 128码率是否拷贝
		BitrateSuper    int    `json:"bitrate_super"`    // 超级会员码率
		Songid          int    `json:"songid"`           // 歌曲ID
		Status          int    `json:"status"`           // 歌曲状态值
		Privilege128    int    `json:"privilege_128"`    // 128码率用户权限
		Level           int    `json:"level"`            // 用户等级
		Hash128         string `json:"hash_128"`         // 128码率哈希值
		PayTypeHigh     int    `json:"pay_type_high"`    // 高品质会员付费类型
		AudioFileHead   int    `json:"audio_file_head"`  // 音频文件头
		HasObbligato    int    `json:"has_obbligato"`    // 是否有伴奏
		FilesizeSuper   int    `json:"filesize_super"`   // 超级会员文件大小
		OldHideHigh     int    `json:"old_hide_high"`    // 高品质是否隐藏
		VideoTimelength int    `json:"video_timelength"` // 视频时长
		ExtnameSuper    string `json:"extname_super"`    // 超级会员扩展名
		PkgPriceFlac    int    `json:"pkg_price_flac"`   // FLAC打包价格
		Timelength128   int    `json:"timelength_128"`   // 128码率时长
		AlbumAudioId    int    `json:"album_audio_id"`   // 专辑音频ID
		TimelengthSuper int    `json:"timelength_super"` // 超级会员时长
		VideoTrack      int    `json:"video_track"`      // 视频轨道
		TopicUrl128     string `json:"topic_url_128"`    // 128码率专题地址
		Hash            string `json:"hash"`             // 哈希值
		HashHigh        string `json:"hash_high"`        // 高品质哈希值
		PkgPrice320     int    `json:"pkg_price_320"`    // 320码率打包价格
		RpPublish       int    `json:"rp_publish"`       // 是否发布歌曲
		TopicUrl        string `json:"topic_url"`        // 歌曲专题地址
		OldCpyHigh      int    `json:"old_cpy_high"`     // 高品质是否拷贝
		PrivilegeFlac   int    `json:"privilege_flac"`   // FLAC 用户权限
		PriceSuper      int    `json:"price_super"`      // 超级会员价格
		PkgPrice        int    `json:"pkg_price"`        // 打包价格
	} `json:"data"`
	Status    int    `json:"status"`     // 状态值
	Errmsg    string `json:"errmsg"`     // 错误信息
	Total     int    `json:"total"`      // 总数
	ErrorCode int    `json:"error_code"` // 错误代码
}
