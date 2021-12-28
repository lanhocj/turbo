import axios from "axios";

export let api = axios.create({
    baseURL: "/api",
    withCredentials: true,
    contentType: 'application/json; charset=utf-8',
    transformRequest: [(data, headers) => {
        let result = Object.create({})

        data.forEach((value,key) => {
            result[key] = value;
        })

        return JSON.stringify(result)
    }]
})