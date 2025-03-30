<template>
    <ElFormItem label="内容：">
        <ElInput v-model="text" type="textarea" :rows="6" placeholder="请输入文本" />
    </ElFormItem>
    <ElFormItem label="声音：">
        <ElSelect v-model="voice">
            <ElOption v-for="item in voices" :value="item" :key="item" />
        </ElSelect>
    </ElFormItem>
    <ElFormItem label="语言：">
        <ElSelect class="lang" v-model="lang">
            <ElOption label="中文" :value="1" />
            <ElOption label="英文" :value="2" />
            <ElOption label="中英混合" :value="3" />
        </ElSelect>
        <ElCheckbox v-model="isGB" v-show="lang != 1">英式英语</ElCheckbox>
    </ElFormItem>
    <ElFormItem label="语速：">
        <ElSlider v-model="speed" :min="0.1" :max="5" :step="0.1" show-input />
    </ElFormItem>
    <ElFormItem>
        <ElButton class="full" type="primary" @click="onGenerateClick">生成</ElButton>
    </ElFormItem>
    <audio :src="url" class="full" controls autoplay></audio>
</template>

<script setup>
import { ElButton, ElCheckbox, ElFormItem, ElInput, ElMessage, ElOption, ElSelect, ElSlider } from 'element-plus';
import { ref } from 'vue';
import { getURL, post } from './request';

const voices = ["af_alloy", "af_aoede", "af_bella", "af_heart", "af_jessica", "af_kore", "af_nicole", "af_nova", "af_river", "af_sarah", "af_sky", "am_adam", "am_echo", "am_eric", "am_fenrir", "am_liam", "am_michael", "am_onyx", "am_puck", "am_santa", "bf_alice", "bf_emma", "bf_isabella", "bf_lily", "bm_daniel", "bm_fable", "bm_george", "bm_lewis", "zf_xiaobei", "zf_xiaoni", "zf_xiaoxia", "zf_xiaoxiao", "zf_xiaoyi", "zm_yunjian", "zm_yunxi", "zm_yunyang"]

const text = ref("")
const voice = ref("zm_yunxi")
const lang = ref(3)
const isGB = ref(false)
const speed = ref(1)
const url = ref("")

function onGenerateClick() {
    if (text.value == "") {
        ElMessage({ message: "请输入内容", type: "error" })
        return
    }
    url.value = ""
    post("/tts", {
        text: text.value,
        voice: voice.value,
        lang: lang.value,
        isGB: isGB.value,
        speed: speed.value
    }, () => {
        url.value = getURL("/tts/download?" + Date.now())
    })
}
</script>

<style scoped>
.lang {
    width: calc(100% - 100px);
    margin-right: 10px;
}

.full {
    width: 100%;
}
</style>