import Vue from "vue";
import camelCase from "lodash/camelCase";
import upperFirst from "lodash/upperFirst";
import {modalAddNode, modalAddUser, VueExtendLoader} from "./extends";
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

Vue.use(VueExtendLoader)

console.log(Vue.prototype)

Vue.prototype.$api = api
Vue.prototype.$validate = validate

let app = new Vue({
    el: "#main",
    data() {
        return {
            data: "hello"
        }
    },
    methods: {
        modalAddNode() {
            this.$modalAddNode().open()
        }
    }
});

console.log(app)