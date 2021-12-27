import Vue from "vue";
import PopupAddNode from "./extends/popup-add-node";

let PopupConstructor = Vue.extend(PopupAddNode)

export const popup = (obj = {}) => {
    let object = Object.assign({
        title: "提示",
        content: "确认吗？",
        showCancelButton: false,
        confirmText: "确定",
        cancelText: "取消"
    }, obj)

    let dom = new PopupConstructor({
        el: document.createElement("div")
    })

    document.body.appendChild(dom.$el)

    Object.keys(object).forEach(key => {
        if (dom.hasOwnProperty(key)) {
            dom[key] = object[key]
        }
    })
}