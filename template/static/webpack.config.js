const {VueLoaderPlugin} = require("vue-loader")

const isProd = process.env.NODE_ENV === 'production'

module.exports = {
    mode: isProd ? "production": "development",
    target: "web",
    output: {
        filename: "app.js"
    },
    module: {
        rules: [
            {
                test: /\.(sc|c)ss/,
                use: [ 'vue-style-loader', 'css-loader', 'sass-loader' ]
            },
            {
                test: /\.js/,
                loader: 'babel-loader'
            },
            {
                test: /\.vue/,
                loader: 'vue-loader'
            }
        ]
    },
    plugins: [
        new VueLoaderPlugin()
    ],
    resolve: {
        alias: {
            vue$: "vue/dist/vue.esm.js",
        },
        extensions: ["*", ".js", ".vue", ".json"],
    },
};