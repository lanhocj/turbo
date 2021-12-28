


/**
 * rule = {
 *     [name] = {
 *         test: /{regex}/,
 *         error_msg: "hello??"
 *     }
 * }
 * @param rule Object
 * @param data FormData
 */
export const validate = (rule = {}, data) => {
    return new Promise((resolve, reject) => {
        data.forEach((value, key) => {
            if (rule.hasOwnProperty(key)) {
                let item = rule[key] | {}

                if (item.test && !(item.test.test(value))) {
                    return reject(item.error_msg || "表单错误")
                }
            }
        })

        return resolve()
    })
}