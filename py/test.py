# Copyright (c) [2023] [JinCanQi]
# [make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
# You can use this software according to the terms and conditions of the Mulan PubL v2.
# You may obtain a copy of Mulan PubL v2 at:
#          http://license.coscl.org.cn/MulanPubL-2.0
# THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
# EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
# MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
# See the Mulan PubL v2 for more details.

import pyaudio
from flask import Flask, Response

app = Flask(__name__)

def generate_audio():
    CHUNK = 1024
    FORMAT = pyaudio.paInt24
    CHANNELS = 1
    RATE = 44100

    p = pyaudio.PyAudio()
    stream = p.open(format=FORMAT, channels=CHANNELS, rate=RATE, output=True)

    with open('42-10-2020-10-23-平安.mp3', 'rb') as f:
        data = f.read(CHUNK)
        while data:
            yield data
            data = f.read(CHUNK)
            if data is None:
                data = f.read(CHUNK)
    #         判断回车键 则将data设置为None
            if input() == '':
                data = None


    stream.stop_stream()
    stream.close()
    p.terminate()

@app.route('/audio')
def audio():
    return Response(generate_audio(), mimetype='audio/mp3')

if __name__ == '__main__':
    app.run()