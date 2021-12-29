import camelCase from "lodash/camelCase";

let requireComponent = require.context('./extends', false, /\w+\.(vue|js)$/)

export const VueExtendLoader = {
    install(Vue, options) {
        requireComponent.keys().forEach(fileName => {
            let component = requireComponent(fileName)

            const name = camelCase(
                // 获取和目录深度无关的文件名
                fileName
                    .split('/')
                    .pop()
                    .replace(/\.\w+$/, '')
            )

            let Constructor = Vue.extend(component.default || component)

            Vue.prototype['$' + name] = (obj = {}) => {
                let dom = new Constructor({
                    el: document.createElement("dev")
                })

                Object.keys(obj).forEach(key => {
                    dom[key] = obj[key]
                })

                document.body.appendChild(dom.$el)
                return dom;
            }
        })
    }
}
