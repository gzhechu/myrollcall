<script setup>
import { reactive, onMounted } from 'vue'
import { GetNameList } from '../../wailsjs/go/main/App'

const data = reactive({
    elements: [],
    selectedName: null, // 用于保存选中的元素
    selectedId: null,
    isRolling: false,
    timerId: null
})

const loadNames = async () => {
    try {
        const nameList = await GetNameList("yourArgument"); // 替换为实际参数
        // console.log(nameList); // 这是一个包含 main.Person 的数组
        // 你可以直接使用 nameList 中的内容
        data.elements = nameList; // 假设你想将其赋值给 data.elements
    } catch (error) {
        console.error("获取名字列表时出错:", error);
    }
};

const rollcall = () => {
    if (data.isRolling) return; // 防止重复点击
    loadNames()
    if (!data || !Array.isArray(data.elements) || data.elements.length <= 0) {
        console.error("数据错误，请先加载名字列表");
        return;
    }
    // console.log(data.elements.length)
    // console.log(data.elements)

    data.isRolling = true;
    let count = 30;

    // 清除之前的定时器
    if (data.timerId) {
        clearTimeout(data.timerId);
    }

    const roll = () => {
        let rdmIdx = Math.floor(Math.random() * data.elements.length);
        let element = data.elements[rdmIdx];
        if (element && element.name) {
            console.log(rdmIdx, element.name, element.id);
            data.selectedName = element.name;
            data.selectedId = element.id;
        }
        count--;

        // 模拟滚动延迟
        data.timerId = setTimeout(() => {
            if (count <= 0) {
                clearTimeout(data.timerId);
                data.isRolling = false;
            } else {
                roll(); // 继续滚动
            }
        }, 50); // 滚动间隔时间
    };
    roll();
};

</script>

<template>
    <main>
        <div id="calledname" class="calledname">{{ data.selectedName }}</div>
        <div id="calledid" class="calledid">{{ data.selectedId }}</div>
        <div id="input" class="input-box">
            <button class="btn" @click="rollcall">点名</button>
        </div>
    </main>
</template>

<style scoped>
.calledname {
    height: 360px;
    line-height: 360px;
    margin: 1.5rem auto;
    font-size: 280px;
}

.calledid {
    height: 72px;
    line-height: 72px;
    margin: 1.5rem auto;
    font-size: 64px;
}


.input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
}

.input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
}
</style>
