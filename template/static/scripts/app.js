import Vue from "vue";
import {upperFirst, camelCase} from "lodash";
import {popup} from "./extends";
import {api} from "./axios";
import {validate} from "./validate";

const requireComponent = require.context(
    // 其组件目录的相对路径
    './components',
    // 是否查询其子目录
    false,
    // 匹配基础组件文件名的正则表达式
    /\w+\.(vue|js)$/
)

requireComponent.keys().forEach(fileName => {
    const componentConfig = requireComponent(fileName)

    const componentName = upperFirst(
        camelCase(
            // 获取和目录深度无关的文件名
            fileName
                .split('/')
                .pop()
                .replace(/\.\w+$/, '')
        )
    )
    Vue.component(componentName, componentConfig.default || componentConfig)
})

Vue.prototype.$popup = popup
Vue.prototype.$api = api
Vue.prototype.$validate = validate

new Vue({
    el: "#main",
    data() {
        return {
            data: "hello"
        }
    },
    methods: {
        addNodePopup() {
            this.$popup()
        }
    }
});
