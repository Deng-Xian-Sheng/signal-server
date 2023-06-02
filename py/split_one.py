# Copyright (c) [2023] [JinCanQi]
# [make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
# You can use this software according to the terms and conditions of the Mulan PubL v2.
# You may obtain a copy of Mulan PubL v2 at:
#          http://license.coscl.org.cn/MulanPubL-2.0
# THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
# EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
# MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
# See the Mulan PubL v2 for more details.

from pydub import AudioSegment
from pydub.silence import split_on_silence
import main

sound = AudioSegment.from_file("./42-10-2020-10-23-平安.mp3", format="mp3")
audio_chunks = split_on_silence(sound, min_silence_len=main.MIN_SILENCE_LEN, silence_thresh=main.SILENCE_THRESH,
                                keep_silence=main.KEEP_SILENCE, seek_step=main.SEEK_STEP)
# 遍历音频片段
for i, segment in enumerate(audio_chunks):
    # 第一个
    # if i == 0:
        # 输出
        segment.set_frame_rate(16000)
        segment.export("split_one_{}.mp3".format(i), format="mp3")