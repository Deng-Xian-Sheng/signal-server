# Copyright (c) [2023] [JinCanQi]
# [make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
# You can use this software according to the terms and conditions of the Mulan PubL v2.
# You may obtain a copy of Mulan PubL v2 at:
#          http://license.coscl.org.cn/MulanPubL-2.0
# THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
# EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
# MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
# See the Mulan PubL v2 for more details.

# 一些有用的信息：

# 请使用python实现以下需求：
# 我需要制作用于深度学习的数据集，数据集由时长为5～15s的若干音频文件组成。
# 我有一个文件夹，里面全是mp3格式的音频文件，音频是通过歌曲提取人声产生的。
# 我需要将这些音频文件剪辑成时长5～15s的音频文件，时长并不是严格的，剪辑时要确保不要把人声截断，应该在上一句与下一句的停顿处剪辑，为了确保不把人声截断，时长不是严格的，可以长一点也可以短一点，但是要尽可能的接近要求（5～15s）；数据集不应该包含无声片段。

# Sure, I can help you with that. Based on the Pydub official documentation, here are some commands and methods that might be useful for your task:
#  1. To open an audio file:
# from pydub import AudioSegment
# sound = AudioSegment.from_file("your_audio_file.mp3", format="mp3")
# 2. To extract a segment of audio from a given start time and end time:
# # Get the first 5 seconds of the audio
# first_five_seconds = sound[:5000]
#  # Get the last 15 seconds of the audio
# last_fifteen_seconds = sound[-15000:]
#  # Get a central 10 second segment starting from 30 seconds into the audio
# ten_seconds = sound[30000:40000]
# 3. To find the pause between words:
# from pydub.silence import split_on_silence
# audio_chunks = split_on_silence(sound, min_silence_len=500, silence_thresh=-50)
# 4. To export the audio to a new file:
# first_five_seconds.export("first_five_seconds.mp3", format="mp3")
# I hope this helps with your task! Let me know if you need any further assistance.

# 多次重复转换采样率可能导致音频质量不佳，转换采样率之前先检查音频采样率
# 说话人识别模型支持mp3 mp3和wav输出的结果没有明显的不同
# 说话人识别模型对于年龄差距较大的目标说话人，识别完全不准确
# 关于分割的 静默阈值 和 无声片段最小长度 需要分别使用压缩器并获取最高最低音量和检测节奏 然后根据获得的数据设置这些值

from pydub import AudioSegment
from pydub.silence import split_on_silence
from pydub.effects import compress_dynamic_range
import librosa
import threading

# AUDIO_FORMAT 音频文件格式
AUDIO_FORMAT = "mp3"

# 要求时长 ms
MIN_DURATION = 5 * 1000
MAX_DURATION = 15 * 1000

# 严格时长约束 ms 大于30s不利于后续处理：说话人识别，普遍不应该大于30s
STRICT_MAX_DURATION = 30 * 1000

# SILENCE_THRESH 静默阈值 dBFS 程序会动态赋值
SILENCE_THRESH = -20

# MIN_SILENCE_LEN 无声片段最小长度 ms 1000ms为1s 程序会动态赋值
MIN_SILENCE_LEN = 1000

# KEEP_SILENCE 分割后，是否在片段的开头和结尾保留一些低于静默阈值的声音，以防止声音突然截断，这可以使声音更自然，可选值 bool || int
# false为不保留，true为全部保留，int为保留的时间 ms，默认500ms
KEEP_SILENCE = 500

# SEEK_STEP 在检测到静默后从非静默点向前搜索的持续时间。这在音频中有少量噪音或音损的情况下很有用。较大的 SEEK_STEP 可能有助于缓解此类噪音
# 默认 1ms
SEEK_STEP = 1

# 淡入淡出时间 ms
FADE_TIME = 500

# 压缩器参数
compress_dynamic_range_params = {
    "threshold": -10,
    "ratio": 2,
    "attack": 30,
    "release": 300
}


# 音频类 用于将音频流变成网络流
class AudioThread(threading.Thread):
    # 输入参数是音频流
    def __init__(self, audio_stream: bytes):
        threading.Thread.__init__(self)
        self.audio_stream = audio_stream

    # 重写run()方法
    def run(self):
        pass


# split_segment 通过淡入淡出将某个音频片段再次分割
# param segment pydub.AudioSegment 音频片段
# return list<pydub.AudioSegment> 分割后的音频片段
def split_segment(segment: AudioSegment) -> list:
    # 获取总时长 ms
    duration = len(segment)

    # 片段数 = 总时长 / MAX_DURATION 返回浮点
    block_number = duration / MAX_DURATION

    # 每个的时长 = 总时长 / 片段树 此处向下取整，余数的时长就被抛弃了，所以后面我们判断是最后一个片段则该片段一直持续到最后
    duration_one = int(duration / block_number)

    # 分割后的音频片段
    segments = []

    # 遍历分割
    for i in range(0, int(block_number)):
        # 起始位置
        start = i * duration_one

        # 结束位置
        if i == block_number - 1:
            # 最后一个片段
            end = duration
        else:
            # 非最后一个片段
            end = start + duration_one

        # 淡入淡出
        segment_fade = segment[start:end].fade_in(FADE_TIME).fade_out(FADE_TIME)

        # 添加到分割后的音频片段
        segments.append(segment_fade)

    return segments


# merge_segment 合并两个音频片段
# param segment1 pydub.AudioSegment 音频片段1
# param segment2 pydub.AudioSegment 音频片段2
# return pydub.AudioSegment 合并后的音频片段
def merge_segment(segment1: AudioSegment, segment2: AudioSegment) -> AudioSegment:
    return segment1 + segment2


# process 处理
# param file_path str 文件路径
def process(file_path: str):
    sound = AudioSegment.from_file(file_path, format=AUDIO_FORMAT)

    sound = compress_dynamic_range(sound, **compress_dynamic_range_params)
    # 峰值的20%作为静默阈值
    SILENCE_THRESH = sound.dBFS * 0.2

    y_librosa, sr_librosa = librosa.load(sound.get_array_of_samples(), sr=sound.frame_rate)
    tempo, _ = librosa.beat.beat_track(y=y_librosa, sr=sr_librosa)
    # 速度 105 的人声 MIN_SILENCE_LEN = 571
    # 速度 126 的人声 MIN_SILENCE_LEN = 476
    # 速度 200 的人声 MIN_SILENCE_LEN = 300
    MIN_SILENCE_LEN = int((60 / tempo) * 1000)

    audio_chunks = split_on_silence(sound, min_silence_len=MIN_SILENCE_LEN, silence_thresh=SILENCE_THRESH,
                                    keep_silence=KEEP_SILENCE, seek_step=SEEK_STEP)

    # 遍历音频片段
    i = 0
    while i < len(audio_chunks):

        # 时长 ms
        duration = len(audio_chunks[i])

        # 严格时长约束
        if duration > STRICT_MAX_DURATION:

            # 抹除当前值
            in_split_segment = audio_chunks.pop(audio_chunks[i])

            # 使用 split_segment() 再次分割，使用deque的方法抹除当前值并把分割后的值追加到当前值的下一个值
            out_split_segment = split_segment(in_split_segment)

            for vv, segment2 in enumerate(out_split_segment):
                audio_chunks.insert(i + vv, segment2)

        if duration < MIN_DURATION:

            # 抹除当前值
            in_merge_segment = audio_chunks.pop(audio_chunks[i])

            # 抹除下一个值
            if i < len(audio_chunks):

                in2_merge_segment = audio_chunks.pop(audio_chunks[i])

                # 将当前和下一个合并
                out_merge_segment = merge_segment(in_merge_segment, in2_merge_segment)

                # 将合并后的值插入到当前值的位置
                audio_chunks.insert(i, out_merge_segment)


# 创建子线程
mythread = AudioThread()

# 启动子线程
mythread.start()

#     TODO

# 等待子线程结束
mythread.join()
